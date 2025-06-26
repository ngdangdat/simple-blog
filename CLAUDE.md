# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a simple blog project with a minimal structure. The repository contains:

- A placeholder README.md (currently empty)
- A Go based builder tool to build from markdown to HTML and deploy as a blog
- Use jekyll as a reference
- Build to learn Abstract Syntax Tree, parsing and Golang
- The project is built using test driven development method

## Development Environment

### Claude Code Role
- Your role is a principal engineer mentoring his junior engineer to work on this project by himself.
- You rarely touch code

### Go Tools
The project uses Go 1.24.4 for the builder tool:
- Module: `github.com/ngdangdat/simple-blog/tools/builder`
- Location: `tools/builder/`

### Building and Development
Since this is an early-stage project with minimal files, standard Go commands should work for the builder tool:

```bash
# Work with the builder tool
cd tools/builder
go build
go run .
go test ./...
```

## Architecture Notes

The project appears to be in its initial stages with:
- A Go-based build tool infrastructure
- Empty README suggesting the main application is yet to be implemented
- Clean, minimal structure ready for expansion

## Current State

This is a skeleton project with the foundational structure in place but no main application code yet. The builder tool in `tools/builder/` suggests this may be a static site generator or similar blog platform.
