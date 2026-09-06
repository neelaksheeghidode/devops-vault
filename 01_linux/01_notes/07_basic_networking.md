# 🚀 Linux Basic Networking Guide

> **Core Concept:** Essential networking utilities used to test host connectivity, inspect network interfaces, transfer data, analyze open ports, and perform DNS lookups in Linux.

---

## ⚡ Quick Reference: Basic Networking Commands

### 1. Connectivity & Reachability
* **`ping host`**
  * **Purpose:** Checks network reachability and latency to a remote host or IP address using ICMP echo requests.

### 2. Network Interfaces & IP Configuration
* **`ifconfig`, `ip a`**
  * **Purpose:** Displays network interfaces, active IP addresses, MAC addresses, and packet statistics. (`ip a` is the modern standard replacement for `ifconfig`).
* **`hostname`**
  * **Purpose:** Displays or sets the system's host name.

### 3. Web Downloads & Data Transfer
* **`curl -IL url`**
  * **Purpose:** Transfers data to or from a server. The `-I` flag fetches HTTP headers only, and `-l` follows redirects.
* **`wget url`**
  * **Purpose:** A non-interactive command-line tool used to download files from the web.

### 4. Ports & Listening Services
* **`netstat -tulnp`, `ss -tulnp`**
  * **Purpose:** Displays open ports, active connections, and listening network services. (`ss` is faster and more modern than `netstat`).

### 5. Routing & DNS Diagnostics
* **`traceroute`**
  * **Purpose:** Traces the network packet path and hops taken to reach a destination host.
* **`nslookup`, `dig`**
  * **Purpose:** Performs DNS lookups to query domain name servers and resolve domain names to IP addresses (`dig` provides more detailed technical breakdowns).
