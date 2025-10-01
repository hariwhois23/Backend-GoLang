# Variables and Declarations in Go

## Variable Declaration

Go provides multiple ways to declare variables.

### 1. Using `var` Keyword

```go
var name string
var age int
var isActive bool

// With initialization
var name string = "John"
var age int = 30
var isActive bool = true

// Type inference
var name = "John"      // string inferred
var age = 30           // int inferred
var isActive = true    // bool inferred
```

### 2. Short Declaration (`:=`)

Most common inside functions:

```go
name := "John"
age := 30
isActive := true
```

**Note:** Short declaration can only be used inside functions, not at package level.

### 3. Multiple Variable Declaration

```go
// Multiple variables of same type
var x, y, z int

// Multiple variables with initialization
var x, y, z int = 1, 2, 3

// Type inference
var x, y, z = 1, 2, 3

// Short declaration
x, y, z := 1, 2, 3

// Different types
var (
    name   string = "John"
    age    int    = 30
    height float64 = 5.9
)
```

## Zero Values

Variables declared without explicit initialization get zero values:

```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // "" (empty string)
var p *int       // nil
var slice []int  // nil
var m map[string]int // nil
```

## Constants

Declared using `const` keyword:

```go
const Pi = 3.14159
const MaxUsers = 100

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusServerError = 500
)

// Typed constants
const (
    Name string = "MyApp"
    Version string = "1.0.0"
)
```

### Iota

Used for creating enumerated constants:

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

// With expressions
const (
    _  = iota             // 0 (ignored)
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB                    // 1 << 20 = 1048576
    GB                    // 1 << 30
    TB                    // 1 << 40
)
```

## Scope Rules

### Package Level (Global)

```go
package main

var globalVar = "accessible everywhere"

func main() {
    fmt.Println(globalVar)
}
```

### Function Level

```go
func example() {
    localVar := "only in this function"
    // localVar not accessible outside
}
```

### Block Level

```go
func example() {
    x := 10
    
    if true {
        y := 20      // only accessible in if block
        fmt.Println(x, y)
    }
    
    // y is not accessible here
    fmt.Println(x)
}
```

## Type Declaration

Create custom types:

```go
// Type alias
type MyInt int
type MyString string

// Struct type
type Person struct {
    Name string
    Age  int
}

// Function type
type Handler func(string) error

// Interface type
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

## Variable Naming Rules

**Valid names:**
```go
name := "John"
firstName := "John"
first_name := "John"
name1 := "John"
_name := "John"
```

**Invalid names:**
```go
1name := "John"      // cannot start with digit
first-name := "John" // cannot use hyphen
```

**Conventions:**
- Use camelCase: `firstName`, `lastName`
- Exported names start with uppercase: `UserName`, `MaxRetries`
- Unexported names start with lowercase: `userName`, `maxRetries`
- Use short names for local variables: `i`, `j`, `err`
- Use descriptive names for package-level variables

## Blank Identifier

Use `_` to ignore values:

```go
// Ignore error
value, _ := strconv.Atoi("123")

// Ignore index
for _, value := range slice {
    fmt.Println(value)
}

// Import for side effects only
import _ "database/sql/driver"
```

## Pointers

```go
var x int = 10
var p *int = &x    // p points to x

fmt.Println(*p)    // dereference: prints 10
*p = 20            // modify through pointer
fmt.Println(x)     // prints 20

// Short declaration
y := 42
ptr := &y
```

## New vs Make

### `new(T)`

Allocates zeroed storage and returns a pointer:

```go
p := new(int)     // *int, points to 0
*p = 42

s := new([]int)   // *[]int, points to nil slice
```

### `make(T)`

Initializes slices, maps, and channels:

```go
slice := make([]int, 5)           // slice of length 5
m := make(map[string]int)         // initialized map
ch := make(chan int, 10)          // buffered channel
```

## Examples

### Basic Usage

```go
package main

import "fmt"

func main() {
    // Various declarations
    var name string = "Alice"
    age := 25
    var height, weight = 5.6, 130
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    fmt.Printf("Height: %.1f, Weight: %d\n", height, weight)
}
```

### Struct Declaration

```go
type User struct {
    ID       int
    Username string
    Email    string
}

func main() {
    // Different ways to initialize
    user1 := User{1, "john", "john@example.com"}
    
    user2 := User{
        ID:       2,
        Username: "jane",
        Email:    "jane@example.com",
    }
    
    var user3 User
    user3.ID = 3
    user3.Username = "bob"
    user3.Email = "bob@example.com"
}
```

### Multiple Return Values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

## Key Takeaways

- Use `:=` for local variables inside functions
- Use `var` for package-level variables or when explicit type is needed
- Variables have zero values by default
- Constants are immutable and evaluated at compile time
- Use `iota` for enumerated constants
- Pointers allow indirect access to values
- Use `make` for slices, maps, and channels
- Use `new` when you need a pointer to a zero value