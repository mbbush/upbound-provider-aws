//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternateContact) DeepCopyInto(out *AlternateContact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternateContact.
func (in *AlternateContact) DeepCopy() *AlternateContact {
	if in == nil {
		return nil
	}
	out := new(AlternateContact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlternateContact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternateContactList) DeepCopyInto(out *AlternateContactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlternateContact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternateContactList.
func (in *AlternateContactList) DeepCopy() *AlternateContactList {
	if in == nil {
		return nil
	}
	out := new(AlternateContactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlternateContactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternateContactObservation) DeepCopyInto(out *AlternateContactObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternateContactObservation.
func (in *AlternateContactObservation) DeepCopy() *AlternateContactObservation {
	if in == nil {
		return nil
	}
	out := new(AlternateContactObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternateContactParameters) DeepCopyInto(out *AlternateContactParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AlternateContactType != nil {
		in, out := &in.AlternateContactType, &out.AlternateContactType
		*out = new(string)
		**out = **in
	}
	if in.EmailAddress != nil {
		in, out := &in.EmailAddress, &out.EmailAddress
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternateContactParameters.
func (in *AlternateContactParameters) DeepCopy() *AlternateContactParameters {
	if in == nil {
		return nil
	}
	out := new(AlternateContactParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternateContactSpec) DeepCopyInto(out *AlternateContactSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternateContactSpec.
func (in *AlternateContactSpec) DeepCopy() *AlternateContactSpec {
	if in == nil {
		return nil
	}
	out := new(AlternateContactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternateContactStatus) DeepCopyInto(out *AlternateContactStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternateContactStatus.
func (in *AlternateContactStatus) DeepCopy() *AlternateContactStatus {
	if in == nil {
		return nil
	}
	out := new(AlternateContactStatus)
	in.DeepCopyInto(out)
	return out
}
