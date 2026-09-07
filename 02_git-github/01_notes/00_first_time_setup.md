# 🚀 Git & GitHub: Quick Reference Guide
*From Zero to First Push*

<div align="center">

[![Git](https://img.shields.io/badge/Tool-Git-F05032?style=for-the-badge&logo=git&logoColor=white)]()
[![GitHub](https://img.shields.io/badge/Platform-GitHub-181717?style=for-the-badge&logo=github&logoColor=white)]()
[![Status](https://img.shields.io/badge/Status-Beginner%20Friendly-brightgreen?style=for-the-badge)]()
[![DevOps](https://img.shields.io/badge/Track-DevOps-2CA5E0?style=for-the-badge&logo=databricks&logoColor=white)]()

---



---



</div>


<p align="center">
  <img src="https://img.shields.io/badge/STEP%200-DOWNLOAD%20%26%20ACCOUNTS-blue?style=for-the-badge&logo=git&logoColor=white" />
</p>

<p align="left">
  <a href="https://git-scm.com"><img src="https://img.shields.io/badge/1.%20Install%20Git%20Bash-F05032?style=flat-square&logo=git&logoColor=white" /></a>
  <a href="https://github.com"><img src="https://img.shields.io/badge/2.%20GitHub%20Account-181717?style=flat-square&logo=github&logoColor=white" /></a>
</p>

* **Generate PAT Token:** Profile ➔ Settings ➔ Developer settings ➔ Personal access tokens (classic) ➔ Generate new token ➔ Tick `repo` checkbox ➔ Generate token
* ⚠️ **Copy & save the token immediately!**

---

---

<p align="center">
  <img src="https://img.shields.io/badge/STEP%201-COMPUTER%20SETUP-blueviolet?style=for-the-badge&logo=git&logoColor=white" />
</p>

<p align="left">
  <img src="https://img.shields.io/badge/Config-Global%20User%20%26%20Email-success?style=flat-square" />
</p>


git config --global user.name "Your GitHub Username"
git config --global user.email "your-email@example.com"

---


---


</div>

<p align="center">
  <img src="https://img.shields.io/badge/STEP%202-INITIALIZE%20PROJECT-orange?style=for-the-badge&logo=git&logoColor=white" />
</p>

* Right-click inside your project folder ➔ Open **Git Bash here**

<p align="left">
  <img src="https://img.shields.io/badge/Command-git%20init-yellow?style=flat-square" />
  <img src="https://img.shields.io/badge/Command-git%20add%20.%20%26%20commit-orange?style=flat-square" />
</p>

```bash
git init
git status
git add .
git commit -m "First commit"