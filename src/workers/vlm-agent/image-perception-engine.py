import os
import json
import importlib
from datetime import datetime

# 命名規則（カバブケース）に合わせて動的にインポート
vlm_model_manager = importlib.import_module("vlm-model-manager")
MoondreamManager = vlm_model_manager.MoondreamManager

class ImagePerceptionEngine:
    """
    画像を解析してメタデータを生成するエンジン
    """
    def __init__(self):
        self.vlm = MoondreamManager()
        self.output_dir = "c:/Users/dance/zone/uniai/data/vlm_db"

    def perceive(self, image_path):
        print(f">>> Processing image: {image_path}")
        
        # 1. 画像の説明を生成
        description = self.vlm.analyze_image(image_path, "Describe what is happening in this image.")
        
        # 2. テキスト（OCR）を抽出
        ocr_text = self.vlm.extract_text(image_path)
        
        # 3. メタデータの構築
        metadata = {
            "source_file": image_path,
            "filename": os.path.basename(image_path),
            "timestamp": datetime.now().isoformat(),
            "description": description,
            "ocr_text": ocr_text,
            "tags": self._generate_tags(description, ocr_text)
        }
        
        # 4. JSONとして保存
        self._save_metadata(metadata)
        return metadata

    def _generate_tags(self, description, ocr_text):
        # 簡易的なタグ生成（本来はLLMにやらせるのが良い）
        common_keywords = ["screenshot", "ui", "code", "chat", "windows", "docker", "python"]
        tags = []
        combined = (description + " " + ocr_text).lower()
        for kw in common_keywords:
            if kw in combined:
                tags.append(kw)
        return list(set(tags))

    def _save_metadata(self, metadata):
        if not os.path.exists(self.output_dir):
            os.makedirs(self.output_dir)
            
        target_path = os.path.join(self.output_dir, f"{metadata['filename']}.json")
        with open(target_path, "w", encoding="utf-8") as f:
            json.dump(metadata, f, indent=4, ensure_ascii=False)
        print(f">>> Metadata saved: {target_path}")

if __name__ == "__main__":
    # エンジンのテスト
    # engine = ImagePerceptionEngine()
    pass
