import os
from kaggle_secrets import UserSecretsClient

# --- 1. 環境変数のセットアップ ---
user_secrets = UserSecretsClient()
github_token = user_secrets.get_secret("GITHUB_TOKEN")
os.environ['NGROK_AUTHTOKEN'] = user_secrets.get_secret("NGROK_AUTHTOKEN")
os.environ['HF_TOKEN'] = user_secrets.get_secret("HF_TOKEN")

# --- 2. リポジトリのクローンと移動 ---
if not os.path.exists('goemon'):
    print("Cloning repository...")
    !git clone https://{github_token}@github.com/bonsai/goemon
    %cd goemon
else:
    print("Repository already exists. Moving to directory.")
    %cd goemon
    # すでにクローン済みの場合は最新を取得
    !git pull origin main

# --- 3. 実行権限の付与と起動 ---
# スクリプト自体が重複チェック（Goのインストール、モデルの存在確認）を行っています
!chmod +x linux-start.sh
!bash linux-start.sh