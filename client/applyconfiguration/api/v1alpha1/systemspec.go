// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// SystemSpecApplyConfiguration represents an declarative configuration of the SystemSpec type for use
// with apply.
type SystemSpecApplyConfiguration struct {
	ID           *string `json:"id,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	ProductSKU   *string `json:"productSku,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
}

// SystemSpecApplyConfiguration constructs an declarative configuration of the SystemSpec type for use with
// apply.
func SystemSpec() *SystemSpecApplyConfiguration {
	return &SystemSpecApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *SystemSpecApplyConfiguration) WithID(value string) *SystemSpecApplyConfiguration {
	b.ID = &value
	return b
}

// WithManufacturer sets the Manufacturer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Manufacturer field is set to the value of the last call.
func (b *SystemSpecApplyConfiguration) WithManufacturer(value string) *SystemSpecApplyConfiguration {
	b.Manufacturer = &value
	return b
}

// WithProductSKU sets the ProductSKU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProductSKU field is set to the value of the last call.
func (b *SystemSpecApplyConfiguration) WithProductSKU(value string) *SystemSpecApplyConfiguration {
	b.ProductSKU = &value
	return b
}

// WithSerialNumber sets the SerialNumber field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SerialNumber field is set to the value of the last call.
func (b *SystemSpecApplyConfiguration) WithSerialNumber(value string) *SystemSpecApplyConfiguration {
	b.SerialNumber = &value
	return b
}
