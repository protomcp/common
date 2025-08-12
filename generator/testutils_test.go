package generator

import (
	"testing"

	"darvaza.org/core"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Test NewField function
func TestNewField(t *testing.T) {
	field := NewField("test_field", 1, TypeString)

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "test_field", field.GetName(), "field name")
	core.AssertEqual(t, int32(1), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_STRING, field.GetType(), "field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL, field.GetLabel(), "field label")
}

// Test NewRepeatedField function
func TestNewRepeatedField(t *testing.T) {
	field := NewRepeatedField("items", 2, TypeInt32)

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "items", field.GetName(), "field name")
	core.AssertEqual(t, int32(2), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_INT32, field.GetType(), "field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, field.GetLabel(), "field label")
}

// Test NewRequiredField function
func TestNewRequiredField(t *testing.T) {
	field := NewRequiredField("id", 3, TypeUInt64)

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "id", field.GetName(), "field name")
	core.AssertEqual(t, int32(3), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_UINT64, field.GetType(), "field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REQUIRED, field.GetLabel(), "field label")
}

// Test NewMessageField function
func TestNewMessageField(t *testing.T) {
	field := NewMessageField("user", 4, ".example.User")

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "user", field.GetName(), "field name")
	core.AssertEqual(t, int32(4), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, field.GetType(), "field type")
	core.AssertEqual(t, ".example.User", field.GetTypeName(), "field type name")
}

// Test NewEnumField function
func TestNewEnumField(t *testing.T) {
	field := NewEnumField("status", 5, ".example.Status")

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "status", field.GetName(), "field name")
	core.AssertEqual(t, int32(5), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_ENUM, field.GetType(), "field type")
	core.AssertEqual(t, ".example.Status", field.GetTypeName(), "field type name")
}

// Test NewMapField function
func TestNewMapField(t *testing.T) {
	field := NewMapField("attributes", 6, ".AttributesEntry")

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "attributes", field.GetName(), "field name")
	core.AssertEqual(t, int32(6), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, field.GetType(), "field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, field.GetLabel(), "field label")
	core.AssertEqual(t, ".AttributesEntry", field.GetTypeName(), "field type name")
}

// Test NewOneOfField function
func TestNewOneOfField(t *testing.T) {
	field := NewOneOfField("choice", 7, TypeBool, 0)

	core.AssertNotNil(t, field, "field")
	core.AssertEqual(t, "choice", field.GetName(), "field name")
	core.AssertEqual(t, int32(7), field.GetNumber(), "field number")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_BOOL, field.GetType(), "field type")
	core.AssertNotNil(t, field.OneofIndex, "oneof index")
	core.AssertEqual(t, int32(0), *field.OneofIndex, "oneof index value")
}

// Test creating complex message structures
func TestComplexMessageCreation(t *testing.T) {
	// Create a message with various field types
	msg := &descriptorpb.DescriptorProto{
		Name: proto.String("TestMessage"),
		Field: []*descriptorpb.FieldDescriptorProto{
			NewField("id", 1, TypeInt64),
			NewRepeatedField("tags", 2, TypeString),
			NewMessageField("user", 3, ".example.User"),
			NewEnumField("status", 4, ".example.Status"),
			NewMapField("metadata", 5, ".MetadataEntry"),
			NewOneOfField("variant", 6, TypeBool, 0),
		},
	}

	core.AssertNotNil(t, msg, "message")
	core.AssertEqual(t, "TestMessage", msg.GetName(), "message name")
	core.AssertEqual(t, 6, len(msg.Field), "field count")

	// Verify each field
	core.AssertEqual(t, "id", msg.Field[0].GetName(), "field 0 name")
	core.AssertEqual(t, "tags", msg.Field[1].GetName(), "field 1 name")
	core.AssertEqual(t, "user", msg.Field[2].GetName(), "field 2 name")
	core.AssertEqual(t, "status", msg.Field[3].GetName(), "field 3 name")
	core.AssertEqual(t, "metadata", msg.Field[4].GetName(), "field 4 name")
	core.AssertEqual(t, "variant", msg.Field[5].GetName(), "field 5 name")

	// Verify field types
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_INT64, msg.Field[0].GetType(), "field 0 type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_STRING, msg.Field[1].GetType(), "field 1 type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, msg.Field[2].GetType(), "field 2 type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_ENUM, msg.Field[3].GetType(), "field 3 type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, msg.Field[4].GetType(), "field 4 type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_BOOL, msg.Field[5].GetType(), "field 5 type")

	// Verify labels
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL, msg.Field[0].GetLabel(), "field 0 label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, msg.Field[1].GetLabel(), "field 1 label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, msg.Field[4].GetLabel(), "field 4 label")
}

// Test field helpers with edge cases
func TestFieldHelpersEdgeCases(t *testing.T) {
	// Test with empty name
	field := NewField("", 1, TypeString)
	core.AssertEqual(t, "", field.GetName(), "empty name")

	// Test with zero number
	field = NewField("test", 0, TypeString)
	core.AssertEqual(t, int32(0), field.GetNumber(), "zero number")

	// Test with negative number
	field = NewField("test", -1, TypeString)
	core.AssertEqual(t, int32(-1), field.GetNumber(), "negative number")

	// Test map field with empty name
	field = NewMapField("", 1, ".Entry")
	core.AssertEqual(t, "", field.GetName(), "empty map name")
	core.AssertEqual(t, ".Entry", field.GetTypeName(), "empty map type name")

	// Test oneof with large index
	field = NewOneOfField("test", 1, TypeString, 999)
	core.AssertNotNil(t, field.OneofIndex, "oneof index")
	core.AssertEqual(t, int32(999), *field.OneofIndex, "large oneof index")
}

// Test NewMessage function
func TestNewMessage(t *testing.T) {
	fields := []*descriptorpb.FieldDescriptorProto{
		NewField("id", 1, TypeInt64),
		NewField("name", 2, TypeString),
	}

	msg := NewMessage("User", fields...)

	core.AssertNotNil(t, msg, "message")
	core.AssertEqual(t, "User", msg.GetName(), "message name")
	core.AssertEqual(t, 2, len(msg.Field), "field count")
	core.AssertEqual(t, "id", msg.Field[0].GetName(), "first field name")
	core.AssertEqual(t, "name", msg.Field[1].GetName(), "second field name")

	// Test with no fields
	emptyMsg := NewMessage("Empty")
	core.AssertNotNil(t, emptyMsg, "empty message")
	core.AssertEqual(t, "Empty", emptyMsg.GetName(), "empty message name")
	core.AssertEqual(t, 0, len(emptyMsg.Field), "empty field count")
}

// Test NewMessageWithNested function
func TestNewMessageWithNested(t *testing.T) {
	fields := []*descriptorpb.FieldDescriptorProto{
		NewField("id", 1, TypeInt64),
	}

	nestedMessages := []*descriptorpb.DescriptorProto{
		NewMessage("Nested", NewField("value", 1, TypeString)),
	}

	nestedEnums := []*descriptorpb.EnumDescriptorProto{
		NewEnum("Status", "UNKNOWN", "ACTIVE", "INACTIVE"),
	}

	msg := NewMessageWithNested("ComplexMessage", fields, nestedMessages, nestedEnums)

	core.AssertNotNil(t, msg, "message")
	core.AssertEqual(t, "ComplexMessage", msg.GetName(), "message name")
	core.AssertEqual(t, 1, len(msg.Field), "field count")
	core.AssertEqual(t, 1, len(msg.NestedType), "nested message count")
	core.AssertEqual(t, 1, len(msg.EnumType), "nested enum count")
	core.AssertEqual(t, "Nested", msg.NestedType[0].GetName(), "nested message name")
	core.AssertEqual(t, "Status", msg.EnumType[0].GetName(), "nested enum name")
}

// Test NewEnum function
func TestNewEnum(t *testing.T) {
	enum := NewEnum("Colour", "RED", "GREEN", "BLUE")

	core.AssertNotNil(t, enum, "enum")
	core.AssertEqual(t, "Colour", enum.GetName(), "enum name")
	core.AssertEqual(t, 3, len(enum.Value), "value count")
	core.AssertEqual(t, "RED", enum.Value[0].GetName(), "first value name")
	core.AssertEqual(t, int32(0), enum.Value[0].GetNumber(), "first value number")
	core.AssertEqual(t, "GREEN", enum.Value[1].GetName(), "second value name")
	core.AssertEqual(t, int32(1), enum.Value[1].GetNumber(), "second value number")
	core.AssertEqual(t, "BLUE", enum.Value[2].GetName(), "third value name")
	core.AssertEqual(t, int32(2), enum.Value[2].GetNumber(), "third value number")

	// Test with no values
	emptyEnum := NewEnum("Empty")
	core.AssertNotNil(t, emptyEnum, "empty enum")
	core.AssertEqual(t, "Empty", emptyEnum.GetName(), "empty enum name")
	core.AssertEqual(t, 0, len(emptyEnum.Value), "empty value count")
}

// Test NewEnumValue function
func TestNewEnumValue(t *testing.T) {
	value := NewEnumValue("SUCCESS", 42)

	core.AssertNotNil(t, value, "enum value")
	core.AssertEqual(t, "SUCCESS", value.GetName(), "value name")
	core.AssertEqual(t, int32(42), value.GetNumber(), "value number")

	// Test with negative number
	negValue := NewEnumValue("ERROR", -1)
	core.AssertNotNil(t, negValue, "negative enum value")
	core.AssertEqual(t, "ERROR", negValue.GetName(), "negative value name")
	core.AssertEqual(t, int32(-1), negValue.GetNumber(), "negative value number")
}

// Test NewService function
func TestNewService(t *testing.T) {
	methods := []*descriptorpb.MethodDescriptorProto{
		NewMethod("GetUser", ".GetUserRequest", ".GetUserResponse"),
		NewMethod("UpdateUser", ".UpdateUserRequest", ".UpdateUserResponse"),
	}

	service := NewService("UserService", methods...)

	core.AssertNotNil(t, service, "service")
	core.AssertEqual(t, "UserService", service.GetName(), "service name")
	core.AssertEqual(t, 2, len(service.Method), "method count")
	core.AssertEqual(t, "GetUser", service.Method[0].GetName(), "first method name")
	core.AssertEqual(t, "UpdateUser", service.Method[1].GetName(), "second method name")

	// Test with no methods
	emptyService := NewService("EmptyService")
	core.AssertNotNil(t, emptyService, "empty service")
	core.AssertEqual(t, "EmptyService", emptyService.GetName(), "empty service name")
	core.AssertEqual(t, 0, len(emptyService.Method), "empty method count")
}

// Test NewMethod function
func TestNewMethod(t *testing.T) {
	method := NewMethod("GetUser", ".GetUserRequest", ".GetUserResponse")

	core.AssertNotNil(t, method, "method")
	core.AssertEqual(t, "GetUser", method.GetName(), "method name")
	core.AssertEqual(t, ".GetUserRequest", method.GetInputType(), "input type")
	core.AssertEqual(t, ".GetUserResponse", method.GetOutputType(), "output type")

	// Test with empty types
	emptyMethod := NewMethod("Empty", "", "")
	core.AssertNotNil(t, emptyMethod, "empty method")
	core.AssertEqual(t, "Empty", emptyMethod.GetName(), "empty method name")
	core.AssertEqual(t, "", emptyMethod.GetInputType(), "empty input type")
	core.AssertEqual(t, "", emptyMethod.GetOutputType(), "empty output type")
}

// Test NewFile function
func TestNewFile(t *testing.T) {
	file := NewFile("user.proto", "example.user")

	core.AssertNotNil(t, file, "file")
	core.AssertEqual(t, "user.proto", file.GetName(), "file name")
	core.AssertEqual(t, "example.user", file.GetPackage(), "package name")

	// Test with empty values
	emptyFile := NewFile("", "")
	core.AssertNotNil(t, emptyFile, "empty file")
	core.AssertEqual(t, "", emptyFile.GetName(), "empty file name")
	core.AssertEqual(t, "", emptyFile.GetPackage(), "empty package name")
}

// Test NewFileWithTypes function
func TestNewFileWithTypes(t *testing.T) {
	messages := []*descriptorpb.DescriptorProto{
		NewMessage("User", NewField("id", 1, TypeInt64)),
		NewMessage("Profile", NewField("bio", 1, TypeString)),
	}

	enums := []*descriptorpb.EnumDescriptorProto{
		NewEnum("Status", "ACTIVE", "INACTIVE"),
	}

	services := []*descriptorpb.ServiceDescriptorProto{
		NewService("UserService", NewMethod("GetUser", ".GetUserRequest", ".GetUserResponse")),
	}

	file := NewFileWithTypes("api.proto", "example.api", messages, enums, services)

	core.AssertNotNil(t, file, "file")
	core.AssertEqual(t, "api.proto", file.GetName(), "file name")
	core.AssertEqual(t, "example.api", file.GetPackage(), "package name")
	core.AssertEqual(t, 2, len(file.MessageType), "message count")
	core.AssertEqual(t, 1, len(file.EnumType), "enum count")
	core.AssertEqual(t, 1, len(file.Service), "service count")
	core.AssertEqual(t, "User", file.MessageType[0].GetName(), "first message name")
	core.AssertEqual(t, "Profile", file.MessageType[1].GetName(), "second message name")
	core.AssertEqual(t, "Status", file.EnumType[0].GetName(), "enum name")
	core.AssertEqual(t, "UserService", file.Service[0].GetName(), "service name")

	// Test with nil slices
	nilFile := NewFileWithTypes("nil.proto", "example", nil, nil, nil)
	core.AssertNotNil(t, nilFile, "nil file")
	core.AssertEqual(t, 0, len(nilFile.MessageType), "nil message count")
	core.AssertEqual(t, 0, len(nilFile.EnumType), "nil enum count")
	core.AssertEqual(t, 0, len(nilFile.Service), "nil service count")
}

// Test NewOneOf function
func TestNewOneOf(t *testing.T) {
	oneof := NewOneOf("response")

	core.AssertNotNil(t, oneof, "oneof")
	core.AssertEqual(t, "response", oneof.GetName(), "oneof name")

	// Test with empty name
	emptyOneOf := NewOneOf("")
	core.AssertNotNil(t, emptyOneOf, "empty oneof")
	core.AssertEqual(t, "", emptyOneOf.GetName(), "empty oneof name")
}

// Test NewFieldWithType function
func TestNewFieldWithType(t *testing.T) {
	field := NewFieldWithType(TypeString)

	core.AssertNotNil(t, field, "field")
	core.AssertNotNil(t, field.Type, "field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_STRING, *field.Type, "field type value")
	core.AssertNil(t, field.Name, "field name should be nil")
	core.AssertNil(t, field.Number, "field number should be nil")
	core.AssertNil(t, field.Label, "field label should be nil")

	// Test with different types
	messageField := NewFieldWithType(TypeMessage)
	core.AssertNotNil(t, messageField, "message field")
	core.AssertNotNil(t, messageField.Type, "message field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, *messageField.Type, "message field type value")

	enumField := NewFieldWithType(TypeEnum)
	core.AssertNotNil(t, enumField, "enum field")
	core.AssertNotNil(t, enumField.Type, "enum field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_ENUM, *enumField.Type, "enum field type value")
}

// Test NewFieldWithLabel function
func TestNewFieldWithLabel(t *testing.T) {
	field := NewFieldWithLabel(descriptorpb.FieldDescriptorProto_LABEL_REPEATED)

	core.AssertNotNil(t, field, "field")
	core.AssertNotNil(t, field.Label, "field label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, *field.Label, "field label value")
	core.AssertNil(t, field.Name, "field name should be nil")
	core.AssertNil(t, field.Number, "field number should be nil")
	// Now includes a default Type field for validity
	core.AssertNotNil(t, field.Type, "field type")
	core.AssertEqual(t, TypeString, *field.Type, "field type default")

	// Test with different labels
	optionalField := NewFieldWithLabel(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL)
	core.AssertNotNil(t, optionalField, "optional field")
	core.AssertNotNil(t, optionalField.Label, "optional field label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
		*optionalField.Label, "optional field label value")
	core.AssertNotNil(t, optionalField.Type, "optional field type")

	requiredField := NewFieldWithLabel(descriptorpb.FieldDescriptorProto_LABEL_REQUIRED)
	core.AssertNotNil(t, requiredField, "required field")
	core.AssertNotNil(t, requiredField.Label, "required field label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REQUIRED,
		*requiredField.Label, "required field label value")
	core.AssertNotNil(t, requiredField.Type, "required field type")
}

// Test NewRepeatedMessageField function
func TestNewRepeatedMessageField(t *testing.T) {
	// Test with type name
	field := NewRepeatedMessageField(".example.Message")

	core.AssertNotNil(t, field, "field")
	core.AssertNotNil(t, field.Label, "field label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, *field.Label, "field label value")
	core.AssertNotNil(t, field.Type, "field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, *field.Type, "field type value")
	core.AssertNotNil(t, field.TypeName, "field type name")
	core.AssertEqual(t, ".example.Message", *field.TypeName, "field type name value")

	// Test with empty type name (should not set TypeName)
	emptyField := NewRepeatedMessageField("")
	core.AssertNotNil(t, emptyField, "empty field")
	core.AssertNotNil(t, emptyField.Label, "empty field label")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED, *emptyField.Label, "empty field label value")
	core.AssertNotNil(t, emptyField.Type, "empty field type")
	core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, *emptyField.Type, "empty field type value")
	core.AssertNil(t, emptyField.TypeName, "empty field type name should be nil")

	// Test for map field pattern
	mapField := NewRepeatedMessageField(".MapEntry")
	core.AssertNotNil(t, mapField, "map field")
	core.AssertNotNil(t, mapField.TypeName, "map field type name")
	core.AssertEqual(t, ".MapEntry", *mapField.TypeName, "map field type name value")
}
