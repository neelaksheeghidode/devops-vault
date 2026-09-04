# Complete Linux & DevOps Commands Cheat Sheet

## 1. Basic System & Info Commands
| Command | What it does |
| :--- | :--- |
| `pwd` | Shows your current working directory |
| `uname -a`, `uname -r` | Shows kernel and system information |
| `date` | Shows current date and time |
| `cal` | Shows the calendar |
| `clear` | Clears the terminal screen |
| `history` | Shows the list of previously run commands |
| `echo "text"` | Prints text to the terminal |
| `man command` | Opens the manual/help page for a command |
| `exit` | Closes the terminal session |

## 2. File & Directory Management
| Command | What it does |
| :--- | :--- |
| `pwd` | Shows current directory |
| `ls`, `ls -l`, `ls -la`, `ls -h` | Lists files (long format, hidden files, human-readable size) |
| `cd`, `cd ..`, `cd ~` | Changes directory, goes up one level, goes home |
| `tree` | Shows folder structure as a tree |
| `find /path -name "file"` | Searches for files by name |
| `locate filename` | Fast file search using an indexed database |
| `which command` | Shows the executable path of a command |
| `whereis command` | Shows binary, source, and man page locations |
| `touch filename` | Creates an empty file |
| `mkdir folder`, `mkdir -p a/b/c` | Creates a directory (with parent folders) |
| `cp file1 file2` | Copies a file |
| `cp -r folder1 folder2` | Copies a folder recursively |
| `mv file1 file2` | Moves or renames a file/folder |
| `rm file` | Deletes a file |
| `rm -r folder`, `rm -rf folder` | Deletes a folder (recursive, force, no confirmation) |
| `rmdir folder` | Removes an empty directory |

## 3. Viewing and Editing Files
| Command | What it does |
| :--- | :--- |
| `cat file` | Shows full file content |
| `less file`, `more file` | Views file page by page |
| `head file`, `head -n 20 file` | Shows first lines of a file |
| `tail file`, `tail -f file` | Shows last lines, `-f` follows live updates |
| `nano file` | Simple terminal text editor |
| `vim file` | Advanced terminal text editor |
| `wc file`, `wc -l file` | Counts words, lines, characters |
| `diff file1 file2` | Compares differences between two files |

## 4. Permissions and Ownership
| Command | What it does |
| :--- | :--- |
| `chmod 755 file` | Changes file permissions |
| `chmod +x file` | Makes a file executable |
| `chown user:group file` | Changes file owner and group |
| `chgrp group file` | Changes group ownership |
| `ls -l` | Shows permissions in rwx format |
| `umask` | Shows or sets default permission mask |

## 5. User & Group Management
| Command | What it does |
| :--- | :--- |
| `useradd username`, `adduser` | Creates a new user |
| `passwd username` | Sets or changes a user's password |
| `userdel username` | Deletes a user |
| `usermod` | Modifies user account details |
| `groupadd groupname` | Creates a new group |
| `groupdel groupname` | Deletes a group |
| `groups` | Checks groups of the current user |
| `whoami` | Shows current logged-in user |
| `who` | Shows all active users on the system |
| `id` | Shows user ID and group IDs |
| `su username` | Switches to another user |
| `sudo command` | Runs a command with administrator privileges |

## 6. Process & Resource Monitoring
| Command | What it does |
| :--- | :--- |
| `ps`, `ps aux`, `ps -ef` | Lists running processes |
| `top`, `htop` | Live interactive view of running processes |
| `kill PID` | Stops a process by its ID |
| `kill -9 PID` | Force kills a process |
| `killall name` | Kills processes by name |
| `jobs` | Lists background jobs in current shell |
| `bg`, `fg` | Sends a job to background/foreground |
| `free -h` | Shows memory (RAM and Swap) usage |
| `uptime` | Shows how long the system has been running |
| `vmstat` | Reports virtual memory statistics |
| `nohup command &` | Keeps process running after terminal closes |

## 7. Basic Networking
| Command | What it does |
| :--- | :--- |
| `ping host` | Checks if a host is reachable |
| `ifconfig`, `ip a` | Shows network interfaces and IPs |
| `curl -IL url` | Fetches data/headers from a URL |
| `wget url` | Downloads a file from a URL |
| `hostname` | Shows or sets the machine's hostname |
| `netstat -tulnp`, `ss -tulnp` | Shows open ports and listening services |
| `traceroute` | Traces network packet path to a host |
| `nslookup`, `dig` | Performs DNS lookups |

## 8. Package Management & Services
| Command | What it does |
| :--- | :--- |
| `apt update && apt upgrade` | Updates packages (Debian/Ubuntu) |
| `apt install package` | Installs a package (Debian/Ubuntu) |
| `yum install`, `dnf install` | Installs a package (RHEL/CentOS/Fedora) |
| `apt remove package` | Removes a package |
| `dpkg -l`, `rpm -qa` | Lists installed packages |
| `systemctl start service` | Starts a system service |
| `systemctl stop service` | Stops a system service |
| `systemctl restart service` | Restarts a system service |
| `systemctl status service` | Checks status of a service |
| `systemctl enable service` | Enables service to start on boot |

## 9. Advanced Text Processing & Search (AWK, GREP, SED)
| Command | What it does |
| :--- | :--- |
| `grep "pattern" file` | Searches for specific text or patterns in files |
| `grep -i`, `grep -r` | Case-insensitive search, recursive search in directories |
| `awk '{print $1}'` | Advanced text processing and column extraction |
| `sed 's/old/new/g'` | Stream editor to find and replace text |
| `cut -d',' -f1 file` | Extracts specific columns from files/output |
| `sort` | Sorts lines of text files |
| `uniq` | Filters out or detects duplicate lines |
| `tr` | Translates, squeezes, or deletes characters |
| `tee file` | Outputs to screen and saves to a file simultaneously |
| `xargs` | Builds and executes command-line arguments |

## 10. Compression and Archiving
| Command | What it does |
| :--- | :--- |
| `tar -cvf archive.tar folder` | Creates a uncompressed tar archive |
| `tar -xvf archive.tar` | Extracts a tar archive |
| `tar -czvf archive.tar.gz folder` | Creates a compressed `.tar.gz` archive |
| `tar -xzvf archive.tar.gz` | Extracts a `.tar.gz` archive |
| `zip -r archive.zip folder` | Creates a zip archive |
| `unzip archive.zip` | Extracts a zip archive |

## 11. System Information & Disk Usage
| Command | What it does |
| :--- | :--- |
| `uname -a` | Shows kernel and system info |
| `whoami` | Shows current logged-in user |
| `date` | Shows current date and time |
| `uptime` | Shows system uptime |
| `df -h` | Shows disk space usage in human-readable format |
| `du -sh folder` | Shows total size of a specific folder |
| `free -h` | Shows memory usage |

## 12. Disk Partitioning & LVM (Logical Volume Manager)
| Command | What it does |
| :--- | :--- |
| `fdisk -l` | Lists all disk partitions |
| `lsblk` | Shows block devices and partitions in a tree layout |
| `blkid` | Shows UUIDs and file system types of partitions |
| `pvcreate`, `pvs`, `pvdisplay` | Creates and manages Physical Volumes (PV) |
| `vgcreate`, `vgs`, `vgdisplay` | Creates and manages Volume Groups (VG) |
| `lvcreate`, `lvs`, `lvdisplay` | Creates and manages Logical Volumes (LV) |
| `lvextend` | Extends the size of a Logical Volume |
| `mkfs.ext4 /dev/sdX` | Formats a partition with ext4 file system |
| `mount`, `umount` | Mounts or unmounts a storage partition |
| `resize2fs` | Resizes the file system size after extending LV |
