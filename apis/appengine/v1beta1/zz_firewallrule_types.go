/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallRuleInitParameters struct {

	// The action to take if this rule matches.
	// Possible values are: UNSPECIFIED_ACTION, ALLOW, DENY.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional string description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange *string `json:"sourceRange,omitempty" tf:"source_range,omitempty"`
}

type FirewallRuleObservation struct {

	// The action to take if this rule matches.
	// Possible values are: UNSPECIFIED_ACTION, ALLOW, DENY.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional string description of this rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format apps/{{project}}/firewall/ingressRules/{{priority}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange *string `json:"sourceRange,omitempty" tf:"source_range,omitempty"`
}

type FirewallRuleParameters struct {

	// The action to take if this rule matches.
	// Possible values are: UNSPECIFIED_ACTION, ALLOW, DENY.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional string description of this rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/appengine/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Application in appengine to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Application in appengine to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	// +kubebuilder:validation:Optional
	SourceRange *string `json:"sourceRange,omitempty" tf:"source_range,omitempty"`
}

// FirewallRuleSpec defines the desired state of FirewallRule
type FirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// FirewallRuleStatus defines the observed state of FirewallRule.
type FirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRule is the Schema for the FirewallRules API. A single firewall rule that is evaluated against incoming traffic and provides an action to take on matched requests.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceRange) || (has(self.initProvider) && has(self.initProvider.sourceRange))",message="spec.forProvider.sourceRange is a required parameter"
	Spec   FirewallRuleSpec   `json:"spec"`
	Status FirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleList contains a list of FirewallRules
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallRule_Kind             = "FirewallRule"
	FirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRule_Kind}.String()
	FirewallRule_KindAPIVersion   = FirewallRule_Kind + "." + CRDGroupVersion.String()
	FirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRule{}, &FirewallRuleList{})
}
