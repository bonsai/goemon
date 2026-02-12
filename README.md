# Goemon: Local AI Concierge Swarm

Unified Multi-Modal Agent framework integrating TTS, SD, ML, and VLM, orchestrated by Go.

## Vision
"General Purpose AI Concierge" - A locally operated swarm of specialized small models fulfilling complex creative requests via a Mail-driven architecture. Goemon serves as the brain (Orchestrator) and nervous system (Mail/DB) for the underlying AI muscles.

## Architecture
- **Orchestrator (Go)**: Handles high-concurrency tasks like mail monitoring, file system watching, database management, and worker process control.
- **Workers (Python)**: Executes heavy ML inference tasks (Stable Diffusion, Moondream2 VLM, GLM-based prompt generation).
- **Communication**: Asynchronous agent-to-agent communication via a local Mail (IMAP/SMTP) protocol.

## Directory Structure
- `src/cmd/goemon/`: Main entry point and CLI Command Center.
- `src/cmd/cli/`: Interactive UI tools (Dashboard, Mail UI, Metrics).
- `src/core/agent/`: Agent definitions (Foreman, SpellWriter, MailAgent).
- `src/core/db/`: Persistent storage (SQLite for Image, Mail, and Vector data).
- `src/core/bridge/`: Go-Python integration bridge.
- `src/workers/`: Specialized AI task workers.
- `docs/adr/`: Architecture Decision Records.

## Swarm of Agents
- **Foreman ([foreman.go](src/core/agent/foreman.go))**: System orchestrator and background logic controller.
- **SpellWriter ([spell-writer.go](src/core/agent/spell-writer.go))**: Prompt generator that "steals" fragments from memory (`vec.db`) and requests GLM to cast new spells.
- **Image Service ([service.go](src/core/image/service.go))**: Manages T2I (Baker) and I2T (Watcher) worker lifecycle.
- **Mail Agent ([mail-agent.go](src/core/agent/mail-agent.go))**: Handles agent-specific mailbox logic and identity.

## Databases
- `imagestack.db`: Metadata for generated images and their lifecycle.
- `mail.db`: Persistent record of all agent communications.
- `vec.db`: Vector-ready storage for long-term memory fragments.

## Usage
Build the Goemon Command Center:
```bash
go build -o goemon.exe ./src/cmd/goemon
./goemon.exe
```

The CLI menu provides access to:
1.  **System Dashboard**: Real-time monitoring of workers and mail.
2.  **Mail Interface**: Manually send prompts or receive generated artifacts.
3.  **Data Metrics**: Visualize the system's "Taijukei" (Growth metrics).
4.  **Foreman Control**: Start/Stop the background orchestration logic.

## Development
- **Design Decisions**: Refer to [docs/adr/](docs/adr/) for detailed architectural evolution.
- **Python Bridge**: Worker-specific logic often resides in `python/` subdirectories within core packages.
