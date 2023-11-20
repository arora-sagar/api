//go:build !ignore_autogenerated

/*
Copyright 2023 Nephio.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMFID) DeepCopyInto(out *AMFID) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMFID.
func (in *AMFID) DeepCopy() *AMFID {
	if in == nil {
		return nil
	}
	out := new(AMFID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	if in.ConfigRefs != nil {
		in, out := &in.ConfigRefs, &out.ConfigRefs
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigStatus) DeepCopyInto(out *ConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigStatus.
func (in *ConfigStatus) DeepCopy() *ConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNNInfo) DeepCopyInto(out *DNNInfo) {
	*out = *in
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNNInfo.
func (in *DNNInfo) DeepCopy() *DNNInfo {
	if in == nil {
		return nil
	}
	out := new(DNNInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSSAI) DeepCopyInto(out *NSSAI) {
	*out = *in
	if in.SD != nil {
		in, out := &in.SD, &out.SD
		*out = new(string)
		**out = **in
	}
	in.DNNInfo.DeepCopyInto(&out.DNNInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSSAI.
func (in *NSSAI) DeepCopy() *NSSAI {
	if in == nil {
		return nil
	}
	out := new(NSSAI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLMN) DeepCopyInto(out *PLMN) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLMN.
func (in *PLMN) DeepCopy() *PLMN {
	if in == nil {
		return nil
	}
	out := new(PLMN)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PLMN) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLMNID) DeepCopyInto(out *PLMNID) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLMNID.
func (in *PLMNID) DeepCopy() *PLMNID {
	if in == nil {
		return nil
	}
	out := new(PLMNID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLMNInfo) DeepCopyInto(out *PLMNInfo) {
	*out = *in
	out.PLMNID = in.PLMNID
	if in.NSSAI != nil {
		in, out := &in.NSSAI, &out.NSSAI
		*out = make([]NSSAI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLMNInfo.
func (in *PLMNInfo) DeepCopy() *PLMNInfo {
	if in == nil {
		return nil
	}
	out := new(PLMNInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLMNList) DeepCopyInto(out *PLMNList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PLMN, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLMNList.
func (in *PLMNList) DeepCopy() *PLMNList {
	if in == nil {
		return nil
	}
	out := new(PLMNList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PLMNList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLMNSpec) DeepCopyInto(out *PLMNSpec) {
	*out = *in
	if in.PLMNInfo != nil {
		in, out := &in.PLMNInfo, &out.PLMNInfo
		*out = make([]PLMNInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLMNSpec.
func (in *PLMNSpec) DeepCopy() *PLMNSpec {
	if in == nil {
		return nil
	}
	out := new(PLMNSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PLMNStatus) DeepCopyInto(out *PLMNStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PLMNStatus.
func (in *PLMNStatus) DeepCopy() *PLMNStatus {
	if in == nil {
		return nil
	}
	out := new(PLMNStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServedGUAMI) DeepCopyInto(out *ServedGUAMI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServedGUAMI.
func (in *ServedGUAMI) DeepCopy() *ServedGUAMI {
	if in == nil {
		return nil
	}
	out := new(ServedGUAMI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServedGUAMI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServedGUAMIList) DeepCopyInto(out *ServedGUAMIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServedGUAMI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServedGUAMIList.
func (in *ServedGUAMIList) DeepCopy() *ServedGUAMIList {
	if in == nil {
		return nil
	}
	out := new(ServedGUAMIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServedGUAMIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServedGUAMISpec) DeepCopyInto(out *ServedGUAMISpec) {
	*out = *in
	out.PLMNID = in.PLMNID
	out.AMFID = in.AMFID
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServedGUAMISpec.
func (in *ServedGUAMISpec) DeepCopy() *ServedGUAMISpec {
	if in == nil {
		return nil
	}
	out := new(ServedGUAMISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServedGUAMIStatus) DeepCopyInto(out *ServedGUAMIStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServedGUAMIStatus.
func (in *ServedGUAMIStatus) DeepCopy() *ServedGUAMIStatus {
	if in == nil {
		return nil
	}
	out := new(ServedGUAMIStatus)
	in.DeepCopyInto(out)
	return out
}