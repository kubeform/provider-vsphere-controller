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

type PortGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortGroupSpec   `json:"spec,omitempty"`
	Status            PortGroupStatus `json:"status,omitempty"`
}

type PortGroupSpecPorts struct {
	// The linkable identifier for this port entry.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The MAC addresses of the network service of the virtual machine connected on this port.
	// +optional
	MacAddresses []string `json:"macAddresses,omitempty" tf:"mac_addresses"`
	// Type type of the entity connected on this port. Possible values are host (VMKkernel), systemManagement (service console), virtualMachine, or unknown.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PortGroupSpec struct {
	State *PortGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource PortGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PortGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// List of active network adapters used for load balancing.
	// +optional
	ActiveNics []string `json:"activeNics,omitempty" tf:"active_nics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	// +optional
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	// +optional
	AllowMACChanges *bool `json:"allowMACChanges,omitempty" tf:"allow_mac_changes"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	// +optional
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	// +optional
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon"`
	// The effective network policy after inheritance. Note that this will look similar to, but is not the same, as the policy attributes defined in this resource.
	// +optional
	ComputedPolicy *map[string]string `json:"computedPolicy,omitempty" tf:"computed_policy"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	// +optional
	Failback *bool `json:"failback,omitempty" tf:"failback"`
	// The managed object ID of the host to set the virtual switch up on.
	HostSystemID *string `json:"hostSystemID" tf:"host_system_id"`
	// The linkable identifier for this port group.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// The name of the port group.
	Name *string `json:"name" tf:"name"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	// +optional
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches"`
	// The ports that currently exist and are used on this port group.
	// +optional
	Ports []PortGroupSpecPorts `json:"ports,omitempty" tf:"ports"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	// +optional
	ShapingAverageBandwidth *int64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	// +optional
	ShapingBurstSize *int64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size"`
	// Enable traffic shaping on this virtual switch or port group.
	// +optional
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	// +optional
	ShapingPeakBandwidth *int64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth"`
	// List of standby network adapters used for failover.
	// +optional
	StandbyNics []string `json:"standbyNics,omitempty" tf:"standby_nics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	// +optional
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy"`
	// The name of the virtual switch to bind this port group to.
	VirtualSwitchName *string `json:"virtualSwitchName" tf:"virtual_switch_name"`
	// The VLAN ID/trunk mode for this port group. An ID of 0 denotes no tagging, an ID of 1-4094 tags with the specific ID, and an ID of 4095 enables trunk mode, allowing the guest to manage its own tagging.
	// +optional
	VlanID *int64 `json:"vlanID,omitempty" tf:"vlan_id"`
}

type PortGroupStatus struct {
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

// PortGroupList is a list of PortGroups
type PortGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PortGroup CRD objects
	Items []PortGroup `json:"items,omitempty"`
}
