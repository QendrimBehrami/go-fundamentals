# Go Fundamentals

[![Go Version](https://img.shields.io/badge/Go-1.24.5-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/QendrimBehrami/go-fundamentals)](https://goreportcard.com/report/github.com/QendrimBehrami/go-fundamentals)
[![Tests](https://github.com/QendrimBehrami/go-fundamentals/workflows/Tests/badge.svg)](https://github.com/QendrimBehrami/go-fundamentals/actions)

A comprehensive collection of Go programming exercises and examples designed to cover most fundamental aspects of the Go programming language. This is a personal learning project that will eventually include extensive examples and exercises covering the core concepts of Go.

## Project Structure

- `hello-world/` - Basic Hello World program with multi-language support
- `integers/` - Working with integers, basic arithmetic operations
- `iteration/` - For loops, iteration patterns, and string building
- `arrays-and-slices/` - Arrays, slices, variadic functions, and related algorithms
- `structs-methods-interfaces/` - Structs, methods, interfaces, and geometric calculations
- `pointers-and-errors/` - Pointers, memory management, and error handling patterns
- `maps/` - Working with maps, custom dictionary types, error handling for lookups

_(More modules will be added as the project progresses)_

## Getting Started

### Prerequisites

- Go 1.21 or newer installed ([download](https://golang.org/dl/))
- This project uses Go 1.24.5 with modern features like the new benchmark API

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

The `hello-world` example demonstrates:

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

The `ìntegers` example demonstrates:

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

The `iteration` example demonstrates:

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

The `arrays-and-slices` example demonstrates:

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
- Implementing methods with both value and pointer receivers
- Using interfaces for polymorphism (`Shape` interface)
- Advanced mathematical calculations (Heron's formula for triangle area)
- Table-driven tests with subtests for comprehensive testing
- Performance benchmarking with Go 1.24's new `b.Loop()` API

```go
// Example usage - Shape interface polymorphism
shapes := []Shape{
    Rectangle{Width: 12, Height: 6},
    Circle{Radius: 10},
    Triangle{SideA: 3, SideB: 4, SideC: 5}, // Classic 3-4-5 right triangle
}

for _, shape := range shapes {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
}
// Output:
// Area: 72.00, Perimeter: 36.00
// Area: 314.16, Perimeter: 62.83
// Area: 6.00, Perimeter: 12.00
```

**Key Concepts Covered:**

- **Interface satisfaction**: All shapes implement `Area()` and `Perimeter()`
- **Mathematical accuracy**: Triangle uses Heron's formula with three sides
- **Test-driven development**: Comprehensive tests with known values (3-4-5 triangle)
- **Performance testing**: Benchmarks showing ~1ns/op for simple operations

### Pointers and Errors

The `pointers-and-errors` example demonstrates:

- Understanding pointers and memory addresses
- Value vs pointer receivers for methods
- Error handling patterns in Go
- Creating custom error types
- Bitcoin wallet implementation with balance management
- Deposit, withdraw, and balance operations
- Error scenarios and edge case handling

```go
// Example usage - Bitcoin wallet with error handling
wallet := Wallet{}

wallet.Deposit(Bitcoin(10))
fmt.Println(wallet.Balance()) // Output: 10 BTC

err := wallet.Withdraw(Bitcoin(5))
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println(wallet.Balance()) // Output: 5 BTC
}

// Attempting to withdraw more than available
err = wallet.Withdraw(Bitcoin(100))
if err != nil {
    fmt.Println(err) // Output: insufficient funds
}
```

**Key Concepts Covered:**

- **Pointer receivers**: Methods that modify the receiver must use pointers
- **Error handling**: Go's idiomatic error handling with explicit error checking
- **Custom types**: Creating domain-specific types like `Bitcoin` for type safety
- **Test-driven development**: Comprehensive error scenario testing
- **Memory management**: Understanding when to use pointers vs values

### Maps

The `maps` example demonstrates:

- Using Go's built-in map type
- Implementing custom dictionary types
- Error handling for missing keys
- Table-driven tests for search and error scenarios

```go
// Example usage
dictionary := Dictionary{"test": "this is just a test"}

result, err := dictionary.Search("test")
fmt.Println(result) // Output: this is just a test

_, err = dictionary.Search("unknown")
if err != nil {
    fmt.Println(err) // Output: could not find the word you were looking for
}
```

## Performance & Benchmarking

This project includes comprehensive benchmarking to understand Go's performance characteristics:

````bash
```bash
# Run benchmarks for all modules
go test -bench=. ./...

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem ./...

# Example benchmark results (structs-methods-interfaces on Apple M3):
# BenchmarkRectangleArea-8   	1000000000	         1.000 ns/op
# BenchmarkCircleArea-8      	1000000000	         1.000 ns/op
# BenchmarkTriangleArea-8    	1000000000	         1.005 ns/op
````

**Performance Insights:**

- Apple M3 processor delivers exceptional performance for mathematical operations
- Go compiler optimizations make even complex calculations (Heron's formula) extremely fast
- All geometric calculations complete in ~1 nanosecond, showing the power of modern hardware
- The difference between simple arithmetic and `math.Sqrt()` is negligible on ARM64

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
- Benchmark testing with Go 1.24's new API
- Arrays and slices: declaration, usage, and algorithms
- Variadic functions and edge case handling
- Structs and methods (value and pointer receivers)
- Interface design and polymorphism
- Advanced mathematical implementations (Heron's formula)
- Performance benchmarking and optimization analysis
- Pointers and memory management
- Error handling patterns and custom error types
- Type safety with custom types (Bitcoin example)

**Planned Topics:**

- Variables, constants, and basic data types
- Maps and advanced data structures
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

🟢 **Completed:** Hello World, Integers, Iteration, Arrays & Slices, Structs/Methods/Interfaces, Pointers & Errors chapters with comprehensive testing  
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
