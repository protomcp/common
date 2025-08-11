package generator

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Test utility functions for creating protobuf descriptors
// These are only used in tests and provide convenient shortcuts

// NewField creates a field descriptor with the given name and type
func NewField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
		Label:  &label,
		Type:   &fieldType,
	}
}

// NewRepeatedField creates a repeated field descriptor
func NewRepeatedField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_REPEATED
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
		Label:  &label,
		Type:   &fieldType,
	}
}

// NewRequiredField creates a required field descriptor (proto2)
func NewRequiredField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_REQUIRED
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
		Label:  &label,
		Type:   &fieldType,
	}
}

// NewMessageField creates a message type field descriptor
func NewMessageField(name string, number int32, typeName string) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(number),
		Label:    &label,
		Type:     &msgType,
		TypeName: proto.String(typeName),
	}
}

// NewEnumField creates an enum type field descriptor
func NewEnumField(name string, number int32, typeName string) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	enumType := descriptorpb.FieldDescriptorProto_TYPE_ENUM
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(number),
		Label:    &label,
		Type:     &enumType,
		TypeName: proto.String(typeName),
	}
}

// NewMapField creates a map field descriptor
func NewMapField(name string, number int32, entryTypeName string) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_REPEATED
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(number),
		Label:    &label,
		Type:     &msgType,
		TypeName: proto.String(entryTypeName),
	}
}

// NewOneOfField creates a field that's part of a oneof
func NewOneOfField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type,
	oneOfIndex int32) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	return &descriptorpb.FieldDescriptorProto{
		Name:       proto.String(name),
		Number:     proto.Int32(number),
		Label:      &label,
		Type:       &fieldType,
		OneofIndex: proto.Int32(oneOfIndex),
	}
}

// NewMessage creates a message descriptor with the given name
func NewMessage(name string, fields ...*descriptorpb.FieldDescriptorProto) *descriptorpb.DescriptorProto {
	return &descriptorpb.DescriptorProto{
		Name:  proto.String(name),
		Field: fields,
	}
}

// NewMessageWithNested creates a message with nested types
func NewMessageWithNested(name string, fields []*descriptorpb.FieldDescriptorProto,
	nestedMessages []*descriptorpb.DescriptorProto,
	nestedEnums []*descriptorpb.EnumDescriptorProto) *descriptorpb.DescriptorProto {
	return &descriptorpb.DescriptorProto{
		Name:       proto.String(name),
		Field:      fields,
		NestedType: nestedMessages,
		EnumType:   nestedEnums,
	}
}

// NewEnum creates an enum descriptor with the given name and values
func NewEnum(name string, values ...string) *descriptorpb.EnumDescriptorProto {
	enumValues := make([]*descriptorpb.EnumValueDescriptorProto, len(values))
	for i, v := range values {
		enumValues[i] = &descriptorpb.EnumValueDescriptorProto{
			Name:   proto.String(v),
			Number: proto.Int32(int32(i)),
		}
	}
	return &descriptorpb.EnumDescriptorProto{
		Name:  proto.String(name),
		Value: enumValues,
	}
}

// NewEnumValue creates an enum value descriptor
func NewEnumValue(name string, number int32) *descriptorpb.EnumValueDescriptorProto {
	return &descriptorpb.EnumValueDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
	}
}

// NewService creates a service descriptor with the given name
func NewService(name string, methods ...*descriptorpb.MethodDescriptorProto) *descriptorpb.ServiceDescriptorProto {
	return &descriptorpb.ServiceDescriptorProto{
		Name:   proto.String(name),
		Method: methods,
	}
}

// NewMethod creates a method descriptor
func NewMethod(name, inputType, outputType string) *descriptorpb.MethodDescriptorProto {
	return &descriptorpb.MethodDescriptorProto{
		Name:       proto.String(name),
		InputType:  proto.String(inputType),
		OutputType: proto.String(outputType),
	}
}

// NewFile creates a file descriptor with the given name and package
func NewFile(name, pkg string) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:    proto.String(name),
		Package: proto.String(pkg),
	}
}

// NewFileWithTypes creates a file descriptor with various types
func NewFileWithTypes(name, pkg string,
	messages []*descriptorpb.DescriptorProto,
	enums []*descriptorpb.EnumDescriptorProto,
	services []*descriptorpb.ServiceDescriptorProto) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:        proto.String(name),
		Package:     proto.String(pkg),
		MessageType: messages,
		EnumType:    enums,
		Service:     services,
	}
}

// NewOneOf creates a oneof descriptor
func NewOneOf(name string) *descriptorpb.OneofDescriptorProto {
	return &descriptorpb.OneofDescriptorProto{
		Name: proto.String(name),
	}
}

// NewFieldWithType creates a minimal field descriptor with only type set
// Useful for testing type-checking functions
func NewFieldWithType(fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	return &descriptorpb.FieldDescriptorProto{
		Type: &fieldType,
	}
}

// NewFieldWithLabel creates a minimal field descriptor with only label set
// Useful for testing cardinality functions
func NewFieldWithLabel(label descriptorpb.FieldDescriptorProto_Label) *descriptorpb.FieldDescriptorProto {
	return &descriptorpb.FieldDescriptorProto{
		Label: &label,
	}
}

// NewRepeatedMessageField creates a repeated message field with optional type name
// Useful for map field testing and other repeated message scenarios
func NewRepeatedMessageField(typeName string) *descriptorpb.FieldDescriptorProto {
	label := descriptorpb.FieldDescriptorProto_LABEL_REPEATED
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	field := &descriptorpb.FieldDescriptorProto{
		Label: &label,
		Type:  &msgType,
	}
	if typeName != "" {
		field.TypeName = proto.String(typeName)
	}
	return field
}

// Field type shortcuts for common types
var (
	TypeDouble   = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
	TypeFloat    = descriptorpb.FieldDescriptorProto_TYPE_FLOAT
	TypeInt64    = descriptorpb.FieldDescriptorProto_TYPE_INT64
	TypeUInt64   = descriptorpb.FieldDescriptorProto_TYPE_UINT64
	TypeInt32    = descriptorpb.FieldDescriptorProto_TYPE_INT32
	TypeFixed64  = descriptorpb.FieldDescriptorProto_TYPE_FIXED64
	TypeFixed32  = descriptorpb.FieldDescriptorProto_TYPE_FIXED32
	TypeBool     = descriptorpb.FieldDescriptorProto_TYPE_BOOL
	TypeString   = descriptorpb.FieldDescriptorProto_TYPE_STRING
	TypeGroup    = descriptorpb.FieldDescriptorProto_TYPE_GROUP
	TypeMessage  = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	TypeBytes    = descriptorpb.FieldDescriptorProto_TYPE_BYTES
	TypeUInt32   = descriptorpb.FieldDescriptorProto_TYPE_UINT32
	TypeEnum     = descriptorpb.FieldDescriptorProto_TYPE_ENUM
	TypeSFixed32 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED32
	TypeSFixed64 = descriptorpb.FieldDescriptorProto_TYPE_SFIXED64
	TypeSInt32   = descriptorpb.FieldDescriptorProto_TYPE_SINT32
	TypeSInt64   = descriptorpb.FieldDescriptorProto_TYPE_SINT64
)
