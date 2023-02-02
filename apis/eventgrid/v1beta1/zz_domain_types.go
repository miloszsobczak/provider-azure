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

type DomainObservation struct {

	// The Endpoint associated with the EventGrid Domain.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the EventGrid Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`
}

type DomainParameters struct {

	// Whether to create the domain topic when the first event subscription at the scope of the domain topic is created. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoCreateTopicWithFirstSubscription *bool `json:"autoCreateTopicWithFirstSubscription,omitempty" tf:"auto_create_topic_with_first_subscription,omitempty"`

	// Whether to delete the domain topic when the last event subscription at the scope of the domain topic is deleted. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoDeleteTopicWithLastSubscription *bool `json:"autoDeleteTopicWithLastSubscription,omitempty" tf:"auto_delete_topic_with_last_subscription,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// One or more inbound_ip_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	InboundIPRule []InboundIPRuleParameters `json:"inboundIpRule,omitempty" tf:"inbound_ip_rule,omitempty"`

	// A input_mapping_default_values block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InputMappingDefaultValues []InputMappingDefaultValuesParameters `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values,omitempty"`

	// A input_mapping_fields block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InputMappingFields []InputMappingFieldsParameters `json:"inputMappingFields,omitempty" tf:"input_mapping_fields,omitempty"`

	// Specifies the schema in which incoming events will be published to this domain. Allowed values are CloudEventSchemaV1_0, CustomEventSchema, or EventGridSchema. Defaults to eventgridschema. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InputSchema *string `json:"inputSchema,omitempty" tf:"input_schema,omitempty"`

	// Whether local authentication methods is enabled for the EventGrid Domain. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Whether or not public network access is allowed for this server. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which the EventGrid Domain exists. Changing this forces a new resource to be created.
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

type IdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Event Grid Domain.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Event Grid Domain. Possible values are SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InboundIPRuleObservation struct {
}

type InboundIPRuleParameters struct {

	// The action to take when the rule is matched. Possible values are Allow.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The IP mask (CIDR) to match on.
	// +kubebuilder:validation:Optional
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask"`
}

type InputMappingDefaultValuesObservation struct {
}

type InputMappingDefaultValuesParameters struct {

	// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type InputMappingFieldsObservation struct {
}

type InputMappingFieldsParameters struct {

	// Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`

	// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventTime *string `json:"eventTime,omitempty" tf:"event_time,omitempty"`

	// Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API. Manages an EventGrid Domain
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
