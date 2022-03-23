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
func (in *AwsEventbridgeObservation) DeepCopyInto(out *AwsEventbridgeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsEventbridgeObservation.
func (in *AwsEventbridgeObservation) DeepCopy() *AwsEventbridgeObservation {
	if in == nil {
		return nil
	}
	out := new(AwsEventbridgeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsEventbridgeParameters) DeepCopyInto(out *AwsEventbridgeParameters) {
	*out = *in
	if in.ConfigAccountID != nil {
		in, out := &in.ConfigAccountID, &out.ConfigAccountID
		*out = new(string)
		**out = **in
	}
	if in.ConfigRegion != nil {
		in, out := &in.ConfigRegion, &out.ConfigRegion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsEventbridgeParameters.
func (in *AwsEventbridgeParameters) DeepCopy() *AwsEventbridgeParameters {
	if in == nil {
		return nil
	}
	out := new(AwsEventbridgeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventProcessorsObservation) DeepCopyInto(out *EventProcessorsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventProcessorsObservation.
func (in *EventProcessorsObservation) DeepCopy() *EventProcessorsObservation {
	if in == nil {
		return nil
	}
	out := new(EventProcessorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventProcessorsParameters) DeepCopyInto(out *EventProcessorsParameters) {
	*out = *in
	if in.AwsEventbridge != nil {
		in, out := &in.AwsEventbridge, &out.AwsEventbridge
		*out = make([]AwsEventbridgeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventProcessorsParameters.
func (in *EventProcessorsParameters) DeepCopy() *EventProcessorsParameters {
	if in == nil {
		return nil
	}
	out := new(EventProcessorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerObservation) DeepCopyInto(out *TriggerObservation) {
	*out = *in
	if in.ConfigScheduleType != nil {
		in, out := &in.ConfigScheduleType, &out.ConfigScheduleType
		*out = new(string)
		**out = **in
	}
	if in.FunctionName != nil {
		in, out := &in.FunctionName, &out.FunctionName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TriggerID != nil {
		in, out := &in.TriggerID, &out.TriggerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerObservation.
func (in *TriggerObservation) DeepCopy() *TriggerObservation {
	if in == nil {
		return nil
	}
	out := new(TriggerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerParameters) DeepCopyInto(out *TriggerParameters) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.ConfigCollection != nil {
		in, out := &in.ConfigCollection, &out.ConfigCollection
		*out = new(string)
		**out = **in
	}
	if in.ConfigDatabase != nil {
		in, out := &in.ConfigDatabase, &out.ConfigDatabase
		*out = new(string)
		**out = **in
	}
	if in.ConfigFullDocument != nil {
		in, out := &in.ConfigFullDocument, &out.ConfigFullDocument
		*out = new(bool)
		**out = **in
	}
	if in.ConfigFullDocumentBefore != nil {
		in, out := &in.ConfigFullDocumentBefore, &out.ConfigFullDocumentBefore
		*out = new(bool)
		**out = **in
	}
	if in.ConfigMatch != nil {
		in, out := &in.ConfigMatch, &out.ConfigMatch
		*out = new(string)
		**out = **in
	}
	if in.ConfigOperationType != nil {
		in, out := &in.ConfigOperationType, &out.ConfigOperationType
		*out = new(string)
		**out = **in
	}
	if in.ConfigOperationTypes != nil {
		in, out := &in.ConfigOperationTypes, &out.ConfigOperationTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConfigProject != nil {
		in, out := &in.ConfigProject, &out.ConfigProject
		*out = new(string)
		**out = **in
	}
	if in.ConfigProviders != nil {
		in, out := &in.ConfigProviders, &out.ConfigProviders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConfigSchedule != nil {
		in, out := &in.ConfigSchedule, &out.ConfigSchedule
		*out = new(string)
		**out = **in
	}
	if in.ConfigServiceID != nil {
		in, out := &in.ConfigServiceID, &out.ConfigServiceID
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.EventProcessors != nil {
		in, out := &in.EventProcessors, &out.EventProcessors
		*out = make([]EventProcessorsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerParameters.
func (in *TriggerParameters) DeepCopy() *TriggerParameters {
	if in == nil {
		return nil
	}
	out := new(TriggerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}
