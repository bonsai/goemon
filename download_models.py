import os
from huggingface_hub import snapshot_download

def download_models():
    models = {
        "moondream2": "vikhyatk/moondream2",
        # GLM-4-9B (約15GB) や Anything-V5 (約4GB) は巨大です。
        # ディスク制限 (Kaggle: 20GB永続 / Colab: 15GB共有) を考慮し、
        # 必要に応じてコメントを外して使用してください。
        # "anything-v5": "stablediffusionapi/anything-v5",
        # "glm-4-9b-chat": "THUDM/glm-4-9b-chat"
    }
    
    base_path = "./models"
    os.makedirs(base_path, exist_ok=True)
    
    for name, repo_id in models.items():
        print(f"Downloading {name} from {repo_id}...")
        target_dir = os.path.join(base_path, name)
        if not os.path.exists(target_dir):
            snapshot_download(
                repo_id=repo_id,
                local_dir=target_dir,
                local_dir_use_symlinks=False,
                revision="main"
            )
            print(f"Finished downloading {name}")
        else:
            print(f"{name} already exists, skipping.")

if __name__ == "__main__":
    download_models()
