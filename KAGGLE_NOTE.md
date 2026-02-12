# Kaggle Swarm Console Setup Guide

Git リポジトリから直接 `goemon` を起動するための Kaggle Notebook 手順です。

## 1. 事前準備 (Kaggle Secrets)
リポジトリがプライベートなため、クローンには GitHub のトークンが必要です。
1. [GitHub Settings -> Tokens (classic)](https://github.com/settings/tokens) で `repo` 権限のあるトークンを作成。
2. Kaggle の **Add-ons -> Secrets** に以下を登録：
   - `NGROK_AUTHTOKEN`: [ngrok ダッシュボード](https://dashboard.ngrok.com/get-started/your-authtoken)から取得。
   - `GITHUB_TOKEN`: 作成した GitHub トークン。

## 2. 起動用コード (Consolidated Python Cell)
以下のコードを Kaggle のセルに貼り付けて実行してください。これ一つでセットアップから起動まで完了します。

```python
import os
from kaggle_secrets import UserSecretsClient

# --- 1. 環境変数のセットアップ ---
user_secrets = UserSecretsClient()
# 各種トークンの取得
github_token = user_secrets.get_secret("GITHUB_TOKEN")
os.environ['NGROK_AUTHTOKEN'] = user_secrets.get_secret("NGROK_AUTHTOKEN")
os.environ['HF_TOKEN'] = user_secrets.get_secret("HF_TOKEN")

# --- 2. リポジトリのクローンと移動 ---
if not os.path.exists('goemon'):
    print("Cloning repository...")
    !git clone https://{github_token}@github.com/bonsai/goemon
else:
    print("Repository already exists. Skipping clone.")

%cd goemon

# 最新の状態に更新 (必要に応じて)
!git pull origin main

# --- 3. 実行権限の付与と起動 ---
# このスクリプトが Go のインストール、モデルのダウンロード、ビルド、起動をすべて行います
!chmod +x kaggle-start.sh
!bash kaggle-start.sh
```

## 3. Web UI へのアクセス
1. スクリプトの出力ログに表示される **QR コード** をスマホでスキャンします。
2. または、`https://xxxx-xxxx.ngrok-free.app` という URL をクリックします。
3. パスワード `1234` で自動ログインされ、管理画面が開きます。

## 注意点
- Kaggle の **Settings -> Internet** が `On` になっていることを確認してください。
- GPU を使用する場合は **Settings -> Accelerator** を `T4 x2` または `P100` に設定してください。
- `kaggle-start.sh` はバックグラウンドで ngrok を動かし続けるため、セルは実行状態のままになります。
