//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package policymanager

import (
	"fybrik.io/fybrik/pkg/model/datacatalog"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GetPolicyDecisionsRequest) DeepCopyInto(out *GetPolicyDecisionsRequest) {
	*out = *in
	in.Context.DeepCopyInto(&out.Context)
	out.Action = in.Action
	in.Resource.DeepCopyInto(&out.Resource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetPolicyDecisionsRequest.
func (in *GetPolicyDecisionsRequest) DeepCopy() *GetPolicyDecisionsRequest {
	if in == nil {
		return nil
	}
	out := new(GetPolicyDecisionsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GetPolicyDecisionsResponse) DeepCopyInto(out *GetPolicyDecisionsResponse) {
	*out = *in
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		*out = make([]ResultItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetPolicyDecisionsResponse.
func (in *GetPolicyDecisionsResponse) DeepCopy() *GetPolicyDecisionsResponse {
	if in == nil {
		return nil
	}
	out := new(GetPolicyDecisionsResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestAction) DeepCopyInto(out *RequestAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestAction.
func (in *RequestAction) DeepCopy() *RequestAction {
	if in == nil {
		return nil
	}
	out := new(RequestAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(datacatalog.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultItem) DeepCopyInto(out *ResultItem) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultItem.
func (in *ResultItem) DeepCopy() *ResultItem {
	if in == nil {
		return nil
	}
	out := new(ResultItem)
	in.DeepCopyInto(out)
	return out
}
