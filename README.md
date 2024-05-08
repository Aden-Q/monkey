# [WIP] Monkey

Monkey is an interpreted language written in Go. *This project is still under development.*

## Dependencies

+ `direnv`
+ `just` (not necessary)
+ Go 1.21+
+ ginkgo (if you want to run local unit tests)
+ golangci-lint (if you want to do local lint)

## Features

+ Programs can run in REPL or as scripts
+ Identifiers consist of alphabet letters or underscore
+ Statements are explicitly end by seminlon `;`
+ Comments start with double slash `//`
+ Syntax errors are caught in a single run
+ Conditional statements (`if` and `switch` keywords)
+ Loops (`for`, `range`, and `while` keywords)
+ First-class functions
+ If expressions
+ Prefix operators (binary expressions)
+ Infix operators (unary expressions)
+ Postfix operators
+ Tree-walking interpreter
+ Works on 64-bit machines
+ Not production ready yet, need a few more details to be addressed
+ Function that allows for closures

## Components

+ Token set
+ Lexer
+ Abstract Syntax Tree (AST)
+ Pratt parser based on context-free grammars and the Backus-Naur-Form
+ Tree-walking interpreter (evaluator), future: bytecode, VM, and JIT compilation

## TODOs

+ [ ] feat: Unicode
+ [ ] feat: parsing line, column number for better visibility
+ [ ] feat: hexical notation and octal notation for integers
+ [ ] feat: formatting and prettier in REPL
+ [ ] feat: support for multiple types: boolean, float, struct, string, byte, etc
+ [ ] feat: support for collection types: array, map, set
+ [ ] feat: add support for variadic functions
+ [ ] feat: add support for anonymous functions
+ [x] feat: add <=, >= operators
+ [ ] feat: add logical operators ||， &&
+ [ ] feat: add bitwise operators ^, |, &
+ [ ] refactor: unary operators, binary operators, ternary operators
+ [ ] feat: use Cobra to enable multiple modes when launching the REPL
+ [ ] feat: use quit(), exit(), or Ctrl-D to exit elegantly
+ [ ] feat: support for concurrency primitives such as Mutex, RWMutex, atomic
+ [ ] feat: support for comments
+ [ ] docs: a diagram for the full REPL loop including the AST used
+ [ ] check whether we need the token field in AST
+ [ ] test: increase test coverage to at least 80%
+ [ ] chore: dockerize it and publish an image to Docker Hub
+ [ ] feat: check how to trigger the GC during runtime (in REPL)
+ [ ] feat: return can only be used in functions, do not allow plain return in REPL
+ [ ] feat: if expression with multiple branches
+ [ ] feat: check peek token type in parseExpression
+ [ ] feat: check the difference between if expression and if statement
+ [ ] feat: empty statement with only a single ;
+ [ ] feat: scroll in command history with up and down keys
+ [ ] feat: PrettyPrint, color, AST, etc
+ [ ] feat: sys call such as print(...)
+ [ ] feat: add a helper for available functions
+ [ ] feat: consider how to write a GC
+ [ ] feat: switch between multiple value representation system using some flag
+ [ ] feat: class (object system)
+ [ ] feat: configuration with yaml or envrc
+ [ ] feat: edge cases for those operators
+ [ ] feat: integer division operator and float division operator
+ [ ] feat: reference integer literal as constant, simulate some static memory space for literals of integer, strings, etc.
+ [ ] feat: immutable types
+ [ ] feat: integer overflow problem
+ [ ] feat: command history and navigate in REPL using left, right, up, bottom
+ [ ] feat: configuration as env vars, default + direnv
+ [ ] feat: semantics for left and right arrows in REPL
+ [ ] feat: semantics for up and down arrows in REPL
+ [ ] feat: lexing, parsing, evaluation of nil expression statement
+ [ ] feat: slow startup
+ [ ] ci: deploy as a web app
+ [ ] fix: do not allow plain return in REPL outside of a function, report an error instead

## References

+ *Writing An Interpreter In Go, Thorsten Ball*
+ *Top Down Operator Precedence, Vaughan Pratt*
+ *The Structure and Interpretation of Computer Programs (SICP), Harold Abelson*

## License

[MIT License](./LICENSE)
