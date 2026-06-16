# DataCenter-Guard 🛡️

A comprehensive, interactive reference guide and design tool for secure greenfield data centre implementation.

## Overview

DataCenter-Guard is an educational platform designed to help organizations understand and implement secure data centre design patterns. It covers network architecture, security layers, privileged access management, logging & monitoring, compliance frameworks, and incident response.

### Key Features
- 🏗️ **Interactive Architecture Builder** - Drag-and-drop network topology designer
- 🔐 **PAM & Security Guides** - Best practices for privileged access management
- 📊 **SIEM & Monitoring** - Comprehensive logging and monitoring setup guides
- ✅ **Compliance Frameworks** - NIST, CIS, ISO 27001, HIPAA, PCI-DSS checklists
- 🛠️ **Tool Reference Database** - 50+ security tools with comparisons
- 📋 **IaC Templates** - Terraform, Ansible, and CloudFormation examples
- 📈 **Trade-off Calculator** - Cost vs Security analysis
- 🚨 **Incident Response** - IR playbook templates

## Tech Stack

- **Frontend**: Vue 3 + Vite, Pinia, Tailwind CSS, D3.js
- **Backend**: Go + Gin Framework
- **Deployment**: Docker, Netlify (frontend), DigitalOcean (full-stack)

## Quick Start

### Local Development

```bash
# Clone the repository
git clone https://github.com/SvenNellerz/datacenter-guard.git
cd datacenter-guard

# Using Docker Compose (recommended)
docker-compose up

# Frontend: http://localhost:5173
# Backend API: http://localhost:8080
```

### Manual Setup

**Frontend**:
```bash
cd frontend
npm install
npm run dev
```

**Backend**:
```bash
cd backend
go mod download
go run main.go
```

## Project Structure

```
datacenter-guard/
├── frontend/              # Vue 3 + Vite application
├── backend/              # Go REST API
├── docs/                 # Documentation
└── docker-compose.yml    # Development setup
```

## Documentation

- [SETUP.md](docs/SETUP.md) - Detailed setup instructions
- [ARCHITECTURE.md](docs/ARCHITECTURE.md) - System architecture overview
- [API.md](docs/API.md) - API documentation
- [DEPLOYMENT.md](docs/DEPLOYMENT.md) - Deployment guides

## Features

### 1. Dashboard
Central hub with quick navigation and key statistics

### 2. Network Architecture
- Visual topology editor
- DMZ configuration
- Network segmentation patterns
- Template library

### 3. PAM (Privileged Access Management)
- Tool comparisons (Vault, CyberArk, Okta)
- Workflow diagrams
- Setup guides

### 4. Monitoring & Logging
- SIEM solutions comparison
- Alert templates
- Centralized logging setup

### 5. Compliance
- Framework-based checklists
- Progress tracking
- Alignment mapping (NIST, CIS, ISO 27001, HIPAA, PCI-DSS)

### 6. Tools Reference
- Searchable database (50+ tools)
- Cost/complexity analysis
- Feature comparison matrix

### 7. Configuration Templates
- Terraform modules (AWS/Azure/GCP)
- Ansible playbooks
- Firewall rules
- SIEM configurations

### 8. Incident Response
- IR playbook templates
- Breach procedures
- Evidence collection guides

## API Endpoints

```
GET  /api/health              - Health check
GET  /api/tools               - List all tools
GET  /api/tools/:id           - Get tool details
GET  /api/frameworks          - List compliance frameworks
GET  /api/frameworks/:id      - Get framework details
POST /api/frameworks/:id/checklist - Generate checklist
GET  /api/templates           - Get configuration templates
POST /api/configs/generate    - Generate Terraform/Ansible
GET  /api/playbooks           - Get incident response playbooks
```

## Deployment

### Netlify (Frontend Only)
```bash
npm run build
# Deploy dist/ directory to Netlify
```

### DigitalOcean (Full Stack)
See [DEPLOYMENT.md](docs/DEPLOYMENT.md) for detailed instructions.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Apache License 2.0

## Support

For issues, questions, or suggestions, please open an issue on GitHub.

---

**Built with ❤️ for secure data centre design**
