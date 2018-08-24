// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*API)(nil), (*kubeadm.API)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_API_To_kubeadm_API(a.(*API), b.(*kubeadm.API), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.API)(nil), (*API)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_API_To_v1alpha2_API(a.(*kubeadm.API), b.(*API), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuditPolicyConfiguration)(nil), (*kubeadm.AuditPolicyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(a.(*AuditPolicyConfiguration), b.(*kubeadm.AuditPolicyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.AuditPolicyConfiguration)(nil), (*AuditPolicyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha2_AuditPolicyConfiguration(a.(*kubeadm.AuditPolicyConfiguration), b.(*AuditPolicyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BootstrapToken)(nil), (*kubeadm.BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_BootstrapToken_To_kubeadm_BootstrapToken(a.(*BootstrapToken), b.(*kubeadm.BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.BootstrapToken)(nil), (*BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_BootstrapToken_To_v1alpha2_BootstrapToken(a.(*kubeadm.BootstrapToken), b.(*BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BootstrapTokenString)(nil), (*kubeadm.BootstrapTokenString)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_BootstrapTokenString_To_kubeadm_BootstrapTokenString(a.(*BootstrapTokenString), b.(*kubeadm.BootstrapTokenString), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.BootstrapTokenString)(nil), (*BootstrapTokenString)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_BootstrapTokenString_To_v1alpha2_BootstrapTokenString(a.(*kubeadm.BootstrapTokenString), b.(*BootstrapTokenString), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Etcd)(nil), (*kubeadm.Etcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Etcd_To_kubeadm_Etcd(a.(*Etcd), b.(*kubeadm.Etcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.Etcd)(nil), (*Etcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_Etcd_To_v1alpha2_Etcd(a.(*kubeadm.Etcd), b.(*Etcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExternalEtcd)(nil), (*kubeadm.ExternalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ExternalEtcd_To_kubeadm_ExternalEtcd(a.(*ExternalEtcd), b.(*kubeadm.ExternalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.ExternalEtcd)(nil), (*ExternalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_ExternalEtcd_To_v1alpha2_ExternalEtcd(a.(*kubeadm.ExternalEtcd), b.(*ExternalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HostPathMount)(nil), (*kubeadm.HostPathMount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_HostPathMount_To_kubeadm_HostPathMount(a.(*HostPathMount), b.(*kubeadm.HostPathMount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.HostPathMount)(nil), (*HostPathMount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_HostPathMount_To_v1alpha2_HostPathMount(a.(*kubeadm.HostPathMount), b.(*HostPathMount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InitConfiguration)(nil), (*kubeadm.InitConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_InitConfiguration_To_kubeadm_InitConfiguration(a.(*InitConfiguration), b.(*kubeadm.InitConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.InitConfiguration)(nil), (*InitConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_InitConfiguration_To_v1alpha2_InitConfiguration(a.(*kubeadm.InitConfiguration), b.(*InitConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JoinConfiguration)(nil), (*kubeadm.JoinConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_JoinConfiguration_To_kubeadm_JoinConfiguration(a.(*JoinConfiguration), b.(*kubeadm.JoinConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.JoinConfiguration)(nil), (*JoinConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_JoinConfiguration_To_v1alpha2_JoinConfiguration(a.(*kubeadm.JoinConfiguration), b.(*JoinConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalEtcd)(nil), (*kubeadm.LocalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LocalEtcd_To_kubeadm_LocalEtcd(a.(*LocalEtcd), b.(*kubeadm.LocalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.LocalEtcd)(nil), (*LocalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_LocalEtcd_To_v1alpha2_LocalEtcd(a.(*kubeadm.LocalEtcd), b.(*LocalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Networking)(nil), (*kubeadm.Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Networking_To_kubeadm_Networking(a.(*Networking), b.(*kubeadm.Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.Networking)(nil), (*Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_Networking_To_v1alpha2_Networking(a.(*kubeadm.Networking), b.(*Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeRegistrationOptions)(nil), (*kubeadm.NodeRegistrationOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(a.(*NodeRegistrationOptions), b.(*kubeadm.NodeRegistrationOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.NodeRegistrationOptions)(nil), (*NodeRegistrationOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions(a.(*kubeadm.NodeRegistrationOptions), b.(*NodeRegistrationOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*kubeadm.API)(nil), (*API)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_API_To_v1alpha2_API(a.(*kubeadm.API), b.(*API), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*kubeadm.InitConfiguration)(nil), (*InitConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_InitConfiguration_To_v1alpha2_InitConfiguration(a.(*kubeadm.InitConfiguration), b.(*InitConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*API)(nil), (*kubeadm.API)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_API_To_kubeadm_API(a.(*API), b.(*kubeadm.API), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*InitConfiguration)(nil), (*kubeadm.InitConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_InitConfiguration_To_kubeadm_InitConfiguration(a.(*InitConfiguration), b.(*kubeadm.InitConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_API_To_kubeadm_API(in *API, out *kubeadm.API, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	// WARNING: in.ControlPlaneEndpoint requires manual conversion: does not exist in peer-type
	out.BindPort = in.BindPort
	return nil
}

func autoConvert_kubeadm_API_To_v1alpha2_API(in *kubeadm.API, out *API, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.BindPort = in.BindPort
	return nil
}

func autoConvert_v1alpha2_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in *AuditPolicyConfiguration, out *kubeadm.AuditPolicyConfiguration, s conversion.Scope) error {
	out.Path = in.Path
	out.LogDir = in.LogDir
	out.LogMaxAge = (*int32)(unsafe.Pointer(in.LogMaxAge))
	return nil
}

// Convert_v1alpha2_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in *AuditPolicyConfiguration, out *kubeadm.AuditPolicyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in, out, s)
}

func autoConvert_kubeadm_AuditPolicyConfiguration_To_v1alpha2_AuditPolicyConfiguration(in *kubeadm.AuditPolicyConfiguration, out *AuditPolicyConfiguration, s conversion.Scope) error {
	out.Path = in.Path
	out.LogDir = in.LogDir
	out.LogMaxAge = (*int32)(unsafe.Pointer(in.LogMaxAge))
	return nil
}

// Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha2_AuditPolicyConfiguration is an autogenerated conversion function.
func Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha2_AuditPolicyConfiguration(in *kubeadm.AuditPolicyConfiguration, out *AuditPolicyConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_AuditPolicyConfiguration_To_v1alpha2_AuditPolicyConfiguration(in, out, s)
}

func autoConvert_v1alpha2_BootstrapToken_To_kubeadm_BootstrapToken(in *BootstrapToken, out *kubeadm.BootstrapToken, s conversion.Scope) error {
	out.Token = (*kubeadm.BootstrapTokenString)(unsafe.Pointer(in.Token))
	out.Description = in.Description
	out.TTL = (*v1.Duration)(unsafe.Pointer(in.TTL))
	out.Expires = (*v1.Time)(unsafe.Pointer(in.Expires))
	out.Usages = *(*[]string)(unsafe.Pointer(&in.Usages))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1alpha2_BootstrapToken_To_kubeadm_BootstrapToken is an autogenerated conversion function.
func Convert_v1alpha2_BootstrapToken_To_kubeadm_BootstrapToken(in *BootstrapToken, out *kubeadm.BootstrapToken, s conversion.Scope) error {
	return autoConvert_v1alpha2_BootstrapToken_To_kubeadm_BootstrapToken(in, out, s)
}

func autoConvert_kubeadm_BootstrapToken_To_v1alpha2_BootstrapToken(in *kubeadm.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	out.Token = (*BootstrapTokenString)(unsafe.Pointer(in.Token))
	out.Description = in.Description
	out.TTL = (*v1.Duration)(unsafe.Pointer(in.TTL))
	out.Expires = (*v1.Time)(unsafe.Pointer(in.Expires))
	out.Usages = *(*[]string)(unsafe.Pointer(&in.Usages))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_kubeadm_BootstrapToken_To_v1alpha2_BootstrapToken is an autogenerated conversion function.
func Convert_kubeadm_BootstrapToken_To_v1alpha2_BootstrapToken(in *kubeadm.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	return autoConvert_kubeadm_BootstrapToken_To_v1alpha2_BootstrapToken(in, out, s)
}

func autoConvert_v1alpha2_BootstrapTokenString_To_kubeadm_BootstrapTokenString(in *BootstrapTokenString, out *kubeadm.BootstrapTokenString, s conversion.Scope) error {
	out.ID = in.ID
	out.Secret = in.Secret
	return nil
}

// Convert_v1alpha2_BootstrapTokenString_To_kubeadm_BootstrapTokenString is an autogenerated conversion function.
func Convert_v1alpha2_BootstrapTokenString_To_kubeadm_BootstrapTokenString(in *BootstrapTokenString, out *kubeadm.BootstrapTokenString, s conversion.Scope) error {
	return autoConvert_v1alpha2_BootstrapTokenString_To_kubeadm_BootstrapTokenString(in, out, s)
}

func autoConvert_kubeadm_BootstrapTokenString_To_v1alpha2_BootstrapTokenString(in *kubeadm.BootstrapTokenString, out *BootstrapTokenString, s conversion.Scope) error {
	out.ID = in.ID
	out.Secret = in.Secret
	return nil
}

// Convert_kubeadm_BootstrapTokenString_To_v1alpha2_BootstrapTokenString is an autogenerated conversion function.
func Convert_kubeadm_BootstrapTokenString_To_v1alpha2_BootstrapTokenString(in *kubeadm.BootstrapTokenString, out *BootstrapTokenString, s conversion.Scope) error {
	return autoConvert_kubeadm_BootstrapTokenString_To_v1alpha2_BootstrapTokenString(in, out, s)
}

func autoConvert_v1alpha2_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	out.Local = (*kubeadm.LocalEtcd)(unsafe.Pointer(in.Local))
	out.External = (*kubeadm.ExternalEtcd)(unsafe.Pointer(in.External))
	return nil
}

// Convert_v1alpha2_Etcd_To_kubeadm_Etcd is an autogenerated conversion function.
func Convert_v1alpha2_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	return autoConvert_v1alpha2_Etcd_To_kubeadm_Etcd(in, out, s)
}

func autoConvert_kubeadm_Etcd_To_v1alpha2_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	out.Local = (*LocalEtcd)(unsafe.Pointer(in.Local))
	out.External = (*ExternalEtcd)(unsafe.Pointer(in.External))
	return nil
}

// Convert_kubeadm_Etcd_To_v1alpha2_Etcd is an autogenerated conversion function.
func Convert_kubeadm_Etcd_To_v1alpha2_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	return autoConvert_kubeadm_Etcd_To_v1alpha2_Etcd(in, out, s)
}

func autoConvert_v1alpha2_ExternalEtcd_To_kubeadm_ExternalEtcd(in *ExternalEtcd, out *kubeadm.ExternalEtcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_v1alpha2_ExternalEtcd_To_kubeadm_ExternalEtcd is an autogenerated conversion function.
func Convert_v1alpha2_ExternalEtcd_To_kubeadm_ExternalEtcd(in *ExternalEtcd, out *kubeadm.ExternalEtcd, s conversion.Scope) error {
	return autoConvert_v1alpha2_ExternalEtcd_To_kubeadm_ExternalEtcd(in, out, s)
}

func autoConvert_kubeadm_ExternalEtcd_To_v1alpha2_ExternalEtcd(in *kubeadm.ExternalEtcd, out *ExternalEtcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_kubeadm_ExternalEtcd_To_v1alpha2_ExternalEtcd is an autogenerated conversion function.
func Convert_kubeadm_ExternalEtcd_To_v1alpha2_ExternalEtcd(in *kubeadm.ExternalEtcd, out *ExternalEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_ExternalEtcd_To_v1alpha2_ExternalEtcd(in, out, s)
}

func autoConvert_v1alpha2_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.Writable = in.Writable
	out.PathType = corev1.HostPathType(in.PathType)
	return nil
}

// Convert_v1alpha2_HostPathMount_To_kubeadm_HostPathMount is an autogenerated conversion function.
func Convert_v1alpha2_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	return autoConvert_v1alpha2_HostPathMount_To_kubeadm_HostPathMount(in, out, s)
}

func autoConvert_kubeadm_HostPathMount_To_v1alpha2_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.Writable = in.Writable
	out.PathType = corev1.HostPathType(in.PathType)
	return nil
}

// Convert_kubeadm_HostPathMount_To_v1alpha2_HostPathMount is an autogenerated conversion function.
func Convert_kubeadm_HostPathMount_To_v1alpha2_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	return autoConvert_kubeadm_HostPathMount_To_v1alpha2_HostPathMount(in, out, s)
}

func autoConvert_v1alpha2_InitConfiguration_To_kubeadm_InitConfiguration(in *InitConfiguration, out *kubeadm.InitConfiguration, s conversion.Scope) error {
	out.BootstrapTokens = *(*[]kubeadm.BootstrapToken)(unsafe.Pointer(&in.BootstrapTokens))
	if err := Convert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	// WARNING: in.API requires manual conversion: does not exist in peer-type
	// WARNING: in.KubeProxy requires manual conversion: does not exist in peer-type
	// WARNING: in.Etcd requires manual conversion: does not exist in peer-type
	// WARNING: in.KubeletConfiguration requires manual conversion: does not exist in peer-type
	// WARNING: in.Networking requires manual conversion: does not exist in peer-type
	// WARNING: in.KubernetesVersion requires manual conversion: does not exist in peer-type
	// WARNING: in.APIServerExtraArgs requires manual conversion: does not exist in peer-type
	// WARNING: in.ControllerManagerExtraArgs requires manual conversion: does not exist in peer-type
	// WARNING: in.SchedulerExtraArgs requires manual conversion: does not exist in peer-type
	// WARNING: in.APIServerExtraVolumes requires manual conversion: does not exist in peer-type
	// WARNING: in.ControllerManagerExtraVolumes requires manual conversion: does not exist in peer-type
	// WARNING: in.SchedulerExtraVolumes requires manual conversion: does not exist in peer-type
	// WARNING: in.APIServerCertSANs requires manual conversion: does not exist in peer-type
	// WARNING: in.CertificatesDir requires manual conversion: does not exist in peer-type
	// WARNING: in.ImageRepository requires manual conversion: does not exist in peer-type
	// WARNING: in.UnifiedControlPlaneImage requires manual conversion: does not exist in peer-type
	// WARNING: in.AuditPolicyConfiguration requires manual conversion: does not exist in peer-type
	// WARNING: in.FeatureGates requires manual conversion: does not exist in peer-type
	// WARNING: in.ClusterName requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_kubeadm_InitConfiguration_To_v1alpha2_InitConfiguration(in *kubeadm.InitConfiguration, out *InitConfiguration, s conversion.Scope) error {
	// WARNING: in.ClusterConfiguration requires manual conversion: does not exist in peer-type
	out.BootstrapTokens = *(*[]BootstrapToken)(unsafe.Pointer(&in.BootstrapTokens))
	if err := Convert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_JoinConfiguration_To_kubeadm_JoinConfiguration(in *JoinConfiguration, out *kubeadm.JoinConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.DiscoveryTimeout = (*v1.Duration)(unsafe.Pointer(in.DiscoveryTimeout))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.ClusterName = in.ClusterName
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.ControlPlane = in.ControlPlane
	out.AdvertiseAddress = in.AdvertiseAddress
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha2_JoinConfiguration_To_kubeadm_JoinConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_JoinConfiguration_To_kubeadm_JoinConfiguration(in *JoinConfiguration, out *kubeadm.JoinConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_JoinConfiguration_To_kubeadm_JoinConfiguration(in, out, s)
}

func autoConvert_kubeadm_JoinConfiguration_To_v1alpha2_JoinConfiguration(in *kubeadm.JoinConfiguration, out *JoinConfiguration, s conversion.Scope) error {
	if err := Convert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.DiscoveryTimeout = (*v1.Duration)(unsafe.Pointer(in.DiscoveryTimeout))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.ClusterName = in.ClusterName
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.ControlPlane = in.ControlPlane
	out.AdvertiseAddress = in.AdvertiseAddress
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_kubeadm_JoinConfiguration_To_v1alpha2_JoinConfiguration is an autogenerated conversion function.
func Convert_kubeadm_JoinConfiguration_To_v1alpha2_JoinConfiguration(in *kubeadm.JoinConfiguration, out *JoinConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_JoinConfiguration_To_v1alpha2_JoinConfiguration(in, out, s)
}

func autoConvert_v1alpha2_LocalEtcd_To_kubeadm_LocalEtcd(in *LocalEtcd, out *kubeadm.LocalEtcd, s conversion.Scope) error {
	out.Image = in.Image
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_v1alpha2_LocalEtcd_To_kubeadm_LocalEtcd is an autogenerated conversion function.
func Convert_v1alpha2_LocalEtcd_To_kubeadm_LocalEtcd(in *LocalEtcd, out *kubeadm.LocalEtcd, s conversion.Scope) error {
	return autoConvert_v1alpha2_LocalEtcd_To_kubeadm_LocalEtcd(in, out, s)
}

func autoConvert_kubeadm_LocalEtcd_To_v1alpha2_LocalEtcd(in *kubeadm.LocalEtcd, out *LocalEtcd, s conversion.Scope) error {
	out.Image = in.Image
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_kubeadm_LocalEtcd_To_v1alpha2_LocalEtcd is an autogenerated conversion function.
func Convert_kubeadm_LocalEtcd_To_v1alpha2_LocalEtcd(in *kubeadm.LocalEtcd, out *LocalEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_LocalEtcd_To_v1alpha2_LocalEtcd(in, out, s)
}

func autoConvert_v1alpha2_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_v1alpha2_Networking_To_kubeadm_Networking is an autogenerated conversion function.
func Convert_v1alpha2_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	return autoConvert_v1alpha2_Networking_To_kubeadm_Networking(in, out, s)
}

func autoConvert_kubeadm_Networking_To_v1alpha2_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_kubeadm_Networking_To_v1alpha2_Networking is an autogenerated conversion function.
func Convert_kubeadm_Networking_To_v1alpha2_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	return autoConvert_kubeadm_Networking_To_v1alpha2_Networking(in, out, s)
}

func autoConvert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in *NodeRegistrationOptions, out *kubeadm.NodeRegistrationOptions, s conversion.Scope) error {
	out.Name = in.Name
	out.CRISocket = in.CRISocket
	out.Taints = *(*[]corev1.Taint)(unsafe.Pointer(&in.Taints))
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	return nil
}

// Convert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions is an autogenerated conversion function.
func Convert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in *NodeRegistrationOptions, out *kubeadm.NodeRegistrationOptions, s conversion.Scope) error {
	return autoConvert_v1alpha2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in, out, s)
}

func autoConvert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions(in *kubeadm.NodeRegistrationOptions, out *NodeRegistrationOptions, s conversion.Scope) error {
	out.Name = in.Name
	out.CRISocket = in.CRISocket
	out.Taints = *(*[]corev1.Taint)(unsafe.Pointer(&in.Taints))
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	return nil
}

// Convert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions is an autogenerated conversion function.
func Convert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions(in *kubeadm.NodeRegistrationOptions, out *NodeRegistrationOptions, s conversion.Scope) error {
	return autoConvert_kubeadm_NodeRegistrationOptions_To_v1alpha2_NodeRegistrationOptions(in, out, s)
}
