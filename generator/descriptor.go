package generator

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// AsMessage attempts to cast the descriptor to a message descriptor and validates it has a Name.
// Returns the message descriptor and true if successful and Name is not nil/empty, nil and false otherwise.
// A DescriptorProto without a Name is considered invalid as every message must have a name.
func AsMessage(desc proto.Message) (*descriptorpb.DescriptorProto, bool) {
	msgDesc, ok := desc.(*descriptorpb.DescriptorProto)
	if ok && isPointerNonZero(msgDesc.Name) {
		return msgDesc, true
	}
	return nil, false
}

// IsMessage checks if the descriptor is a message descriptor.
// Returns true if the descriptor is a DescriptorProto, false otherwise.
func IsMessage(desc proto.Message) bool {
	_, ok := AsMessage(desc)
	return ok
}

// AsMessageWithName attempts to cast the descriptor to a message descriptor and checks if it has the given name.
// Returns the message descriptor and true if successful and name matches (empty name matches any).
func AsMessageWithName(desc proto.Message, name string) (*descriptorpb.DescriptorProto, bool) {
	msgDesc, ok := AsMessage(desc)
	switch {
	case !ok:
		// Wrong type
		return nil, false
	case name == "", isPointerEqual(msgDesc.Name, name):
		// Match or any accepted
		return msgDesc, true
	default:
		// Name doesn't match
		return nil, false
	}
}

// IsMessageWithName checks if the descriptor is a message type with the given name.
// Returns true if the descriptor is a DescriptorProto with matching name, false otherwise.
func IsMessageWithName(desc proto.Message, name string) bool {
	_, ok := AsMessageWithName(desc, name)
	return ok
}

// AsFieldType attempts to cast the descriptor to a field descriptor and validates it has a Type.
// Returns the field descriptor and true if successful and Type is not nil, nil and false otherwise.
// A FieldDescriptorProto without a Type is considered invalid as every field must have a type.
func AsFieldType(desc proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := desc.(*descriptorpb.FieldDescriptorProto)
	if ok && fieldDesc.Type != nil {
		return fieldDesc, true
	}
	return nil, false
}

// IsFieldType checks if the descriptor is a valid field descriptor.
// Returns true if the descriptor is a FieldDescriptorProto with Type set, false otherwise.
func IsFieldType(desc proto.Message) bool {
	_, ok := AsFieldType(desc)
	return ok
}

// AsServiceType attempts to cast the descriptor to a service descriptor and validates it has a Name.
// Returns the service descriptor and true if successful and Name is not nil/empty, nil and false otherwise.
// A ServiceDescriptorProto without a Name is considered invalid as every service must have a name.
func AsServiceType(desc proto.Message) (*descriptorpb.ServiceDescriptorProto, bool) {
	svcDesc, ok := desc.(*descriptorpb.ServiceDescriptorProto)
	if ok && isPointerNonZero(svcDesc.Name) {
		return svcDesc, true
	}
	return nil, false
}

// IsServiceType checks if the descriptor is a service descriptor.
// Returns true if the descriptor is a ServiceDescriptorProto, false otherwise.
func IsServiceType(desc proto.Message) bool {
	_, ok := AsServiceType(desc)
	return ok
}

// AsMethodType attempts to cast the descriptor to a method descriptor and validates essential fields.
// Returns the method descriptor and true if successful with valid Name, InputType and OutputType.
// A MethodDescriptorProto without these fields is considered invalid.
func AsMethodType(desc proto.Message) (*descriptorpb.MethodDescriptorProto, bool) {
	methodDesc, ok := desc.(*descriptorpb.MethodDescriptorProto)
	switch {
	case !ok:
		return nil, false
	case !isPointerNonZero(methodDesc.Name):
		return nil, false
	case !isPointerNonZero(methodDesc.InputType):
		return nil, false
	case !isPointerNonZero(methodDesc.OutputType):
		return nil, false
	default:
		return methodDesc, true
	}
}

// IsMethodType checks if the descriptor is a method descriptor.
// Returns true if the descriptor is a MethodDescriptorProto, false otherwise.
func IsMethodType(desc proto.Message) bool {
	_, ok := AsMethodType(desc)
	return ok
}

// AsEnumType attempts to cast the descriptor to an enum descriptor and validates it has a Name.
// Returns the enum descriptor and true if successful and Name is not nil/empty, nil and false otherwise.
// An EnumDescriptorProto without a Name is considered invalid as every enum must have a name.
func AsEnumType(desc proto.Message) (*descriptorpb.EnumDescriptorProto, bool) {
	enumDesc, ok := desc.(*descriptorpb.EnumDescriptorProto)
	if ok && isPointerNonZero(enumDesc.Name) {
		return enumDesc, true
	}
	return nil, false
}

// IsEnumType checks if the descriptor is an enum descriptor.
// Returns true if the descriptor is an EnumDescriptorProto, false otherwise.
func IsEnumType(desc proto.Message) bool {
	_, ok := AsEnumType(desc)
	return ok
}

// AsFileType attempts to cast the descriptor to a file descriptor and validates it has a Name.
// Returns the file descriptor and true if successful and Name is not nil/empty, nil and false otherwise.
// A FileDescriptorProto without a Name is considered invalid as every file must have a name.
func AsFileType(desc proto.Message) (*descriptorpb.FileDescriptorProto, bool) {
	fileDesc, ok := desc.(*descriptorpb.FileDescriptorProto)
	if ok && isPointerNonZero(fileDesc.Name) {
		return fileDesc, true
	}
	return nil, false
}

// IsFileType checks if the descriptor is a file descriptor.
// Returns true if the descriptor is a FileDescriptorProto, false otherwise.
func IsFileType(desc proto.Message) bool {
	_, ok := AsFileType(desc)
	return ok
}

// AsRepeatedField checks if the field is repeated and returns it as a field descriptor.
// Returns the field descriptor and true if it's a repeated field, nil and false otherwise.
func AsRepeatedField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	switch {
	case !ok:
		// Wrong type
		return nil, false
	case !isPointerEqual(fieldDesc.Label, LabelRepeated):
		// Missing label or not repeated
		return nil, false
	default:
		// Is repeated
		return fieldDesc, true
	}
}

// IsRepeatedField checks if the field is repeated.
// Returns true if the field has LABEL_REPEATED, false otherwise.
func IsRepeatedField(field proto.Message) bool {
	_, ok := AsRepeatedField(field)
	return ok
}

// AsMapField checks if the field is a map field and returns it as a field descriptor.
// In protobuf, map fields are represented as repeated message fields
// where the message type is a special map entry type.
// This function uses a heuristic: type names ending with "Entry".
// For definitive checking, use AsMapFieldWithMessage.
// Returns the field descriptor and true if it's a map field, nil and false otherwise.
func AsMapField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	// Must be a repeated message field
	fieldDesc, ok := AsRepeatedField(field)
	switch {
	case !ok:
		return nil, false
	case !isPointerEqual(fieldDesc.Type, TypeMessage):
		return nil, false
	case !isPointerNonZero(fieldDesc.TypeName):
		return nil, false
	case !strings.HasSuffix(*fieldDesc.TypeName, "Entry"):
		return nil, false
	default:
		// Check if the type name indicates a map entry.
		// This is a heuristic - map entry messages typically end with "Entry".
		// This is not definitive - false positives/negatives are possible.
		return fieldDesc, true
	}
}

// AsMapFieldWithMessage checks if the field is a map field by examining the map entry's message descriptor.
// This is the definitive check as it verifies the map_entry=true option.
// Returns the field descriptor and true if it's a map field, nil and false otherwise.
func AsMapFieldWithMessage(field proto.Message, entryMsg proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsMapField(field)
	switch {
	case !ok:
		// Not a map field according to heuristic
		return nil, false
	case entryMsg == nil:
		// no message to validate against, accept
		return fieldDesc, true
	}

	// Check for map_entry option in the message descriptor
	msg, ok := AsMessage(entryMsg)
	switch {
	case !ok, msg.Options == nil, !isPointerNonZero(msg.Options.MapEntry):
		return nil, false
	default:
		return fieldDesc, true
	}
}

// IsMapField checks if the field is a map field.
// Returns true if the field represents a protobuf map, false otherwise.
func IsMapField(field proto.Message) bool {
	_, ok := AsMapField(field)
	return ok
}

// IsMapFieldWithMessage checks if the field is a map field with definitive message descriptor check.
// Returns true if the field represents a protobuf map with map_entry=true, false otherwise.
func IsMapFieldWithMessage(field proto.Message, entryMsg proto.Message) bool {
	_, ok := AsMapFieldWithMessage(field, entryMsg)
	return ok
}

// AsOneOfField checks if the field is part of a oneof and returns it as a field descriptor.
// Returns the field descriptor and true if it's part of a oneof, nil and false otherwise.
func AsOneOfField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	switch {
	case !ok:
		return nil, false
	case fieldDesc.OneofIndex == nil:
		// A field is part of a oneof if it has a OneofIndex set
		return nil, false
	default:
		return fieldDesc, true
	}
}

// IsOneOfField checks if the field is part of a oneof.
// Returns true if the field has a OneofIndex set, false otherwise.
func IsOneOfField(field proto.Message) bool {
	_, ok := AsOneOfField(field)
	return ok
}

// AsOptionalField checks if the field is optional and returns it as a field descriptor.
// Returns the field descriptor and true if it has LABEL_OPTIONAL, nil and false otherwise.
//
// Note: This returns true for all fields with LABEL_OPTIONAL, which includes:
//   - Proto2 optional fields
//   - Proto3 optional fields (with 'optional' keyword, also have Proto3Optional=true)
//   - Proto3 singular fields (without 'optional' keyword, have Proto3Optional=false/nil)
//
// To distinguish proto3 'optional' fields specifically, check the Proto3Optional field.
func AsOptionalField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	switch {
	case !ok:
		// Wrong type
		return nil, false
	case !isPointerEqual(fieldDesc.Label, LabelOptional):
		// Missing label or not optional
		return nil, false
	default:
		// Is optional
		return fieldDesc, true
	}
}

// IsOptionalField checks if the field is optional.
// Returns true if the field has LABEL_OPTIONAL, false otherwise.
//
// Note: This returns true for all fields with LABEL_OPTIONAL, which includes:
//   - Proto2 optional fields
//   - Proto3 optional fields (with 'optional' keyword, also have Proto3Optional=true)
//   - Proto3 singular fields (without 'optional' keyword, have Proto3Optional=false/nil)
//
// To distinguish proto3 'optional' fields specifically, check the Proto3Optional field.
func IsOptionalField(field proto.Message) bool {
	_, ok := AsOptionalField(field)
	return ok
}

// AsRequiredField checks if the field is required and returns it as a field descriptor.
// Returns the field descriptor and true if it's required, nil and false otherwise (proto2 only).
func AsRequiredField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	switch {
	case !ok:
		// Wrong type
		return nil, false
	case !isPointerEqual(fieldDesc.Label, LabelRequired):
		// Missing label or not required
		return nil, false
	default:
		// Is required
		return fieldDesc, true
	}
}

// IsRequiredField checks if the field is required (proto2 only).
// Returns true if the field has LABEL_REQUIRED, false otherwise.
func IsRequiredField(field proto.Message) bool {
	_, ok := AsRequiredField(field)
	return ok
}

// AsScalarField checks if the field is a scalar type and returns it as a field descriptor.
// Returns the field descriptor and true if it's a scalar field, nil and false otherwise.
func AsScalarField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if !ok {
		// Wrong type or missing Type field
		return nil, false
	}

	switch *fieldDesc.Type {
	case TypeDouble, TypeFloat,
		TypeInt64, TypeUInt64, TypeInt32, TypeUInt32,
		TypeFixed64, TypeFixed32,
		TypeSFixed32, TypeSFixed64,
		TypeSInt32, TypeSInt64,
		TypeBool, TypeString, TypeBytes:
		return fieldDesc, true
	default:
		return nil, false
	}
}

// IsScalarField checks if the field is a scalar type.
// Returns true if the field is a scalar protobuf type, false otherwise.
func IsScalarField(field proto.Message) bool {
	_, ok := AsScalarField(field)
	return ok
}

// AsMessageField checks if the field is a message type and returns it as a field descriptor.
// Returns the field descriptor and true if it's a TYPE_MESSAGE field, nil and false otherwise.
// Note: This does not include TYPE_GROUP fields. Use AsGroupField for groups.
func AsMessageField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if ok && isPointerEqual(fieldDesc.Type, TypeMessage) {
		return fieldDesc, true
	}
	return nil, false
}

// IsMessageField checks if the field is a message type.
// Returns true if the field is TYPE_MESSAGE, false otherwise.
// Note: This does not include TYPE_GROUP fields. Use IsGroupField for groups.
func IsMessageField(field proto.Message) bool {
	_, ok := AsMessageField(field)
	return ok
}

// AsGroupField checks if the field is a group type and returns it as a field descriptor.
// Returns the field descriptor and true if it's a TYPE_GROUP field, nil and false otherwise.
// Note: Groups are deprecated in proto2 and not supported in proto3. Use message types instead.
func AsGroupField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if ok && isPointerEqual(fieldDesc.Type, TypeGroup) {
		return fieldDesc, true
	}
	return nil, false
}

// IsGroupField checks if the field is a group type.
// Returns true if the field is TYPE_GROUP, false otherwise.
// Note: Groups are deprecated in proto2 and not supported in proto3. Use message types instead.
func IsGroupField(field proto.Message) bool {
	_, ok := AsGroupField(field)
	return ok
}

// AsEnumField checks if the field is an enum type and returns it as a field descriptor.
// Returns the field descriptor and true if it's an enum field, nil and false otherwise.
func AsEnumField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool) {
	fieldDesc, ok := AsFieldType(field)
	if ok && isPointerEqual(fieldDesc.Type, TypeEnum) {
		return fieldDesc, true
	}
	return nil, false
}

// IsEnumField checks if the field is an enum type.
// Returns true if the field is TYPE_ENUM, false otherwise.
func IsEnumField(field proto.Message) bool {
	_, ok := AsEnumField(field)
	return ok
}

func isPointerEqual[T comparable](p *T, v T) bool {
	return p != nil && *p == v
}

func isPointerNonZero[T comparable](p *T) bool {
	var zero T
	return p != nil && *p != zero
}
