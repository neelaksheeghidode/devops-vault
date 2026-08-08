# 🚀 Day 01: Linux Architecture & Cloud Networking Fundamentals

## 📌 Course Overview & Context
This repository contains comprehensive notes and practical breakdowns from **Day 01** of the Linux for DevOps series. Day 01 focuses on establishing strong foundational knowledge in web networking, server infrastructure, application types, Linux OS internal architecture, and system diagnostic concepts essential for a DevOps Engineer.

---

## 🌐 Section 1: Internet Architecture & Networking Basics

### 1. How Does the Internet Work?
- **The Misconception:** The internet does not primarily run on slow satellite connections.
- **The Reality:** Internet connectivity relies on massive networks of **undersea/submarine optical fiber cables** spanning oceans and continents.
- **Data Centers:** Centralized facilities packed with high-performance physical servers that store data, host application services, and process global web requests via network backbones.
- **ISPs (Internet Service Providers):** Telecom companies (e.g., AT&T, Reliance Jio, Airtel) own/lease fiber infrastructure and provide access to end-users via cellular towers and broadband connections.

### 2. Client-Server Architecture
- **Client:** Any end-user device or software (Web Browser, Mobile App, Postman) requesting data over a network.
- **Server:** A specialized computer or process running 24/7 designed to receive incoming requests, execute logic, and send back responses.

---

## 💻 Section 2: Application Types & Web Infrastructure

### 1. Classification of Applications
- **Standalone Applications:** 
  - Software that runs locally on a single machine without requiring an active network connection or database interaction (e.g., Calculator, Paint, basic text editors).
- **Web Applications:** 
  - Network-enabled software accessed via web browsers that processes dynamic data over HTTP/HTTPS protocols (e.g., YouTube, Amazon, Gmail).

### 2. Web Servers vs. Application Servers
| Feature | Web Server | Application Server |
| :--- | :--- | :--- |
| **Primary Purpose** | Serves static content (HTML, CSS, JS, Images) & routes incoming HTTP requests. | Executes complex business logic, backend code, and database transactions. |
| **Data Type** | Static assets. | Dynamic responses. |
| **Popular Tools** | Nginx, Apache HTTP Server. | Node.js, Apache Tomcat, Gunicorn, Django. |

### 3. Application Support & Maintenance
- **Application Support:** Monitoring, troubleshooting server logs, handling runtime errors, and ensuring high system uptime.
- **Application Maintenance:** Regularly updating dependencies, patching security vulnerabilities, scaling server resources, and deploying new features without causing service downtime.

---

## 🐧 Section 3: Linux Operating System Architecture

### 1. What is Linux?
Linux is an open-source, Unix-like operating system kernel created by Linus Torvalds. It is licensed under the **GPL (General Public License)**, making it free to modify, distribute, and use across enterprise cloud infrastructure.

### 2. Linux vs. Windows
- **GUI vs. CLI:** Windows is optimized for graphical desktop use; Linux is lightweight, highly customizable, and optimized for command-line interface (CLI) execution.
- **Resource Overhead:** Linux servers run without heavy desktop environments, leaving maximum CPU and RAM available for applications.
- **Licensing & Costs:** Windows Server requires paid licensing; most Linux distributions (Ubuntu, Debian, Rocky Linux) are completely open-source and free.

### 3. Core Operating System Architecture Layers

[ User / DevOps Engineer ]
│
▼
[ Shell (Terminal Interface) ]
│
▼
[ Kernel (Core System Controller) ]
│
▼
[ Hardware (CPU, RAM, Storage, Network) ]


### 4. Key OS Components Explained
- **Bootloader:** The initial program (e.g., GRUB) that loads the Linux kernel into main memory (RAM) during computer startup.
- **Kernel:** The primary core of the OS that directly communicates with hardware, managing system memory, process scheduling, and device drivers.
- **Shell:** A command interpreter (e.g., `bash`, `zsh`) that acts as a bridge/gateway between the user and the kernel by converting user commands into kernel system calls.
- **Desktop Environment (DE):** Optional graphical layer (e.g., GNOME, KDE) for desktop users. Servers run *headless* (without DE) to save system resources.

### 5. Remote Server Access Tools
To manage remote Linux servers located in data centers or cloud platforms (AWS, Azure):
- **SSH (Secure Shell - Port 22):** Encrypted network protocol for terminal access.
- **Tools:** PuTTY, MobaXterm, OpenSSH (built-in terminal), VS Code Remote-SSH.

---

## ⚙️ Section 4: System Exploration & Process Lifecycle

### 1. Linux Process States
Every program running in Linux is assigned a unique PID (Process ID) and cycles through the following states:
1. **Running (`R`):** The process is currently executing on the CPU or is in the runnable queue.
2. **Interruptible Sleep (`S`):** The process is waiting for an event or resource (e.g., user input or network response).
3. **Uninterruptible Sleep (`D`):** The process is waiting directly for hardware/disk I/O operations and cannot be interrupted.
4. **Zombie (`Z`):** A terminated process whose parent process has not yet read its exit status, leaving an entry in the process table.

### 2. Essential System Diagnostic Commands
```bash
# Display absolute path of current working directory
pwd

# List directory contents with hidden files and permissions
ls -la

# View currently logged-in user identity
whoami

# Check memory (RAM) utilization in MBs/GBs
free -m

# Inspect disk storage usage across mounted filesystems
df -h

# Monitor real-time process execution and CPU/RAM usage
top
