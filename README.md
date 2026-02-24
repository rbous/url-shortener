# url-shortener
> IN-DEVELOPMENT

A minimal, containerized URL shortener and redirect service built with Python and FastAPI. Hosted at [url.mrie.dev](https://url.mrie.dev).

> **Try it:** [url.mrie.dev/prayer](https://url.mrie.dev/prayer)

## Overview

Maps short slugs to full URLs and performs HTTP redirects. Sending a GET request to `/{slug}` redirects the client to the configured destination URL.

## Tech Stack

- **Python 3.9** + **FastAPI** — web framework
- **Uvicorn** — ASGI server
- **Docker** / **Docker Compose** — containerization
- **GitHub Actions** — CI/CD (build, push, and sign image to GHCR)
- **Cosign** (Sigstore) — supply chain security via keyless image signing
- **Cloudflare Tunnel** — exposes the service to the public URL without opening inbound ports

## Getting Started

### Run locally

```bash
pip install fastapi uvicorn
uvicorn logic:app --host 0.0.0.0 --port 8000
```

### Run with Docker

```bash
docker build -t url-shortener .
docker run -p 8000:8000 url-shortener
```

### Run with Docker Compose (uses pre-built image from GHCR)

```bash
docker compose up
```

The service is available at `http://localhost:8000`.

## API

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/{slug}` | Redirects to the mapped URL (`302`), or returns `404` if not found |

Interactive API docs are available at `/docs` (Swagger UI) and `/redoc` (ReDoc).

## Configuration

URL mappings are defined in the `data` dictionary in `logic.py`:

```python
data = {
    "prayer": "https://mrie.dev/prayertimes/default",
}
```

Add or remove entries there to manage redirects.

## CI/CD

The GitHub Actions workflow (`.github/workflows/docker-publish.yml`) triggers on:
- Push to `main`
- Semver tags (`v*.*.*`)
- Pull requests to `main`
- Daily scheduled run

On push to `main` or a version tag, the workflow builds and pushes the image to `ghcr.io/rbous/url-shortener` and signs it with Cosign keyless signing.

## Current Limitations

- URL mappings are hardcoded (no database or persistence layer)
- No API endpoints for creating or deleting mappings
- No authentication

