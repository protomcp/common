# AGENT.md

This file provides guidance to AI agents when working with code in this
repository. For developers and general project information, please refer to
[README.md](README.md) first.

## Repository Overview

`common` is a shared utilities library for protomcp.org projects. It provides
common testing helpers, error handling patterns, and utility functions that are
used across multiple repositories including `nanorpc` and `protomcp`.

## Prerequisites

Before starting development, ensure you have:

- Go 1.23 or later installed (check with `go version`).
- `make` command available (usually pre-installed on Unix systems).
- `pnpm` for JavaScript/TypeScript tooling (preferred over npm).
- Git configured for proper line endings.

## Common Development Commands

```bash
# Full build cycle (get deps, generate, tidy, build)
make all

# Run tests
make test

# Run tests with coverage
make test GOTEST_FLAGS="-cover"

# Run tests with verbose output and coverage
make test GOTEST_FLAGS="-v -cover"

# Generate coverage reports
make coverage

# Format code and tidy dependencies (run before committing)
make tidy

# Clean build artifacts
make clean

# Update dependencies
make up

# Run go:generate directives
make generate
```

## Build System Features

### Unified Build System

This project uses the same build system as other protomcp.org repositories:

- **Module**: `protomcp.org/common`
- **No submodules**: Single module structure for simplicity
- **Shared tooling**: Same linting and quality tools as siblings

### Tool Integration

The build system includes comprehensive tooling:

#### Linting and Quality

- **golangci-lint**: Go code linting with version selection.
- **revive**: Additional Go linting with custom rules.
- **markdownlint**: Markdown formatting and style checking.
- **shellcheck**: Shell script analysis.
- **cspell**: Spell checking for documentation and code.
- **languagetool**: Grammar checking for Markdown files.

#### Coverage and Testing

- **Coverage collection**: Automated test coverage reporting.
- **Test execution**: Standard Go testing with coverage support.

#### Development Tools

- **Whitespace fixing**: Automated trailing whitespace removal.
- **EOF handling**: Ensures files end with newlines.
- **Dynamic tool detection**: Tools auto-detected via pnpx.

### Configuration Files

Tool configurations are stored in `internal/build/`:

- `markdownlint.json`: Markdown linting rules (80-char lines)
- `cspell.json`: Spell checking dictionary and rules
- `languagetool.cfg`: Grammar checking configuration
- `revive.toml`: Go linting rules and thresholds

## Project Structure

### Core Components

- **Common utilities**: General-purpose helper functions
- **Test utilities**: Shared testing helpers and assertions
- **Error patterns**: Consistent error handling across projects

### Key Features

- **Minimal dependencies**: Keep external dependencies to a minimum
- **Interface-based**: Prefer interfaces for flexibility
- **Well-tested**: High test coverage for reliability
- **Documentation**: All public APIs must be documented

## Development Workflow

### MANDATORY: Test-Driven Development (TDD)

**ALL DEVELOPMENT MUST FOLLOW TDD**:

1. **Write failing tests first** - Define expected behaviour in tests
2. **Implement minimal code** - Write just enough to pass tests
3. **Refactor for quality** - Improve code while maintaining tests
4. **Repeat cycle** - Continue until feature is complete

**Test Infrastructure**:

- Focus on table-driven tests for comprehensive coverage
- Use subtests for better test organization
- Test utilities will be added as needed in future

### Before Starting Work

1. **Understand purpose**: This is a shared library - changes affect multiple
   projects.
2. **Check dependencies**: Avoid adding external dependencies unless absolutely
   necessary.
3. **Review existing code**: Follow established patterns and conventions.
4. **Write tests first**: Always start with failing tests.

### Code Quality Standards

The project enforces quality through:

- **Go standards**: Standard Go conventions and formatting
- **Field alignment**: Structs optimized for memory efficiency

  ```bash
  # Fix field alignment issues
  GOXTOOLS="golang.org/x/tools/go/analysis/passes"
  FA="$GOXTOOLS/fieldalignment/cmd/fieldalignment"
  go run "$FA@latest" -fix ./...
  ```

- **Linting rules**: Comprehensive linting via golangci-lint and revive
- **Test coverage**: Aim for high test coverage
- **Documentation**: All public APIs must be documented

## Testing Guidelines

### Test Structure

- **Table-driven tests**: Preferred for comprehensive coverage
- **Subtests**: Use `t.Run()` for better organization
- **Test helpers**: Use standard Go testing patterns
- **Isolation**: Tests should not depend on external resources

### Running Tests

```bash
# Run all tests
make test

# Run tests with race detection
make test GOTEST_FLAGS="-race"

# Run specific tests
make test GOTEST_FLAGS="-run TestSpecific"

# Generate coverage
make coverage
```

## Important Notes

### Design Principles

- **Simplicity**: Keep interfaces and implementations simple
- **Reusability**: Design for use across multiple projects
- **Stability**: Changes here affect multiple repositories
- **Performance**: Consider performance implications

### Development Environment

- Always use `pnpm` instead of `npm` for JavaScript/TypeScript tooling
- No protocol buffer files in this repository
- Focus on pure Go utilities and helpers

## Pre-commit Checklist

1. **ALWAYS run `make tidy` first** - Fix ALL issues before committing:
   - Go code formatting and whitespace clean-up
   - Markdown files checked with markdownlint and cspell
   - Shell scripts checked with shellcheck
2. **Verify all tests pass** with `make test`
3. **Check coverage** with `make coverage` if adding new code
4. **Update documentation** if changing public APIs

## Git Usage Guidelines

**CRITICAL**: Always follow these git practices to avoid accidental commits:

1. **NEVER use bulk operations** - Always explicitly specify files:

   ```bash
   # CORRECT - explicitly specify files
   git add file1.go file2.go
   git commit -s file1.go file2.go -m "commit message"

   # WRONG - bulk staging/committing
   git add .
   git add -A
   git add -u
   git commit -s -m "commit message"
   git commit -a -m "commit message"
   ```

2. **Use `-s` when doing commits** - Don't take credit for the work

3. **Check what you're committing**:

   ```bash
   git status --porcelain  # Check current state
   git diff --cached       # Review staged changes before committing
   ```

4. **Atomic commits** - Each commit should contain only related changes for a
   single purpose

## Common Patterns

### Error Handling

When this library includes error handling utilities:

```go
// Consistent error creation
return fmt.Errorf("failed to process: %w", err)

// Nil checks
if obj == nil {
    return errors.New("nil receiver")
}
```

### Testing Helpers

When creating test utilities:

```go
// Table-driven test helper
func TestSomething(t *testing.T) {
    tests := []struct {
        name string
        input string
        want string
    }{
        // test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## Troubleshooting

### Common Issues

1. **Import cycles**:
   - Keep dependencies minimal
   - Avoid importing from projects that depend on common

2. **Test failures**:
   - Run tests locally before pushing
   - Check for race conditions with `-race` flag

3. **Tool detection failures**:
   - Install tools globally with `pnpm install -g <tool>`
   - Check that pnpx is available and functional
   - Tools fall back to no-op if not found

### Getting Help

- Check existing code for patterns
- Review test files for usage examples
- Keep changes minimal and focused
- Test thoroughly before proposing changes

This project provides the foundation for other protomcp.org repositories,
so stability and reliability are paramount.
