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
	histoPath := filepath.Join(outputDirectory, fmt.Sprintf("%s.histo", name))
	out, err := os.Create(histoPath)
	if err != nil {
		return fmt.Errorf("Error while report: %v", err)
	}
	reporter.Report(out)
	log.Printf("Wrote load test histogram: %s\n", histoPath)
	return nil
}
