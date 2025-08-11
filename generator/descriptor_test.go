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

func isRepeatedFieldTestCases() []boolCheckTestCase {
	repeatedField := NewFieldWithLabel(descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
	optionalField := NewFieldWithLabel(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL)

	return []boolCheckTestCase{
		newIsRepeatedFieldTrue("repeated field", repeatedField),
		newIsRepeatedFieldFalse("optional field", optionalField),
		newIsRepeatedFieldFalse("nil field", nil),
		newIsRepeatedFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
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
		newIsMapFieldTrue("map field", mapField),
		newIsMapFieldTrue("map field with options", mapFieldWithOptions),
		newIsMapFieldFalse("regular repeated field", regularRepeated),
		newIsMapFieldFalse("regular message field", regularMessage),
		newIsMapFieldFalse("repeated message no typename", repeatedMessageNoTypeName),
		newIsMapFieldFalse("repeated message no entry in name", repeatedMessageNoEntry),
		newIsMapFieldFalse("nil field", nil),
		newIsMapFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}
}

func TestIsMapField(t *testing.T) {
	core.RunTestCases(t, isMapFieldTestCases())
}

func isOneOfFieldTestCases() []boolCheckTestCase {
	oneOfField := NewOneOfField("variant", 1, TypeString, 0)
	regularField := NewField("regular", 2, TypeString)

	return []boolCheckTestCase{
		newIsOneOfFieldTrue("OneOf field", oneOfField),
		newIsOneOfFieldFalse("regular field", regularField),
		newIsOneOfFieldFalse("nil field", nil),
		newIsOneOfFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
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

	return []boolCheckTestCase{
		newIsOptionalFieldTrue("optional field", optionalField),
		newIsOptionalFieldTrue("proto3 optional field", proto3Optional),
		newIsOptionalFieldFalse("required field", requiredField),
		newIsOptionalFieldFalse("nil field", nil),
		newIsOptionalFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}
}

func TestIsOptionalField(t *testing.T) {
	core.RunTestCases(t, isOptionalFieldTestCases())
}

func isRequiredFieldTestCases() []boolCheckTestCase {
	requiredField := NewRequiredField("required", 1, TypeString)
	optionalField := NewField("optional", 2, TypeString)

	return []boolCheckTestCase{
		newIsRequiredFieldTrue("required field", requiredField),
		newIsRequiredFieldFalse("optional field", optionalField),
		newIsRequiredFieldFalse("nil field", nil),
		newIsRequiredFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
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
		NewField("uint64", 4, TypeUInt64),
		NewField("int32", 5, TypeInt32),
		NewField("fixed64", 6, TypeFixed64),
		NewField("fixed32", 7, TypeFixed32),
		NewField("bool", 8, TypeBool),
		NewField("string", 9, TypeString),
		NewField("bytes", 10, TypeBytes),
		NewField("uint32", 11, TypeUInt32),
		NewField("sfixed32", 12, TypeSFixed32),
		NewField("sfixed64", 13, TypeSFixed64),
		NewField("sint32", 14, TypeSInt32),
		NewField("sint64", 15, TypeSInt64),
	}

	testCases := []boolCheckTestCase{}

	// Add positive test cases for all scalar types
	for i, field := range scalarFields {
		testCases = append(testCases, newIsScalarFieldTrue("scalar field "+string(rune('A'+i)), field))
	}

	// Add negative test cases
	messageField := NewMessageField("message", 16, ".example.Message")
	enumField := NewEnumField("enum", 17, ".example.Enum")
	groupField := NewField("group", 18, TypeGroup)

	testCases = append(testCases,
		newIsScalarFieldFalse("message field", messageField),
		newIsScalarFieldFalse("enum field", enumField),
		newIsScalarFieldFalse("group field", groupField),
		newIsScalarFieldFalse("nil field", nil),
		newIsScalarFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
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

	return []boolCheckTestCase{
		newIsMessageFieldTrue("message field", messageField),
		newIsMessageFieldTrue("group field", groupField), // GROUP is also a message type
		newIsMessageFieldFalse("scalar field", scalarField),
		newIsMessageFieldFalse("enum field", enumField),
		newIsMessageFieldFalse("nil field", nil),
		newIsMessageFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}
}

func TestIsMessageField(t *testing.T) {
	core.RunTestCases(t, isMessageFieldTestCases())
}

func isEnumFieldTestCases() []boolCheckTestCase {
	enumField := NewFieldWithType(TypeEnum)
	scalarField := NewFieldWithType(TypeInt32)
	messageField := NewFieldWithType(TypeMessage)

	return []boolCheckTestCase{
		newIsEnumFieldTrue("enum field", enumField),
		newIsEnumFieldFalse("scalar field", scalarField),
		newIsEnumFieldFalse("message field", messageField),
		newIsEnumFieldFalse("nil field", nil),
		newIsEnumFieldFalse("wrong type", &descriptorpb.DescriptorProto{}),
	}
}

func TestIsEnumField(t *testing.T) {
	core.RunTestCases(t, isEnumFieldTestCases())
}
