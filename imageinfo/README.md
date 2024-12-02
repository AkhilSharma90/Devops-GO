# Docker Image Analysis Tool

this is a powerful utility for analyzing Docker images, specifically designed to help developers and security professionals identify potential security issues and understand the detailed structure of their Docker images.

## Features

- **Image Layer and Structure Analysis**: Examine the composition of Docker image layers.
- **Secrets Detection**: Identify potential secrets and sensitive files within images.
- **Environment Variable Display**: Expose all environment variables set in the image.
- **Port Exposure Display**: List all ports exposed by the image.
- **User Permissions Identification**: Show which user the image runs as and warn if it's root.
- **Dockerfile Reconstruction**: Recreate Dockerfile instructions from the image.
- **Layer Extraction**: Extract and save image layers for in-depth inspection.
- **Multiple Image Support**: Analyze more than one image at a time.
- **Verbose Mode**: Get detailed output for thorough analysis.

## Prerequisites

- **Go**: Version 1.19 or higher.
- **Docker**: Installed and running on your system.
- **Docker API Access**: Required for image analysis.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone
   cd
   ```

2. **Build the Project**:
   ```bash
   go build .
   ```

## Usage

### Basic Analysis

To analyze a single Docker image:
```bash
./executable nginx:latest
```

### Verbose Mode

For more detailed information about the image:
```bash
./executable -v nginx:latest
```

### Extract Layers

To save image layers to the current directory for inspection:
```bash
./executable -x nginx:latest
```

### Analyze Multiple Images

1. Create a text file (`images.txt`) with image names:
   ```bash
   echo "nginx:latest" > images.txt
   echo "ubuntu:20.04" >> images.txt
   ```

2. Run the analysis:
   ```bash
   ./executable -f images.txt
   ```

### Use Specific Docker API Version

If a specific Docker API version is needed:
```bash
./executable -sV=1.41 nginx:latest
```

## Command Line Options

- **-v**: Enable verbose output.
- **-f**: Specify a file containing images to analyze.
- **-x**: Extract layers to the current directory.
- **-filter**: Filter out noise (enabled by default).
- **-sV**: Set a specific Docker API version.

## Output Explanation

provides information in several key sections:

- **Basic Information**: Displays Docker version and GraphDriver details.
- **Environment Variables**: Lists all environment variables set in the image.
- **Open Ports**: Shows all exposed ports.
- **User Information**: Indicates the user the image runs as and warns if it's root.
- **Potential Secrets**: Lists files that might contain sensitive information and their paths.
- **Dockerfile Reconstruction**: Displays the commands used to build the image and layer info.

### Example Output

```bash
Analyzing nginx:latest
Docker Version: 20.10.12
GraphDriver: overlay2
Environment Variables:
| PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
| NGINX_VERSION=1.21.6
Open Ports:
| 80/tcp
Image user:
| User is root
Potential secrets:
| Found match etc/ssl/private/ssl-cert-snakeoil.key (Private key file)
```

## Security Considerations

- Requires access to the Docker daemon.
- Running as root might be required for full functionality.
- Be cautious with extracted layers as they might contain sensitive data.

## Troubleshooting

### Docker Connection Issues

Check if Docker is running:
```bash
docker ps
```

Try running with root privileges:
```bash
sudo ./executable nginx:latest
```

### API Version Mismatch

Use a specific API version if needed:
```bash
./executable -sV=1.41 nginx:latest
```

### Image Pull Issues

Manually pull the image first:
```bash
docker pull nginx:latest
```