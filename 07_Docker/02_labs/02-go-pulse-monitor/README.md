cat << 'EOF' > README.md
# 🚀 Go Pulse Monitor

[![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Multi--Stage-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![Status](https://img.shields.io/badge/Status-Operational-success?style=for-the-badge)]()

A high-performance, lightweight cloud-native microservice built in **Go (Golang)** that exposes enterprise-grade system telemetry via a secure `/health` endpoint. Designed using production-ready **multi-stage Docker builds** for minimal footprint and maximum security.

---

## 🛠️ Tech Stack & Architecture
* **Backend:** Go (Standard `net/http` package)
* **Containerization:** Docker (Multi-stage build pattern: Alpine compiler + Alpine runtime)
* **Output:** Structured JSON telemetry data

---

## 📡 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/health` | Returns real-time system diagnostics (OS, Architecture, CPU count, Go version, Timestamp) |

---

## 🚀 Quick Start Guide

### 1. Build the Docker Image
```bash
docker build -t go-pulse-monitor:v1 .

docker run --rm -p 8080:8080 go-pulse-monitor:v1

curl http://localhost:8080/health
```
