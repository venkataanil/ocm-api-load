package report

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Write(name, outputDirectory string, metrics *vegeta.Metrics) error {
	reporter := vegeta.NewJSONReporter(metrics)
	path := filepath.Join(outputDirectory, fmt.Sprintf("%s.json", name))
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Error while report: %v", err)
	}
	if err := reporter.Report(out); err != nil {
		return err
	}
	log.Printf("Wrote load test report '%s' to %s\n", name, path)
	return nil
}
