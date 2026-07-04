# gRPC

## 1. What are the advantages of gRPC compared to REST?

**gRPC** is a high-performance RPC framework developed by Google. It uses **HTTP/2** for transport and **Protocol Buffers (protobuf)** for data serialization.

### Advantages of gRPC

- Higher performance due to binary serialization (protobuf).
- Smaller message size compared to JSON.
- Lower latency.
- Built-in support for HTTP/2 features:
  - Multiplexing
  - Header compression
  - Persistent connections
- Strongly typed API defined in `.proto` files.
- Automatic generation of client and server code.
- Supports bidirectional streaming.
- Cross-platform and supports many programming languages.

---

## 2. How does Protocol Buffers (protobuf) work, and why is it good?

**Protocol Buffers (protobuf)** is Google's language-neutral binary serialization format.

The API and data structures are defined in a `.proto` file.

Example:

```proto
syntax = "proto3";

message User {
  int64 id = 1;
  string name = 2;
}
```

The `protoc` compiler generates code for multiple languages (Go, Java, C++, Python, etc.).

### Advantages

- Compact binary format.
- Faster serialization/deserialization than JSON.
- Strong typing.
- Backward and forward compatibility by preserving field numbers.
- Automatic code generation.

---

## 3. When should you use gRPC, and when should you use REST?

### Use gRPC when:

- Communication is between internal microservices.
- High performance and low latency are important.
- Streaming is required.
- Strong contracts and code generation are desired.
- Multiple services communicate frequently.

Examples:
- Microservices
- Real-time systems
- Internal APIs
- Distributed systems

---

### Use REST when:

- Building public APIs.
- Clients are web browsers or third-party applications.
- Human-readable JSON is preferred.
- Simplicity and broad compatibility are priorities.

Examples:
- Public APIs
- Mobile applications
- Web frontends
- External integrations

---

## gRPC vs REST

| Feature | gRPC | REST |
|--------|------|------|
| Protocol | HTTP/2 | HTTP/1.1 (typically), HTTP/2 supported |
| Data format | Protocol Buffers (binary) | JSON (usually) |
| Performance | High | Moderate |
| Payload size | Small | Larger |
| Streaming | Native support | Limited (WebSockets/SSE usually required) |
| Type safety | Strong | Weaker |
| Code generation | Built-in | Usually manual or via OpenAPI |
| Browser support | Limited (requires gRPC-Web) | Excellent |

---

## Summary

- **gRPC** is ideal for internal service-to-service communication where performance, type safety, and streaming are important.
- **Protocol Buffers** provide compact, fast, and strongly typed binary serialization with automatic code generation.
- **REST** is a better choice for public APIs, browser clients, and integrations where JSON and wide compatibility are more important than maximum performance.