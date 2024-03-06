/*
Copyright 2022 Gravitational, Inc.

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

package modules

import (
	"context"
	"crypto"
	"testing"
	"time"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/api/utils/keys"
)

// TestModules implements the Modules interface for testing.
//
// Setting Test* fields will return those values from interface
// methods. IsBoringBinary and PrintVersion functions return the
// same values from default modules.
//
// See SetTestModules for an example.
type TestModules struct {
	// TestBuildType is returned from the BuiltType function.
	TestBuildType string
	// TestFeatures is returned from the Features function.
	TestFeatures Features

	defaultModules

	// MockAttestationData is fake attestation data to return
	// during tests when hardware key support is enabled.
	MockAttestationData *keys.AttestationData
}

// SetTestModules sets the value returned from GetModules to testModules
// and reverts the change in the test cleanup function.
// It must not be used in parallel tests.
//
//	func TestWithFakeModules(t *testing.T) {
//	   modules.SetTestModules(t, &modules.TestModules{
//	     TestBuildType: modules.BuildEnterprise,
//	     TestFeatures: modules.Features{
//	        Cloud: true,
//	     },
//	   })
//
//	   // test implementation
//
//	   // cleanup will revert module changes after test completes
//	}
func SetTestModules(t *testing.T, testModules Modules) {
	defaultModules := GetModules()
	t.Cleanup(func() { SetModules(defaultModules) })
	t.Setenv("TELEPORT_TEST_NOT_SAFE_FOR_PARALLEL", "true")
	SetModules(testModules)
}

// PrintVersion prints teleport version
func (m *TestModules) PrintVersion() {
	m.defaultModules.PrintVersion()
}

// IsBoringBinary checks if the binary was compiled with BoringCrypto.
func (m *TestModules) IsBoringBinary() bool {
	return m.defaultModules.IsBoringBinary()
}

// Features returns supported features.
func (m *TestModules) Features() Features {
	return m.TestFeatures
}

// BuildType returns build type (OSS or Enterprise).
func (m *TestModules) BuildType() string {
	return m.TestBuildType
}

// AttestHardwareKey attests a hardware key.
func (m *TestModules) AttestHardwareKey(ctx context.Context, obj interface{}, as *keys.AttestationStatement, pk crypto.PublicKey, d time.Duration) (*keys.AttestationData, error) {
	// If the given attestation statement is non-nil, that means the client passed it to the
	// server to attest the key. If MockAttestationData is set, return that to attest the key.
	if as != nil && m.MockAttestationData != nil {
		return m.MockAttestationData, nil
	}
	return nil, trace.NotFound("no attestation data for the given key")
}
