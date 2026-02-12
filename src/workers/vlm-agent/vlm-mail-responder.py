import os
import json
import time
import importlib

# 命名規則（カバブケース）に合わせて動的にインポート
memory_vector_store = importlib.import_module("memory-vector-store")
MemoryVectorStore = memory_vector_store.MemoryVectorStore

class VlmMailResponder:
    """
    メール経由の問い合わせに回答するプロセッサ
    """
    def __init__(self):
        self.store = MemoryVectorStore()
        self.mail_in_dir = "c:/Users/dance/zone/uniai/data/mail_in"
        self.mail_out_dir = "c:/Users/dance/zone/uniai/data/mail_out"

    def process_queries(self):
        """
        未処理の検索メールを処理する
        """
        if not os.path.exists(self.mail_in_dir):
            os.makedirs(self.mail_in_dir)
        if not os.path.exists(self.mail_out_dir):
            os.makedirs(self.mail_out_dir)

        mails = [f for f in os.listdir(self.mail_in_dir) if f.endswith(".json")]
        for mail_file in mails:
            mail_path = os.path.join(self.mail_in_dir, mail_file)
            with open(mail_path, "r", encoding="utf-8") as f:
                data = json.load(f)

            if "query" in data:
                query = data["query"]
                print(f">>> Processing search query: {query}")
                results = self.store.search(query)
                
                self._send_response(mail_file, query, results)
                # 処理済みメールを削除またはアーカイブ
                os.remove(mail_path)

    def _send_response(self, original_filename, query, results):
        response_file = os.path.join(self.mail_out_dir, f"res_{original_filename}")
        
        if results:
            content = f"「{query}」に関連する写真が {len(results)} 件見つかりました。\n\n"
            for r in results[:3]: # 上位3件を表示
                content += f"- ファイル: {r['filename']}\n"
                content += f"  内容: {r['description'][:100]}...\n"
                content += f"  パス: {r['source_file']}\n\n"
        else:
            content = f"残念ながら「{query}」に関連する写真は見つかりませんでした。"

        response_data = {
            "to": "user",
            "subject": f"Re: {query} についての検索結果",
            "body": content,
            "timestamp": time.time()
        }

        with open(response_file, "w", encoding="utf-8") as f:
            json.dump(response_data, f, indent=4, ensure_ascii=False)
        print(f">>> Response sent: {response_file}")

if __name__ == "__main__":
    # responder = VlmMailResponder()
    # responder.process_queries()
    pass
