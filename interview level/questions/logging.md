# Logging

## 1. What is logging used for, and how is it applied?

**Logging** is the process of recording information about an application's execution. It helps developers monitor, debug, and troubleshoot applications in production.

### Common use cases

- Debugging application errors.
- Monitoring application behavior.
- Tracking requests and user actions.
- Auditing important events.
- Diagnosing performance issues.

### Common log levels

- **DEBUG** – Detailed information for debugging.
- **INFO** – General information about application events.
- **WARN** – Unexpected situations that are not errors.
- **ERROR** – Errors that affect part of the application.
- **FATAL** – Critical errors that cause the application to terminate.

### Best practices

- Use structured logging (e.g., JSON) for easier searching and analysis.
- Include contextual information such as request ID, user ID, and timestamps.
- Avoid logging sensitive information (passwords, tokens, personal data).
- Use appropriate log levels.
- Send logs to centralized logging systems (e.g., ELK Stack, Grafana Loki, Splunk, or Cloud Logging).

### Example (Go)

```go
log.Printf("User %d logged in", userID)
```

or with a structured logger:

```go
logger.Info("User logged in", "user_id", userID)
```