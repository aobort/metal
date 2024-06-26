// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// OOBSpecApplyConfiguration represents an declarative configuration of the OOBSpec type for use
// with apply.
type OOBSpecApplyConfiguration struct {
	MACAddress      *string                            `json:"macAddress,omitempty"`
	EndpointRef     *v1.LocalObjectReference           `json:"endpointRef,omitempty"`
	SecretRef       *v1.LocalObjectReference           `json:"secretRef,omitempty"`
	Protocol        *ProtocolApplyConfiguration        `json:"protocol,omitempty"`
	Flags           map[string]string                  `json:"flags,omitempty"`
	ConsoleProtocol *ConsoleProtocolApplyConfiguration `json:"consoleProtocol,omitempty"`
}

// OOBSpecApplyConfiguration constructs an declarative configuration of the OOBSpec type for use with
// apply.
func OOBSpec() *OOBSpecApplyConfiguration {
	return &OOBSpecApplyConfiguration{}
}

// WithMACAddress sets the MACAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MACAddress field is set to the value of the last call.
func (b *OOBSpecApplyConfiguration) WithMACAddress(value string) *OOBSpecApplyConfiguration {
	b.MACAddress = &value
	return b
}

// WithEndpointRef sets the EndpointRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndpointRef field is set to the value of the last call.
func (b *OOBSpecApplyConfiguration) WithEndpointRef(value v1.LocalObjectReference) *OOBSpecApplyConfiguration {
	b.EndpointRef = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *OOBSpecApplyConfiguration) WithSecretRef(value v1.LocalObjectReference) *OOBSpecApplyConfiguration {
	b.SecretRef = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *OOBSpecApplyConfiguration) WithProtocol(value *ProtocolApplyConfiguration) *OOBSpecApplyConfiguration {
	b.Protocol = value
	return b
}

// WithFlags puts the entries into the Flags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Flags field,
// overwriting an existing map entries in Flags field with the same key.
func (b *OOBSpecApplyConfiguration) WithFlags(entries map[string]string) *OOBSpecApplyConfiguration {
	if b.Flags == nil && len(entries) > 0 {
		b.Flags = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Flags[k] = v
	}
	return b
}

// WithConsoleProtocol sets the ConsoleProtocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConsoleProtocol field is set to the value of the last call.
func (b *OOBSpecApplyConfiguration) WithConsoleProtocol(value *ConsoleProtocolApplyConfiguration) *OOBSpecApplyConfiguration {
	b.ConsoleProtocol = value
	return b
}
