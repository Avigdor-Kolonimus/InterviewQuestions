# Common questions

## What is Docker?
Docker is a platform for packaging, delivering, and running applications in isolated containers. Containers include the application and all its dependencies, ensuring consistent behavior across environments.

## What is Postman?
Postman is an API testing and development tool. It allows developers to send HTTP requests, inspect responses, manage collections, and automate API tests.

## What are database migrations?
Migrations are version-controlled changes to a database schema. They allow creating, modifying, and rolling back tables, indexes, constraints, and other database objects.

## What are sessions and JWT tokens?

### Session
A session stores user state on the server. The client usually keeps only a session ID in a cookie.

### JWT (JSON Web Token)
A JWT is a signed token containing user data and claims. It is stored on the client and sent with each request. The server validates the signature instead of storing session data.

---

# Scheduler

## What is the difference between preemptive and cooperative multitasking?

### Preemptive
The scheduler can interrupt a running task and switch to another one automatically.

### Cooperative
A task voluntarily yields control to allow other tasks to run.

Go uses a preemptive scheduler.

---

# Context

## What is a context?
A context carries deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines.

## Why is context used?
- Cancellation of operations
- Timeouts and deadlines
- Passing request-scoped metadata

## How does WithCancel work?

```go
ctx, cancel := context.WithCancel(parent)
```

Calling:

```go
cancel()
```

closes the `Done()` channel and notifies all derived contexts.

## How do you use context?

```go
select {
case <-ctx.Done():
    return ctx.Err()
case result := <-ch:
    return result
}
```

Commonly used in HTTP handlers, database queries, RPC calls, and goroutines.

## How is context implemented?
Contexts form a tree. Child contexts inherit cancellation from their parent.

## Types of contexts

```go
context.Background()
context.TODO()
context.WithCancel()
context.WithTimeout()
context.WithDeadline()
context.WithValue()
```

---

# Defer

## What is defer?
`defer` schedules a function call to run when the surrounding function returns.

## Why use defer?
Typically used for cleanup:

```go
defer file.Close()
defer mu.Unlock()
```

## What is the execution order of multiple defers?
LIFO (Last In, First Out).

```go
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
```

Output:

```text
3
2
1
```

## Does defer run before or after return?
The return value is evaluated first, then deferred functions execute, and only after that does the function actually return.

---

# Race Conditions

## What is a race condition?
A race condition occurs when multiple goroutines access the same data concurrently and at least one access is a write without proper synchronization.

## When does a deadlock occur?
A deadlock occurs when goroutines are waiting for each other indefinitely and none can make progress.

Common causes:
- Reading from a channel with no sender
- Sending to a channel with no receiver
- Incorrect mutex usage
- Circular waiting

## How can race conditions be prevented?
- Mutexes (`sync.Mutex`, `sync.RWMutex`)
- Channels
- Atomic operations (`sync/atomic`)
- Avoiding shared mutable state

## Does the `-race` flag detect races at compile time or runtime?
Runtime.

Example:

```bash
go run -race main.go
```

The race detector monitors memory accesses while the program executes.

---

# Select and Switch

## Does select guarantee execution order?
No.

If multiple cases are ready, Go chooses one pseudo-randomly.

## What is the difference between switch and select?

### switch
Works with values and expressions.

```go
switch x {
case 1:
}
```

### select
Works with channel operations.

```go
select {
case <-ch1:
case <-ch2:
}
```

---

# Errors and Panic

## What is panic?
A panic is a runtime error that stops normal execution and begins stack unwinding.

## What is used to handle panic?
`recover()`.

## Where should recover be placed?
Inside a deferred function.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}()
```

## When should errors be used vs panic?

### Errors
Expected failures:
- Validation errors
- Network errors
- Database errors

### Panic
Programmer mistakes or unrecoverable conditions:
- Nil pointer dereference
- Corrupted internal state
- Impossible situations

---

# OOP in Go

## How is OOP implemented in Go?
Go does not have classical OOP. Instead it provides:
- Structs
- Methods
- Interfaces
- Composition

## How is inheritance implemented in Go?
Go does not support inheritance.

Instead it uses composition (embedding):

```go
type Animal struct{}

func (Animal) Eat() {}

type Dog struct {
    Animal
}
```

`Dog` automatically gets access to `Animal`'s methods.

This is often described as:
> "Composition over inheritance."