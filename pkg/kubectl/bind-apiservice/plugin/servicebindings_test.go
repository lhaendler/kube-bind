/*
Copyright 2022 The Kube Bind Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import (
	"bytes"
	"os"
	"testing"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	kubebindv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

func TestHumanReadablePromt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		testData       kubebindv1alpha1.PermissionClaim
		expectedOutput string
	}{
		{"Test Provider, Required, Create.Donate, OnConflict.ProviderOverwrittes, ", //TODO
			kubebindv1alpha1.PermissionClaim{
				GroupResource: kubebindv1alpha1.GroupResource{
					Group:    "",
					Resource: "secret",
				},
				Version: "v1",
				Selector: kubebindv1alpha1.ResourceSelector{
					Name:  "mangodb-secret",
					Owner: kubebindv1alpha1.Provider,
				},
				Required: true,
				Create: &kubebindv1alpha1.CreateOptions{
					Donate: true,
				},
				OnConflict: &kubebindv1alpha1.OnConflictOptions{
					ProviderOverwrites:              true,
					RecreateWhenConsumerSideDeleted: true,
				},
				Update: &kubebindv1alpha1.UpdateOptions{
					Fields: []string{"stringData.credentials", "stringData.config"},
				},
			},
			"Create and modify the following user-owned secret objects (apiVersion: \"v1\") on your cluster referenced with:\n" +
				"	name: \"mangodb-secret\"\n" +
				"Conflicting objects will be overwritten and created objects will be recreated upon deletion.\n" +
				"The following fields of the objects will still be managed by the provider:\n" +
				"	\"stringData.credentials\"\n" +
				"	\"stringData.config\"\n" +
				"Accepting this Permission is required in order to proceed.\n" +
				"Do you accept this Permission? [No,Yes]\n",
		},
		{"Test 1",
			kubebindv1alpha1.PermissionClaim{
				GroupResource: kubebindv1alpha1.GroupResource{
					Group:    "mangodb",
					Resource: "customConfig",
				},
				Version: "v1alpha1",
				Selector: kubebindv1alpha1.ResourceSelector{
					Name:  "mangodb-CC",
					Owner: kubebindv1alpha1.Consumer,
				},
				Required: false,
				Adopt:    true,
				OnConflict: &kubebindv1alpha1.OnConflictOptions{
					ProviderOverwrites:              false,
					RecreateWhenConsumerSideDeleted: false,
				},
				Update: &kubebindv1alpha1.UpdateOptions{
					Preserving:     []string{"config.ID", "spec.data.dashboardToken"},
					AlwaysRecreate: true,
				},
			},
			"The provider wants to become owner of the following user-created customConfig objects (apiVersion: \"mangodb/v1alpha1\") referenced with:\n" +
				"	name: \"mangodb-CC\"\n" +
				"The provider will be able to access, modify and delete said objects on your cluster.\n" +
				"The following fields of the objects will still be managed by the user:\n" +
				"	\"config.ID\"\n" +
				"	\"spec.data.dashboardToken\"\n" +
				"Modification of said objects will by handled by deletion and recreation of said objects.\n" +
				"Accepting this Permission is optional.\n" +
				"Do you accept this Permission? [No,Yes]\n",
		},
		{"Test 2",
			kubebindv1alpha1.PermissionClaim{
				GroupResource: kubebindv1alpha1.GroupResource{
					Group:    "mangodb",
					Resource: "customConfig",
				},
				Version: "v1alpha1",
				Selector: kubebindv1alpha1.ResourceSelector{
					Owner: kubebindv1alpha1.Consumer,
				},
				Required: true,
				Adopt:    false,
			},
			"The provider wants read access to all user-created customConfig objects (apiVersion: \"mangodb/v1alpha1\").\n" +
				"The provider will not be able to modify or delete said objects.\n" +
				"Accepting this Permission is required in order to proceed.\n" +
				"Do you accept this Permission? [No,Yes]\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var output bytes.Buffer
			var input bytes.Buffer
			input.WriteString("y\n")
			opts := NewBindAPIServiceOptions(genericclioptions.IOStreams{In: &input, Out: &output, ErrOut: os.Stderr})
			b, err := opts.promptYesNo(tt.testData)
			if output.String() != tt.expectedOutput {
				t.Errorf("Expected IO Output did not match. got: \"\n%s\"\nwanted: \"\n%s\"\n", output.String(), tt.expectedOutput)
			}
			if b == false || (err != nil) {
				t.Errorf("Expected Return value did not match. got: \"%v\", \"%v\"", b, err)
			}
		})
	}
}
