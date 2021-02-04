package tests

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Run(
	attacker *vegeta.Attacker,
	metrics vegeta.Metrics,
	rate vegeta.Pacer,
	outputDirectory string,
	duration time.Duration) error {

	return TestListClusters(attacker,
		metrics,
		rate,
		outputDirectory,
		duration)
}
