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

type WebBackendServiceIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type WebBackendServiceIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type WebBackendServiceIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type WebBackendServiceIAMMemberInitParameters struct {
	Condition []WebBackendServiceIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type WebBackendServiceIAMMemberObservation struct {
	Condition []WebBackendServiceIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	WebBackendService *string `json:"webBackendService,omitempty" tf:"web_backend_service,omitempty"`
}

type WebBackendServiceIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []WebBackendServiceIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.BackendService
	// +kubebuilder:validation:Optional
	WebBackendService *string `json:"webBackendService,omitempty" tf:"web_backend_service,omitempty"`

	// Reference to a BackendService in compute to populate webBackendService.
	// +kubebuilder:validation:Optional
	WebBackendServiceRef *v1.Reference `json:"webBackendServiceRef,omitempty" tf:"-"`

	// Selector for a BackendService in compute to populate webBackendService.
	// +kubebuilder:validation:Optional
	WebBackendServiceSelector *v1.Selector `json:"webBackendServiceSelector,omitempty" tf:"-"`
}

// WebBackendServiceIAMMemberSpec defines the desired state of WebBackendServiceIAMMember
type WebBackendServiceIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebBackendServiceIAMMemberParameters `json:"forProvider"`
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
	InitProvider WebBackendServiceIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// WebBackendServiceIAMMemberStatus defines the observed state of WebBackendServiceIAMMember.
type WebBackendServiceIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebBackendServiceIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebBackendServiceIAMMember is the Schema for the WebBackendServiceIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type WebBackendServiceIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   WebBackendServiceIAMMemberSpec   `json:"spec"`
	Status WebBackendServiceIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebBackendServiceIAMMemberList contains a list of WebBackendServiceIAMMembers
type WebBackendServiceIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebBackendServiceIAMMember `json:"items"`
}

// Repository type metadata.
var (
	WebBackendServiceIAMMember_Kind             = "WebBackendServiceIAMMember"
	WebBackendServiceIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebBackendServiceIAMMember_Kind}.String()
	WebBackendServiceIAMMember_KindAPIVersion   = WebBackendServiceIAMMember_Kind + "." + CRDGroupVersion.String()
	WebBackendServiceIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(WebBackendServiceIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&WebBackendServiceIAMMember{}, &WebBackendServiceIAMMemberList{})
}
