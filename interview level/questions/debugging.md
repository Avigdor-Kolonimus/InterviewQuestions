# Application Debugging

## 1. What debugging and profiling tools are available in Go (pprof)?

Go provides built-in tools for debugging and profiling through the **`runtime/pprof`** and **`net/http/pprof`** packages.

---

## 🔧 Main tool: pprof

**pprof** is Go’s official profiling tool used to analyze performance and detect bottlenecks.

It can collect different types of profiles:

### Types of profiles

- **CPU profile**
  - Shows where CPU time is spent.
  - Helps detect hot functions.

- **Heap profile**
  - Shows memory allocations and current heap usage.
  - Useful for finding memory leaks.

- **Goroutine profile**
  - Shows all running goroutines.
  - Helps detect leaks or blocked goroutines.

- **Threadcreate profile**
  - Shows OS threads created by the runtime.

- **Block profile**
  - Shows where goroutines are blocked (e.g., on channels, mutexes).

---

## 🌐 HTTP pprof (most commonly used in services)

You can enable it in a web service:

```go
import _ "net/http/pprof"
```

Then run:

```
http://localhost:6060/debug/pprof/
```

Or use CLI:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

---

## 🧪 Example usage

### CPU profiling

```bash
go tool pprof cpu.prof
```

### Heap profiling

```bash
go tool pprof heap.prof
```

### Web UI

Inside pprof:

```bash
(pprof) web
```

This opens a graph showing function call hotspots.

---

## 🛠 Other Go debugging tools

### 1. `runtime/trace`
- More detailed execution tracing than pprof.
- Shows scheduling, goroutines, syscalls.

```bash
go tool trace trace.out
```

---

### 2. `delve` (dlv)
- Full-featured debugger for Go.
- Supports breakpoints, step execution, variable inspection.

```bash
dlv debug
```

---

### 3. Logging & metrics tools
- Structured logging (zap, logrus)
- Metrics (Prometheus + Grafana)

---

## Summary

- **pprof** → performance profiling (CPU, memory, goroutines)
- **trace** → deep runtime execution analysis
- **delve** → interactive debugging
- **metrics/logging** → production monitoring