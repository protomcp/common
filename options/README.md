# Options Module (PLACEHOLDER)

<!-- cspell:ignore protoc GOOS arduino zstd -->

[![Go Reference][godoc-badge]][godoc-link]
[![Go Report Card][goreportcard-badge]][goreportcard-link]
[![codecov][codecov-badge]][codecov-link]

> [!WARNING]
> **PLACEHOLDER MODULE - NOT YET IMPLEMENTED**
>
> This module is a placeholder showing the planned architecture for a flexible
> override system for protobuf options, allowing plugins to modify option values
> from compiled proto descriptors through multiple configuration layers.

## Philosophy

Protocol buffer files define options at various levels (file, message, field,
service, method) through the protobuf options mechanism. These options are
compiled into descriptors and passed to plugins. This module provides a way to
override those compiled option values without modifying the original `.proto`
files, enabling environment-specific, target-specific, and use-case-specific
code generation.

## Core Concept

The module operates on the principle of **option override layers**:

1. **Base Layer**: Options as defined in `.proto` files and compiled by protoc.
2. **Override Layers**: Multiple sources that can modify these base values.
3. **Resolution**: Final values computed by applying overrides in priority
   order.

## How It Works

### Option Discovery

When a plugin receives a `CodeGeneratorRequest`, it contains compiled
descriptors with all options already resolved by protoc:

```go
func (p *Plugin) Generate(req *plugin.CodeGeneratorRequest) {
    for _, file := range req.ProtoFile {
        // File options from the descriptor
        fileOpts := file.GetOptions()

        // Message options
        for _, msg := range file.MessageType {
            msgOpts := msg.GetOptions()

            // Field options
            for _, field := range msg.Field {
                fieldOpts := field.GetOptions()
            }
        }

        // Service options
        for _, svc := range file.Service {
            svcOpts := svc.GetOptions()

            // Method options
            for _, method := range svc.Method {
                methodOpts := method.GetOptions()
            }
        }
    }
}
```

### Override Sources

The module allows overriding these compiled options through multiple sources:

```go
type OverrideSource interface {
    // Get override value for a specific option on a specific element
    GetOverride(
        element proto.Message,    // The descriptor (file, message, field, etc.)
        option *proto.ExtensionDesc, // The option to override
    ) (value any, found bool)

    // Priority determines override order (higher wins)
    Priority() int
}
```

## Override Hierarchy

Overrides are applied in the following priority order (lowest to highest):

### 1. Compiled Proto Options (Priority: 0)

The base values from the compiled `.proto` files:

```protobuf
message User {
    option (my_opt.message_opt) = "original";

    string id = 1 [(my_opt.field_opt) = "base"];
}
```

### 2. Configuration Files (Priority: 100)

External configuration files that override options by pattern:

```yaml
# overrides.yaml
overrides:
  - selector: "User"           # Message name
    options:
      my_opt.message_opt: "configured"

  - selector: "User.id"        # Field path
    options:
      my_opt.field_opt: "overridden"

  - selector: "*.timestamp"    # Pattern matching
    options:
      my_opt.field_opt: "timestamp_special"
```

### 3. Environment Variables (Priority: 200)

Environment-based overrides for deployment flexibility:

```bash
# Override specific option values
export PROTO_OPT_my_opt_message_opt="from_env"
export PROTO_OPT_User_id_my_opt_field_opt="env_override"
```

### 4. Options Files (Priority: 300)

Dedicated `.options` files alongside `.proto` files:

```text
# user.options
User my_opt.message_opt:"from_options_file"
User.id my_opt.field_opt:"options_file_override"
*.timestamp my_opt.field_opt:"timestamp_handling"
```

### 5. Command-Line Arguments (Priority: 400)

Direct overrides via protoc plugin parameters:

```bash
protoc --my-plugin_out=\
my_opt.message_opt=override,\
User.id.my_opt.field_opt=direct:\
. user.proto
```

### 6. Context-Aware Overrides (Priority: 500)

Dynamic overrides based on build context:

```go
overrides.RegisterContextual(
    func(ctx *Context, element proto.Message) map[string]any {
    if ctx.Target == "embedded" && generator.IsMessageType(element, "User") {
        return map[string]any{
            "my_opt.buffer_size": 256,
        }
    }
    return nil
})
```

## Implementation

### Helper Utilities

The module provides option-specific helpers, while general descriptor utilities
are available in `protomcp.org/common/generator`:

```go
// Option-specific helpers
func HasOption(desc proto.Message, ext *proto.ExtensionDesc) bool
func GetOptionValue[T any](
    desc proto.Message, ext *proto.ExtensionDesc) (T, bool)
func GetEffectiveOption[T any](
    registry *Registry, desc proto.Message, ext *proto.ExtensionDesc) (T, bool)

// Option override helpers
func ApplyOverride(
    desc proto.Message, ext *proto.ExtensionDesc, value any) error
func ClearOverride(desc proto.Message, ext *proto.ExtensionDesc) error
func GetOverrideSource(
    registry *Registry, desc proto.Message, ext *proto.ExtensionDesc) string

// Pattern matching for option selectors
func MatchesSelector(selector string, path string) bool
func CompileSelector(selector string) (*Selector, error)
func SelectorFromPath(elements ...string) string
```

Note: General descriptor utilities like `IsMessageType`, `GetFieldPath`, and
`ForEachField` are provided by `protomcp.org/common/generator`.

### Basic Usage

```go
import (
    "protomcp.org/common/generator"
    "protomcp.org/common/options"
    "google.golang.org/protobuf/proto"
)

func (p *Plugin) Generate(req *plugin.CodeGeneratorRequest) {
    // Create override registry
    registry := options.NewRegistry()

    // Add override sources
    registry.AddSource(options.NewYAMLSource("overrides.yaml"))
    registry.AddSource(options.NewEnvSource("PROTO_OPT_"))
    registry.AddSource(options.NewOptionsFileSource())
    registry.AddSource(options.NewArgsSource(req.GetParameter()))

    // Process files with overrides
    for _, file := range req.ProtoFile {
        // Get effective options for the file
        fileOpts := registry.GetFileOptions(file)

        for _, msg := range file.MessageType {
            // Get effective options for the message
            msgOpts := registry.GetMessageOptions(file, msg)

            for _, field := range msg.Field {
                // Get effective options for the field
                fieldOpts := registry.GetFieldOptions(file, msg, field)

                // Use the overridden values
                maxSize := options.GetExtension(fieldOpts, my_opt.E_MaxSize)
                if maxSize != nil {
                    // Generate code based on overridden max_size
                }
            }
        }
    }
}
```

### Accessing Overridden Options

```go
// Helper to get option value with overrides applied
func (r *Registry) GetOptionValue(
    descriptor proto.Message,
    extension *proto.ExtensionDesc,
) any {
    // Start with base value from compiled proto
    baseValue := proto.GetExtension(descriptor.GetOptions(), extension)

    // Apply overrides in priority order
    for _, source := range r.sources {
        if override, found := source.GetOverride(descriptor, extension); found {
            baseValue = override
        }
    }

    return baseValue
}
```

### Pattern Matching

The system supports flexible pattern matching for selectors:

```go
type Selector interface {
    Matches(path ElementPath) bool
}

// Examples:
// "User" - matches User message
// "User.id" - matches id field in User message
// "*.timestamp" - matches any field named timestamp
// "**.id" - matches id field at any nesting level
// "User.*" - matches all fields in User message
// "package.*.Message" - matches Message in any sub-package
```

## Configuration File Format

### YAML Format

```yaml
# overrides.yaml
version: "1.0"

# Global context
context:
  target: embedded
  platform: linux

# Override rules
overrides:
  # Message-level override
  - selector: "MyMessage"
    options:
      deprecated: true
      my_opt.generate: false

  # Field-level override with context
  - selector: "MyMessage.data"
    when:
      target: embedded
    options:
      my_opt.max_size: 256

  # Pattern-based override
  - selector: "*.id"
    options:
      my_opt.field_type: "fixed64"

  # Package-wide override
  - selector: "my_package.**"
    options:
      java_package: "com.example.my_package"
```

### Options File Format

```text
# myfile.options
# Simple key-value format for direct overrides

# Message options
MyMessage deprecated:true
MyMessage my_opt.validate:true

# Field options
MyMessage.id my_opt.field_type:fixed64
MyMessage.data my_opt.max_size:1024

# Pattern matching
*.timestamp my_opt.timestamp_format:"RFC3339"
```

## Naming Conventions

Following protobuf conventions:

- **Proto Options**: Use `underscore_case` (e.g., `my_opt.field_type`)
- **Plugin Names**: Use `kebab-case` for command-line (e.g., `my-plugin`)
- **Package Names**: Use `underscore_case` or dots
  (e.g., `my_package.sub_package`)
- **Selectors**: Use dots for hierarchy (e.g., `Message.field.subfield`)
- **Environment Variables**: Use `UPPER_SNAKE_CASE`
  (e.g., `PROTO_OPT_MY_OPTION`)

## Advanced Features

### Option Hooks and Transformations

Apply transformations to nodes based on their options:

```go
// Register a hook that runs on every element with certain options
registry.RegisterHook(func(element proto.Message, options proto.Message) error {
    // Check if element has a specific option set
    transform, ok := proto.GetExtension(options, my_opt.E_Transform).(bool)
    if ok && transform {
        // Apply transformation to the element
        applyTransformation(element)
    }
    return nil
})

// Register a typed hook for specific descriptor types
// cspell:ignore descriptorpb
registry.RegisterFieldHook(func(
    file *descriptorpb.FileDescriptorProto,
    msg *descriptorpb.DescriptorProto,
    field *descriptorpb.FieldDescriptorProto,
) error {
    opts := registry.GetFieldOptions(file, msg, field)

    // Transform based on option values
    if sizeVal := proto.GetExtension(opts, my_opt.E_MaxSize); sizeVal != nil {
        if size, ok := sizeVal.(int32); ok && size > 0 && size <= 256 {
            // Modify field based on max_size option
            fieldType := descriptorpb.FieldDescriptorProto_TYPE_FIXED32
            field.Type = &fieldType
        }
    }

    return nil
})

// Register a visitor pattern for traversing with options
registry.Walk(func(
    path ElementPath, element proto.Message, opts proto.Message) error {
    // path provides context: ["file.proto", "MyMessage", "my_field"]
    // element is the descriptor being visited
    // opts are the effective options after overrides

    switch descriptor := element.(type) {
    case *descriptorpb.FieldDescriptorProto:
        lazyVal := proto.GetExtension(opts, my_opt.E_Lazy)
        if lazyVal != nil {
            if lazy, ok := lazyVal.(bool); ok && lazy {
                // Mark field for lazy loading
                markLazy(path, descriptor)
            }
        }
    }

    return nil
})

// Batch transformations for efficiency
registry.RegisterBatchHook(func(elements []ElementWithOptions) error {
    // Process multiple elements together
    for _, element := range elements {
        if needsIndexing(element.Options) {
            addToIndex(element.Element, element.Path)
        }
    }
    return generateIndex()
})
```

### Conditional Hooks

Apply hooks based on conditions:

```go
// Only run in certain contexts
registry.RegisterConditionalHook(
    func(ctx *Context) bool {
        return ctx.Target == "database"
    },
    func(element proto.Message, options proto.Message) error {
        // This hook only runs when targeting database generation
        if proto.GetExtension(options, my_opt.E_Indexed) == true {
            generateDatabaseIndex(element)
        }
        return nil
    },
)

// Chain multiple conditions
registry.RegisterHook(
    WithCondition(HasOption(my_opt.E_Validate)).
    And(TargetIs("production")).
    Do(func(element proto.Message, options proto.Message) error {
        generateValidation(element)
        return nil
    }),
)
```

### Option Validation

Validate that overrides are compatible with option definitions:

```go
registry.AddValidator(func(
    descriptor proto.Message,
    option *proto.ExtensionDesc,
    value any,
) error {
    // Validate type compatibility
    if !options.IsCompatibleType(option, value) {
        return fmt.Errorf("incompatible type for %s", option.Name)
    }

    // Custom validation logic
    if option == my_opt.E_MaxSize {
        if size := value.(int32); size < 0 {
            return fmt.Errorf("max_size must be non-negative")
        }
    }

    return nil
})
```

### Option Transformation

Transform option values during resolution:

```go
registry.AddTransformer(func(
    descriptor proto.Message,
    option *proto.ExtensionDesc,
    value any,
) any {
    // Example: expand environment variables in string options
    if str, ok := value.(string); ok {
        return os.ExpandEnv(str)
    }
    return value
})
```

### Debugging Support

```go
// Enable debug logging
registry.EnableDebug()

// Trace option resolution
trace := registry.TraceOption(descriptor, my_opt.E_FieldOpt)
for _, step := range trace {
    log.Printf("Source: %s, Value: %v, Applied: %v",
        step.Source, step.Value, step.Applied)
}

// Export effective options
effective := registry.ExportEffective()
data, _ := yaml.Marshal(effective)
os.WriteFile("effective-options.yaml", data, 0644)
```

## Best Practices

1. **Use Patterns Wisely**: Broad patterns like `*` can have unintended effects.
   Be specific when possible.

2. **Document Overrides**: Include comments explaining why overrides are needed.

3. **Validate Early**: Validate overrides during plugin initialization to catch
   errors before code generation.

4. **Layer Appropriately**: Use the right override level for your use case:
   - Configuration files for project-wide settings
   - Environment variables for deployment variations
   - Command-line for one-off changes

5. **Test Override Combinations**: Test that your overrides work correctly in
   different contexts.

## Examples

### Target-Specific Overrides

```yaml
# Embedded target gets smaller sizes
overrides:
  - selector: "*.data"
    when:
      target: embedded
    options:
      my_opt.max_size: 256

  - selector: "*.data"
    when:
      target: server
    options:
      my_opt.max_size: 65536
```

### Development vs Production

```yaml
# Development enables validation
overrides:
  - selector: "**"
    when:
      build_type: debug
    options:
      my_opt.validate: true
      my_opt.logging: verbose

  - selector: "**"
    when:
      build_type: release
    options:
      my_opt.validate: false
      my_opt.logging: errors_only
```

## Compatibility

The options module is designed to work with:

- Standard protobuf options mechanism
- Custom options via extensions
- Any protoc plugin that processes descriptors
- Proto2 and Proto3 files

## Future Enhancements

- Remote configuration sources (etcd, consul, etc.)
- Hot-reloading of override configurations
- GUI for managing override rules
- Automatic conflict detection and resolution
- Override impact analysis tools

---

**NOTE:** This entire module is currently a placeholder demonstrating the
intended architecture and planned features. Actual implementation will be added
in future releases as the protomcp.org ecosystem develops and concrete
requirements for option override functionality emerge.

[godoc-badge]: https://pkg.go.dev/badge/protomcp.org/common/options.svg
[godoc-link]: https://pkg.go.dev/protomcp.org/common/options
[goreportcard-badge]: https://goreportcard.com/badge/protomcp.org/common
[goreportcard-link]: https://goreportcard.com/report/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg?flag=options
[codecov-link]: https://codecov.io/gh/protomcp/common?flag=options
