// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// InventoryStatusApplyConfiguration represents an declarative configuration of the InventoryStatus type for use
// with apply.
type InventoryStatusApplyConfiguration struct {
	InventoryStatuses *InventoryStatusesApplyConfiguration `json:"inventoryStatuses,omitempty"`
}

// InventoryStatusApplyConfiguration constructs an declarative configuration of the InventoryStatus type for use with
// apply.
func InventoryStatus() *InventoryStatusApplyConfiguration {
	return &InventoryStatusApplyConfiguration{}
}

// WithInventoryStatuses sets the InventoryStatuses field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InventoryStatuses field is set to the value of the last call.
func (b *InventoryStatusApplyConfiguration) WithInventoryStatuses(value *InventoryStatusesApplyConfiguration) *InventoryStatusApplyConfiguration {
	b.InventoryStatuses = value
	return b
}