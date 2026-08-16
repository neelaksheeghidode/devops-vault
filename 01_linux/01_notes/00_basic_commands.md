# BASIC COMMANDS

### 1.1 Moving Around the File System [ ]

| Command | What it does |
| --- | --- |
| `pwd` | Shows your current directory |
| `ls`, `ls -l`, `ls -la` | Lists files (long format, includes hidden files) |
| `cd`, `cd ..`, `cd ~` | Changes directory, goes up one level, goes home |
| `tree` | Shows folder structure as a tree |
| `find /path -name "file"` | Searches for files by name |
| `locate filename` | Fast file search using an indexed database |
| `which command` | Shows the path of a command |
| `whereis command` | Shows binary, source, and man page locations |

### 1.2 Creating, Copying, Moving, Deleting [x]

| Command | What it does |
| --- | --- |
| `touch file` | Creates an empty file |
| `mkdir folder`, `mkdir -p a/b/c` | Creates a directory (with parent folders) |
| `cp file1 file2` | Copies a file |
| `cp -r folder1 folder2` | Copies a folder recursively |
| `mv file1 file2` | Moves or renames a file |
| `rm file` | Deletes a file |
| `rm -r folder`, `rm -rf folder` | Deletes a folder (force, no confirmation) |
| `rmdir folder` | Removes an empty folder |

### 1.3 Viewing and Editing Files [ ]

| Command | What it does |
| --- | --- |
| `cat file` | Shows full file content |
| `less file`, `more file` | Views file page by page |
| `head file`, `head -n 20 file` | Shows first lines of a file |
| `tail file`, `tail -f file` | Shows last lines, `-f` follows live updates |
| `nano file` | Simple terminal text editor |
| `vim file` | Advanced terminal text editor |
| `wc file`, `wc -l file` | Counts words, lines, characters |
| `diff file1 file2` | Compares two files |

### 1.4 Permissions and Ownership [ ]

| Command | What it does |
| --- | --- |
| `chmod 755 file` | Changes file permissions |
| `chmod +x file` | Makes a file executable |
| `chown user:group file` | Changes file owner and group |
| `ls -l` | Shows permissions in `rwx` format |
| `umask` | Shows or sets default permission mask |

### 1.5 Process Basics [ ]

| Command | What it does |
| --- | --- |
| `ps`, `ps aux` | Lists running processes |
| `top` | Live view of running processes |
| `kill PID` | Stops a process by its ID |
| `kill -9 PID` | Force kills a process |
| `jobs` | Lists background jobs in current shell |
| `bg`, `fg` | Sends a job to background/foreground |

### 1.6 Basic Networking [ ]

| Command | What it does |
| --- | --- |
| `ping host` | Checks if a host is reachable |
| `ifconfig` / `ip a` | Shows network interfaces and IPs |
| `curl url` | Fetches data from a URL |
| `wget url` | Downloads a file from a URL |
| `hostname` | Shows the machine's hostname |
| `netstat -tulnp` | Shows open ports and listening services |

### 1.7 Package Management [ ]

| Command | What it does |
| --- | --- |
| `apt update && apt upgrade` | Updates packages (Debian/Ubuntu) |
| `apt install package` | Installs a package (Debian/Ubuntu) |
| `yum install package` / `dnf install package` | Installs a package (RHEL/CentOS/Fedora) |
| `apt remove package` | Removes a package |
| `dpkg -l` / `rpm -qa` | Lists installed packages |

### 1.8 Compression and Archiving [ ]

| Command | What it does |
| --- | --- |
| `tar -cvf archive.tar folder` | Creates a tar archive |
| `tar -xvf archive.tar` | Extracts a tar archive |
| `tar -czvf archive.tar.gz folder` | Creates a compressed tar.gz archive |
| `tar -xzvf archive.tar.gz` | Extracts a tar.gz archive |
| `zip -r archive.zip folder` | Creates a zip archive |
| `unzip archive.zip` | Extracts a zip archive |

### 1.9 System Information [ ]

| Command | What it does |
| --- | --- |
| `uname -a` | Shows kernel and system info |
| `whoami` | Shows current logged-in user |
| `date` | Shows current date and time |
| `uptime` | Shows how long the system has been running |
| `df -h` | Shows disk space usage |
| `du -sh folder` | Shows size of a folder |
| `free -h` | Shows memory usage |
