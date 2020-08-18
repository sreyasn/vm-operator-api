// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineClassBindingSpec defines the desired state of VirtualMachineClassBinding
type VirtualMachineClassBindingSpec struct {
}

// VirtualMachineClassBindingStatus defines the observed state of VirtualMachineClassBinding
type VirtualMachineClassBindingStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vmclassbinding

// VirtualMachineClassBinding is a binding object responsible for
// defining a VirtualMachineClass and a namespace the class is associated with
type VirtualMachineClassBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineClassBindingSpec    `json:"spec,omitempty"`
	Status VirtualMachineClassBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineClassBindingList contains a list of VirtualMachineClassBinding
type VirtualMachineClassBindingList  struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentSource `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&VirtualMachineClassBinding{}, &VirtualMachineClassBindingList{})
}