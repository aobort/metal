// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/metal/api/v1alpha1"
)

// AggregateItemApplyConfiguration represents an declarative configuration of the AggregateItem type for use
// with apply.
type AggregateItemApplyConfiguration struct {
	SourcePath *v1alpha1.JSONPath      `json:"sourcePath,omitempty"`
	TargetPath *v1alpha1.JSONPath      `json:"targetPath,omitempty"`
	Aggregate  *v1alpha1.AggregateType `json:"aggregate,omitempty"`
}

// AggregateItemApplyConfiguration constructs an declarative configuration of the AggregateItem type for use with
// apply.
func AggregateItem() *AggregateItemApplyConfiguration {
	return &AggregateItemApplyConfiguration{}
}

// WithSourcePath sets the SourcePath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SourcePath field is set to the value of the last call.
func (b *AggregateItemApplyConfiguration) WithSourcePath(value v1alpha1.JSONPath) *AggregateItemApplyConfiguration {
	b.SourcePath = &value
	return b
}

// WithTargetPath sets the TargetPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetPath field is set to the value of the last call.
func (b *AggregateItemApplyConfiguration) WithTargetPath(value v1alpha1.JSONPath) *AggregateItemApplyConfiguration {
	b.TargetPath = &value
	return b
}

// WithAggregate sets the Aggregate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Aggregate field is set to the value of the last call.
func (b *AggregateItemApplyConfiguration) WithAggregate(value v1alpha1.AggregateType) *AggregateItemApplyConfiguration {
	b.Aggregate = &value
	return b
}
