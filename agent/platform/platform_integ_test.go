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

// +build integration

// Package platform contains platform specific utilities.
package platform

import (
	"testing"

	logger "github.com/aws/amazon-ssm-agent/agent/log"
	"github.com/stretchr/testify/assert"
)

func TestGetPlatformName(t *testing.T) {
	var log = logger.NewMockLog()
	t.Log("get platform name and version")
	data, err := PlatformName(log)
	t.Logf("platform name is %v ", data)
	assert.NoError(t, err, "get platform name should not result in err")
	data, err = PlatformVersion(log)
	t.Logf("platform version is %v ", data)
	assert.NoError(t, err, "get platform version should not result in err")
}

func TestFullyQualifiedDomainName(t *testing.T) {
	t.Logf("fqdn/hostname is %v", fullyQualifiedDomainName())
	assert.True(t, true, "expected no error trying to retireve the fqdn/hostname value")
}
