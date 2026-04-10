# 05 - Services (Networking Basics - Pods ko expose karna)

---

## 🔰 Problem (Real World)

Tumne Deployment banaya:

👉 pods chal rahe ✔

👉 scaling bhi ho rahi ✔

Lekin…

❌ user app tak kaise pohanchay?

❌ pod ka IP change hota rehta hai

❌ direct pod access unreliable hai

---

## 🚀 Solution

👉 Service

---

## 🧠 Simple Definition

Service = stable network endpoint

👉 jo pods ko expose karta hai

---

## 📊 Core Idea

Pods = dynamic (IP change hota rehta)

Service = static (stable IP / DNS)

---

## 🧠 Real World Soch

Service = load balancer + stable entry point

---

## 📦 Types of Services

1. ClusterIP (default)
2. NodePort
3. LoadBalancer

---

## 🔵 1. ClusterIP (Default)

👉 sirf cluster ke andar accessible

---

### Use Case

* backend services
* internal communication

---

### Example YAML

```
apiVersion: v1
kind: Service
metadata:
  name: webapp-service
spec:
  type: ClusterIP
  selector:
    app: web
  ports:
  - port: 80
    targetPort: 80
```

---

## 🧠 Breakdown

* selector → kaunse pods connect karne hain
* port → service port
* targetPort → container port

---

## 🟢 2. NodePort

👉 external access allow karta hai

---

### Concept

Node ka port open hota hai

👉 user us port se access karta hai

---

### Example YAML

```
apiVersion: v1
kind: Service
metadata:
  name: webapp-nodeport
spec:
  type: NodePort
  selector:
    app: web
  ports:
  - port: 80
    targetPort: 80
    nodePort: 30007
```
---

## 🧠 Access

http://NodeIP<NodeIP>:30007

---

## ⚠️ Limitations

* port range limited (30000–32767)
* production mein rarely direct use

---

## 🟣 3. LoadBalancer

👉 cloud environment ke liye

---

### Concept

Cloud provider (AWS, GCP)

👉 external load balancer create karta hai

---

### Example YAML

```
apiVersion: v1
kind: Service
metadata:
  name: webapp-lb
spec:
  type: LoadBalancer
  selector:
    app: web
  ports:
  - port: 80
    targetPort: 80
```

---

## 🧠 Real World

👉 AWS ELB create hota hai automatically

---

## ⚙️ Commands

Create:

kubectl apply -f service.yaml

---

Check:

kubectl get services

---

Details:

kubectl describe service webapp-service

---

## 🔄 How Service Works

Step by step:

1. selector match karta hai pods
2. endpoints create hotay hain
3. traffic distribute hota hai

---

## 🧠 Load Balancing

Service automatically:

👉 traffic multiple pods mein distribute karta hai

---

## 📊 Real Scenario

3 pods running:

User request:

👉 Service → Pod1

Next request:

👉 Service → Pod2

---

## 🧠 Important Concept

Service = label selector based routing

---

## 🔍 Debugging

kubectl get svc

kubectl describe svc svc-name

---

Check endpoints:

kubectl get endpoints

---

## ❌ Common Errors

1. selector mismatch

   👉 no endpoints

2. wrong port mapping

   👉 app unreachable

---

## 🧠 Debug Case

kubectl get svc

Service exists ✔

But app not working ❌

👉 check:

kubectl get endpoints

Empty?

👉 selector issue

---

## 🌐 Port Forward (Testing Hack)

kubectl port-forward deployment/webapp 8080:80

---

## 🧠 Use Case

Local testing

without exposing service

---

## 🔥 Production Thinking

* ClusterIP → internal
* NodePort → testing
* LoadBalancer → production

---

## 🧩 DNS Concept

Service ka name hi DNS ban jata hai

Example:

webapp-service.default.svc.cluster.local

---

## 🧠 Real DevOps Use

Frontend → backend call:

http://webapp-service

---

## 🚀 Mini Scenario

Tumhari app:

Frontend (React)

Backend (API)

Backend expose via Service

Frontend call karega:

👉 Service name use karke

---

## 🎯 Summary

* Service = stable endpoint
* Pods = dynamic
* Load balancing built-in
* Types:

  * ClusterIP
  * NodePort
  * LoadBalancer

---

## 🚀 Senior Advice

👉 "No Service = No App"

👉 Agar Service nahi samjhi

→ Networking fail

→ App useless

---

END OF FILE 💀
