# 🚀 Linux Disk Partitioning & LVM Management Guide

> **Core Concept:** Advanced storage administration utilities used to manage disk partitions, configure Logical Volume Managers (LVM), format filesystems, and mount storage devices.

---

## ⚡ Quick Reference: Disk Partitioning & LVM Commands

### 1. Partition Inspection & Identification
* **`fdisk -l`**
  * **Purpose:** Lists all available disk partitions and their partition tables.
* **`lsblk`**
  * **Purpose:** Displays all block devices and partitions in a clean, visual tree layout.
* **`blkid`**
  * **Purpose:** Displays unique identifiers (UUIDs) and filesystem types of available partitions.

### 2. LVM Core Components (PV, VG, LV)
* **`pvcreate`, `pvs`, `pvdisplay`**
  * **Purpose:** Creates, lists, and displays detailed information about Physical Volumes (PV) from raw physical disks.
* **`vgcreate`, `vgs`, `vgdisplay`**
  * **Purpose:** Creates and manages Volume Groups (VG) by combining physical volumes into a single storage pool.
* **`lvcreate`, `lvs`, `lvdisplay`**
  * **Purpose:** Creates and manages Logical Volumes (LV), allowing dynamic allocation of storage slices from volume groups.

### 3. LVM Expansion & Filesystem Resizing
* **`lvextend`**
  * **Purpose:** Dynamically extends or increases the size of a Logical Volume.
* **`resize2fs`**
  * **Purpose:** Resizes the underlying filesystem to match the new expanded size of a Logical Volume (used after `lvextend`).

### 4. Formatting & Mounting
* **`mkfs.ext4 /dev/sdX`**
  * **Purpose:** Formats a disk partition or logical volume with the ext4 filesystem.
* **`mount`, `umount`**
  * **Purpose:** Attaches (mounts) or detaches (unmounts) a storage partition to a directory path in the Linux file system hierarchy.
