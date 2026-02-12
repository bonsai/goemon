# ADR-018: VLM Agent - Perception & Memory Integration (V3 Vision)

## Status
Proposed

## Context
The user has a vast collection of "ideas and memories" in the form of screenshots and photos (e.g., [memo](file:///c:/Users/dance/zone/vlma/docs/memo)). Currently, these are just static files. We need a way to transform these images into a searchable "Latent Space" (Perception) that can be queried via the Mail UI.

## Goals
- **Automatic Perception**: A background agent (VLM Agent) that monitors folders and extracts metadata (OCR, Object Detection, Description).
- **Memory Management**: Store extracted metadata in a vector database for semantic search.
- **Conversational Retrieval**: Use the "Mail" interface for users to ask questions like "Do I have a picture of X?" or "What was the idea in that screenshot about Y?".
- **Idea Recycling**: Turn old screenshots into new prompts or concepts for the V2/V3 generation pipeline.

## Proposed Architecture

### 1. The Watcher (Folder Monitor)
- **Role**: Monitors `data/vlm_watch` for new images.
- **Implementation**: Python script using `watchdog` or periodic polling.

### 2. The Processor (Batch Perception)
- **Role**: Runs a local VLM (e.g., Moondream2, LLaVA-v1.5-7b) to process images in batches.
- **Output**: Generates a JSON metadata file for each image containing:
    - `description`: Natural language summary.
    - `ocr_text`: Extracted text from screenshots.
    - `tags`: Key objects and concepts.
    - `timestamp`: When the memory was captured.

### 3. The Memory DB (Vector Space)
- **Role**: Indexes the metadata.
- **Implementation**: Use a lightweight vector store (e.g., FAISS or ChromaDB) running locally.
- **Integration**: Convert descriptions and OCR text into embeddings (using a small local model like `all-MiniLM-L6-v2`).

### 4. The Mail UI Integration
- **Role**: Acts as the interface for the user.
- **Flow**:
    1. User sends a Mail: "〇〇な写真はあるかい？" (Is there a photo of X?).
    2. VLM Agent queries the Vector DB.
    3. VLM Agent returns a Mail with:
        - A list of matching images.
        - A summary of why they match.
        - Links to the original files.

## Privacy & Locality
- **No Cloud**: All processing happens on the user's machine.
- **Small Models**: Prioritize models that can run on a standard GPU or even CPU (quantized).

## Next Steps
1. Implement the `vlm-processor.py` to extract metadata using a local VLM.
2. Setup the FAISS index for vector search.
3. Integrate the search capability into the `mail-order-agent.py`.
