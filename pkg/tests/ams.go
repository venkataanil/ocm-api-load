package tests

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nimrodshn/cs-load-test/pkg/helpers"
	"github.com/nimrodshn/cs-load-test/pkg/result"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestSelfAccessToken(attacker *vegeta.Attacker, outputDirectory string, duration time.Duration, testID string) error {

	testName := "self-access-token"
	rate := vegeta.ConstantPacer{Freq: 1000, Per: time.Hour}
	fileName := fmt.Sprintf("%s_%s", testID, testName)

	target := vegeta.Target{
		Method: http.MethodPost,
		URL:    helpers.SelfAccessTokenEndpoint,
	}

	targeter := vegeta.NewStaticTargeter(target)
	resultFile, err := createFile(fileName, outputDirectory)
	defer resultFile.Close()
	if err != nil {
		return err
	}

	// Display some info about the test being ran to catch obvious issues
	// and include context
	fmt.Printf("Test: %s\n", testName)
	fmt.Printf("Output File: %s/%s\n", outputDirectory, fileName)

	for res := range attacker.Attack(targeter, rate, duration, testName) {
		result.Write(res, resultFile)
	}

	return nil
}

func TestListSubscriptions(attacker *vegeta.Attacker, outputDirectory string, duration time.Duration, testID string) error {

	testName := "list-subscriptions"
	rate := vegeta.ConstantPacer{Freq: 2000, Per: time.Hour}
	fileName := fmt.Sprintf("%s_%s", testID, testName)

	target := vegeta.Target{
		Method: http.MethodGet,
		URL:    helpers.ListSubscriptionsEndpoint,
	}

	targeter := vegeta.NewStaticTargeter(target)
	resultFile, err := createFile(fileName, outputDirectory)
	defer resultFile.Close()
	if err != nil {
		return err
	}

	// Display some info about the test being ran to catch obvious issues
	// and include context
	fmt.Printf("Test: %s\n", testName)
	fmt.Printf("Output File: %s/%s\n", outputDirectory, fileName)

	for res := range attacker.Attack(targeter, rate, duration, testName) {
		result.Write(res, resultFile)
	}

	return nil

}

func TestAccessReview(attacker *vegeta.Attacker, outputDirectory string, duration time.Duration, testID string) error {

	testName := "access-review"
	rate := vegeta.ConstantPacer{Freq: 100, Per: time.Second}
	fileName := fmt.Sprintf("%s_%s", testID, testName)

	target := vegeta.Target{
		Method: http.MethodPost,
		URL:    helpers.AccessReviewEndpoint,
		Body:   []byte("{'account_username': 'rhn-support-tiwillia', 'action': 'get', 'resource_type': 'Subscription'}"),
	}

	targeter := vegeta.NewStaticTargeter(target)
	resultFile, err := createFile(fileName, outputDirectory)
	defer resultFile.Close()
	if err != nil {
		return err
	}

	// Display some info about the test being ran to catch obvious issues
	// and include contextq
	fmt.Printf("Test: %s\n", testName)
	fmt.Printf("Output File: %s/%s\n", outputDirectory, fileName)

	for res := range attacker.Attack(targeter, rate, duration, testName) {
		result.Write(res, resultFile)
	}

	return nil

}

func TestRegisterNewCluster(attacker *vegeta.Attacker, outputDirectory string, duration time.Duration, testID string) error {

	testName := "new-cluster-registration"
	rate := vegeta.ConstantPacer{Freq: 1000, Per: time.Hour}
	fileName := fmt.Sprintf("%s_%s", testID, testName)

	target := vegeta.Target{
		Method: http.MethodPost,
		URL:    helpers.AccessReviewEndpoint,
		Body:   []byte("{'account_username': 'rhn-support-tiwillia', 'action': 'get', 'resource_type': 'Subscription'}"),
	}

	targeter := vegeta.NewStaticTargeter(target)
	resultFile, err := createFile(fileName, outputDirectory)
	defer resultFile.Close()
	if err != nil {
		return err
	}

	// Display some info about the test being ran to catch obvious issues
	// and include contextq
	fmt.Printf("Test: %s\n", testName)
	fmt.Printf("Output File: %s/%s\n", outputDirectory, fileName)

	for res := range attacker.Attack(targeter, rate, duration, testName) {
		result.Write(res, resultFile)
	}

	return nil

}
