# Calculator — Full-Stack Application

A full-stack calculator with a **React (TypeScript)** frontend and a **Go** backend REST API.

## Project Structure

```
.
├── backend/               # Go REST API
│   ├── calculator/        # Pure arithmetic functions (business logic)
│   ├── handlers/          # HTTP handlers (transport layer)
│   ├── middleware/         # CORS & request logging
│   ├── main.go            # Entrypoint
│   ├── Dockerfile
│   └── go.mod
├── frontend/              # React TypeScript SPA
│   ├── src/
│   │   ├── components/    # Calculator UI component
│   │   ├── __tests__/     # Unit tests
│   │   ├── api.ts         # API client
│   │   └── App.tsx
│   ├── public/
│   ├── Dockerfile
│   └── package.json
├── docker-compose.yml     # Run both services together
└── README.md
```

---

## Prerequisites

| Tool       | Version  | Purpose         |
|------------|----------|-----------------|
| Go         | 1.22+    | Backend         |
| Node.js    | 18+      | Frontend        |
| Docker     | 24+      | (Optional) Deploy |

---

## Getting Started

### Clone the repository

```bash
git clone https://github.com/Prajwal509/Full-Stack-Calculator-App.git
cd Full-Stack-Calculator-App
```

### 1. Backend

```bash
cd backend
go run .
```

The server starts on **http://localhost:8080**.

### 2. Frontend (in a separate terminal)

```bash
cd frontend
npm install
npm start
```

Opens on **http://localhost:3000**. The CRA proxy forwards `/api` requests to the backend at `:8080`.

---

## Running Tests

### Backend

```bash
cd backend
go test ./... -v
```

**With coverage report:**

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Frontend

```bash
cd frontend
npm test                  # Watch mode
npm run test:ci           # Single run (CI)
```

---

## API Reference

All endpoints are prefixed with `/api`.

### `POST /api/calculate`

Perform a calculation.

**Request body (JSON):**

| Field       | Type     | Required | Description                        |
|-------------|----------|----------|------------------------------------|
| `operation` | `string` | Yes      | One of the operations listed below |
| `a`         | `number` | Yes      | First operand                      |
| `b`         | `number` | Depends  | Second operand (not needed for `sqrt`) |

**Supported operations:**

| Operation       | Description                 | Requires `b` |
|-----------------|-----------------------------|---------------|
| `add`           | a + b                       | Yes           |
| `subtract`      | a − b                       | Yes           |
| `multiply`      | a × b                       | Yes           |
| `divide`        | a ÷ b                       | Yes           |
| `exponentiate`  | a ^ b                       | Yes           |
| `sqrt`          | √a                          | No            |
| `percentage`    | b% of a → (a × b / 100)    | Yes           |

#### Examples

```bash
# Addition
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","a":10,"b":3}'
# → {"result":13}

# Division
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":3}'
# → {"result":3.3333333333333335}

# Division by zero
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":0}'
# → 400 {"error":"division by zero"}

# Square root
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"sqrt","a":16}'
# → {"result":4}

# Percentage (15% of 200)
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"percentage","a":200,"b":15}'
# → {"result":30}
```

### `GET /api/health`

Health check endpoint.

```bash
curl http://localhost:8080/api/health
# → {"status":"ok"}
```

---

## Docker Deployment

### Prerequisites

| Tool            | Version | Purpose                    |
|-----------------|---------|----------------------------|
| Docker          | 24+     | Containerise both services |
| Docker Compose  | v2+     | Orchestrate multi-container setup |

#### macOS

Install [Docker Desktop](https://docs.docker.com/desktop/install/mac-install/) (includes Docker Compose):

```bash
brew install --cask docker
```

After installing, launch **Docker Desktop** from Applications and wait for the engine to start.

#### Windows

Install [Docker Desktop](https://docs.docker.com/desktop/install/windows-install/) (includes Docker Compose):

```powershell
winget install Docker.DockerDesktop
```

After installing, launch **Docker Desktop** and ensure WSL 2 backend is enabled (Settings → General).

#### Linux (Ubuntu/Debian)

```bash
# Remove old versions
sudo apt-get remove docker docker-engine docker.io containerd runc

# Install Docker
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

# (Optional) Run Docker without sudo
sudo usermod -aG docker $USER
newgrp docker
```

#### Verify installation

```bash
docker --version
docker compose version
```

### Start the application

From the project root (where `docker-compose.yml` lives):

```bash
docker compose up --build
```

This will:

1. **Build the backend image** — compiles the Go binary in a multi-stage Alpine build.
2. **Build the frontend image** — runs `npm ci && npm run build`, then serves the production bundle with Nginx.
3. **Start both containers** — the frontend container waits for the backend to be ready (`depends_on`).

Once you see logs from both services, open:

- Frontend: **http://localhost:3000**
- Backend (direct): **http://localhost:8080**

The frontend Nginx reverse-proxies `/api` requests to the backend container, so all API calls go through port 3000.

### Run in the background

```bash
docker compose up --build -d
```

View logs:

```bash
docker compose logs -f            # all services
docker compose logs -f backend    # backend only
docker compose logs -f frontend   # frontend only
```

### Stop the application

```bash
docker compose down               # stop & remove containers
docker compose down --volumes     # also remove any volumes
```

### Rebuild after code changes

```bash
docker compose up --build
```

Docker caches layers, so only changed layers are rebuilt.

### Useful commands

```bash
docker compose ps                 # list running containers
docker compose exec backend sh    # shell into the backend container
docker compose exec frontend sh   # shell into the frontend container
```

### Port conflicts

If ports 3000 or 8080 are already in use, either stop the conflicting process or override the host ports in `docker-compose.yml`:

```yaml
ports:
  - "4000:3000"   # map host 4000 → container 3000
```

---

## Design Decisions

### Architecture

- **Layered backend**: Business logic (`calculator/`) is separated from HTTP transport (`handlers/`). This makes the arithmetic functions trivially testable without HTTP and allows swapping the transport layer (e.g., gRPC) without touching logic.
- **Single endpoint**: All operations go through `POST /api/calculate` with an `operation` field, rather than separate endpoints per operation. This keeps the API surface small and the frontend simple — one fetch function handles everything.
- **Standard library only**: The backend uses only Go's standard library (`net/http`, `encoding/json`, `math`). No external router or framework is needed for this scope, keeping dependencies at zero.

### Frontend

- **BEM-style CSS**: Component styling uses a `calculator__*` naming convention with no CSS-in-JS dependency. Keeps the styling simple and predictable.
- **API abstraction**: All backend calls go through `api.ts`, making it easy to mock in tests and swap the base URL for different environments via `REACT_APP_API_URL`.
- **Input validation**: The frontend validates inputs before sending to the backend, providing instant feedback. The backend also validates independently (defense in depth).

### Testing

- **Backend**: Table-driven tests cover the arithmetic functions. Handler tests use `httptest` for in-process HTTP testing — no real server needed.
- **Frontend**: Component tests use React Testing Library with the API module mocked. API tests mock `fetch` directly. Both layers test happy paths and error cases.

### Error Handling

- Division by zero and negative square roots return descriptive `400` errors.
- Invalid JSON and missing fields return clear error messages.
- The frontend displays API errors inline without crashing.

### Assumptions

- Floating-point arithmetic uses Go's `float64` / JavaScript `number` — standard IEEE 754 precision.
- The `percentage` operation computes "b% of a" (i.e., `a * b / 100`).
- CORS is set to allow all origins (`*`) for local development. In production, this should be restricted.
