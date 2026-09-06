# 🚀 Linux System Information & Disk Usage Guide

> **Core Concept:** Essential diagnostics utilities used to monitor core system details, check available disk storage, analyze folder sizes, and inspect system memory usage.

---

## ⚡ Quick Reference: System Info & Disk Usage Commands

### 1. System & Identity Information
* **`uname -a`**
  * **Purpose:** Displays comprehensive system and kernel version information.
* **`whoami`**
  * **Purpose:** Shows the username of the currently logged-in user.
* **`date`**
  * **Purpose:** Shows the current system date, time, and timezone.
* **`uptime`**
  * **Purpose:** Shows how long the system has been running, active users, and load averages.

### 2. Disk Storage & Folder Sizes
* **`df -h`**
  * **Purpose:** Displays available and used disk space across all mounted filesystems in a human-readable format (`-h`).
* **`du -sh folder`**
  * **Purpose:** Shows the total disk size consumed by a specific folder or file in a summary format (`-s` for summary, `-h` for human-readable).

### 3. Memory Monitoring
* **`free -h`**
  * **Purpose:** Displays current system RAM and Swap memory usage in a clear, human-readable format.
