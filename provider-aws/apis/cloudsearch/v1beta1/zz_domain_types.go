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
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DocumentServiceEndpoint *string `json:"documentServiceEndpoint,omitempty" tf:"document_service_endpoint,omitempty"`

	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SearchServiceEndpoint *string `json:"searchServiceEndpoint,omitempty" tf:"search_service_endpoint,omitempty"`
}

type DomainParameters struct {

	// +kubebuilder:validation:Optional
	EndpointOptions []EndpointOptionsParameters `json:"endpointOptions,omitempty" tf:"endpoint_options,omitempty"`

	// +kubebuilder:validation:Optional
	IndexField []IndexFieldParameters `json:"indexField,omitempty" tf:"index_field,omitempty"`

	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ScalingParameters []ScalingParametersParameters `json:"scalingParameters,omitempty" tf:"scaling_parameters,omitempty"`
}

type EndpointOptionsObservation struct {
}

type EndpointOptionsParameters struct {

	// +kubebuilder:validation:Optional
	EnforceHTTPS *bool `json:"enforceHttps,omitempty" tf:"enforce_https,omitempty"`

	// +kubebuilder:validation:Optional
	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty" tf:"tls_security_policy,omitempty"`
}

type IndexFieldObservation struct {
}

type IndexFieldParameters struct {

	// +kubebuilder:validation:Optional
	AnalysisScheme *string `json:"analysisScheme,omitempty" tf:"analysis_scheme,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Facet *bool `json:"facet,omitempty" tf:"facet,omitempty"`

	// +kubebuilder:validation:Optional
	Highlight *bool `json:"highlight,omitempty" tf:"highlight,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Return *bool `json:"return,omitempty" tf:"return,omitempty"`

	// +kubebuilder:validation:Optional
	Search *bool `json:"search,omitempty" tf:"search,omitempty"`

	// +kubebuilder:validation:Optional
	Sort *bool `json:"sort,omitempty" tf:"sort,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScalingParametersObservation struct {
}

type ScalingParametersParameters struct {

	// +kubebuilder:validation:Optional
	DesiredInstanceType *string `json:"desiredInstanceType,omitempty" tf:"desired_instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredPartitionCount *float64 `json:"desiredPartitionCount,omitempty" tf:"desired_partition_count,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredReplicationCount *float64 `json:"desiredReplicationCount,omitempty" tf:"desired_replication_count,omitempty"`
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

// Domain is the Schema for the Domains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
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
