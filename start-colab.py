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
!bash linux-start.sh
