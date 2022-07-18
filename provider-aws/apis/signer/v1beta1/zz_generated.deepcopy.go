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
func (in *RevocationRecordObservation) DeepCopyInto(out *RevocationRecordObservation) {
	*out = *in
	if in.RevocationEffectiveFrom != nil {
		in, out := &in.RevocationEffectiveFrom, &out.RevocationEffectiveFrom
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = new(string)
		**out = **in
	}
	if in.RevokedBy != nil {
		in, out := &in.RevokedBy, &out.RevokedBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordObservation.
func (in *RevocationRecordObservation) DeepCopy() *RevocationRecordObservation {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationRecordParameters) DeepCopyInto(out *RevocationRecordParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordParameters.
func (in *RevocationRecordParameters) DeepCopy() *RevocationRecordParameters {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodObservation) DeepCopyInto(out *SignatureValidityPeriodObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodObservation.
func (in *SignatureValidityPeriodObservation) DeepCopy() *SignatureValidityPeriodObservation {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodParameters) DeepCopyInto(out *SignatureValidityPeriodParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodParameters.
func (in *SignatureValidityPeriodParameters) DeepCopy() *SignatureValidityPeriodParameters {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfile) DeepCopyInto(out *SigningProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfile.
func (in *SigningProfile) DeepCopy() *SigningProfile {
	if in == nil {
		return nil
	}
	out := new(SigningProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileList) DeepCopyInto(out *SigningProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SigningProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileList.
func (in *SigningProfileList) DeepCopy() *SigningProfileList {
	if in == nil {
		return nil
	}
	out := new(SigningProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileObservation) DeepCopyInto(out *SigningProfileObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PlatformDisplayName != nil {
		in, out := &in.PlatformDisplayName, &out.PlatformDisplayName
		*out = new(string)
		**out = **in
	}
	if in.RevocationRecord != nil {
		in, out := &in.RevocationRecord, &out.RevocationRecord
		*out = make([]RevocationRecordObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VersionArn != nil {
		in, out := &in.VersionArn, &out.VersionArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileObservation.
func (in *SigningProfileObservation) DeepCopy() *SigningProfileObservation {
	if in == nil {
		return nil
	}
	out := new(SigningProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileParameters) DeepCopyInto(out *SigningProfileParameters) {
	*out = *in
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SignatureValidityPeriod != nil {
		in, out := &in.SignatureValidityPeriod, &out.SignatureValidityPeriod
		*out = make([]SignatureValidityPeriodParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileParameters.
func (in *SigningProfileParameters) DeepCopy() *SigningProfileParameters {
	if in == nil {
		return nil
	}
	out := new(SigningProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileSpec) DeepCopyInto(out *SigningProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileSpec.
func (in *SigningProfileSpec) DeepCopy() *SigningProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SigningProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileStatus) DeepCopyInto(out *SigningProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileStatus.
func (in *SigningProfileStatus) DeepCopy() *SigningProfileStatus {
	if in == nil {
		return nil
	}
	out := new(SigningProfileStatus)
	in.DeepCopyInto(out)
	return out
}
