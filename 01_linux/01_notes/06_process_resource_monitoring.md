# 🚀 Linux Process & Resource Monitoring Guide

> **Core Concept:** Essential utilities used to monitor running processes, track system resource usage (RAM, CPU, Swap), manage background jobs, and troubleshoot system performance.

---

## ⚡ Quick Reference: Process & Resource Monitoring

### 1. Process Listing & Inspection
* **`ps`, `ps aux`, `ps -ef`**
  * **Purpose:** Lists currently running processes. `ps aux` and `ps -ef` provide detailed, system-wide snapshots of all active processes.

### 2. Interactive Process Viewers
* **`top`, `htop`**
  * **Purpose:** Provides a live, interactive, real-time view of running processes, CPU load, and memory usage. (`htop` offers a much more visual and user-friendly color interface).

### 3. Terminating Processes
* **`kill PID`**
  * **Purpose:** Gracefully stops a running process using its unique Process ID (PID).
* **`kill -9 PID`**
  * **Purpose:** Force kills a process immediately using the SIGKILL signal (use when a process is frozen).
* **`killall name`**
  * **Purpose:** Terminates all running processes matching a specific program or command name.

### 4. Job Control & Background Execution
* **`jobs`**
  * **Purpose:** Lists background jobs currently running in your active shell session.
* **`bg`, `fg`**
  * **Purpose:** Manages job states. `bg` sends a suspended job to the background, and `fg` brings a background job back to the foreground.
* **`nohup command &`**
  * **Purpose:** Executes a command that keeps running in the background even after you close the terminal session or disconnect from SSH.

### 5. System Resource & Performance Monitoring
* **`free -h`**
  * **Purpose:** Displays system memory (RAM) and Swap usage in a clear, human-readable format.
* **`uptime`**
  * **Purpose:** Shows how long the system has been running, the number of logged-in users, and system load averages.
* **`vmstat`**
  * **Purpose:** Reports virtual memory statistics, CPU activity, processes, and block I/O performance.
