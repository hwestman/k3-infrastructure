/*
Copyright 2019 Rancher Labs, Inc.

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

// Code generated by codegen. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=upgrade.cattle.io
package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlanList is a list of Plan resources
type PlanList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata"`

    Items []Plan `json:"items"`
}

func NewPlan(namespace, name string, obj Plan) *Plan {
    obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Plan").ToAPIVersionAndKind()
    obj.Name = name
    obj.Namespace = namespace
    return &obj
}
