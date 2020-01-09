# Go Training

![](https://github.com/maxine-krebs/golang-training/workflows/Test/badge.svg)

## How to use this Training

Each module will have a `test/` folder. To verify your knowledge, please fork this repo and ensure that each `test/` passes CI.

## Organization

Each module has the following sections:
- Concepts
- Exercises
- Tips
- Further Reading

## Syllabus

[Module 1 - Mechanics](1-Mechanics/Readme.md)
- [Go Setup](1-Mechanics/1.1.md)
  - Local install
  - System validation
- [Hello World](1-Mechanics/1.2.md)
  - Line by line walkthrough
  - Compile and run
- Project Organization
  - [Packages](1-Mechanics/1.3.0.md)
  - [GOPATH](1-Mechanics/1.3.1.md)
  - [Package Management](1-Mechanics/1.3.2.md)

Module 2 - Basics
- [Types](2-Basics/2.1.md)
  - Basic types
  - Aliases
- [Variables](2-Basics/2.2.md)
  - Declaration and Assignment
  - Zero values
  - iota
- [Functions](2-Basics/2.3.md)
  - Arguments
  - Multi-valued returns
  - Anonymous functions
  - Closures
- [Control Flow](2-Basics/2.4.md)
  - if/else
  - for/break/continue
  - short statements
  - defer
  - panic/recover

Module 3 - Data Structures
- [Pointers](1-Mechanics/1.3.1.md)
  - Dereferencing
- [Structs](3-Data/3.2.md)
  - Definition
  - Declaration/Assignment
  - Fields
  - Nested Structures
  - Methods
- [Container Types](3-Data/3.3.md)
  - Slices
  - Lists
  - Maps
  - Looping

Module 4 - Design
- [Interfaces](4-Design/4.1.md)
- [Errors](4-Design/4.2.md)
- [Logging](4-Design/4.3.md)
- [Init Function](4-Design/4.4.md)
- [Testing](4-Design/4.5.md)

Module 5 - Concurrency
- [Goroutines](5-Concurrency/5.1.md)
- [Channels](5-Concurrency/5.2.md)
  - Unbuffered / Buffered
  - Directional Arguments
  - Select
  - Close
- [Sync](5-Concurrency/5.3.md)
  - Mutex
  - Once
  - WaitGroup

## Issues and Errors

Is something explained poorly? Or flat out wrong? [Let me know!](https://github.com/cgetzen/golang-training/issues/new)
