import os
from kaggle_secrets import UserSecretsClient

# --- 1. 環境変数のセットアップ ---
user_secrets = UserSecretsClient()

def set_env_secret(env_name, secret_name):
    val = user_secrets.get_secret(secret_name)
    if val:
        os.environ[env_name] = val
        print(f"✅ {env_name} is set.")
    else:
        print(f"⚠️ {env_name} is NOT set in Kaggle Secrets.")

set_env_secret('GITHUB_TOKEN', 'GITHUB_TOKEN')
set_env_secret('NGROK_AUTHTOKEN', 'NGROK_AUTHTOKEN')
set_env_secret('HF_TOKEN', 'HF_TOKEN')

github_token = os.environ.get('GITHUB_TOKEN', '')

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