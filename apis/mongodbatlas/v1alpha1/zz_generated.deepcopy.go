//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedConfigurationObservation) DeepCopyInto(out *AdvancedConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedConfigurationObservation.
func (in *AdvancedConfigurationObservation) DeepCopy() *AdvancedConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AdvancedConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedConfigurationParameters) DeepCopyInto(out *AdvancedConfigurationParameters) {
	*out = *in
	if in.DefaultReadConcern != nil {
		in, out := &in.DefaultReadConcern, &out.DefaultReadConcern
		*out = new(string)
		**out = **in
	}
	if in.DefaultWriteConcern != nil {
		in, out := &in.DefaultWriteConcern, &out.DefaultWriteConcern
		*out = new(string)
		**out = **in
	}
	if in.FailIndexKeyTooLong != nil {
		in, out := &in.FailIndexKeyTooLong, &out.FailIndexKeyTooLong
		*out = new(bool)
		**out = **in
	}
	if in.JavascriptEnabled != nil {
		in, out := &in.JavascriptEnabled, &out.JavascriptEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MinimumEnabledTLSProtocol != nil {
		in, out := &in.MinimumEnabledTLSProtocol, &out.MinimumEnabledTLSProtocol
		*out = new(string)
		**out = **in
	}
	if in.NoTableScan != nil {
		in, out := &in.NoTableScan, &out.NoTableScan
		*out = new(bool)
		**out = **in
	}
	if in.OplogSizeMb != nil {
		in, out := &in.OplogSizeMb, &out.OplogSizeMb
		*out = new(float64)
		**out = **in
	}
	if in.SampleRefreshIntervalBiConnector != nil {
		in, out := &in.SampleRefreshIntervalBiConnector, &out.SampleRefreshIntervalBiConnector
		*out = new(float64)
		**out = **in
	}
	if in.SampleSizeBiConnector != nil {
		in, out := &in.SampleSizeBiConnector, &out.SampleSizeBiConnector
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedConfigurationParameters.
func (in *AdvancedConfigurationParameters) DeepCopy() *AdvancedConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AdvancedConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auditing) DeepCopyInto(out *Auditing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auditing.
func (in *Auditing) DeepCopy() *Auditing {
	if in == nil {
		return nil
	}
	out := new(Auditing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Auditing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditingList) DeepCopyInto(out *AuditingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Auditing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditingList.
func (in *AuditingList) DeepCopy() *AuditingList {
	if in == nil {
		return nil
	}
	out := new(AuditingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuditingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditingObservation) DeepCopyInto(out *AuditingObservation) {
	*out = *in
	if in.ConfigurationType != nil {
		in, out := &in.ConfigurationType, &out.ConfigurationType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditingObservation.
func (in *AuditingObservation) DeepCopy() *AuditingObservation {
	if in == nil {
		return nil
	}
	out := new(AuditingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditingParameters) DeepCopyInto(out *AuditingParameters) {
	*out = *in
	if in.AuditAuthorizationSuccess != nil {
		in, out := &in.AuditAuthorizationSuccess, &out.AuditAuthorizationSuccess
		*out = new(bool)
		**out = **in
	}
	if in.AuditFilter != nil {
		in, out := &in.AuditFilter, &out.AuditFilter
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditingParameters.
func (in *AuditingParameters) DeepCopy() *AuditingParameters {
	if in == nil {
		return nil
	}
	out := new(AuditingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditingSpec) DeepCopyInto(out *AuditingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditingSpec.
func (in *AuditingSpec) DeepCopy() *AuditingSpec {
	if in == nil {
		return nil
	}
	out := new(AuditingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditingStatus) DeepCopyInto(out *AuditingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditingStatus.
func (in *AuditingStatus) DeepCopy() *AuditingStatus {
	if in == nil {
		return nil
	}
	out := new(AuditingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BiConnectorConfigObservation) DeepCopyInto(out *BiConnectorConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BiConnectorConfigObservation.
func (in *BiConnectorConfigObservation) DeepCopy() *BiConnectorConfigObservation {
	if in == nil {
		return nil
	}
	out := new(BiConnectorConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BiConnectorConfigParameters) DeepCopyInto(out *BiConnectorConfigParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ReadPreference != nil {
		in, out := &in.ReadPreference, &out.ReadPreference
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BiConnectorConfigParameters.
func (in *BiConnectorConfigParameters) DeepCopy() *BiConnectorConfigParameters {
	if in == nil {
		return nil
	}
	out := new(BiConnectorConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ConnectionStrings != nil {
		in, out := &in.ConnectionStrings, &out.ConnectionStrings
		*out = make([]ConnectionStringsObservation, len(*in))
		copy(*out, *in)
	}
	if in.ContainerID != nil {
		in, out := &in.ContainerID, &out.ContainerID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MongoDBVersion != nil {
		in, out := &in.MongoDBVersion, &out.MongoDBVersion
		*out = new(string)
		**out = **in
	}
	if in.MongoURI != nil {
		in, out := &in.MongoURI, &out.MongoURI
		*out = new(string)
		**out = **in
	}
	if in.MongoURIUpdated != nil {
		in, out := &in.MongoURIUpdated, &out.MongoURIUpdated
		*out = new(string)
		**out = **in
	}
	if in.MongoURIWithOptions != nil {
		in, out := &in.MongoURIWithOptions, &out.MongoURIWithOptions
		*out = new(string)
		**out = **in
	}
	if in.ProviderEncryptEBSVolumeFlag != nil {
		in, out := &in.ProviderEncryptEBSVolumeFlag, &out.ProviderEncryptEBSVolumeFlag
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotBackupPolicy != nil {
		in, out := &in.SnapshotBackupPolicy, &out.SnapshotBackupPolicy
		*out = make([]SnapshotBackupPolicyObservation, len(*in))
		copy(*out, *in)
	}
	if in.SrvAddress != nil {
		in, out := &in.SrvAddress, &out.SrvAddress
		*out = new(string)
		**out = **in
	}
	if in.StateName != nil {
		in, out := &in.StateName, &out.StateName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.AdvancedConfiguration != nil {
		in, out := &in.AdvancedConfiguration, &out.AdvancedConfiguration
		*out = make([]AdvancedConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoScalingComputeEnabled != nil {
		in, out := &in.AutoScalingComputeEnabled, &out.AutoScalingComputeEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutoScalingComputeScaleDownEnabled != nil {
		in, out := &in.AutoScalingComputeScaleDownEnabled, &out.AutoScalingComputeScaleDownEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutoScalingDiskGbEnabled != nil {
		in, out := &in.AutoScalingDiskGbEnabled, &out.AutoScalingDiskGbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BackingProviderName != nil {
		in, out := &in.BackingProviderName, &out.BackingProviderName
		*out = new(string)
		**out = **in
	}
	if in.BackupEnabled != nil {
		in, out := &in.BackupEnabled, &out.BackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BiConnector != nil {
		in, out := &in.BiConnector, &out.BiConnector
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.BiConnectorConfig != nil {
		in, out := &in.BiConnectorConfig, &out.BiConnectorConfig
		*out = make([]BiConnectorConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudBackup != nil {
		in, out := &in.CloudBackup, &out.CloudBackup
		*out = new(bool)
		**out = **in
	}
	if in.ClusterType != nil {
		in, out := &in.ClusterType, &out.ClusterType
		*out = new(string)
		**out = **in
	}
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(float64)
		**out = **in
	}
	if in.EncryptionAtRestProvider != nil {
		in, out := &in.EncryptionAtRestProvider, &out.EncryptionAtRestProvider
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]LabelsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MongoDBMajorVersion != nil {
		in, out := &in.MongoDBMajorVersion, &out.MongoDBMajorVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumShards != nil {
		in, out := &in.NumShards, &out.NumShards
		*out = new(float64)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.PitEnabled != nil {
		in, out := &in.PitEnabled, &out.PitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderAutoScalingComputeMaxInstanceSize != nil {
		in, out := &in.ProviderAutoScalingComputeMaxInstanceSize, &out.ProviderAutoScalingComputeMaxInstanceSize
		*out = new(string)
		**out = **in
	}
	if in.ProviderAutoScalingComputeMinInstanceSize != nil {
		in, out := &in.ProviderAutoScalingComputeMinInstanceSize, &out.ProviderAutoScalingComputeMinInstanceSize
		*out = new(string)
		**out = **in
	}
	if in.ProviderBackupEnabled != nil {
		in, out := &in.ProviderBackupEnabled, &out.ProviderBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProviderDiskIops != nil {
		in, out := &in.ProviderDiskIops, &out.ProviderDiskIops
		*out = new(float64)
		**out = **in
	}
	if in.ProviderDiskTypeName != nil {
		in, out := &in.ProviderDiskTypeName, &out.ProviderDiskTypeName
		*out = new(string)
		**out = **in
	}
	if in.ProviderEncryptEBSVolume != nil {
		in, out := &in.ProviderEncryptEBSVolume, &out.ProviderEncryptEBSVolume
		*out = new(bool)
		**out = **in
	}
	if in.ProviderInstanceSizeName != nil {
		in, out := &in.ProviderInstanceSizeName, &out.ProviderInstanceSizeName
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.ProviderRegionName != nil {
		in, out := &in.ProviderRegionName, &out.ProviderRegionName
		*out = new(string)
		**out = **in
	}
	if in.ProviderVolumeType != nil {
		in, out := &in.ProviderVolumeType, &out.ProviderVolumeType
		*out = new(string)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(float64)
		**out = **in
	}
	if in.ReplicationSpecs != nil {
		in, out := &in.ReplicationSpecs, &out.ReplicationSpecs
		*out = make([]ReplicationSpecsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VersionReleaseSystem != nil {
		in, out := &in.VersionReleaseSystem, &out.VersionReleaseSystem
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStringsObservation) DeepCopyInto(out *ConnectionStringsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStringsObservation.
func (in *ConnectionStringsObservation) DeepCopy() *ConnectionStringsObservation {
	if in == nil {
		return nil
	}
	out := new(ConnectionStringsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStringsParameters) DeepCopyInto(out *ConnectionStringsParameters) {
	*out = *in
	if in.AwsPrivateLink != nil {
		in, out := &in.AwsPrivateLink, &out.AwsPrivateLink
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AwsPrivateLinkSrv != nil {
		in, out := &in.AwsPrivateLinkSrv, &out.AwsPrivateLinkSrv
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = make([]PrivateEndpointParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateSrv != nil {
		in, out := &in.PrivateSrv, &out.PrivateSrv
		*out = new(string)
		**out = **in
	}
	if in.Standard != nil {
		in, out := &in.Standard, &out.Standard
		*out = new(string)
		**out = **in
	}
	if in.StandardSrv != nil {
		in, out := &in.StandardSrv, &out.StandardSrv
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStringsParameters.
func (in *ConnectionStringsParameters) DeepCopy() *ConnectionStringsParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectionStringsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointsObservation) DeepCopyInto(out *EndpointsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointsObservation.
func (in *EndpointsObservation) DeepCopy() *EndpointsObservation {
	if in == nil {
		return nil
	}
	out := new(EndpointsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointsParameters) DeepCopyInto(out *EndpointsParameters) {
	*out = *in
	if in.EndpointID != nil {
		in, out := &in.EndpointID, &out.EndpointID
		*out = new(string)
		**out = **in
	}
	if in.ProviderName != nil {
		in, out := &in.ProviderName, &out.ProviderName
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointsParameters.
func (in *EndpointsParameters) DeepCopy() *EndpointsParameters {
	if in == nil {
		return nil
	}
	out := new(EndpointsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsObservation) DeepCopyInto(out *LabelsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsObservation.
func (in *LabelsObservation) DeepCopy() *LabelsObservation {
	if in == nil {
		return nil
	}
	out := new(LabelsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsParameters) DeepCopyInto(out *LabelsParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsParameters.
func (in *LabelsParameters) DeepCopy() *LabelsParameters {
	if in == nil {
		return nil
	}
	out := new(LabelsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoliciesObservation) DeepCopyInto(out *PoliciesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoliciesObservation.
func (in *PoliciesObservation) DeepCopy() *PoliciesObservation {
	if in == nil {
		return nil
	}
	out := new(PoliciesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoliciesParameters) DeepCopyInto(out *PoliciesParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PolicyItem != nil {
		in, out := &in.PolicyItem, &out.PolicyItem
		*out = make([]PolicyItemParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoliciesParameters.
func (in *PoliciesParameters) DeepCopy() *PoliciesParameters {
	if in == nil {
		return nil
	}
	out := new(PoliciesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemObservation) DeepCopyInto(out *PolicyItemObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemObservation.
func (in *PolicyItemObservation) DeepCopy() *PolicyItemObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyItemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyItemParameters) DeepCopyInto(out *PolicyItemParameters) {
	*out = *in
	if in.FrequencyInterval != nil {
		in, out := &in.FrequencyInterval, &out.FrequencyInterval
		*out = new(float64)
		**out = **in
	}
	if in.FrequencyType != nil {
		in, out := &in.FrequencyType, &out.FrequencyType
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetentionUnit != nil {
		in, out := &in.RetentionUnit, &out.RetentionUnit
		*out = new(string)
		**out = **in
	}
	if in.RetentionValue != nil {
		in, out := &in.RetentionValue, &out.RetentionValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyItemParameters.
func (in *PolicyItemParameters) DeepCopy() *PolicyItemParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyItemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointObservation) DeepCopyInto(out *PrivateEndpointObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointObservation.
func (in *PrivateEndpointObservation) DeepCopy() *PrivateEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointParameters) DeepCopyInto(out *PrivateEndpointParameters) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]EndpointsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SrvConnectionString != nil {
		in, out := &in.SrvConnectionString, &out.SrvConnectionString
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointParameters.
func (in *PrivateEndpointParameters) DeepCopy() *PrivateEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionsConfigObservation) DeepCopyInto(out *RegionsConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionsConfigObservation.
func (in *RegionsConfigObservation) DeepCopy() *RegionsConfigObservation {
	if in == nil {
		return nil
	}
	out := new(RegionsConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionsConfigParameters) DeepCopyInto(out *RegionsConfigParameters) {
	*out = *in
	if in.AnalyticsNodes != nil {
		in, out := &in.AnalyticsNodes, &out.AnalyticsNodes
		*out = new(float64)
		**out = **in
	}
	if in.ElectableNodes != nil {
		in, out := &in.ElectableNodes, &out.ElectableNodes
		*out = new(float64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.ReadOnlyNodes != nil {
		in, out := &in.ReadOnlyNodes, &out.ReadOnlyNodes
		*out = new(float64)
		**out = **in
	}
	if in.RegionName != nil {
		in, out := &in.RegionName, &out.RegionName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionsConfigParameters.
func (in *RegionsConfigParameters) DeepCopy() *RegionsConfigParameters {
	if in == nil {
		return nil
	}
	out := new(RegionsConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSpecsObservation) DeepCopyInto(out *ReplicationSpecsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSpecsObservation.
func (in *ReplicationSpecsObservation) DeepCopy() *ReplicationSpecsObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicationSpecsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSpecsParameters) DeepCopyInto(out *ReplicationSpecsParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NumShards != nil {
		in, out := &in.NumShards, &out.NumShards
		*out = new(float64)
		**out = **in
	}
	if in.RegionsConfig != nil {
		in, out := &in.RegionsConfig, &out.RegionsConfig
		*out = make([]RegionsConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ZoneName != nil {
		in, out := &in.ZoneName, &out.ZoneName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSpecsParameters.
func (in *ReplicationSpecsParameters) DeepCopy() *ReplicationSpecsParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationSpecsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotBackupPolicyObservation) DeepCopyInto(out *SnapshotBackupPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotBackupPolicyObservation.
func (in *SnapshotBackupPolicyObservation) DeepCopy() *SnapshotBackupPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SnapshotBackupPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotBackupPolicyParameters) DeepCopyInto(out *SnapshotBackupPolicyParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.NextSnapshot != nil {
		in, out := &in.NextSnapshot, &out.NextSnapshot
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]PoliciesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReferenceHourOfDay != nil {
		in, out := &in.ReferenceHourOfDay, &out.ReferenceHourOfDay
		*out = new(float64)
		**out = **in
	}
	if in.ReferenceMinuteOfHour != nil {
		in, out := &in.ReferenceMinuteOfHour, &out.ReferenceMinuteOfHour
		*out = new(float64)
		**out = **in
	}
	if in.RestoreWindowDays != nil {
		in, out := &in.RestoreWindowDays, &out.RestoreWindowDays
		*out = new(float64)
		**out = **in
	}
	if in.UpdateSnapshots != nil {
		in, out := &in.UpdateSnapshots, &out.UpdateSnapshots
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotBackupPolicyParameters.
func (in *SnapshotBackupPolicyParameters) DeepCopy() *SnapshotBackupPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SnapshotBackupPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Team) DeepCopyInto(out *Team) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Team.
func (in *Team) DeepCopy() *Team {
	if in == nil {
		return nil
	}
	out := new(Team)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Team) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamList) DeepCopyInto(out *TeamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Team, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamList.
func (in *TeamList) DeepCopy() *TeamList {
	if in == nil {
		return nil
	}
	out := new(TeamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TeamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamObservation) DeepCopyInto(out *TeamObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamObservation.
func (in *TeamObservation) DeepCopy() *TeamObservation {
	if in == nil {
		return nil
	}
	out := new(TeamObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamParameters) DeepCopyInto(out *TeamParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Usernames != nil {
		in, out := &in.Usernames, &out.Usernames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamParameters.
func (in *TeamParameters) DeepCopy() *TeamParameters {
	if in == nil {
		return nil
	}
	out := new(TeamParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamSpec) DeepCopyInto(out *TeamSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamSpec.
func (in *TeamSpec) DeepCopy() *TeamSpec {
	if in == nil {
		return nil
	}
	out := new(TeamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamStatus) DeepCopyInto(out *TeamStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamStatus.
func (in *TeamStatus) DeepCopy() *TeamStatus {
	if in == nil {
		return nil
	}
	out := new(TeamStatus)
	in.DeepCopyInto(out)
	return out
}
