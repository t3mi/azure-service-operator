// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240402preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM, ManagedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM runs a test to see if a specific instance of ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM(subject ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM instances for property testing - lazily instantiated
// by ManagedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator()
var managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator gopter.Gen

// ManagedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator returns a generator of ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM instances for property testing.
// We first initialize managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator() gopter.Gen {
	if managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator != nil {
		return managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM(generators)
	managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM(generators)
	managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM{}), generators)

	return managedClusters_TrustedAccessRoleBinding_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusters_TrustedAccessRoleBinding_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TrustedAccessRoleBindingProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_TrustedAccessRoleBindingProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrustedAccessRoleBindingProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrustedAccessRoleBindingProperties_STATUS_ARM, TrustedAccessRoleBindingProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrustedAccessRoleBindingProperties_STATUS_ARM runs a test to see if a specific instance of TrustedAccessRoleBindingProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTrustedAccessRoleBindingProperties_STATUS_ARM(subject TrustedAccessRoleBindingProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrustedAccessRoleBindingProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of TrustedAccessRoleBindingProperties_STATUS_ARM instances for property testing - lazily instantiated by
// TrustedAccessRoleBindingProperties_STATUS_ARMGenerator()
var trustedAccessRoleBindingProperties_STATUS_ARMGenerator gopter.Gen

// TrustedAccessRoleBindingProperties_STATUS_ARMGenerator returns a generator of TrustedAccessRoleBindingProperties_STATUS_ARM instances for property testing.
func TrustedAccessRoleBindingProperties_STATUS_ARMGenerator() gopter.Gen {
	if trustedAccessRoleBindingProperties_STATUS_ARMGenerator != nil {
		return trustedAccessRoleBindingProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrustedAccessRoleBindingProperties_STATUS_ARM(generators)
	trustedAccessRoleBindingProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TrustedAccessRoleBindingProperties_STATUS_ARM{}), generators)

	return trustedAccessRoleBindingProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTrustedAccessRoleBindingProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrustedAccessRoleBindingProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_ARM_Canceled,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_ARM_Deleting,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_ARM_Failed,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_ARM_Succeeded,
		TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_ARM_Updating))
	gens["Roles"] = gen.SliceOf(gen.AlphaString())
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
}
