import torch
import time
import os
from diffusers import StableDiffusionPipeline

class Baker:
    def __init__(self, model_id, device=None):
        self.model_id = model_id
        self.device = device if device else ("cuda" if torch.cuda.is_available() else "cpu")
        self.pipe = None
        # Lazy load model to save memory
        # self.load_model()

    def load_model(self):
        if self.pipe is not None:
            return
            
        print(f"Loading Stable Diffusion model: {self.model_id} on {self.device}...")
        
        # Determine cache directory: C:\Models for Windows Host, 'models' for Docker/Linux
        cache_dir = "models"
        if os.name == 'nt':
            cache_dir = r"C:\Models"
            
        try:
            # Use local cache to persist downloads
            self.pipe = StableDiffusionPipeline.from_pretrained(
                self.model_id,
                torch_dtype=torch.float32,
                cache_dir=cache_dir
            )
            self.pipe.to(self.device)
            if self.device == "cuda":
                self.pipe.enable_attention_slicing()
            print("Model loaded.")
        except Exception as e:
            print(f"Failed to load model: {e}")
            raise

    def unload_model(self):
        if self.pipe is not None:
            print("Unloading Baker model...")
            del self.pipe
            self.pipe = None
            if self.device == "cuda":
                torch.cuda.empty_cache()
            import gc
            gc.collect()
            print("Baker model unloaded.")

    def bake(self, prompt):
        self.load_model()
        print(f"Baking image for prompt: {prompt[:50]}...")
        with torch.no_grad():
            image = self.pipe(prompt).images[0]
        
        timestamp = int(time.time())
        output_dir = "outputs"
        os.makedirs(output_dir, exist_ok=True)
        image_path = os.path.join(output_dir, f"generated_{timestamp}.png")
        image.save(image_path)
        print(f"Image baked at {image_path}")
        return image_path
