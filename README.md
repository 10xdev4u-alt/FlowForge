# FlowForge

A high-performance, scalable, AI-powered workflow builder (mini-Zapier clone).

## Features
- **Visual DAG Builder:** React Flow-powered canvas for designing complex automation.
- **Distributed Engine:** Powered by Asynq and Redis for reliable, scalable execution.
- **AI Integration:** OpenAI-powered nodes for intelligent data processing.
- **Type-Safe Persistence:** PostgreSQL with sqlc for high-performance data access.
- **Stateless Auth:** JWT-based authentication with bcrypt security.

## Tech Stack
- **Backend:** Golang (net/http + go-chi), PostgreSQL (via sqlc), Redis, Asynq.
- **Frontend:** React (Vite), TypeScript, TailwindCSS, Zustand, React Flow.
- **Infrastructure:** Docker, Docker Compose, Makefiles.

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Go 1.24+
- Node.js 18+

### Setup
1. Clone the repository.
2. Spin up the infrastructure:
   ```bash
   docker-compose up -d
   ```
3. Initialize the database (if using golang-migrate):
   ```bash
   make migrate-up
   ```
4. Start the backend:
   ```bash
   go run main.go
   ```
5. Start the frontend:
   ```bash
   cd client
   npm install
   npm run dev
   ```

## Architecture
FlowForge uses a modern micro-services architecture where the HTTP API serves as a thin layer that enqueues tasks into a distributed Redis-backed queue (Asynq). Multiple worker instances can then process these tasks concurrently, enabling massive scale.
