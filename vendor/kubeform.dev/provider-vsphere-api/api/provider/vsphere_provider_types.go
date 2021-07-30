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

package provider

type VsphereSpec struct {
	// If set, VMware vSphere client will permit unverifiable SSL certificates.
	// +optional
	AllowUnverifiedSsl *bool `json:"allowUnverifiedSsl,omitempty" tf:"allow_unverified_ssl"`
	// API timeout in minutes (Default: 5)
	// +optional
	ApiTimeout *int64 `json:"apiTimeout,omitempty" tf:"api_timeout"`
	// govmomi debug
	// +optional
	ClientDebug *bool `json:"clientDebug,omitempty" tf:"client_debug"`
	// govmomi debug path for debug
	// +optional
	ClientDebugPath *string `json:"clientDebugPath,omitempty" tf:"client_debug_path"`
	// govmomi debug path for a single run
	// +optional
	ClientDebugPathRun *string `json:"clientDebugPathRun,omitempty" tf:"client_debug_path_run"`
	// The user password for vSphere API operations.
	Password *string `json:"password" tf:"password"`
	// Persist vSphere client sessions to disk
	// +optional
	PersistSession *bool `json:"persistSession,omitempty" tf:"persist_session"`
	// The directory to save vSphere REST API sessions to
	// +optional
	RestSessionPath *string `json:"restSessionPath,omitempty" tf:"rest_session_path"`
	// The user name for vSphere API operations.
	User *string `json:"user" tf:"user"`
	// +optional
	// Deprecated
	VcenterServer *string `json:"vcenterServer,omitempty" tf:"vcenter_server"`
	// Keep alive interval for the VIM session in minutes
	// +optional
	VimKeepAlive *int64 `json:"vimKeepAlive,omitempty" tf:"vim_keep_alive"`
	// The directory to save vSphere SOAP API sessions to
	// +optional
	VimSessionPath *string `json:"vimSessionPath,omitempty" tf:"vim_session_path"`
	// The vSphere Server name for vSphere API operations.
	// +optional
	VsphereServer *string `json:"vsphereServer,omitempty" tf:"vsphere_server"`
}
