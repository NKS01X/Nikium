<p align="center">
  <img src="./.vscode/extensions/nikium-syntax/icons/nikium-icon.svg" width="200" alt="Nikium Logo">
</p>

# 🌟 Nikium Documentation

<p align="center">
  <strong>A brutally simple, blazing fast, and elegantly designed modern programming language.</strong>
</p>

<p align="center">
  <a href="#1-quick-install">Installation</a> •
  <a href="#2-the-tutorial">Tutorial</a> •
  <a href="#3-standard-library">Standard Library</a> •
  <a href="#4-architecture--internals">Architecture</a>
</p>

---

## 🚀 1. Quick Install

Install Nikium in seconds. No need to install Go or compile any code.

**Mac / Linux:**
```bash
curl -fsSL https://raw.githubusercontent.com/NKS01X/Nikium/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/NKS01X/Nikium/main/install.ps1 | iex
```

*(Download standalone executable for your OS directly from [Releases page](https://github.com/NKS01X/Nikium/releases)).*

---

## 📖 2. The Tutorial

Nikium is built for developers who want structural discipline of system-level languages combined with fluid, unburdened writing rhythm of modern scripting languages.

### Variables and Declarations
Choose between fast, dynamic code or strict, statically-guaranteed code.

```nikium
// Implicit stack / dynamic type inference
x = 10;
language = "Nikium";
is_fast = true;

// Statically guaranteed allocation
y:i64 = 20;
name:String = "Nikium";
```

### Memory Management (The Core Feature)
Nikium bridges gap between Python's elegance and C++'s memory control by offering deterministic structure instantiation.

* **Stack Allocation:** Use standard `.` access.
* **Heap/Pointer Allocation:** Use strictly enforced `->` access.

```nikium
struct Point {
    x: i64;
    y: i64;

    init() {
        print "Point created";
    }

    destroy() {
        print "Point destroyed";
    }
}

// Stack allocated Point
p1 = Point(10, 20);
p1.x = 15;

// C++ Style Pointers and memory mapping
p2* ptrName = new Point(30, 40);

// Guaranteed type check at parsing time, zero runtime overhead
ptrName->x = 50; 
```

### Control Flow & Scope Safety
Nikium supports loops evaluated strictly in localized, temporary enclosed environment to prevent scope leaking into main stack.

```nikium
for (i = 0; i < 5; i++) {
    print "Highly localized execution: " + i;
}
// 'i' is garbage collected here.

x = 10;
if (x > 5) {
    print "Condition true";
} else {
    print "Condition false";
}
```

### Data Structures
Arrays and Hashes are first-class primitives executing under `O(1)` hashing bounds against internal slice arrays.

```nikium
arr = [1, 2, 3];
hash = {"api": "v1", "turbo": true};

print hash["api"]; // Outputs: v1
print arr[0]; // Outputs: 1

// Modifying data structures
push(arr, 4);
hash["version"] = 2;
```

### Functions and Generics
Nikium supports first-class functions, closures, and C++ style generic templates.

```nikium
// Standard function
fn add(a, b) {
    return a + b;
}

// Generic function
fn max<T>(a: T, b: T) {
    if (a > b) {
        return a;
    }
    return b;
}
```

---

## 📚 3. Standard Library

Nikium incorporates dynamic file-loading standard library approach. By calling `load "stdlib/module.nik"`, execution layer intercepts filesystem, spawns fresh sub-parser context, evaluates file silently, and merges compiled AST object references into working space. **No heavy JIT requirements.**

### `math.nik` (O(1) mathematical bindings)
| Function | Signature | Description | Example |
|---|---|---|---|
| `abs` | `abs(a)` | Absolute value | `abs(-5)` -> `5` |
| `min` / `max` | `min(a, b)` | Smaller/Larger of two values | `min(3, 7)` -> `3` |
| `pow` | `pow(base, exp)` | Integer exponentiation | `pow(2, 3)` -> `8` |
| `clamp` | `clamp(val, lo, hi)` | Clamp value to `[lo, hi]` | `clamp(15, 0, 10)` -> `10` |

### `arrayutils.nik` (Native slice iteration)
*Requires `len` builtin.*

| Function | Signature | Description |
|---|---|---|
| `map` | `map(arr, f)` | Apply `f` to each element, return new array |
| `filter` | `filter(arr, pred)` | Keep elements where `pred` returns true |
| `reduce` | `reduce(arr, f, init)` | Fold array into single value |
| `contains` | `contains(arr, val)` | True if `val` in array |
| `sum` | `sum(arr)` | Sum of integer array |
| `reverse` | `reverse(arr)` | Reversed copy of array |
| `indexOf` | `indexOf(arr, val)` | Index of `val` or -1 |

**Example Usage:**
```nikium
load "stdlib/arrayutils.nik";
arr = [1, 2, 3, 4, 5];
doubled = map(arr, fn(x) { return x * 2; });
```

### `stringutils.nik` (Native `[]byte` pointer mapping)
*Requires `ord` and `chr` builtins.*

| Function | Signature | Description |
|---|---|---|
| `upper` / `lower` | `upper(s)` | Uppercase/Lowercase string |
| `trim` | `trim(s)` | Strip leading/trailing spaces and tabs |
| `repeat` | `repeat(s, n)` | Repeat string `n` times |
| `split` | `split(s, sep)` | Split string by separator char |
| `indexOf` | `indexOf(s, sub)` | Index of substring `sub` or -1 |

### `input.nik` & Native Builtins
No `load` statement required for builtins. Available globally.

| Name | Description |
|---|---|
| `readLine()` / `readInt()` | Read raw line or integer from stdin (`input.nik`) |
| `len(x)` | Length of string or array |
| `push(arr, val)` | Append to array, return new array |
| `Print(...)` | Print values to stdout |
| `ord(c)` / `chr(n)` | Char to ASCII integer / ASCII integer to char |

---

## 🏛️ 4. Architecture & Internals

Definitive guide on Nikium's internals.

### The "God Nodes"
Core entities powering Nikium's AST execution cycle:

| Component | Centrality | Purpose in Core Loop |
|-----------|------------|---------------------------|
| `Parser` | **Bridge** | Builds AST recursively via Pratt Parsing. Connects raw syntax to executable semantics. |
| `Eval()` | **Hub** | Heart of interpreter. Recursively evaluates AST against environment. Highly optimized. |
| `String` | **Connector**| Powers primary object type, handles memory referencing for text elements. |
| `NewEnv` | **Allocator**| Handles dynamic scope building, closure capturing, memory allocations in real-time. |

### The Pratt Parser
Compiler frontend uses recursive descent setup known as **Pratt Parsing**.
1. Associates parsing functions directly to `TokenTypes`.
2. Computes precedence rules instantly (e.g., `ASTERISK` runs before `PLUS`).
3. Scales Infinitely: Adding new token requires adding single line: `registerPrefix(token, handler)`.

### Runtime Optimizations

* **High-Speed Increments:** `++` operator intercepts at *PrefixExpression* phase. Fetches raw integer reference directly from `Environment` map, increments natively in Go's integer space, immediately updates environment pointer—bypassing binary tree traversal entirely.
* **Associative Arrays for State:** Variables access environments via highly efficient Go map structure, avoiding deep linear scope chaining. Closures use `NewEnclosedEnvironment`, allocates only when explicitly needed for captured state memory.
* **Bitwise Logic:** Bitwise shifts (`<<`, `>>`) run faster than multiplication, optimized straight down to hardware-level execution rules. Logical operations (`&&`, `||`) short-circuit lazily, stopping execution tree walk exact moment truth states known.
* **Precise Error Reporting:** Evaluator captures location metadata (line, column). Error messages descriptive (e.g., "Cannot add String to Int on line 4, col 12").

---
<p align="center">
  <em>Forged by NIKHIL</em>
</p>
