# Advanced PostgreSQL

## 1. Explain the types of indexes in PostgreSQL and when to use each one (especially B-Tree). What are the pros and cons of each index type?

Indexes improve query performance by allowing PostgreSQL to find rows without scanning the entire table.

---

## B-Tree (default index)

**B-Tree** is the default and most commonly used index type in PostgreSQL.

It is a balanced tree where keys are stored in sorted order, allowing efficient lookups.

### Best for

- Equality searches (`=`)
- Range queries (`<`, `>`, `<=`, `>=`)
- `BETWEEN`
- `ORDER BY`
- `MIN()` / `MAX()`
- Prefix searches (`LIKE 'abc%'`)

Example:

```sql
CREATE INDEX idx_users_email
ON users(email);
```

### Time complexity

- Search: **O(log n)**
- Insert/Delete: **O(log n)**

### Advantages

- General-purpose index
- Supports sorting
- Excellent for equality and range queries
- Works for most workloads

### Disadvantages

- Not suitable for full-text search
- Not ideal for multi-dimensional data
- Larger storage overhead than some specialized indexes

---

## Hash Index

Optimized for equality comparisons only.

```sql
CREATE INDEX idx_hash
ON users USING HASH(email);
```

### Advantages

- Very fast for `=` lookups

### Disadvantages

- No range queries
- No ordering support
- Less versatile than B-Tree

---

## GiST (Generalized Search Tree)

Supports custom indexing strategies.

Commonly used for:

- Geospatial data (PostGIS)
- Range types
- Nearest-neighbor searches

### Advantages

- Flexible
- Supports many data types

### Disadvantages

- More complex
- Usually slower than B-Tree for simple lookups

---

## GIN (Generalized Inverted Index)

Designed for values containing multiple elements.

Common use cases:

- Full-text search
- JSONB
- Arrays

Example:

```sql
CREATE INDEX idx_json
ON documents
USING GIN(data);
```

### Advantages

- Excellent for JSON and arrays
- Fast full-text search

### Disadvantages

- Larger indexes
- Slower inserts and updates

---

## BRIN (Block Range Index)

Stores metadata about ranges of table pages instead of individual rows.

Best for very large tables where values are naturally ordered.

Examples:

- Timestamp
- Auto-increment IDs

### Advantages

- Very small indexes
- Fast to build
- Low maintenance

### Disadvantages

- Lower lookup precision
- Less effective on randomly distributed data

---

## Index Comparison

| Index | Best For | Advantages | Disadvantages |
|------|----------|------------|---------------|
| B-Tree | Equality, ranges, sorting | General-purpose, fast | Not for full-text or complex data |
| Hash | Equality only | Fast `=` lookups | No ranges or sorting |
| GiST | Spatial and custom data | Flexible | More complex, slower for simple queries |
| GIN | JSON, arrays, full-text | Excellent search performance | Large indexes, slower writes |
| BRIN | Huge ordered tables | Tiny indexes, fast creation | Lower precision |

---

## 2. How do transactions work in PostgreSQL? What transaction types exist? What are ACID and transaction isolation levels?

A **transaction** is a sequence of SQL operations executed as a single unit.

```sql
BEGIN;

UPDATE accounts
SET balance = balance - 100
WHERE id = 1;

UPDATE accounts
SET balance = balance + 100
WHERE id = 2;

COMMIT;
```

If an error occurs:

```sql
ROLLBACK;
```

---

## ACID Properties

### Atomicity

Either all operations succeed or none do.

---

### Consistency

A transaction moves the database from one valid state to another.

---

### Isolation

Concurrent transactions should not interfere incorrectly with each other.

---

### Durability

Once committed, changes survive crashes and restarts.

---

## Isolation Levels

### Read Uncommitted

- Dirty reads allowed by the SQL standard.
- PostgreSQL treats this as **Read Committed**.

---

### Read Committed (default)

- Prevents dirty reads.
- Non-repeatable reads and phantom reads are possible.

---

### Repeatable Read

- Prevents dirty and non-repeatable reads.
- Uses MVCC snapshots.
- Phantom reads are prevented in PostgreSQL's implementation.

---

### Serializable

- Highest isolation level.
- Transactions behave as if executed one at a time.
- May fail with serialization errors that require retrying.

---

## 3. What are sharding, partitioning, and replication, and when are they used?

### Partitioning

Splits a **single table** into smaller partitions.

Examples:

- By date
- By ID
- By region

Benefits:

- Faster queries
- Easier maintenance
- Smaller indexes

---

### Sharding

Splits the **entire database** across multiple servers.

Each shard stores part of the data.

Benefits:

- Horizontal scaling
- Increased storage capacity
- Higher write throughput

Trade-offs:

- More complex queries
- Cross-shard joins are difficult
- Operational complexity

---

### Replication

Copies data from one server to one or more replicas.

Types:

- **Streaming replication**
- **Logical replication**

Benefits:

- High availability
- Read scaling
- Disaster recovery

---

## Comparison

| Feature | Partitioning | Sharding | Replication |
|--------|--------------|----------|-------------|
| Scope | Table | Database | Database copy |
| Main goal | Performance | Horizontal scaling | Availability |
| Data duplication | No | No | Yes |

---

## 4. How do you use `EXPLAIN` and `EXPLAIN ANALYZE` to analyze query performance, and what should you look for in the execution plan?

### EXPLAIN

Shows the query execution plan without executing the query.

```sql
EXPLAIN
SELECT *
FROM users
WHERE id = 100;
```

---

### EXPLAIN ANALYZE

Executes the query and displays actual execution statistics.

```sql
EXPLAIN ANALYZE
SELECT *
FROM users
WHERE id = 100;
```

---

## What to look for

### Sequential Scan

```
Seq Scan
```

- Entire table is scanned.
- May indicate a missing or unused index.

---

### Index Scan

```
Index Scan
```

- Uses an index.
- Usually much faster for selective queries.

---

### Bitmap Index Scan

Useful when many rows match and combining index lookups with heap access is more efficient.

---

### Estimated vs Actual Rows

Compare:

```
rows=100
actual rows=10000
```

Large differences suggest outdated statistics.

Run:

```sql
ANALYZE;
```

---

### Cost

Example:

```
cost=0.29..8.31
```

Lower estimated cost generally indicates a cheaper execution plan.

---

### Execution Time

```
Execution Time: 12 ms
```

The most important metric for real performance.

---

### Buffers (optional)

```sql
EXPLAIN (ANALYZE, BUFFERS)
```

Shows:

- Shared buffer hits
- Disk reads
- Cache efficiency

---

## Summary

- **B-Tree** is the default index and the best choice for most equality, range, and sorting queries.
- PostgreSQL transactions provide **ACID** guarantees and support multiple isolation levels using **MVCC**.
- **Partitioning** improves performance for large tables, **sharding** enables horizontal scaling across servers, and **replication** increases availability and read scalability.
- Use **EXPLAIN** to inspect the execution plan and **EXPLAIN ANALYZE** to measure actual performance. Focus on scan types, row estimates, execution time, and whether indexes are being used effectively.