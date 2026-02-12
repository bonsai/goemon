#!/bin/bash

# =================================================================
# Goemon Swarm Console - Kaggle Auto-Launcher
# =================================================================

echo "Starting Goemon Swarm on Kaggle..."

# 0. Go のインストール (Kaggle には標準で入っていないため)
if ! command -v go &> /dev/null; then
    echo "Step 0: Installing Go..."
    GO_VERSION="1.22.0"
    wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    rm go${GO_VERSION}.linux-amd64.tar.gz
    go version
else
    echo "Go is already installed."
fi

# 1. モデルのダウンロード (GLM, Anything, Moondream2)
echo "Step 1: Downloading models..."
python3 download_models.py

# 2. Go アプリケーションのビルド
echo "Step 2: Building Goemon..."
# Linux 環境なので .exe は不要
go build -o goemon src/cmd/goemon/main.go
chmod +x goemon

# 3. Goemon 実行
echo "Step 3: Launching Goemon Swarm..."
# NGROK_AUTHTOKEN は環境変数として渡される前提
./goemon
