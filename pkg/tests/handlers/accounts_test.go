package handlers

import (
	"context"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/cloud-bulldozer/ocm-api-load/pkg/logging"
	"github.com/cloud-bulldozer/ocm-api-load/pkg/types"
)

func Test_clusterAuthorizationsBody(t *testing.T) {
	type args struct {
		ctx       context.Context
		clusterID string
		options   *types.TestOptions
	}
	logger, _ := logging.NewGoLoggerBuilder().Build()
	tests := []struct {
		name string
		args args
	}{

		{"randomize-1", args{context.TODO(), "clusterID-1", &types.TestOptions{Logger: logger}}},
		{"randomize-2", args{context.TODO(), "clusterID-2", &types.TestOptions{Logger: logger}}},
		{"randomize-3", args{context.TODO(), "clusterID-3", &types.TestOptions{Logger: logger}}},
		{"randomize-4", args{context.TODO(), "clusterID-4", &types.TestOptions{Logger: logger}}},
		{"randomize-5", args{context.TODO(), "clusterID-5", &types.TestOptions{Logger: logger}}},
		{"randomize-6", args{context.TODO(), "clusterID-6", &types.TestOptions{Logger: logger}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clusterAuthorizationsBody(tt.args.ctx, tt.args.clusterID, tt.args.options)
			if got == nil {
				t.Errorf("clusterAuthorizationsBody() = %v", got)
			} else {
				if !strings.Contains(string(got), tt.args.clusterID) {
					t.Errorf("Clusterid not in body = %v, Expected = %v", string(got), tt.args.clusterID)
				}
			}
		})
	}
}

func Test_randomizeCount(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name     string
		mincount int
		maxcount int
	}{
		{"Testing range of random count generated-1", 4, 10},
		{"Testing range of random count generated-2", 4, 10},
		{"Testing range of random count generated-3", 4, 10},
		{"Testing range of random count generated-4", 4, 10},
		{"Testing range of random count generated-5", 4, 10},
		{"Testing range of random count generated-6", 4, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomizeCount()
			if !(got >= tt.mincount && got <= tt.maxcount) {
				t.Errorf("randomizeCount() = %v, Expected count to be within range %v and %v", got, tt.mincount, tt.maxcount)
			}
		})
	}
}

func Test_randomizeResourceName(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	resourcenames := [5]string{"m3.2xlarge", "m4.large", "m4.xlarge", "m5.large", "m5.xlarge"}
	tests := []struct {
		name          string
		resourcenames [5]string
	}{
		{"Testing to check if output returned is from AWSResourcename-1", resourcenames},
		{"Testing to check if output returned is from AWSResourcename-2", resourcenames},
		{"Testing to check if output returned is from AWSResourcename-3", resourcenames},
		{"Testing to check if output returned is from AWSResourcename-4", resourcenames},
		{"Testing to check if output returned is from AWSResourcename-5", resourcenames},
		{"Testing to check if output returned is from AWSResourcename-6", resourcenames},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomizeResourceName()
			result := false
			for _, resourcename := range tt.resourcenames {
				if resourcename == got {
					result = true
					break
				}
			}
			if !result {
				t.Errorf("randomizeResourceName() = %v, Expected from list %v ", got, resourcenames)
			}
		})
	}
}
