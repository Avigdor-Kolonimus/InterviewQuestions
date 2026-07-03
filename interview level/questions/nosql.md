# NoSQL

## 1. Explain how NoSQL databases differ from relational databases and when they are preferable.

**NoSQL** databases are non-relational databases designed to handle large volumes of data, flexible schemas, and horizontal scalability.

### Differences between NoSQL and relational databases

| NoSQL | Relational Database |
|-------|----------------------|
| Schema-less or flexible schema | Fixed schema |
| Data stored as documents, key-value pairs, columns, or graphs | Data stored in tables (rows and columns) |
| Optimized for horizontal scaling | Typically scales vertically |
| Usually provides eventual consistency (depending on the database) | Strong consistency and ACID transactions |
| Limited joins | Powerful SQL queries and joins |

---

## Types of NoSQL databases

- **Key-value** (Redis)
- **Document** (MongoDB)
- **Column-family** (Cassandra)
- **Graph** (Neo4j)

---

## When to use NoSQL

NoSQL databases are preferred when:

- The data structure changes frequently.
- High read/write throughput is required.
- Horizontal scaling across many servers is needed.
- The application handles large volumes of semi-structured or unstructured data.
- Low latency is critical (e.g., caching).

Typical use cases:

- Caching (Redis)
- Session storage
- Real-time analytics
- Social networks
- IoT applications
- Content management systems
- Big data applications

---

## When to use a relational database

Relational databases are a better choice when:

- Data has a well-defined structure.
- Complex SQL queries and joins are required.
- Strong consistency and ACID transactions are essential.
- Data integrity is a top priority.

Typical use cases:

- Banking systems
- Financial applications
- E-commerce orders and payments
- ERP/CRM systems

---

## Summary

- **NoSQL** databases provide flexible schemas, high performance, and horizontal scalability, making them ideal for large-scale and rapidly changing data.
- **Relational databases** provide structured data storage, strong consistency, and powerful querying capabilities, making them ideal for transactional systems.
```