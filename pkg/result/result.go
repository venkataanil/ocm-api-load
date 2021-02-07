package result

import (
	"os"

	vegeta "github.com/tsenart/vegeta/v12/lib"
	"k8s.io/apimachinery/pkg/util/json"
)

func Write(result *vegeta.Result, file *os.File) error {
	json.NewEncoder(file).Encode(result)
	return nil
}
