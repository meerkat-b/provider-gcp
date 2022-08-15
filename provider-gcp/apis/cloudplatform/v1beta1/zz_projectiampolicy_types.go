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

type ProjectIAMPolicyObservation struct {

	// The etag of the project's IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProjectIAMPolicyParameters struct {

	// The google_iam_policy data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// The project id of the target project. This is not
	// inferred from the provider.
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

// ProjectIAMPolicySpec defines the desired state of ProjectIAMPolicy
type ProjectIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectIAMPolicyParameters `json:"forProvider"`
}

// ProjectIAMPolicyStatus defines the observed state of ProjectIAMPolicy.
type ProjectIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMPolicy is the Schema for the ProjectIAMPolicys API. Collection of resources to manage IAM policy for a project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIAMPolicySpec   `json:"spec"`
	Status            ProjectIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMPolicyList contains a list of ProjectIAMPolicys
type ProjectIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	ProjectIAMPolicy_Kind             = "ProjectIAMPolicy"
	ProjectIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectIAMPolicy_Kind}.String()
	ProjectIAMPolicy_KindAPIVersion   = ProjectIAMPolicy_Kind + "." + CRDGroupVersion.String()
	ProjectIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ProjectIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectIAMPolicy{}, &ProjectIAMPolicyList{})
}
