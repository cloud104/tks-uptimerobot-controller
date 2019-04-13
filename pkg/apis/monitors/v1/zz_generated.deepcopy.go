// +build !ignore_autogenerated

/*

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
// Code generated by main. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeRobot) DeepCopyInto(out *UptimeRobot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeRobot.
func (in *UptimeRobot) DeepCopy() *UptimeRobot {
	if in == nil {
		return nil
	}
	out := new(UptimeRobot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UptimeRobot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeRobotHosts) DeepCopyInto(out *UptimeRobotHosts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeRobotHosts.
func (in *UptimeRobotHosts) DeepCopy() *UptimeRobotHosts {
	if in == nil {
		return nil
	}
	out := new(UptimeRobotHosts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeRobotList) DeepCopyInto(out *UptimeRobotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UptimeRobot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeRobotList.
func (in *UptimeRobotList) DeepCopy() *UptimeRobotList {
	if in == nil {
		return nil
	}
	out := new(UptimeRobotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UptimeRobotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeRobotSpec) DeepCopyInto(out *UptimeRobotSpec) {
	*out = *in
	out.StatusPage = in.StatusPage
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]UptimeRobotHosts, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeRobotSpec.
func (in *UptimeRobotSpec) DeepCopy() *UptimeRobotSpec {
	if in == nil {
		return nil
	}
	out := new(UptimeRobotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeRobotStatus) DeepCopyInto(out *UptimeRobotStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeRobotStatus.
func (in *UptimeRobotStatus) DeepCopy() *UptimeRobotStatus {
	if in == nil {
		return nil
	}
	out := new(UptimeRobotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UptimeStatusPage) DeepCopyInto(out *UptimeStatusPage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UptimeStatusPage.
func (in *UptimeStatusPage) DeepCopy() *UptimeStatusPage {
	if in == nil {
		return nil
	}
	out := new(UptimeStatusPage)
	in.DeepCopyInto(out)
	return out
}
