# Go Language: Beginner's Installation Guide

## What is Go?

Go (also called Golang) is an open-source programming language developed by Google. It's designed for simplicity, efficiency, and excellent support for concurrent programming.

## Installation

### macOS

**Option 1: Using Installer**
1. Download the PKG file from [go.dev/dl](https://go.dev/dl/)
2. Run the package installer
3. Installation path: `/usr/local/go`

**Option 2: Using Homebrew**
```bash
brew install go
```

Verify installation:
```bash
go version
```

### Go Modules (Modern Approach)
Create a project anywhere:
```bash
mkdir my-project
cd my-project
go mod init example.com/my-project
```

## Your First Go Program

Create a file named `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:
```bash
go run main.go
```

Build an executable:
```bash
go build main.go
./main  # On Linux/Mac

```

## Essential Commands

```bash
go run file.go        # Run a Go file
go build             # Compile packages
go install           # Compile and install packages
go test              # Run tests
go fmt               # Format code
go get               # Download and install packages
go mod init          # Initialize a new module
go mod tidy          # Clean up dependencies
go doc               # Show documentation
```

## Common Use Cases

- Web servers and APIs
- Cloud services and microservices
- DevOps tools (Docker, Kubernetes built with Go)
- Command-line tools
- Network programming
- System programming

## Why Choose Go?

✅ Fast execution and compilation  
✅ Easy to learn and read  
✅ Excellent for concurrent programming  
✅ Strong community and ecosystem  
✅ Great for cloud-native applications  
✅ Single binary deployment  
✅ Backed by Google  

---
