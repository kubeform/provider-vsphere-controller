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

type ClusterVmDependencyRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterVmDependencyRuleSpec   `json:"spec,omitempty"`
	Status            ClusterVmDependencyRuleStatus `json:"status,omitempty"`
}

type ClusterVmDependencyRuleSpec struct {
	State *ClusterVmDependencyRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterVmDependencyRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterVmDependencyRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The managed object ID of the cluster.
	ComputeClusterID *string `json:"computeClusterID" tf:"compute_cluster_id"`
	// The name of the VM group that this rule depends on. The VMs defined in the group specified by vm_group_name will not be started until the VMs in this group are started.
	DependencyVmGroupName *string `json:"dependencyVmGroupName" tf:"dependency_vm_group_name"`
	// Enable this rule in the cluster.
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// When true, prevents any virtual machine operations that may violate this rule.
	// +optional
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory"`
	// The unique name of the virtual machine group in the cluster.
	Name *string `json:"name" tf:"name"`
	// The name of the VM group that is the subject of this rule. The VMs defined in this group will not be started until the VMs in the group specified by dependency_vm_group_name are started.
	VmGroupName *string `json:"vmGroupName" tf:"vm_group_name"`
}

type ClusterVmDependencyRuleStatus struct {
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

// ClusterVmDependencyRuleList is a list of ClusterVmDependencyRules
type ClusterVmDependencyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterVmDependencyRule CRD objects
	Items []ClusterVmDependencyRule `json:"items,omitempty"`
}
