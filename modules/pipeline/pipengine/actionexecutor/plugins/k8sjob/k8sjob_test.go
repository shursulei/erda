// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sjob

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/providers/clusterinfo"
	"github.com/erda-project/erda/pkg/k8sclient"
)

func Test_isBuildkitHit(t *testing.T) {
	utCount := 10000

	tests := []struct {
		Rate   int
		Result int
	}{
		{
			Rate: 0,
		},
		{
			Rate: 10,
		},
		{
			Rate: 50,
		},
		{
			Rate: 80,
		},
		{
			Rate: 100,
		},
	}

	for _, test := range tests {
		for i := 0; i < utCount; i++ {
			if isRateHit(test.Rate) {
				test.Result++
			}
		}
		t.Logf("rate: %d, result: %v", test.Rate, test.Result*100/utCount)
	}
}

func Test_generateKubeJob(t *testing.T) {
	defer monkey.UnpatchAll()

	monkey.Patch(k8sclient.NewWithTimeOut, func(_ string, _ time.Duration) (*k8sclient.K8sClient, error) {
		return nil, nil
	})

	monkey.Patch(clusterinfo.GetClusterInfoByName, func(clusterName string) (apistructs.ClusterInfo, error) {
		return apistructs.ClusterInfo{CM: apistructs.ClusterInfoData{
			"BUILDKIT_ENABLE": "false",
		}}, nil
	})

	j, err := New("fake-job", "fake-cluster", apistructs.ClusterInfo{})
	assert.NoError(t, err)
	assert.Equal(t, string(j.Name()), "fake-job")

	_, err = j.generateKubeJob(apistructs.JobFromUser{
		Name:      "fake-job",
		Namespace: metav1.NamespaceDefault,
	}, nil)
	assert.NoError(t, err)
}

func TestCheckLabels(t *testing.T) {
	type args struct {
		source map[string]string
		target map[string]string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				source: map[string]string{
					"alibabacloud.com/eci": "true",
				},
				target: map[string]string{},
			},
			want: false,
		},
		{
			name: "test02",
			args: args{
				source: map[string]string{},
				target: map[string]string{
					"alibabacloud.com/eci": "true",
				},
			},
			want: false,
		},
		{
			name: "test03",
			args: args{
				source: map[string]string{
					"alibabacloud.com/eci": "true",
				},
				target: map[string]string{
					"alibabacloud.com/eci": "true",
					"erda.cloud/csi":       "true",
				},
			},
			want: true,
		},
		{
			name: "test04",
			args: args{
				source: map[string]string{},
				target: map[string]string{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkLabels(tt.args.source, tt.args.target)
			if got != tt.want {
				t.Errorf("checkLabels() got = %v, want %v", got, tt.want)
			}
		})
	}

}
