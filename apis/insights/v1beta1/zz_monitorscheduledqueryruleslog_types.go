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

type CriteriaDimensionObservation struct {
}

type CriteriaDimensionParameters struct {

	// Name of the dimension.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Operator for dimension values, - 'Include'.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of dimension values.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorScheduledQueryRulesLogCriteriaObservation struct {
}

type MonitorScheduledQueryRulesLogCriteriaParameters struct {

	// A dimension block as defined below.
	// +kubebuilder:validation:Required
	Dimension []CriteriaDimensionParameters `json:"dimension" tf:"dimension,omitempty"`

	// Name of the metric. Supported metrics are listed in the Azure Monitor Microsoft.OperationalInsights/workspaces metrics namespace.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`
}

type MonitorScheduledQueryRulesLogObservation struct {

	// The ID of the scheduled query rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorScheduledQueryRulesLogParameters struct {

	// A list of IDs of Resources referred into query.
	// +kubebuilder:validation:Optional
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// A criteria block as defined below.
	// +kubebuilder:validation:Required
	Criteria []MonitorScheduledQueryRulesLogCriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// The resource URI over which log search query is to be run.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDRef *v1.Reference `json:"dataSourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDSelector *v1.Selector `json:"dataSourceIdSelector,omitempty" tf:"-"`

	// The description of the scheduled query rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this scheduled query rule is enabled. Default is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure Region where the resource should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the resource group in which to create the scheduled query rule instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MonitorScheduledQueryRulesLogSpec defines the desired state of MonitorScheduledQueryRulesLog
type MonitorScheduledQueryRulesLogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorScheduledQueryRulesLogParameters `json:"forProvider"`
}

// MonitorScheduledQueryRulesLogStatus defines the observed state of MonitorScheduledQueryRulesLog.
type MonitorScheduledQueryRulesLogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorScheduledQueryRulesLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesLog is the Schema for the MonitorScheduledQueryRulesLogs API. Manages a LogToMetricAction Scheduled Query Rules resources within Azure Monitor
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorScheduledQueryRulesLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorScheduledQueryRulesLogSpec   `json:"spec"`
	Status            MonitorScheduledQueryRulesLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesLogList contains a list of MonitorScheduledQueryRulesLogs
type MonitorScheduledQueryRulesLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorScheduledQueryRulesLog `json:"items"`
}

// Repository type metadata.
var (
	MonitorScheduledQueryRulesLog_Kind             = "MonitorScheduledQueryRulesLog"
	MonitorScheduledQueryRulesLog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorScheduledQueryRulesLog_Kind}.String()
	MonitorScheduledQueryRulesLog_KindAPIVersion   = MonitorScheduledQueryRulesLog_Kind + "." + CRDGroupVersion.String()
	MonitorScheduledQueryRulesLog_GroupVersionKind = CRDGroupVersion.WithKind(MonitorScheduledQueryRulesLog_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorScheduledQueryRulesLog{}, &MonitorScheduledQueryRulesLogList{})
}
