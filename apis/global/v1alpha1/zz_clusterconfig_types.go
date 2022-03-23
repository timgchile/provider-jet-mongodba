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

type ClusterConfigObservation struct {
	CustomZoneMapping map[string]*string `json:"customZoneMapping,omitempty" tf:"custom_zone_mapping,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterConfigParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	CustomZoneMappings []CustomZoneMappingsParameters `json:"customZoneMappings,omitempty" tf:"custom_zone_mappings,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedNamespaces []ManagedNamespacesParameters `json:"managedNamespaces,omitempty" tf:"managed_namespaces,omitempty"`

	// +crossplane:generate:reference:type=github.com/timgchile/provider-jet-mongodba/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/timgchile/provider-jet-mongodba/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type CustomZoneMappingsObservation struct {
}

type CustomZoneMappingsParameters struct {

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ManagedNamespacesObservation struct {
}

type ManagedNamespacesParameters struct {

	// +kubebuilder:validation:Required
	Collection *string `json:"collection" tf:"collection,omitempty"`

	// +kubebuilder:validation:Required
	CustomShardKey *string `json:"customShardKey" tf:"custom_shard_key,omitempty"`

	// +kubebuilder:validation:Required
	DB *string `json:"db" tf:"db,omitempty"`

	// +kubebuilder:validation:Optional
	IsCustomShardKeyHashed *bool `json:"isCustomShardKeyHashed,omitempty" tf:"is_custom_shard_key_hashed,omitempty"`

	// +kubebuilder:validation:Optional
	IsShardKeyUnique *bool `json:"isShardKeyUnique,omitempty" tf:"is_shard_key_unique,omitempty"`
}

// ClusterConfigSpec defines the desired state of ClusterConfig
type ClusterConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterConfigParameters `json:"forProvider"`
}

// ClusterConfigStatus defines the observed state of ClusterConfig.
type ClusterConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterConfig is the Schema for the ClusterConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbajet}
type ClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterConfigSpec   `json:"spec"`
	Status            ClusterConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterConfigList contains a list of ClusterConfigs
type ClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterConfig `json:"items"`
}

// Repository type metadata.
var (
	ClusterConfig_Kind             = "ClusterConfig"
	ClusterConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterConfig_Kind}.String()
	ClusterConfig_KindAPIVersion   = ClusterConfig_Kind + "." + CRDGroupVersion.String()
	ClusterConfig_GroupVersionKind = CRDGroupVersion.WithKind(ClusterConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterConfig{}, &ClusterConfigList{})
}
