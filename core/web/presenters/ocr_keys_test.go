package presenters

// TODO - RYAN

// func TestOCRKeysBundleResource(t *testing.T) {
// 	timestamp := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

// 	var (
// 		ocrKeyBundleID = "f126a0bd09f6a9e2d8cd4d2fcb623cb094f6e487c8a22dfb061911d64b770026"
// 		password       = "p4SsW0rD1!@#_"
// 	)

// 	ocrKeyBundleIDSha256, err := models.Sha256HashFromHex(ocrKeyBundleID)
// 	require.NoError(t, err)

// 	pk, err := ocrkey.NewV2()
// 	require.NoError(t, err)

// 	bundle := ocrkey.EncryptedKeyBundle{
// 		ID:                    ocrKeyBundleIDSha256,
// 		OnChainSigningAddress: pkEncrypted.OnChainSigningAddress,
// 		OffChainPublicKey:     pkEncrypted.OffChainPublicKey,
// 		ConfigPublicKey:       pkEncrypted.ConfigPublicKey,
// 		CreatedAt:             timestamp,
// 		UpdatedAt:             timestamp,
// 	}

// 	r := NewOCRKeysBundleResource(bundle)
// 	b, err := jsonapi.Marshal(r)
// 	require.NoError(t, err)

// 	expected := fmt.Sprintf(`
// 	{
// 		"data":{
// 			"type":"encryptedKeyBundles",
// 			"id":"%s",
// 			"attributes":{
// 				"onChainSigningAddress": "%s",
// 				"offChainPublicKey": "%s",
// 				"configPublicKey": "%s",
// 				"createdAt":"2000-01-01T00:00:00Z",
// 				"updatedAt":"2000-01-01T00:00:00Z",
// 				"deletedAt":null
// 			}
// 		}
// 	}`,
// 		ocrKeyBundleID,
// 		pkEncrypted.OnChainSigningAddress.String(),
// 		pkEncrypted.OffChainPublicKey.String(),
// 		pkEncrypted.ConfigPublicKey.String(),
// 	)

// 	assert.JSONEq(t, expected, string(b))

// 	// With a deleted field
// 	bundle.DeletedAt = gorm.DeletedAt(sql.NullTime{Time: timestamp, Valid: true})

// 	r = NewOCRKeysBundleResource(bundle)
// 	b, err = jsonapi.Marshal(r)
// 	require.NoError(t, err)

// 	expected = fmt.Sprintf(`
// 	{
// 		"data": {
// 			"type":"encryptedKeyBundles",
// 			"id":"%s",
// 			"attributes":{
// 				"onChainSigningAddress": "%s",
// 				"offChainPublicKey": "%s",
// 				"configPublicKey": "%s",
// 				"createdAt":"2000-01-01T00:00:00Z",
// 				"updatedAt":"2000-01-01T00:00:00Z",
// 				"deletedAt":"2000-01-01T00:00:00Z"
// 			}
// 		}
// 	}`,
// 		ocrKeyBundleID,
// 		pkEncrypted.OnChainSigningAddress.String(),
// 		pkEncrypted.OffChainPublicKey.String(),
// 		pkEncrypted.ConfigPublicKey.String(),
// 	)

// 	assert.JSONEq(t, expected, string(b))
// }
