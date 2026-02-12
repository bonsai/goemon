# Pull Request: UniAI to Goemon Transformation & Web Swarm Console Implementation

## 概要
本プロジェクトを `uniai` から `goemon` へとリブランドし、コアシステムの復旧、および Kaggle 等のクラウド環境で GPU 資源を最大限活用するための **Web Swarm Console** を実装しました。

## 主な変更点

### 1. プロジェクトのリブランドと基盤修復
- **リネーム**: 全ての Go インポートパスを `uniai` から `goemon` へ置換。
- **ファイル復元**: 破損していた 17 個の `.go` ファイルのうち、DB、エージェント、CLI に関わる主要な 12 ファイルを ADR に基づき再実装。
- **README 更新**: プロジェクト名を `goemon` に変更し、アーキテクチャ図とディレクトリ構造を最新化。

### 2. Web Swarm Console (Go-based Web UI)
Node.js を一切使用せず、Go の標準機能 (`html/template`) のみで構築された軽量な Web GUI を実装しました。
- **Anything V5 統合**: プロンプト 1 回の送信で 9 枚の画像を同時生成。
- **Moondream2 (VLM) 統合**: ギャラリーから画像を選択し、内容について質問（Perception）が可能。
- **GLM-4 統合**: ランダムなテキストや「呪文（Spell）」の生成を Web 越しに実行。
- **リアルタイム更新**: JavaScript による 5 秒おきポーリングにより、生成された画像を即座にギャラリーへ反映。

### 3. クラウド / Docker 対応
- **Multi-stage Dockerfile**: `goemon.exe` をビルドし、Web ポート 8080 を公開する Docker 設定。
- **モデル自動ダウンロード**: GLM, Anything, Moondream2 の重みを Hugging Face から自動取得する `download_models.py` を追加。
- **Kaggle 起動スクリプト**: `ngrok` を使用して、Kaggle の閉じた環境から外部ブラウザへセキュアにトンネルを繋ぐ `kaggle-start.sh` を作成。

## 動作確認方法
1. `go build -o goemon.exe src/cmd/goemon/main.go` でビルド。
2. `./goemon.exe` を実行し、ブラウザで `http://localhost:8080` にアクセス。
3. または Kaggle 上で `!bash kaggle-start.sh` を実行し、発行された ngrok URL にアクセス。

---
この PR により、CLI ベースだったシステムが、クラウド上の GPU を効率的に操れる強力な Web インターフェースを持つ Swarm システムへと進化しました。
