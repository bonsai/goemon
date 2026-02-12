# Google Colab Swarm Console Setup Guide

Google Colab で `goemon` を起動するための手順です。

## 1. 事前準備 (Colab Secrets)
1. 左側のパネルにある **鍵アイコン (Secrets)** をクリックします。
2. 以下のシークレットを追加し、**Notebook access** をオンにします：
   - `NGROK_AUTHTOKEN`: [ngrok ダッシュボード](https://dashboard.ngrok.com/get-started/your-authtoken)から取得。
   - `GITHUB_TOKEN`: GitHub の `repo` 権限付きトークン。
   - `HF_TOKEN`: (任意) Hugging Face のトークン。

## 2. 起動用コード
以下のコードをセルに貼り付けて実行してください。

```python
import os
from google.colab import userdata

# --- 1. 環境変数のセットアップ ---
os.environ['NGROK_AUTHTOKEN'] = userdata.get("NGROK_AUTHTOKEN")
os.environ['HF_TOKEN'] = userdata.get("HF_TOKEN")
github_token = userdata.get("GITHUB_TOKEN")

# --- 2. リポジトリのクローンと移動 ---
if not os.path.exists('goemon'):
    !git clone https://{github_token}@github.com/bonsai/goemon
%cd goemon
!git pull origin main

# --- 3. 実行権限の付与と起動 ---
!chmod +x linux-start.sh
!bash linux-start.sh
```

## 3. Web UI へのアクセス
1. 実行後、ログに表示される **QR コード** または **ngrok の URL** をクリックします。
2. パスワード `1234` でログインします。

## 注意点
- **ランタイムのタイプ**: GPU (T4 など) を使用することをお勧めします。
- **ディスク**: Colab は Kaggle よりディスク容量に余裕があるため、多くのモデルを試せます。
