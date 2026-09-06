# 🚀 Linux Advanced Text Processing & Search Guide

> **Core Concept:** Powerful command-line text manipulation and filtering utilities used to search patterns, extract columns, transform text streams, and automate data processing.

---

## ⚡ Quick Reference: AWK, GREP, SED & Stream Tools

### 1. Pattern Matching & Searching (`grep`)
* **`grep "pattern" file`**
  * **Purpose:** Searches for specific text patterns or regular expressions inside files.
* **`grep -i`, `grep -r`**
  * **Purpose:** Modifies search behavior. `-i` makes the search case-insensitive, and `-r` performs a recursive search across all files within directories.

### 2. Advanced Text Extraction & Editing (`awk` & `sed`)
* **`awk '{print $1}'`**
  * **Purpose:** An advanced text processing and pattern scanning language used heavily for column extraction and data parsing.
* **`sed 's/old/new/g'`**
  * **Purpose:** A stream editor (SED) used to find and replace text patterns globally (`/g`) across streams or files.
* **`cut -d',' -f1 file`**
  * **Purpose:** Extracts specific columns or fields from files based on a designated delimiter (e.g., a comma).

### 3. Sorting, Filtering & Character Manipulation
* **`sort`**
  * **Purpose:** Sorts lines of text files alphabetically or numerically.
* **`uniq`**
  * **Purpose:** Filters out or detects duplicate consecutive lines in a text stream.
* **`tr`**
  * **Purpose:** Translates, squeezes, and/or deletes specific characters from standard input.

### 4. Stream Splitting & Argument Building (`tee` & `xargs`)
* **`tee file`**
  * **Purpose:** Reads standard input and writes it to both the screen (stdout) and one or more files simultaneously.
* **`xargs`**
  * **Purpose:** Builds and executes command-line arguments from standard input, allowing output from one command to be passed as arguments to another.
