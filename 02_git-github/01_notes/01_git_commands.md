# Git & Linux Command Reference Guide

| Command | Explanation & Use Case |
| :--- | :--- |
| `git --version` | Displays the currently installed version of Git to verify installation and check compatibility. |
| `git config --global user.name "Name"` | Sets the global username associated with all your local Git commits across the system. |
| `git config --global user.email "Email"` | Sets the global email address attached to your commits for author identification. |
| `git config --list` | Lists all current Git configuration settings loaded across local, global, and system scopes. |
| `git init` | Initializes a brand new, empty Git repository inside the current local project directory. |
| `ls -la` | Lists all files (including hidden files starting with `.`) in detailed long format. |
| `git status` | Displays the current state of the working directory, showing staged, unstaged, and untracked files. |
| `git add <file>` | Stages specific changes or files from the working directory, preparing them for the next commit. |
| `git add .` | Stages all modified, new, and deleted files in the current directory and subdirectories. |
| `git commit -m "msg"` | Records the staged snapshots permanently into the local repository history with a descriptive message. |
| `git log` | Displays the chronological commit history of the repository with hashes, authors, and dates. |
| `git log --oneline` | Displays a condensed, single-line summary view of the commit history for quick scanning. |
| `git diff` | Shows unstaged line-by-line differences between your working directory and the last commit. |
| `git checkout <hash>` | Switches the working tree to a specific historical commit or branch (used for inspecting past states). |
| `git branch` | Lists all local branches in the repository and highlights the currently active branch. |
| `git branch <name>` | Creates a new branch pointing to the current commit without switching over to it. |
| `git switch <name>` | Safely switches your working directory context to an existing target branch. |
| `git switch -c <name>` | Creates a brand new branch and immediately switches your working context into it. |
| `git merge <name>` | Merges history changes from the specified branch into your current active branch. |
| `git remote add origin <url>` | Links your local repository to a remote server repository URL under the shortname `origin`. |
| `git remote -v` | Verifies and lists all remote repository URLs configured for fetch and push operations. |
| `git push origin <branch>` | Uploads your local branch commits and history up to the remote repository server. |
| `git push -u origin <branch>` | Pushes local branch commits and links the local branch to track the remote upstream branch. |
| `git pull origin <branch>` | Fetches and automatically merges changes from the specified remote branch into your local branch. |
| `git clone <url>` | Downloads a complete copy of an existing remote repository, including all files and history. |
| `git stash` | Temporarily shelves your modified, unstaged changes to clear your workspace. |
| `git stash pop` | Restores the most recently stashed changes back into your working directory and removes them from the stash stack. |
| `git reset <hash>` | Resets your current branch pointer back to a specified commit while keeping working changes intact. |
| `git reset --hard <hash>` | Completely resets your branch pointer and wipes out all working directory changes to match a past commit. |
| `git revert <hash>` | Creates a brand new commit that explicitly undoes the changes introduced by a specific prior commit. |
| `git config --global --unset <key>` | Removes a specific global configuration setting from your Git environment. |
| `git rm <file>` | Removes files from your working tree and stages the removal for the next commit. |
| `git rm --cached <file>` | Stops tracking a file and removes it from the staging area while keeping it locally on disk. |
| `git mv <old> <new>` | Moves or renames a file or directory and automatically stages the change. |
| `git stash list` | Lists all current stashed changes currently saved in your local stash stack queue. |
| `git stash drop` | Manually deletes a specific or the most recent stash entry from your stash stack. |
| `git stash clear` | Completely deletes and wipes out all saved stashes from your local storage stack. |
| `git tag <tag-name>` | Creates a lightweight pointer tag at the current commit to mark important releases. |
| `git tag -a <tag> -m "msg"` | Creates an annotated tag with metadata, author info, and a custom message attached. |
| `git push origin --tags` | Pushes and uploads all your local release tags up to the remote repository. |
| `git fetch` | Downloads all history, branches, and objects from the remote repository without merging. |
