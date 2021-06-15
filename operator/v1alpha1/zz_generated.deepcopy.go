// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegatedAuthentication) DeepCopyInto(out *DelegatedAuthentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegatedAuthentication.
func (in *DelegatedAuthentication) DeepCopy() *DelegatedAuthentication {
	if in == nil {
		return nil
	}
	out := new(DelegatedAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegatedAuthorization) DeepCopyInto(out *DelegatedAuthorization) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegatedAuthorization.
func (in *DelegatedAuthorization) DeepCopy() *DelegatedAuthorization {
	if in == nil {
		return nil
	}
	out := new(DelegatedAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNS) DeepCopyInto(out *ExternalDNS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNS.
func (in *ExternalDNS) DeepCopy() *ExternalDNS {
	if in == nil {
		return nil
	}
	out := new(ExternalDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalDNS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSCRDSourceOptions) DeepCopyInto(out *ExternalDNSCRDSourceOptions) {
	*out = *in
	if in.LabelFilter != nil {
		in, out := &in.LabelFilter, &out.LabelFilter
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSCRDSourceOptions.
func (in *ExternalDNSCRDSourceOptions) DeepCopy() *ExternalDNSCRDSourceOptions {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSCRDSourceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSDomain) DeepCopyInto(out *ExternalDNSDomain) {
	*out = *in
	in.ExternalDNSDomainUnion.DeepCopyInto(&out.ExternalDNSDomainUnion)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSDomain.
func (in *ExternalDNSDomain) DeepCopy() *ExternalDNSDomain {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSDomainUnion) DeepCopyInto(out *ExternalDNSDomainUnion) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSDomainUnion.
func (in *ExternalDNSDomainUnion) DeepCopy() *ExternalDNSDomainUnion {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSDomainUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSList) DeepCopyInto(out *ExternalDNSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalDNS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSList.
func (in *ExternalDNSList) DeepCopy() *ExternalDNSList {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalDNSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSOperatorCondition) DeepCopyInto(out *ExternalDNSOperatorCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSOperatorCondition.
func (in *ExternalDNSOperatorCondition) DeepCopy() *ExternalDNSOperatorCondition {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSOperatorCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSProvider) DeepCopyInto(out *ExternalDNSProvider) {
	*out = *in
	out.ExternalDNSProviderUnion = in.ExternalDNSProviderUnion
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSProvider.
func (in *ExternalDNSProvider) DeepCopy() *ExternalDNSProvider {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSProviderUnion) DeepCopyInto(out *ExternalDNSProviderUnion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSProviderUnion.
func (in *ExternalDNSProviderUnion) DeepCopy() *ExternalDNSProviderUnion {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSProviderUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSServiceSourceOptions) DeepCopyInto(out *ExternalDNSServiceSourceOptions) {
	*out = *in
	if in.PublishPolicy != nil {
		in, out := &in.PublishPolicy, &out.PublishPolicy
		*out = new(ServicePublishPolicy)
		**out = **in
	}
	if in.ServiceType != nil {
		in, out := &in.ServiceType, &out.ServiceType
		*out = make([]v1.ServiceType, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSServiceSourceOptions.
func (in *ExternalDNSServiceSourceOptions) DeepCopy() *ExternalDNSServiceSourceOptions {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSServiceSourceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSSource) DeepCopyInto(out *ExternalDNSSource) {
	*out = *in
	in.ExternalDNSSourceUnion.DeepCopyInto(&out.ExternalDNSSourceUnion)
	if in.HostnameAnnotationPolicy != nil {
		in, out := &in.HostnameAnnotationPolicy, &out.HostnameAnnotationPolicy
		*out = new(HostnameAnnotationPolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSSource.
func (in *ExternalDNSSource) DeepCopy() *ExternalDNSSource {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSSourceUnion) DeepCopyInto(out *ExternalDNSSourceUnion) {
	*out = *in
	if in.AnnotationFilter != nil {
		in, out := &in.AnnotationFilter, &out.AnnotationFilter
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ExternalDNSServiceSourceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.CRD != nil {
		in, out := &in.CRD, &out.CRD
		*out = new(ExternalDNSCRDSourceOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSSourceUnion.
func (in *ExternalDNSSourceUnion) DeepCopy() *ExternalDNSSourceUnion {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSSourceUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSSpec) DeepCopyInto(out *ExternalDNSSpec) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]ExternalDNSDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Provider.DeepCopyInto(&out.Provider)
	in.Source.DeepCopyInto(&out.Source)
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSSpec.
func (in *ExternalDNSSpec) DeepCopy() *ExternalDNSSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSStatus) DeepCopyInto(out *ExternalDNSStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ExternalDNSOperatorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Provider.DeepCopyInto(&out.Provider)
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSStatus.
func (in *ExternalDNSStatus) DeepCopy() *ExternalDNSStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenerationHistory) DeepCopyInto(out *GenerationHistory) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenerationHistory.
func (in *GenerationHistory) DeepCopy() *GenerationHistory {
	if in == nil {
		return nil
	}
	out := new(GenerationHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericOperatorConfig) DeepCopyInto(out *GenericOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ServingInfo.DeepCopyInto(&out.ServingInfo)
	out.LeaderElection = in.LeaderElection
	out.Authentication = in.Authentication
	out.Authorization = in.Authorization
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericOperatorConfig.
func (in *GenericOperatorConfig) DeepCopy() *GenericOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(GenericOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GenericOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageContentSourcePolicy) DeepCopyInto(out *ImageContentSourcePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageContentSourcePolicy.
func (in *ImageContentSourcePolicy) DeepCopy() *ImageContentSourcePolicy {
	if in == nil {
		return nil
	}
	out := new(ImageContentSourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageContentSourcePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageContentSourcePolicyList) DeepCopyInto(out *ImageContentSourcePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageContentSourcePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageContentSourcePolicyList.
func (in *ImageContentSourcePolicyList) DeepCopy() *ImageContentSourcePolicyList {
	if in == nil {
		return nil
	}
	out := new(ImageContentSourcePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageContentSourcePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageContentSourcePolicySpec) DeepCopyInto(out *ImageContentSourcePolicySpec) {
	*out = *in
	if in.RepositoryDigestMirrors != nil {
		in, out := &in.RepositoryDigestMirrors, &out.RepositoryDigestMirrors
		*out = make([]RepositoryDigestMirrors, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageContentSourcePolicySpec.
func (in *ImageContentSourcePolicySpec) DeepCopy() *ImageContentSourcePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ImageContentSourcePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfig) DeepCopyInto(out *LoggingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfig.
func (in *LoggingConfig) DeepCopy() *LoggingConfig {
	if in == nil {
		return nil
	}
	out := new(LoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	if in.LastFailedDeploymentErrors != nil {
		in, out := &in.LastFailedDeploymentErrors, &out.LastFailedDeploymentErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorCondition) DeepCopyInto(out *OperatorCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorCondition.
func (in *OperatorCondition) DeepCopy() *OperatorCondition {
	if in == nil {
		return nil
	}
	out := new(OperatorCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorSpec) DeepCopyInto(out *OperatorSpec) {
	*out = *in
	out.Logging = in.Logging
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorSpec.
func (in *OperatorSpec) DeepCopy() *OperatorSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorStatus) DeepCopyInto(out *OperatorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OperatorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CurrentAvailability != nil {
		in, out := &in.CurrentAvailability, &out.CurrentAvailability
		*out = new(VersionAvailability)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetAvailability != nil {
		in, out := &in.TargetAvailability, &out.TargetAvailability
		*out = new(VersionAvailability)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorStatus.
func (in *OperatorStatus) DeepCopy() *OperatorStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryDigestMirrors) DeepCopyInto(out *RepositoryDigestMirrors) {
	*out = *in
	if in.Mirrors != nil {
		in, out := &in.Mirrors, &out.Mirrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryDigestMirrors.
func (in *RepositoryDigestMirrors) DeepCopy() *RepositoryDigestMirrors {
	if in == nil {
		return nil
	}
	out := new(RepositoryDigestMirrors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticPodOperatorStatus) DeepCopyInto(out *StaticPodOperatorStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	if in.NodeStatuses != nil {
		in, out := &in.NodeStatuses, &out.NodeStatuses
		*out = make([]NodeStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticPodOperatorStatus.
func (in *StaticPodOperatorStatus) DeepCopy() *StaticPodOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(StaticPodOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionAvailability) DeepCopyInto(out *VersionAvailability) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Generations != nil {
		in, out := &in.Generations, &out.Generations
		*out = make([]GenerationHistory, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionAvailability.
func (in *VersionAvailability) DeepCopy() *VersionAvailability {
	if in == nil {
		return nil
	}
	out := new(VersionAvailability)
	in.DeepCopyInto(out)
	return out
}
