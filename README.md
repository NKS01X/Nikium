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
*   **Rich Data Types:** Supports `i64` integers, `string`s, `bool`eans, and arrays.
*   **Logical & Relational Operators:** Features logic operators (`&&`, `||`) and standard relationals.
*   **Standard Library:** Ships with stdlib modules for math, arrays, strings, and I/O.
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

### Comments

Single-line comments start with `//`.

```nikium
// This is a comment
x = 10; // This is also a comment
```

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
*   `string`: A sequence of characters enclosed in double quotes (`"`). Supports escape sequences: `\n` (newline), `\t` (tab), `\\` (backslash), `\"` (double quote).
*   `bool`: `true` or `false`
*   `array`: Arrays enclosed in `[]`

### 3. Operators

*   **Arithmetic:** `+`, `-`, `*`, `/`, `%`
*   **Bitwise:** `<<`, `>>`
*   **Relational:** `==`, `!=`, `<`, `>`
*   **Logical:** `&&`, `||`, `!`

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

Use `break` to exit early and `continue` to skip iteration:

```nikium
i = 0;
while i < 5 {
    i = i + 1;
    if i == 2 {
        continue;
    }
    if i == 4 {
        break;
    }
    print i;
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

### 6. Built-in Functions

| Name | Description |
|---|---|
| `len(x)` | Length of string or array |
| `push(arr, val)` | Append to array, return new array |
| `Print(...)` | Print values to stdout |
| `readline()` | Read line from stdin |
| `readchar()` | Read single char from stdin |
| `ord(c)` | Char → ASCII integer |
| `chr(n)` | ASCII integer → char |

## Standard Library

Stdlib lives in `stdlib/`. Full reference: [`stdlib/README.md`](stdlib/README.md).

### math.nik — Math Utilities

Provides mathematical operations.

```nikium
load "stdlib/math.nik";

print min(10, 5);      // 5
print max(10, 5);      // 10
print pow(2, 3);       // 8
print clamp(10, 0, 5); // 5
```

### array.nik — Array Utilities

Provides array formatting and transformations.

```nikium
load "stdlib/array.nik";

arr = [1, 2, 3];
print sum(arr);         // 6
print contains(arr, 2); // true
print reverse(arr);     // [3, 2, 1]
print indexOf(arr, 2);  // 1
```

### string.nik — String Utilities

Provides string manipulation functionalities.

```nikium
load "stdlib/string.nik";

print upper("hello");          // "HELLO"
print repeat("hi ", 3);        // "hi hi hi "
print startsWith("hello", "h");// true
print indexOf("hello", "ll");  // 2
print split("a,b", ",");       // ["a", "b"]
```

### I/O Modules

| Module | Functions |
|---|---|
| `input.nik` | `readLine()`, `readString()`, `readInt()`, `readArray()` |

## Example Program

```nikium
// example.nik
load "stdlib/math.nik";
load "stdlib/string.nik";
load "stdlib/array.nik";

print "--- Nikium Example Program ---";

name:string = "World";
x = 10;
y = 20;
active:bool = true;

print "Hello, " + name + "!";

z = x + y;
print "The sum of x and y is:";
print z;

print "The min of x and y is:";
print min(x, y);

up_name = upper(name);
print "Uppercase name:";
print up_name;

arr = [1, 2, 3, 4, 5];
print "Sum of array:";
print sum(arr);

i = 0;
while i < 5 {
    print i;
    i = i + 1;
}

if (active) {
    print "System is active.";
}

print "--- End of Program ---";
```

## Author

Nikium was created by **NIKHIL**.
