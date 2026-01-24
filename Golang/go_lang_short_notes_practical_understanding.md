## 1. What is Go (Golang)?

- Go is a **compiled, statically typed language** created at **Google**.
- Designed to be:
  - As **easy as Python** to write
  - As **fast and safe as C/C++** to run
- Built for:
  - **Backend systems**
  - **Cloud tools**
  - **Distributed & concurrent systems**

**Mental Model:**  
👉 *Python ki simplicity + C ki performance = Go*

---

## 2. Why Go Exists (Real Reason)

Google engineers were tired of:
- Slow builds (C/C++)
- Complex syntax
- Poor concurrency models

Go fixes this by:
- Very **fast compilation**
- **Simple syntax** (less magic)
- **First‑class concurrency** (goroutines)

**Industry use:**
- Docker 🐳
- Kubernetes ☸️
- Terraform
- Prometheus

---

## 3. Go Program Structure (Minimum Viable)

Every Go program needs:
- `package main`
- `func main()` → execution starts here

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

**Rule:**  
If there is no `main()` → program won’t run.

---

## 4. Go is Statically Typed (But Smart)

- Every variable has a **type**
- Type can be:
  - **Explicit** (`var x int = 10`)
  - **Inferred** (`x := 10`)

```go
name := "Nasir"   // compiler infers string
age := 19         // compiler infers int
```

**Why this is good:**
- Fewer runtime bugs
- Better performance
- Safer refactoring

---

## 5. Core Data Types (What You’ll Actually Use)

### Numbers
- `int` → whole numbers
- `float64` → decimals

### Text
- `string`

### Boolean
- `bool` → `true / false`

**Rule:** Go does NOT auto‑convert types.
```go
var x int = 10
var y float64 = float64(x) // explicit conversion
```

---

## 6. Variables – Clean Rules

### Declaration styles
```go
var city string = "London"

city := "London"   // preferred (inside functions)
```

### Zero Values (Important!)
If you don’t assign a value, Go assigns **zero value**:
- `int` → `0`
- `float` → `0.0`
- `string` → `""`
- `bool` → `false`

This avoids **undefined behavior**.

---

## 7. Printing & Debugging (Daily Use)

- `fmt.Print()` → same line
- `fmt.Println()` → new line
- `fmt.Printf()` → formatted output

```go
fmt.Printf("User %s scored %d marks", name, marks)
```

**Senior tip:**  
Use `Printf` for logs. It’s cleaner.

---

## 8. Conditions (Decision Making)

### if / else
```go
if age >= 18 {
    fmt.Println("Eligible")
} else {
    fmt.Println("Not eligible")
}
```

### switch (very clean in Go)
```go
switch status {
case "running":
    fmt.Println("Service OK")
default:
    fmt.Println("Unknown")
}
```

**Why switch is loved:**
- No break needed
- Very readable

---

## 9. Loops – Only ONE Loop in Go 😎

Go has **only `for` loop**.

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

Other patterns:
- while‑like loop
- infinite loop

Simple. No confusion.

---

## 10. Arrays (Basics Only)

- Fixed size
- Same data type

```go
marks := [3]int{10, 20, 30}
```

**Note:** In real projects → mostly **slices**, not arrays.

---

## 11. Big Picture (DevOps Lens)

Why Go matters for **DevOps / Cloud engineers**:
- Fast CLIs
- Low memory usage
- Easy deployment (single binary)
- Perfect for containers

**If you know Go:**
- Docker internals make sense
- Kubernetes code is readable
- Cloud tools feel less magical

---
