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

type Vnic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VnicSpec   `json:"spec,omitempty"`
	Status            VnicStatus `json:"status,omitempty"`
}

type VnicSpecIpv4 struct {
	// Use DHCP to configure the interface's IPv4 stack.
	// +optional
	Dhcp *bool `json:"dhcp,omitempty" tf:"dhcp"`
	// IP address of the default gateway, if DHCP is not set.
	// +optional
	Gw *string `json:"gw,omitempty" tf:"gw"`
	// address of the interface, if DHCP is not set.
	// +optional
	Ip *string `json:"ip,omitempty" tf:"ip"`
	// netmask of the interface, if DHCP is not set.
	// +optional
	Netmask *string `json:"netmask,omitempty" tf:"netmask"`
}

type VnicSpecIpv6 struct {
	// List of IPv6 addresses
	// +optional
	Addresses []string `json:"addresses,omitempty" tf:"addresses"`
	// Use IPv6 Autoconfiguration (RFC2462).
	// +optional
	Autoconfig *bool `json:"autoconfig,omitempty" tf:"autoconfig"`
	// Use DHCP to configure the interface's IPv4 stack.
	// +optional
	Dhcp *bool `json:"dhcp,omitempty" tf:"dhcp"`
	// IP address of the default gateway, if DHCP or autoconfig is not set.
	// +optional
	Gw *string `json:"gw,omitempty" tf:"gw"`
}

type VnicSpec struct {
	State *VnicSpecResource `json:"state,omitempty" tf:"-"`

	Resource VnicSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VnicSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Key of the distributed portgroup the nic will connect to
	// +optional
	DistributedPortGroup *string `json:"distributedPortGroup,omitempty" tf:"distributed_port_group"`
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	// +optional
	DistributedSwitchPort *string `json:"distributedSwitchPort,omitempty" tf:"distributed_switch_port"`
	// ESX host the interface belongs to
	Host *string `json:"host" tf:"host"`
	// +optional
	Ipv4 *VnicSpecIpv4 `json:"ipv4,omitempty" tf:"ipv4"`
	// +optional
	Ipv6 *VnicSpecIpv6 `json:"ipv6,omitempty" tf:"ipv6"`
	// MAC address of the interface.
	// +optional
	Mac *string `json:"mac,omitempty" tf:"mac"`
	// MTU of the interface.
	// +optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu"`
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'provisioning'
	// +optional
	Netstack *string `json:"netstack,omitempty" tf:"netstack"`
	// portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	// +optional
	Portgroup *string `json:"portgroup,omitempty" tf:"portgroup"`
}

type VnicStatus struct {
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

// VnicList is a list of Vnics
type VnicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Vnic CRD objects
	Items []Vnic `json:"items,omitempty"`
}
