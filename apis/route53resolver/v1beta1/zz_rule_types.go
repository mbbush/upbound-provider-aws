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

type RuleInitParameters struct {

	// DNS queries for this domain name are forwarded to the IP addresses that are specified using target_ip.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using target_ip.
	// This argument should only be specified for FORWARD type rules.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53resolver/v1beta1.Endpoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResolverEndpointID *string `json:"resolverEndpointId,omitempty" tf:"resolver_endpoint_id,omitempty"`

	// Reference to a Endpoint in route53resolver to populate resolverEndpointId.
	// +kubebuilder:validation:Optional
	ResolverEndpointIDRef *v1.Reference `json:"resolverEndpointIdRef,omitempty" tf:"-"`

	// Selector for a Endpoint in route53resolver to populate resolverEndpointId.
	// +kubebuilder:validation:Optional
	ResolverEndpointIDSelector *v1.Selector `json:"resolverEndpointIdSelector,omitempty" tf:"-"`

	// The rule type. Valid values are FORWARD, SYSTEM and RECURSIVE.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for FORWARD type rules.
	TargetIP []TargetIPInitParameters `json:"targetIp,omitempty" tf:"target_ip,omitempty"`
}

type RuleObservation struct {

	// The ARN (Amazon Resource Name) for the resolver rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// DNS queries for this domain name are forwarded to the IP addresses that are specified using target_ip.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The ID of the resolver rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using target_ip.
	// This argument should only be specified for FORWARD type rules.
	ResolverEndpointID *string `json:"resolverEndpointId,omitempty" tf:"resolver_endpoint_id,omitempty"`

	// The rule type. Valid values are FORWARD, SYSTEM and RECURSIVE.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
	// Values are NOT_SHARED, SHARED_BY_ME or SHARED_WITH_ME
	ShareStatus *string `json:"shareStatus,omitempty" tf:"share_status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for FORWARD type rules.
	TargetIP []TargetIPObservation `json:"targetIp,omitempty" tf:"target_ip,omitempty"`
}

type RuleParameters struct {

	// DNS queries for this domain name are forwarded to the IP addresses that are specified using target_ip.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using target_ip.
	// This argument should only be specified for FORWARD type rules.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53resolver/v1beta1.Endpoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResolverEndpointID *string `json:"resolverEndpointId,omitempty" tf:"resolver_endpoint_id,omitempty"`

	// Reference to a Endpoint in route53resolver to populate resolverEndpointId.
	// +kubebuilder:validation:Optional
	ResolverEndpointIDRef *v1.Reference `json:"resolverEndpointIdRef,omitempty" tf:"-"`

	// Selector for a Endpoint in route53resolver to populate resolverEndpointId.
	// +kubebuilder:validation:Optional
	ResolverEndpointIDSelector *v1.Selector `json:"resolverEndpointIdSelector,omitempty" tf:"-"`

	// The rule type. Valid values are FORWARD, SYSTEM and RECURSIVE.
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for FORWARD type rules.
	// +kubebuilder:validation:Optional
	TargetIP []TargetIPParameters `json:"targetIp,omitempty" tf:"target_ip,omitempty"`
}

type TargetIPInitParameters struct {

	// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The port at ip that you want to forward DNS queries to. Default value is 53.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol for the resolver endpoint. Valid values can be found in the AWS documentation. Default value is Do53.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type TargetIPObservation struct {

	// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The port at ip that you want to forward DNS queries to. Default value is 53.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol for the resolver endpoint. Valid values can be found in the AWS documentation. Default value is Do53.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type TargetIPParameters struct {

	// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip" tf:"ip,omitempty"`

	// The port at ip that you want to forward DNS queries to. Default value is 53.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol for the resolver endpoint. Valid values can be found in the AWS documentation. Default value is Do53.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Provides a Route53 Resolver rule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainName) || (has(self.initProvider) && has(self.initProvider.domainName))",message="spec.forProvider.domainName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleType) || (has(self.initProvider) && has(self.initProvider.ruleType))",message="spec.forProvider.ruleType is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
