# Replication Controller (RC)

## What is Replication Controller?

A Replication Controller ensures that a specific number of Pods are always running.

👉 If a Pod crashes → it creates a new one automatically  
👉 If too many Pods exist → it removes extra ones  

This is called **self-healing + desired state management**

---

## Why do we need it?

Without RC:
- If a Pod dies → app goes down ❌
- If load increases → no scaling ❌

With RC:
- Always keeps app running ✅
- Maintains fixed number of replicas ✅

---

## Real-world DevOps Example

Imagine:

You run a shopping website (Amazon-like app)

You want:
- 3 copies of your web app always running

If 1 Pod crashes:
- RC immediately creates a new one
- users never see downtime

---

## How it works (Simple Flow)

User → RC → Checks Pods → If missing → Creates new Pod

---

## Key Parts of YAML

apiVersion: v1  
kind: ReplicationController  
metadata:  
  name: my-app  

---

spec:  
  replicas: 3  
  selector:  
    app: my-app  

---

template:  
  metadata:  
    labels:  
      app: my-app  

---

containers:  
  - name: nginx  
    image: nginx  

---

## Line by Line Explanation

### replicas: 3
👉 Means: always keep 3 Pods running

---

### selector
👉 RC uses this to find matching Pods

---

### labels
👉 Tags applied to Pods so RC can identify them

---

### template
👉 Blueprint for creating new Pods

---

## Important Concept

RC does NOT manage existing Pods  
It only ensures desired number exists

---

## Basic Commands

kubectl apply -f rc.yml  
kubectl get replicationcontroller  
kubectl delete replicationcontroller my-app  
kubectl get pods  

---

## Real DevOps Use Case

- Small legacy applications
- Simple scaling before Deployments existed

---

## Modern Reality (VERY IMPORTANT)

👉 RC is OLD method  
👉 Now replaced by ReplicaSet + Deployment

But:
- still asked in CKA exam
- still important for understanding basics

---

## Interview Questions

Q: What happens if a Pod dies in RC?  
👉 RC creates a new Pod automatically

Q: Difference between Pod and RC?  
👉 Pod = single instance  
👉 RC = ensures multiple instances always running

---

## Summary

- RC maintains desired number of Pods
- Provides self-healing
- Uses labels + selectors
- Mostly replaced by ReplicaSet in modern Kubernetes