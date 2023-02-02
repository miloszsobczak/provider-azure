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

type CompositeIndexObservation struct {
}

type CompositeIndexParameters struct {

	// One or more index blocks as defined below.
	// +kubebuilder:validation:Required
	Index []IndexParameters `json:"index" tf:"index,omitempty"`
}

type ConflictResolutionPolicyObservation struct {
}

type ConflictResolutionPolicyParameters struct {

	// The conflict resolution path in the case of LastWriterWins mode.
	// +kubebuilder:validation:Optional
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path,omitempty"`

	// The procedure to resolve conflicts in the case of custom mode.
	// +kubebuilder:validation:Optional
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure,omitempty"`

	// Indicates the conflict resolution mode. Possible values include: LastWriterWins, Custom.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type GremlinGraphAutoscaleSettingsObservation struct {
}

type GremlinGraphAutoscaleSettingsParameters struct {

	// The maximum throughput of the Gremlin graph (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type GremlinGraphObservation struct {

	// The ID of the CosmosDB Gremlin Graph.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The configuration of the indexing policy. One or more index_policy blocks as defined below.
	// +kubebuilder:validation:Optional
	IndexPolicy []IndexPolicyObservation `json:"indexPolicy,omitempty" tf:"index_policy,omitempty"`
}

type GremlinGraphParameters struct {

	// The name of the CosmosDB Account to create the Gremlin Graph within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// An autoscale_settings block as defined below. Requires partition_key_path to be set.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []GremlinGraphAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// A conflict_resolution_policy blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ConflictResolutionPolicy []ConflictResolutionPolicyParameters `json:"conflictResolutionPolicy,omitempty" tf:"conflict_resolution_policy,omitempty"`

	// The name of the Cosmos DB Graph Database in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=GremlinDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a GremlinDatabase to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a GremlinDatabase to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The default time to live (TTL) of the Gremlin graph. If the value is missing or set to "-1", items don’t expire.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The configuration of the indexing policy. One or more index_policy blocks as defined below.
	// +kubebuilder:validation:Optional
	IndexPolicy []IndexPolicyParameters `json:"indexPolicy,omitempty" tf:"index_policy,omitempty"`

	// Define a partition key. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	PartitionKeyPath *string `json:"partitionKeyPath" tf:"partition_key_path,omitempty"`

	// Define a partition key version. Changing this forces a new resource to be created. Possible values are 1and 2. This should be set to 2 in order to use large partition keys.
	// +kubebuilder:validation:Optional
	PartitionKeyVersion *float64 `json:"partitionKeyVersion,omitempty" tf:"partition_key_version,omitempty"`

	// The name of the resource group in which the Cosmos DB Gremlin Graph is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The throughput of the Gremlin graph (RU/s). Must be set in increments of 100. The minimum value is 400.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// One or more unique_key blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	UniqueKey []UniqueKeyParameters `json:"uniqueKey,omitempty" tf:"unique_key,omitempty"`
}

type IndexObservation struct {
}

type IndexParameters struct {

	// Order of the index. Possible values are Ascending or Descending.
	// +kubebuilder:validation:Required
	Order *string `json:"order" tf:"order,omitempty"`

	// Path for which the indexing behaviour applies to. According to the service design, all spatial types including LineString, MultiPolygon, Point, and Polygon will be applied to the path.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IndexPolicyObservation struct {

	// One or more spatial_index blocks as defined below.
	// +kubebuilder:validation:Optional
	SpatialIndex []SpatialIndexObservation `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type IndexPolicyParameters struct {

	// Indicates if the indexing policy is automatic. Defaults to true.
	// +kubebuilder:validation:Optional
	Automatic *bool `json:"automatic,omitempty" tf:"automatic,omitempty"`

	// One or more composite_index blocks as defined below.
	// +kubebuilder:validation:Optional
	CompositeIndex []CompositeIndexParameters `json:"compositeIndex,omitempty" tf:"composite_index,omitempty"`

	// List of paths to exclude from indexing. Required if indexing_mode is Consistent or Lazy.
	// +kubebuilder:validation:Optional
	ExcludedPaths []*string `json:"excludedPaths,omitempty" tf:"excluded_paths,omitempty"`

	// List of paths to include in the indexing. Required if indexing_mode is Consistent or Lazy.
	// +kubebuilder:validation:Optional
	IncludedPaths []*string `json:"includedPaths,omitempty" tf:"included_paths,omitempty"`

	// Indicates the indexing mode. Possible values include: Consistent, Lazy, None.
	// +kubebuilder:validation:Required
	IndexingMode *string `json:"indexingMode" tf:"indexing_mode,omitempty"`

	// One or more spatial_index blocks as defined below.
	// +kubebuilder:validation:Optional
	SpatialIndex []SpatialIndexParameters `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type SpatialIndexObservation struct {
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type SpatialIndexParameters struct {

	// Path for which the indexing behaviour applies to. According to the service design, all spatial types including LineString, MultiPolygon, Point, and Polygon will be applied to the path.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type UniqueKeyObservation struct {
}

type UniqueKeyParameters struct {

	// A list of paths to use for this unique key. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Paths []*string `json:"paths" tf:"paths,omitempty"`
}

// GremlinGraphSpec defines the desired state of GremlinGraph
type GremlinGraphSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GremlinGraphParameters `json:"forProvider"`
}

// GremlinGraphStatus defines the observed state of GremlinGraph.
type GremlinGraphStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GremlinGraphObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GremlinGraph is the Schema for the GremlinGraphs API. Manages a Gremlin Graph within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type GremlinGraph struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GremlinGraphSpec   `json:"spec"`
	Status            GremlinGraphStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GremlinGraphList contains a list of GremlinGraphs
type GremlinGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GremlinGraph `json:"items"`
}

// Repository type metadata.
var (
	GremlinGraph_Kind             = "GremlinGraph"
	GremlinGraph_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GremlinGraph_Kind}.String()
	GremlinGraph_KindAPIVersion   = GremlinGraph_Kind + "." + CRDGroupVersion.String()
	GremlinGraph_GroupVersionKind = CRDGroupVersion.WithKind(GremlinGraph_Kind)
)

func init() {
	SchemeBuilder.Register(&GremlinGraph{}, &GremlinGraphList{})
}
