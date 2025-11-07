# Run Instructions

## Prerequisites
- **Go**: Ensure Go is installed. This project was tested with Go **1.25.3**. Download from [https://go.dev/dl/](https://go.dev/dl/).
- **Docker**: Install Docker Desktop or Docker Engine from [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/).
- **Project Files**: The repository already includes `devices.csv`, `openapi.json`, and `results.txt`.
- **Device Simulator**: Download the appropriate simulator for your OS/architecture from the assignment links. In my case on Windows AMD64, I needed to add a `.exe` extension to allow it to run in the terminal.

---

## Environment Setup
Before running the API (locally or via Docker), you must create a `.env` file in the project root.  
This configuration file defines mandatory environment variables for the API server.

1. Copy the example file:
```bash
cp .env.example .env
```
2. Open `.env` and update as needed. For example:
```
API_PORT=6733
API_VERSION=v1
```

---

## Running Locally (without Docker)
1. Install Go dependencies:
```bash
go mod tidy
```
2. Start the API server:
```bash
go run .
```
3. In a second terminal, run the device simulator, providing the API port defined in .env:
```bash
./device-simulator-win-amd64.exe 6733
```

---

## Running with Docker
1. Build and start the service using Docker Compose:
```bash
docker compose up --build
```
This command automatically loads the `.env` file to set the API port and version.

2. The server will listen on the port defined in `.env`. For example:
```
Server listening on :6733 (version: v1)
```

3. Run the simulator against the same port:
```bash
./device-simulator-win-amd64.exe 6733
```

4. To stop and remove containers:
```bash
docker compose down
```

---

## Running Unit Tests
**Additional learning experience** – For curiosity and exploration, a few simple unit tests were implemented to understand Go’s testing workflow.

1. From the project root, run all unit tests:
```bash
go test ./...
```
