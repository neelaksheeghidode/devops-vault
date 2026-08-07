#!/bin/bash

echo "========================================"
echo "🚀 Starting Automated Git Push Cycle..."
echo "========================================"

# Ask for commit message
read -p "📝 Enter commit message: " msg

# Execute Git Commands
echo -e "\n⏳ Staging files..."
git add .

echo "💾 Committing changes..."
git commit -m "$msg"

echo "🚀 Pushing to GitHub..."
git push

echo "========================================"
echo "✅ Push Completed Successfully!"
echo "========================================"
