# Nikium Standard Library

Import any stdlib file at the top of your `.nik` script using `load`.

---

## math.nik
| Function | Signature | Description |
|---|---|---|
| `abs` | `abs(a)` | Absolute value |
| `min` | `min(a, b)` | Smaller of two values |
| `max` | `max(a, b)` | Larger of two values |
| `pow` | `pow(base, exp)` | Integer exponentiation |
| `clamp` | `clamp(val, lo, hi)` | Clamp value to `[lo, hi]` |

```nikium
load "stdlib/math.nik";
print abs(-5);         // 5
print min(10, 5);      // 5
print max(10, 5);      // 10
print pow(2, 3);       // 8
print clamp(10, 0, 5); // 5
```

---

## arrayutils.nik
| Function | Signature | Description |
|---|---|---|
| `map` | `map(arr, f)` | Apply `f` to each element, return new array |
| `filter` | `filter(arr, pred)` | Keep elements where `pred` returns true |
| `reduce` | `reduce(arr, f, init)` | Fold array into single value |
| `contains` | `contains(arr, val)` | True if `val` in array |
| `sum` | `sum(arr)` | Sum of integer array |
| `reverse` | `reverse(arr)` | Reversed copy of array |
| `indexOf` | `indexOf(arr, val)` | Index of `val` or -1 |

```nikium
load "stdlib/arrayutils.nik";
arr = [1, 2, 3];
print map(arr, fn(x) { return x * 2; }); // [2, 4, 6]
print filter(arr, fn(x) { return x > 1; }); // [2, 3]
print reduce(arr, fn(acc, x) { return acc + x; }, 0); // 6
print contains(arr, 2); // true
print sum(arr);         // 6
print reverse(arr);     // [3, 2, 1]
print indexOf(arr, 2);  // 1
```

> Requires: `len` builtin (array support added in evaluator)

---

## stringutils.nik
| Function | Signature | Description |
|---|---|---|
| `upper` | `upper(s)` | Uppercase string |
| `lower` | `lower(s)` | Lowercase string |
| `trim` | `trim(s)` | Strip leading/trailing spaces and tabs |
| `repeat` | `repeat(s, n)` | Repeat string `n` times |
| `startsWith` | `startsWith(s, prefix)` | True if `s` starts with `prefix` |
| `endsWith` | `endsWith(s, suffix)` | True if `s` ends with `suffix` |
| `split` | `split(s, sep)` | Split string by separator char |
| `indexOf` | `indexOf(s, sub)` | Index of substring `sub` or -1 |

```nikium
load "stdlib/stringutils.nik";
print upper("hello");          // "HELLO"
print lower("HELLO");          // "hello"
print trim("  hi  ");          // "hi"
print repeat("hi ", 3);        // "hi hi hi "
print startsWith("hello", "h");// true
print endsWith("hello", "o");  // true
print split("a,b", ",");       // ["a", "b"]
print indexOf("hello", "ll");  // 2
```

> Requires: `ord` and `chr` builtins

---

## input.nik
| Function | Signature | Description |
|---|---|---|
| `readLine` | `readLine()` | Read raw line from stdin |
| `readString` | `readString()` | Read first whitespace-delimited token |
| `readInt` | `readInt()` | Read integer from stdin |
| `readArray` | `readArray()` | Read space-separated integers from stdin |

```nikium
load "stdlib/input.nik";
// Read a line of text
line = readLine(); 

// Read a token (stops at space)
token = readString();

// Read single integer
num = readInt();

// Read array of space separated ints
arr = readArray();
```

---

## Builtins (native)
| Name | Description |
|---|---|
| `len(x)` | Length of string or array |
| `push(arr, val)` | Append to array, return new array |
| `Print(...)` | Print values to stdout |
| `readline()` | Read line from stdin |
| `readchar()` | Read single char from stdin |
| `ord(c)` | Char to ASCII integer |
| `chr(n)` | ASCII integer to char |

---

## sql.nik
A SQL DSL for building queries using chainable functions, producing a struct with `sql` string and `args` array.

| Function | Signature | Description |
|---|---|---|
| `Select` | `Select(fields)` | Initialize a SELECT query |
| `Insert` | `Insert(table, fields, args)` | Initialize an INSERT query |
| `Update` | `Update(table, fields, args)` | Initialize an UPDATE query |
| `From`   | `From(q, table)` | Set the FROM table |
| `Where`  | `Where(q, field, op, value)` | Append a WHERE condition |
| `Limit`  | `Limit(q, limit)` | Set the LIMIT clause |
| `Offset` | `Offset(q, offset)` | Set the OFFSET clause |
| `toSql`  | `toSql(q)` | Compile the query to a SQL string and arguments |

### Example Usage
```nikium
load "stdlib/sql.nik";

q = Select(["id", "name", "email"]);
From(q, "users");
Where(q, "age", ">", 18);
Where(q, "status", "=", "active");
Limit(q, 10);

res = toSql(q);
print res.sql;  // SELECT id, name, email FROM users WHERE age > $1 AND status = $2 LIMIT $3;
print res.args; // [18, "active", 10]
```

### Internal Structs
The builder functions construct query and condition structs internally:
```nikium
query = struct {
    type: "",       // select | insert | update | delete
    table: "",
    alias: "",
    fields: [],
    joins: [],
    where: [],      // array of condition structs
    args: [],
    groupBy: [],
    having: "",
    orderBy: [],
    limit: 0,
    offset: 0,
    distinct: false
};

condition = struct {
    field: "",
    op: "=",
    value: ""
};

order = struct {
    field: "",
    direction: ""   // asc | desc
};

join = struct {
    type: "",    // inner, left, right
    table: "",
    on: ""
};
```
