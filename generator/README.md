# Generator Module (PLACEHOLDER)

[![Go Reference][godoc-badge]][godoc-link]
[![Go Report Card][goreportcard-badge]][goreportcard-link]
[![codecov][codecov-badge]][codecov-link]

> [!WARNING]
> **PLACEHOLDER MODULE - NOT YET IMPLEMENTED**
>
> This module is a placeholder showing the planned architecture for common
> utilities for protoc plugin code generators. It will provide helpers for
> working with protocol buffer descriptors, managing code generation, and
> traversing descriptor trees.

## Overview

**PLANNED FUNCTIONALITY (NOT YET AVAILABLE):**

The generator module will provide essential utilities that all protoc plugins
need when generating code from protocol buffer descriptors. It will serve as a
foundation for building robust code generators with consistent patterns.

## Core Utilities

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

## Integration with Options Module

The options module relies heavily on these utilities:

1. **Type Checking**: For conditional option application
2. **Path Building**: For selector matching
3. **Traversal**: For applying hooks to matching elements
4. **Context**: Shared context for environment-aware options

Example usage in options module:

```go
// options module using generator utilities
func (r *Registry) applyHooks(file proto.Message) error {
    var hookErr error

    err := generator.ForEachMessage(file, func(msg proto.Message) bool {
        // Apply hooks to this message and its fields
        hookErr = r.applyMessageHooks(file, msg)
        return hookErr == nil // continue if no error
    })

    if hookErr != nil {
        return hookErr
    }
    return err
}

func (r *Registry) applyMessageHooks(file, msg proto.Message) error {
    msgPath := generator.GetFullName(msg)

    // Apply hook to message itself if it matches
    if MatchesSelector(r.selector, msgPath) {
        if err := r.runHooks(msg, r.GetMessageOptions(file, msg)); err != nil {
            return err
        }
    }

    // Apply hooks to each field
    var fieldErr error
    err := generator.ForEachField(msg, func(field proto.Message) bool {
        fieldPath := generator.GetFieldPath(file, msg, field)
        if MatchesSelector(r.selector, fieldPath) {
            opts := r.GetFieldOptions(file, msg, field)
            fieldErr = r.runHooks(field, opts)
            return fieldErr == nil // continue if no error
        }
        return true // continue to next field
    })

    if fieldErr != nil {
        return fieldErr
    }
    return err
}
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

---

**NOTE:** This entire module is currently a placeholder demonstrating the
intended architecture and planned features. Actual implementation will be added
in future releases as the protomcp.org ecosystem develops and the need for
these utilities becomes concrete.

[godoc-badge]: https://pkg.go.dev/badge/protomcp.org/common/generator.svg
[godoc-link]: https://pkg.go.dev/protomcp.org/common/generator
[goreportcard-badge]: https://goreportcard.com/badge/protomcp.org/common
[goreportcard-link]: https://goreportcard.com/report/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg?flag=generator
[codecov-link]: https://codecov.io/gh/protomcp/common?flag=generator
