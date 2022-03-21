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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Index) DeepCopyInto(out *Index) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Index.
func (in *Index) DeepCopy() *Index {
	if in == nil {
		return nil
	}
	out := new(Index)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Index) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexList) DeepCopyInto(out *IndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Index, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexList.
func (in *IndexList) DeepCopy() *IndexList {
	if in == nil {
		return nil
	}
	out := new(IndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexObservation) DeepCopyInto(out *IndexObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IndexID != nil {
		in, out := &in.IndexID, &out.IndexID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexObservation.
func (in *IndexObservation) DeepCopy() *IndexObservation {
	if in == nil {
		return nil
	}
	out := new(IndexObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexParameters) DeepCopyInto(out *IndexParameters) {
	*out = *in
	if in.Analyzer != nil {
		in, out := &in.Analyzer, &out.Analyzer
		*out = new(string)
		**out = **in
	}
	if in.Analyzers != nil {
		in, out := &in.Analyzers, &out.Analyzers
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CollectionName != nil {
		in, out := &in.CollectionName, &out.CollectionName
		*out = new(string)
		**out = **in
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.MappingsDynamic != nil {
		in, out := &in.MappingsDynamic, &out.MappingsDynamic
		*out = new(bool)
		**out = **in
	}
	if in.MappingsFields != nil {
		in, out := &in.MappingsFields, &out.MappingsFields
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.SearchAnalyzer != nil {
		in, out := &in.SearchAnalyzer, &out.SearchAnalyzer
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Synonyms != nil {
		in, out := &in.Synonyms, &out.Synonyms
		*out = make([]SynonymsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexParameters.
func (in *IndexParameters) DeepCopy() *IndexParameters {
	if in == nil {
		return nil
	}
	out := new(IndexParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexSpec) DeepCopyInto(out *IndexSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexSpec.
func (in *IndexSpec) DeepCopy() *IndexSpec {
	if in == nil {
		return nil
	}
	out := new(IndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexStatus) DeepCopyInto(out *IndexStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexStatus.
func (in *IndexStatus) DeepCopy() *IndexStatus {
	if in == nil {
		return nil
	}
	out := new(IndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SynonymsObservation) DeepCopyInto(out *SynonymsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SynonymsObservation.
func (in *SynonymsObservation) DeepCopy() *SynonymsObservation {
	if in == nil {
		return nil
	}
	out := new(SynonymsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SynonymsParameters) DeepCopyInto(out *SynonymsParameters) {
	*out = *in
	if in.Analyzer != nil {
		in, out := &in.Analyzer, &out.Analyzer
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SourceCollection != nil {
		in, out := &in.SourceCollection, &out.SourceCollection
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SynonymsParameters.
func (in *SynonymsParameters) DeepCopy() *SynonymsParameters {
	if in == nil {
		return nil
	}
	out := new(SynonymsParameters)
	in.DeepCopyInto(out)
	return out
}
