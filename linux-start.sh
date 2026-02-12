#!/bin/bash

# =================================================================
# Goemon Swarm Console - Cloud Auto-Launcher (Kaggle/Colab)
# =================================================================

echo "Starting Goemon Swarm on Cloud Environment..."

# 0. Go のインストール (標準で入っていない環境用)
export GO_HOME="$HOME/opt/go"
export PATH="$GO_HOME/bin:$PATH"

if ! command -v go &> /dev/null; then
    echo "Step 0: Installing Go locally..."
    GO_VERSION="1.22.0"
    mkdir -p "$HOME/opt"
    wget -q https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
    tar -C "$HOME/opt" -xzf go${GO_VERSION}.linux-amd64.tar.gz
    rm go${GO_VERSION}.linux-amd64.tar.gz
    go version
else
    echo "Go is already installed at $(command -v go)"
fi

# 1. モデルのダウンロード (GLM, Anything, Moondream2)
echo "Step 1: Checking models..."
# すでにモデルがある場合はスキップするロジックは python 側にあるためそのまま実行
python3 download_models.py

# 2. Go アプリケーションのビルド
echo "Step 2: Building Goemon..."
# 依存関係の解決
go mod tidy

go build -o goemon src/cmd/goemon/main.go
chmod +x goemon

# 3. Goemon 実行
echo "Step 3: Launching Goemon Swarm..."
# Linux 環境に合わせた実行
./goemon
