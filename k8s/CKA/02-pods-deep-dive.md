# 02 - Pods Deep Dive (Lifecycle, States, Real-World Debugging)

---

## 🔰 Pod kya hota hai?

Pod = Kubernetes ka **smallest deployable unit**

Simple:
👉 Pod ke andar 1 ya multiple containers hote hain

---

## 🧠 Real World Soch

Tum directly container deploy nahi karte
👉 tum Pod deploy karte ho

Pod = wrapper around container

---

## 📦 Basic Pod Structure (YAML)

apiVersion: v1
kind: Pod
metadata:
name: nginx-pod
labels:
app: web
spec:
containers:

* name: nginx
  image: nginx
  ports:

  * containerPort: 80

---

## 🧠 Breakdown

* apiVersion → Kubernetes API version
* kind → resource type (Pod)
* metadata → name + labels
* spec → actual configuration

---

## ⚙️ Pod Lifecycle (VERY IMPORTANT)

Pod ka lifecycle hota hai:

1. Pending
2. ContainerCreating
3. Running
4. Succeeded / Failed

---

## 🔄 States Samajho

### 1. Pending

👉 Pod create ho gaya, but start nahi hua

Reasons:

* node available nahi
* resources kam

---

### 2. ContainerCreating

👉 Image pull ho rahi hai

---

### 3. Running

👉 Sab kuch sahi chal raha

---

### 4. Failed

👉 Container crash ho gaya

---

### 5. CrashLoopBackOff (REAL PAIN)

👉 Container bar bar crash ho raha

---

## 🧠 Real World Example

Scenario:
App start hoti hai → error → crash → restart → crash

👉 Kubernetes retry karta rehta hai

---

## 🔍 Debugging Flow (Production Reality)

Step 1:
kubectl get pods

Check:

* STATUS

---

Step 2:
kubectl describe pod <name>

Check:

* Events
* errors

---

Step 3:
kubectl logs <pod-name>

👉 Actual error yahan milta hai

---

## 📊 Logs Deep Dive

kubectl logs webapp

👉 single container

---

Multi-container pod:

kubectl logs webapp -c container-name

---

Follow logs (live):

kubectl logs -f webapp

---

## 🧠 Real DevOps Scenario

Issue:
👉 "App down"

Tum:
kubectl logs → error: DB connection failed

Conclusion:
👉 issue pod mein nahi → backend dependency mein hai

---

## 🔁 Restart Behavior

Kubernetes automatically restart karta hai container

Restart policies:

* Always (default)
* OnFailure
* Never

---

## 📦 Multi-Container Pod (Intro)

Example:

containers:

* nginx
* redis

---

## 🧠 Use Case

* App + logging agent
* App + sidecar

---

## ⚠️ Important Concept

Pod ke andar:
👉 sab containers same:

* IP address
* network
* storage (if shared)

---

## 🌐 Networking Insight

Pod = 1 IP

Even if:
👉 3 containers inside

---

## 📉 Common Errors (Real Life)

### ❌ ImagePullBackOff

Reason:

* wrong image name
* private repo

---

### ❌ CrashLoopBackOff

Reason:

* app error
* config issue

---

### ❌ Pending

Reason:

* no resources

---

## 🧠 Real Debug Case

kubectl describe pod nginx

Event:
"FailedScheduling"

👉 Node pe jagah nahi

---

## 🔥 Production Rule

👉 NEVER deploy single Pod manually

Use:

* Deployment
* ReplicaSet

---

## 🎯 Why?

Pod:
❌ not self-healing alone
❌ not scalable

---

## 🚀 Mini Scenario

Tumne pod banaya:

kubectl run webapp --image=nginx

User load aya:
👉 pod crash

👉 no backup

👉 site down

---

## 🧠 Correct Way

Deployment use karo
👉 multiple pods
👉 auto-healing

---

## 🧩 Pod vs Container

Container = app
Pod = wrapper + environment

---

## 🧠 Final Mental Model

Think like this:

Pod = mini server
Container = app inside server

---

## 🎯 Summary

* Pod = smallest unit
* lifecycle samajhna MUST
* logs = main debugging tool
* CrashLoopBackOff = common issue
* production = NEVER use raw pods

---

## 🚀 Senior Advice

👉 Tumhara kaam hai:

"Pod ko chalana nahi
Pod ko samajhna"

---


END OF FILE 💀
