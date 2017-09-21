// +build !ignore_autogenerated

/*
Copyright 2017 The Stash Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	reflect "reflect"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/client-go/pkg/api/v1"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_AzureSpec, InType: reflect.TypeOf(&AzureSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_B2Spec, InType: reflect.TypeOf(&B2Spec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Backend, InType: reflect.TypeOf(&Backend{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_FileGroup, InType: reflect.TypeOf(&FileGroup{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_GCSSpec, InType: reflect.TypeOf(&GCSSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_LocalSpec, InType: reflect.TypeOf(&LocalSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_RestServerSpec, InType: reflect.TypeOf(&RestServerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Restic, InType: reflect.TypeOf(&Restic{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ResticList, InType: reflect.TypeOf(&ResticList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ResticSpec, InType: reflect.TypeOf(&ResticSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ResticStatus, InType: reflect.TypeOf(&ResticStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_RetentionPolicy, InType: reflect.TypeOf(&RetentionPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_S3Spec, InType: reflect.TypeOf(&S3Spec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_SwiftSpec, InType: reflect.TypeOf(&SwiftSpec{})},
	)
}

// DeepCopy_v1alpha1_AzureSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_AzureSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AzureSpec)
		out := out.(*AzureSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_B2Spec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_B2Spec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*B2Spec)
		out := out.(*B2Spec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_Backend is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Backend(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Backend)
		out := out.(*Backend)
		*out = *in
		if in.Local != nil {
			in, out := &in.Local, &out.Local
			*out = new(LocalSpec)
			if err := DeepCopy_v1alpha1_LocalSpec(*in, *out, c); err != nil {
				return err
			}
		}
		if in.S3 != nil {
			in, out := &in.S3, &out.S3
			*out = new(S3Spec)
			**out = **in
		}
		if in.GCS != nil {
			in, out := &in.GCS, &out.GCS
			*out = new(GCSSpec)
			**out = **in
		}
		if in.Azure != nil {
			in, out := &in.Azure, &out.Azure
			*out = new(AzureSpec)
			**out = **in
		}
		if in.Swift != nil {
			in, out := &in.Swift, &out.Swift
			*out = new(SwiftSpec)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_FileGroup is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_FileGroup(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FileGroup)
		out := out.(*FileGroup)
		*out = *in
		if in.Tags != nil {
			in, out := &in.Tags, &out.Tags
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if err := DeepCopy_v1alpha1_RetentionPolicy(&in.RetentionPolicy, &out.RetentionPolicy, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_GCSSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_GCSSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GCSSpec)
		out := out.(*GCSSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_LocalSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_LocalSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSpec)
		out := out.(*LocalSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.VolumeSource); err != nil {
			return err
		} else {
			out.VolumeSource = *newVal.(*v1.VolumeSource)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_RestServerSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_RestServerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RestServerSpec)
		out := out.(*RestServerSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_Restic is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Restic(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Restic)
		out := out.(*Restic)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_ResticSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_ResticStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ResticList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ResticList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResticList)
		out := out.(*ResticList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Restic, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Restic(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ResticSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ResticSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResticSpec)
		out := out.(*ResticSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Selector); err != nil {
			return err
		} else {
			out.Selector = *newVal.(*meta_v1.LabelSelector)
		}
		if in.FileGroups != nil {
			in, out := &in.FileGroups, &out.FileGroups
			*out = make([]FileGroup, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_FileGroup(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if err := DeepCopy_v1alpha1_Backend(&in.Backend, &out.Backend, c); err != nil {
			return err
		}
		if in.VolumeMounts != nil {
			in, out := &in.VolumeMounts, &out.VolumeMounts
			*out = make([]v1.VolumeMount, len(*in))
			copy(*out, *in)
		}
		if newVal, err := c.DeepCopy(&in.Resources); err != nil {
			return err
		} else {
			out.Resources = *newVal.(*v1.ResourceRequirements)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ResticStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ResticStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResticStatus)
		out := out.(*ResticStatus)
		*out = *in
		if in.FirstBackupTime != nil {
			in, out := &in.FirstBackupTime, &out.FirstBackupTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.LastBackupTime != nil {
			in, out := &in.LastBackupTime, &out.LastBackupTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.LastSuccessfulBackupTime != nil {
			in, out := &in.LastSuccessfulBackupTime, &out.LastSuccessfulBackupTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

// DeepCopy_v1alpha1_RetentionPolicy is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_RetentionPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RetentionPolicy)
		out := out.(*RetentionPolicy)
		*out = *in
		if in.KeepTags != nil {
			in, out := &in.KeepTags, &out.KeepTags
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_S3Spec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_S3Spec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*S3Spec)
		out := out.(*S3Spec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_SwiftSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_SwiftSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SwiftSpec)
		out := out.(*SwiftSpec)
		*out = *in
		return nil
	}
}
