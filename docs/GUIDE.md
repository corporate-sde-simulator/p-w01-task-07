# Learning Guide - Go (Golang)

> **Welcome to Product-Track Week 1, Task 7!**
> This is a **learning task**. This guide teaches you the Go (Golang) concepts you need and tells you exactly where to look and what to fix. Take your time and read carefully.

---

## What You Need To Do (Summary)

1. **Read** `TICKET.md` to understand the task
2. **Read** this guide to learn the Go (Golang) syntax you'll need
3. **Find the bugs** in the `src/` files (look for `// BUG:` or `# BUG:` comments)
4. **Fix the bugs** using what you learned
5. **Run the tests** to verify your fix: `go test -v ./...`
6. **Add new tests** in the `tests/` folder to prove your fix works

---

## Go (Golang) Quick Reference

### Variables and Types
```go
name := "Alice"                  // Short declaration (type inferred)
var count int = 42               // Explicit type
price := 19.99                   // float64
items := []int{1, 2, 3}         // slice (like a dynamic array)
config := map[string]string{     // map (dictionary)
    "key": "value",
}
isActive := true                 // bool
```

### Functions
```go
func greet(name string) string {
    return "Hello, " + name + "!"
}

// Multiple return values (very common in Go!)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil   // nil = no error
}

// Calling:
result, err := divide(10, 3)
if err != nil {
    fmt.Println("Error:", err)
}
```

### Structs (like classes)
```go
type Calculator struct {
    history []int                // field (lowercase = private)
}

// Constructor function (Go doesn't have constructors)
func NewCalculator() *Calculator {
    return &Calculator{history: []int{}}
}

// Method (function attached to a struct)
func (c *Calculator) Add(a, b int) int {
    result := a + b
    c.history = append(c.history, result)
    return result
}

func (c *Calculator) GetHistory() []int {
    return c.history
}

// Using it:
calc := NewCalculator()
calc.Add(2, 3)
fmt.Println(calc.GetHistory())  // [5]
```

### Maps (Key-Value Storage)
```go
user := map[string]string{"name": "Alice"}
user["name"]                     // Access: "Alice"
user["email"] = "alice@test.com" // Add/update
value, ok := user["name"]       // Check if exists (ok = true/false)
delete(user, "email")           // Remove
len(user)                       // Length
```

### Slices (Dynamic Arrays)
```go
items := []int{1, 2, 3}
items = append(items, 4)        // Add: [1, 2, 3, 4]
len(items)                      // Length: 4
for i, item := range items {   // Loop with index
    fmt.Println(i, item)
}
```

### Error Handling (Go uses explicit error returns)
```go
result, err := someFunction()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
// Use result safely here
```

### Concurrency (Goroutines & Mutexes)
```go
import "sync"

var mu sync.Mutex

func safeIncrement(counter *int) {
    mu.Lock()           // Lock before writing
    *counter++
    mu.Unlock()         // Unlock after writing
}

// Or use defer:
func safeRead(counter *int) int {
    mu.Lock()
    defer mu.Unlock()   // Automatically unlocks when function returns
    return *counter
}
```

### How to Run Tests
```bash
# From the task folder:
go test -v ./...

# With race detector:
go test -race -v ./...
```

### How to Add a Test
```go
func TestSomethingSpecific(t *testing.T) {
    obj := NewProcessor()
    result, err := obj.Process(input)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}
```

---

## Project Structure

| File | Purpose |
|------|---------|
| `TICKET.md` | Your task assignment - **read this first!** |
| `src/configStore.go` | Main source code - **has bugs to fix** |
| `src/fileWatcher.go` | Supporting code - **may also have bugs** |
| `tests/configStore_test.go` | Test file - **add your tests here** |
| `docs/DESIGN.md` | Architecture decisions (background reading) |
| `docs/GUIDE.md` | This learning guide |
| `.context/` | Meeting notes and PR comments (background context) |

---

## Where to Look and What to Fix

### Bug #1 - in `src/configStore.go` (around line 36)
**What's wrong:** Only looks at top-level key, ignores nested traversal

**How to find it:** Open `src/configStore.go` and look for the `BUG:` comment near line 36. Read the comment carefully - it describes exactly what's broken.

**Hint:** Think about what the correct behavior should be (described in `TICKET.md` under Requirements), then compare it to what the code actually does.

### Bug #2 - in `src/configStore.go` (around line 78)
**What's wrong:** Replaces entire subtree instead of deep merging

**How to find it:** Open `src/configStore.go` and look for the `BUG:` comment near line 78. Read the comment carefully - it describes exactly what's broken.

**Hint:** Think about what the correct behavior should be (described in `TICKET.md` under Requirements), then compare it to what the code actually does.

### Bug #3 - in `src/configStore.go` (around line 128)
**What's wrong:** representation.

**How to find it:** Open `src/configStore.go` and look for the `BUG:` comment near line 128. Read the comment carefully - it describes exactly what's broken.

**Hint:** Think about what the correct behavior should be (described in `TICKET.md` under Requirements), then compare it to what the code actually does.

### Bug #4 - in `src/fileWatcher.go` (around line 68)
**What's wrong:** Never updates lastModTime after reload

**How to find it:** Open `src/fileWatcher.go` and look for the `BUG:` comment near line 68. Read the comment carefully - it describes exactly what's broken.

**Hint:** Think about what the correct behavior should be (described in `TICKET.md` under Requirements), then compare it to what the code actually does.


---

## Step-by-Step Workflow

1. Open a terminal and navigate to this task folder
2. Read `TICKET.md` completely
3. Open `src/` files and find the `BUG:` comments
4. Fix each bug (the comment tells you what's wrong)
5. Run the tests:
   ```bash
   go test -v ./...
   ```
6. If tests pass, you're done with the fix
7. **Bonus:** Add a new test to `tests/` that specifically tests the bug you fixed

---

## Common Mistakes to Avoid

- Don't change the function signatures (method names, parameters) - only fix the logic inside
- Don't delete the `BUG:` comments - they help reviewers understand what you changed
- Make sure all existing tests still pass after your changes
- If you're stuck, re-read the `TICKET.md` requirements section carefully
