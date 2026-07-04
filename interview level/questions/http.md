# HTTP/HTTPS

## 1. What is the difference between HTTP and HTTPS?

**HTTP (HyperText Transfer Protocol)** is a protocol for transferring data between a client and a server.

**HTTPS (HyperText Transfer Protocol Secure)** is HTTP running over **TLS**, which encrypts the communication.

### Differences

| HTTP | HTTPS |
|------|-------|
| Data is sent in plain text | Data is encrypted |
| No authentication | Server identity is verified using a certificate |
| Vulnerable to eavesdropping and tampering | Protects against eavesdropping and man-in-the-middle attacks |
| Default port: 80 | Default port: 443 |

---

## 2. What is TLS, and why is a certificate needed?

**TLS (Transport Layer Security)** is a cryptographic protocol that secures communication between a client and a server.

It provides:

- **Encryption** – prevents others from reading transmitted data.
- **Integrity** – ensures data has not been modified in transit.
- **Authentication** – verifies the server's identity.

### TLS Certificate

A **TLS certificate** is issued by a trusted **Certificate Authority (CA)**.

It contains:
- The server's public key.
- The server's identity (domain name).
- The CA's digital signature.

During the TLS handshake:
1. The server sends its certificate.
2. The client verifies that the certificate is valid and trusted.
3. A secure encrypted connection is established.

---

## 3. What is the difference between HTTP/1.1, HTTP/2, and HTTP/3?

### HTTP/1.1

- One request per connection at a time (pipelining exists but is rarely used).
- Multiple TCP connections are often needed.
- Text-based protocol.
- Suffers from head-of-line blocking at the application level.

---

### HTTP/2

- Uses **binary framing** instead of plain text.
- Supports **multiplexing** (multiple requests over a single TCP connection).
- Header compression (HPACK).
- Server Push (rarely used today).
- Lower latency and better network utilization than HTTP/1.1.

---

### HTTP/3

- Uses **QUIC** instead of TCP.
- QUIC runs over **UDP**.
- Eliminates TCP head-of-line blocking.
- Faster connection establishment.
- Better performance on unreliable or high-latency networks (e.g., mobile).

---

## HTTP Versions Comparison

| Feature | HTTP/1.1 | HTTP/2 | HTTP/3 |
|--------|----------|---------|---------|
| Transport | TCP | TCP | QUIC (UDP) |
| Multiplexing | ❌ | ✅ | ✅ |
| Header compression | ❌ | HPACK | QPACK |
| Binary protocol | ❌ | ✅ | ✅ |
| Connection setup | Slow | Faster | Fastest |
| Head-of-line blocking | Yes | At TCP layer | Eliminated at transport layer |

---

## Summary

- **HTTP** transmits data in plain text, while **HTTPS** secures communication using **TLS**.
- **TLS** provides encryption, integrity, and authentication through digital certificates.
- **HTTP/2** improves performance with multiplexing and binary framing.
- **HTTP/3** uses QUIC over UDP, reducing latency and eliminating TCP head-of-line blocking.