# UniAI: Universal AI Agent

Unified Multi-Modal Agent framework integrating TTS, SD, ML, and VLM.

## Vision
A "General Purpose AI Concierge" that operates entirely locally, utilizing a swarm of specialized small models to fulfill complex creative requests via a Mail-driven architecture.

## Roadmap
- **V1**: Foundation (CLI, Docker, Simple UI) - *Completed in snsw/sdml*
- **V2**: Integration (TTS + SD + ML) - *In Progress*
- **V3**: Perception (VLM Integration) - *Planned*

## Directory Structure
- `docs/`: Architecture Decision Records (ADR) and guides.
- `src/core/`: Universal engine logic (Mail, Docker, UI components).
- `src/workers/`: Specialized AI task workers (TTS, ImageGen, etc.).
- `docker/`: Dockerfiles and orchestration scripts.
- `models/`: Local model storage.
- `data/`: Task queues and delivered artifacts.

## Architecture
See [ADR-016](docs/adr-016-v2-ux-agent-worker.md) and [ADR-017](docs/adr-017-universal-agent-v2-v3-integration.md) for details.
