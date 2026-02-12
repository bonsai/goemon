import os
from huggingface_hub import snapshot_download

def download_models():
    models = {
        "moondream2": "vikhyatk/moondream2",
        "anything-v5": "stablediffusionapi/anything-v5",
        "glm-4-9b-chat": "THUDM/glm-4-9b-chat" # 例としてGLM-4を指定
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
