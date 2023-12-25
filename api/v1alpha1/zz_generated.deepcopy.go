//go:build !ignore_autogenerated

/*
Copyright 2021 NVIDIA

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedState) DeepCopyInto(out *AppliedState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedState.
func (in *AppliedState) DeepCopy() *AppliedState {
	if in == nil {
		return nil
	}
	out := new(AppliedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapNameReference) DeepCopyInto(out *ConfigMapNameReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapNameReference.
func (in *ConfigMapNameReference) DeepCopy() *ConfigMapNameReference {
	if in == nil {
		return nil
	}
	out := new(ConfigMapNameReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePluginSpec) DeepCopyInto(out *DevicePluginSpec) {
	*out = *in
	in.ImageSpecWithConfig.DeepCopyInto(&out.ImageSpecWithConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePluginSpec.
func (in *DevicePluginSpec) DeepCopy() *DevicePluginSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrainSpec) DeepCopyInto(out *DrainSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrainSpec.
func (in *DrainSpec) DeepCopy() *DrainSpec {
	if in == nil {
		return nil
	}
	out := new(DrainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverUpgradePolicySpec) DeepCopyInto(out *DriverUpgradePolicySpec) {
	*out = *in
	if in.WaitForCompletion != nil {
		in, out := &in.WaitForCompletion, &out.WaitForCompletion
		*out = new(WaitForCompletionSpec)
		**out = **in
	}
	if in.DrainSpec != nil {
		in, out := &in.DrainSpec, &out.DrainSpec
		*out = new(DrainSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverUpgradePolicySpec.
func (in *DriverUpgradePolicySpec) DeepCopy() *DriverUpgradePolicySpec {
	if in == nil {
		return nil
	}
	out := new(DriverUpgradePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostDeviceNetwork) DeepCopyInto(out *HostDeviceNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostDeviceNetwork.
func (in *HostDeviceNetwork) DeepCopy() *HostDeviceNetwork {
	if in == nil {
		return nil
	}
	out := new(HostDeviceNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostDeviceNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostDeviceNetworkList) DeepCopyInto(out *HostDeviceNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostDeviceNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostDeviceNetworkList.
func (in *HostDeviceNetworkList) DeepCopy() *HostDeviceNetworkList {
	if in == nil {
		return nil
	}
	out := new(HostDeviceNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostDeviceNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostDeviceNetworkSpec) DeepCopyInto(out *HostDeviceNetworkSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostDeviceNetworkSpec.
func (in *HostDeviceNetworkSpec) DeepCopy() *HostDeviceNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(HostDeviceNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostDeviceNetworkStatus) DeepCopyInto(out *HostDeviceNetworkStatus) {
	*out = *in
	if in.AppliedStates != nil {
		in, out := &in.AppliedStates, &out.AppliedStates
		*out = make([]AppliedState, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostDeviceNetworkStatus.
func (in *HostDeviceNetworkStatus) DeepCopy() *HostDeviceNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(HostDeviceNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBKubernetesSpec) DeepCopyInto(out *IBKubernetesSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBKubernetesSpec.
func (in *IBKubernetesSpec) DeepCopy() *IBKubernetesSpec {
	if in == nil {
		return nil
	}
	out := new(IBKubernetesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPoIBNetwork) DeepCopyInto(out *IPoIBNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPoIBNetwork.
func (in *IPoIBNetwork) DeepCopy() *IPoIBNetwork {
	if in == nil {
		return nil
	}
	out := new(IPoIBNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPoIBNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPoIBNetworkList) DeepCopyInto(out *IPoIBNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPoIBNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPoIBNetworkList.
func (in *IPoIBNetworkList) DeepCopy() *IPoIBNetworkList {
	if in == nil {
		return nil
	}
	out := new(IPoIBNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPoIBNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPoIBNetworkSpec) DeepCopyInto(out *IPoIBNetworkSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPoIBNetworkSpec.
func (in *IPoIBNetworkSpec) DeepCopy() *IPoIBNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(IPoIBNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPoIBNetworkStatus) DeepCopyInto(out *IPoIBNetworkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPoIBNetworkStatus.
func (in *IPoIBNetworkStatus) DeepCopy() *IPoIBNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(IPoIBNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContainerResources != nil {
		in, out := &in.ContainerResources, &out.ContainerResources
		*out = make([]ResourceRequirements, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpecWithConfig) DeepCopyInto(out *ImageSpecWithConfig) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpecWithConfig.
func (in *ImageSpecWithConfig) DeepCopy() *ImageSpecWithConfig {
	if in == nil {
		return nil
	}
	out := new(ImageSpecWithConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacvlanNetwork) DeepCopyInto(out *MacvlanNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacvlanNetwork.
func (in *MacvlanNetwork) DeepCopy() *MacvlanNetwork {
	if in == nil {
		return nil
	}
	out := new(MacvlanNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MacvlanNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacvlanNetworkList) DeepCopyInto(out *MacvlanNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MacvlanNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacvlanNetworkList.
func (in *MacvlanNetworkList) DeepCopy() *MacvlanNetworkList {
	if in == nil {
		return nil
	}
	out := new(MacvlanNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MacvlanNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacvlanNetworkSpec) DeepCopyInto(out *MacvlanNetworkSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacvlanNetworkSpec.
func (in *MacvlanNetworkSpec) DeepCopy() *MacvlanNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(MacvlanNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MacvlanNetworkStatus) DeepCopyInto(out *MacvlanNetworkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MacvlanNetworkStatus.
func (in *MacvlanNetworkStatus) DeepCopy() *MacvlanNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(MacvlanNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultusSpec) DeepCopyInto(out *MultusSpec) {
	*out = *in
	in.ImageSpecWithConfig.DeepCopyInto(&out.ImageSpecWithConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultusSpec.
func (in *MultusSpec) DeepCopy() *MultusSpec {
	if in == nil {
		return nil
	}
	out := new(MultusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NICFeatureDiscoverySpec) DeepCopyInto(out *NICFeatureDiscoverySpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NICFeatureDiscoverySpec.
func (in *NICFeatureDiscoverySpec) DeepCopy() *NICFeatureDiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(NICFeatureDiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVIPAMSpec) DeepCopyInto(out *NVIPAMSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVIPAMSpec.
func (in *NVIPAMSpec) DeepCopy() *NVIPAMSpec {
	if in == nil {
		return nil
	}
	out := new(NVIPAMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicClusterPolicy) DeepCopyInto(out *NicClusterPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicClusterPolicy.
func (in *NicClusterPolicy) DeepCopy() *NicClusterPolicy {
	if in == nil {
		return nil
	}
	out := new(NicClusterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicClusterPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicClusterPolicyList) DeepCopyInto(out *NicClusterPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NicClusterPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicClusterPolicyList.
func (in *NicClusterPolicyList) DeepCopy() *NicClusterPolicyList {
	if in == nil {
		return nil
	}
	out := new(NicClusterPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicClusterPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicClusterPolicySpec) DeepCopyInto(out *NicClusterPolicySpec) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OFEDDriver != nil {
		in, out := &in.OFEDDriver, &out.OFEDDriver
		*out = new(OFEDDriverSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RdmaSharedDevicePlugin != nil {
		in, out := &in.RdmaSharedDevicePlugin, &out.RdmaSharedDevicePlugin
		*out = new(DevicePluginSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SriovDevicePlugin != nil {
		in, out := &in.SriovDevicePlugin, &out.SriovDevicePlugin
		*out = new(DevicePluginSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IBKubernetes != nil {
		in, out := &in.IBKubernetes, &out.IBKubernetes
		*out = new(IBKubernetesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecondaryNetwork != nil {
		in, out := &in.SecondaryNetwork, &out.SecondaryNetwork
		*out = new(SecondaryNetworkSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NvIpam != nil {
		in, out := &in.NvIpam, &out.NvIpam
		*out = new(NVIPAMSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NicFeatureDiscovery != nil {
		in, out := &in.NicFeatureDiscovery, &out.NicFeatureDiscovery
		*out = new(NICFeatureDiscoverySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicClusterPolicySpec.
func (in *NicClusterPolicySpec) DeepCopy() *NicClusterPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NicClusterPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicClusterPolicyStatus) DeepCopyInto(out *NicClusterPolicyStatus) {
	*out = *in
	if in.AppliedStates != nil {
		in, out := &in.AppliedStates, &out.AppliedStates
		*out = make([]AppliedState, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicClusterPolicyStatus.
func (in *NicClusterPolicyStatus) DeepCopy() *NicClusterPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(NicClusterPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OFEDDriverSpec) DeepCopyInto(out *OFEDDriverSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.StartupProbe != nil {
		in, out := &in.StartupProbe, &out.StartupProbe
		*out = new(PodProbeSpec)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(PodProbeSpec)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(PodProbeSpec)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OfedUpgradePolicy != nil {
		in, out := &in.OfedUpgradePolicy, &out.OfedUpgradePolicy
		*out = new(DriverUpgradePolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CertConfig != nil {
		in, out := &in.CertConfig, &out.CertConfig
		*out = new(ConfigMapNameReference)
		**out = **in
	}
	if in.RepoConfig != nil {
		in, out := &in.RepoConfig, &out.RepoConfig
		*out = new(ConfigMapNameReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OFEDDriverSpec.
func (in *OFEDDriverSpec) DeepCopy() *OFEDDriverSpec {
	if in == nil {
		return nil
	}
	out := new(OFEDDriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodProbeSpec) DeepCopyInto(out *PodProbeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodProbeSpec.
func (in *PodProbeSpec) DeepCopy() *PodProbeSpec {
	if in == nil {
		return nil
	}
	out := new(PodProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryNetworkSpec) DeepCopyInto(out *SecondaryNetworkSpec) {
	*out = *in
	if in.Multus != nil {
		in, out := &in.Multus, &out.Multus
		*out = new(MultusSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CniPlugins != nil {
		in, out := &in.CniPlugins, &out.CniPlugins
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IPoIB != nil {
		in, out := &in.IPoIB, &out.IPoIB
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IpamPlugin != nil {
		in, out := &in.IpamPlugin, &out.IpamPlugin
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryNetworkSpec.
func (in *SecondaryNetworkSpec) DeepCopy() *SecondaryNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(SecondaryNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitForCompletionSpec) DeepCopyInto(out *WaitForCompletionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitForCompletionSpec.
func (in *WaitForCompletionSpec) DeepCopy() *WaitForCompletionSpec {
	if in == nil {
		return nil
	}
	out := new(WaitForCompletionSpec)
	in.DeepCopyInto(out)
	return out
}
