# System Design

## 1. What are the Transactional Outbox and Saga patterns, and what are they used for?

### Transactional Outbox

The **Transactional Outbox** pattern ensures that database updates and event publishing happen reliably without using distributed transactions.

### Problem

Suppose an Order Service:

1. Saves an order to PostgreSQL.
2. Publishes an event to Kafka.

If the database transaction commits but the Kafka publish fails, the system becomes inconsistent.

---

### Solution

Instead of publishing directly to Kafka:

1. Save the business data.
2. Save an event to an **Outbox** table **in the same database transaction**.
3. Commit the transaction.
4. A background worker reads the Outbox table and publishes events to Kafka.
5. After successful delivery, the event is marked as processed or deleted.

```
Client
   │
   ▼
Order Service
   │
   ├── Orders table
   └── Outbox table
        │
        ▼
 Outbox Worker
        │
        ▼
      Kafka
```

### Advantages

- No lost events
- No distributed transactions
- Eventual consistency
- Reliable event delivery

### Disadvantages

- Additional Outbox table
- Background worker required
- Slight increase in system complexity

---

## Saga Pattern

A **Saga** manages distributed transactions across multiple microservices.

Instead of one global transaction, it executes a sequence of local transactions.

If one step fails, previously completed steps are compensated.

Example:

```
Create Order
      │
Reserve Inventory
      │
Charge Payment
      │
Send Notification
```

If payment fails:

```
Cancel Inventory Reservation
Cancel Order
```

These are called **compensating transactions**.

---

### Saga implementations

#### Choreography

- Services communicate through events.
- No central coordinator.

Advantages:

- Loosely coupled
- Easy to extend

Disadvantages:

- Harder to understand and debug
- Event chains become complex

---

#### Orchestration

A dedicated **Saga Orchestrator** controls the workflow.

```
Orchestrator
     │
 ├── Order Service
 ├── Payment Service
 ├── Inventory Service
 └── Notification Service
```

Advantages:

- Centralized control
- Easier monitoring

Disadvantages:

- Orchestrator becomes an additional component.

---

## 2. Explain some common architectural patterns (Circuit Breaker, Retry, CQRS).

### Circuit Breaker

Prevents repeatedly calling an unhealthy service.

States:

- **Closed** → requests pass normally.
- **Open** → requests fail immediately.
- **Half-Open** → limited requests test whether the service has recovered.

Benefits:

- Prevents cascading failures.
- Reduces load on unhealthy services.
- Improves system resilience.

---

### Retry

Automatically retries failed operations.

Common strategies:

- Fixed delay
- Exponential backoff
- Exponential backoff with jitter

Retry is suitable only for **transient failures**, such as:

- Network timeouts
- Temporary service outages

Avoid retrying permanent errors (e.g., invalid input).

---

### CQRS (Command Query Responsibility Segregation)

Separates **write operations** from **read operations**.

```
          Client
         /      \
 Commands      Queries
     │            │
Write Model   Read Model
     │            │
 Database     Read Database
```

Benefits:

- Independent scaling of reads and writes.
- Optimized read models.
- Better performance for read-heavy systems.

Disadvantages:

- More infrastructure.
- Eventual consistency.
- Increased complexity.

---

## 3. What scaling patterns do you know (vertical and horizontal scaling)?

### Vertical Scaling (Scale Up)

Increase the resources of a single machine.

Examples:

- More CPU
- More RAM
- Faster storage

Advantages:

- Simple to implement.
- No application changes.

Disadvantages:

- Hardware limits.
- Single point of failure.
- Can become expensive.

---

### Horizontal Scaling (Scale Out)

Add more servers instead of upgrading one.

Usually combined with a load balancer.

```
          Load Balancer
         /      |      \
      Server1 Server2 Server3
```

Advantages:

- High availability.
- Fault tolerance.
- Virtually unlimited scalability.

Disadvantages:

- More complex architecture.
- Requires stateless services or distributed state management.
- Load balancing and data consistency become important.

---

## Vertical vs Horizontal Scaling

| Feature | Vertical Scaling | Horizontal Scaling |
|--------|------------------|--------------------|
| Add resources | Bigger server | More servers |
| Complexity | Low | Higher |
| Fault tolerance | Low | High |
| Scalability | Limited | Excellent |
| Cost | Can become expensive | Better long-term scalability |

---

## Summary

- **Transactional Outbox** guarantees reliable event publishing by storing events in the same database transaction as business data.
- **Saga** manages distributed transactions using local transactions and compensating actions instead of a global transaction.
- **Circuit Breaker** prevents repeated calls to failing services, **Retry** handles transient failures, and **CQRS** separates read and write models for better scalability.
- **Vertical scaling** upgrades a single machine, while **horizontal scaling** adds more machines, providing better scalability and fault tolerance.