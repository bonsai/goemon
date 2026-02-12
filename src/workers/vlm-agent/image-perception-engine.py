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

    def perceive(self, file_path):
        ext = os.path.splitext(file_path)[1].lower()
        if ext == ".pdf":
            return self.perceive_pdf(file_path)
        else:
            return self.perceive_image(file_path)

    def perceive_image(self, image_path):
        print(f">>> Processing image: {image_path}")
        # 1. 画像の説明を生成
        description = self.vlm.analyze_image(image_path, "Describe what is happening in this image.")
        # 2. テキスト（OCR）を抽出
        ocr_text = self.vlm.extract_text(image_path)
        return self._create_and_save_metadata(image_path, description, ocr_text)

    def perceive_pdf(self, pdf_path):
        print(f">>> Processing PDF: {pdf_path}")
        # PDFからテキストを直接抽出
        ocr_text = self.vlm.extract_text_from_pdf(pdf_path)
        # 最初の100文字程度を説明の代わりに使う（将来的に要約LLMを入れると良い）
        description = f"PDF Document: {ocr_text[:200]}..."
        return self._create_and_save_metadata(pdf_path, description, ocr_text)

    def _create_and_save_metadata(self, source_file, description, ocr_text):
        metadata = {
            "source_file": source_file,
            "filename": os.path.basename(source_file),
            "timestamp": datetime.now().isoformat(),
            "description": description,
            "ocr_text": ocr_text,
            "tags": self._generate_tags(description, ocr_text)
        }
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
