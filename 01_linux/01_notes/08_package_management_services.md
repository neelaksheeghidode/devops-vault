# 🚀 Linux Package Management & Services Guide

> **Core Concept:** Essential administrative utilities used to manage software installations, updates, and system daemon services across different Linux distributions.

---

## ⚡ Quick Reference: Package Management & Services Commands

### 1. System Updates & Maintenance (Debian / Ubuntu)
* **`apt update && apt upgrade`**
  * **Purpose:** Refreshes the local package index database and upgrades all installed software packages to their latest available versions.

### 2. Installing Packages
* **`apt install package`**
  * **Purpose:** Downloads and installs a software package on Debian, Ubuntu, or derivative distributions.
* **`yum install`, `dnf install`**
  * **Purpose:** Installs software packages on Red Hat Enterprise Linux (RHEL), CentOS, or Fedora distributions (`dnf` is the modern replacement for `yum`).

### 3. Removing Packages & Listing Software
* **`apt remove package`**
  * **Purpose:** Uninstalls and removes a package from the system while retaining its configuration files.
* **`dpkg -l`, `rpm -qa`**
  * **Purpose:** Lists all software packages currently installed on the system (`dpkg` for Debian/Ubuntu, `rpm` for RHEL/CentOS).

### 4. Managing System Services (`systemctl`)
* **`systemctl start service`**
  * **Purpose:** Immediately starts a system service daemon.
* **`systemctl stop service`**
  * **Purpose:** Stops a running system service.
* **`systemctl restart service`**
  * **Purpose:** Restarts a service (stops and starts it back up sequentially).
* **`systemctl status service`**
  * **Purpose:** Checks the current execution status, logs, and health of a service.
* **`systemctl enable service`**
  * **Purpose:** Configures a system service to automatically start on boot when the machine powers up.
