import os
from kaggle_secrets import UserSecretsClient
from IPython import get_ipython

# --- 1. 環境変数のセットアップ ---
user_secrets = UserSecretsClient()
ipython = get_ipython()

def set_env_secret(env_name, secret_name):
    val = user_secrets.get_secret(secret_name)
    if val:
        os.environ[env_name] = val
        # Jupyter/Kaggle のシェル環境にも確実に反映させる
        if ipython:
            ipython.run_line_magic('env', f'{env_name}={val}')
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

# --- 4. ngrok URL の監視と表示 ---
import time
import threading
from IPython.display import display, HTML

def watch_ngrok_url():
    url_file = 'ngrok_url.txt'
    # 古いファイルがあれば消しておく
    if os.path.exists(url_file):
        os.remove(url_file)
        
    print("⏳ Waiting for ngrok URL...")
    while True:
        if os.path.exists(url_file):
            try:
                with open(url_file, 'r') as f:
                    url = f.read().strip()
                if url:
                    display(HTML(f"""
                        <div style="padding:20px; background-color:#e1f5fe; border-radius:10px; border:2px solid #01579b; margin:20px 0;">
                            <h2 style="color:#01579b; margin-top:0;">🚀 Goemon Swarm Online!</h2>
                            <p>Mobile/Public URL: <a href="{url}?pass=1234" target="_blank" style="font-size:1.2em; font-weight:bold; color:#d81b60;">{url}</a></p>
                            <p style="font-size:0.9em; color:#555;">(Auto-login pass included)</p>
                        </div>
                    """))
                    break
            except Exception as e:
                pass
        time.sleep(1)

# 監視スレッドを開始
threading.Thread(target=watch_ngrok_url, daemon=True).start()

# 起動
!bash linux-start.sh
