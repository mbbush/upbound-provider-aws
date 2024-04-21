// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedPrefixListEntryInitParameters struct {

	// CIDR block of this entry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("cidr_block",false)
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Reference to a VPC in ec2 to populate cidr.
	// +kubebuilder:validation:Optional
	CidrRef *v1.Reference `json:"cidrRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate cidr.
	// +kubebuilder:validation:Optional
	CidrSelector *v1.Selector `json:"cidrSelector,omitempty" tf:"-"`

	// Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the prefix list.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// Reference to a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDRef *v1.Reference `json:"prefixListIdRef,omitempty" tf:"-"`

	// Selector for a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDSelector *v1.Selector `json:"prefixListIdSelector,omitempty" tf:"-"`
}

type ManagedPrefixListEntryObservation struct {

	// CIDR block of this entry.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the managed prefix list entry.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the prefix list.
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`
}

type ManagedPrefixListEntryParameters struct {

	// CIDR block of this entry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("cidr_block",false)
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Reference to a VPC in ec2 to populate cidr.
	// +kubebuilder:validation:Optional
	CidrRef *v1.Reference `json:"cidrRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate cidr.
	// +kubebuilder:validation:Optional
	CidrSelector *v1.Selector `json:"cidrSelector,omitempty" tf:"-"`

	// Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the prefix list.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// Reference to a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDRef *v1.Reference `json:"prefixListIdRef,omitempty" tf:"-"`

	// Selector for a ManagedPrefixList in ec2 to populate prefixListId.
	// +kubebuilder:validation:Optional
	PrefixListIDSelector *v1.Selector `json:"prefixListIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ManagedPrefixListEntrySpec defines the desired state of ManagedPrefixListEntry
type ManagedPrefixListEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPrefixListEntryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ManagedPrefixListEntryInitParameters `json:"initProvider,omitempty"`
}

// ManagedPrefixListEntryStatus defines the observed state of ManagedPrefixListEntry.
type ManagedPrefixListEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPrefixListEntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedPrefixListEntry is the Schema for the ManagedPrefixListEntrys API. Use the
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ManagedPrefixListEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedPrefixListEntrySpec   `json:"spec"`
	Status            ManagedPrefixListEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrefixListEntryList contains a list of ManagedPrefixListEntrys
type ManagedPrefixListEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPrefixListEntry `json:"items"`
}

// Repository type metadata.
var (
	ManagedPrefixListEntry_Kind             = "ManagedPrefixListEntry"
	ManagedPrefixListEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPrefixListEntry_Kind}.String()
	ManagedPrefixListEntry_KindAPIVersion   = ManagedPrefixListEntry_Kind + "." + CRDGroupVersion.String()
	ManagedPrefixListEntry_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPrefixListEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPrefixListEntry{}, &ManagedPrefixListEntryList{})
}
