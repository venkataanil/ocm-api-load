package tests

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type testCase func(attacker *vegeta.Attacker,
	metrics map[string]*vegeta.Metrics,
	rate vegeta.Pacer,
	outputDirectory string,
	duration time.Duration) error

func Run(
	attacker *vegeta.Attacker,
	metrics map[string]*vegeta.Metrics,
	rate vegeta.Pacer,
	outputDirectory string,
	duration time.Duration) error {

	testCases := []testCase{TestCreateCluster, TestListClusters}

	for _, testCase := range testCases {
		err := testCase(attacker,
			metrics,
			rate,
			outputDirectory,
			duration)
		if err != nil {
			return err
		}
	}
	return nil
}
