package fluxmonitorv2

import (
	"time"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	coreorm "github.com/smartcontractkit/chainlink/core/store/orm"
)

func ValidatedFluxMonitorSpec(config *coreorm.Config, ts string) (job.Job, error) {
	var j = job.Job{
		Pipeline: *pipeline.NewTaskDAG(),
	}
	var spec job.FluxMonitorSpec
	tree, err := toml.Load(ts)
	if err != nil {
		return j, err
	}
	err = tree.Unmarshal(&j)
	if err != nil {
		return j, err
	}
	err = tree.Unmarshal(&spec)
	if err != nil {
		return j, err
	}
	j.FluxMonitorSpec = &spec

	if j.Type != job.FluxMonitor {
		return j, errors.Errorf("unsupported type %s", j.Type)
	}
	if j.SchemaVersion != uint32(1) {
		return j, errors.Errorf("the only supported schema version is currently 1, got %v", j.SchemaVersion)
	}

	// Find the smallest of all the timeouts
	// and ensure the polling period is greater than that.
	minTaskTimeout, aTimeoutSet, err := j.Pipeline.MinTimeout()
	if err != nil {
		return j, err
	}
	timeouts := []time.Duration{
		config.DefaultHTTPTimeout().Duration(),
		time.Duration(j.MaxTaskDuration),
	}
	if aTimeoutSet {
		timeouts = append(timeouts, minTaskTimeout)
	}
	var minTimeout time.Duration = 1<<63 - 1
	for _, timeout := range timeouts {
		if timeout < minTimeout {
			minTimeout = timeout
		}
	}

	if !validatePollTimer(spec.PollTimerDisabled, minTimeout, spec.PollTimerPeriod) {
		return j, errors.Errorf("pollTimer.period must be equal or greater than %v, got %v", minTimeout, spec.PollTimerPeriod)
	}

	if !validateJitter(spec.PollJitter, spec.PollTimerPeriod) {
		return j, errors.Errorf("PollJitter must be less than or equal to PollTimerPeriod in seconds")
	}

	return j, nil
}

// validateJitter validates the jitter divisor is not greater than the polling
// interval in seconds
func validateJitter(pollJitter int32, pollTimerPeriod time.Duration) bool {
	if pollJitter > int32(pollTimerPeriod.Seconds()) {
		return false
	}

	return true
}

// validatePollTime validates the period is greater than the min timeout for an
// enabled poll timer.
func validatePollTimer(disabled bool, minTimeout time.Duration, period time.Duration) bool {
	// Disabled timers do not need to validate the period
	if disabled {
		return true
	}

	return period >= minTimeout
}
