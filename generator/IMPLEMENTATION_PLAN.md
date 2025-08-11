# Generator Utilities Module Implementation Plan

## Executive Summary

This document outlines the implementation plan for the
`protomcp.org/common/generator` submodule, which provides helper utilities for
implementing protoc code generators. This module offers convenience functions
on top of standard protobuf packages to simplify the development of protoc
plugins. Phase 1 (descriptor type checking) is now complete with 98.2% test
coverage, following strict TDD principles.

## Context and Background

### Project Structure

- **Repository**: `protomcp.org/common`
- **Module**: `generator/` (submodule)
- **Status**: Phase 1 Complete (Descriptor Type Checking)
- **Purpose**: Provide helper utilities for implementing protoc code
  generators
- **Target Users**: Developers building protoc plugins and code generators

### Key Constraints

1. **British English**: Use GB English spelling throughout (e.g.,
   "organisation", "behaviour", "centre")
2. **No Directory Changes**: All operations must be executed from the current
   working directory
3. **Submodule Handling**:
   - Root `Makefile` uses `.tmp/gen.mk` for dynamic submodule targets
   - Use `go -C generator/` for Go commands (root and generator are separate
     modules)
   - Root module cannot directly import generator module code
4. **TDD Mandatory**: Write failing tests first, then implement minimal code
   to pass
5. **Testing Standards**: Follow MANDATORY requirements from
   `darvaza.org/core/TESTING.md`

### Dependencies

- `google.golang.org/protobuf/types/descriptorpb` - Protocol buffer descriptor
  types
- `google.golang.org/protobuf/proto` - Protocol buffer runtime
- `darvaza.org/core` - Testing utilities (test dependency only)

### Utilities from darvaza.org/core to Use

Based on analysis of `darvaza.org/core`, we should leverage these existing
utilities instead of reimplementing:

**Testing Infrastructure**:

- `TestCase` interface and `RunTestCases` - For table-driven/data-driven
  tests only
- All `Assert*` functions - For test assertions
- `MockT` - For testing our test utilities if needed
- Regular `t.Run()` - For non-data-driven test organisation and subtests

**Error Handling**:

- `ErrNilReceiver`, `ErrInvalid`, `ErrNotExists` - Standard errors
- `Wrap`, `Wrapf` - Error wrapping utilities
- `IsNil`, `IsZero` - Validation helpers
- `CompoundError` - For aggregating multiple errors

**Slice Utilities**:

- `S[T]()` - Slice creation helper
- `SliceContains`, `SliceContainsFn` - For checking slice membership
- `SliceCopy`, `SliceCopyFn` - For copying slices
- `SliceMap` - For transforming slices

**Generic Utilities**:

- `Coalesce` - For finding first non-zero value
- `IIf` - Ternary operator
- `Must`, `MustOK` - For panic on error
- `Maybe`, `MaybeOK` - For ignoring errors

**Context Management**:

- `ContextKey[T]` - Type-safe context keys (if needed for build context)

**Note**: We do NOT need to reimplement any string case conversion
(ToSnakeCase, ToCamelCase, etc.) if they don't exist in core - we'll implement
only what's needed for protobuf naming conventions.

## Implementation Phases

### Phase 1: Foundation - Descriptor Type Checking

**Files**: `descriptor.go`, `descriptor_test.go`

**Functionality**:

- Type identification functions (`IsMessageType`, `IsFieldType`,
  `IsServiceType`, etc.)
- Field characteristic checks (`IsRepeatedField`, `IsMapField`,
  `IsOneOfField`, etc.)

**Test Approach**:

- Create test cases with mock descriptors
- Test each type checking function with valid and invalid inputs
- Use table-driven tests with `TestCase` interface

### Phase 2: Path and Naming Utilities

**Files**: `naming.go`, `naming_test.go`

**Functionality**:

- Path construction (`GetFullName`, `GetFieldPath`, `GetQualifiedName`)
- Name conversion (`ToGoName`, `ToCamelCase`, `ToSnakeCase`, `ToKebabCase`)
- Package and namespace handling

**Test Approach**:

- Test various naming conventions and edge cases
- Include tests for empty strings, special characters
- Verify correct path construction for nested types

### Phase 3: Descriptor Traversal

**Files**: `traversal.go`, `traversal_test.go`

**Functionality**:

- Find functions (`FindMessage`, `FindField`, `FindEnum`, `FindService`)
- ForEach iterators following Go 1.23 convention (return true to continue)
- Nested traversal support

**Test Approach**:

- Create complex descriptor hierarchies for testing
- Test early termination (return false)
- Verify correct traversal order

### Phase 4: Visitor Pattern

**Files**: `visitor.go`, `visitor_test.go`

**Functionality**:

- `Visitor` interface definition
- `Walk` and `WalkMessage` functions
- `PathVisitor` for path-aware traversal
- Error propagation and handling

**Test Approach**:

- Implement mock visitors for testing
- Test visitor call order and error handling
- Verify path construction in PathVisitor

### Phase 5: Type Resolution

**Files**: `types.go`, `types_test.go`

**Functionality**:

- Type information extraction (`GetFieldType`, `GetFieldTypeName`)
- Type classification (`IsScalarType`, `IsMessageType`, `IsEnumType`)
- Type resolution (`ResolveTypeName`, `GetTypeDescriptor`)

**Test Approach**:

- Test all protobuf type variants
- Include tests for custom message types
- Verify correct type resolution across files

### Phase 6: Context Management

**Files**: `context.go`, `context_test.go`

**Functionality**:

- `Context` struct with builder pattern
- Target, platform, build type configuration
- Feature flags and environment variables
- Metadata storage for plugin-specific data

**Test Approach**:

- Test builder pattern methods
- Verify immutability where appropriate
- Test context merging and overrides

### Phase 7: Dependency Analysis

**Files**: `dependencies.go`, `dependencies_test.go`

**Functionality**:

- Dependency tracking (`GetFileDependencies`, `GetMessageDependencies`)
- Import management (`GetRequiredImports`, `ResolveImportPath`)
- Circular dependency detection

**Test Approach**:

- Create test files with various import patterns
- Test circular dependency scenarios
- Verify correct import resolution

### Phase 8: Code Generation Helpers

**Files**: `codegen.go`, `codegen_test.go`

**Functionality**:

- `CodeBuffer` with indentation management
- `GeneratedFile` structure
- Comment extraction and formatting
- Tag parsing from comments

**Test Approach**:

- Test indentation with nested structures
- Verify comment formatting preserves content
- Test tag extraction with various formats

### Phase 9: Integration and Documentation

**Files**: `README.md`, `doc.go`

**Tasks**:

- Remove placeholder warnings from README.md
- Create comprehensive doc.go with package documentation
- Add integration examples
- Update badges and links

## Testing Requirements

### Mandatory Compliance (from TESTING.md)

1. **TestCase Interface Validations**: Add `var _ TestCase = ...` for all
   test case types that implement data-driven tests
2. **Factory Functions**: Create `newXxxTestCase()` for all test case types
   - Use semantic factory variants to reduce parameters (e.g.,
     `newIsServiceTypeTrue()` and `newIsServiceTypeFalse()`)
   - Factory functions allow logical parameter order while structs maintain
     field alignment for memory efficiency
3. **Factory Usage**: Use factory functions, no naked struct literals
4. **RunTestCases Usage**: Use `core.RunTestCases(t, cases)` for data-driven
   tests (not required for all tests, only table-driven ones)
5. **Anonymous Functions**: No `t.Run()` anonymous functions >3 lines
6. **Test Case List Factories**: Complex test lists use factory functions

### Key Testing Patterns

**Factory Variants**: Create multiple factory functions for the same test case
type to reduce complexity:

- Base factory with all parameters
- Semantic variants that encode common patterns (e.g., success/failure cases)
- Reduces boolean parameter confusion
- Makes test intent clearer

**Example**:

```go
// Base factory
func newDescriptorTypeTestCase(name string, desc proto.Message,
    expected bool, checkFn func(proto.Message) bool) descriptorTypeTestCase

// Semantic variants for clearer intent
func newIsServiceTypeTrue(name string,
    desc proto.Message) descriptorTypeTestCase
func newIsServiceTypeFalse(name string,
    desc proto.Message) descriptorTypeTestCase
```

### Coverage Goals

- Minimum 80% code coverage
- 100% coverage for critical paths
- All error paths tested

## Build and Quality Checks

### Local Development Commands

```bash
# Run tests for generator submodule
make -C generator/ test

# Run tests with coverage
make -C generator/ test GOTEST_FLAGS="-cover"

# Generate coverage report
make -C generator/ coverage

# Run linting and formatting
make tidy

# Build the module
make -C generator/ all
```

### Pre-commit Checklist

- [ ] All tests passing
- [ ] Coverage meets requirements
- [ ] `make tidy` runs clean
- [ ] No linting errors
- [ ] Documentation updated
- [ ] Examples working

## Success Criteria

1. **Functional Requirements**
   - All planned utilities implemented
   - Compatible with standard protobuf packages
   - Usable by options module and other protoc plugin implementations
   - Simplifies common generator implementation tasks

2. **Quality Requirements**
   - Cognitive complexity ≤7
   - Cyclomatic complexity ≤10
   - Test coverage ≥80%
   - All linting checks pass

3. **Documentation Requirements**
   - Comprehensive godoc comments
   - Working examples in documentation
   - README updated without placeholders

## Handover Notes

### Current State (2025-08-11)

- Module structure exists and is functional
- README.md updated with implemented functionality
- Build system integrated via parent Makefile
- **Phase 1 Status**: ✅ COMPLETE
  - `descriptor_test.go`: Comprehensive tests with refactored structure
  - `descriptor.go`: Full implementation with AsFoo/IsFoo pattern
  - All tests passing with 100% coverage
  - Refactored to use AsFoo pattern for type-safe casting
  - Added AsMessage/IsMessage for generic message type checking
- **Test Utilities**: ✅ COMPLETE
  - `testutils.go`: Comprehensive test helper functions
  - `testutils_test.go`: Full test coverage
  - Added minimal constructors (NewFieldWithType, NewFieldWithLabel, etc.)
  - 100% test coverage achieved
- **Test Refactoring**: ✅ COMPLETE
  - All descriptor tests refactored to use factory pattern
  - Eliminated 40+ manual Enum() calls
  - Improved test organisation with factory functions
  - Documentation converted to tables for better readability

### Phase 1 Completion Summary

**Work Completed**:

- Implemented 28 functions (14 AsFoo, 14 IsFoo) for type checking
- Refactored test infrastructure reducing duplication by ~60%
- Achieved 98.2% test coverage with all tests passing
- Updated documentation removing placeholder status
- Fixed all linting and formatting issues

**Technical Decisions**:

- **AsFoo/IsFoo Pattern**: Type-safe casting and checking combined
- **Map Field Detection**: Simplified heuristic for "Entry" type names
- **Test Refactoring**: Consolidated test types using generic boolCheckTestCase

**Commit Steps for Phase 1**:

1. **Create commit message using Write tool**:
   - Use Write tool to create `.commit-msg` in current directory
   - Follow strict GB English spelling
   - End all sentences with periods except titles

2. **Ensure code quality**:

   ```bash
   make tidy
   make coverage-generator  # Output in .tmp/coverage/
   ```

3. **Commit files individually** (NEVER bulk stage):

   ```bash
   git commit -sF .commit-msg generator/descriptor.go \
       generator/descriptor_test.go
   git commit -sF .commit-msg generator/README.md
   git commit -sF .commit-msg generator/go.mod generator/go.sum
   git commit -sF .commit-msg internal/build/cspell.json
   ```

4. **Clean up**:

   ```bash
   rm .commit-msg
   ```

**Note**: IMPLEMENTATION_PLAN.md is NOT committed (internal tracking only).

**Next Steps (Phase 3 - Descriptor Traversal)**:

1. Implement Find functions for locating descriptors:
   - FindMessage(file, name) - Find message by name in file
   - FindField(msg, name) - Find field by name in message
   - FindEnum(file, name) - Find enum by name in file
   - FindService(file, name) - Find service by name in file
   - FindMethod(service, name) - Find method by name in service

2. Implement ForEach iterators (Go 1.23 convention):
   - ForEachField(msg, func) - Iterate over fields
   - ForEachMessage(file, func) - Iterate over messages
   - ForEachEnum(file, func) - Iterate over enums
   - ForEachService(file, func) - Iterate over services
   - ForEachMethod(service, func) - Iterate over methods
   - ForEachNestedMessage(msg, func) - Iterate nested messages
   - ForEachNestedEnum(msg, func) - Iterate nested enums

3. Follow established patterns:
   - Use AsFoo helpers from Phase 1 for type checking
   - Write comprehensive tests first (TDD)
   - Use testutils helpers for test descriptor creation
   - Maintain 100% test coverage
   - Return true to continue, false to stop iteration

### Key Files to Review

- `generator/README.md` - Architectural blueprint
- `../AGENT.md` - Development guidelines
- `../../darvaza.org/core/TESTING.md` - Testing requirements
- `.tmp/gen.mk` - Generated makefile for submodule

### Important Reminders

- Never use `cd` or change directories
- Always use `make -C generator/` for submodule operations
- Run `make tidy` before any commits
- Use explicit git adds, never `git add .`
- Follow British English spelling conventions
- Keep dependencies minimal

## Risk Mitigation

### Technical Risks

1. **Descriptor API Complexity**: Mitigate with comprehensive testing
2. **Performance with Large Schemas**: Add benchmarks in later phase
3. **Proto2 vs Proto3 Differences**: Test both variants

### Process Risks

1. **Scope Creep**: Stick to documented functionality only
2. **Breaking Changes**: Design for stability from start
3. **Integration Issues**: Test with real protoc plugins early

## Timeline Estimate

Based on complexity and testing requirements:

- Phase 1-2: 2-3 hours (Foundation and naming)
- Phase 3-4: 3-4 hours (Traversal and visitor)
- Phase 5-6: 2-3 hours (Types and context)
- Phase 7-8: 3-4 hours (Dependencies and codegen)
- Phase 9: 1-2 hours (Integration and documentation)

**Total Estimate**: 11-16 hours of focused development

## Appendix: Example Usage

```go
// Example: Using generator utilities to implement a protoc plugin
package main

import (
    "protomcp.org/common/generator"
    "google.golang.org/protobuf/types/descriptorpb"
    "google.golang.org/protobuf/types/pluginpb"
)

// MyGenerator implements a custom code generator using the utilities
type MyGenerator struct {
    buffer *generator.CodeBuffer
}

func (g *MyGenerator) Generate(req *pluginpb.CodeGeneratorRequest) (
    *pluginpb.CodeGeneratorResponse, error) {
    response := &pluginpb.CodeGeneratorResponse{}

    for _, file := range req.ProtoFile {
        // Use traversal utilities to process messages
        err := generator.ForEachMessage(file, func(msg proto.Message) bool {
            if generator.IsMessageType(msg, "MyMessage") {
                // Use naming utilities
                goName := generator.ToGoName(generator.GetMessageName(msg))

                // Use code buffer for output
                g.buffer = generator.NewCodeBuffer()
                g.buffer.P("type %s struct {", goName)
                g.buffer.In()

                // Process fields using traversal
                generator.ForEachField(msg, func(field proto.Message) bool {
                    g.generateField(field)
                    return true
                })

                g.buffer.Out()
                g.buffer.P("}")
            }
            return true // continue iteration
        })

        if err != nil {
            return nil, err
        }
    }

    return response, nil
}

func (g *MyGenerator) generateField(field proto.Message) {
    // Use type resolution utilities
    fieldType := generator.GetFieldTypeName(field)
    fieldName := generator.ToGoName(generator.GetFieldName(field))

    // Use type checking utilities
    if generator.IsRepeatedField(field) {
        g.buffer.P("%s []%s", fieldName, fieldType)
    } else {
        g.buffer.P("%s %s", fieldName, fieldType)
    }
}
```

---

**Document Version**: 1.0
**Created**: 2025-08-10
**Author**: AI Assistant
**Status**: Ready for Implementation
