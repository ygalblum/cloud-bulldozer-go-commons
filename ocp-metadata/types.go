// Copyright 2023 The go-commons Authors.
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

package ocpmetadata

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// OCP specific constants
const (
	running            = "Running"
	completedUpdate    = "Completed"
	workerNodeSelector = "node-role.kubernetes.io/worker=,node-role.kubernetes.io/infra!=,node-role.kubernetes.io/workload!="
	monitoringNs       = "openshift-monitoring"
	tokenExpiration    = 10 * time.Hour
)

var routeGVR = schema.GroupVersionResource{
	Group:    "route.openshift.io",
	Version:  "v1",
	Resource: "routes",
}

var ingressControllerGRV = schema.GroupVersionResource{
	Group:    "operator.openshift.io",
	Version:  "v1",
	Resource: "ingresscontrollers",
}

var vmiGVR = schema.GroupVersionResource{
	Group:    "kubevirt.io",
	Version:  "v1",
	Resource: "virtualmachineinstances",
}

// infraObj
// TODO at the moment can be used to decode some AWS platform specific information from the infrastructure object
// like region and resourceTags (which is actually used to detect if this is a ROSA cluster)
// similar information is not found for other platforms like GCP or Azure.
// To collect such information we shall use a different approach, i.e:
// Using well known node labels like topology.kubernetes.io/region to get the cloud region
type infraObj struct {
	Status struct {
		InfrastructureName string `json:"infrastructureName"`
		Platform           string `json:"platform"`
		Type               string `json:"type"`
		PlatformStatus     struct {
			Aws struct {
				Region       string `json:"region"`
				ResourceTags []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"resourceTags"`
			} `json:"aws"`
			Type string `json:"type"`
		} `json:"platformStatus"`
	} `json:"status"`
}

// Type to store version info
type versionObj struct {
	ocpVersion      string
	ocpMajorVersion string
	k8sVersion      string
}

// Type to store cluster info
type clusterVersion struct {
	Status struct {
		History []struct {
			State   string `json:"state"`
			Version string `json:"version"`
		} `json:"history"`
	} `json:"status"`
}

// Type to store cluster metadata
type ClusterMetadata struct {
	MetricName       string `json:"metricName,omitempty"`
	Platform         string `json:"platform,omitempty"`
	ClusterType      string `json:"clusterType,omitempty"`
	OCPVersion       string `json:"ocpVersion,omitempty"`
	OCPMajorVersion  string `json:"ocpMajorVersion,omitempty"`
	K8SVersion       string `json:"k8sVersion,omitempty"`
	MasterNodesType  string `json:"masterNodesType,omitempty"`
	WorkerNodesType  string `json:"workerNodesType,omitempty"`
	MasterNodesCount int    `json:"masterNodesCount,omitempty"`
	InfraNodesType   string `json:"infraNodesType,omitempty"`
	WorkerNodesCount int    `json:"workerNodesCount,omitempty"`
	InfraNodesCount  int    `json:"infraNodesCount,omitempty"`
	OtherNodesCount  int    `json:"otherNodesCount,omitempty"`
	TotalNodes       int    `json:"totalNodes,omitempty"`
	SDNType          string `json:"sdnType,omitempty"`
	ClusterName      string `json:"clusterName,omitempty"`
	Region           string `json:"region,omitempty"`
	Fips             bool   `json:"fips,omitempty"`
	Publish          string `json:"publish,omitempty"`
	WorkerArch       string `json:"workerArch,omitempty"`
	ControlPlaneArch string `json:"controlPlaneArch,omitempty"`
	Ipsec            bool   `json:"ipsec,omitempty"`
	IpsecMode        string `json:"ipsecMode,omitempty"`
}
