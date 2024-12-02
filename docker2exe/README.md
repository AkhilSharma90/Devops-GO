# image2exe

`image2exe` is a command-line tool designed to convert Docker images into standalone executable files. This tool simplifies the process of packaging and distributing Docker-based applications as single binary files, allowing users to run containerized applications without needing to manage Docker commands.

## Features

- Converts any Docker image into a standalone executable
- Supports cross-platform (Windows, Linux, macOS)
- Custom working directory support
- Environment variable passing
- Volume mounting capabilities
- Optional image embedding for complete self-containment

## Prerequisites

- Go 1.16 or later
- Docker (installed and running)
- Make (for building from source)

## Installation

### From Source

1. Clone the repository:

   ```bash
   git clone
   cd image2exe
   ```

2. Build the project:

   ```bash
   go build -o image2exe
   ```

## Usage

### Basic Usage

Convert a Docker image to an executable:

```bash
./image2exe --name <executable-name> --image <docker-image>
```

**Example:**

```bash
./image2exe --name myapp --image nginx
```

### Command-line Options

- `--name` (Required): Name of the output executable
- `--image` (Required): Docker image to convert
- `--output`: Output directory (default: `./dist`)
- `--module`: Go module name for the generated code
- `--workdir`, `-w`: Working directory inside the container
- `--env`, `-e`: Environment variables to pass
- `--volume`, `-v`: Volumes to mount
- `--target`, `-t`: Target platforms (default: all supported platforms)
- `--embed`: Embed Docker image in the binary

### Examples

Basic conversion with Alpine:

```bash
./image2exe --name alpine-test --image alpine:latest
```

Specify working directory and environment variables:

```bash
./image2exe --name myapp \
  --image nginx \
  --workdir /app \
  --env "ENV=production" \
  --env "DEBUG=true"
```

Mount volumes:

```bash
./image2exe --name dataapp \
  --image postgres \
  --volume "/host/data:/container/data"
```

Build for a specific platform:

```bash
./image2exe --name winapp \
  --image nginx \
  --target windows/amd64
```

Embed image in executable:

```bash
./image2exe --name standalone \
  --image alpine:latest \
  --embed
```

## Running Generated Executables

After generation, executables will be located in the `dist` directory. You can run them like any normal executable:

**Windows:**

```powershell
.\dist\myapp-windows-amd64.exe [commands...]
```

**Linux/MacOS:**

```bash
./dist/myapp-linux-amd64 [commands...]
```

**Example commands:**

```bash
# Run a simple echo command
./myapp-linux-amd64 echo "Hello World"

# Start an interactive shell
./myapp-linux-amd64 sh

# Run a complex command
./myapp-linux-amd64 sh -c "ls -la && echo 'Done!'"
```

## Supported Platforms

By default, `image2exe` generates executables for:

- Windows (amd64)
- Linux (amd64)
- macOS (amd64, arm64)

You can target specific platforms using the `--target` flag.

## Building from Source

Install dependencies:

```bash
go mod tidy
```

Build the project:

```bash
go build -o image2exe
```

## Troubleshooting

### Docker not running

- Ensure Docker Desktop (Windows/macOS) or Docker daemon (Linux) is running
- Verify with `docker ps`

### Permission Issues

- Ensure you have the proper Docker permissions
- On Linux, ensure your user is part of the Docker group

### Image Pull Failures

- Check your internet connectivity
- Verify the image name is correct
- Ensure you're logged in if using private repositories

### Execution Issues

- Verify the executable matches your platform
- Check if Docker is properly installed and running
- Run with elevated privileges if necessary