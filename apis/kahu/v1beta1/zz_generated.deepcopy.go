//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupCondition) DeepCopyInto(out *BackupCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupCondition.
func (in *BackupCondition) DeepCopy() *BackupCondition {
	if in == nil {
		return nil
	}
	out := new(BackupCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocation) DeepCopyInto(out *BackupLocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocation.
func (in *BackupLocation) DeepCopy() *BackupLocation {
	if in == nil {
		return nil
	}
	out := new(BackupLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationList) DeepCopyInto(out *BackupLocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupLocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationList.
func (in *BackupLocationList) DeepCopy() *BackupLocationList {
	if in == nil {
		return nil
	}
	out := new(BackupLocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationSpec) DeepCopyInto(out *BackupLocationSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationSpec.
func (in *BackupLocationSpec) DeepCopy() *BackupLocationSpec {
	if in == nil {
		return nil
	}
	out := new(BackupLocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	out.ReclaimPolicy = in.ReclaimPolicy
	if in.MetadataLocation != nil {
		in, out := &in.MetadataLocation, &out.MetadataLocation
		*out = new(BackupLocation)
		(*in).DeepCopyInto(*out)
	}
	in.PreExecHook.DeepCopyInto(&out.PreExecHook)
	in.PostExecHook.DeepCopyInto(&out.PostExecHook)
	if in.IncludedProviders != nil {
		in, out := &in.IncludedProviders, &out.IncludedProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeProviders != nil {
		in, out := &in.ExcludeProviders, &out.ExcludeProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedNamespaces != nil {
		in, out := &in.IncludedNamespaces, &out.IncludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedNamespaces != nil {
		in, out := &in.ExcludedNamespaces, &out.ExcludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedResources != nil {
		in, out := &in.IncludedResources, &out.IncludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	if in.LastBackup != nil {
		in, out := &in.LastBackup, &out.LastBackup
		*out = (*in).DeepCopy()
	}
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]BackupCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimPolicyType) DeepCopyInto(out *ReclaimPolicyType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimPolicyType.
func (in *ReclaimPolicyType) DeepCopy() *ReclaimPolicyType {
	if in == nil {
		return nil
	}
	out := new(ReclaimPolicyType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceHookSpec) DeepCopyInto(out *ResourceHookSpec) {
	*out = *in
	if in.IncludedNamespaces != nil {
		in, out := &in.IncludedNamespaces, &out.IncludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedNamespaces != nil {
		in, out := &in.ExcludedNamespaces, &out.ExcludedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedResources != nil {
		in, out := &in.IncludedResources, &out.IncludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceHookSpec.
func (in *ResourceHookSpec) DeepCopy() *ResourceHookSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceHookSpec)
	in.DeepCopyInto(out)
	return out
}
