# Nikium

<!-- <div align="center">
  <img src="https://via.placeholder.com/400x200.png?text=Nikium+Language" alt="Nikium Banner">
</div> -->

<p align="center">
  <strong>A simple, elegant, and modern programming language.</strong>
</p>

---

## Introduction

Nikium is a simple, statically-typed programming language with a clean and minimal syntax. It is designed to be easy to learn and use, while still being powerful enough to build real-world applications. Nikium is an interpreted language, and this repository contains the source code for the Nikium interpreter, written in Go.

## Features

*   **Simple and Clean Syntax:** Nikium's syntax is designed to be easy to read and write, with a minimal set of keywords and a consistent structure.
*   **Static Typing:** Nikium is a statically-typed language, which means that all variables must have a type. This helps to catch errors at compile time and makes your code more robust.
*   **Indentation-based Blocks:** Nikium uses indentation to define code blocks, similar to Python. This makes the code more readable and eliminates the need for curly braces.
*   **Control Flow:** Nikium supports `if-else` statements and `while` loops for controlling the flow of execution.
*   **REPL:** Nikium comes with a REPL (Read-Eval-Print Loop) for interactive programming.
*   **File Execution:** You can also execute Nikium programs from a file.

## Getting Started

### Prerequisites

To run the Nikium interpreter, you will need to have Go installed on your system. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

### Installation

To build the Nikium interpreter, clone this repository and run the following command in the root directory:

```bash
go build
```

This will create an executable file named `Nikium` (or `Nikium.exe` on Windows) in the root directory.

### Running the REPL

To start the interactive REPL, run the following command:

```bash
./Nikium
```

You will be greeted with the Nikium banner and a prompt where you can start typing your code.

### Running a File

To execute a Nikium program from a file, pass the file path as a command-line argument:

```bash
./Nikium your_program.nik
```

## Language Syntax

### 1. Variable Declaration

You can declare variables with or without a type annotation.

```nikium
// With type annotation
x:i32 = 10;
name:string = "Nikium";

// Without type annotation (type is inferred)
y = 20;
```

### 2. Data Types

Nikium supports the following data types:

*   `i32`: 32-bit signed integer
*   `i64`: 64-bit signed integer
*   `string`: A sequence of characters enclosed in double quotes (`"`)

### 3. Operators

Nikium supports common arithmetic and relational operators.

*   **Arithmetic:** `+`, `-`, `*`
*   **Relational:** `==`, `!=`, `<`, `>`

### 4. Control Flow

#### If-Else Statements

Conditional logic is handled with `if-else` statements.

```nikium
x = 10;

if x > 5
    print "x is greater than 5";
else
    print "x is not greater than 5";
```

#### While Loops

Nikium supports `while` loops for repeated execution of a block of code.

```nikium
i = 0;
while i < 5
    print i;
    i = i + 1;
```

### 5. Built-in Functions

#### `print`

The `print` function is used to display the value of an expression.

```nikium
print "Hello from Nikium!";
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

// Print a greeting
print "Hello, " + name + "!";

// Perform some calculations
z = x + y;
print "The sum of x and y is: ";
print z;

// Use a loop to print numbers
i = 0;
while i < 5
    print i;
    i = i + 1;

print "--- End of Program ---";
```

## Author

Nikium was created by **NIKHIL**.
