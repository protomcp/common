# Generator Module

[![Go Reference][godoc-badge]][godoc-link]
[![Go Report Card][goreportcard-badge]][goreportcard-link]
[![codecov][codecov-badge]][codecov-link]

## Overview

The generator module provides utilities for protoc plugin development, starting
with test utilities for creating descriptor objects in tests.

## Test Utilities

Helper functions for creating descriptor objects in tests:

### Field Creation Functions

Complete field constructors (with name and number):

| Function | Purpose | Parameters | Returns |
|----------|---------|------------|---------|
| `NewField` | Create optional field | `name string, number int32, fieldType descriptorpb.FieldDescriptorProto_Type` | `*descriptorpb.FieldDescriptorProto` |
| `NewRepeatedField` | Create repeated field | `name string, number int32, fieldType descriptorpb.FieldDescriptorProto_Type` | `*descriptorpb.FieldDescriptorProto` |
| `NewRequiredField` | Create required field (proto2) | `name string, number int32, fieldType descriptorpb.FieldDescriptorProto_Type` | `*descriptorpb.FieldDescriptorProto` |
| `NewMessageField` | Create message type field | `name string, number int32, typeName string` | `*descriptorpb.FieldDescriptorProto` |
| `NewEnumField` | Create enum type field | `name string, number int32, typeName string` | `*descriptorpb.FieldDescriptorProto` |
| `NewMapField` | Create map field | `name string, number int32, entryTypeName string` | `*descriptorpb.FieldDescriptorProto` |
| `NewOneOfField` | Create oneof field | `name string, number int32, fieldType descriptorpb.FieldDescriptorProto_Type, oneofIndex int32` | `*descriptorpb.FieldDescriptorProto` |

Minimal field constructors (for testing specific properties):

| Function | Purpose | Parameters | Returns |
|----------|---------|------------|---------|
| `NewFieldWithType` | Create field with type only | `fieldType descriptorpb.FieldDescriptorProto_Type` | `*descriptorpb.FieldDescriptorProto` |
| `NewFieldWithLabel` | Create field with label only | `label descriptorpb.FieldDescriptorProto_Label` | `*descriptorpb.FieldDescriptorProto` |
| `NewRepeatedMessageField` | Create repeated message field | `typeName string` | `*descriptorpb.FieldDescriptorProto` |

### Descriptor Creation Functions

| Function | Purpose | Parameters | Returns |
|----------|---------|------------|---------|
| `NewMessage` | Create message descriptor | `name string, fields ...*descriptorpb.FieldDescriptorProto` | `*descriptorpb.DescriptorProto` |
| `NewMessageWithNested` | Create message with nested types | `name string, messages []*descriptorpb.DescriptorProto, enums []*descriptorpb.EnumDescriptorProto` | `*descriptorpb.DescriptorProto` |
| `NewEnum` | Create enum descriptor | `name string, values ...string` | `*descriptorpb.EnumDescriptorProto` |
| `NewEnumValue` | Create enum value descriptor | `name string, number int32` | `*descriptorpb.EnumValueDescriptorProto` |
| `NewService` | Create service descriptor | `name string, methods ...*descriptorpb.MethodDescriptorProto` | `*descriptorpb.ServiceDescriptorProto` |
| `NewMethod` | Create method descriptor | `name, inputType, outputType string` | `*descriptorpb.MethodDescriptorProto` |
| `NewFile` | Create file descriptor | `name, pkg string, messages []*descriptorpb.DescriptorProto` | `*descriptorpb.FileDescriptorProto` |
| `NewFileWithTypes` | Create file with types | `name, pkg string, messages []*descriptorpb.DescriptorProto, enums []*descriptorpb.EnumDescriptorProto, services []*descriptorpb.ServiceDescriptorProto` | `*descriptorpb.FileDescriptorProto` |
| `NewOneOf` | Create oneof descriptor | `name string` | `*descriptorpb.OneofDescriptorProto` |

### Type Constants

<!-- cspell:ignore SINT SFIXED -->

| Constant | Type | Value |
|----------|------|-------|
| `TypeDouble` | Floating point | `descriptorpb.FieldDescriptorProto_TYPE_DOUBLE` |
| `TypeFloat` | Floating point | `descriptorpb.FieldDescriptorProto_TYPE_FLOAT` |
| `TypeInt64` | Signed integer | `descriptorpb.FieldDescriptorProto_TYPE_INT64` |
| `TypeUInt64` | Unsigned integer | `descriptorpb.FieldDescriptorProto_TYPE_UINT64` |
| `TypeInt32` | Signed integer | `descriptorpb.FieldDescriptorProto_TYPE_INT32` |
| `TypeFixed64` | Fixed size | `descriptorpb.FieldDescriptorProto_TYPE_FIXED64` |
| `TypeFixed32` | Fixed size | `descriptorpb.FieldDescriptorProto_TYPE_FIXED32` |
| `TypeBool` | Boolean | `descriptorpb.FieldDescriptorProto_TYPE_BOOL` |
| `TypeString` | String | `descriptorpb.FieldDescriptorProto_TYPE_STRING` |
| `TypeBytes` | Binary | `descriptorpb.FieldDescriptorProto_TYPE_BYTES` |
| `TypeUInt32` | Unsigned integer | `descriptorpb.FieldDescriptorProto_TYPE_UINT32` |
| `TypeSFixed32` | Signed fixed | `descriptorpb.FieldDescriptorProto_TYPE_SFIXED32` |
| `TypeSFixed64` | Signed fixed | `descriptorpb.FieldDescriptorProto_TYPE_SFIXED64` |
| `TypeSInt32` | Signed integer | `descriptorpb.FieldDescriptorProto_TYPE_SINT32` |
| `TypeSInt64` | Signed integer | `descriptorpb.FieldDescriptorProto_TYPE_SINT64` |
| `TypeGroup` | Group (deprecated) | `descriptorpb.FieldDescriptorProto_TYPE_GROUP` |
| `TypeMessage` | Message | `descriptorpb.FieldDescriptorProto_TYPE_MESSAGE` |
| `TypeEnum` | Enum | `descriptorpb.FieldDescriptorProto_TYPE_ENUM` |

### Example Usage

```go
import (
    "testing"

    "darvaza.org/core"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/descriptorpb"
    "protomcp.org/common/generator"
)

func TestMyGenerator(t *testing.T) {
    // Creating complete message descriptors for testing
    msg := &descriptorpb.DescriptorProto{
        Name: proto.String("User"),
        Field: []*descriptorpb.FieldDescriptorProto{
            generator.NewField("id", 1, generator.TypeInt64),
            generator.NewRepeatedField("tags", 2, generator.TypeString),
            generator.NewMessageField("profile", 3, ".example.Profile"),
            generator.NewEnumField("status", 4, ".example.Status"),
            generator.NewMapField("metadata", 5, ".MetadataEntry"),
            generator.NewOneOfField("variant", 6, generator.TypeBool, 0),
        },
    }
    // Test generator with the constructed message
    core.AssertNotNil(t, msg, "message")
    core.AssertEqual(t, "User", msg.GetName(), "message name")
}

// Minimal field descriptors for testing specific behaviour
func TestFieldProperties(t *testing.T) {
    // Create fields with only the properties needed for testing
    messageField := generator.NewFieldWithType(generator.TypeMessage)
    repeatedField := generator.NewFieldWithLabel(
        descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
    mapField := generator.NewRepeatedMessageField(".MapEntry")

    // Test field properties using core assertions
    core.AssertEqual(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
        messageField.GetType(), "field type")
    core.AssertEqual(t, descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
        repeatedField.GetLabel(), "field label")
    core.AssertNotNil(t, mapField.TypeName, "map field type name")
}
```

## Planned Core Functionality

The following sections describe planned functionality that will be added
incrementally as the protomcp.org ecosystem develops.

### Descriptor Type Checking

Essential for the options module and other generators:

```go
// Type checking - needed by options for conditional logic
func IsMessageType(desc proto.Message, name string) bool
func IsFieldType(desc proto.Message) bool
func IsServiceType(desc proto.Message) bool
func IsMethodType(desc proto.Message) bool
func IsEnumType(desc proto.Message) bool
func IsFileType(desc proto.Message) bool

// Specific type checks
func IsRepeatedField(field proto.Message) bool
func IsMapField(field proto.Message) bool
func IsOneofField(field proto.Message) bool
func IsOptionalField(field proto.Message) bool
func IsRequiredField(field proto.Message) bool
```

### Path and Naming Utilities

Required by options module for selector matching:

```go
// Path construction - needed for option selectors
func GetMessageName(desc proto.Message) string
func GetFieldPath(file, msg, field proto.Message) string
func GetFullName(desc proto.Message) string
func GetPackage(file proto.Message) string
func GetQualifiedName(desc proto.Message) string

// Name manipulation
func ToGoName(name string) string
func ToCamelCase(name string) string
func ToSnakeCase(name string) string
func ToKebabCase(name string) string
```

### Descriptor Traversal

Critical for options module's hook system:

```go
// Finding descriptors - needed for option resolution
func FindMessage(file proto.Message, name string) proto.Message
func FindField(msg proto.Message, name string) proto.Message
func FindEnum(file proto.Message, name string) proto.Message
func FindService(file proto.Message, name string) proto.Message
func FindMethod(service proto.Message, name string) proto.Message

// Iteration helpers - needed for applying option hooks
// Callback returns true to continue, false to stop (Go 1.23 convention)
func ForEachField(msg proto.Message, fn func(field proto.Message) bool) error
func ForEachMessage(file proto.Message, fn func(msg proto.Message) bool) error
func ForEachEnum(file proto.Message, fn func(enum proto.Message) bool) error
func ForEachService(file proto.Message, fn func(svc proto.Message) bool) error
func ForEachMethod(svc proto.Message, fn func(method proto.Message) bool) error

// Nested traversal
func ForEachNestedMessage(
    msg proto.Message, fn func(nested proto.Message) bool) error
func ForEachNestedEnum(
    msg proto.Message, fn func(enum proto.Message) bool) error
```

### Tree Walking

Simple functions for walking descriptor trees:

```go
// Walk descriptor tree with callback functions
func Walk(file proto.Message, onMessage func(proto.Message) error) error
func WalkMessage(msg proto.Message, onField func(proto.Message) error) error

// Path-aware walking
func WalkWithPath(file proto.Message,
    callback func(path []string, desc proto.Message) error) error
```

### Type Resolution

Needed by options for type-safe option handling:

```go
// Type information
func GetFieldType(field proto.Message) descriptorpb.FieldDescriptorProto_Type
func GetFieldTypeName(field proto.Message) string
func IsScalarType(field proto.Message) bool
func IsMessageType(field proto.Message) bool
func IsEnumType(field proto.Message) bool

// Type lookup
func ResolveTypeName(file proto.Message, typeName string) proto.Message
func GetTypeDescriptor(field proto.Message) proto.Message
```

### Dependency Analysis

Useful for both options and code generation:

```go
// Dependency tracking
func GetFileDependencies(file proto.Message) []string
func GetMessageDependencies(msg proto.Message) []string
func GetFieldDependencies(field proto.Message) []string

// Import management
func GetRequiredImports(file proto.Message) []string
func ResolveImportPath(file proto.Message, typeName string) string
```

## Comment Access

Helper functions for reading comments from descriptors:

```go
// Comment reading
func GetLeadingComments(desc proto.Message) string
func GetTrailingComments(desc proto.Message) string
```

## Best Practices

1. **Nil Safety**: All functions handle nil inputs gracefully.
2. **Error Propagation**: Iteration functions stop on first error.
3. **Type Safety**: Use type assertions with checks for proto.Message types.
4. **Performance**: Cache frequently accessed paths and names.
5. **Compatibility**: Support both proto2 and proto3 descriptors.

## Future Extensions

- Additional traversal helpers.
- Dependency graph visualization.
- Performance optimizations for large schemas.
- Parallel traversal support.

[godoc-badge]: https://pkg.go.dev/badge/protomcp.org/common/generator.svg
[godoc-link]: https://pkg.go.dev/protomcp.org/common/generator
[goreportcard-badge]: https://goreportcard.com/badge/protomcp.org/common
[goreportcard-link]: https://goreportcard.com/report/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg?flag=generator
[codecov-link]: https://codecov.io/gh/protomcp/common?flag=generator
