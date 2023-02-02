/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DataProtectionReplicationObservation struct {
}

type DataProtectionReplicationParameters struct {

	// The endpoint type, default value is dst for destination.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Location of the primary volume. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	RemoteVolumeLocation *string `json:"remoteVolumeLocation" tf:"remote_volume_location,omitempty"`

	// Resource ID of the primary volume.
	// +crossplane:generate:reference:type=Volume
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RemoteVolumeResourceID *string `json:"remoteVolumeResourceId,omitempty" tf:"remote_volume_resource_id,omitempty"`

	// Reference to a Volume to populate remoteVolumeResourceId.
	// +kubebuilder:validation:Optional
	RemoteVolumeResourceIDRef *v1.Reference `json:"remoteVolumeResourceIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate remoteVolumeResourceId.
	// +kubebuilder:validation:Optional
	RemoteVolumeResourceIDSelector *v1.Selector `json:"remoteVolumeResourceIdSelector,omitempty" tf:"-"`

	// Replication frequency, supported values are '10minutes', 'hourly', 'daily', values are case sensitive.
	// +kubebuilder:validation:Required
	ReplicationFrequency *string `json:"replicationFrequency" tf:"replication_frequency,omitempty"`
}

type DataProtectionSnapshotPolicyObservation struct {
}

type DataProtectionSnapshotPolicyParameters struct {

	// Resource ID of the snapshot policy to apply to the volume.
	// +crossplane:generate:reference:type=SnapshotPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SnapshotPolicyID *string `json:"snapshotPolicyId,omitempty" tf:"snapshot_policy_id,omitempty"`

	// Reference to a SnapshotPolicy to populate snapshotPolicyId.
	// +kubebuilder:validation:Optional
	SnapshotPolicyIDRef *v1.Reference `json:"snapshotPolicyIdRef,omitempty" tf:"-"`

	// Selector for a SnapshotPolicy to populate snapshotPolicyId.
	// +kubebuilder:validation:Optional
	SnapshotPolicyIDSelector *v1.Selector `json:"snapshotPolicyIdSelector,omitempty" tf:"-"`
}

type ExportPolicyRuleObservation struct {
}

type ExportPolicyRuleParameters struct {

	// A list of allowed clients IPv4 addresses.
	// +kubebuilder:validation:Required
	AllowedClients []*string `json:"allowedClients" tf:"allowed_clients,omitempty"`

	// A list of allowed protocols. Valid values include CIFS, NFSv3, or NFSv4.1. Only one value is supported at this time. This replaces the previous arguments: cifs_enabled, nfsv3_enabled and nfsv4_enabled.
	// +kubebuilder:validation:Optional
	ProtocolsEnabled []*string `json:"protocolsEnabled,omitempty" tf:"protocols_enabled,omitempty"`

	// Is root access permitted to this volume?
	// +kubebuilder:validation:Optional
	RootAccessEnabled *bool `json:"rootAccessEnabled,omitempty" tf:"root_access_enabled,omitempty"`

	// The index number of the rule.
	// +kubebuilder:validation:Required
	RuleIndex *float64 `json:"ruleIndex" tf:"rule_index,omitempty"`

	// Is the file system on unix read only?
	// +kubebuilder:validation:Optional
	UnixReadOnly *bool `json:"unixReadOnly,omitempty" tf:"unix_read_only,omitempty"`

	// Is the file system on unix read and write?
	// +kubebuilder:validation:Optional
	UnixReadWrite *bool `json:"unixReadWrite,omitempty" tf:"unix_read_write,omitempty"`
}

type VolumeObservation struct {

	// The ID of the NetApp Volume.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of IPv4 Addresses which should be used to mount the volume.
	MountIPAddresses []*string `json:"mountIpAddresses,omitempty" tf:"mount_ip_addresses,omitempty"`
}

type VolumeParameters struct {

	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// Is the NetApp Volume enabled for Azure VMware Solution (AVS) datastore purpose. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AzureVMwareDataStoreEnabled *bool `json:"azureVmwareDataStoreEnabled,omitempty" tf:"azure_vmware_data_store_enabled,omitempty"`

	// Creates volume from snapshot. Following properties must be the same as the original volume where the snapshot was taken from: protocols, subnet_id, location, service_level, resource_group_name, account_name and pool_name. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Snapshot
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceID *string `json:"createFromSnapshotResourceId,omitempty" tf:"create_from_snapshot_resource_id,omitempty"`

	// Reference to a Snapshot to populate createFromSnapshotResourceId.
	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceIDRef *v1.Reference `json:"createFromSnapshotResourceIdRef,omitempty" tf:"-"`

	// Selector for a Snapshot to populate createFromSnapshotResourceId.
	// +kubebuilder:validation:Optional
	CreateFromSnapshotResourceIDSelector *v1.Selector `json:"createFromSnapshotResourceIdSelector,omitempty" tf:"-"`

	// A data_protection_replication block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataProtectionReplication []DataProtectionReplicationParameters `json:"dataProtectionReplication,omitempty" tf:"data_protection_replication,omitempty"`

	// A data_protection_snapshot_policy block as defined below.
	// +kubebuilder:validation:Optional
	DataProtectionSnapshotPolicy []DataProtectionSnapshotPolicyParameters `json:"dataProtectionSnapshotPolicy,omitempty" tf:"data_protection_snapshot_policy,omitempty"`

	// One or more export_policy_rule block defined below.
	// +kubebuilder:validation:Optional
	ExportPolicyRule []ExportPolicyRuleParameters `json:"exportPolicyRule,omitempty" tf:"export_policy_rule,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Indicates which network feature to use, accepted values are Basic or Standard, it defaults to Basic if not defined. This is a feature in public preview and for more information about it and how to register, please refer to Configure network features for an Azure NetApp Files volume. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	NetworkFeatures *string `json:"networkFeatures,omitempty" tf:"network_features,omitempty"`

	// The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Pool
	// +kubebuilder:validation:Optional
	PoolName *string `json:"poolName,omitempty" tf:"pool_name,omitempty"`

	// Reference to a Pool to populate poolName.
	// +kubebuilder:validation:Optional
	PoolNameRef *v1.Reference `json:"poolNameRef,omitempty" tf:"-"`

	// Selector for a Pool to populate poolName.
	// +kubebuilder:validation:Optional
	PoolNameSelector *v1.Selector `json:"poolNameSelector,omitempty" tf:"-"`

	// The target volume protocol expressed as a list. Supported single value include CIFS, NFSv3, or NFSv4.1. If argument is not defined it will default to NFSv3. Changing this forces a new resource to be created and data will be lost. Dual protocol scenario is supported for CIFS and NFSv3, for more information, please refer to Create a dual-protocol volume for Azure NetApp Files document.
	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Volume security style, accepted values are Unix or Ntfs. If not provided, single-protocol volume is created defaulting to Unix if it is NFSv3 or NFSv4.1 volume, if CIFS, it will default to Ntfs. In a dual-protocol volume, if not provided, its value will be Ntfs. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecurityStyle *string `json:"securityStyle,omitempty" tf:"security_style,omitempty"`

	// The target performance of the file system. Valid values include Premium, Standard, or Ultra. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ServiceLevel *string `json:"serviceLevel" tf:"service_level,omitempty"`

	// Specifies whether the .snapshot (NFS clients) or ~snapshot (SMB clients) path of a volume is visible, default value is true.
	// +kubebuilder:validation:Optional
	SnapshotDirectoryVisible *bool `json:"snapshotDirectoryVisible,omitempty" tf:"snapshot_directory_visible,omitempty"`

	// The maximum Storage Quota allowed for a file system in Gigabytes.
	// +kubebuilder:validation:Required
	StorageQuotaInGb *float64 `json:"storageQuotaInGb" tf:"storage_quota_in_gb,omitempty"`

	// The ID of the Subnet the NetApp Volume resides in, which must have the Microsoft.NetApp/volumes delegation. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Throughput of this volume in Mibps.
	// +kubebuilder:validation:Optional
	ThroughputInMibps *float64 `json:"throughputInMibps,omitempty" tf:"throughput_in_mibps,omitempty"`

	// A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	VolumePath *string `json:"volumePath" tf:"volume_path,omitempty"`

	// Specifies the Availability Zone in which the Volume should be located. Possible values are 1, 2 and 3. Changing this forces a new resource to be created. This feature is currently in preview, for more information on how to enable it, please refer to Manage availability zone volume placement for Azure NetApp Files.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Volume is the Schema for the Volumes API. Manages a NetApp Volume.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec"`
	Status            VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
