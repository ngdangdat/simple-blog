# Simple Blog

A minimal blog builder written in Go that converts Markdown to HTML. Built as a learning project to explore Abstract Syntax Trees, parsing, and test-driven development.

## Features

- Markdown to HTML conversion
- Jekyll-inspired static site generation
- Test-driven development approach

## Getting Started

```bash
# Navigate to the builder tool
cd tools/builder

# Run the builder
go run .

# Run tests
go test ./...

# Build the tool
go build
```

## Architecture

The project uses a Go-based builder tool located in `tools/builder/` that parses Markdown files and generates static HTML output.

## Development

This project follows test-driven development practices and is designed to teach parsing concepts through practical implementation.