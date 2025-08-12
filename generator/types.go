package generator

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Field type aliases for common protobuf field types.
// These provide shorter, more readable names for the descriptorpb constants.
// cspell:ignore SFIXED SINT
const (
	// Numeric types
	TypeDouble   = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
	TypeFloat    = descriptorpb.FieldDescriptorProto_TYPE_FLOAT
	TypeInt64    = descriptorpb.FieldDescriptorProto_TYPE_INT64
	TypeUInt64   = descriptorpb.FieldDescriptorProto_TYPE_UINT64
	TypeInt32    = descriptorpb.FieldDescriptorProto_TYPE_INT32
	TypeFixed64  = descriptorpb.FieldDescriptorProto_TYPE_FIXED64
	TypeFixed32  = descriptorpb.FieldDescriptorProto_TYPE_FIXED32
	TypeSFixed32 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED32
	TypeSFixed64 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED64
	TypeSInt32   = descriptorpb.FieldDescriptorProto_TYPE_SINT32
	TypeSInt64   = descriptorpb.FieldDescriptorProto_TYPE_SINT64
	TypeUInt32   = descriptorpb.FieldDescriptorProto_TYPE_UINT32

	// Boolean and string types
	TypeBool   = descriptorpb.FieldDescriptorProto_TYPE_BOOL
	TypeString = descriptorpb.FieldDescriptorProto_TYPE_STRING
	TypeBytes  = descriptorpb.FieldDescriptorProto_TYPE_BYTES

	// Complex types
	TypeGroup   = descriptorpb.FieldDescriptorProto_TYPE_GROUP
	TypeMessage = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	TypeEnum    = descriptorpb.FieldDescriptorProto_TYPE_ENUM
)

// Label type aliases for field cardinality
const (
	LabelOptional = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	LabelRequired = descriptorpb.FieldDescriptorProto_LABEL_REQUIRED
	LabelRepeated = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
)
