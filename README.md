<p align="center">
  <img src="./.vscode/extensions/nikium-syntax/icons/nikium-icon.svg" width="200" alt="Nikium Logo">
</p>

# 🌟 Nikium: The Ultimate Language 

<p align="center">
  <strong>A brutally simple, blazing fast, and elegantly designed modern programming language.</strong>
</p>

---

## 🚀 Quick Install

Install Nikium in seconds. No need to install Go or compile any code.

**Mac / Linux:**
```bash
curl -fsSL https://raw.githubusercontent.com/NKS01X/Nikium/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/NKS01X/Nikium/main/install.ps1 | iex
```

---

## 📖 The Architecture Book: How Nikium Works

Welcome to the definitive guide on Nikium. This document goes far beyond basic usage—it is a deep dive into the language's core abstractions, architecture, and runtime optimizations, serving as a comprehensive book on the language design.

### 🏛️ The God Nodes of Nikium

Through our rigorous analysis (via our internal `graphify` knowledge system), we've isolated the core entities—the "God Nodes"—that power Nikium's AST execution cycle. These components create a highly responsive structure:

| Component | Centrality / Connections | Purpose in the Core Loop |
|-----------|-------------------------|---------------------------|
| `Parser` | **44 edges** (Bridge) | Takes lexed tokens and builds the AST recursively via Pratt Parsing. Connects raw syntax to executable semantics. |
| `Eval()` | **33 edges** (Hub) | The heart of the interpreter. Recursively evaluates the AST against the environment. Highly optimized. |
| `String` | **24 edges** (Connector)| Powers the primary object type and handles memory referencing for text elements. |
| `New()` / `NewEnv` | **37 edges combined** | Handles dynamic scope building, closure capturing, and memory allocations in real-time. |

---

## ⚡ Why is Nikium Better? Comparisons & Optimizations

Nikium intentionally bridges the gap between C++'s memory control and Python's elegance.

### 1. Pointer vs Stack Semantics (C++ Style Control)
Unlike typical dynamic languages where everything is implicitly reference-counted under the hood, Nikium offers deterministic structure instantiation.

*   **Stack Allocation**: Using `.` access
*   **Heap/Pointer Allocation**: Using `->` access (Strictly enforced)

#### Feature Comparison
| Language | Pointer Support | Ease of Parsing | Runtime Overhead |
|----------|-----------------|-----------------|------------------|
| Python   | ❌ No           | 🥇 High         | 🐢 High          |
| C++      | ✅ Yes          | 🥉 Low          | 🏎️ Zero          |
| **Nikium**| ✅ **Yes**       | 🥇 **High**      | 🚀 **Minimal**    |

**How it Optimizes:**
By explicitly distinguishing `ptr->prop` and `obj.prop`, `evalPropertyAccessExpression` performs a guaranteed type check at parsing time, dropping massive runtime overhead usually involved in extreme dynamic property resolution.

### 2. High-Speed Increments
In typical interpreted scripting languages, `x = x + 1` requires AST traversal of an `AssignExpression`, evaluating an `Identifier`, traversing a `BinaryExpression`, and applying `left + right`.

**Nikium's Optimization:**
The `++` operator intercepts at the *PrefixExpression* phase (see `evalIncExpression`). It fetches the raw integer reference directly from the `Environment` map, increments it natively in Go's integer space, and immediately updates the environment pointer—bypassing binary tree traversal entirely.

### 3. Associative Arrays for State
Variables access environments via a highly efficient Go map structure, avoiding deep linear scope chaining when possible. Closures use `NewEnclosedEnvironment`, which only allocates when explicitly needed for captured state memory.

---

## 🛠️ Comprehensive Syntax Guide

### 🧱 1. Memory and Declarations

Declare types strictly or leverage the type inference engine.

```nikium
// Implicit stack / dynamic type inference
x = 10;

// Statically guaranteed allocation
y:i64 = 20;

// C++ Style Pointers and memory mapping
p* ptrName = new Type(args);
ptrName->value = 50;
```

### 🧮 2. Operators & Bitwise Speed

Support for standard arithmetic is coupled with deep bitwise logic. Bitwise shifts (`<<`, `>>`) run *faster* than multiplication, optimized straight down to hardware-level execution rules.

**Logical Short-Circuiting**: `&&` and `||` evaluate lazily in Nikium, stopping execution tree walk the exact moment truth states are known.

---

## 🔁 3. Control Flows and Logic Arrays

### The `for` Loop Implementation
While traditional `while` loops exist, Nikium supports `for(init; cond; post)` loops evaluated strictly in a localized, temporary `loopEnv` (enclosed environment) to prevent scope leaking strings into the main stack.

```nikium
for (i = 0; i < 5; i++) {
    // Highly localized AST execution
    print i;
}
```

### Deep Array Integrations
Arrays and Hashes are first-class primitives, executing under `O(1)` hashing bounds for structured components against internal slice arrays.

```nikium
arr = [1, 2, 3];
hash = {"api": "v1", "turbo": true};
print hash["api"];
```

---

## 🏗️ The Pratt Parser Architecture

The compiler frontend uses a sophisticated recursive descent setup known as **Pratt Parsing**. This allows it to:
1. Associate parsing functions (`prefixParseFn` / `infixParseFn`) directly to `TokenTypes`.
2. Compute precedence rules instantly (e.g., `ASTERISK` runs before `PLUS`).
3. Scale Infinitely: Adding a new token only requires adding a single line `registerPrefix(token, handler)`.

---

## 📚 Standard Library Interop

Nikium incorporates a dynamic file-loading standard library approach. By calling `load "stdlib/module.nik"`, the execution layer intercepts the filesystem, spawns a completely fresh sub-parser context, evaluates the file silently, and merges the compiled AST object references into your working space. No heavy JIT requirements.

| Module | Core Logic Provided | Speed Profile |
|--------|---------------------|---------------|
| `math.nik` | Min, Max, Pow, Clamp | O(1) mathematical bindings |
| `stringutils`| Upper, Repeat, Split | Native `[]byte` pointer mapping |
| `arrayutils`| Sum, Reverse, IndexOf| Direct native slice iteration |

---

## 🏁 Conclusion

Nikium is built for developers who want the brutal simplicity and structural discipline of system-level languages combined with the fluid, unburdened writing rhythm of modern scripting languages.

*Forged by NIKHIL*
