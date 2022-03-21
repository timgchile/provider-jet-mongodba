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

type APIKeysObservation struct {
}

type APIKeysParameters struct {

	// +kubebuilder:validation:Required
	APIKeyID *string `json:"apiKeyId" tf:"api_key_id,omitempty"`

	// +kubebuilder:validation:Required
	RoleNames []*string `json:"roleNames" tf:"role_names,omitempty"`
}

type ProjectObservation struct {
	ClusterCount *float64 `json:"clusterCount,omitempty" tf:"cluster_count,omitempty"`

	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProjectParameters struct {

	// +kubebuilder:validation:Optional
	APIKeys []APIKeysParameters `json:"apiKeys,omitempty" tf:"api_keys,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	OrgID *string `json:"orgId" tf:"org_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectOwnerID *string `json:"projectOwnerId,omitempty" tf:"project_owner_id,omitempty"`

	// +kubebuilder:validation:Optional
	Teams []TeamsParameters `json:"teams,omitempty" tf:"teams,omitempty"`

	// +kubebuilder:validation:Optional
	WithDefaultAlertsSettings *bool `json:"withDefaultAlertsSettings,omitempty" tf:"with_default_alerts_settings,omitempty"`
}

type TeamsObservation struct {
}

type TeamsParameters struct {

	// +kubebuilder:validation:Required
	RoleNames []*string `json:"roleNames" tf:"role_names,omitempty"`

	// +kubebuilder:validation:Required
	TeamID *string `json:"teamId" tf:"team_id,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbajet}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
