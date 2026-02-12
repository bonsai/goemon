FROM python:3.10-slim

# Set Python to unbuffered mode to see logs immediately
ENV PYTHONUNBUFFERED=1

WORKDIR /app

# Install system dependencies in a single layer
RUN apt-get update && apt-get install -y \
    git \
    libgl1 \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

# Install torch separately with pip options to handle large downloads
RUN pip install --no-cache-dir \
    --timeout 1000 \
    --retries 5 \
    torch

# Install other ML packages
RUN pip install --no-cache-dir \
    accelerate \
    diffusers \
    pillow \
    safetensors \
    transformers

# Additional dependencies for VLM and other features
RUN pip install --no-cache-dir \
    einops \
    timm \
    sentencepiece

# Install huggingface_hub for model downloading
RUN pip install --no-cache-dir huggingface_hub

# Copy model downloader
COPY download_models.py .

# Create models directory
RUN mkdir -p /app/models

# Download models during build (Optional: can be moved to entrypoint for runtime download)
# RUN python download_models.py

# Copy source code from new structure
# Note: Build context is project root
COPY src/workers/Agent/main.py \
     src/workers/Agent/manager.py \
     ./
COPY src/workers/Image_Baker/baker.py \
     src/workers/Image_Baker/inspector.py \
     ./
COPY src/workers/Mail_Manager/postman.py ./

# Run the script
CMD ["python", "main.py"]
