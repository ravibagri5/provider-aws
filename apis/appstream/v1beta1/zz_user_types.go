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

type UserObservation struct {

	// ARN of the appstream user.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date and time, in UTC and extended RFC 3339 format, when the user was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// Unique ID of the appstream user.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserParameters struct {

	// Authentication type for the user. You must specify USERPOOL. Valid values: API, SAML, USERPOOL
	// +kubebuilder:validation:Required
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type,omitempty"`

	// Specifies whether the user in the user pool is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// First name, or given name, of the user.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last name, or surname, of the user.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Send an email notification.
	// +kubebuilder:validation:Optional
	SendEmailNotification *bool `json:"sendEmailNotification,omitempty" tf:"send_email_notification,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Provides an AppStream user
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
