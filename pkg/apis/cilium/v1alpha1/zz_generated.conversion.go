//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1alpha1

import (
	unsafe "unsafe"

	cilium "github.com/gardener/gardener-extension-networking-cilium/pkg/apis/cilium"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BPFSocketLBHostnsOnly)(nil), (*cilium.BPFSocketLBHostnsOnly)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BPFSocketLBHostnsOnly_To_cilium_BPFSocketLBHostnsOnly(a.(*BPFSocketLBHostnsOnly), b.(*cilium.BPFSocketLBHostnsOnly), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.BPFSocketLBHostnsOnly)(nil), (*BPFSocketLBHostnsOnly)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_BPFSocketLBHostnsOnly_To_v1alpha1_BPFSocketLBHostnsOnly(a.(*cilium.BPFSocketLBHostnsOnly), b.(*BPFSocketLBHostnsOnly), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EgressGateway)(nil), (*cilium.EgressGateway)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EgressGateway_To_cilium_EgressGateway(a.(*EgressGateway), b.(*cilium.EgressGateway), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.EgressGateway)(nil), (*EgressGateway)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_EgressGateway_To_v1alpha1_EgressGateway(a.(*cilium.EgressGateway), b.(*EgressGateway), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Hubble)(nil), (*cilium.Hubble)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Hubble_To_cilium_Hubble(a.(*Hubble), b.(*cilium.Hubble), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.Hubble)(nil), (*Hubble)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_Hubble_To_v1alpha1_Hubble(a.(*cilium.Hubble), b.(*Hubble), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPv6)(nil), (*cilium.IPv6)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPv6_To_cilium_IPv6(a.(*IPv6), b.(*cilium.IPv6), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.IPv6)(nil), (*IPv6)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_IPv6_To_v1alpha1_IPv6(a.(*cilium.IPv6), b.(*IPv6), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeProxy)(nil), (*cilium.KubeProxy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeProxy_To_cilium_KubeProxy(a.(*KubeProxy), b.(*cilium.KubeProxy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.KubeProxy)(nil), (*KubeProxy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_KubeProxy_To_v1alpha1_KubeProxy(a.(*cilium.KubeProxy), b.(*KubeProxy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkConfig)(nil), (*cilium.NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkConfig_To_cilium_NetworkConfig(a.(*NetworkConfig), b.(*cilium.NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.NetworkConfig)(nil), (*NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_NetworkConfig_To_v1alpha1_NetworkConfig(a.(*cilium.NetworkConfig), b.(*NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Nodeport)(nil), (*cilium.Nodeport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Nodeport_To_cilium_Nodeport(a.(*Nodeport), b.(*cilium.Nodeport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.Nodeport)(nil), (*Nodeport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_Nodeport_To_v1alpha1_Nodeport(a.(*cilium.Nodeport), b.(*Nodeport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Overlay)(nil), (*cilium.Overlay)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Overlay_To_cilium_Overlay(a.(*Overlay), b.(*cilium.Overlay), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.Overlay)(nil), (*Overlay)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_Overlay_To_v1alpha1_Overlay(a.(*cilium.Overlay), b.(*Overlay), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SnatOutOfCluster)(nil), (*cilium.SnatOutOfCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SnatOutOfCluster_To_cilium_SnatOutOfCluster(a.(*SnatOutOfCluster), b.(*cilium.SnatOutOfCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.SnatOutOfCluster)(nil), (*SnatOutOfCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_SnatOutOfCluster_To_v1alpha1_SnatOutOfCluster(a.(*cilium.SnatOutOfCluster), b.(*SnatOutOfCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SnatToUpstreamDNS)(nil), (*cilium.SnatToUpstreamDNS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SnatToUpstreamDNS_To_cilium_SnatToUpstreamDNS(a.(*SnatToUpstreamDNS), b.(*cilium.SnatToUpstreamDNS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cilium.SnatToUpstreamDNS)(nil), (*SnatToUpstreamDNS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cilium_SnatToUpstreamDNS_To_v1alpha1_SnatToUpstreamDNS(a.(*cilium.SnatToUpstreamDNS), b.(*SnatToUpstreamDNS), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_BPFSocketLBHostnsOnly_To_cilium_BPFSocketLBHostnsOnly(in *BPFSocketLBHostnsOnly, out *cilium.BPFSocketLBHostnsOnly, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_BPFSocketLBHostnsOnly_To_cilium_BPFSocketLBHostnsOnly is an autogenerated conversion function.
func Convert_v1alpha1_BPFSocketLBHostnsOnly_To_cilium_BPFSocketLBHostnsOnly(in *BPFSocketLBHostnsOnly, out *cilium.BPFSocketLBHostnsOnly, s conversion.Scope) error {
	return autoConvert_v1alpha1_BPFSocketLBHostnsOnly_To_cilium_BPFSocketLBHostnsOnly(in, out, s)
}

func autoConvert_cilium_BPFSocketLBHostnsOnly_To_v1alpha1_BPFSocketLBHostnsOnly(in *cilium.BPFSocketLBHostnsOnly, out *BPFSocketLBHostnsOnly, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_BPFSocketLBHostnsOnly_To_v1alpha1_BPFSocketLBHostnsOnly is an autogenerated conversion function.
func Convert_cilium_BPFSocketLBHostnsOnly_To_v1alpha1_BPFSocketLBHostnsOnly(in *cilium.BPFSocketLBHostnsOnly, out *BPFSocketLBHostnsOnly, s conversion.Scope) error {
	return autoConvert_cilium_BPFSocketLBHostnsOnly_To_v1alpha1_BPFSocketLBHostnsOnly(in, out, s)
}

func autoConvert_v1alpha1_EgressGateway_To_cilium_EgressGateway(in *EgressGateway, out *cilium.EgressGateway, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_EgressGateway_To_cilium_EgressGateway is an autogenerated conversion function.
func Convert_v1alpha1_EgressGateway_To_cilium_EgressGateway(in *EgressGateway, out *cilium.EgressGateway, s conversion.Scope) error {
	return autoConvert_v1alpha1_EgressGateway_To_cilium_EgressGateway(in, out, s)
}

func autoConvert_cilium_EgressGateway_To_v1alpha1_EgressGateway(in *cilium.EgressGateway, out *EgressGateway, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_EgressGateway_To_v1alpha1_EgressGateway is an autogenerated conversion function.
func Convert_cilium_EgressGateway_To_v1alpha1_EgressGateway(in *cilium.EgressGateway, out *EgressGateway, s conversion.Scope) error {
	return autoConvert_cilium_EgressGateway_To_v1alpha1_EgressGateway(in, out, s)
}

func autoConvert_v1alpha1_Hubble_To_cilium_Hubble(in *Hubble, out *cilium.Hubble, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_Hubble_To_cilium_Hubble is an autogenerated conversion function.
func Convert_v1alpha1_Hubble_To_cilium_Hubble(in *Hubble, out *cilium.Hubble, s conversion.Scope) error {
	return autoConvert_v1alpha1_Hubble_To_cilium_Hubble(in, out, s)
}

func autoConvert_cilium_Hubble_To_v1alpha1_Hubble(in *cilium.Hubble, out *Hubble, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_Hubble_To_v1alpha1_Hubble is an autogenerated conversion function.
func Convert_cilium_Hubble_To_v1alpha1_Hubble(in *cilium.Hubble, out *Hubble, s conversion.Scope) error {
	return autoConvert_cilium_Hubble_To_v1alpha1_Hubble(in, out, s)
}

func autoConvert_v1alpha1_IPv6_To_cilium_IPv6(in *IPv6, out *cilium.IPv6, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_IPv6_To_cilium_IPv6 is an autogenerated conversion function.
func Convert_v1alpha1_IPv6_To_cilium_IPv6(in *IPv6, out *cilium.IPv6, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPv6_To_cilium_IPv6(in, out, s)
}

func autoConvert_cilium_IPv6_To_v1alpha1_IPv6(in *cilium.IPv6, out *IPv6, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_IPv6_To_v1alpha1_IPv6 is an autogenerated conversion function.
func Convert_cilium_IPv6_To_v1alpha1_IPv6(in *cilium.IPv6, out *IPv6, s conversion.Scope) error {
	return autoConvert_cilium_IPv6_To_v1alpha1_IPv6(in, out, s)
}

func autoConvert_v1alpha1_KubeProxy_To_cilium_KubeProxy(in *KubeProxy, out *cilium.KubeProxy, s conversion.Scope) error {
	out.ServiceHost = (*string)(unsafe.Pointer(in.ServiceHost))
	out.ServicePort = (*int32)(unsafe.Pointer(in.ServicePort))
	return nil
}

// Convert_v1alpha1_KubeProxy_To_cilium_KubeProxy is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxy_To_cilium_KubeProxy(in *KubeProxy, out *cilium.KubeProxy, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxy_To_cilium_KubeProxy(in, out, s)
}

func autoConvert_cilium_KubeProxy_To_v1alpha1_KubeProxy(in *cilium.KubeProxy, out *KubeProxy, s conversion.Scope) error {
	out.ServiceHost = (*string)(unsafe.Pointer(in.ServiceHost))
	out.ServicePort = (*int32)(unsafe.Pointer(in.ServicePort))
	return nil
}

// Convert_cilium_KubeProxy_To_v1alpha1_KubeProxy is an autogenerated conversion function.
func Convert_cilium_KubeProxy_To_v1alpha1_KubeProxy(in *cilium.KubeProxy, out *KubeProxy, s conversion.Scope) error {
	return autoConvert_cilium_KubeProxy_To_v1alpha1_KubeProxy(in, out, s)
}

func autoConvert_v1alpha1_NetworkConfig_To_cilium_NetworkConfig(in *NetworkConfig, out *cilium.NetworkConfig, s conversion.Scope) error {
	out.Debug = (*bool)(unsafe.Pointer(in.Debug))
	out.PSPEnabled = (*bool)(unsafe.Pointer(in.PSPEnabled))
	out.KubeProxy = (*cilium.KubeProxy)(unsafe.Pointer(in.KubeProxy))
	out.Hubble = (*cilium.Hubble)(unsafe.Pointer(in.Hubble))
	out.TunnelMode = (*cilium.TunnelMode)(unsafe.Pointer(in.TunnelMode))
	out.Store = (*cilium.Store)(unsafe.Pointer(in.Store))
	out.IPv6 = (*cilium.IPv6)(unsafe.Pointer(in.IPv6))
	out.BPFSocketLBHostnsOnly = (*cilium.BPFSocketLBHostnsOnly)(unsafe.Pointer(in.BPFSocketLBHostnsOnly))
	out.EgressGateway = (*cilium.EgressGateway)(unsafe.Pointer(in.EgressGateway))
	out.MTU = (*int)(unsafe.Pointer(in.MTU))
	out.Devices = *(*[]string)(unsafe.Pointer(&in.Devices))
	out.LoadBalancingMode = (*cilium.LoadBalancingMode)(unsafe.Pointer(in.LoadBalancingMode))
	out.IPv4NativeRoutingCIDREnabled = (*bool)(unsafe.Pointer(in.IPv4NativeRoutingCIDREnabled))
	out.Overlay = (*cilium.Overlay)(unsafe.Pointer(in.Overlay))
	out.SnatToUpstreamDNS = (*cilium.SnatToUpstreamDNS)(unsafe.Pointer(in.SnatToUpstreamDNS))
	out.SnatOutOfCluster = (*cilium.SnatOutOfCluster)(unsafe.Pointer(in.SnatOutOfCluster))
	return nil
}

// Convert_v1alpha1_NetworkConfig_To_cilium_NetworkConfig is an autogenerated conversion function.
func Convert_v1alpha1_NetworkConfig_To_cilium_NetworkConfig(in *NetworkConfig, out *cilium.NetworkConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkConfig_To_cilium_NetworkConfig(in, out, s)
}

func autoConvert_cilium_NetworkConfig_To_v1alpha1_NetworkConfig(in *cilium.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	out.Debug = (*bool)(unsafe.Pointer(in.Debug))
	out.PSPEnabled = (*bool)(unsafe.Pointer(in.PSPEnabled))
	out.KubeProxy = (*KubeProxy)(unsafe.Pointer(in.KubeProxy))
	out.Hubble = (*Hubble)(unsafe.Pointer(in.Hubble))
	out.TunnelMode = (*TunnelMode)(unsafe.Pointer(in.TunnelMode))
	out.Store = (*Store)(unsafe.Pointer(in.Store))
	out.IPv6 = (*IPv6)(unsafe.Pointer(in.IPv6))
	out.BPFSocketLBHostnsOnly = (*BPFSocketLBHostnsOnly)(unsafe.Pointer(in.BPFSocketLBHostnsOnly))
	out.EgressGateway = (*EgressGateway)(unsafe.Pointer(in.EgressGateway))
	out.MTU = (*int)(unsafe.Pointer(in.MTU))
	out.Devices = *(*[]string)(unsafe.Pointer(&in.Devices))
	out.LoadBalancingMode = (*LoadBalancingMode)(unsafe.Pointer(in.LoadBalancingMode))
	out.IPv4NativeRoutingCIDREnabled = (*bool)(unsafe.Pointer(in.IPv4NativeRoutingCIDREnabled))
	out.Overlay = (*Overlay)(unsafe.Pointer(in.Overlay))
	out.SnatToUpstreamDNS = (*SnatToUpstreamDNS)(unsafe.Pointer(in.SnatToUpstreamDNS))
	out.SnatOutOfCluster = (*SnatOutOfCluster)(unsafe.Pointer(in.SnatOutOfCluster))
	return nil
}

// Convert_cilium_NetworkConfig_To_v1alpha1_NetworkConfig is an autogenerated conversion function.
func Convert_cilium_NetworkConfig_To_v1alpha1_NetworkConfig(in *cilium.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	return autoConvert_cilium_NetworkConfig_To_v1alpha1_NetworkConfig(in, out, s)
}

func autoConvert_v1alpha1_Nodeport_To_cilium_Nodeport(in *Nodeport, out *cilium.Nodeport, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Mode = cilium.NodePortMode(in.Mode)
	return nil
}

// Convert_v1alpha1_Nodeport_To_cilium_Nodeport is an autogenerated conversion function.
func Convert_v1alpha1_Nodeport_To_cilium_Nodeport(in *Nodeport, out *cilium.Nodeport, s conversion.Scope) error {
	return autoConvert_v1alpha1_Nodeport_To_cilium_Nodeport(in, out, s)
}

func autoConvert_cilium_Nodeport_To_v1alpha1_Nodeport(in *cilium.Nodeport, out *Nodeport, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Mode = NodePortMode(in.Mode)
	return nil
}

// Convert_cilium_Nodeport_To_v1alpha1_Nodeport is an autogenerated conversion function.
func Convert_cilium_Nodeport_To_v1alpha1_Nodeport(in *cilium.Nodeport, out *Nodeport, s conversion.Scope) error {
	return autoConvert_cilium_Nodeport_To_v1alpha1_Nodeport(in, out, s)
}

func autoConvert_v1alpha1_Overlay_To_cilium_Overlay(in *Overlay, out *cilium.Overlay, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_Overlay_To_cilium_Overlay is an autogenerated conversion function.
func Convert_v1alpha1_Overlay_To_cilium_Overlay(in *Overlay, out *cilium.Overlay, s conversion.Scope) error {
	return autoConvert_v1alpha1_Overlay_To_cilium_Overlay(in, out, s)
}

func autoConvert_cilium_Overlay_To_v1alpha1_Overlay(in *cilium.Overlay, out *Overlay, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_Overlay_To_v1alpha1_Overlay is an autogenerated conversion function.
func Convert_cilium_Overlay_To_v1alpha1_Overlay(in *cilium.Overlay, out *Overlay, s conversion.Scope) error {
	return autoConvert_cilium_Overlay_To_v1alpha1_Overlay(in, out, s)
}

func autoConvert_v1alpha1_SnatOutOfCluster_To_cilium_SnatOutOfCluster(in *SnatOutOfCluster, out *cilium.SnatOutOfCluster, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_SnatOutOfCluster_To_cilium_SnatOutOfCluster is an autogenerated conversion function.
func Convert_v1alpha1_SnatOutOfCluster_To_cilium_SnatOutOfCluster(in *SnatOutOfCluster, out *cilium.SnatOutOfCluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_SnatOutOfCluster_To_cilium_SnatOutOfCluster(in, out, s)
}

func autoConvert_cilium_SnatOutOfCluster_To_v1alpha1_SnatOutOfCluster(in *cilium.SnatOutOfCluster, out *SnatOutOfCluster, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_SnatOutOfCluster_To_v1alpha1_SnatOutOfCluster is an autogenerated conversion function.
func Convert_cilium_SnatOutOfCluster_To_v1alpha1_SnatOutOfCluster(in *cilium.SnatOutOfCluster, out *SnatOutOfCluster, s conversion.Scope) error {
	return autoConvert_cilium_SnatOutOfCluster_To_v1alpha1_SnatOutOfCluster(in, out, s)
}

func autoConvert_v1alpha1_SnatToUpstreamDNS_To_cilium_SnatToUpstreamDNS(in *SnatToUpstreamDNS, out *cilium.SnatToUpstreamDNS, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_SnatToUpstreamDNS_To_cilium_SnatToUpstreamDNS is an autogenerated conversion function.
func Convert_v1alpha1_SnatToUpstreamDNS_To_cilium_SnatToUpstreamDNS(in *SnatToUpstreamDNS, out *cilium.SnatToUpstreamDNS, s conversion.Scope) error {
	return autoConvert_v1alpha1_SnatToUpstreamDNS_To_cilium_SnatToUpstreamDNS(in, out, s)
}

func autoConvert_cilium_SnatToUpstreamDNS_To_v1alpha1_SnatToUpstreamDNS(in *cilium.SnatToUpstreamDNS, out *SnatToUpstreamDNS, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_cilium_SnatToUpstreamDNS_To_v1alpha1_SnatToUpstreamDNS is an autogenerated conversion function.
func Convert_cilium_SnatToUpstreamDNS_To_v1alpha1_SnatToUpstreamDNS(in *cilium.SnatToUpstreamDNS, out *SnatToUpstreamDNS, s conversion.Scope) error {
	return autoConvert_cilium_SnatToUpstreamDNS_To_v1alpha1_SnatToUpstreamDNS(in, out, s)
}
