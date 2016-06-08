// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Amazon Software License (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/asl/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// +build windows

// Package main represents the entry point of the ssm agent updater.
package main

import (
	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/update/processor"
)

// updateRoot returns the platform specific path to update artifacts
func updateRoot(detail *processor.UpdateDetail) error {
	detail.UpdateRoot = appconfig.UpdaterArtifactsRoot
	return nil
}
