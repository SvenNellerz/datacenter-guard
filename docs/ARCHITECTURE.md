# Architecture

## System Overview
DataCenter Guard is a reference platform that combines a Vue 3 frontend with a Go API. The frontend presents architecture guidance, compliance checklists, templates, tool comparisons, and incident response content. The backend serves the same curated data through REST endpoints for future expansion.

## Components
### Frontend
- Vue 3 single-page application
- Pinia for lightweight client-side state
- Vue Router for section navigation
- Tailwind CSS for styling
- Static JSON data assets for tools, frameworks, templates, and playbooks

### Backend
- Gin-based REST API
- Handlers grouped by resource type
- Service layer that loads JSON assets from the frontend data directory
- CORS middleware for local frontend development

### Data Layer
- Repository-local JSON files in `frontend/src/assets/data/`
- No external database dependency in the bootstrap version
- Static content can later be migrated to object storage or a database-backed CMS

## Data Flow
1. A user opens the Vue SPA.
2. Views either consume local JSON directly or call `/api/*` endpoints.
3. Gin handlers invoke service loaders.
4. Service loaders deserialize JSON assets and return typed data.
5. Responses are serialized as JSON API payloads.

## API Architecture
- `/api/health` for service health
- `/api/tools` and `/api/tools/:id` for tool reference data
- `/api/frameworks` and `/api/frameworks/:id/checklist` for compliance content
- `/api/configs/templates` and `/api/configs/generate` for template browsing and generation
- `/api/playbooks` and `/api/architectures` for IR and design reference data

## Security Considerations
- Restrict CORS origins for deployed environments.
- Treat generated configuration templates as starting points and review before production use.
- Protect privileged guidance, exports, and any future user-generated content with authentication if the platform becomes multi-user.
- Validate and sanitize any future dynamic template inputs before rendering or exporting them.
