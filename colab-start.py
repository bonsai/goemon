import os

try:
    from google.colab import userdata
    IN_COLAB = True
except ImportError:
    IN_COLAB = False

# --- 1. 環境変数のセットアップ ---
if IN_COLAB:
    print("Running on Google Colab")
    os.environ['NGROK_AUTHTOKEN'] = userdata.get("NGROK_AUTHTOKEN")
    os.environ['HF_TOKEN'] = userdata.get("HF_TOKEN")
    github_token = userdata.get("GITHUB_TOKEN")
else:
    print("Not in Colab. Please use start.py for Kaggle.")
    exit(1)

# --- 2. リポジトリのクローンと移動 ---
if not os.path.exists('goemon'):
    print("Cloning repository...")
    !git clone https://{github_token}@github.com/bonsai/goemon
    %cd goemon
else:
    print("Repository already exists. Moving to directory.")
    %cd goemon
    !git pull origin main

# --- 3. 実行権限の付与と起動 ---
!chmod +x linux-start.sh
!bash linux-start.sh
