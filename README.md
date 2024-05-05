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
+ Loops (`for` and `while` keywords)
+ First-class functions
+ If expressions
+ Prefix operators (binary expressions)
+ Infix operators (unary expressions)
+ Postfix operators

## Components

+ Token set
+ Lexer
+ Abstract Syntax Tree (AST)
+ Pratt parser based on context-free grammars and the Backus-Naur-Form

## TODOs

+ [ ] feat: Unicode
+ [ ] feat: parsing line, column number for better visibility
+ [ ] feat: hexical notation and octal notation for integers
+ [ ] feat: formatting and prettier in REPL
+ [ ] feat: support for multiple types: boolean, float, struct, string, byte, etc
+ [ ] feat: support for collection types: array, map, set
+ [ ] feat: add support for variadic functions
+ [ ] feat: add support for anonymous functions
+ [ ] feat: add <=, >= operators
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

## References

+ *Top Down Operator Precedence*, Vaughan Pratt


## License

[MIT License](./LICENSE)
