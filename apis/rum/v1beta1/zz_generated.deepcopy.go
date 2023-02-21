//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitor) DeepCopyInto(out *AppMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitor.
func (in *AppMonitor) DeepCopy() *AppMonitor {
	if in == nil {
		return nil
	}
	out := new(AppMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorConfigurationObservation) DeepCopyInto(out *AppMonitorConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorConfigurationObservation.
func (in *AppMonitorConfigurationObservation) DeepCopy() *AppMonitorConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AppMonitorConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorConfigurationParameters) DeepCopyInto(out *AppMonitorConfigurationParameters) {
	*out = *in
	if in.AllowCookies != nil {
		in, out := &in.AllowCookies, &out.AllowCookies
		*out = new(bool)
		**out = **in
	}
	if in.EnableXray != nil {
		in, out := &in.EnableXray, &out.EnableXray
		*out = new(bool)
		**out = **in
	}
	if in.ExcludedPages != nil {
		in, out := &in.ExcludedPages, &out.ExcludedPages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.FavoritePages != nil {
		in, out := &in.FavoritePages, &out.FavoritePages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.GuestRoleArn != nil {
		in, out := &in.GuestRoleArn, &out.GuestRoleArn
		*out = new(string)
		**out = **in
	}
	if in.IdentityPoolID != nil {
		in, out := &in.IdentityPoolID, &out.IdentityPoolID
		*out = new(string)
		**out = **in
	}
	if in.IncludedPages != nil {
		in, out := &in.IncludedPages, &out.IncludedPages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SessionSampleRate != nil {
		in, out := &in.SessionSampleRate, &out.SessionSampleRate
		*out = new(float64)
		**out = **in
	}
	if in.Telemetries != nil {
		in, out := &in.Telemetries, &out.Telemetries
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorConfigurationParameters.
func (in *AppMonitorConfigurationParameters) DeepCopy() *AppMonitorConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AppMonitorConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorList) DeepCopyInto(out *AppMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorList.
func (in *AppMonitorList) DeepCopy() *AppMonitorList {
	if in == nil {
		return nil
	}
	out := new(AppMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorObservation) DeepCopyInto(out *AppMonitorObservation) {
	*out = *in
	if in.AppMonitorID != nil {
		in, out := &in.AppMonitorID, &out.AppMonitorID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CwLogGroup != nil {
		in, out := &in.CwLogGroup, &out.CwLogGroup
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorObservation.
func (in *AppMonitorObservation) DeepCopy() *AppMonitorObservation {
	if in == nil {
		return nil
	}
	out := new(AppMonitorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorParameters) DeepCopyInto(out *AppMonitorParameters) {
	*out = *in
	if in.AppMonitorConfiguration != nil {
		in, out := &in.AppMonitorConfiguration, &out.AppMonitorConfiguration
		*out = make([]AppMonitorConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomEvents != nil {
		in, out := &in.CustomEvents, &out.CustomEvents
		*out = make([]CustomEventsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CwLogEnabled != nil {
		in, out := &in.CwLogEnabled, &out.CwLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorParameters.
func (in *AppMonitorParameters) DeepCopy() *AppMonitorParameters {
	if in == nil {
		return nil
	}
	out := new(AppMonitorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorSpec) DeepCopyInto(out *AppMonitorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorSpec.
func (in *AppMonitorSpec) DeepCopy() *AppMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(AppMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppMonitorStatus) DeepCopyInto(out *AppMonitorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppMonitorStatus.
func (in *AppMonitorStatus) DeepCopy() *AppMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(AppMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEventsObservation) DeepCopyInto(out *CustomEventsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEventsObservation.
func (in *CustomEventsObservation) DeepCopy() *CustomEventsObservation {
	if in == nil {
		return nil
	}
	out := new(CustomEventsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEventsParameters) DeepCopyInto(out *CustomEventsParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEventsParameters.
func (in *CustomEventsParameters) DeepCopy() *CustomEventsParameters {
	if in == nil {
		return nil
	}
	out := new(CustomEventsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsDestination) DeepCopyInto(out *MetricsDestination) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsDestination.
func (in *MetricsDestination) DeepCopy() *MetricsDestination {
	if in == nil {
		return nil
	}
	out := new(MetricsDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsDestination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsDestinationList) DeepCopyInto(out *MetricsDestinationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricsDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsDestinationList.
func (in *MetricsDestinationList) DeepCopy() *MetricsDestinationList {
	if in == nil {
		return nil
	}
	out := new(MetricsDestinationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsDestinationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsDestinationObservation) DeepCopyInto(out *MetricsDestinationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsDestinationObservation.
func (in *MetricsDestinationObservation) DeepCopy() *MetricsDestinationObservation {
	if in == nil {
		return nil
	}
	out := new(MetricsDestinationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsDestinationParameters) DeepCopyInto(out *MetricsDestinationParameters) {
	*out = *in
	if in.AppMonitorName != nil {
		in, out := &in.AppMonitorName, &out.AppMonitorName
		*out = new(string)
		**out = **in
	}
	if in.AppMonitorNameRef != nil {
		in, out := &in.AppMonitorNameRef, &out.AppMonitorNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AppMonitorNameSelector != nil {
		in, out := &in.AppMonitorNameSelector, &out.AppMonitorNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	if in.DestinationArn != nil {
		in, out := &in.DestinationArn, &out.DestinationArn
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleArn != nil {
		in, out := &in.IAMRoleArn, &out.IAMRoleArn
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleArnRef != nil {
		in, out := &in.IAMRoleArnRef, &out.IAMRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IAMRoleArnSelector != nil {
		in, out := &in.IAMRoleArnSelector, &out.IAMRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsDestinationParameters.
func (in *MetricsDestinationParameters) DeepCopy() *MetricsDestinationParameters {
	if in == nil {
		return nil
	}
	out := new(MetricsDestinationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsDestinationSpec) DeepCopyInto(out *MetricsDestinationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsDestinationSpec.
func (in *MetricsDestinationSpec) DeepCopy() *MetricsDestinationSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsDestinationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsDestinationStatus) DeepCopyInto(out *MetricsDestinationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsDestinationStatus.
func (in *MetricsDestinationStatus) DeepCopy() *MetricsDestinationStatus {
	if in == nil {
		return nil
	}
	out := new(MetricsDestinationStatus)
	in.DeepCopyInto(out)
	return out
}
