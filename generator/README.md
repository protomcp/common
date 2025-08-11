# Generator Module

[![Go Reference][godoc-badge]][godoc-link]
[![Go Report Card][goreportcard-badge]][goreportcard-link]
[![codecov][codecov-badge]][codecov-link]

## Overview

The generator module provides essential utilities that all protoc plugins need
when generating code from protocol buffer descriptors. It serves as a foundation
for building robust code generators with consistent patterns.

## Status

### ✅ Implemented

- **Descriptor Type Checking** - Complete set of type checking and casting
  functions
- **Field Characteristic Checking** - Comprehensive field property detection

### 🚧 Planned

- Path and Naming Utilities
- Descriptor Traversal
- Visitor Pattern
- Type Resolution
- Context Management
- Dependency Analysis
- Code Generation Helpers

## Core Utilities

### Descriptor Type Checking

All type checking functions follow the `AsFoo`/`IsFoo` pattern where `AsFoo`
returns the typed descriptor and a boolean, while `IsFoo` is a convenience
wrapper that only returns the boolean.

#### Type Casting and Checking

```go
// Cast to specific descriptor types
func AsMessage(desc proto.Message) (*descriptorpb.DescriptorProto, bool)
func AsFieldType(desc proto.Message) (*descriptorpb.FieldDescriptorProto, bool)
func AsServiceType(desc proto.Message) (
    *descriptorpb.ServiceDescriptorProto, bool)
func AsMethodType(desc proto.Message) (
    *descriptorpb.MethodDescriptorProto, bool)
func AsEnumType(desc proto.Message) (*descriptorpb.EnumDescriptorProto, bool)
func AsFileType(desc proto.Message) (*descriptorpb.FileDescriptorProto, bool)

// Cast to message descriptor and verify name
func AsMessageType(desc proto.Message, name string) (
    *descriptorpb.DescriptorProto, bool)

// Boolean checks (wrappers around AsFoo functions)
func IsMessage(desc proto.Message) bool
func IsMessageType(desc proto.Message, name string) bool
func IsFieldType(desc proto.Message) bool
func IsServiceType(desc proto.Message) bool
func IsMethodType(desc proto.Message) bool
func IsEnumType(desc proto.Message) bool
func IsFileType(desc proto.Message) bool
```

#### Field Characteristics

```go
// Cast and check field characteristics
func AsRepeatedField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)
func AsMapField(field proto.Message) (*descriptorpb.FieldDescriptorProto, bool)
func AsOneOfField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)
func AsOptionalField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)
func AsRequiredField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)
func AsScalarField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)
func AsMessageField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)
func AsEnumField(field proto.Message) (
    *descriptorpb.FieldDescriptorProto, bool)

// Boolean checks (wrappers around AsFoo functions)
func IsRepeatedField(field proto.Message) bool
func IsMapField(field proto.Message) bool
func IsOneOfField(field proto.Message) bool
func IsOptionalField(field proto.Message) bool
func IsRequiredField(field proto.Message) bool
func IsScalarField(field proto.Message) bool
func IsMessageField(field proto.Message) bool
func IsEnumField(field proto.Message) bool
```

#### Usage Example

```go
// Using the AsFoo pattern for type-safe access
if field, ok := AsFieldType(desc); ok {
    // field is now typed as *descriptorpb.FieldDescriptorProto
    if repeatedField, ok := AsRepeatedField(field); ok {
        fmt.Printf("Field %s is repeated\n", repeatedField.GetName())
    }
}

// Simple boolean check when you don't need the typed object
if IsMessageType(desc, "MyMessage") {
    // Process message type
}

// Chaining checks for specific field types
if field, ok := AsFieldType(desc); ok {
    switch {
    case IsScalarField(field):
        // Handle scalar field
    case IsMessageField(field):
        // Handle message field
    case IsEnumField(field):
        // Handle enum field
    }
}
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

### Visitor Pattern

Used by options module for walking descriptor tree with hooks:

```go
// Visitor interface
type Visitor interface {
    VisitFile(file *descriptorpb.FileDescriptorProto) error
    VisitMessage(msg *descriptorpb.DescriptorProto) error
    VisitField(field *descriptorpb.FieldDescriptorProto) error
    VisitEnum(enum *descriptorpb.EnumDescriptorProto) error
    VisitEnumValue(value *descriptorpb.EnumValueDescriptorProto) error
    VisitService(svc *descriptorpb.ServiceDescriptorProto) error
    VisitMethod(method *descriptorpb.MethodDescriptorProto) error
}

// Walk descriptor tree
func Walk(file proto.Message, visitor Visitor) error
func WalkMessage(msg proto.Message, visitor Visitor) error

// Path-aware walking
type PathVisitor interface {
    Visit(path []string, desc proto.Message) error
}

func WalkWithPath(file proto.Message, visitor PathVisitor) error
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

### Context Management

Shared context for options and generators:

```go
// Build context
type Context struct {
    Target      string              // e.g., "embedded", "server", "wasm"
    Platform    string              // e.g., "linux", "darwin", "windows"
    BuildType   string              // e.g., "debug", "release", "profile"
    Features    map[string]bool     // Feature flags
    Variables   map[string]string   // Environment variables
    Metadata    any                 // Plugin-specific context
}

func NewContext() *Context
func (c *Context) WithTarget(target string) *Context
func (c *Context) WithPlatform(platform string) *Context
func (c *Context) WithBuildType(buildType string) *Context
func (c *Context) WithFeature(name string, enabled bool) *Context
func (c *Context) WithVariable(key, value string) *Context
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

## Code Generation Helpers

### Output Management

```go
// Code buffer with formatting
type CodeBuffer struct {
    // Internal buffer management
}

func NewCodeBuffer() *CodeBuffer
func (b *CodeBuffer) P(format string, args ...any)  // Print line
func (b *CodeBuffer) In()                            // Indent
func (b *CodeBuffer) Out()                           // Outdent
func (b *CodeBuffer) String() string                 // Get output

// File generation
type GeneratedFile struct {
    Name    string
    Content string
}

func NewGeneratedFile(name string) *GeneratedFile
```

### Comment Extraction

```go
// Comment handling
func GetLeadingComments(desc proto.Message) string
func GetTrailingComments(desc proto.Message) string
func FormatComments(comments string, prefix string) string
func ExtractTags(comments string) map[string]string
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

1. **Null Safety**: All functions handle nil inputs gracefully.
2. **Error Propagation**: Iteration functions stop on first error.
3. **Type Safety**: Use type assertions with checks for proto.Message types.
4. **Performance**: Cache frequently accessed paths and names.
5. **Compatibility**: Support both proto2 and proto3 descriptors.

## Future Extensions

- AST manipulation utilities
- Code formatting helpers
- Template engine integration
- Dependency graph visualization
- Performance profiling for large schemas
- Parallel processing support

---

## Implementation Status

**Phase 1 Complete**: Descriptor type checking and field characteristic
detection are fully implemented with 98.2% test coverage.

**Next Phases**: Additional functionality will be added incrementally as the
protomcp.org ecosystem develops.

[godoc-badge]: https://pkg.go.dev/badge/protomcp.org/common/generator.svg
[godoc-link]: https://pkg.go.dev/protomcp.org/common/generator
[goreportcard-badge]: https://goreportcard.com/badge/protomcp.org/common
[goreportcard-link]: https://goreportcard.com/report/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg?flag=generator
[codecov-link]: https://codecov.io/gh/protomcp/common?flag=generator
