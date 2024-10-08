// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

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

func Test_DiskAccess_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess_Spec_ARM, DiskAccess_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess_Spec_ARM runs a test to see if a specific instance of DiskAccess_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess_Spec_ARM(subject DiskAccess_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess_Spec_ARM
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

// Generator of DiskAccess_Spec_ARM instances for property testing - lazily instantiated by
// DiskAccess_Spec_ARMGenerator()
var diskAccess_Spec_ARMGenerator gopter.Gen

// DiskAccess_Spec_ARMGenerator returns a generator of DiskAccess_Spec_ARM instances for property testing.
// We first initialize diskAccess_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskAccess_Spec_ARMGenerator() gopter.Gen {
	if diskAccess_Spec_ARMGenerator != nil {
		return diskAccess_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_Spec_ARM(generators)
	diskAccess_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DiskAccess_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDiskAccess_Spec_ARM(generators)
	diskAccess_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DiskAccess_Spec_ARM{}), generators)

	return diskAccess_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskAccess_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskAccess_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskAccess_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
}

func Test_ExtendedLocation_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocation_ARM, ExtendedLocation_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocation_ARM runs a test to see if a specific instance of ExtendedLocation_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocation_ARM(subject ExtendedLocation_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation_ARM
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

// Generator of ExtendedLocation_ARM instances for property testing - lazily instantiated by
// ExtendedLocation_ARMGenerator()
var extendedLocation_ARMGenerator gopter.Gen

// ExtendedLocation_ARMGenerator returns a generator of ExtendedLocation_ARM instances for property testing.
func ExtendedLocation_ARMGenerator() gopter.Gen {
	if extendedLocation_ARMGenerator != nil {
		return extendedLocation_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocation_ARM(generators)
	extendedLocation_ARMGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation_ARM{}), generators)

	return extendedLocation_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocation_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocation_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_ARM_EdgeZone))
}
