# API Documentation

## Base URL
- Local: `http://localhost:8080/api`

## Endpoints

### GET `/health`
Returns service status and runtime metadata.

Example response:
```json
{
  "status": "healthy",
  "service": "datacenter-guard-api",
  "version": "1.0.0"
}
```

### GET `/tools`
Query parameters:
- `category`
- `pricing`
- `search`

Example response:
```json
{
  "success": true,
  "count": 53,
  "data": [{ "id": "hashicorp-vault", "name": "HashiCorp Vault" }]
}
```

### GET `/tools/:id`
Returns a single tool by ID.

### GET `/frameworks`
Returns all supported compliance frameworks.

### GET `/frameworks/:id/checklist`
Returns checklist categories and normalized checklist items.

Example response:
```json
{
  "success": true,
  "data": [
    {
      "name": "Identify",
      "items": [
        { "id": "nist-csf-0-0", "title": "Maintain a current inventory...", "checked": false }
      ]
    }
  ]
}
```

### GET `/configs/templates`
Returns all stored template definitions.

### POST `/configs/generate`
Request body:
```json
{
  "type": "terraform",
  "template_id": "terraform-aws-vpc",
  "parameters": {
    "management_cidr": "10.0.0.0/24"
  }
}
```

Example response:
```json
{
  "success": true,
  "data": {
    "content": "terraform {...}",
    "filename": "terraform-aws-vpc.tf"
  }
}
```

### GET `/playbooks`
Returns incident response playbooks.

### GET `/architectures`
Returns reference architecture templates.

### POST `/export`
Placeholder endpoint for future export workflows.
