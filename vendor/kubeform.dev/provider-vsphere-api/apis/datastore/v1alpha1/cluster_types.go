/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of custom attributes to set on this resource.
	// +optional
	CustomAttributes *map[string]string `json:"customAttributes,omitempty" tf:"custom_attributes"`
	// The managed object ID of the datacenter to put the datastore cluster in.
	DatacenterID *string `json:"datacenterID" tf:"datacenter_id"`
	// The name of the folder to locate the datastore cluster in.
	// +optional
	Folder *string `json:"folder,omitempty" tf:"folder"`
	// Name for the new storage pod.
	Name *string `json:"name" tf:"name"`
	// Advanced configuration options for storage DRS.
	// +optional
	SdrsAdvancedOptions *map[string]string `json:"sdrsAdvancedOptions,omitempty" tf:"sdrs_advanced_options"`
	// The default automation level for all virtual machines in this storage cluster.
	// +optional
	SdrsAutomationLevel *string `json:"sdrsAutomationLevel,omitempty" tf:"sdrs_automation_level"`
	// When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
	// +optional
	SdrsDefaultIntraVmAffinity *bool `json:"sdrsDefaultIntraVmAffinity,omitempty" tf:"sdrs_default_intra_vm_affinity"`
	// Enable storage DRS for this datastore cluster.
	// +optional
	SdrsEnabled *bool `json:"sdrsEnabled,omitempty" tf:"sdrs_enabled"`
	// The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	// +optional
	SdrsFreeSpaceThreshold *int64 `json:"sdrsFreeSpaceThreshold,omitempty" tf:"sdrs_free_space_threshold"`
	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to freeSpace, drs_free_space_threshold is used.
	// +optional
	SdrsFreeSpaceThresholdMode *string `json:"sdrsFreeSpaceThresholdMode,omitempty" tf:"sdrs_free_space_threshold_mode"`
	// The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to balance the space.
	// +optional
	SdrsFreeSpaceUtilizationDifference *int64 `json:"sdrsFreeSpaceUtilizationDifference,omitempty" tf:"sdrs_free_space_utilization_difference"`
	// Overrides the default automation settings when correcting I/O load imbalances.
	// +optional
	SdrsIoBalanceAutomationLevel *string `json:"sdrsIoBalanceAutomationLevel,omitempty" tf:"sdrs_io_balance_automation_level"`
	// The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this datastore.
	// +optional
	SdrsIoLatencyThreshold *int64 `json:"sdrsIoLatencyThreshold,omitempty" tf:"sdrs_io_latency_threshold"`
	// Enable I/O load balancing for this datastore cluster.
	// +optional
	SdrsIoLoadBalanceEnabled *bool `json:"sdrsIoLoadBalanceEnabled,omitempty" tf:"sdrs_io_load_balance_enabled"`
	// The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
	// +optional
	SdrsIoLoadImbalanceThreshold *int64 `json:"sdrsIoLoadImbalanceThreshold,omitempty" tf:"sdrs_io_load_imbalance_threshold"`
	// The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to move VMs off of a datastore.
	// +optional
	SdrsIoReservableIopsThreshold *int64 `json:"sdrsIoReservableIopsThreshold,omitempty" tf:"sdrs_io_reservable_iops_threshold"`
	// The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
	// +optional
	SdrsIoReservablePercentThreshold *int64 `json:"sdrsIoReservablePercentThreshold,omitempty" tf:"sdrs_io_reservable_percent_threshold"`
	// The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
	// +optional
	SdrsIoReservableThresholdMode *string `json:"sdrsIoReservableThresholdMode,omitempty" tf:"sdrs_io_reservable_threshold_mode"`
	// The storage DRS poll interval, in minutes.
	// +optional
	SdrsLoadBalanceInterval *int64 `json:"sdrsLoadBalanceInterval,omitempty" tf:"sdrs_load_balance_interval"`
	// Overrides the default automation settings when correcting storage and VM policy violations.
	// +optional
	SdrsPolicyEnforcementAutomationLevel *string `json:"sdrsPolicyEnforcementAutomationLevel,omitempty" tf:"sdrs_policy_enforcement_automation_level"`
	// Overrides the default automation settings when correcting affinity rule violations.
	// +optional
	SdrsRuleEnforcementAutomationLevel *string `json:"sdrsRuleEnforcementAutomationLevel,omitempty" tf:"sdrs_rule_enforcement_automation_level"`
	// Overrides the default automation settings when correcting disk space imbalances.
	// +optional
	SdrsSpaceBalanceAutomationLevel *string `json:"sdrsSpaceBalanceAutomationLevel,omitempty" tf:"sdrs_space_balance_automation_level"`
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	// +optional
	SdrsSpaceUtilizationThreshold *int64 `json:"sdrsSpaceUtilizationThreshold,omitempty" tf:"sdrs_space_utilization_threshold"`
	// Overrides the default automation settings when generating recommendations for datastore evacuation.
	// +optional
	SdrsVmEvacuationAutomationLevel *string `json:"sdrsVmEvacuationAutomationLevel,omitempty" tf:"sdrs_vm_evacuation_automation_level"`
	// A list of tag IDs to apply to this object.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type ClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
