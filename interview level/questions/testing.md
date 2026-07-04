# Testing

## 1. What are unit, mock, and integration tests? How do they differ?

### Unit Tests

A **unit test** verifies the behavior of a single unit of code (usually a function or method) in isolation.

Characteristics:
- Fast to run.
- No external dependencies.
- Easy to write and maintain.
- Used to verify business logic.

Example:
Testing a function that calculates a discount without accessing a database or making HTTP requests.

---

### Mock Tests

A **mock test** is a type of unit test that replaces external dependencies with **mock objects** or **fake implementations**.

Mocks are used to simulate:
- Databases
- HTTP services
- Message queues
- File systems
- Third-party APIs

Benefits:
- Tests remain fast and deterministic.
- Allows testing error scenarios that are difficult to reproduce with real services.
- Verifies interactions with dependencies.

Example:
Mocking a `UserRepository` instead of connecting to a real database.

---

### Integration Tests

An **integration test** verifies that multiple components work together correctly.

Unlike unit tests, integration tests use real dependencies whenever possible.

Examples:
- Application + PostgreSQL
- Application + Redis
- HTTP API + Database
- Service-to-service communication

Characteristics:
- Slower than unit tests.
- More realistic.
- Validate configuration and interaction between components.

---

## Comparison

| Feature | Unit Test | Mock Test | Integration Test |
|---------|-----------|-----------|------------------|
| Tests a single unit | ✅ | ✅ | ❌ |
| Uses real dependencies | ❌ | ❌ | ✅ |
| Uses mocks | ❌ | ✅ | Sometimes |
| Fast | ✅ | ✅ | ❌ |
| Suitable for CI | ✅ | ✅ | ✅ (usually fewer tests) |
| Purpose | Verify business logic | Verify logic with mocked dependencies | Verify components work together |

### Rule of thumb

- **Unit tests** verify individual functions.
- **Mock tests** verify code that depends on external services without using the real services.
- **Integration tests** verify that the entire system or multiple components work together correctly.