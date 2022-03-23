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

type ProviderSnapshotObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MasterKeyUUID *string `json:"masterKeyUuid,omitempty" tf:"master_key_uuid,omitempty"`

	MongodVersion *string `json:"mongodVersion,omitempty" tf:"mongod_version,omitempty"`

	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StorageSizeBytes *float64 `json:"storageSizeBytes,omitempty" tf:"storage_size_bytes,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProviderSnapshotParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/timgchile/provider-jet-mongodba/apis/mongodba/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/timgchile/provider-jet-mongodba/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	RetentionInDays *float64 `json:"retentionInDays" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// ProviderSnapshotSpec defines the desired state of ProviderSnapshot
type ProviderSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderSnapshotParameters `json:"forProvider"`
}

// ProviderSnapshotStatus defines the observed state of ProviderSnapshot.
type ProviderSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderSnapshot is the Schema for the ProviderSnapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbajet}
type ProviderSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderSnapshotSpec   `json:"spec"`
	Status            ProviderSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderSnapshotList contains a list of ProviderSnapshots
type ProviderSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderSnapshot `json:"items"`
}

// Repository type metadata.
var (
	ProviderSnapshot_Kind             = "ProviderSnapshot"
	ProviderSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProviderSnapshot_Kind}.String()
	ProviderSnapshot_KindAPIVersion   = ProviderSnapshot_Kind + "." + CRDGroupVersion.String()
	ProviderSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(ProviderSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&ProviderSnapshot{}, &ProviderSnapshotList{})
}
