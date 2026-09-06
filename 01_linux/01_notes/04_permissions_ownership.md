# 🚀 Linux Permissions & Ownership Guide

> **Core Concept:** Essential security utilities used to control access rights, define file execution permissions, and manage user/group ownership in Linux.

---

## ⚡ Quick Reference: Permissions & Ownership Commands

### 1. Modifying Permissions (`chmod`)
* **`chmod 755 file`**
  * **Purpose:** Changes the access permissions of a file or directory using numeric (octal) mode.
* **`chmod +x file`**
  * **Purpose:** Adds execute permissions to a file, making it runnable as a script or program.

### 2. Managing Ownership (`chown` & `chgrp`)
* **`chown user:group file`**
  * **Purpose:** Changes both the owner user and the group ownership of a file or directory.
* **`chgrp group file`**
  * **Purpose:** Changes only the group ownership of a file or directory.

### 3. Inspection & Default Masks (`ls -l` & `umask`)
* **`ls -l`**
  * **Purpose:** Lists files in long format, displaying permissions in `rwx` (read, write, execute) format.
* **`umask`**
  * **Purpose:** Displays or sets the default permission mask used when new files and directories are created.
