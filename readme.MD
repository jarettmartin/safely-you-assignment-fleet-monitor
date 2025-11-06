# Run Instructions

## Prerequisites
- **Go**: Ensure Go is installed. This project was tested with Go **1.25.3**. Download from [https://go.dev/dl/](https://go.dev/dl/).
- **Project Files**: The repository already includes `devices.csv`, `openapi.json`, and `results.txt`.
- **Device Simulator**: Download the appropriate simulator for your OS/architecture from the assignment links. In my case on Windows AMD64, I needed to add a `.exe` extension to allow it to run in the terminal.

## Installation
1. Clone the repository and move into the directory:
```bash
git clone git@github.com:jarettmartin/safely-you-assignment-fleet-monitor.git
cd .\safely-you-assignment-fleet-monitor\
```
2. Install Go dependencies:
```bash
go mod tidy
```

## Running the Project
1. Start the API server from a terminal in the project root:
```bash
go run .
```
2. Run the device simulator from another terminal, providing the port your API is listening on. Example (Windows AMD64, port 6733):
```bash
.\device-simulator-win-amd64.exe 6733
```
