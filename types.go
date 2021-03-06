/* vim:set sw=8 ts=8 noet:
 *
 * Copyright (c) 2017 Torchbox Ltd.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely. This software is provided 'as-is', without any express or implied
 * warranty.
 */

package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DatabaseSpec struct {
	Type   string `json:"type"`
	Secret string `json:"secretName"`
	Class  string `json:"class"`
}

type DatabaseStatus struct {
	Phase  string `json:"phase"`
	Error  string `json:"error,omitempty"`
	Server string `json:"server,omitempty"`
}

type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   DatabaseSpec   `json:"spec"`
	Status DatabaseStatus `json:"status,omitempty"`
}

type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Database `json:"items"`
}
