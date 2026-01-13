# AI Assistant Prompts

This document contains the key prompts and interactions used during the development of this solution.

## Initial Analysis Prompt

```
explain and analyze problem in file README.md
- create plan to resolve problem
- show at least 2 solutions, compare and explain why i should pick this
- make sure everything in README.md have to do
- the language i want is Golang
```

## Solution Design

Three solutions were proposed and compared:

### Solution 1: Sequential Stream Processing (Recommended)
- Single-pass file reading with buffered I/O
- O(unique campaigns) memory complexity
- Simple, reliable, and maintainable

### Solution 2: Concurrent Chunk Processing
- Parallel processing with goroutines
- Higher throughput on multi-core systems
- More complex chunk boundary handling

### Solution 3: Memory-Mapped File
- mmap-based file access
- Platform-specific considerations
- Overkill for 1GB file size

## Key Design Decisions

1. **Standard library only**: No external dependencies for maximum portability
2. **Streaming approach**: Memory-efficient processing for large files
3. **Buffered I/O**: 64KB buffer for optimal disk read performance
4. **Clear separation**: Models, aggregator, and output modules

## Implementation Notes

The AI assistant was used to:
1. Analyze the problem requirements
2. Design multiple solution approaches
3. Generate the initial codebase structure
4. Create unit tests
5. Write documentation
