# Go Fundamentals

A comprehensive collection of Go programming exercises and examples designed to cover most fundamental aspects of the Go programming language. This is a personal learning project that will eventually include extensive examples and exercises covering the core concepts of Go.

## Project Structure

- `hello-world/` - Basic Hello World program with multi-language support
- `integers/` - Working with integers, basic arithmetic operations
- `iteration/` - For loops, iteration patterns, and string building
- `arrays-and-slices/` - Arrays, slices, variadic functions, and related algorithms
- `structs-methods-interfaces/` - Structs, methods, interfaces, and geometric calculations

_(More modules will be added as the project progresses)_

## Getting Started

### Prerequisites

- Go 1.18 or newer installed ([download](https://golang.org/dl/))

### Clone the repository

```bash
git clone https://github.com/QendrimBehrami/go-fundamentals.git
cd go-fundamentals
```

### Running Examples & Tests

To run the Hello World example:

```bash
cd hello-world
go run hello.go
```

To run tests for a specific module (e.g., arrays-and-slices):

```bash
cd arrays-and-slices
go test
```

To run all tests in the project:

```bash
go test ./...
```

3. Run the integers example:

```bash
cd integers
go test
```

4. Run the iteration example:

````bash
cd iteration
go test
```bash
cd arrays-and-slices
go test
````

### Running Tests

To run all tests in the project:

```bash
go test ./...
```

To run tests for a specific package:

```bash
cd hello-world
go test
```

To run tests with verbose output:

```bash
go test -v
```

## Examples

### Hello World

The hello-world example demonstrates:

- Basic function creation
- String manipulation
- Multi-language greetings (English, Spanish, French)
- Test-driven development with Go

```go
// Example usage
fmt.Println(Hello("Chris", "English"))   // Output: Hello, Chris
fmt.Println(Hello("Elodie", "Spanish"))  // Output: Hola, Elodie
fmt.Println(Hello("Louis", "French"))    // Output: Bonjour, Louis
```

### Integers

The integers example demonstrates:

- Working with integer types
- Basic arithmetic operations
- Function parameters and return values
- Example tests for documentation

```go
// Example usage
sum := Add(1, 5)
fmt.Println(sum)  // Output: 6
```

### Iteration

The iteration example demonstrates:

- For loops and range iteration
- String building with `strings.Builder`
- Benchmark testing
- Performance considerations

```go
// Example usage
result := Repeat("a", 5)
fmt.Println(result)  // Output: aaaaa
```

### Arrays and Slices

The arrays-and-slices example demonstrates:

- Declaring and using arrays and slices
- Variadic functions
- Summing collections
- Slicing and copying data
- Handling empty collections
- Test-driven development with edge cases

```go
// Example usage
numbers := []int{1, 2, 3, 4, 5}
sum := Sum(numbers)
fmt.Println(sum) // Output: 15

sums := SumAll([]int{1,2}, []int{0,9})
fmt.Println(sums) // Output: [3 9]

sumsTails := SumAllTails([]int{1,2,3}, []int{0,9})
fmt.Println(sumsTails) // Output: [5 9]
```

### Structs, Methods, and Interfaces

The `structs-methods-interfaces` example demonstrates:

- Defining structs for geometric shapes (Rectangle, Circle, Triangle)
- Implementing methods on structs
- Using interfaces for polymorphism
- Table-driven tests and subtests

```go
// Example usage
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area()) // Output: 50

var s Shape = Circle{Radius: 3}
fmt.Println(s.Area()) // Output: 28.274333882308138
```

## Learning Goals

This project aims to cover comprehensive Go fundamentals including:

**Currently Implemented:**

- Package structure and imports
- Functions and return values
- Constants and variables
- Control structures (switch statements)
- Testing with the `testing` package
- Test-driven development (TDD)
- Integer types and basic arithmetic
- Example tests for documentation
- For loops and iteration patterns
- String building and performance optimization
- Benchmark testing
- Arrays and slices: declaration, usage, and algorithms
- Variadic functions and edge case handling

**Planned Topics:**

- Variables, constants, and basic data types
- Arrays, slices, and maps
- Structs and methods
- Interfaces and polymorphism
- Pointers and memory management
- Error handling patterns
- Goroutines and concurrency
- Channels and communication
- File I/O operations
- JSON handling
- HTTP client/server programming
- Database interactions
- Package organization and modules
- Build tags and compilation
- Benchmarking and profiling

## Progress

🟢 **Completed:** Hello World, Integers, Iteration, Arrays & Slices, Structs/Methods/Interfaces chapters with comprehensive testing  
🟡 **In Progress:** Expanding fundamental concepts  
🔴 **Planned:** Advanced topics and real-world applications

_This project is under active development and will be expanded regularly with new examples and exercises._

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [The Go Programming Language Specification](https://golang.org/ref/spec)
- [Go Tour](https://tour.golang.org/)
