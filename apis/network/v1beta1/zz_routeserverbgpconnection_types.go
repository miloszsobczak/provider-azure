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

type RouteServerBGPConnectionObservation struct {

	// The ID of the Route Server Bgp Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The peer autonomous system number for the Route Server Bgp Connection. Changing this forces a new resource to be created.
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// The peer ip address for the Route Server Bgp Connection. Changing this forces a new resource to be created.
	PeerIP *string `json:"peerIp,omitempty" tf:"peer_ip,omitempty"`

	// The ID of the Route Server within which this Bgp connection should be created. Changing this forces a new resource to be created.
	RouteServerID *string `json:"routeServerId,omitempty" tf:"route_server_id,omitempty"`
}

type RouteServerBGPConnectionParameters struct {

	// The peer autonomous system number for the Route Server Bgp Connection. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// The peer ip address for the Route Server Bgp Connection. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PeerIP *string `json:"peerIp,omitempty" tf:"peer_ip,omitempty"`

	// The ID of the Route Server within which this Bgp connection should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.RouteServer
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RouteServerID *string `json:"routeServerId,omitempty" tf:"route_server_id,omitempty"`

	// Reference to a RouteServer in network to populate routeServerId.
	// +kubebuilder:validation:Optional
	RouteServerIDRef *v1.Reference `json:"routeServerIdRef,omitempty" tf:"-"`

	// Selector for a RouteServer in network to populate routeServerId.
	// +kubebuilder:validation:Optional
	RouteServerIDSelector *v1.Selector `json:"routeServerIdSelector,omitempty" tf:"-"`
}

// RouteServerBGPConnectionSpec defines the desired state of RouteServerBGPConnection
type RouteServerBGPConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteServerBGPConnectionParameters `json:"forProvider"`
}

// RouteServerBGPConnectionStatus defines the observed state of RouteServerBGPConnection.
type RouteServerBGPConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteServerBGPConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteServerBGPConnection is the Schema for the RouteServerBGPConnections API. Manages a BGP Connection for a Route Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RouteServerBGPConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.peerAsn)",message="peerAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.peerIp)",message="peerIp is a required parameter"
	Spec   RouteServerBGPConnectionSpec   `json:"spec"`
	Status RouteServerBGPConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteServerBGPConnectionList contains a list of RouteServerBGPConnections
type RouteServerBGPConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteServerBGPConnection `json:"items"`
}

// Repository type metadata.
var (
	RouteServerBGPConnection_Kind             = "RouteServerBGPConnection"
	RouteServerBGPConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteServerBGPConnection_Kind}.String()
	RouteServerBGPConnection_KindAPIVersion   = RouteServerBGPConnection_Kind + "." + CRDGroupVersion.String()
	RouteServerBGPConnection_GroupVersionKind = CRDGroupVersion.WithKind(RouteServerBGPConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteServerBGPConnection{}, &RouteServerBGPConnectionList{})
}
