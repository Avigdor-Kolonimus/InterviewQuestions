# Practical Architecture

## 1. How do you decompose a monolith into microservices, and what are the advantages and disadvantages?

A **monolith** is an application where all functionality is deployed as a single unit.

A **microservice architecture** splits the application into multiple independent services, each responsible for a specific business capability.

### How to decompose a monolith

A common approach is to identify **bounded contexts** (business domains) using Domain-Driven Design (DDD).

Examples:

- User Service
- Order Service
- Payment Service
- Inventory Service
- Notification Service

Each microservice should:

- Own its database.
- Have a single responsibility.
- Be independently deployable.
- Expose APIs to other services.

Avoid splitting by technical layers (e.g., one service for controllers and another for repositories). Instead, split by **business capabilities**.

---

### Advantages

- Independent deployment
- Better scalability
- Fault isolation
- Smaller codebases
- Technology flexibility
- Easier ownership by individual teams

---

### Disadvantages

- Increased operational complexity
- More network communication
- Distributed transactions are difficult
- More monitoring and logging required
- Higher infrastructure costs

---

## 2. How do microservices communicate?

There are two main communication models.

### Synchronous communication

One service directly calls another and waits for a response.

Protocols:

- HTTP/REST
- gRPC

Advantages:

- Simple
- Immediate response

Disadvantages:

- Tight coupling
- Higher latency
- Failures can propagate between services

---

### Asynchronous communication

Services exchange events through a message broker.

Examples:

- Kafka
- RabbitMQ

Advantages:

- Loose coupling
- Better scalability
- Increased resilience

Disadvantages:

- More complex architecture
- Eventual consistency
- Harder debugging

---

### Typical architecture

```
Client
   │
API Gateway
   │
 ├── User Service
 ├── Order Service
 ├── Payment Service
 └── Notification Service
        │
     Kafka/RabbitMQ
```

---

## 3. What is API idempotency, and why is it important when retrying requests?

An operation is **idempotent** if executing it multiple times has the same effect as executing it once.

### HTTP methods

| Method | Idempotent |
|--------|------------|
| GET | ✅ |
| PUT | ✅ |
| DELETE | ✅ |
| POST | ❌ (usually) |
| PATCH | Depends |

Example:

```
DELETE /users/1
```

Calling it multiple times leaves the resource deleted.

---

### Why idempotency is important

Clients may retry requests due to:

- Network failures
- Timeouts
- Retries by load balancers
- Client reconnects

Without idempotency:

- Duplicate orders
- Duplicate payments
- Duplicate emails

may occur.

---

### Idempotency Key

A common solution is an **Idempotency-Key**.

```
POST /payments

Idempotency-Key: abc123
```

The server stores the key and returns the original response for repeated requests with the same key instead of processing the operation again.

---

## 4. How do you optimize PostgreSQL queries?

### Use indexes

Create indexes for frequently filtered or joined columns.

```sql
CREATE INDEX idx_users_email
ON users(email);
```

---

### Analyze execution plans

Use:

```sql
EXPLAIN
```

or

```sql
EXPLAIN ANALYZE
```

Check for:

- Sequential scans
- Index scans
- Execution time
- Estimated vs actual rows

---

### Select only required columns

Avoid:

```sql
SELECT *
```

Prefer:

```sql
SELECT id, name
FROM users;
```

---

### Optimize joins

- Join indexed columns.
- Avoid unnecessary joins.
- Filter rows before joining when possible.

---

### Use pagination

Instead of loading all rows:

```sql
LIMIT 100
OFFSET 0
```

For very large datasets, **keyset (cursor-based) pagination** is often more efficient than large `OFFSET` values.

---

### Keep statistics up to date

Run:

```sql
ANALYZE;
```

or

```sql
VACUUM ANALYZE;
```

so the query planner has accurate statistics.

---

### Optimize schema

- Choose appropriate data types.
- Normalize data to reduce redundancy.
- Denormalize selectively for read-heavy workloads.

---

### Avoid unnecessary work

- Eliminate redundant subqueries.
- Use proper filtering.
- Cache frequently requested data (e.g., Redis).

---

## Summary

- Split a monolith into microservices based on **business domains**, with each service owning its own data and responsibility.
- Use **HTTP/gRPC** for synchronous communication and **Kafka/RabbitMQ** for asynchronous, event-driven communication.
- **Idempotency** ensures that retrying a request does not create duplicate side effects, often implemented using an **Idempotency-Key**.
- Optimize PostgreSQL by using indexes, analyzing execution plans with `EXPLAIN ANALYZE`, selecting only needed data, optimizing joins, maintaining statistics, and designing efficient schemas.