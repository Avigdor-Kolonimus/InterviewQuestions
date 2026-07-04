# CI/CD

## 1. What is CI/CD and why is it needed in software development?

**CI/CD** stands for:

- **CI (Continuous Integration)** – continuous integration
- **CD (Continuous Delivery / Continuous Deployment)** – continuous delivery or deployment

### CI (Continuous Integration)

CI is a practice where developers frequently merge code changes into a shared repository.

Each change is automatically:
- built
- tested
- validated

**Goal:**
Detect bugs early and avoid integration issues.

---

### CD (Continuous Delivery / Deployment)

CD ensures that code changes are automatically prepared for release or even deployed to production.

- **Continuous Delivery** → code is always ready to be deployed (manual approval required)
- **Continuous Deployment** → every passing change is automatically deployed to production

---

### Why CI/CD is needed

- Faster and safer releases
- Early bug detection
- Reduced manual work
- Consistent build and deployment process
- Better collaboration in teams

---

## 2. What are the main stages in a typical CI/CD pipeline?

A standard CI/CD pipeline usually includes:

### 1. Source stage
- Code is pushed to a repository (GitHub, GitLab, etc.)

### 2. Build stage
- Application is compiled or built (e.g., Go build, Docker image build)

### 3. Test stage
- Unit tests, integration tests, linting

### 4. Analysis stage (optional)
- Code quality checks (SonarQube, static analysis)

### 5. Packaging stage
- Build artifacts or Docker images are created

### 6. Deploy stage
- Deployment to staging or production environment

### 7. Monitoring stage
- Logs, metrics, alerts (Prometheus, Grafana, etc.)

---

## 3. What is the difference between GitLab and GitHub? Why do companies use GitLab?

### GitHub

- Most popular Git-based platform
- Strong open-source ecosystem
- GitHub Actions for CI/CD
- Easier collaboration for public projects

### GitLab

- All-in-one DevOps platform
- Built-in CI/CD (very powerful GitLab CI)
- More advanced built-in DevOps features
- Better support for self-hosting (on-premise GitLab)

---

### Key differences

| Feature | GitHub | GitLab |
|--------|--------|--------|
| CI/CD | GitHub Actions | GitLab CI (more integrated) |
| DevOps tools | External integrations | Built-in |
| Self-hosting | Limited | Strong support |
| Popularity | Higher | Lower but strong in enterprise |

---

### Why companies often choose GitLab

- Full DevOps lifecycle in one tool
- Strong built-in CI/CD system
- Better control for private/self-hosted infrastructure
- More enterprise-oriented features out of the box
- Easier compliance and security control

---

### Summary

- **CI/CD** automates integration, testing, and deployment.
- CI ensures code correctness continuously.
- CD automates release delivery or deployment.
- GitLab is often preferred in enterprises due to its all-in-one DevOps approach.