import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
from PIL import Image
import os
import fitz  # PyMuPDF

class MoondreamManager:
    """
    moondream モデルのロードと推論を管理するクラス
    """
    def __init__(self, model_id="vikhyatk/moondream2", revision="2024-03-06"):
        self.model_id = model_id
        self.device = "cuda" if torch.cuda.is_available() else "cpu"
        self.model = None
        self.tokenizer = None
        self.revision = revision

    def load_model(self):
        print(f">>> Loading Moondream model ({self.model_id})...")
        self.tokenizer = AutoTokenizer.from_pretrained(self.model_id, revision=self.revision)
        self.model = AutoModelForCausalLM.from_pretrained(
            self.model_id, 
            trust_remote_code=True, 
            revision=self.revision
        ).to(self.device)
        self.model.eval()
        print(">>> Model loaded successfully.")

    def analyze_image(self, image_path, prompt="Describe this image in detail."):
        if self.model is None:
            self.load_model()
            
        image = Image.open(image_path).convert("RGB")
        enc_image = self.model.encode_image(image)
        answer = self.model.answer_question(enc_image, prompt, self.tokenizer)
        return answer

    def extract_text(self, image_path):
        """
        画像からテキストを抽出（OCR的な振る舞い）
        """
        prompt = "Read all the text visible in this image and list it."
        return self.analyze_image(image_path, prompt)

    def extract_text_from_pdf(self, pdf_path):
        """
        PDFからテキストを抽出する
        """
        print(f">>> Extracting text from PDF: {pdf_path}")
        text = ""
        try:
            doc = fitz.open(pdf_path)
            for page in doc:
                text += page.get_text()
            doc.close()
        except Exception as e:
            print(f"[!] Error reading PDF: {e}")
        return text

if __name__ == "__main__":
    # テスト実行
    manager = MoondreamManager()
    # manager.load_model() # 遅延ロードのテスト
