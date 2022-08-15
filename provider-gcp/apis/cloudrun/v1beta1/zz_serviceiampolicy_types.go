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

type ServiceIAMPolicyObservation struct {

	// The etag of the IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceIAMPolicyParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// Used to find the parent resource to bind the IAM policy to
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

// ServiceIAMPolicySpec defines the desired state of ServiceIAMPolicy
type ServiceIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceIAMPolicyParameters `json:"forProvider"`
}

// ServiceIAMPolicyStatus defines the observed state of ServiceIAMPolicy.
type ServiceIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIAMPolicy is the Schema for the ServiceIAMPolicys API. Collection of resources to manage IAM policy for Cloud Run Service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceIAMPolicySpec   `json:"spec"`
	Status            ServiceIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIAMPolicyList contains a list of ServiceIAMPolicys
type ServiceIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	ServiceIAMPolicy_Kind             = "ServiceIAMPolicy"
	ServiceIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceIAMPolicy_Kind}.String()
	ServiceIAMPolicy_KindAPIVersion   = ServiceIAMPolicy_Kind + "." + CRDGroupVersion.String()
	ServiceIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ServiceIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceIAMPolicy{}, &ServiceIAMPolicyList{})
}
