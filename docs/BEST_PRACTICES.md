# Contribution and Best Practices

## Contributing
- Keep changes focused and scoped to the feature or dataset you are updating.
- Preserve JSON validity and consistent formatting in reference data files.
- Update documentation when endpoints, templates, or workflows change.

## Frontend Practices
- Prefer reusable Vue components for repeated UI patterns.
- Keep state in Pinia only when it needs to be shared.
- Use Tailwind utility classes consistently with the existing dark theme.

## Backend Practices
- Keep handlers thin and move data logic into services.
- Return structured success/error payloads using shared response helpers.
- Add validation before introducing new dynamic inputs or export features.

## Security Practices
- Never commit secrets or environment-specific credentials.
- Review template output before using it in production.
- Document assumptions, default ports, and trust boundaries for new assets.

## Review Checklist
- Does the change keep the SPA and API data models aligned?
- Are new JSON files valid and complete?
- Are docs updated for user-visible changes?
- Have build or test steps been considered for the affected area?
