//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Alibaba Group Holding Limited.

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

package v1

import (
	"github.com/alibaba/polardbx-operator/api/v1/polardbx"
	"github.com/alibaba/polardbx-operator/api/v1/xstore"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXCluster) DeepCopyInto(out *PolarDBXCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXCluster.
func (in *PolarDBXCluster) DeepCopy() *PolarDBXCluster {
	if in == nil {
		return nil
	}
	out := new(PolarDBXCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolarDBXCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterKnobs) DeepCopyInto(out *PolarDBXClusterKnobs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterKnobs.
func (in *PolarDBXClusterKnobs) DeepCopy() *PolarDBXClusterKnobs {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterKnobs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolarDBXClusterKnobs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterKnobsList) DeepCopyInto(out *PolarDBXClusterKnobsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolarDBXClusterKnobs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterKnobsList.
func (in *PolarDBXClusterKnobsList) DeepCopy() *PolarDBXClusterKnobsList {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterKnobsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolarDBXClusterKnobsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterKnobsSpec) DeepCopyInto(out *PolarDBXClusterKnobsSpec) {
	*out = *in
	if in.Knobs != nil {
		in, out := &in.Knobs, &out.Knobs
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterKnobsSpec.
func (in *PolarDBXClusterKnobsSpec) DeepCopy() *PolarDBXClusterKnobsSpec {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterKnobsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterKnobsStatus) DeepCopyInto(out *PolarDBXClusterKnobsStatus) {
	*out = *in
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterKnobsStatus.
func (in *PolarDBXClusterKnobsStatus) DeepCopy() *PolarDBXClusterKnobsStatus {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterKnobsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterList) DeepCopyInto(out *PolarDBXClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolarDBXCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterList.
func (in *PolarDBXClusterList) DeepCopy() *PolarDBXClusterList {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolarDBXClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterSpec) DeepCopyInto(out *PolarDBXClusterSpec) {
	*out = *in
	out.ProtocolVersion = in.ProtocolVersion
	in.Topology.DeepCopyInto(&out.Topology)
	in.Config.DeepCopyInto(&out.Config)
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]polardbx.PrivilegeItem, len(*in))
		copy(*out, *in)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(polardbx.Security)
		(*in).DeepCopyInto(*out)
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		*out = new(polardbx.RestoreSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterSpec.
func (in *PolarDBXClusterSpec) DeepCopy() *PolarDBXClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXClusterStatus) DeepCopyInto(out *PolarDBXClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]polardbx.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StatusForPrint.DeepCopyInto(&out.StatusForPrint)
	in.ReplicaStatus.DeepCopyInto(&out.ReplicaStatus)
	if in.SpecSnapshot != nil {
		in, out := &in.SpecSnapshot, &out.SpecSnapshot
		*out = new(polardbx.SpecSnapshot)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXClusterStatus.
func (in *PolarDBXClusterStatus) DeepCopy() *PolarDBXClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PolarDBXClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXMonitor) DeepCopyInto(out *PolarDBXMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXMonitor.
func (in *PolarDBXMonitor) DeepCopy() *PolarDBXMonitor {
	if in == nil {
		return nil
	}
	out := new(PolarDBXMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolarDBXMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXMonitorList) DeepCopyInto(out *PolarDBXMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolarDBXMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXMonitorList.
func (in *PolarDBXMonitorList) DeepCopy() *PolarDBXMonitorList {
	if in == nil {
		return nil
	}
	out := new(PolarDBXMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolarDBXMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXMonitorSpec) DeepCopyInto(out *PolarDBXMonitorSpec) {
	*out = *in
	out.MonitorInterval = in.MonitorInterval
	out.ScrapeTimeout = in.ScrapeTimeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXMonitorSpec.
func (in *PolarDBXMonitorSpec) DeepCopy() *PolarDBXMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(PolarDBXMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolarDBXMonitorStatus) DeepCopyInto(out *PolarDBXMonitorStatus) {
	*out = *in
	if in.MonitorSpecSnapshot != nil {
		in, out := &in.MonitorSpecSnapshot, &out.MonitorSpecSnapshot
		*out = new(PolarDBXMonitorSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolarDBXMonitorStatus.
func (in *PolarDBXMonitorStatus) DeepCopy() *PolarDBXMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(PolarDBXMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XStore) DeepCopyInto(out *XStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XStore.
func (in *XStore) DeepCopy() *XStore {
	if in == nil {
		return nil
	}
	out := new(XStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XStoreList) DeepCopyInto(out *XStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XStoreList.
func (in *XStoreList) DeepCopy() *XStoreList {
	if in == nil {
		return nil
	}
	out := new(XStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XStoreSpec) DeepCopyInto(out *XStoreSpec) {
	*out = *in
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]xstore.Privilege, len(*in))
		copy(*out, *in)
	}
	in.Topology.DeepCopyInto(&out.Topology)
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XStoreSpec.
func (in *XStoreSpec) DeepCopy() *XStoreSpec {
	if in == nil {
		return nil
	}
	out := new(XStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XStoreStatus) DeepCopyInto(out *XStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]xstore.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BoundVolumes != nil {
		in, out := &in.BoundVolumes, &out.BoundVolumes
		*out = make(map[string]*xstore.HostPathVolume, len(*in))
		for key, val := range *in {
			var outVal *xstore.HostPathVolume
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(xstore.HostPathVolume)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LastVolumeSizeUpdateTime != nil {
		in, out := &in.LastVolumeSizeUpdateTime, &out.LastVolumeSizeUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.ObservedTopology != nil {
		in, out := &in.ObservedTopology, &out.ObservedTopology
		*out = new(xstore.Topology)
		(*in).DeepCopyInto(*out)
	}
	if in.ObservedConfig != nil {
		in, out := &in.ObservedConfig, &out.ObservedConfig
		*out = new(xstore.Config)
		(*in).DeepCopyInto(*out)
	}
	if in.LastLogPurgeTime != nil {
		in, out := &in.LastLogPurgeTime, &out.LastLogPurgeTime
		*out = (*in).DeepCopy()
	}
	if in.PodPorts != nil {
		in, out := &in.PodPorts, &out.PodPorts
		*out = make(map[string]xstore.PodPorts, len(*in))
		for key, val := range *in {
			var outVal map[string]int32
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(xstore.PodPorts, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XStoreStatus.
func (in *XStoreStatus) DeepCopy() *XStoreStatus {
	if in == nil {
		return nil
	}
	out := new(XStoreStatus)
	in.DeepCopyInto(out)
	return out
}
