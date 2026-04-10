.

🧠 STEP 1 — CORE COMPONENTS (CLEAR CONCEPT)
📦 1. Pod

👉 Smallest unit
👉 Container ko wrap karta hai

Real life:

Docker container = worker
Pod = worker ka room

🚀 2. Deployment

👉 Pods ko manage karta hai:

kitne pods hone chahiye
update kaise hoga
rollback kaise hoga
🌐 3. Service

👉 Network layer
👉 Pods ko expose karta hai

Problem solve:
❌ Pod IP change hota rehta
✅ Service stable IP deta

⚙️ 4. ConfigMap

👉 Non-sensitive config store karta hai
Example:

APP_COLOR
ENV
🔐 5. Secret

👉 Sensitive data store karta hai
Example:

DB password
API keys
🔗 6. envFrom

👉 ConfigMap/Secret ko container ke environment variables bana deta hai


📄 STEP 2 — COMPLETE REAL YAML (PRODUCTION STYLE)
🚀 FULL APP (Deployment + Service + Config + Secret)
# 🔹 CONFIGMAP
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  APP_COLOR: blue
  APP_ENV: production

---

# 🔹 SECRET
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
type: Opaque
data:
  DB_PASSWORD: cGFzc3dvcmQ=   # base64 encoded

---

# 🔹 DEPLOYMENT
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webapp-deployment

spec:
  replicas: 3   # desired pods

  selector:
    matchLabels:
      app: webapp   # MUST match template labels

  template:
    metadata:
      labels:
        app: webapp

    spec:
      containers:
      - name: webapp-container
        image: nginx:latest   # container image

        ports:
        - containerPort: 80

        envFrom:
        - configMapRef:
            name: app-config
        - secretRef:
            name: app-secret

---

# 🔹 SERVICE
apiVersion: v1
kind: Service
metadata:
  name: webapp-service

spec:
  selector:
    app: webapp

  ports:
  - protocol: TCP
    port: 80
    targetPort: 80

  type: NodePort
🔍 STEP 3 — LINE BY LINE BREAKDOWN
🔹 COMMON FIELDS (IMPORTANT)

These are in almost EVERY YAML:

🔸 apiVersion

👉 Kubernetes API version
Example:

v1 → basic objects
apps/v1 → deployments
🔸 kind

👉 Resource type
Example:

Pod
Deployment
Service
🔸 metadata

👉 Information about object

name
labels
🔸 spec

👉 Actual configuration (MOST IMPORTANT)

🔍 CONFIGMAP BREAKDOWN
data:
  APP_COLOR: blue

👉 Ye environment variable banega inside container

🔍 SECRET BREAKDOWN
data:
  DB_PASSWORD: cGFzc3dvcmQ=

👉 base64 encoded
👉 container ke andar use hoga

🔍 DEPLOYMENT (DEEP)
🔸 replicas
replicas: 3

👉 Always 3 pods running

🔸 selector
matchLabels:
  app: webapp

👉 Deployment kis pods ko manage karega

⚠️ MUST MATCH template labels

🔸 template

👉 Pod ka blueprint

🔸 containers
image: nginx:latest

👉 Docker image

🔸 ports
containerPort: 80

👉 Container kis port pe run karega

🔸 envFrom
envFrom:
- configMapRef:
    name: app-config

👉 ConfigMap ke variables inject

🔍 SERVICE (VERY IMPORTANT)
🔸 selector
app: webapp

👉 Service kis pod ko target karega

🔸 port vs targetPort
port: 80
targetPort: 80

👉 port = service port
👉 targetPort = container port

🔸 type
NodePort

👉 External access enable karta

🔗 STEP 4 — CONNECTION FLOW (MOST IMPORTANT)
🔥 FULL FLOW
Deployment create hota
Deployment → ReplicaSet create karta
ReplicaSet → Pods create karta
Pods → container run karte (nginx)
ConfigMap + Secret → env variables inject
Service → pods ko expose karta
🌍 REAL USER FLOW

User:
👉 hits Node IP:Port

Service:
👉 receives request

Selector:
👉 finds pods (label match)

Pod:
👉 nginx response deta

🧠 STEP 5 — UNDERSTANDING LEVEL (CHECK YOURSELF)

Agar tum ye explain kar sako:

👉 “User request kaise container tak jaati hai?”

Then:
✔ You understand Kubernetes

⚠️ COMMON MISTAKES (IMPORTANT)

❌ selector != labels
👉 pods connect nahi honge

❌ wrong image
👉 ImagePullBackOff

❌ secret not base64
👉 error

❌ service not matching labels
👉 no traffic