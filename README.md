# Monkey

An interpreted language written in Go

## Dependencies

+ `direnv`
+ `just` (not necessary)
+ Go 1.21+
+ ginkgo (if you want to run local unit tests)
+ golangci-lint (if you want to do local linting)

## Conventions

+ Identifiers only consist of alphabet letters or underscore

## TODOs

- [ ] feat: add Unicode support
- [x] feat: add a return keyword for function
- [ ] feat: add line, column number when getting errors
- [ ] feat: add support for hex notation and octal notation for integers
- [ ] feat: formatting and prettier
- [ ] feat: add support for basic types: boolean, float, struct, string, byte
- [ ] feat: add support for collection types: array, map
- [ ] feat: add support for function anonymous arguments
- [ ] feat: add support for variadic functions
- [ ] feat: add support for anonymous functions
- [ ] feat: add <=, >= operators
- [ ] feat: add logical operators ||， &&
- [ ] feat: add bitwise operators ^, |, &
- [ ] refactor: unary operators, binary operators, ternary operators
- [ ] feat: use Cobra to enable multiple modes when launching the REPL
- [ ] feat: use quit(), exit(), or Ctrl-D to exit
- [ ] ci: fix the release GHA workflow
- [ ] feat: concurrency, Mutex, RWMutex







## License

[MIT License](./LICENSE)
