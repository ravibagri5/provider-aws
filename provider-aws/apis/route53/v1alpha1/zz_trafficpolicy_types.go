/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TrafficPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type TrafficPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	Document *string `json:"document" tf:"document,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// TrafficPolicySpec defines the desired state of TrafficPolicy
type TrafficPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficPolicyParameters `json:"forProvider"`
}

// TrafficPolicyStatus defines the observed state of TrafficPolicy.
type TrafficPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficPolicy is the Schema for the TrafficPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TrafficPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficPolicySpec   `json:"spec"`
	Status            TrafficPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficPolicyList contains a list of TrafficPolicys
type TrafficPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficPolicy `json:"items"`
}

// Repository type metadata.
var (
	TrafficPolicy_Kind             = "TrafficPolicy"
	TrafficPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficPolicy_Kind}.String()
	TrafficPolicy_KindAPIVersion   = TrafficPolicy_Kind + "." + CRDGroupVersion.String()
	TrafficPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TrafficPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficPolicy{}, &TrafficPolicyList{})
}