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

# 3. ngrok のセットアップ (ポート 8080 を外部公開)
# Kaggle のシークレットから NGROK_AUTH_TOKEN を取得することを推奨
if [ -z "$NGROK_AUTH_TOKEN" ]; then
    echo "Warning: NGROK_AUTH_TOKEN is not set. ngrok might not start correctly."
    echo "Please set it via Kaggle Secrets or environment variable."
else
    echo "Step 3: Setting up ngrok tunnel..."
    curl -s https://ngrok-agent.s3.amazonaws.com/ngrok.asc | sudo tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null
    echo "deb https://ngrok-agent.s3.amazonaws.com buster main" | sudo tee /etc/apt/sources.list.d/ngrok.list
    sudo apt update && sudo apt install ngrok
    ngrok config add-authtoken $NGROK_AUTH_TOKEN
    
    # バックグラウンドで ngrok を起動
    ngrok http 8080 --log=stdout > ngrok.log &
    sleep 5
    
    # 公開 URL を表示
    PUBLIC_URL=$(curl -s http://localhost:4040/api/tunnels | grep -o 'https://[^"]*')
    echo "===================================================="
    echo "GOEMON WEB CONSOLE IS LIVE AT:"
    echo "$PUBLIC_URL"
    echo "===================================================="
fi

# 4. Goemon 実行
echo "Step 4: Launching Goemon Swarm..."
./goemon.exe
