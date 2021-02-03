package load

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"testing"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"

	. "github.com/onsi/ginkgo" // nolint
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega" // nolint

	vegeta "github.com/tsenart/vegeta/v12/lib"
	common "gitlab.cee.redhat.com/service/uhc-clusters-service/test/common"
)

var connection *sdk.Connection
var connAttacker func(*vegeta.Attacker)
var metrics vegeta.Metrics
var rate vegeta.Rate
var duration time.Duration

var testArgs struct {
	load          bool
	tokenURL      string
	gatewayURL    string
	clientID      string
	clientSecret  string
	token         string
	durationInMin int
	rate          int
}

func init() {
	flag.BoolVar(
		&testArgs.load,
		"load",
		false,
		"Run load tests.",
	)
	flag.StringVar(
		&testArgs.tokenURL,
		"token-url",
		"https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		"Token URL.",
	)
	flag.StringVar(
		&testArgs.gatewayURL,
		"gateway-url",
		"http://localhost:8001",
		"Gateway URL.",
	)
	flag.StringVar(
		&testArgs.clientID,
		"client-id",
		"cloud-services",
		"OpenID client identifier.",
	)
	flag.StringVar(
		&testArgs.clientSecret,
		"client-secret",
		"",
		"OpenID client secret.",
	)
	flag.StringVar(
		&testArgs.token,
		"token",
		"",
		"Offline token for authentication.",
	)
	flag.IntVar(
		&testArgs.durationInMin,
		"duration-in-min",
		5,
		"How long should the load test take.",
	)
	flag.IntVar(
		&testArgs.rate,
		"rate",
		100,
		"How many times (per second) should the endpoint be hit.",
	)

}

var _ = BeforeSuite(func() {

	// waitForBackendToBeReady()
})

func CreateConnection() (*sdk.Connection, error) {
	// Create a logger:
	logger, err := createLogger()
	if err != nil {
		return nil, err
	}

	// Create the connection:
	return sdk.NewConnectionBuilder().
		Logger(logger).
		Insecure(true).
		URL(testArgs.gatewayURL).
		Client(testArgs.clientID, testArgs.clientSecret).
		Tokens(testArgs.token).
		TransportWrapper(func(wrapped http.RoundTripper) http.RoundTripper {
			return &common.CleanClustersTransport{Wrapped: wrapped}
		}).
		Build()
}

func createLogger() (sdk.Logger, error) {
	return sdk.NewStdLoggerBuilder().
		Streams(GinkgoWriter, GinkgoWriter).
		Debug(false).
		Build()
}

func TestLoad(t *testing.T) {
	flag.Parse()
	common.SetupLogger()
	if testArgs.load {
		validateArgs()
		rate = vegeta.Rate{Freq: testArgs.durationInMin, Per: time.Second}
		duration = time.Duration(testArgs.durationInMin) * time.Minute
		connAttacker = vegeta.Client(&http.Client{Transport: connection})

		RegisterFailHandler(Fail)
		junitReporter := reporters.NewJUnitReporter("junit.xml")
		RunSpecsWithDefaultAndCustomReporters(t, "API Test Suite", []Reporter{junitReporter})
	}
}

func TestCreateCluster(t *testing.T) {
	body, err := v1.NewCluster().Name("load-test").MultiAZ(false).Build()
	Expect(err).NotTo(HaveOccurred())
	var raw bytes.Buffer
	err = v1.MarshalCluster(body, &raw)

	attacker := vegeta.NewAttacker(connAttacker)
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(testArgs.gatewayURL + "/api/clusters_mgmt/v1/clusters"),
		Body:   raw.Bytes(),
	})
	for res := range attacker.Attack(targeter, rate, duration, "Create") {
		metrics.Add(res)
	}
}

func validateArgs() {
	common.CheckEmpty(testArgs.gatewayURL, "gateway-url")
	common.CheckEmpty(testArgs.token, "token")
}
