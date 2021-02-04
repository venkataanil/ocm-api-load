package tests

import (
	"bytes"
	"net/http"
	"time"

	"github.com/nimrodshn/cs-load-test/pkg/helpers"
	"github.com/nimrodshn/cs-load-test/pkg/report"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestListClusters(attacker *vegeta.Attacker,
	metrics vegeta.Metrics,
	rate vegeta.Pacer,
	outputDirectory string,
	duration time.Duration) error {

	fakeClusterProps := map[string]string{
		"fake_cluster": "true",
	}
	body, err := v1.NewCluster().
		Name("load-test").
		Properties(fakeClusterProps).
		MultiAZ(false).Build()
	if err != nil {
		return err
	}
	var raw bytes.Buffer
	err = v1.MarshalCluster(body, &raw)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodGet,
		URL:    helpers.ClustersEndpoint,
		Body:   nil,
	})
	for res := range attacker.Attack(targeter, rate, duration, "Create") {
		metrics.Add(res)
	}
	metrics.Close()

	return report.Write("create-cluster-replort",
		outputDirectory,
		&metrics)
}

func TestCreateCluster(attacker *vegeta.Attacker,
	metrics vegeta.Metrics,
	rate vegeta.Pacer,
	outputDirectory string,
	duration time.Duration) error {

	fakeClusterProps := map[string]string{
		"fake_cluster": "true",
	}
	body, err := v1.NewCluster().
		Name("load-test").
		Properties(fakeClusterProps).
		MultiAZ(false).Build()
	if err != nil {
		return err
	}
	var raw bytes.Buffer
	err = v1.MarshalCluster(body, &raw)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodPost,
		URL:    helpers.ClustersEndpoint,
		Body:   raw.Bytes(),
	})
	for res := range attacker.Attack(targeter, rate, duration, "Create") {
		metrics.Add(res)
	}
	metrics.Close()

	return report.Write("create-cluster-replort",
		outputDirectory,
		&metrics)
}
