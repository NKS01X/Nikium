# Nikium

<p align="center">
  <strong>A simple, elegant, and modern programming language.</strong>
</p>

---

## Introduction

Nikium is a simple, elegant, and modern programming language designed to be easy to learn and use while remaining powerful enough for real-world applications. It is statically-typed, meaning variables can have a type declaration, which helps catch errors early and makes programs more robust. Nikium is designed to balance simplicity, readability, and practical functionality, making it ideal for beginners, scripting, and rapid prototyping.

The language is interpreted by an interpreter written in Go. This repository contains the source code for that interpreter.

## Features

*   **Simple and Clean Syntax:** Nikium's syntax is designed to be easy to read and write.
*   **Static Typing:** Declare variable types to catch errors early. Type inference is also supported.
*   **Curly-Brace Blocks:** Nikium uses curly braces to define code blocks for `if`, `else`, `while`, and functions.
*   **Control Flow:** Supports `if-else` statements and `while` loops.
*   **Functions:** First-class functions are supported with the `fn` keyword.
*   **Rich Data Types:** Supports `i64` integers, `string`s, and `bool`eans.
*   **REPL:** Nikium comes with a REPL (Read-Eval-Print Loop) for interactive programming.
*   **File Execution:** Execute Nikium programs from `.nik` files.

## Getting Started

### Prerequisites

To run the Nikium interpreter, you need Go installed on your system. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

### Installation

Clone this repository and run the following command in the root directory:

```bash
go build
```

This will create an executable file named `Nikium` (or `Nikium.exe` on Windows) in the root directory.

### Running the REPL

To start the interactive REPL, run the command:

```bash
./Nikium
```

### Running a File

To execute a Nikium program from a file, pass the file path as an argument:

```bash
./Nikium your_program.nik
```

## Language Syntax

### 1. Variable Declaration

You can declare variables with or without a type annotation. If a type is not provided, it will be inferred.

```nikium
// With type annotation
x:i64 = 10;
name:string = "Nikium";

// Without type annotation (type is inferred)
y = 20;
is_active = true;
```

### 2. Data Types

*   `i64`: 64-bit signed integer
*   `string`: A sequence of characters enclosed in double quotes (`"`)
*   `bool`: `true` or `false`

### 3. Operators

*   **Arithmetic:** `+`, `-`, `*`, `/`
*   **Relational:** `==`, `!=`, `<`, `>`

### 4. Control Flow

#### If-Else Statements

```nikium
x = 10;

if x > 5 {
    print "x is greater than 5";
} else {
    print "x is not greater than 5";
}
```

#### While Loops

```nikium
i = 0;
while i < 5 {
    print i;
    i = i + 1;
}
```

### 5. Functions

Functions are defined using the `fn` keyword.

```nikium
add = fn(a, b) {
    return a + b;
};

print add(3, 4); // Prints 7
```

### 6. Built-in Functions & Statements

#### `print`

The `print` statement is used to display the value of an expression.

```nikium
print "Hello from Nikium!";
```

#### `abs`

The `abs` function returns the absolute value of an integer.

```nikium
print abs(-10); // Prints 10
```

## Example Program

Here is a complete example program that demonstrates some of the features of the Nikium language:

```nikium
// example.nik

print "--- Nikium Example Program ---";

// Variable declaration
name:string = "World";
x = 10;
y = 20;
active:bool = true;

// Print a greeting
print "Hello, " + name + "!";

// Perform some calculations
z = x + y;
print "The sum of x and y is:";
print z;

// Use a loop to print numbers
i = 0;
while i < 5 {
    print i;
    i = i + 1;
}

// Use a conditional
if (active) {
    print "System is active.";
}

print "--- End of Program ---";
```

## Author

Nikium was created by **NIKHIL**.
