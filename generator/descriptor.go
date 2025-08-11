package generator

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// AsMessageType attempts to cast the descriptor to a message descriptor and checks if it has the given name
// Returns the message descriptor and true if successful and name matches, nil and false otherwise
func AsMessageType(desc proto.Message, typeName string) (*descriptorpb.DescriptorProto, bool) {
	if desc == nil || typeName == "" {
		return nil, false
	}

	msgDesc, ok := desc.(*descriptorpb.DescriptorProto)
	if !ok {
		return nil, false
	}

	if msgDesc.Name == nil || *msgDesc.Name != typeName {
		return nil, false
	}

	return msgDesc, true
}

// IsMessageType checks if the descriptor is a message type with the given name
func IsMessageType(desc proto.Message, typeName string) bool {
	_, ok := AsMessageType(desc, typeName)
	return ok
}

// AsFieldType attempts to cast the descriptor to a field descriptor
// Returns the field descriptor and true if successful, nil and false otherwise
func AsFieldType(desc proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	if desc == nil {
		return nil, false
	}
	fieldDesc, ok := desc.(*descriptorpb.FieldDescriptorProto)
	return fieldDesc, ok
}

// IsFieldType checks if the descriptor is a field descriptor
func IsFieldType(desc proto.Message) bool {
	_, ok := AsFieldType(desc)
	return ok
}

// AsServiceType attempts to cast the descriptor to a service descriptor
// Returns the service descriptor and true if successful, nil and false otherwise
func AsServiceType(desc proto.Message) (*descriptorpb.ServiceDescriptorProto, bool) {
	if desc == nil {
		return nil, false
	}
	svcDesc, ok := desc.(*descriptorpb.ServiceDescriptorProto)
	return svcDesc, ok
}

// IsServiceType checks if the descriptor is a service descriptor
func IsServiceType(desc proto.Message) bool {
	_, ok := AsServiceType(desc)
	return ok
}

// AsMethodType attempts to cast the descriptor to a method descriptor
// Returns the method descriptor and true if successful, nil and false otherwise
func AsMethodType(desc proto.Message) (*descriptorpb.MethodDescriptorProto, bool) {
	if desc == nil {
		return nil, false
	}
	methodDesc, ok := desc.(*descriptorpb.MethodDescriptorProto)
	return methodDesc, ok
}

// IsMethodType checks if the descriptor is a method descriptor
func IsMethodType(desc proto.Message) bool {
	_, ok := AsMethodType(desc)
	return ok
}

// AsEnumType attempts to cast the descriptor to an enum descriptor
// Returns the enum descriptor and true if successful, nil and false otherwise
func AsEnumType(desc proto.Message) (*descriptorpb.EnumDescriptorProto, bool) {
	if desc == nil {
		return nil, false
	}
	enumDesc, ok := desc.(*descriptorpb.EnumDescriptorProto)
	return enumDesc, ok
}

// IsEnumType checks if the descriptor is an enum descriptor
func IsEnumType(desc proto.Message) bool {
	_, ok := AsEnumType(desc)
	return ok
}

// AsFileType attempts to cast the descriptor to a file descriptor
// Returns the file descriptor and true if successful, nil and false otherwise
func AsFileType(desc proto.Message) (*descriptorpb.FileDescriptorProto, bool) {
	if desc == nil {
		return nil, false
	}
	fileDesc, ok := desc.(*descriptorpb.FileDescriptorProto)
	return fileDesc, ok
}

// IsFileType checks if the descriptor is a file descriptor
func IsFileType(desc proto.Message) bool {
	_, ok := AsFileType(desc)
	return ok
}

// AsRepeatedField checks if the field is repeated and returns it as a field descriptor
// Returns the field descriptor and true if it's a repeated field, nil and false otherwise
func AsRepeatedField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok || fieldDesc.Label == nil {
		return nil, false
	}

	if *fieldDesc.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		return fieldDesc, true
	}
	return nil, false
}

// IsRepeatedField checks if the field is repeated
func IsRepeatedField(field proto.Message) bool {
	_, ok := AsRepeatedField(field)
	return ok
}

// AsMapField checks if the field is a map field and returns it as a field descriptor
// In protobuf, map fields are represented as repeated message fields
// where the message type is a special map entry type
// Returns the field descriptor and true if it's a map field, nil and false otherwise
func AsMapField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok {
		return nil, false
	}

	// Must be repeated
	if fieldDesc.Label == nil || *fieldDesc.Label != descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		return nil, false
	}

	// Must be a message type
	if fieldDesc.Type == nil || *fieldDesc.Type != descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		return nil, false
	}

	// Check if the type name indicates a map entry
	// In real protobuf, map entries have a special option, but for this simplified
	// implementation, we check if the type name contains "Entry"
	if fieldDesc.TypeName == nil {
		return nil, false
	}

	typeName := *fieldDesc.TypeName
	if strings.Contains(typeName, "Entry") {
		return fieldDesc, true
	}
	return nil, false
}

// IsMapField checks if the field is a map field
func IsMapField(field proto.Message) bool {
	_, ok := AsMapField(field)
	return ok
}

// AsOneOfField checks if the field is part of a oneof and returns it as a field descriptor
// Returns the field descriptor and true if it's part of a oneof, nil and false otherwise
func AsOneOfField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok {
		return nil, false
	}

	// A field is part of a oneof if it has a OneofIndex set
	if fieldDesc.OneofIndex != nil {
		return fieldDesc, true
	}
	return nil, false
}

// IsOneOfField checks if the field is part of a oneof
func IsOneOfField(field proto.Message) bool {
	_, ok := AsOneOfField(field)
	return ok
}

// AsOptionalField checks if the field is optional and returns it as a field descriptor
// Returns the field descriptor and true if it's optional, nil and false otherwise
func AsOptionalField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok || fieldDesc.Label == nil {
		return nil, false
	}

	if *fieldDesc.Label == descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL {
		return fieldDesc, true
	}
	return nil, false
}

// IsOptionalField checks if the field is optional
func IsOptionalField(field proto.Message) bool {
	_, ok := AsOptionalField(field)
	return ok
}

// AsRequiredField checks if the field is required and returns it as a field descriptor
// Returns the field descriptor and true if it's required, nil and false otherwise (proto2 only)
func AsRequiredField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok || fieldDesc.Label == nil {
		return nil, false
	}

	if *fieldDesc.Label == descriptorpb.FieldDescriptorProto_LABEL_REQUIRED {
		return fieldDesc, true
	}
	return nil, false
}

// IsRequiredField checks if the field is required (proto2 only)
func IsRequiredField(field proto.Message) bool {
	_, ok := AsRequiredField(field)
	return ok
}

// AsScalarField checks if the field is a scalar type and returns it as a field descriptor
// Returns the field descriptor and true if it's a scalar field, nil and false otherwise
func AsScalarField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok || fieldDesc.Type == nil {
		return nil, false
	}

	switch *fieldDesc.Type {
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE,
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT,
		descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
		descriptorpb.FieldDescriptorProto_TYPE_BOOL,
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_BYTES,
		descriptorpb.FieldDescriptorProto_TYPE_UINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return fieldDesc, true
	default:
		return nil, false
	}
}

// IsScalarField checks if the field is a scalar type
func IsScalarField(field proto.Message) bool {
	_, ok := AsScalarField(field)
	return ok
}

// AsMessageField checks if the field is a message type and returns it as a field descriptor
// Returns the field descriptor and true if it's a message field, nil and false otherwise
func AsMessageField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok || fieldDesc.Type == nil {
		return nil, false
	}

	if *fieldDesc.Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE ||
		*fieldDesc.Type == descriptorpb.FieldDescriptorProto_TYPE_GROUP {
		return fieldDesc, true
	}
	return nil, false
}

// IsMessageField checks if the field is a message type
func IsMessageField(field proto.Message) bool {
	_, ok := AsMessageField(field)
	return ok
}

// AsEnumField checks if the field is an enum type and returns it as a field descriptor
// Returns the field descriptor and true if it's an enum field, nil and false otherwise
func AsEnumField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok || fieldDesc.Type == nil {
		return nil, false
	}

	if *fieldDesc.Type == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
		return fieldDesc, true
	}
	return nil, false
}

// IsEnumField checks if the field is an enum type
func IsEnumField(field proto.Message) bool {
	_, ok := AsEnumField(field)
	return ok
}
