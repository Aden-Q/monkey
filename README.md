# [WIP] Monkey

Monkey is an interpreted language written in Go. *This project is still under development.*

## Dependencies

+ `direnv`
+ `just` (not necessary)
+ Go 1.21+
+ ginkgo (if you want to run local unit tests)
+ golangci-lint (if you want to do local lint)

## Features

+ Programs can be ran in REPL or as scripts
+ Identifiers that consist of alphabet letters or underscore
+ Able to catch all syntax errors in a single run
+ Able to evaluate arithmetic expressions
+ Able to evaluate conditional statements (`if` and `select` keywords)
+ Able to evaluate loops (`for` and `while` keywords)

## Components

+ Token set
+ Lexer
+ Abstract Syntax Tree (AST)
+ Pratt parser

## TODOs

- [ ] feat: Unicode
- [ ] feat: parsing line, column number for better visibility
- [ ] feat: hexical notation and octal notation for integers
- [ ] feat: formatting and prettier in REPL
- [ ] feat: support for multiple types: boolean, float, struct, string, byte, etc
- [ ] feat: support for collection types: array, map, set
- [ ] feat: add support for variadic functions
- [ ] feat: add support for anonymous functions
- [ ] feat: add <=, >= operators
- [ ] feat: add logical operators ||， &&
- [ ] feat: add bitwise operators ^, |, &
- [ ] refactor: unary operators, binary operators, ternary operators
- [ ] feat: use Cobra to enable multiple modes when launching the REPL
- [ ] feat: use quit(), exit(), or Ctrl-D to exit elegantly
- [ ] feat: support for concurrency primitives such as Mutex, RWMutex, atomic
- [ ] feat: support for comments



## License

[MIT License](./LICENSE)
