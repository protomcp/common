# Common

[![Go Reference][godoc-badge]][godoc-link]
[![codecov][codecov-badge]][codecov-link]

Common provides shared utilities and testing helpers for protomcp.org projects.

## Overview

The `common` package serves as a foundation for other protomcp.org repositories,
providing reusable components for testing, error handling, and general utilities
that are shared between `nanorpc`, `protomcp`, and related projects.

## Features

- **Test Utilities**: Common testing helpers and assertions
- **Shared Types**: Reusable data structures and interfaces
- **Error Handling**: Consistent error patterns across projects
- **Utility Functions**: General-purpose helper functions

## Usage

This package is designed to be imported by other protomcp.org projects:

```go
import (
    "protomcp.org/common"
    "protomcp.org/common/testutils"
)
```

## Development

For development guidelines, please refer to [AGENT.md](AGENT.md).

## License

See [LICENCE.txt](LICENCE.txt) for licensing information.

[godoc-badge]: https://pkg.go.dev/badge/protomcp.org/common.svg
[godoc-link]: https://pkg.go.dev/protomcp.org/common
[codecov-badge]: https://codecov.io/gh/protomcp/common/graph/badge.svg
[codecov-link]: https://codecov.io/gh/protomcp/common
