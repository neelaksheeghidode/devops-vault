# 🚀 Linux Basic System & Information Commands Guide

> **Core Concept:** Essential utilities and commands used to query system state, check environment details, and manage terminal sessions in Linux.

---

## ⚡ Quick Reference: System & Info Commands

### 1. Working Directory & Environment
* **`pwd` (Print Working Directory)**
  * **Purpose:** Displays the absolute path of the directory you are currently working in.
  * **Syntax:** `pwd`

### 2. System & Kernel Information
* **`uname -a`, `uname -r`**
  * **Purpose:** Displays core system and Linux kernel version details. `-a` provides all system info, while `-r` specifically outputs the kernel release version.
  * **Syntax:** `uname -a` | `uname -r`

### 3. Date, Time & Calendar
* **`date`**
  * **Purpose:** Shows the current system date, time, timezone, and year.
  * **Syntax:** `date`
* **`cal`**
  * **Purpose:** Displays a formatted monthly calendar in the terminal.
  * **Syntax:** `cal`

### 4. Terminal Utility & Navigation
* **`clear`**
  * **Purpose:** Clears the terminal screen for a clean working space.
  * **Syntax:** `clear`
* **`history`**
  * **Purpose:** Displays a numbered list of previously executed commands in your current session.
  * **Syntax:** `history`
* **`echo "text"`**
  * **Purpose:** Prints the specified text or variable value directly to the terminal output.
  * **Syntax:** `echo "Hello DevOps"`

### 5. Help & Session Control
* **`man command` (Manual)**
  * **Purpose:** Opens the official manual/help documentation page for any Linux command.
  * **Syntax:** `man ls` or `man pwd`
* **`exit`**
  * **Purpose:** Safely terminates and closes the current terminal session or shell.
  * **Syntax:** `exit`

---

## 💡 Best Practices & Pro Tips

* **Documentation Lookup:** Instead of guessing command flags, always rely on `man <command>` to read the native manual pages.
* **Command Recall:** Use the `history` command combined with `grep` (e.g., `history | grep docker`) to quickly find commands you ran previously.
