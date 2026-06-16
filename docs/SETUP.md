# Setup Guide

## Prerequisites
- Node.js 18+ and npm
- Go 1.21+
- Docker and Docker Compose

## Quick Start with Docker Compose
1. Clone the repository.
2. From the repository root run `docker compose up --build`.
3. Open the frontend at `http://localhost:5173`.
4. Access the API health endpoint at `http://localhost:8080/api/health`.

## Manual Frontend Setup
1. Change into `frontend/`.
2. Install dependencies with `npm install`.
3. Start development mode with `npm run dev`.
4. Build production assets with `npm run build`.

## Manual Backend Setup
1. Change into `backend/`.
2. Download modules with `go mod download`.
3. Start the API with `go run .`.
4. Build a binary with `go build -o datacenter-guard-api .`.

## Environment Variables
- `PORT`: backend listen port. Defaults to `8080`.
- `GIN_MODE`: use `debug` locally and `release` in production.

## Development Workflow
- Update static reference data under `frontend/src/assets/data/`.
- Keep frontend views aligned with API response structures.
- Use infrastructure templates as reference assets, not direct production deployment artifacts.
- Run frontend build and backend build/test steps before opening a pull request.
