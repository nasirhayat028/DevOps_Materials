# Kubernetes Pods

## What is a Pod?
A Pod is the smallest deployable unit in Kubernetes.

It can run:
- One container (most common)
- Multiple containers (tight coupling)

---

## Why Pods exist?
Kubernetes does not manage containers directly.

Pods solve:
- Networking (shared IP)
- Storage (shared volumes)
- Lifecycle management

---

## Pod Lifecycle Flow

kubectl run → API Server → Scheduler → Node → Kubelet → Container Runtime → Pod runs

---

## Basic Commands

kubectl get pods  
kubectl run nginx --image=nginx  
kubectl describe pod <pod-name>  
kubectl get pods -o wide  
kubectl delete pod <pod-name>

---

## YAML Example

apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
  - name: nginx
    image: nginx

---

## YAML Explanation (Line by Line)

apiVersion: v1  
→ Kubernetes API version (stable core objects)

kind: Pod  
→ Defines this object as a Pod

metadata:  
→ Information about the object (name, labels, etc.)

name: nginx-pod  
→ Name of the Pod

spec:  
→ Desired state of the Pod

containers:  
→ List of containers inside the Pod

name: nginx  
→ Container name

image: nginx  
→ Docker image to run

---

## Key Concepts

- Pod is ephemeral (temporary)
- All containers in a Pod share same network
- Pods are NOT directly used for scaling
- Use Deployments for production

---

## Real-world Example

- Web server (nginx)
- Logging sidecar container

Both run inside same Pod

---

## Interview Questions

Q: Why do we need Pods?  
A: To provide abstraction over containers for networking, storage, and lifecycle management.

Q: Can a Pod have multiple containers?  
A: Yes, if they are tightly coupled.