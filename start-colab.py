import os

try:
    from google.colab import userdata
    IN_COLAB = True
except ImportError:
    IN_COLAB = False

# --- 1. 環境変数のセットアップ ---
if IN_COLAB:
    print("Running on Google Colab")
    from IPython import get_ipython
    ipython = get_ipython()

    def set_colab_env(env_name, secret_name):
        try:
            val = userdata.get(secret_name)
            if val:
                os.environ[env_name] = val
                if ipython:
                    ipython.run_line_magic('env', f'{env_name}={val}')
                print(f"✅ {env_name} is set.")
            else:
                print(f"⚠️ {env_name} is NOT set in Colab Secrets.")
        except Exception:
            print(f"⚠️ Could not find secret {secret_name}")

    set_colab_env('NGROK_AUTHTOKEN', 'NGROK_AUTHTOKEN')
    set_colab_env('HF_TOKEN', 'HF_TOKEN')
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

# --- 4. ngrok URL の監視と表示 ---
import time
import threading
from IPython.display import display, HTML

def watch_ngrok_url():
    url_file = 'ngrok_url.txt'
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
                        <div style="padding:20px; background-color:#f1f8e9; border-radius:10px; border:2px solid #33691e; margin:20px 0;">
                            <h2 style="color:#33691e; margin-top:0;">🚀 Goemon Swarm Online! (Colab)</h2>
                            <p>Mobile/Public URL: <a href="{url}?pass=1234" target="_blank" style="font-size:1.2em; font-weight:bold; color:#d81b60;">{url}</a></p>
                            <p style="font-size:0.9em; color:#555;">(Auto-login pass included)</p>
                        </div>
                    """))
                    break
            except Exception as e:
                pass
        time.sleep(1)

threading.Thread(target=watch_ngrok_url, daemon=True).start()

!bash linux-start.sh
