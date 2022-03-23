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

type ContainerObservation struct {
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty" tf:"azure_subscription_id,omitempty"`

	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	GCPProjectID *string `json:"gcpProjectId,omitempty" tf:"gcp_project_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	Provisioned *bool `json:"provisioned,omitempty" tf:"provisioned,omitempty"`

	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VnetName *string `json:"vnetName,omitempty" tf:"vnet_name,omitempty"`
}

type ContainerParameters struct {

	// +kubebuilder:validation:Required
	AtlasCidrBlock *string `json:"atlasCidrBlock" tf:"atlas_cidr_block,omitempty"`

	// +crossplane:generate:reference:type=github.com/timgchile/provider-jet-mongodba/apis/mongodba/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/timgchile/provider-jet-mongodba/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerParameters `json:"forProvider"`
}

// ContainerStatus defines the observed state of Container.
type ContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Container is the Schema for the Containers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbajet}
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerSpec   `json:"spec"`
	Status            ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Containers
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

// Repository type metadata.
var (
	Container_Kind             = "Container"
	Container_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Container_Kind}.String()
	Container_KindAPIVersion   = Container_Kind + "." + CRDGroupVersion.String()
	Container_GroupVersionKind = CRDGroupVersion.WithKind(Container_Kind)
)

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
