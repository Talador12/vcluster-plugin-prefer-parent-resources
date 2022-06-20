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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageVersionConditionApplyConfiguration represents an declarative configuration of the StorageVersionCondition type for use
// with apply.
type StorageVersionConditionApplyConfiguration struct {
	Type               *v1alpha1.StorageVersionConditionType `json:"type,omitempty"`
	Status             *v1alpha1.ConditionStatus             `json:"status,omitempty"`
	ObservedGeneration *int64                                `json:"observedGeneration,omitempty"`
	LastTransitionTime *v1.Time                              `json:"lastTransitionTime,omitempty"`
	Reason             *string                               `json:"reason,omitempty"`
	Message            *string                               `json:"message,omitempty"`
}

// StorageVersionConditionApplyConfiguration constructs an declarative configuration of the StorageVersionCondition type for use with
// apply.
func StorageVersionCondition() *StorageVersionConditionApplyConfiguration {
	return &StorageVersionConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *StorageVersionConditionApplyConfiguration) WithType(value v1alpha1.StorageVersionConditionType) *StorageVersionConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *StorageVersionConditionApplyConfiguration) WithStatus(value v1alpha1.ConditionStatus) *StorageVersionConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *StorageVersionConditionApplyConfiguration) WithObservedGeneration(value int64) *StorageVersionConditionApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *StorageVersionConditionApplyConfiguration) WithLastTransitionTime(value v1.Time) *StorageVersionConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *StorageVersionConditionApplyConfiguration) WithReason(value string) *StorageVersionConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *StorageVersionConditionApplyConfiguration) WithMessage(value string) *StorageVersionConditionApplyConfiguration {
	b.Message = &value
	return b
}
