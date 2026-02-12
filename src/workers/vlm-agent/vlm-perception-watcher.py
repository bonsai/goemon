import os
import time
import importlib
import sys

# 自ディレクトリをパスに追加してハイフン付きファイルをインポート可能にする
sys.path.append(os.path.dirname(__file__))

# 命名規則（カバブケース）に合わせて動的にインポート
image_perception_engine = importlib.import_module("image-perception-engine")
ImagePerceptionEngine = image_perception_engine.ImagePerceptionEngine

vlm_mail_responder = importlib.import_module("vlm-mail-responder")
VlmMailResponder = vlm_mail_responder.VlmMailResponder

class VlmPerceptionWatcher:
    """
    フォルダを監視し、画像解析とメール応答を統合するメインオーケストレーター
    """
    def __init__(self):
        self.watch_dir = "c:/Users/dance/zone/uniai/data/vlm_watch"
        self.db_dir = "c:/Users/dance/zone/uniai/data/vlm_db"
        self.engine = ImagePerceptionEngine()
        self.responder = VlmMailResponder()

    def run_forever(self):
        print(f"[*] VLM Perception Watcher started.")
        print(f"[*] Watching: {self.watch_dir}")
        
        if not os.path.exists(self.watch_dir):
            os.makedirs(self.watch_dir)

        while True:
            # 1. 新しい画像の解析
            self._check_new_images()
            
            # 2. メール問い合わせの処理
            self.responder.process_queries()
            
            time.sleep(10)

    def _check_new_images(self):
        files = [f for f in os.listdir(self.watch_dir) if f.lower().endswith(('.png', '.jpg', '.jpeg', '.pdf'))]
        for f in files:
            image_path = os.path.join(self.watch_dir, f)
            metadata_path = os.path.join(self.db_dir, f"{f}.json")
            
            # まだ解析されていない画像があれば処理
            if not os.path.exists(metadata_path):
                try:
                    self.engine.perceive(image_path)
                except Exception as e:
                    print(f"[!] Error processing {f}: {e}")

if __name__ == "__main__":
    watcher = VlmPerceptionWatcher()
    watcher.run_forever()
