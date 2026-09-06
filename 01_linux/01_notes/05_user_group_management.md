# 🚀 Linux User & Group Management Guide

> **Core Concept:** Essential administrative utilities used to create users, manage passwords, organize groups, and control user switching and privileges in Linux.

---

## ⚡ Quick Reference: User & Group Management Commands

### 1. User Creation & Deletion
* **`useradd username`, `adduser`**
  * **Purpose:** Creates a brand new user account on the system. (`adduser` is an interactive script, while `useradd` is a low-level utility).
* **`passwd username`**
  * **Purpose:** Sets or updates the password for a user account.
* **`usermod`**
  * **Purpose:** Modifies existing user account attributes (like username, home directory, shells, or groups).
* **`userdel username`**
  * **Purpose:** Deletes an existing user account from the system.

### 2. Group Management
* **`groupadd groupname`**
  * **Purpose:** Creates a new security group.
* **`groupdel groupname`**
  * **Purpose:** Deletes an existing security group.

### 3. Identity & Session Information
* **`groups`**
  * **Purpose:** Checks the group memberships of the current or specified user.
* **`whoami`**
  * **Purpose:** Displays the username of the currently logged-in user.
* **`who`**
  * **Purpose:** Shows a list of all active users currently logged into the system.
* **`id`**
  * **Purpose:** Displays the numeric User ID (UID) and Group ID (GID) along with supplementary groups for a user.

### 4. Privilege Escalation & Switching
* **`su username`**
  * **Purpose:** Switches the current terminal session to another user account (Switch User).
* **`sudo command`**
  * **Purpose:** Executes a specific command with elevated administrator (root) privileges.
