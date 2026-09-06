# 🚀 Linux File Viewing & Editing Guide

> **Core Concept:** Essential utilities used to inspect file contents, view live logs, edit text files, and compare differences directly from the terminal.

---

## ⚡ Quick Reference: Viewing & Editing Commands

### 1. Viewing Content (Full & Page-by-Page)
* **`cat file`**
  * **Purpose:** Displays the full content of a file directly in the terminal output.
* **`less file`, `more file`**
  * **Purpose:** Views file contents page-by-page. `less` is more advanced as it lets you scroll up and down smoothly.

### 2. Viewing Specific Portions (Head & Tail)
* **`head file`, `head -n 20 file`**
  * **Purpose:** Shows the first lines of a file (default is 10 lines). Use `-n 20` to specify the exact number of lines.
* **`tail file`, `tail -f file`**
  * **Purpose:** Shows the last lines of a file. The `-f` flag "follows" live updates, which is extremely useful for monitoring log files in real-time.

### 3. Terminal Text Editors
* **`nano file`**
  * **Purpose:** A simple, beginner-friendly terminal text editor with shortcut keys displayed at the bottom.
* **`vim file`**
  * **Purpose:** An advanced, powerful terminal text editor used by engineers for rapid configuration and coding.

### 4. Analysis & Comparison
* **`wc file`, `wc -l file`**
  * **Purpose:** Counts words, lines, and characters in a file. The `-l` flag specifically counts the total number of lines.
* **`diff file1 file2`**
  * **Purpose:** Compares two files line by line and highlights the exact differences between them.
