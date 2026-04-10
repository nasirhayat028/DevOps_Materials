# 🚀 KUBERNETES DEEP MASTER NOTES (STEP-BY-STEP UNDERSTANDING)

---

# 🧠 PART 1 — WHAT IS KUBERNETES (REAL UNDERSTANDING)

---

## ❓ Problem Before Kubernetes

Imagine:

You built a web app:

* Backend (Node.js)
* Database (MySQL)

You run using Docker.

Now real problems start:

❌ Server crash → app down
❌ Traffic increase → app slow
❌ Manual restart → time waste
❌ Scaling manually → messy
❌ Different environments → inconsistency

---

## ✅ Solution: Kubernetes

Kubernetes = **container management system**

👉 It automatically:

* Runs containers
* Restarts if crash (self-healing)
* Scales up/down
* Distributes load
* Manages networking

---

## 🏢 Real Life Example (IMPORTANT)

Think of Kubernetes like:

👉 **Factory Manager**

* Containers = workers
* Pods = worker groups
* Nodes = machines
* Kubernetes = manager

Manager ensures:

* Workers always running
* If worker dies → replace
* If work increases → hire more

---

# 🧱 PART 2 — CORE BUILDING BLOCKS

---

## 1. POD (MOST IMPORTANT)

### 🧠 What:

Pod = smallest unit in Kubernetes

👉 Usually contains:

* 1 container

---

### ❓ Why:

Kubernetes never runs container directly
👉 Always inside Pod

---

### 🌍 Real Example:

You deploy:

* nginx server

👉 Kubernetes creates:

* Pod → inside it nginx container

---

---

## 2. NODE

### 🧠 What:

Node = machine (VM or server)

---

### 🌍 Example:

* AWS EC2
* Local VM

---

---

## 3. CLUSTER

### 🧠 What:

Cluster = group of nodes

---

---

## 4. DEPLOYMENT

### 🧠 What:

Manages:

* Pods
* Scaling
* Updates

---

---

## 5. SERVICE

### 🧠 What:

Provides stable network access

---

---

# ⚡ PART 3 — YOUR COMMANDS (DEEP BREAKDOWN)

---

# 🔹 1. VIEW PODS

```bash
kubectl get pods
```

---

### 🧠 What:

Shows all running pods

---

### ⚙️ Behind:

kubectl → API Server → etcd → response

---

### 🌍 Real Use:

* Check app health
* Debug issues

---

---

# 🔹 2. CREATE POD

```bash
kubectl run nginx --image nginx
```

---

### 🧠 What:

Creates Pod with nginx container

---

### ⚙️ Behind:

* Pod created
* Scheduler assigns node
* Image pulled
* Container started

---

### 🌍 Real Example:

Deploy quick test server

---

---

# 🔹 3. WRONG IMAGE CASE

```bash
kubectl run redis --image redis123
```

---

### 🧠 What happens:

❌ Image not found

---

### Result:

Pod → `ImagePullBackOff`

---

### 💡 Learning:

Always use valid image

---

---

# 🔹 4. DESCRIBE POD

```bash
kubectl describe pod <name>
```

---

### 🧠 What:

Full details

---

### Includes:

* Image
* Events
* Errors
* Node

---

### 🌍 Use:

Debugging tool

---

---

# 🔹 5. EXTENDED VIEW

```bash
kubectl get pods -o wide
```

---

### 🧠 Extra Info:

* Node name
* Pod IP

---

---

# 🔹 6. DELETE POD

```bash
kubectl delete pod webapp
```

---

### 🧠 What:

Deletes pod

---

### ⚠️ Important:

If controlled by deployment → auto recreated

---

---

# 🔹 7. UPDATE IMAGE

```bash
kubectl set image pod redis redis=redis:latest 
```

---

### 🧠 What:

Updates container image

---

### ⚠️ Reality:

Pods are not meant to be updated manually
👉 Use Deployment instead

---

---

# 🔥 PART 4 — REPLICATION CONTROLLER (VERY IMPORTANT)

---

## YAML:

```yaml
replicas: 3
```

---

### 🧠 Concept:

Maintain 3 pods always

---

### ⚙️ Behind:

Loop runs continuously:

IF pods < 3 → create
IF pods > 3 → delete

---

### 🌍 Real Example:

Your app needs:

* 3 instances for traffic

---

---

## FULL FLOW:

```bash
kubectl apply -f replicases.yml
```

👉 Creates controller + pods

---

```bash
kubectl get replicationcontroller
```

👉 Check status

---

```bash
kubectl delete replicationcontroller my-resume
```

👉 Deletes all pods

---

---

# 🔥 PART 5 — DEPLOYMENT vs REPLICASET

---

### Flow:

Deployment → ReplicaSet → Pods

---

### 🧠 Why Deployment:

* Rolling updates
* Rollback
* Scaling

---

---

# 🔥 PART 6 — SERVICES (NETWORKING)

---

```bash
kubectl get svc
kubectl create -f service.yml
```

---

### 🧠 Problem:

Pods IP changes

---

### Solution:

Service gives:
👉 Stable IP / DNS

---

### 🌍 Real Example:

User → Service → Pod

---

---

# 🔥 PART 7 — NAMESPACES

---

```bash
kubectl get pods --namespace=research
```

---

### 🧠 What:

Logical separation

---

### 🌍 Example:

* dev
* testing
* prod

---

---

# 🔥 PART 8 — IMPERATIVE vs DECLARATIVE

---

### Imperative:

Quick commands

---

### Declarative:

YAML based (production)

---

👉 Always prefer:
✔ Declarative

---

---

# 🔥 PART 9 — LABELS & SELECTORS

---

```bash
kubectl get pods --selector app=App1
```

---

### 🧠 Concept:

Labels = tags
Selectors = filters

---

---

# 🔥 PART 10 — TAINTS & TOLERATIONS

---

## 🧠 Problem:

Control which pod goes where

---

## Taint:

```bash
kubectl taint nodes node01 spray=mortein:NoSchedule
```

👉 Block pods

---

## Toleration:

Pod allows scheduling

---

### 🌍 Real Example:

* DB node only for DB pods

---

---

# 🔥 PART 11 — NODE LABELS

---

```bash
kubectl label node node01 color=blue
```

---

### 🧠 Use:

Assign workload to specific node

---

---

# 🔥 PART 12 — RESOURCES

---

### 🧠 Why:

Control CPU / Memory

---

### 🌍 Example:

Prevent one app from consuming all RAM

---

---

# 🔥 PART 13 — PRIORITY CLASSES

---

### 🧠 What:

Important pods run first

---

### 🌍 Example:

* Payment system → high priority
* Logs → low priority

---

---

# 🔥 PART 14 — MONITORING

---

```bash
kubectl top pod
```

---

### 🧠 What:

Shows CPU / memory usage

---

---

# 🔥 PART 15 — LOGGING

---

```bash
kubectl logs -f pod
```

---

### 🧠 What:

See live logs

---

---

# 🔥 PART 16 — ROLLING UPDATES

---

```bash
kubectl rollout undo deployment/nginx
```

---

### 🧠 What:

Rollback version

---

---

# 🔥 PART 17 — CONFIGMAP

---

### 🧠 What:

Store non-sensitive config

---

---

# 🔥 PART 18 — SECRETS

---

### 🧠 What:

Store sensitive data

---

---

# 🧠 FINAL UNDERSTANDING (VERY IMPORTANT)

---

Kubernetes is NOT about commands.

It is about:

👉 Desired State System

You say:
"I want 3 pods running"

Kubernetes ensures:
✔ Always 3 running

---

# 🚀 YOUR CURRENT LEVEL

You now understand:

✔ Pods
✔ Controllers
✔ Deployment
✔ Services
✔ Scheduling
✔ Configs

👉 You are entering **real DevOps zone**

---

# ⚡ NEXT STEP (IMPORTANT)

Now we go deeper:

👉 Send ONE topic:

* Deployment YAML
* Service YAML
* Real app architecture

I will:

* Break it line-by-line
* Show production thinking
* Give real-world scenario

---



# 🚀 KUBERNETES DEEP MASTER NOTES (ADVANCED + COMPLETE)

---

# 🧠 1. REPLICATION CONTROLLER (VERY IMPORTANT)

## 🧠 What is it?

ReplicationController ensures ke **fixed number of Pods hamesha running rahen**.

## ❓ Why needed?

Real world:

* Pod crash ho gaya
* Node down ho gaya

👉 Without controller = downtime
👉 With controller = auto recovery

## ⚙️ How it works?

Continuous loop:

* Desired = 3
* Current = 2
  👉 Create 1 new pod

## 🌍 Real Example

E-commerce backend:

* 3 instances running
* 1 crash → auto new pod

## 📦 Commands

```
kubectl apply -f replicases.yml
kubectl get replicationcontroller
kubectl delete replicationcontroller my-resume
```

---

# 🧠 2. DEPLOYMENT vs REPLICASET

## ReplicaSet

* Maintains number of pods

## Deployment

* Manages ReplicaSet
* Adds features:

  * Rolling updates
  * Rollback
  * Versioning

## ⚙️ Flow

Deployment → ReplicaSet → Pods

## 🌍 Real Example

App v1 → update to v2
Deployment replaces pods gradually

## 📦 Commands

```
kubectl get deployments
kubectl get replicasets
kubectl describe deployment <name>
```

---

# 🧠 3. SERVICES (NETWORKING)

## ❓ Problem

Pods IP changes

## ✅ Solution

Service = stable IP/DNS

## Types

* ClusterIP
* NodePort
* LoadBalancer

## 🌍 Real Example

User → Service → Pod

## 📦 Commands

```
kubectl get svc
kubectl create -f service.yml
kubectl describe service
```

---

# 🧠 4. NAMESPACES

## What?

Logical isolation

## Example

* dev
* staging
* prod

## Commands

```
kubectl get namespaces
kubectl get pods -n research
kubectl get pods --all-namespaces
```

---

# 🧠 5. IMPERATIVE vs DECLARATIVE

## Imperative

Quick commands

```
kubectl run nginx
```

## Declarative

Production standard

```
kubectl apply -f file.yaml
```

---

# 🧠 6. LABELS & SELECTORS

## Labels

Key-value tags

## Selectors

Filter resources

## Commands

```
kubectl get pods --selector app=frontend
kubectl get all --selector env=prod
```

---

# 🧠 7. TAINTS & TOLERATIONS

## Taint

Block pods from node

## Toleration

Allow specific pods

## Commands

```
kubectl taint nodes node01 spray=mortein:NoSchedule
kubectl describe node node01
```

---

# 🧠 8. NODE LABELS

## What?

Assign workloads to nodes

## Commands

```
kubectl label node node01 color=blue
kubectl get nodes --show-labels
```

---

# 🧠 9. RESOURCES

## What?

Limit CPU & RAM

## Why?

Prevent resource exhaustion

---

# 🧠 10. PRIORITY CLASSES

## What?

Important pods run first

## Commands

```
kubectl get priorityclasses
kubectl apply -f high-priority.yaml
```

---

# 🧠 11. MONITORING

## What?

Track CPU & Memory

## Commands

```
kubectl top node
kubectl top pod
```

---

# 🧠 12. LOGGING

## What?

See application logs

## Commands

```
kubectl logs pod-name
kubectl logs -f pod-name
```

---

# 🚀 FINAL MINDSET

Kubernetes = Desired State Engine

You define → Kubernetes maintains

---

# 🚀 NEXT STEP

Ask for:

* Deployment YAML deep dive
* Service YAML deep dive

We go production level next 🔥
    