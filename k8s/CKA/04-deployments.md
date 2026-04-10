# 04 - Deployments (Production-Level Workloads, Rolling Updates, Zero Downtime)

---

## 🔰 Problem (Real World)

Tumne ReplicaSet banaya:

👉 pods manage ho rahe hain ✔

👉 scaling bhi ho rahi ✔

Lekin…

❌ update kaise karoge bina downtime ke?

❌ version control kaise hoga?

❌ rollback kaise hoga?

---

## 🚀 Solution

👉 Deployment

---

## 🧠 Simple Definition

Deployment = higher-level controller

👉 jo ReplicaSet ko manage karta hai

👉 aur updates safely perform karta hai

---

## 📊 Architecture
```
Deployment
↓
ReplicaSet
↓
Pods
```

---

## 📦 Basic Deployment YAML

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webapp-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      containers:
      - name: nginx
        image: nginx
```

```
spec:
  containers:
  - name: nginx
    image: nginx
```

---

## 🧠 Breakdown

* replicas → kitne pods chahiye
* template → pod ka blueprint
* selector → kin pods ko manage karna hai

---

## ⚙️ Commands

Create:

kubectl apply -f deployment.yaml

---

Check:

kubectl get deployments

kubectl get replicasets

kubectl get pods

---

Details:

kubectl describe deployment webapp-deployment

---

## 🔄 Rolling Updates (GAME CHANGER)

👉 Deployment automatically update karta hai pods

👉 ek ek karke (zero downtime)

---

## 🧠 Real Scenario

Version v1 chal rahi hai

Tum v2 deploy karte ho

Deployment:

1. new pod create karega (v2)
2. old pod remove karega (v1)
3. repeat

---

## 📈 Update Command

kubectl set image deployment/webapp-deployment nginx=nginx:1.25

---

## 🔍 Check Update Status

kubectl rollout status deployment/webapp-deployment

---

## 📜 History

kubectl rollout history deployment/webapp-deployment

---

## 🔙 Rollback (LIFESAVER)

kubectl rollout undo deployment/webapp-deployment

---

## 🧠 Real DevOps Scenario

Tumne new version deploy ki

👉 bug aa gaya

👉 site crash

Solution:

👉 rollback

👉 within seconds previous version back

---

## ⚙️ Strategy Types

Default:
👉 RollingUpdate

Options:

* RollingUpdate
* Recreate (downtime)

---

## 🧠 Advanced (Important)

RollingUpdate settings:

* maxUnavailable
* maxSurge

---

## Example:
```
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
```

---

## 🧠 Meaning

maxUnavailable = kitne pods down ho sakte

maxSurge = kitne extra pods ban sakte

---

## 📊 Real World Thinking

👉 balance between:

* speed
* availability

---

## 🔥 Zero Downtime Concept

User ko kabhi downtime feel nahi hota

👉 kyunki always kuch pods running rehte hain

---

## 🧩 Scaling

kubectl scale deployment webapp-deployment --replicas=5

---

## 🧠 Auto Healing

Agar pod crash ho:
👉 Deployment → ReplicaSet → new pod create

---

## ❌ Common Mistakes

1. selector mismatch
2. wrong image name
3. forgetting rollout status

---

## 🧠 Debugging

kubectl describe deployment <name>

kubectl get pods

kubectl logs <pod>

---

## 🚀 Real Production Scenario

Tum DevOps engineer ho:

App live hai
Users: 10,000+

Tum new feature deploy karte ho:

👉 Deployment safely update karta hai

👉 no downtime

👉 rollback possible

---

## 🧠 Final Mental Model

Deployment = Manager
ReplicaSet = Supervisor
Pod = Worker

---

## 🎯 Summary

* Deployment = production standard
* rolling updates = zero downtime
* rollback = safety net
* scaling = easy

---

## 🚀 Senior Advice

👉 Kubernetes = Deployment-driven system

👉 Agar Deployment nahi samjha
→ Kubernetes nahi samjha

---

END OF FILE 💀
