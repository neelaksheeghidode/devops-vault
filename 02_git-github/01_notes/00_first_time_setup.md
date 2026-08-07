# 🚀 Quick Reference Guide from Zero to First Push

---

## 🛠️ Step 0: Download & Create Accounts

- [ ] **Install Git Bash** → [git-scm.com](https://git-scm.com/)
- [ ] **Create GitHub Account** → [github.com](https://github.com/)
- [ ] **Generate Personal Access Token (PAT):**
  1. `Profile` ➔ `Settings` ➔ `Developer settings`
  2. `Personal access tokens (classic)` ➔ `Generate new token`
  3. Tick `repo` checkbox ➔ Click **Generate token**
  4. ⚠️ **Copy & save the token immediately!**



---

## 👤 Step 1: First-Time Computer Setup

```bash
git config --global user.name "Your GitHub Username"
git config --global user.email "your-email@example.com"
```



## 📁 Step 2: Initialize Project & Save Locally

Right-click in your project folder ➔ Open Git Bash here

```bash
git init
git status
git add .
git commit -m "First commit"
```


## 🌐 Step 3: Create Repository on GitHub

    Click + (Top Right) ➔ New repository

    Enter Repository name

    Select Public

    ❌ Do NOT check "Add a README file"

    Click Create repository

    Copy the repo URL (https://github.com/username/repo-name.git)



## 🔗 Step 4: Link & Push to GitHub

```Bash

git branch -M main
git remote add origin YOUR_COPIED_GITHUB_URL
git push -u origin main
```



 🔑 Note:

 When prompted, enter your GitHub Username, and paste your PAT Token as the password (text won't be visible while pasting).



## 🔄 Daily Workflow (For Every New Update)


```Bash

git add .
git commit -m "Your update message"
git push
```


