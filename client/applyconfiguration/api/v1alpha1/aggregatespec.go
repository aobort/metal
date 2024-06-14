// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// AggregateSpecApplyConfiguration represents an declarative configuration of the AggregateSpec type for use
// with apply.
type AggregateSpecApplyConfiguration struct {
	Aggregates []AggregateItemApplyConfiguration `json:"aggregates,omitempty"`
}

// AggregateSpecApplyConfiguration constructs an declarative configuration of the AggregateSpec type for use with
// apply.
func AggregateSpec() *AggregateSpecApplyConfiguration {
	return &AggregateSpecApplyConfiguration{}
}

// WithAggregates adds the given value to the Aggregates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Aggregates field.
func (b *AggregateSpecApplyConfiguration) WithAggregates(values ...*AggregateItemApplyConfiguration) *AggregateSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAggregates")
		}
		b.Aggregates = append(b.Aggregates, *values[i])
	}
	return b
}