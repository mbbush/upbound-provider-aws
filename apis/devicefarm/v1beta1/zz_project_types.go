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

type ProjectObservation struct {

	// The Amazon Resource Name of this project
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProjectParameters struct {

	// Sets the execution timeout value (in minutes) for a project. All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
	// +kubebuilder:validation:Optional
	DefaultJobTimeoutMinutes *float64 `json:"defaultJobTimeoutMinutes,omitempty" tf:"default_job_timeout_minutes,omitempty"`

	// The name of the project
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API. Provides a Devicefarm project
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}