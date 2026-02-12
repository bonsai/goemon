import os
import json
import numpy as np

class VectorMemoryAgent:
    """
    メタデータの検索とベクトル管理を行うエージェント
    """
    def __init__(self):
        self.db_dir = "c:/Users/dance/zone/uniai/data/vlm_db"
        self.memories = []

    def load_all_memories(self):
        """
        保存されたすべてのメタデータをロード
        """
        self.memories = []
        if not os.path.exists(self.db_dir):
            return
            
        for file in os.listdir(self.db_dir):
            if file.endswith(".json"):
                with open(os.path.join(self.db_dir, file), "r", encoding="utf-8") as f:
                    self.memories.append(json.load(f))
        print(f">>> Loaded {len(self.memories)} memories from DB.")

    def search(self, query):
        """
        簡易的なキーワード検索（将来的にベクトル検索にアップグレード）
        """
        self.load_all_memories()
        results = []
        query = query.lower()
        
        for mem in self.memories:
            score = 0
            # 説明文、OCRテキスト、タグから検索
            text_pool = (mem['description'] + " " + mem['ocr_text'] + " " + " ".join(mem['tags'])).lower()
            
            if query in text_pool:
                score += 1
                results.append({"memory": mem, "score": score})
                
        # スコア順にソート
        results.sort(key=lambda x: x['score'], reverse=True)
        return [r['memory'] for r in results]

if __name__ == "__main__":
    # 検索のテスト
    # store = MemoryVectorStore()
    # print(store.search("OneDrive"))
    pass
