package generator

import (
	"testing"

	"darvaza.org/core"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Compile-time verification that test case types implement TestCase interface
var _ core.TestCase = isMessageTypeTestCase{}
var _ core.TestCase = boolCheckTestCase{}

// Test case for IsMessageType which has a different signature
type isMessageTypeTestCase struct {
	name     string
	desc     proto.Message
	typeName string
	expected bool
}

func newIsMessageTypeTestCase(name string, desc proto.Message, typeName string, expected bool) isMessageTypeTestCase {
	return isMessageTypeTestCase{
		name:     name,
		desc:     desc,
		typeName: typeName,
		expected: expected,
	}
}

func (tc isMessageTypeTestCase) Name() string {
	return tc.name
}

func (tc isMessageTypeTestCase) Test(t *testing.T) {
	t.Helper()
	result := IsMessageType(tc.desc, tc.typeName)
	core.AssertEqual(t, tc.expected, result, "IsMessageType")
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

// Semantic factory variants for different test functions
// These reduce boolean parameter confusion and make test intent clearer

// IsMessage factories
func newIsMessageTrue(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, true, IsMessage, "IsMessage")
}

func newIsMessageFalse(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, false, IsMessage, "IsMessage")
}

// IsFieldType factories
func newIsFieldTypeTrue(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, true, IsFieldType, "IsFieldType")
}

func newIsFieldTypeFalse(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, false, IsFieldType, "IsFieldType")
}

// IsServiceType factories
func newIsServiceTypeTrue(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, true, IsServiceType, "IsServiceType")
}

func newIsServiceTypeFalse(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, false, IsServiceType, "IsServiceType")
}

// IsMethodType factories
func newIsMethodTypeTrue(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, true, IsMethodType, "IsMethodType")
}

func newIsMethodTypeFalse(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, false, IsMethodType, "IsMethodType")
}

// IsEnumType factories
func newIsEnumTypeTrue(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, true, IsEnumType, "IsEnumType")
}

func newIsEnumTypeFalse(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, false, IsEnumType, "IsEnumType")
}

// IsFileType factories
func newIsFileTypeTrue(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, true, IsFileType, "IsFileType")
}

func newIsFileTypeFalse(name string, desc proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, desc, false, IsFileType, "IsFileType")
}

// IsRepeatedField factories
func newIsRepeatedFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsRepeatedField, "IsRepeatedField")
}

func newIsRepeatedFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsRepeatedField, "IsRepeatedField")
}

// IsMapField factories
func newIsMapFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsMapField, "IsMapField")
}

func newIsMapFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsMapField, "IsMapField")
}

// IsOneOfField factories
func newIsOneOfFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsOneOfField, "IsOneOfField")
}

func newIsOneOfFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsOneOfField, "IsOneOfField")
}

// IsOptionalField factories
func newIsOptionalFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsOptionalField, "IsOptionalField")
}

func newIsOptionalFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsOptionalField, "IsOptionalField")
}

// IsRequiredField factories
func newIsRequiredFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsRequiredField, "IsRequiredField")
}

func newIsRequiredFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsRequiredField, "IsRequiredField")
}

// IsScalarField factories
func newIsScalarFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsScalarField, "IsScalarField")
}

func newIsScalarFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsScalarField, "IsScalarField")
}

// IsMessageField factories
func newIsMessageFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsMessageField, "IsMessageField")
}

func newIsMessageFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsMessageField, "IsMessageField")
}

// IsEnumField factories
func newIsEnumFieldTrue(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, true, IsEnumField, "IsEnumField")
}

func newIsEnumFieldFalse(name string, field proto.Message) boolCheckTestCase {
	return newBoolCheckTestCase(name, field, false, IsEnumField, "IsEnumField")
}

// Test functions

func TestIsMessage(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsMessageTrue("message descriptor", &descriptorpb.DescriptorProto{}),
		newIsMessageFalse("field descriptor", &descriptorpb.FieldDescriptorProto{}),
		newIsMessageFalse("service descriptor", &descriptorpb.ServiceDescriptorProto{}),
		newIsMessageFalse("nil descriptor", nil),
	}

	core.RunTestCases(t, testCases)
}

func TestIsMessageType(t *testing.T) {
	msgDesc := &descriptorpb.DescriptorProto{
		Name: proto.String("TestMessage"),
	}

	testCases := []isMessageTypeTestCase{
		newIsMessageTypeTestCase("matching message type", msgDesc, "TestMessage", true),
		newIsMessageTypeTestCase("non-matching message type", msgDesc, "OtherMessage", false),
		newIsMessageTypeTestCase("nil descriptor", nil, "TestMessage", false),
		newIsMessageTypeTestCase("empty type name", msgDesc, "", false),
		newIsMessageTypeTestCase("wrong descriptor type", &descriptorpb.FieldDescriptorProto{}, "TestMessage", false),
	}

	core.RunTestCases(t, testCases)
}

func TestIsFieldType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsFieldTypeTrue("field descriptor", &descriptorpb.FieldDescriptorProto{}),
		newIsFieldTypeFalse("message descriptor", &descriptorpb.DescriptorProto{}),
		newIsFieldTypeFalse("service descriptor", &descriptorpb.ServiceDescriptorProto{}),
		newIsFieldTypeFalse("nil descriptor", nil),
	}

	core.RunTestCases(t, testCases)
}

func TestIsServiceType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsServiceTypeTrue("service descriptor", &descriptorpb.ServiceDescriptorProto{}),
		newIsServiceTypeFalse("message descriptor", &descriptorpb.DescriptorProto{}),
		newIsServiceTypeFalse("field descriptor", &descriptorpb.FieldDescriptorProto{}),
		newIsServiceTypeFalse("nil descriptor", nil),
	}

	core.RunTestCases(t, testCases)
}

func TestIsMethodType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsMethodTypeTrue("method descriptor", &descriptorpb.MethodDescriptorProto{}),
		newIsMethodTypeFalse("service descriptor", &descriptorpb.ServiceDescriptorProto{}),
		newIsMethodTypeFalse("message descriptor", &descriptorpb.DescriptorProto{}),
		newIsMethodTypeFalse("nil descriptor", nil),
	}

	core.RunTestCases(t, testCases)
}

func TestIsEnumType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsEnumTypeTrue("enum descriptor", &descriptorpb.EnumDescriptorProto{}),
		newIsEnumTypeFalse("enum value descriptor", &descriptorpb.EnumValueDescriptorProto{}),
		newIsEnumTypeFalse("message descriptor", &descriptorpb.DescriptorProto{}),
		newIsEnumTypeFalse("nil descriptor", nil),
	}

	core.RunTestCases(t, testCases)
}

func TestIsFileType(t *testing.T) {
	testCases := []boolCheckTestCase{
		newIsFileTypeTrue("file descriptor", &descriptorpb.FileDescriptorProto{}),
		newIsFileTypeFalse("message descriptor", &descriptorpb.DescriptorProto{}),
		newIsFileTypeFalse("service descriptor", &descriptorpb.ServiceDescriptorProto{}),
		newIsFileTypeFalse("nil descriptor", nil),
	}

	core.RunTestCases(t, testCases)
}

func TestIsRepeatedField(t *testing.T) {
	repeatedField := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
	}

	optionalField := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
	}

	testCases := []boolCheckTestCase{
		newIsRepeatedFieldTrue("repeated field", repeatedField),
		newIsRepeatedFieldFalse("optional field", optionalField),
		newIsRepeatedFieldFalse("nil field", nil),
		newIsRepeatedFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}

func TestIsMapField(t *testing.T) {
	// In protobuf, map fields are represented as repeated message fields
	// where the message type is a map entry (has map_entry option set to true)
	// For testing purposes, we'll check if it's a repeated message field
	// with a type name that contains "Entry" (simplified check)
	mapField := &descriptorpb.FieldDescriptorProto{
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String(".MapEntry"),
	}

	// Another map field example
	mapFieldWithOptions := &descriptorpb.FieldDescriptorProto{
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String(".MyMapEntry"),
		Options:  &descriptorpb.FieldOptions{},
	}

	regularRepeated := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		Type:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
	}

	regularMessage := &descriptorpb.FieldDescriptorProto{
		Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String(".SomeMessage"),
	}

	testCases := []boolCheckTestCase{
		newIsMapFieldTrue("map field", mapField),
		newIsMapFieldTrue("map field with options", mapFieldWithOptions),
		newIsMapFieldFalse("regular repeated field", regularRepeated),
		newIsMapFieldFalse("regular message field", regularMessage),
		newIsMapFieldFalse("nil field", nil),
		newIsMapFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}

func TestIsOneOfField(t *testing.T) {
	oneOfIndex := int32(0)
	oneOfField := &descriptorpb.FieldDescriptorProto{
		OneofIndex: &oneOfIndex,
	}

	regularField := &descriptorpb.FieldDescriptorProto{}

	testCases := []boolCheckTestCase{
		newIsOneOfFieldTrue("OneOf field", oneOfField),
		newIsOneOfFieldFalse("regular field", regularField),
		newIsOneOfFieldFalse("nil field", nil),
		newIsOneOfFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}

func TestIsOptionalField(t *testing.T) {
	optionalField := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
	}

	requiredField := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum(),
	}

	proto3Optional := &descriptorpb.FieldDescriptorProto{
		Label:          descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Proto3Optional: proto.Bool(true),
	}

	testCases := []boolCheckTestCase{
		newIsOptionalFieldTrue("optional field", optionalField),
		newIsOptionalFieldTrue("proto3 optional field", proto3Optional),
		newIsOptionalFieldFalse("required field", requiredField),
		newIsOptionalFieldFalse("nil field", nil),
		newIsOptionalFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}

func TestIsRequiredField(t *testing.T) {
	requiredField := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum(),
	}

	optionalField := &descriptorpb.FieldDescriptorProto{
		Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
	}

	testCases := []boolCheckTestCase{
		newIsRequiredFieldTrue("required field", requiredField),
		newIsRequiredFieldFalse("optional field", optionalField),
		newIsRequiredFieldFalse("nil field", nil),
		newIsRequiredFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}

func TestIsScalarField(t *testing.T) {
	scalarFields := []proto.Message{
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_FLOAT.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_FIXED64.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_FIXED32.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_BYTES.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_SFIXED32.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_SFIXED64.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_SINT32.Enum()},
		&descriptorpb.FieldDescriptorProto{Type: descriptorpb.FieldDescriptorProto_TYPE_SINT64.Enum()},
	}

	testCases := []boolCheckTestCase{}

	// Add positive test cases for all scalar types
	for i, field := range scalarFields {
		testCases = append(testCases, newIsScalarFieldTrue("scalar field "+string(rune('A'+i)), field))
	}

	// Add negative test cases
	messageField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
	}
	enumField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
	}
	groupField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_GROUP.Enum(),
	}

	testCases = append(testCases,
		newIsScalarFieldFalse("message field", messageField),
		newIsScalarFieldFalse("enum field", enumField),
		newIsScalarFieldFalse("group field", groupField),
		newIsScalarFieldFalse("nil field", nil),
		newIsScalarFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	)

	core.RunTestCases(t, testCases)
}

func TestIsMessageField(t *testing.T) {
	messageField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
	}

	groupField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_GROUP.Enum(),
	}

	scalarField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
	}

	enumField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
	}

	testCases := []boolCheckTestCase{
		newIsMessageFieldTrue("message field", messageField),
		newIsMessageFieldTrue("group field", groupField), // GROUP is also a message type
		newIsMessageFieldFalse("scalar field", scalarField),
		newIsMessageFieldFalse("enum field", enumField),
		newIsMessageFieldFalse("nil field", nil),
		newIsMessageFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}

func TestIsEnumField(t *testing.T) {
	enumField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
	}

	scalarField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
	}

	messageField := &descriptorpb.FieldDescriptorProto{
		Type: descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
	}

	testCases := []boolCheckTestCase{
		newIsEnumFieldTrue("enum field", enumField),
		newIsEnumFieldFalse("scalar field", scalarField),
		newIsEnumFieldFalse("message field", messageField),
		newIsEnumFieldFalse("nil field", nil),
		newIsEnumFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}

	core.RunTestCases(t, testCases)
}
