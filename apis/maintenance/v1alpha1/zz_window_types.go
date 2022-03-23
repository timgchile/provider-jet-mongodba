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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WindowObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StartAsap *bool `json:"startAsap,omitempty" tf:"start_asap,omitempty"`
}

type WindowParameters struct {

	// +kubebuilder:validation:Optional
	AutoDefer *bool `json:"autoDefer,omitempty" tf:"auto_defer,omitempty"`

	// +kubebuilder:validation:Optional
	AutoDeferOnceEnabled *bool `json:"autoDeferOnceEnabled,omitempty" tf:"auto_defer_once_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	Defer *bool `json:"defer,omitempty" tf:"defer,omitempty"`

	// +kubebuilder:validation:Optional
	HourOfDay *float64 `json:"hourOfDay,omitempty" tf:"hour_of_day,omitempty"`

	// +kubebuilder:validation:Optional
	NumberOfDeferrals *float64 `json:"numberOfDeferrals,omitempty" tf:"number_of_deferrals,omitempty"`

	// +crossplane:generate:reference:type=github.com/timgchile/provider-jet-mongodba/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/timgchile/provider-jet-mongodba/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

// WindowSpec defines the desired state of Window
type WindowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WindowParameters `json:"forProvider"`
}

// WindowStatus defines the observed state of Window.
type WindowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WindowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Window is the Schema for the Windows API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbajet}
type Window struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WindowSpec   `json:"spec"`
	Status            WindowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WindowList contains a list of Windows
type WindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Window `json:"items"`
}

// Repository type metadata.
var (
	Window_Kind             = "Window"
	Window_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Window_Kind}.String()
	Window_KindAPIVersion   = Window_Kind + "." + CRDGroupVersion.String()
	Window_GroupVersionKind = CRDGroupVersion.WithKind(Window_Kind)
)

func init() {
	SchemeBuilder.Register(&Window{}, &WindowList{})
}
