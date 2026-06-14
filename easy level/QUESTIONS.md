# Go Interview Questions & Answers

## Arrays vs Slices

### What is the difference between arrays and slices in Go?

An **array** has a fixed size, which is part of its type.

```go
var arr [5]int
```

Arrays are value types, so assigning or passing an array copies all its elements.

A **slice** is a dynamic view over an underlying array.

```go
s := []int{1, 2, 3}
```

Copying a slice copies only the slice header, not the underlying data.

---

### What is a slice under the hood?

A slice is a small descriptor containing:

```go
type slice struct {
    ptr *T
    len int
    cap int
}
```

* `ptr` points to the underlying array
* `len` is the number of elements in the slice
* `cap` is the capacity of the underlying array

---

### How does slice reallocation work during append?

If there is enough capacity, `append` reuses the existing backing array.

```go
s := make([]int, 3, 10)
s = append(s, 4)
```

If capacity is exhausted, Go allocates a new array, copies existing elements, and returns a new slice pointing to the new array.

---

## Maps (Hash Tables)

### How is a map implemented in Go?

A Go map is implemented as a hash table consisting of:

* Buckets
* A hash function
* Overflow buckets
* Metadata

Each bucket can store up to 8 key-value pairs.

---

### How does map lookup work?

Go computes a hash for the key:

```go
hash(key)
```

The hash determines the target bucket.

The runtime then searches within that bucket for the key.

---

### What are overflow buckets?

When a bucket becomes full, Go creates overflow buckets linked to the original bucket.

Too many overflow buckets can degrade lookup performance.

---

### Why is a map not thread-safe?

Maps do not use internal synchronization.

Concurrent writes may corrupt internal data structures and cause runtime panics such as:

```text
fatal error: concurrent map writes
```

To safely share maps between goroutines, use:

* `sync.Mutex`
* `sync.RWMutex`
* `sync.Map`

---

### When does map growth and data evacuation occur?

When the load factor becomes too high, Go allocates a larger bucket array.

Instead of moving all data at once, elements are gradually migrated during future map operations. This process is called **evacuation**.

---

### Which types cannot be used as map keys?

Map keys must be **comparable**.

Valid keys:

```go
int
string
bool
pointer
array
struct
```

Invalid keys:

```go
slice
map
function
```

For example:

```go
map[[]int]int
```

does not compile.

---

## Structures and Interfaces

### What is the difference between structs and interfaces?

A **struct** stores data.

```go
type User struct {
    Name string
    Age  int
}
```

An **interface** defines behavior.

```go
type Speaker interface {
    Speak()
}
```

---

### How is polymorphism implemented in Go?

Polymorphism is achieved through interfaces.

```go
type Speaker interface {
    Speak()
}
```

Any type implementing `Speak()` satisfies the interface.

```go
type Dog struct{}

func (Dog) Speak() {
    fmt.Println("Woof")
}
```

```go
type Cat struct{}

func (Cat) Speak() {
    fmt.Println("Meow")
}
```

Both types can be used through the `Speaker` interface.

---

## Pointers

### What are pointers in Go?

A pointer stores the memory address of a value.

```go
x := 10
p := &x
```

---

### Why are pointers used?

Pointers are used to:

* Modify original values
* Avoid copying large objects
* Share data between functions

---

### Why pass pointers to functions?

Without a pointer, a function receives a copy:

```go
func change(v int) {
    v = 5
}
```

With a pointer:

```go
func change(v *int) {
    *v = 5
}
```

the original value is modified.

---

## Strings

### What is a string under the hood?

A string is essentially:

```go
type string struct {
    ptr *byte
    len int
}
```

It stores UTF-8 encoded bytes.

---

### Why are strings immutable in Go?

Immutability provides:

* Safety
* Better memory sharing
* Predictable behavior

Therefore:

```go
s := "hello"
s[0] = 'H'
```

does not compile.

---

### What does `len("Привет")` return?

```go
len("Привет")
```

returns:

```go
12
```

because `len()` counts bytes, not Unicode characters.

Each Cyrillic character occupies 2 bytes in UTF-8.

To count characters:

```go
utf8.RuneCountInString("Привет")
```

returns:

```go
6
```

---

## JSON

### What is JSON?

JSON (JavaScript Object Notation) is a text-based data exchange format.

Example:

```json
{
  "name": "Alex",
  "age": 30
}
```

In Go:

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

Serialization:

```go
json.Marshal()
```

Deserialization:

```go
json.Unmarshal()
```

---

## Concurrency vs Parallelism

### What is the difference between concurrency and parallelism?

**Concurrency** means managing multiple tasks that make progress during overlapping time periods.

Example:

```text
Task A
Task B
Task A
Task B
```

A single CPU core can achieve concurrency through scheduling.

---

**Parallelism** means multiple tasks are executed simultaneously on multiple CPU cores.

```text
CPU1 -> Task A
CPU2 -> Task B
```

---

### Short definition

> Concurrency is about dealing with many tasks at once.
>
> Parallelism is about executing many tasks at the same time.

---

## Goroutines and Synchronization

### What is a goroutine?

A goroutine is a lightweight thread managed by the Go runtime.

```go
go func() {
    fmt.Println("hello")
}()
```

Goroutines are significantly cheaper than OS threads.

---

### What synchronization primitives are available in Go?

Common synchronization mechanisms include:

* Channels
* `sync.WaitGroup`
* `sync.Mutex`
* `sync.RWMutex`
* `sync.Once`
* `sync.Cond`
* Atomic operations
* Context

---

### What is `sync.WaitGroup` used for?

It allows one goroutine to wait for a group of goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()
}()

wg.Wait()
```

---

### What is `sync.Mutex` used for?

A mutex protects shared data from concurrent access.

```go
mu.Lock()

counter++

mu.Unlock()
```

Without synchronization, concurrent access may lead to data races.

---

### When should I use a Mutex vs a Channel?

Use **channels** when goroutines need to communicate and exchange data.

Use a **mutex** when multiple goroutines need safe access to shared state such as:

* Maps
* Caches
* Counters
* Shared structures

A common Go principle is:

> Do not communicate by sharing memory; instead, share memory by communicating.
