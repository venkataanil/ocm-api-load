package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/nimrodshn/cs-load-test/pkg/helpers"
	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"gitlab.cee.redhat.com/service/uhc-clusters-service/test/api"
)

var connection *sdk.Connection
var connAttacker func(*vegeta.Attacker)
var metrics vegeta.Metrics
var rate vegeta.Rate
var duration time.Duration

var args struct {
	load            bool
	tokenURL        string
	gatewayURL      string
	clientID        string
	clientSecret    string
	token           string
	durationInMin   int
	rate            int
	outputDirectory string
}

func init() {
	flag.BoolVar(
		&args.load,
		"load",
		false,
		"Run load tests.",
	)
	flag.StringVar(
		&args.tokenURL,
		"token-url",
		"https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		"Token URL.",
	)
	flag.StringVar(
		&args.gatewayURL,
		"gateway-url",
		"http://localhost:8000",
		"Gateway URL.",
	)
	flag.StringVar(
		&args.clientID,
		"client-id",
		"cloud-services",
		"OpenID client identifier.",
	)
	flag.StringVar(
		&args.clientSecret,
		"client-secret",
		"",
		"OpenID client secret.",
	)
	flag.StringVar(
		&args.token,
		"token",
		"",
		"Offline token for authentication.",
	)
	flag.IntVar(
		&args.durationInMin,
		"duration-in-min",
		1,
		"How long should the load test take.",
	)
	flag.IntVar(
		&args.rate,
		"rate",
		100,
		"How many times (per second) should the endpoint be hit.",
	)
	flag.StringVar(
		&args.outputDirectory,
		"output-path",
		"",
		"path to output results.",
	)
}

func main() {
	flag.Parse()
	connection, err := sdk.NewConnectionBuilder().
		Insecure(true).
		URL(args.gatewayURL).
		Client(args.clientID, args.clientSecret).
		Tokens(args.token).
		TransportWrapper(func(wrapped http.RoundTripper) http.RoundTripper {
			return &helpers.CleanClustersTransport{Wrapped: wrapped}
		}).
		Build()
	if err != nil {
		fmt.Printf("Error creating api connection: %v", err)
		os.Exit(1)
	}
	rate = vegeta.Rate{Freq: args.durationInMin, Per: time.Second}
	duration = time.Duration(args.durationInMin) * time.Minute
	connAttacker = vegeta.Client(&http.Client{Transport: connection})

	if err := TestCreateCluster(); err != nil {
		fmt.Printf("Error running create cluster load test: %v", err)
		os.Exit(1)
	}
}

func TestCreateCluster() error {
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

	attacker := vegeta.NewAttacker(connAttacker)
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodGet,
		URL:    args.gatewayURL + helpers.ClustersEndpoint,
	})
	for res := range attacker.Attack(targeter, rate, duration, "Create") {
		metrics.Add(res)
	}
	reporter := vegeta.NewHDRHistogramPlotReporter(&metrics)
	histoPath := filepath.Join(args.outputDirectory, fmt.Sprintf("%s.histo", "create-clusters-report"))
	out, err := os.Create(histoPath)
	if err != nil {
		return fmt.Errorf("error while report: %v", err)
	}
	reporter.Report(out)
	log.Printf("Wrote load test histogram: %s\n", histoPath)
	return nil
}

func validateArgs() {
	api.CheckEmpty(args.gatewayURL, "gateway-url")
	api.CheckEmpty(args.token, "token")
}
