package generator

import (
	"fmt"
	"testing"

	"darvaza.org/core"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Compile-time verification that test case types implement TestCase interface
var _ core.TestCase = isMessageWithNameTestCase{}
var _ core.TestCase = boolCheckTestCase{}
var _ core.TestCase = mapFieldWithMessageTestCase{}

// Test case for IsMessageWithName which has a different signature
type isMessageWithNameTestCase struct {
	name     string
	desc     proto.Message
	typeName string
	expected bool
}

func newIsMessageWithNameTestCase(name string, desc proto.Message,
	typeName string, expected bool) isMessageWithNameTestCase {
	return isMessageWithNameTestCase{
		name:     name,
		desc:     desc,
		typeName: typeName,
		expected: expected,
	}
}

func (tc isMessageWithNameTestCase) Name() string {
	return tc.name
}

func (tc isMessageWithNameTestCase) Test(t *testing.T) {
	t.Helper()
	result := IsMessageWithName(tc.desc, tc.typeName)
	core.AssertEqual(t, tc.expected, result, "IsMessageWithName")
}

// Generic test case for boolean check functions that take a proto.Message
type boolCheckTestCase struct {
	// Larger fields first for field alignment
	checkFn  func(proto.Message) bool
	desc     proto.Message
	name     string
	fnName   string // For better error messages
	expected bool
}

// Base factory function with all parameters
func newBoolCheckTestCase(name string, desc proto.Message, expected bool,
	checkFn func(proto.Message) bool, fnName string) boolCheckTestCase {
	return boolCheckTestCase{
		name:     name,
		desc:     desc,
		expected: expected,
		checkFn:  checkFn,
		fnName:   fnName,
	}
}

func (tc boolCheckTestCase) Name() string {
	return tc.name
}

func (tc boolCheckTestCase) Test(t *testing.T) {
	t.Helper()
	result := tc.checkFn(tc.desc)
	core.AssertEqual(t, tc.expected, result, tc.fnName)
}

// Type-specific factory functions for better readability

// newIsMessageTestCase creates a test case for IsMessage function
func newIsMessageTestCase(name string, desc proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, expected, IsMessage, "IsMessage")
}

// newIsFieldTypeTestCase creates a test case for IsFieldType function
func newIsFieldTypeTestCase(name string, desc proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, expected, IsFieldType, "IsFieldType")
}

// newIsServiceTypeTestCase creates a test case for IsServiceType function
func newIsServiceTypeTestCase(name string, desc proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, expected, IsServiceType, "IsServiceType")
}

// newIsMethodTypeTestCase creates a test case for IsMethodType function
func newIsMethodTypeTestCase(name string, desc proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, expected, IsMethodType, "IsMethodType")
}

// newIsEnumTypeTestCase creates a test case for IsEnumType function
func newIsEnumTypeTestCase(name string, desc proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, expected, IsEnumType, "IsEnumType")
}

// newIsFileTypeTestCase creates a test case for IsFileType function
func newIsFileTypeTestCase(name string, desc proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, expected, IsFileType, "IsFileType")
}

// newIsRepeatedFieldTestCase creates a test case for IsRepeatedField function
func newIsRepeatedFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsRepeatedField, "IsRepeatedField")
}

// newIsMapFieldTestCase creates a test case for IsMapField function
func newIsMapFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsMapField, "IsMapField")
}

// newIsOneOfFieldTestCase creates a test case for IsOneOfField function
func newIsOneOfFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsOneOfField, "IsOneOfField")
}

// newIsOptionalFieldTestCase creates a test case for IsOptionalField function
func newIsOptionalFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsOptionalField, "IsOptionalField")
}

// newIsRequiredFieldTestCase creates a test case for IsRequiredField function
func newIsRequiredFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsRequiredField, "IsRequiredField")
}

// newIsScalarFieldTestCase creates a test case for IsScalarField function
func newIsScalarFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsScalarField, "IsScalarField")
}

// newIsMessageFieldTestCase creates a test case for IsMessageField function
func newIsMessageFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsMessageField, "IsMessageField")
}

// newIsEnumFieldTestCase creates a test case for IsEnumField function
func newIsEnumFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsEnumField, "IsEnumField")
}

// newIsGroupFieldTestCase creates a test case for IsGroupField function
func newIsGroupFieldTestCase(name string, field proto.Message, expected bool) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, expected, IsGroupField, "IsGroupField")
}

// Test functions

func TestIsMessage(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsMessageTestCase("message descriptor with name",
			&descriptorpb.DescriptorProto{Name: proto.String("TestMessage")}, true),
		newIsMessageTestCase("message descriptor without name",
			&descriptorpb.DescriptorProto{}, false),
		newIsMessageTestCase("field descriptor", &descriptorpb.FieldDescriptorProto{}, false),
		newIsMessageTestCase("service descriptor", &descriptorpb.ServiceDescriptorProto{}, false),
		newIsMessageTestCase("nil descriptor", nil, false),
	}

	core.RunTestCases(t, testCases)
}

func TestIsMessageWithName(t *testing.T) {
	msgDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("TestMessage"),
	}

	testCases := []isMessageWithNameTestCase{
		newIsMessageWithNameTestCase("matching message type", msgDesc, "TestMessage", true),
		newIsMessageWithNameTestCase("non-matching message type", msgDesc, "OtherMessage", false),
		newIsMessageWithNameTestCase("nil descriptor", nil, "TestMessage", false),
		newIsMessageWithNameTestCase("empty type name matches any", msgDesc, "", true),
		newIsMessageWithNameTestCase("wrong descriptor type",
			&descriptorpb.FieldDescriptorProto{}, "TestMessage", false),
	}

	core.RunTestCases(t, testCases)
}

func isFieldTypeTestCases() []boolCheckTestCase {
	fieldType := TypeString
	fieldWithType := &descriptorpb.FieldDescriptorProto{Type: &fieldType}

	return []boolCheckTestCase{
		newIsFieldTypeTestCase("field descriptor with type", fieldWithType, true),
		newIsFieldTypeTestCase("field descriptor without type", &descriptorpb.FieldDescriptorProto{}, false),
		newIsFieldTypeTestCase("message descriptor", &descriptorpb.DescriptorProto{}, false),
		newIsFieldTypeTestCase("service descriptor", &descriptorpb.ServiceDescriptorProto{}, false),
		newIsFieldTypeTestCase("nil descriptor", nil, false),
	}
}

func TestIsFieldType(t *testing.T) {
	core.RunTestCases(t, isFieldTypeTestCases())
}

func TestIsServiceType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsServiceTypeTestCase("service descriptor with name",
			&descriptorpb.ServiceDescriptorProto{Name: proto.String("TestService")}, true),
		newIsServiceTypeTestCase("service descriptor without name",
			&descriptorpb.ServiceDescriptorProto{}, false),
		newIsServiceTypeTestCase("message descriptor", &descriptorpb.DescriptorProto{}, false),
		newIsServiceTypeTestCase("field descriptor", &descriptorpb.FieldDescriptorProto{}, false),
		newIsServiceTypeTestCase("nil descriptor", nil, false),
	}

	core.RunTestCases(t, testCases)
}

func TestIsMethodType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsMethodTypeTestCase("method descriptor with all fields",
			&descriptorpb.MethodDescriptorProto{
				Name:       proto.String("TestMethod"),
				InputType:  proto.String(".TestRequest"),
				OutputType: proto.String(".TestResponse"),
			}, true),
		newIsMethodTypeTestCase("method descriptor without name",
			&descriptorpb.MethodDescriptorProto{
				InputType:  proto.String(".TestRequest"),
				OutputType: proto.String(".TestResponse"),
			}, false),
		newIsMethodTypeTestCase("method descriptor without input type",
			&descriptorpb.MethodDescriptorProto{
				Name:       proto.String("TestMethod"),
				OutputType: proto.String(".TestResponse"),
			}, false),
		newIsMethodTypeTestCase("method descriptor without output type",
			&descriptorpb.MethodDescriptorProto{
				Name:      proto.String("TestMethod"),
				InputType: proto.String(".TestRequest"),
			}, false),
		newIsMethodTypeTestCase("service descriptor", &descriptorpb.ServiceDescriptorProto{}, false),
		newIsMethodTypeTestCase("message descriptor", &descriptorpb.DescriptorProto{}, false),
		newIsMethodTypeTestCase("nil descriptor", nil, false),
	}

	core.RunTestCases(t, testCases)
}

func TestIsEnumType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsEnumTypeTestCase("enum descriptor with name",
			&descriptorpb.EnumDescriptorProto{Name: proto.String("TestEnum")}, true),
		newIsEnumTypeTestCase("enum descriptor without name",
			&descriptorpb.EnumDescriptorProto{}, false),
		newIsEnumTypeTestCase("enum value descriptor", &descriptorpb.EnumValueDescriptorProto{}, false),
		newIsEnumTypeTestCase("message descriptor", &descriptorpb.DescriptorProto{}, false),
		newIsEnumTypeTestCase("nil descriptor", nil, false),
	}

	core.RunTestCases(t, testCases)
}

func TestIsFileType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsFileTypeTestCase("file descriptor with name",
			&descriptorpb.FileDescriptorProto{Name: proto.String("test.proto")}, true),
		newIsFileTypeTestCase("file descriptor without name",
			&descriptorpb.FileDescriptorProto{}, false),
		newIsFileTypeTestCase("message descriptor", &descriptorpb.DescriptorProto{}, false),
		newIsFileTypeTestCase("service descriptor", &descriptorpb.ServiceDescriptorProto{}, false),
		newIsFileTypeTestCase("nil descriptor", nil, false),
	}

	core.RunTestCases(t, testCases)
}

func isRepeatedFieldTestCases() []boolCheckTestCase {
	repeatedField := NewFieldWithLabel(LabelRepeated)
	optionalField := NewFieldWithLabel(LabelOptional)

	return []boolCheckTestCase{
		newIsRepeatedFieldTestCase("repeated field", repeatedField, true),
		newIsRepeatedFieldTestCase("optional field", optionalField, false),
		newIsRepeatedFieldTestCase("nil field", nil, false),
		newIsRepeatedFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsRepeatedField(t *testing.T) {
	core.RunTestCases(t, isRepeatedFieldTestCases())
}

func isMapFieldTestCases() []boolCheckTestCase {
	// In protobuf, map fields are represented as repeated message fields
	// where the message type is a map entry (has map_entry option set to true)
	// For testing purposes, we'll check if it's a repeated message field
	// with a type name that contains "Entry" (simplified check)
	mapField := NewRepeatedMessageField(".MapEntry")

	// Another map field example
	mapFieldWithOptions := NewRepeatedMessageField(".MyMapEntry")
	mapFieldWithOptions.Options = &descriptorpb.FieldOptions{}

	regularRepeated := NewRepeatedField("regular", 1, TypeString)

	regularMessage := NewMessageField("message", 2, ".SomeMessage")

	// Edge cases for AsMapField coverage
	repeatedMessageNoTypeName := NewRepeatedMessageField("") // empty typename

	repeatedMessageNoEntry := NewRepeatedMessageField(".SomeMessage") // Doesn't contain "Entry"

	return []boolCheckTestCase{
		newIsMapFieldTestCase("map field", mapField, true),
		newIsMapFieldTestCase("map field with options", mapFieldWithOptions, true),
		newIsMapFieldTestCase("regular repeated field", regularRepeated, false),
		newIsMapFieldTestCase("regular message field", regularMessage, false),
		newIsMapFieldTestCase("repeated message no typename", repeatedMessageNoTypeName, false),
		newIsMapFieldTestCase("repeated message no entry in name", repeatedMessageNoEntry, false),
		newIsMapFieldTestCase("nil field", nil, false),
		newIsMapFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsMapField(t *testing.T) {
	core.RunTestCases(t, isMapFieldTestCases())
}

// Test case for AsMapFieldWithMessage which has a different signature
type mapFieldWithMessageTestCase struct {
	field    proto.Message
	entryMsg proto.Message
	name     string
	expected bool
}

func newMapFieldWithMessageTestCase(name string, field proto.Message,
	entryMsg proto.Message, expected bool) mapFieldWithMessageTestCase {
	return mapFieldWithMessageTestCase{
		name:     name,
		field:    field,
		entryMsg: entryMsg,
		expected: expected,
	}
}

func (tc mapFieldWithMessageTestCase) Name() string {
	return tc.name
}

func (tc mapFieldWithMessageTestCase) Test(t *testing.T) {
	t.Helper()

	// Test AsMapFieldWithMessage
	field, ok := AsMapFieldWithMessage(tc.field, tc.entryMsg)
	core.AssertEqual(t, tc.expected, ok, "AsMapFieldWithMessage")
	if tc.expected {
		core.AssertNotNil(t, field, "field")
	} else {
		core.AssertNil(t, field, "field")
	}

	// Test IsMapFieldWithMessage
	result := IsMapFieldWithMessage(tc.field, tc.entryMsg)
	core.AssertEqual(t, tc.expected, result, "IsMapFieldWithMessage")
}

func mapFieldWithMessageTestCases() []mapFieldWithMessageTestCase {
	// Create test data
	mapField := NewRepeatedMessageField(".MapEntry")
	regularField := NewField("regular", 1, TypeString)

	// Map entry message with map_entry=true
	mapEntryTrue := proto.Bool(true)
	mapEntryMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("MapEntry"),
		Options: &descriptorpb.MessageOptions{
			MapEntry: mapEntryTrue,
		},
	}

	// Regular message without map_entry option
	regularMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("RegularMessage"),
	}

	// Message with map_entry=false
	mapEntryFalse := proto.Bool(false)
	notMapEntryMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("NotMapEntry"),
		Options: &descriptorpb.MessageOptions{
			MapEntry: mapEntryFalse,
		},
	}

	return []mapFieldWithMessageTestCase{
		newMapFieldWithMessageTestCase("map field with map entry message",
			mapField, mapEntryMsg, true),
		newMapFieldWithMessageTestCase("map field with regular message",
			mapField, regularMsg, false),
		newMapFieldWithMessageTestCase("map field with map_entry=false",
			mapField, notMapEntryMsg, false),
		newMapFieldWithMessageTestCase("regular field with map entry message",
			regularField, mapEntryMsg, false),
		newMapFieldWithMessageTestCase("nil field",
			nil, mapEntryMsg, false),
		newMapFieldWithMessageTestCase("nil entry message fallback to heuristic",
			mapField, nil, true),
		newMapFieldWithMessageTestCase("wrong field type",
			&descriptorpb.DescriptorProto{}, mapEntryMsg, false),
		newMapFieldWithMessageTestCase("wrong entry message type",
			mapField, &descriptorpb.FieldDescriptorProto{}, false),
	}
}

func TestAsMapFieldWithMessage(t *testing.T) {
	core.RunTestCases(t, mapFieldWithMessageTestCases())
}

func isOneOfFieldTestCases() []boolCheckTestCase {
	oneOfField := NewOneOfField("variant", 1, TypeString, 0)
	regularField := NewField("regular", 2, TypeString)

	return []boolCheckTestCase{
		newIsOneOfFieldTestCase("OneOf field", oneOfField, true),
		newIsOneOfFieldTestCase("regular field", regularField, false),
		newIsOneOfFieldTestCase("nil field", nil, false),
		newIsOneOfFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsOneOfField(t *testing.T) {
	core.RunTestCases(t, isOneOfFieldTestCases())
}

func isOptionalFieldTestCases() []boolCheckTestCase {
	optionalField := NewField("optional", 1, TypeString)
	requiredField := NewRequiredField("required", 2, TypeString)

	// Proto3 optional needs special handling
	proto3Optional := NewField("proto3_optional", 3, TypeString)
	proto3Optional.Proto3Optional = proto.Bool(true)

	// Proto3 singular field (has LABEL_OPTIONAL but Proto3Optional=false)
	// This represents a regular proto3 field without the 'optional' keyword
	proto3Singular := NewField("proto3_singular", 4, TypeString)
	proto3Singular.Proto3Optional = proto.Bool(false)

	return []boolCheckTestCase{
		newIsOptionalFieldTestCase("optional field", optionalField, true),
		newIsOptionalFieldTestCase("proto3 optional field", proto3Optional, true),
		newIsOptionalFieldTestCase("proto3 singular field", proto3Singular, true),
		newIsOptionalFieldTestCase("required field", requiredField, false),
		newIsOptionalFieldTestCase("nil field", nil, false),
		newIsOptionalFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsOptionalField(t *testing.T) {
	core.RunTestCases(t, isOptionalFieldTestCases())
}

func isRequiredFieldTestCases() []boolCheckTestCase {
	requiredField := NewRequiredField("required", 1, TypeString)
	optionalField := NewField("optional", 2, TypeString)

	return []boolCheckTestCase{
		newIsRequiredFieldTestCase("required field", requiredField, true),
		newIsRequiredFieldTestCase("optional field", optionalField, false),
		newIsRequiredFieldTestCase("nil field", nil, false),
		newIsRequiredFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsRequiredField(t *testing.T) {
	core.RunTestCases(t, isRequiredFieldTestCases())
}

func isScalarFieldTestCases() []boolCheckTestCase {
	scalarFields := []proto.Message{
		NewField("double", 1, TypeDouble),
		NewField("float", 2, TypeFloat),
		NewField("int64", 3, TypeInt64),
		NewField("int32", 5, TypeInt32),
		NewField("fixed64", 6, TypeFixed64),
		NewField("fixed32", 7, TypeFixed32),
		NewField("bool", 8, TypeBool),
		NewField("string", 9, TypeString),
		NewField("bytes", 10, TypeBytes),
		NewField("u_int32", 11, TypeUInt32),
		NewField("u_int64", 4, TypeUInt64),
		NewField("s_fixed32", 12, TypeSFixed32),
		NewField("s_fixed64", 13, TypeSFixed64),
		NewField("s_int32", 14, TypeSInt32),
		NewField("s_int64", 15, TypeSInt64),
	}

	testCases := []boolCheckTestCase{}

	// Add positive test cases for all scalar types
	for i, field := range scalarFields {
		name := fmt.Sprintf("scalar field %d", i+1)
		testCases = append(testCases, newIsScalarFieldTestCase(name, field, true))
	}

	// Add negative test cases
	messageField := NewMessageField("message", 16, ".example.Message")
	enumField := NewEnumField("enum", 17, ".example.Enum")
	groupField := NewField("group", 18, TypeGroup)

	// With Label but no type
	labelOptional := LabelOptional
	fieldWithoutType := &descriptorpb.FieldDescriptorProto{Label: &labelOptional}

	testCases = append(testCases,
		newIsScalarFieldTestCase("message field", messageField, false),
		newIsScalarFieldTestCase("enum field", enumField, false),
		newIsScalarFieldTestCase("group field", groupField, false),
		newIsScalarFieldTestCase("field without type", fieldWithoutType, false),
		newIsScalarFieldTestCase("nil field", nil, false),
		newIsScalarFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	)

	return testCases
}

func TestIsScalarField(t *testing.T) {
	core.RunTestCases(t, isScalarFieldTestCases())
}

func isMessageFieldTestCases() []boolCheckTestCase {
	messageField := NewFieldWithType(TypeMessage)
	groupField := NewFieldWithType(TypeGroup)
	scalarField := NewFieldWithType(TypeString)
	enumField := NewFieldWithType(TypeEnum)

	// With label but no type
	labelOptional := LabelOptional
	fieldWithoutType := &descriptorpb.FieldDescriptorProto{Label: &labelOptional}

	return []boolCheckTestCase{
		newIsMessageFieldTestCase("message field", messageField, true),
		newIsMessageFieldTestCase("group field", groupField, false), // GROUP is separate from MESSAGE
		newIsMessageFieldTestCase("scalar field", scalarField, false),
		newIsMessageFieldTestCase("enum field", enumField, false),
		newIsMessageFieldTestCase("field without type", fieldWithoutType, false),
		newIsMessageFieldTestCase("nil field", nil, false),
		newIsMessageFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsMessageField(t *testing.T) {
	core.RunTestCases(t, isMessageFieldTestCases())
}

func isGroupFieldTestCases() []boolCheckTestCase {
	groupField := NewFieldWithType(TypeGroup)
	messageField := NewFieldWithType(TypeMessage)
	scalarField := NewFieldWithType(TypeString)
	enumField := NewFieldWithType(TypeEnum)

	labelOptional := LabelOptional
	fieldWithoutType := &descriptorpb.FieldDescriptorProto{Label: &labelOptional} // Has label but no type

	return []boolCheckTestCase{
		newIsGroupFieldTestCase("group field", groupField, true),
		newIsGroupFieldTestCase("message field", messageField, false),
		newIsGroupFieldTestCase("scalar field", scalarField, false),
		newIsGroupFieldTestCase("enum field", enumField, false),
		newIsGroupFieldTestCase("field without type", fieldWithoutType, false),
		newIsGroupFieldTestCase("nil field", nil, false),
		newIsGroupFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsGroupField(t *testing.T) {
	core.RunTestCases(t, isGroupFieldTestCases())
}

func isEnumFieldTestCases() []boolCheckTestCase {
	enumField := NewFieldWithType(TypeEnum)
	scalarField := NewFieldWithType(TypeInt32)
	messageField := NewFieldWithType(TypeMessage)

	return []boolCheckTestCase{
		newIsEnumFieldTestCase("enum field", enumField, true),
		newIsEnumFieldTestCase("scalar field", scalarField, false),
		newIsEnumFieldTestCase("message field", messageField, false),
		newIsEnumFieldTestCase("nil field", nil, false),
		newIsEnumFieldTestCase("wrong type", &descriptorpb.DescriptorProto{}, false),
	}
}

func TestIsEnumField(t *testing.T) {
	core.RunTestCases(t, isEnumFieldTestCases())
}
