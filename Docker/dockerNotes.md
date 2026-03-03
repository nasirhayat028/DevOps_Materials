# 🚀 Docker Complete Notes (Beginner → Advanced)

---

## 1️⃣ What is a Server?

A **server** is a computer that provides services to other computers (clients).

**Example:**  
When you open Instagram, your phone sends a request to Instagram's server.  
The server processes it and sends back images, videos, and data.

---

## 2️⃣ Server Components

- CPU (Processing)
- RAM (Temporary Memory)
- Disk (Storage)
- Network Interface
- Operating System (Linux in DevOps world)

---

## 3️⃣ Scaling a Server

### Vertical Scaling
Increase CPU/RAM of the same server.

### Horizontal Scaling
Add more servers behind a load balancer.

---

## 4️⃣ Linux Process Overview

Everything in Linux runs as a **process**.

Check running processes:

```bash
ps aux
```

Docker containers are isolated Linux processes.

---

## 5️⃣ What is a Container?

A container is:

> A lightweight isolated Linux process with its own filesystem and dependencies.

---

## 6️⃣ Container Features

- Lightweight
- Fast startup
- Portable
- Isolated
- Shares host kernel

---

## 7️⃣ Containerization

Packaging:
- Application
- Dependencies
- Runtime
- Configurations

Into one portable unit.

---

# 🐳 Docker Core Concepts

## 8️⃣ Docker Architecture

- Docker Client
- Docker Daemon
- Docker Images
- Docker Containers
- Docker Registry

---

## 9️⃣ Docker Pull

```bash
docker pull nginx
docker pull nginx:1.25
```

---

## 🔟 Docker Images

List images:

```bash
docker images
```

Images are read-only blueprints used to create containers.

---

## 1️⃣1️⃣ Docker PS

```bash
docker ps
docker ps -a
```

---

## 1️⃣2️⃣ Docker Run

```bash
docker run nginx
docker run -it ubuntu bash
docker run --rm ubuntu ls
```

- `-it` → Interactive terminal  
- `--rm` → Remove container after stop  

---

## 1️⃣3️⃣ Docker Exec

```bash
docker exec -it <container_id> bash
```

Used to run commands inside running containers.

---

## 1️⃣4️⃣ Docker Logs

```bash
docker logs <container_id>
docker logs -f <container_id>
```

---

# 🧱 Dockerfile

## Basic Structure

```Dockerfile
FROM ubuntu:22.04
WORKDIR /app
COPY . .
RUN apt-get update
ENV APP_ENV=production
EXPOSE 3000
CMD ["node", "app.js"]
```

### Explanation

- **FROM** → Base image  
- **WORKDIR** → Set working directory  
- **COPY** → Copy files  
- **RUN** → Execute command during build  
- **ENV** → Set environment variable  
- **EXPOSE** → Document port  
- **CMD** → Default start command  

---

## 🧩 Docker Layers

Each Dockerfile instruction creates a layer.

If one layer changes, Docker rebuilds only that layer and below it.  
This makes builds fast due to caching.

---

## 1️⃣5️⃣ Docker Build

```bash
docker build -t myapp:1.0 .
```

---

## 1️⃣6️⃣ Docker Tag

```bash
docker tag myapp:1.0 username/myapp:1.0
```

---

## 1️⃣7️⃣ Docker Push

```bash
docker push username/myapp:1.0
```

---

# 🌐 Docker Networking

## List Networks

```bash
docker network ls
docker network inspect bridge
```

### Network Modes
- bridge (default)
- host
- none
- overlay (Swarm)

`docker0` is the default Linux bridge interface.

---

# 💾 Docker Volumes

Create volume:

```bash
docker volume create mydata
```

Use volume:

```bash
docker run -v mydata:/data ubuntu
```

Bind mount:

```bash
docker run -v /host/path:/container/path ubuntu
```

- Volume → Managed by Docker  
- Bind Mount → Direct host mapping  

---

# 📦 Docker Compose

## docker-compose.yml

```yaml
version: "3.9"
services:
  web:
    image: nginx
    ports:
      - "8080:80"
  db:
    image: mysql
    environment:
      MYSQL_ROOT_PASSWORD: root
```

Run:

```bash
docker compose up
```

Stop:

```bash
docker compose down
```

---

# 🧠 Docker Internals

Docker uses:

- Linux Namespaces (Isolation)
- cgroups (Resource control)
- Union Filesystem (Layering)
- containerd
- runc

Container = Isolated Linux Process.

---

# 🎯 Final Understanding

Docker is NOT magic.

It is:
> Smart usage of Linux kernel features to isolate processes and package software.