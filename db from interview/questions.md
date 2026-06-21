# Database

### What are the advantages of document-oriented databases over relational databases?

**Document databases** (e.g., MongoDB) store data as documents (typically JSON/BSON) rather than rows and tables.

Advantages:

* Flexible schema (easy to add new fields).
* Better fit for hierarchical and nested data.
* Faster development when requirements change frequently.
* Often easier horizontal scaling.

Relational databases are usually better when:

* Data relationships are important.
* Strong consistency is required.
* Complex joins and transactions are common.

---

### What is a transaction? Why is it needed?

A transaction is a group of operations executed as a single unit of work.

A transaction guarantees that either:

* All operations succeed and are committed.
* None of them are applied (rollback).

Transactions are used to maintain data consistency.

Example: transferring money between accounts.

---

### What is ACID?

ACID is a set of guarantees provided by transactional databases:

* **Atomicity** – all operations succeed or all fail.
* **Consistency** – transactions preserve database integrity.
* **Isolation** – concurrent transactions do not interfere improperly.
* **Durability** – committed changes survive crashes.

---

### What transaction isolation levels exist? Why are they needed?

Isolation levels control how concurrent transactions interact.

#### Read Uncommitted

* Allows dirty reads.
* Rarely used.

#### Read Committed (PostgreSQL default)

* Prevents dirty reads.
* Non-repeatable reads are possible.

#### Repeatable Read

* Prevents dirty and non-repeatable reads.
* Prevents most phantom reads in PostgreSQL.

#### Serializable

* Highest isolation level.
* Transactions behave as if executed sequentially.
* May cause serialization failures and retries.

---

### Tell me about triggers. What are the pros and cons?

A trigger is code automatically executed when INSERT, UPDATE, or DELETE occurs.

Pros:

* Enforces business rules.
* Auditing and history tracking.
* Automatic data synchronization.

Cons:

* Hidden logic.
* Harder debugging.
* Performance overhead.
* Can create unexpected side effects.

---

### What is the difference between WHERE and HAVING?

**WHERE**

* Filters rows before grouping.

**HAVING**

* Filters groups after aggregation.

Example:

```sql
SELECT department_id, COUNT(*)
FROM employees
GROUP BY department_id
HAVING COUNT(*) > 10;
```

Can HAVING be used without GROUP BY?

Yes.

```sql
SELECT COUNT(*)
FROM employees
HAVING COUNT(*) > 100;
```

---

### What is a JOIN? What types exist? Can we avoid RIGHT JOIN?

JOIN combines rows from multiple tables.

Types:

* INNER JOIN
* LEFT JOIN
* RIGHT JOIN
* FULL OUTER JOIN
* CROSS JOIN
* SELF JOIN

RIGHT JOIN can always be rewritten as LEFT JOIN by swapping tables.

---

### Have you worked with query optimizers? How do you optimize queries?

Typical steps:

1. Use EXPLAIN ANALYZE.
2. Check index usage.
3. Reduce scanned rows.
4. Avoid unnecessary SELECT *.
5. Rewrite subqueries if needed.
6. Optimize JOIN order.
7. Consider denormalization.

---

### What can we do if a query has many JOINs?

Options:

* Add proper indexes.
* Reduce the number of joins.
* Denormalize frequently used data.
* Create materialized views.
* Precompute aggregates.
* Cache results.

---

### Why is denormalization needed?

Denormalization reduces the number of joins and improves read performance.

Trade-offs:

* Data duplication.
* More complicated writes.
* Risk of inconsistencies.

---

### What is normalization?

Normalization organizes data to reduce redundancy and anomalies.

Benefits:

* Less duplication.
* Better consistency.
* Easier maintenance.

---

### What is a normal form?

A normal form is a set of rules that reduce redundancy.

Common forms:

* 1NF
* 2NF
* 3NF
* BCNF

Most systems are designed around 3NF.

---

### Have you worked with sharding or partitioning?

**Partitioning**

* Splits a table into smaller pieces within one database.

**Sharding**

* Splits data across multiple database servers.

Partitioning improves manageability.
Sharding improves scalability.

---

### What PostgreSQL indexes do you know?

#### B-Tree

Default index type.

Used for:

* `=`
* `>`
* `<`
* BETWEEN
* ORDER BY

#### Hash

Optimized for equality searches.

#### GIN

Used for:

* JSONB
* Arrays
* Full-text search

#### GiST

Used for:

* Geospatial data
* Ranges

#### BRIN

Useful for huge append-only tables.

---

### Why is having many indexes bad?

Every INSERT, UPDATE, DELETE must update all affected indexes.

Problems:

* Slower writes.
* More storage.
* Longer VACUUM operations.
* Increased maintenance cost.

---

### A query against a partitioned table became slow over time. How would you investigate?

1. Compare EXPLAIN ANALYZE before and after.
2. Check whether partition pruning still works.
3. Check table statistics.
4. Run ANALYZE.
5. Look for data growth.
6. Check index usage.
7. Check for plan changes.

---

### What does an execution plan show?

Execution plan shows:

* How PostgreSQL executes a query.
* Scan types.
* Join methods.
* Cost estimates.
* Row estimates.

#### EXPLAIN

Shows the estimated plan.

#### EXPLAIN ANALYZE

Actually executes the query and shows:

* Real timings.
* Actual row counts.
* Actual execution plan.

---

### How would you choose a sharding key?

A good shard key should:

* Distribute data evenly.
* Avoid hotspots.
* Match common query patterns.
* Minimize cross-shard queries.

Typical examples:

* customer_id
* user_id
* tenant_id

---

### How have you used Redis?

Common use cases:

* Caching.
* Distributed locks.
* Rate limiting.
* Session storage.
* Queues.
* Pub/Sub.

Important concepts:

* TTL.
* Persistence.
* Replication.
* Clustering.

---

### Have you heard of MongoDB? When would you use it?

Good for:

* Flexible schemas.
* Rapid prototyping.
* Event data.
* Catalogs and content management.

Disadvantages:

* Weaker relational capabilities.
* More difficult joins.
* Possible data duplication.

---

### What is replication?

Replication copies data from a primary database to one or more replicas.

Benefits:

* High availability.
* Read scaling.
* Disaster recovery.

---

### What is sent to replicas?

In PostgreSQL physical replication:

* WAL (Write-Ahead Log) records are streamed.

Replicas replay WAL and reproduce changes.

---

### In synchronous replication, is data immediately available on the replica?

Yes.

The primary waits for confirmation from synchronous replicas before committing the transaction.

This provides stronger consistency but increases latency.

---

### Do you use foreign keys? What are they for?

Foreign keys enforce referential integrity.

Advantages:

* Prevent invalid references.
* Protect data consistency.

Disadvantages:

* Additional validation overhead.
* Can complicate large-scale distributed systems.
* May affect write performance.
* Can make some migrations more difficult.

# PostgreSQL Case Study

### Suppose we have four PostgreSQL databases, each using a different transaction isolation level.

Each database stores a counter, and 100 increment requests are sent to it concurrently.

What value will the counter have after all requests complete?

The answer depends on how the increment is implemented.

For a naïve read-modify-write pattern:

```sql
BEGIN;

SELECT value FROM counter;

UPDATE counter
SET value = value + 1;

COMMIT;
```

Results may vary because of lost updates.

For an atomic update:

```sql
UPDATE counter
SET value = value + 1;
```

all isolation levels in PostgreSQL will correctly produce:

```text
100
```

because row-level locking guarantees that increments are serialized safely.

This is usually the key point interviewers expect.
