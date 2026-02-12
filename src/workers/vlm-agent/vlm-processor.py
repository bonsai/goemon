import os
import json
import time
from datetime import datetime

# In a real V3 implementation, this would import a local VLM like Moondream or LLaVA
# For the prototype, we simulate the "Perception" process

WATCH_DIR = "c:/Users/dance/zone/uniai/data/vlm_watch"
DB_DIR = "c:/Users/dance/zone/uniai/data/vlm_db"

def process_image(file_path):
    print(f">>> Perceiving: {file_path}")
    
    # Simulate VLM Processing (OCR + Description)
    # In V3: metadata = vlm_model.analyze(file_path)
    file_name = os.path.basename(file_path)
    
    metadata = {
        "file_path": file_path,
        "processed_at": datetime.now().isoformat(),
        "description": f"A memory captured in {file_name}. Contains ideas about system architecture and UI design.",
        "ocr_text": "Sample OCR text from screenshot: 'Set-ItemProperty -Path HKCU...'",
        "tags": ["screenshot", "idea", "windows", "registry"],
        "vector_id": f"vec_{int(time.time())}"
    }
    
    # Save to "Latent Space" (JSON DB for now)
    db_path = os.path.join(DB_DIR, f"{file_name}.json")
    with open(db_path, "w", encoding="utf-8") as f:
        json.dump(metadata, f, indent=4, ensure_ascii=False)
    
    print(f"OK: Memory indexed at {db_path}")

def watch_loop():
    print(f"VLM Agent is watching: {WATCH_DIR}")
    while True:
        files = [f for f in os.listdir(WATCH_DIR) if f.lower().endswith(('.png', '.jpg', '.jpeg'))]
        for f in files:
            full_path = os.path.join(WATCH_DIR, f)
            # Check if already processed
            if not os.path.exists(os.path.join(DB_DIR, f"{f}.json")):
                process_image(full_path)
        
        time.sleep(5)

if __name__ == "__main__":
    if not os.path.exists(DB_DIR): os.makedirs(DB_DIR)
    watch_loop()
