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

## Current Progress

The developer has successfully completed the initial TDD learning phase:

### Completed
- Basic Go testing setup with table-driven tests
- Simple markdown parsing for headers (`#`, `##`, `###`) 
- Basic bold text parsing (`**bold**`)
- Discovered the limitations of string replacement approach

### Current Challenge
The developer has reached the natural breaking point where simple string replacement fails. The test case `"# test **bold heading**"` → `"<h1>test <b>bold heading</b></h1>"` reveals why AST and proper parsing are needed.

### Learning Progress on Tokenization
The developer has made significant conceptual progress understanding tokenization:

**Key Concepts Mastered:**
- **State tracking**: Using stacks to remember context (inside bold markup, etc.)
- **Lookahead**: Need to read ahead when encountering `#` to distinguish H1/H2/H3
- **Branching logic**: Specialized scanning functions for different patterns
- **Position management**: Functions must return where to continue scanning
- **Token emission timing**: Don't emit immediately, collect until pattern is clear

**Current Understanding:**
- Recognizes that tokenization goes: Raw Text → Tokens → AST  
- Understands the main scanning loop delegates to specialized functions
- Grasps that `scanHeading()`, `scanBold()` functions handle pattern recognition
- Knows functions must return both token type and next scan position

**Ready for Implementation:**
- Token type definitions (HASH, DOUBLE_HASH, TEXT, BOLD_START, etc.)
- Main tokenizer loop structure
- Specialized scanning functions with lookahead

### Next Steps
1. **Implement token types** and basic tokenizer structure
2. **Build scanning functions** for headers and bold text
3. **Write tokenizer tests** to verify token stream output
4. **Design AST structure** based on token stream

## Current State

Active development in `tools/builder/` with conceptual foundation for tokenization complete. Ready to implement the tokenizer with clear understanding of parsing principles.
