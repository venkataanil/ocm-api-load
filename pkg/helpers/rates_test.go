package helpers

import (
	"reflect"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestParseRate(t *testing.T) {
	tests := []struct {
		name        string
		rate        string
		connections int
		wantF       int
		wantD       string
		wantErr     bool
	}{
		{"1", "1/s", 1, 1, "1s", false},
		{"2", "infinity", 1, 0, "", false},
		{"3", "10", 1, 10, "1s", false},
		{"4", "0", 1, 0, "", false},
		{"5", "1/m", 1, 1, "1m", false},
		{"6", "1/h", 1, 1, "1h", false},
		{"7", "1/ms", 1, 1, "1ms", false},
		{"8", "1/ns", 1, 1, "1ns", false},
		{"9", "1/us", 1, 1, "1us", false},
		{"10", "1/µs", 1, 1, "1µs", false},
		{"11", "500/s", 1, 500, "1s", false},
		{"12", "1/t", 1, 0, "", true},
		{"13", "fast", 1, 0, "", true},
		{"1_withConnections", "1/s", 2, 1, "1s", false},
		{"2_withConnections", "infinity", 2, 0, "", false},
		{"3_withConnections", "10", 2, 5, "1s", false},
		{"4_withConnections", "0", 2, 0, "", false},
		{"5_withConnections", "1/m", 2, 1, "1m", false},
		{"6_withConnections", "1/h", 2, 1, "1h", false},
		{"7_withConnections", "1/ms", 3, 1, "1ms", false},
		{"8_withConnections", "1/ns", 10, 1, "1ns", false},
		{"9_withConnections", "1/us", 2, 1, "1us", false},
		{"10_withConnections", "1/µs", 2, 1, "1µs", false},
		{"11_withConnections", "500/s", 4, 125, "1s", false},
		{"12_withConnections", "1/t", 1, 0, "", true},
		{"13_withConnections", "fast", 1, 0, "", true},
		{"14_withConnections", "10", 3, 3, "1s", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRate(tt.rate, tt.connections)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			d, err := time.ParseDuration(tt.wantD)
			if (err != nil) != tt.wantErr && tt.wantD != "" {
				t.Errorf("ParseRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := vegeta.ConstantPacer{Freq: tt.wantF, Per: d}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("ParseRate() = %v, want %v", got, want)
			}
		})
	}
}
