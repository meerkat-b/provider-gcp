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

type NotificationInitParameters struct {

	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: "OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_DELETE", "OBJECT_ARCHIVE"
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix *string `json:"objectNamePrefix,omitempty" tf:"object_name_prefix,omitempty"`

	// The desired content of the Payload. One of "JSON_API_V1" or "NONE".
	PayloadFormat *string `json:"payloadFormat,omitempty" tf:"payload_format,omitempty"`
}

type NotificationObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: "OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_DELETE", "OBJECT_ARCHIVE"
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the created notification.
	NotificationID *string `json:"notificationId,omitempty" tf:"notification_id,omitempty"`

	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix *string `json:"objectNamePrefix,omitempty" tf:"object_name_prefix,omitempty"`

	// The desired content of the Payload. One of "JSON_API_V1" or "NONE".
	PayloadFormat *string `json:"payloadFormat,omitempty" tf:"payload_format,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. projects/my-gcp-project/topics/my-topic or my-topic. If the project is not set in the provider,
	// you will need to use the project-level name.
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type NotificationParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: "OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_DELETE", "OBJECT_ARCHIVE"
	// +kubebuilder:validation:Optional
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	// +kubebuilder:validation:Optional
	ObjectNamePrefix *string `json:"objectNamePrefix,omitempty" tf:"object_name_prefix,omitempty"`

	// The desired content of the Payload. One of "JSON_API_V1" or "NONE".
	// +kubebuilder:validation:Optional
	PayloadFormat *string `json:"payloadFormat,omitempty" tf:"payload_format,omitempty"`

	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. projects/my-gcp-project/topics/my-topic or my-topic. If the project is not set in the provider,
	// you will need to use the project-level name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Reference to a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

// NotificationSpec defines the desired state of Notification
type NotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationParameters `json:"forProvider"`
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
	InitProvider NotificationInitParameters `json:"initProvider,omitempty"`
}

// NotificationStatus defines the observed state of Notification.
type NotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Notification is the Schema for the Notifications API. Creates a new notification configuration on a specified bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.payloadFormat) || (has(self.initProvider) && has(self.initProvider.payloadFormat))",message="spec.forProvider.payloadFormat is a required parameter"
	Spec   NotificationSpec   `json:"spec"`
	Status NotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationList contains a list of Notifications
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

// Repository type metadata.
var (
	Notification_Kind             = "Notification"
	Notification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notification_Kind}.String()
	Notification_KindAPIVersion   = Notification_Kind + "." + CRDGroupVersion.String()
	Notification_GroupVersionKind = CRDGroupVersion.WithKind(Notification_Kind)
)

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
