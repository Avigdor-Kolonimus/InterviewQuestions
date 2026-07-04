# Kafka

## 1. Explain the basic concepts of Kafka: brokers, topics, partitions, producers, and consumers.

### Broker

A **broker** is a Kafka server that stores data and serves client requests.

A Kafka cluster typically consists of multiple brokers for scalability and fault tolerance.

---

### Topic

A **topic** is a logical category or stream of messages.

Examples:

- `orders`
- `payments`
- `notifications`

Producers write messages to topics, and consumers read from them.

---

### Partition

A **partition** is a subset of a topic.

Messages within a partition are **ordered**, while ordering is **not guaranteed across partitions**.

Partitions enable:
- Horizontal scalability
- Parallel processing
- Higher throughput

---

### Producer

A **producer** publishes messages to Kafka topics.

A producer can choose the partition:
- Automatically
- By message key (same key → same partition)
- Explicitly

---

### Consumer

A **consumer** reads messages from topics.

Consumers are usually organized into **consumer groups**.

Within a consumer group:
- Each partition is consumed by only one consumer.
- Different groups can independently consume the same topic.

---

## 2. How can you guarantee "exactly once", "at least once", or "at most once" message delivery?

### At Most Once

- Offset is committed **before** processing.
- Messages may be lost.
- No duplicates.

```
Commit offset → Process message
```

---

### At Least Once

- Process the message first.
- Commit the offset only after successful processing.

```
Process message → Commit offset
```

If a consumer crashes before committing the offset, the message will be processed again.

- No message loss.
- Duplicate processing is possible.

---

### Exactly Once

Kafka supports **Exactly Once Semantics (EOS)** using:

- Idempotent producers
- Transactions
- Transaction-aware consumers

Requirements:
- `enable.idempotence=true`
- Transactions enabled
- Consumers read only committed transactions (`isolation.level=read_committed`)

This prevents both duplicates and message loss.

---

## 3. Why is replication needed in Kafka, and how does it affect availability and consistency?

Kafka replicates partitions across multiple brokers.

Example:

```
Partition 0
 ├── Leader
 ├── Follower
 └── Follower
```

Only the **leader** handles reads and writes.

Followers continuously replicate data from the leader.

### Benefits

- High availability
- Fault tolerance
- Data durability

If the leader fails:

- One of the in-sync replicas (ISR) becomes the new leader.
- Producers and consumers continue working with minimal interruption.

### Trade-off

More replicas improve durability and availability but increase storage usage and replication traffic.

---

## 4. What happens to Kafka data after a system reboot, and why?

Kafka stores messages **on disk**, not only in memory.

Therefore, after restarting:

- Messages remain available.
- Consumers can continue from their committed offsets.

Data is retained according to the topic's **retention policy**, for example:

- Time-based retention
- Size-based retention

Kafka does **not** delete messages immediately after they are consumed.

---

## 5. Can Kafka be used for communication between microservices? How does it differ from HTTP?

Yes. Kafka is commonly used for **asynchronous communication** between microservices.

### Kafka

- Asynchronous messaging
- Producer and consumer are loosely coupled
- High throughput
- Durable message storage
- Retry support
- Event-driven architecture

Examples:
- Order processing
- Notifications
- Analytics
- Audit logs
- Event sourcing

---

### HTTP

- Synchronous request-response
- Client waits for the server's response
- Tight coupling between services
- No built-in message persistence
- Better suited for immediate responses

Examples:
- REST APIs
- Authentication
- CRUD operations
- User-facing requests

---

## Kafka vs HTTP

| Feature | Kafka | HTTP |
|--------|-------|------|
| Communication | Asynchronous | Synchronous |
| Coupling | Loose | Tight |
| Message persistence | Yes | No |
| Retry support | Built-in | Application-level |
| Throughput | Very high | Moderate |
| Ordering | Per partition | Not guaranteed |
| Scalability | Excellent | Good |

---

## Summary

- **Broker** stores data and serves clients.
- **Topic** is a stream of messages.
- **Partition** enables scalability and ordering within a partition.
- **Producer** publishes messages.
- **Consumer** reads messages, often as part of a consumer group.
- Delivery guarantees:
  - **At most once** → no duplicates, possible message loss.
  - **At least once** → no message loss, possible duplicates.
  - **Exactly once** → no duplicates and no message loss (using Kafka transactions and idempotent producers).
- **Replication** improves availability, durability, and fault tolerance.
- Kafka stores data on disk, so messages survive restarts according to the retention policy.
- Kafka is ideal for asynchronous, event-driven communication between microservices, while HTTP is better suited for synchronous request-response interactions.