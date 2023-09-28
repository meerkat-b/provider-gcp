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

type AttestationAuthorityInitParameters struct {

	// This submessage provides human-readable hints about the purpose of
	// the AttestationAuthority. Because the name of a Note acts as its
	// resource reference, it is important to disambiguate the canonical
	// name of the Note (which might be a UUID for security purposes)
	// from "readable" names more suitable for debug output. Note that
	// these hints should NOT be used to look up AttestationAuthorities
	// in security sensitive contexts, such as when looking up
	// Attestations to verify.
	// Structure is documented below.
	Hint []HintInitParameters `json:"hint,omitempty" tf:"hint,omitempty"`
}

type AttestationAuthorityObservation struct {

	// This submessage provides human-readable hints about the purpose of
	// the AttestationAuthority. Because the name of a Note acts as its
	// resource reference, it is important to disambiguate the canonical
	// name of the Note (which might be a UUID for security purposes)
	// from "readable" names more suitable for debug output. Note that
	// these hints should NOT be used to look up AttestationAuthorities
	// in security sensitive contexts, such as when looking up
	// Attestations to verify.
	// Structure is documented below.
	Hint []HintObservation `json:"hint,omitempty" tf:"hint,omitempty"`
}

type AttestationAuthorityParameters struct {

	// This submessage provides human-readable hints about the purpose of
	// the AttestationAuthority. Because the name of a Note acts as its
	// resource reference, it is important to disambiguate the canonical
	// name of the Note (which might be a UUID for security purposes)
	// from "readable" names more suitable for debug output. Note that
	// these hints should NOT be used to look up AttestationAuthorities
	// in security sensitive contexts, such as when looking up
	// Attestations to verify.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Hint []HintParameters `json:"hint" tf:"hint,omitempty"`
}

type HintInitParameters struct {

	// The human readable name of this Attestation Authority, for
	// example "qa".
	HumanReadableName *string `json:"humanReadableName,omitempty" tf:"human_readable_name,omitempty"`
}

type HintObservation struct {

	// The human readable name of this Attestation Authority, for
	// example "qa".
	HumanReadableName *string `json:"humanReadableName,omitempty" tf:"human_readable_name,omitempty"`
}

type HintParameters struct {

	// The human readable name of this Attestation Authority, for
	// example "qa".
	// +kubebuilder:validation:Optional
	HumanReadableName *string `json:"humanReadableName" tf:"human_readable_name,omitempty"`
}

type NoteInitParameters struct {

	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.
	// Structure is documented below.
	AttestationAuthority []AttestationAuthorityInitParameters `json:"attestationAuthority,omitempty" tf:"attestation_authority,omitempty"`

	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// A detailed description of the note
	LongDescription *string `json:"longDescription,omitempty" tf:"long_description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Names of other notes related to this note.
	RelatedNoteNames []*string `json:"relatedNoteNames,omitempty" tf:"related_note_names,omitempty"`

	// URLs associated with this note and related metadata.
	// Structure is documented below.
	RelatedURL []RelatedURLInitParameters `json:"relatedUrl,omitempty" tf:"related_url,omitempty"`

	// A one sentence description of the note.
	ShortDescription *string `json:"shortDescription,omitempty" tf:"short_description,omitempty"`
}

type NoteObservation struct {

	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.
	// Structure is documented below.
	AttestationAuthority []AttestationAuthorityObservation `json:"attestationAuthority,omitempty" tf:"attestation_authority,omitempty"`

	// The time this note was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/notes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of analysis this note describes
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// A detailed description of the note
	LongDescription *string `json:"longDescription,omitempty" tf:"long_description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Names of other notes related to this note.
	RelatedNoteNames []*string `json:"relatedNoteNames,omitempty" tf:"related_note_names,omitempty"`

	// URLs associated with this note and related metadata.
	// Structure is documented below.
	RelatedURL []RelatedURLObservation `json:"relatedUrl,omitempty" tf:"related_url,omitempty"`

	// A one sentence description of the note.
	ShortDescription *string `json:"shortDescription,omitempty" tf:"short_description,omitempty"`

	// The time this note was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type NoteParameters struct {

	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AttestationAuthority []AttestationAuthorityParameters `json:"attestationAuthority,omitempty" tf:"attestation_authority,omitempty"`

	// Time of expiration for this note. Leave empty if note does not expire.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// A detailed description of the note
	// +kubebuilder:validation:Optional
	LongDescription *string `json:"longDescription,omitempty" tf:"long_description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Names of other notes related to this note.
	// +kubebuilder:validation:Optional
	RelatedNoteNames []*string `json:"relatedNoteNames,omitempty" tf:"related_note_names,omitempty"`

	// URLs associated with this note and related metadata.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RelatedURL []RelatedURLParameters `json:"relatedUrl,omitempty" tf:"related_url,omitempty"`

	// A one sentence description of the note.
	// +kubebuilder:validation:Optional
	ShortDescription *string `json:"shortDescription,omitempty" tf:"short_description,omitempty"`
}

type RelatedURLInitParameters struct {

	// Label to describe usage of the URL
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Specific URL associated with the resource.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type RelatedURLObservation struct {

	// Label to describe usage of the URL
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Specific URL associated with the resource.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type RelatedURLParameters struct {

	// Label to describe usage of the URL
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Specific URL associated with the resource.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

// NoteSpec defines the desired state of Note
type NoteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NoteParameters `json:"forProvider"`
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
	InitProvider NoteInitParameters `json:"initProvider,omitempty"`
}

// NoteStatus defines the observed state of Note.
type NoteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NoteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Note is the Schema for the Notes API. A Container Analysis note is a high-level piece of metadata that describes a type of analysis that can be done for a resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Note struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.attestationAuthority) || (has(self.initProvider) && has(self.initProvider.attestationAuthority))",message="spec.forProvider.attestationAuthority is a required parameter"
	Spec   NoteSpec   `json:"spec"`
	Status NoteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NoteList contains a list of Notes
type NoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Note `json:"items"`
}

// Repository type metadata.
var (
	Note_Kind             = "Note"
	Note_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Note_Kind}.String()
	Note_KindAPIVersion   = Note_Kind + "." + CRDGroupVersion.String()
	Note_GroupVersionKind = CRDGroupVersion.WithKind(Note_Kind)
)

func init() {
	SchemeBuilder.Register(&Note{}, &NoteList{})
}
