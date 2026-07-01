# Redis

## 1. How is Redis different from a relational database?

Redis is an **in-memory key-value database**, while a relational database (e.g., PostgreSQL, MySQL) stores data in **tables with predefined schemas**.


| Redis                                         | Relational Database                       |
| --------------------------------------------- | ----------------------------------------- |
| In-memory storage (very fast)                 | Disk-based storage                        |
| Key-value data model                          | Tables, rows, and columns                 |
| Schema-less                                   | Fixed schema                              |
| Optimized for caching and real-time workloads | Optimized for persistent relational data  |
| Limited querying capabilities                 | Powerful SQL queries, joins, transactions |


Redis is commonly used for:

- Caching
- Session storage
- Rate limiting
- Queues
- Leaderboards

---



## 2. What is TTL, and how should it be used in Redis?

**TTL (Time To Live)** is the amount of time a key remains in Redis before it is automatically deleted.

Example:

```bash
SET user:1 "Alice" EX 60
```

or

```bash
EXPIRE user:1 60
```

Useful commands:

```bash
TTL key        # Remaining lifetime
EXPIRE key 60  # Set expiration
PERSIST key    # Remove expiration
```



### Best practices

- Use TTL for **cache entries** to prevent stale data.
- Set expiration when creating the key whenever possible (`SET ... EX`).
- Choose TTL values based on how frequently the data changes.
- Avoid storing cache entries without expiration unless necessary.

---



## 3. How do you clear the Redis cache?



### Delete a single key

```bash
DEL mykey
```



### Delete multiple keys

```bash
DEL key1 key2 key3
```



### Clear the current database

```bash
FLUSHDB
```



### Clear all databases

```bash
FLUSHALL
```



### Asynchronous deletion (recommended for production)

```bash
FLUSHDB ASYNC
```

or

```bash
FLUSHALL ASYNC
```

The `ASYNC` option performs deletion in the background, reducing blocking for large datasets.