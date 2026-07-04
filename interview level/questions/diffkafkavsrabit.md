# Kafka vs RabbitMQ

## 1. What are the main differences between the queue and topic models in Kafka and RabbitMQ?

### Kafka

Kafka is a **distributed event streaming platform** designed for high-throughput data processing.

Characteristics:
- Messages are stored in **topics**, which are divided into **partitions**.
- Messages remain in Kafka after being consumed until the retention period expires.
- Consumers track their own offsets.
- Multiple consumer groups can independently read the same messages.
- Optimized for scalability and throughput.

---

### RabbitMQ

RabbitMQ is a **message broker** implementing the **AMQP** protocol.

Characteristics:
- Producers publish messages to **exchanges**.
- Exchanges route messages to one or more **queues**.
- Consumers read messages from queues.
- Messages are typically removed from the queue after they are acknowledged.
- Optimized for reliable message delivery and flexible routing.

---

## Kafka vs RabbitMQ Architecture

### Kafka

```
Producer
    │
    ▼
  Topic
    │
Partitions
    │
Consumer Groups
```

---

### RabbitMQ

```
Producer
    │
    ▼
 Exchange
    │
 Routing
    ▼
 Queue(s)
    │
Consumer(s)
```

---

## Key differences

| Feature | Kafka | RabbitMQ |
|--------|-------|----------|
| Data model | Topics + Partitions | Exchanges + Queues |
| Message storage | Persistent (until retention expires) | Usually removed after acknowledgment |
| Consumer state | Consumer manages offsets | Broker tracks acknowledgments |
| Throughput | Very high | High |
| Ordering | Guaranteed within a partition | Guaranteed within a queue |
| Replay messages | Yes | Typically no |
| Routing | Simple (topic/partition) | Very flexible (direct, topic, fanout, headers exchanges) |

---

## 2. When should you choose RabbitMQ instead of Kafka, and vice versa?

### Choose RabbitMQ when:

- Reliable task distribution is required.
- Complex routing logic is needed.
- Immediate processing is expected.
- Message replay is unnecessary.
- Work queues or RPC-style communication are used.

Examples:
- Background jobs
- Email processing
- Payment processing
- Task queues
- Workflow orchestration

---

### Choose Kafka when:

- High throughput is required.
- Events must be retained and replayable.
- Multiple services need to consume the same events independently.
- Building event-driven architectures.
- Processing real-time data streams.

Examples:
- Event sourcing
- Audit logs
- Analytics pipelines
- Microservice event buses
- Log aggregation
- Real-time monitoring

---

## RabbitMQ vs Kafka

| Feature | RabbitMQ | Kafka |
|--------|----------|-------|
| Primary use | Message broker | Event streaming platform |
| Communication | Queue-based | Log-based |
| Message replay | No | Yes |
| Event retention | Limited | Configurable retention |
| Throughput | High | Very high |
| Routing capabilities | Excellent | Basic |
| Scalability | Good | Excellent |

---

## Summary

- **RabbitMQ** is best for reliable message delivery, task queues, and complex routing between producers and consumers.
- **Kafka** is best for high-throughput event streaming, durable event storage, and scenarios where multiple consumers need to process the same events independently.
- Choose **RabbitMQ** for traditional messaging and work queues.
- Choose **Kafka** for event-driven systems, streaming, analytics, and large-scale distributed architectures.