package generator

import (
	"testing"

	"darvaza.org/core"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestTypeConstants verifies that our type aliases match the protobuf constants
func TestTypeConstants(t *testing.T) {
	// cspell:ignore SFIXED SINT
	// Test scalar types
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_DOUBLE, TypeDouble, "TypeDouble")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_FLOAT, TypeFloat, "TypeFloat")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_INT64, TypeInt64, "TypeInt64")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_UINT64, TypeUInt64, "TypeUInt64")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_INT32, TypeInt32, "TypeInt32")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_FIXED64, TypeFixed64, "TypeFixed64")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_FIXED32, TypeFixed32, "TypeFixed32")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_BOOL, TypeBool, "TypeBool")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_STRING, TypeString, "TypeString")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_BYTES, TypeBytes, "TypeBytes")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_UINT32, TypeUInt32, "TypeUInt32")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_SFIXED32, TypeSFixed32, "TypeSFixed32")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64, TypeSFixed64, "TypeSFixed64")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_SINT32, TypeSInt32, "TypeSInt32")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_SINT64, TypeSInt64, "TypeSInt64")

	// Test complex types
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_GROUP, TypeGroup, "TypeGroup")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, TypeMessage, "TypeMessage")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_ENUM, TypeEnum, "TypeEnum")
}

// TestLabelConstants verifies that our label aliases match the protobuf constants
func TestLabelConstants(t *testing.T) {
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL, LabelOptional, "LabelOptional")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REQUIRED, LabelRequired, "LabelRequired")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, LabelRepeated, "LabelRepeated")
}
