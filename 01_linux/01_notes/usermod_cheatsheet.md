# 🚀 Linux `usermod` Command Master Guide & Flags Chart

> **Description:** Comprehensive reference chart for modifying user account attributes in Linux. Designed for Markdown notes (`.md`) to be viewed or edited in Vim/Neovim.

---

## 📋 Quick Reference: `usermod` Flags Chart

| Flag / Badge | Description & Purpose | Syntax / Example |
| :--- | :--- | :--- |
| `<code>-c, --comment</code>` | Changes the **GECOS field** (user description/comment in `/etc/passwd`). | `sudo usermod -c "Senior DevOps Engineer" chetan` |
| `<code>-d, --home</code>` | Changes the user's **home directory** path. (Use <code>-m</code> to move contents). | `sudo usermod -d /home/nini -m chetan` |
| `<code>-e, --expiredate</code>` | Sets the **account expiration date** (Format: `YYYY-MM-DD`). | `sudo usermod -e 2026-12-31 chetan` |
| `<code>-f, --inactive</code>` | Sets days of **password inactivity** before account locks. | `sudo usermod -f 30 chetan` |
| `<code>-g, --gid</code>` | Changes the user's **primary group** (by GID or group name). | `sudo usermod -g chetanpur chetan` |
| `<code>-G, --groups</code>` | Defines the list of **supplementary/secondary groups** (replaces existing). | `sudo usermod -G sudo,docker chetan` |
| `<code>-a, --append</code>` | **Must be used with -G** to append groups without removing current ones. | `sudo usermod -aG docker chetan` |
| `<code>-l, --login</code>` | Changes the user's **login name** (username). | `sudo usermod -l newname oldname` |
| `<code>-L, --lock</code>` | **Locks** the user account (adds `!` before password hash in `/etc/shadow`). | `sudo usermod -L chetan` |
| `<code>-U, --unlock</code>` | **Unlocks** a locked user account (removes the `!` prefix). | `sudo usermod -U chetan` |
| `<code>-m, --move-home</code>` | Moves current home directory contents to the new home dir (used with `-d`). | `sudo usermod -d /home/new -m chetan` |
| `<code>-p, --password</code>` | Sets the **encrypted password** directly (not recommended for plaintext). | `sudo usermod -p '$6$...' chetan` |
| `<code>-s, --shell</code>` | Changes the user's **default login shell** (e.g., `/bin/bash`, `/bin/zsh`). | `sudo usermod -s /bin/bash chetan` |
| `<code>-u, --uid</code>` | Changes the user's **UID** (User ID number). | `sudo usermod -u 1005 chetan` |

---

## ⚡ Important Best Practices & Tips

1. **The Dangerous `-G` vs Safe `-aG` Rule:**
   - ⚠️ `sudo usermod -G group1 user` will **overwrite** all supplementary groups and remove the user from any other groups they were part of.
   - ✅ Always use `sudo usermod -aG group1 user` (`-a` for append) when adding a user to a new group (like `docker` or `sudo`).

2. **Verifying Changes:**
   - Check account info: `cat /etc/passwd | grep username`
   - Check password/expiry info: `sudo chage -l username`
   - Check group memberships: `groups username`
