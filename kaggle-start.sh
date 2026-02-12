#!/bin/bash

# =================================================================
# Goemon Swarm Console - Kaggle Auto-Launcher
# =================================================================

echo "Starting Goemon Swarm on Kaggle..."

# 1. モデルのダウンロード (GLM, Anything, Moondream2)
echo "Step 1: Downloading models..."
python3 download_models.py

# 2. Go アプリケーションのビルド
echo "Step 2: Building Goemon..."
go build -o goemon.exe src/cmd/goemon/main.go

# 3. Goemon 実行
echo "Step 3: Launching Goemon Swarm..."
# NGROK_AUTHTOKEN は環境変数として渡される前提
./goemon.exe
