package generator

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Test utility functions for creating protobuf descriptors
// These are only used in tests and provide convenient shortcuts

// NewField creates a field descriptor with the given name and type.
// Returns a FieldDescriptorProto with LABEL_OPTIONAL.
func NewField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	label := LabelOptional
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
		Label:  &label,
		Type:   &fieldType,
	}
}

// NewRepeatedField creates a repeated field descriptor.
// Returns a FieldDescriptorProto with LABEL_REPEATED.
func NewRepeatedField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	label := LabelRepeated
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
		Label:  &label,
		Type:   &fieldType,
	}
}

// NewRequiredField creates a required field descriptor (proto2).
// Returns a FieldDescriptorProto with LABEL_REQUIRED.
func NewRequiredField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	label := LabelRequired
	return &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
		Label:  &label,
		Type:   &fieldType,
	}
}

// NewMessageField creates a message type field descriptor.
// Returns a FieldDescriptorProto with TYPE_MESSAGE and the given type name.
func NewMessageField(name string, number int32, typeName string) *descriptorpb.FieldDescriptorProto {
	label := LabelOptional
	msgType := TypeMessage
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(number),
		Label:    &label,
		Type:     &msgType,
		TypeName: proto.String(typeName),
	}
}

// NewEnumField creates an enum type field descriptor.
// Returns a FieldDescriptorProto with TYPE_ENUM and the given type name.
func NewEnumField(name string, number int32, typeName string) *descriptorpb.FieldDescriptorProto {
	label := LabelOptional
	enumType := TypeEnum
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(number),
		Label:    &label,
		Type:     &enumType,
		TypeName: proto.String(typeName),
	}
}

// NewMapField creates a map field descriptor.
// Note: This creates a repeated message field with the given entry type name.
// The actual map entry message with map_entry option should be defined separately.
func NewMapField(name string, number int32, entryTypeName string) *descriptorpb.FieldDescriptorProto {
	label := LabelRepeated
	msgType := TypeMessage
	return &descriptorpb.FieldDescriptorProto{
		Name:     proto.String(name),
		Number:   proto.Int32(number),
		Label:    &label,
		Type:     &msgType,
		TypeName: proto.String(entryTypeName),
	}
}

// NewOneOfField creates a field that's part of a oneof.
// Returns a FieldDescriptorProto with the given oneof index.
func NewOneOfField(name string, number int32,
	fieldType descriptorpb.FieldDescriptorProto_Type,
	oneOfIndex int32) *descriptorpb.FieldDescriptorProto {
	label := LabelOptional
	return &descriptorpb.FieldDescriptorProto{
		Name:       proto.String(name),
		Number:     proto.Int32(number),
		Label:      &label,
		Type:       &fieldType,
		OneofIndex: proto.Int32(oneOfIndex),
	}
}

// NewMessage creates a message descriptor with the given name.
// Returns a DescriptorProto with the provided fields.
func NewMessage(name string, fields ...*descriptorpb.FieldDescriptorProto) *descriptorpb.DescriptorProto {
	return &descriptorpb.DescriptorProto{
		Name:  proto.String(name),
		Field: fields,
	}
}

// NewMessageWithNested creates a message with nested types.
// Returns a DescriptorProto with nested messages and enums.
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

// NewEnum creates an enum descriptor with the given name and values.
// Returns an EnumDescriptorProto with values numbered from 0.
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

// NewEnumValue creates an enum value descriptor.
// Returns an EnumValueDescriptorProto with the given name and number.
func NewEnumValue(name string, number int32) *descriptorpb.EnumValueDescriptorProto {
	return &descriptorpb.EnumValueDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(number),
	}
}

// NewService creates a service descriptor with the given name.
// Returns a ServiceDescriptorProto with the provided methods.
func NewService(name string, methods ...*descriptorpb.MethodDescriptorProto) *descriptorpb.ServiceDescriptorProto {
	return &descriptorpb.ServiceDescriptorProto{
		Name:   proto.String(name),
		Method: methods,
	}
}

// NewMethod creates a method descriptor.
// Returns a MethodDescriptorProto with the given input and output types.
func NewMethod(name, inputType, outputType string) *descriptorpb.MethodDescriptorProto {
	return &descriptorpb.MethodDescriptorProto{
		Name:       proto.String(name),
		InputType:  proto.String(inputType),
		OutputType: proto.String(outputType),
	}
}

// NewFile creates a file descriptor with the given name and package.
// Returns a FileDescriptorProto with basic metadata.
func NewFile(name, pkg string) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:    proto.String(name),
		Package: proto.String(pkg),
	}
}

// NewFileWithTypes creates a file descriptor with various types.
// Returns a FileDescriptorProto with messages, enums, and services.
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

// NewOneOf creates a oneof descriptor.
// Returns a OneofDescriptorProto with the given name.
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

// NewFieldWithLabel creates a minimal field descriptor with label and a default type
// Useful for testing cardinality functions. Uses TYPE_STRING as default type.
func NewFieldWithLabel(label descriptorpb.FieldDescriptorProto_Label) *descriptorpb.FieldDescriptorProto {
	defaultType := TypeString
	return &descriptorpb.FieldDescriptorProto{
		Label: &label,
		Type:  &defaultType,
	}
}

// NewRepeatedMessageField creates a repeated message field with optional type name
// Useful for map field testing and other repeated message scenarios
func NewRepeatedMessageField(typeName string) *descriptorpb.FieldDescriptorProto {
	label := LabelRepeated
	msgType := TypeMessage
	field := &descriptorpb.FieldDescriptorProto{
		Label: &label,
		Type:  &msgType,
	}
	if typeName != "" {
		field.TypeName = proto.String(typeName)
	}
	return field
}
