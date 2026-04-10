# 01 - kubectl Core Commands (Beginner → Advanced)

---

## 🔰 kubectl kya hai?

kubectl = Kubernetes ka **remote control**

Simple words:
👉 Ye tool use hota hai cluster se baat karne ke liye

Jaise:

* Docker → docker CLI
* Git → git CLI
* Kubernetes → kubectl

---

## 🧠 Real World Soch (DevOps POV)

Production mein:

* Tum directly servers pe login nahi karte
* Tum cluster ko control karte ho using kubectl

👉 Tumhara role:
"System ko manually nahi, commands se control karna"

---

## ⚙️ Basic Syntax

kubectl <action> <resource> <options>

Example:
kubectl get pods

Breakdown:

* get → action
* pods → resource

---

## 📦 1. Pods ko dekhna

kubectl get pods

👉 Ye tumhe batata:

* kaunse pods chal rahe hain
* unka status (Running / Pending / CrashLoop)

---

### 🔍 More Detail View

kubectl get pods -o wide

👉 Extra info:

* Node name (kis machine pe chal raha)
* IP address

---

### 🧠 Real World

Agar app down hai:
👉 sabse pehle: kubectl get pods

---

## 🚀 2. Pod create karna (Imperative way)

kubectl run nginx --image=nginx

👉 Ye kya karta?

* ek pod create karta hai
* nginx container run karta hai

---

### ⚠️ Reality Check

Production mein:
👉 direct "kubectl run" kam use hota hai

👉 mostly YAML files use hoti hain (Declarative)

---

## 🔍 3. Pod details dekhna

kubectl describe pod <pod-name>

Example:
kubectl describe pod nginx

---

### 📊 Ye kya deta?

* Events (errors)
* Image name
* Ports
* Restart count

---

### 🧠 Real World Debugging

Agar pod crash ho:
👉 use karo:

kubectl describe pod <name>

---

## 🗑️ 4. Pod delete karna

kubectl delete pod <pod-name>

Example:
kubectl delete pod webapp

---

### ⚠️ Important

Agar pod deployment ka part hai:
👉 wo dobara create ho jayega (self-healing)

---

## 🔄 5. Image update karna (Live change)

kubectl set image pod/<pod-name> <container>=<image>

Example:
kubectl set image pod/redis redis=redis:latest

---

### 🧠 Reality

👉 Direct pod update rarely use hota hai
👉 Mostly deployments use hote hain

---

## 📊 6. Multiple resources ek sath dekhna

kubectl get all

👉 Show karega:

* pods
* services
* deployments
* etc.

---

## 🎯 7. Specific resource type

kubectl get:

* pods
* deployments
* services
* replicasets

Examples:
kubectl get deployments
kubectl get services

---

## 🧠 Pattern samjho

Har cheez Kubernetes mein ek "resource" hai

---

## 🔎 8. Filtering (Selectors)

kubectl get pods --selector app=App1

👉 Ye sirf wo pods dikhata jinke paas label match karta hai

---

### 🧠 Real World

Production mein:
👉 hundreds pods hote hain

👉 filtering is MUST

---

## 📚 9. Documentation CLI se

kubectl explain pods

Deep dive:
kubectl explain pods.spec

---

### 🧠 Use Case

👉 Interview + YAML likhne ke time

---

## 🔁 10. Force replace

kubectl replace --force -f nginx.yaml

👉 Delete + recreate

---

### ⚠️ Risk

👉 downtime ho sakta hai

---

## 🧠 DevOps Insight

Ye tab use karo jab:

* config badly broken ho
* fast reset chahiye

---

## 📦 11. Namespaces ke sath kaam

kubectl get pods --namespace=research

kubectl get pods --all-namespaces

---

### 🧠 Real World

Namespaces = environments

Example:

* dev
* staging
* prod

---

## 🧪 12. Resource discovery

kubectl api-resources

👉 Ye batata:

* cluster mein kaun kaun se resources available hain

---

## 🧠 Advanced Thinking

👉 Kubernetes = API driven system
👉 kubectl = API client

---

## 🔥 Golden Debugging Flow (Remember This)

1. kubectl get pods
2. kubectl describe pod <name>
3. kubectl logs <name>

👉 90% issues yahin solve hote hain

---

## 🚀 Mini Real World Scenario

Tum DevOps engineer ho.

Issue:
👉 "Website down hai"

Steps:

1. kubectl get pods
   → check status

2. kubectl describe pod
   → check errors

3. kubectl logs
   → actual issue find

---

## 🧠 Final Summary

* kubectl = control tool
* get = check state
* describe = debug
* delete = remove
* run = quick test
* set image = update

---

## 🎯 Senior Advice

👉 Commands yaad mat karo
👉 pattern samjho:

Observe → Debug → Fix

---

END OF FILE 🚀
