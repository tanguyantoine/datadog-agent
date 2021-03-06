// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2020 Datadog, Inc.

package traps

import (
	"testing"

	"github.com/soniah/gosnmp"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	Configure(t, Config{
		Port:             1234,
		CommunityStrings: []string{"public"},
	})
	config, err := ReadConfig()
	assert.NoError(t, err)
	assert.Equal(t, uint16(1234), config.Port)
	assert.Equal(t, defaultStopTimeout, config.StopTimeout)

	params := config.BuildV2Params()
	assert.Equal(t, uint16(1234), params.Port)
	assert.Equal(t, gosnmp.Version2c, params.Version)
	assert.Equal(t, "udp", params.Transport)
	assert.NotNil(t, params.Logger)
}

func TestDefaultPort(t *testing.T) {
	Configure(t, Config{
		CommunityStrings: []string{"public"},
	})
	config, err := ReadConfig()
	assert.NoError(t, err)
	assert.Equal(t, defaultPort, config.Port)
}

func TestCommunityStringsEmpty(t *testing.T) {
	Configure(t, Config{
		CommunityStrings: []string{},
	})
	_, err := ReadConfig()
	assert.Error(t, err)
}

func TestCommunityStringsMissing(t *testing.T) {
	Configure(t, Config{})
	_, err := ReadConfig()
	assert.Error(t, err)
}
