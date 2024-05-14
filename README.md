# [WIP] Monkey

Monkey is an interpreted language written in Go. *This project is still under development.*

## Dependencies

+ `direnv`
+ `just` (not necessary)
+ Go 1.21+
+ ginkgo (if you want to run local unit tests)
+ golangci-lint (if you want to do local lint)

## Usage

Open your shell, then REPL ready for you:

```bash
➜  ~ just run
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
Hello xxx! This is the Monkey programming language!
>>> 
```

## Demo

### Functions

Anonymous function and function binding:

```bash
>>> fn(x, y) { return x * y; } (5, 6);
30
>>> let add = fn(x, y) { return x + y; };
>>> add(2, 8);
10
```

### Conditionals

Simple if condition with an else branch:

```bash
>>> if (10 > 5) { 5; } else { 10; };
5
```

## Conventions and Features

+ Programs can run in REPL or as scripts
+ Identifiers consist of alphabet letters or underscore
+ Statements are explicitly end by seminlon `;`
+ Comments start with double slash `//`
+ Conditional statements (`if` and `switch` keywords)
+ Loops (`for`, `range`, and `while` keywords)
+ First-class functions and closures
+ Prefix operators (binary expressions)
+ Infix operators (unary expressions)
+ Postfix operators
+ A Pratt Parser implementation
+ A Tree-walking interpreter
+ Use Go's GC to prevent memory leak

## Components

+ Token set
+ Lexer
+ Abstract Syntax Tree (AST)
+ Pratt parser based on context-free grammars and the Backus-Naur-Form
+ Tree-walking interpreter (evaluator), future: bytecode, VM, and JIT compilation
+ An object system

## TODOs

+ [ ] docs: doc everything related to usage and implementation details
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
+ [ ] feat: support for concurrency primitives such as Mutex, RWMutex, atomic
+ [ ] feat: support for comments
+ [ ] docs: a diagram for the full REPL loop including the AST used
+ [ ] check whether we need the token field in AST
+ [ ] test: increase test coverage to at least 80%
+ [ ] feat: return can only be used in functions, do not allow plain return in REPL
+ [ ] feat: if expression with multiple branches
+ [ ] feat: check peek token type in parseExpression
+ [ ] feat: check the difference between if expression and if statement
+ [ ] feat: empty statement with only a single ;
+ [ ] feat: scroll in command history with up and down keys
+ [ ] feat: PrettyPrint, color, AST, etc
+ [ ] feat: sys call such as print(...)
+ [ ] feat: add a helper for available functions
+ [ ] feat: switch between multiple value representation system using some flag
+ [ ] feat: class (object system)
+ [ ] feat: configuration with yaml or envrc
+ [ ] feat: edge cases for those operators
+ [ ] feat: integer division operator and float division operator
+ [ ] feat: reference integer literal as constant, simulate some static memory space for literals of integer, strings, etc.
+ [ ] feat: integer overflow problem
+ [ ] feat: command history and navigate in REPL using left, right, up, bottom
+ [ ] feat: configuration as env vars, default + direnv
+ [ ] feat: semantics for left and right arrows in REPL
+ [ ] feat: semantics for up and down arrows in REPL
+ [ ] feat: lexing, parsing, evaluation of nil expression statement
+ [ ] fix: slow startup issue
+ [ ] fix: do not allow plain return in REPL outside of a function, report an error instead
+ [ ] docs: add a section for all control structures
+ [ ] feat: semantics of scope
+ [ ] feat: pointer/reference semantics
+ [ ] feat: func multiple return values
+ [ ] refactor: evalExpressions (int, error)
+ [ ] feat: parallel assignment
+ [ ] feat: escaping characters and error when " mismatches
+ [ ] perf: add immutable constants to the envvironment to reduce memory allocation
+ [ ] feat: add a print builtin function
+ [ ] feat: add quit(), exit() builtin functions to exit elegantly
+ [ ] feat: add a static/dynamic type system
+ [ ] ci: build and publish as a pkg on Docker Hub
+ [ ] feat: dot as an operator (similar to infix index expression)
+ [ ] feat: mutable array implementation for efficient push/pop

## References

+ *Writing An Interpreter In Go, Thorsten Ball*
+ *Top Down Operator Precedence, Vaughan Pratt*
+ *The Structure and Interpretation of Computer Programs (SICP), Harold Abelson*

## License

[MIT License](./LICENSE)
