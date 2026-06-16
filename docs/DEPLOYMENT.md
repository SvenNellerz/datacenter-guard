# Deployment Guide

## Frontend on Netlify
1. Connect the repository to Netlify.
2. Set the base directory to `frontend`.
3. Set the build command to `npm ci && npm run build`.
4. Set the publish directory to `frontend/dist`.
5. Configure the backend API URL through reverse proxy or environment-specific rewrites if the API is hosted separately.

## Full Stack on DigitalOcean
1. Provision a Droplet or App Platform service.
2. Install Docker and Docker Compose on the target host.
3. Clone the repository and run `docker compose up --build -d`.
4. Place the frontend behind a reverse proxy such as Nginx or Caddy.
5. Terminate TLS at the proxy and forward `/api` traffic to the Go service.

## Production Considerations
- Set `GIN_MODE=release`.
- Restrict allowed CORS origins.
- Add centralized logging and monitoring.
- Back up any future persistent content store.
- Use CI/CD secrets for deployment tokens and avoid embedding credentials in the repository.
