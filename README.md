---

## 📦 Docker Image Optimization Best Practices

---

### 🏗️ 1. Use a Minimal Base Image

* Start with a lightweight base to reduce image size.
* Prefer `alpine`, `busybox`, or `scratch` when possible.

**Example:**

```dockerfile
FROM alpine
```

---

### 🔄 2. Use Multi-Stage Builds

* Separate build and runtime environments to include only what’s needed.

**Example:**

```dockerfile
# Build Stage
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

# Final Image
FROM alpine
COPY --from=builder /app/myapp /usr/local/bin/myapp
CMD ["myapp"]
```

---

### 🧹 3. Remove Unnecessary Files

* Clean up package manager caches, temp files, and build artifacts.

**Example:**

```dockerfile
RUN apt-get update && apt-get install -y curl \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* /tmp/*
```

---

### 🧾 4. Combine RUN Instructions

* Merge related `RUN` commands to reduce intermediate image layers.

**Example:**

```dockerfile
RUN apk add --no-cache curl bash \
 && rm -rf /var/cache/apk/*
```

---

### 📂 5. Use `.dockerignore`

* Prevent unnecessary files (e.g. `.git`, `node_modules`, logs) from being added to the image build context.

**Example `.dockerignore`:**

```
.git
node_modules
*.log
Dockerfile.dev
```

---

### 🧰 6. Avoid Installing Unnecessary Packages

* Install only essential dependencies required at runtime.
* Don’t include build tools or editors in production images.

---

### 📤 7. Use Specific COPY/ADD Paths

* Avoid copying the entire project directory unless needed.

**Bad:**

```dockerfile
COPY . .
```

**Good:**

```dockerfile
COPY app.py /app/
COPY requirements.txt /app/
```

---

### 📉 8. Minimize Layers

* Reduce the total number of image layers for better caching and efficiency.
* Merge related operations when appropriate.

---
