package vrfkey

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/utils"
	bm "github.com/smartcontractkit/chainlink/core/utils/big_math"
	"go.dedis.ch/kyber/v3"
)

var suite = secp256k1.NewBlakeKeccackSecp256k1()

type Raw []byte

func (rawKey Raw) Key() KeyV2 {
	rawKeyInt := new(big.Int).SetBytes(rawKey)
	k := secp256k1.IntToScalar(rawKeyInt)
	key, err := keyFromScalar(k)
	if err != nil {
		panic(err)
	}
	return key
}

type KeyV2 struct {
	k         kyber.Scalar
	PublicKey secp256k1.PublicKey
}

func NewV2() (KeyV2, error) {
	k := suite.Scalar().Pick(suite.RandomStream())
	return keyFromScalar(k)
}

func MustNewV2XXXTestingOnly(k *big.Int) KeyV2 {
	rv, err := keyFromScalar(secp256k1.IntToScalar(k))
	if err != nil {
		panic(err)
	}
	return rv
}

func (key KeyV2) ID() string {
	return hexutil.Encode(key.PublicKey[:])
}

func (key KeyV2) Raw() Raw {
	return secp256k1.ToInt(key.k).Bytes()
}

// GenerateProofWithNonce allows external nonce generation for testing purposes
//
// As with signatures, using nonces which are in any way predictable to an
// adversary will leak your secret key! Most people should use GenerateProof
// instead.
func (k KeyV2) GenerateProofWithNonce(seed, nonce *big.Int) (Proof, error) {
	secretKey := secp256k1.ScalarToHash(k.k).Big()
	if !(secp256k1.RepresentsScalar(secretKey) && seed.BitLen() <= 256) {
		return Proof{}, fmt.Errorf("badly-formatted key or seed")
	}
	skAsScalar := secp256k1.IntToScalar(secretKey)
	publicKey := Secp256k1Curve.Point().Mul(skAsScalar, nil)
	h, err := HashToCurve(publicKey, seed, func(*big.Int) {})
	if err != nil {
		return Proof{}, errors.Wrap(err, "vrf.makeProof#HashToCurve")
	}
	gamma := Secp256k1Curve.Point().Mul(skAsScalar, h)
	sm := secp256k1.IntToScalar(nonce)
	u := Secp256k1Curve.Point().Mul(sm, Generator)
	uWitness := secp256k1.EthereumAddress(u)
	v := Secp256k1Curve.Point().Mul(sm, h)
	c := ScalarFromCurvePoints(h, publicKey, gamma, uWitness, v)
	// (m - c*secretKey) % GroupOrder
	s := bm.Mod(bm.Sub(nonce, bm.Mul(c, secretKey)), secp256k1.GroupOrder)
	if e := checkCGammaNotEqualToSHash(c, gamma, s, h); e != nil {
		return Proof{}, e
	}
	outputHash := utils.MustHash(string(append(RandomOutputHashPrefix,
		secp256k1.LongMarshal(gamma)...)))
	rv := Proof{
		PublicKey: publicKey,
		Gamma:     gamma,
		C:         c,
		S:         s,
		Seed:      seed,
		Output:    outputHash.Big(),
	}
	valid, err := rv.VerifyVRFProof()
	if !valid || err != nil {
		panic("constructed invalid proof")
	}
	return rv, nil
}

// GenerateProof returns gamma, plus proof that gamma was constructed from seed
// as mandated from the given secretKey, with public key secretKey*Generator
//
// secretKey and seed must be less than secp256k1 group order. (Without this
// constraint on the seed, the samples and the possible public keys would
// deviate very slightly from uniform distribution.)
func (k KeyV2) GenerateProof(seed *big.Int) (Proof, error) {
	for {
		nonce, err := rand.Int(rand.Reader, secp256k1.GroupOrder)
		if err != nil {
			return Proof{}, err
		}
		proof, err := k.GenerateProofWithNonce(seed, nonce)
		switch {
		case err == ErrCGammaEqualsSHash:
			// This is cryptographically impossible, but if it were ever to happen, we
			// should try again with a different nonce.
			continue
		case err != nil: // Any other error indicates failure
			return Proof{}, err
		default:
			return proof, err // err should be nil
		}
	}
}

func keyFromScalar(k kyber.Scalar) (KeyV2, error) {
	rawPublicKey, err := secp256k1.ScalarToPublicPoint(k).MarshalBinary()
	if err != nil {
		return KeyV2{}, err
	}
	if len(rawPublicKey) != secp256k1.CompressedPublicKeyLength {
		return KeyV2{}, fmt.Errorf("public key %x has wrong length", rawPublicKey)
	}
	var publicKey secp256k1.PublicKey
	copy(publicKey[:], rawPublicKey)
	return KeyV2{
		k:         k,
		PublicKey: publicKey,
	}, nil
}