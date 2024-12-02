## Kubernetes Ping Operator

The Kubernetes Ping Operator is a custom operator designed to automate ping tests to specified hostnames by creating and managing Kubernetes jobs. The operator leverages custom resources to enable users to define the parameters of their ping tests, such as the target hostname and the number of attempts.

### Overview

This operator provides a streamlined way to:
- Define ping tests using Kubernetes custom resources.
- Specify the target hostname for the ping test.
- Set the number of attempts for each ping test.
- Automate the execution of ping tests through Kubernetes jobs.

### Prerequisites

Before using the operator, make sure you have the following software installed on your system. These instructions cater primarily to Windows users utilizing WSL2.

#### For Windows Users:

1. **WSL2 (Windows Subsystem for Linux)**
   - Open PowerShell as Administrator and run:
     ```powershell
     wsl --install
     ```

2. **Docker Desktop**
   - Download and install Docker Desktop.
   - Enable WSL2 integration under Docker Desktop settings.

3. **Go Programming Language**
   - In WSL2, run:
     ```bash
     sudo apt update
     sudo apt install golang-go
     ```

4. **Kubernetes Tools**
   - **kubectl**: 
     ```bash
     curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
     chmod +x kubectl
     sudo mv kubectl /usr/local/bin/
     ```
   - **Minikube**:
     ```bash
     curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
     chmod +x minikube-linux-amd64
     sudo mv minikube-linux-amd64 /usr/local/bin/minikube
     ```

5. **Operator SDK**
   - Download the binary:
     ```bash
     curl -LO https://github.com/operator-framework/operator-sdk/releases/download/v1.31.0/operator-sdk_linux_amd64
     chmod +x operator-sdk_linux_amd64
     sudo mv operator-sdk_linux_amd64 /usr/local/bin/operator-sdk
     ```

### Installation and Setup

1. **Clone the Repository**
   ```bash
   git clone
   cd ping-operator
   ```

2. **Start Minikube**
   ```bash
   minikube start --driver=docker --cpus=2 --memory=2048
   ```

3. **Build and Run the Operator**
   ```bash
   # Install dependencies
   go mod tidy

   # Generate Kubernetes manifests
   make manifests

   # Install CRDs
   make install

   # Run the operator
   make run
   ```

### Usage

#### Create a Basic Ping Test

To create a basic ping test, open a new terminal and apply the following resource:

```bash
kubectl apply -f - <<EOF
apiVersion: monitors.demo.io/v1beta1
kind: Ping
metadata:
  name: ping-sample
spec:
  hostname: "www.google.com"
  attempts: 1
EOF
```

#### Check the Results

Verify the ping test by observing the Kubernetes resources created:

```bash
# View the Ping resource
kubectl get ping

# Check the created job
kubectl get jobs

# View the pod running the ping
kubectl get pods

# Check the ping results
kubectl logs <pod-name>
```

### Customization Examples

#### Ping a Different Website with More Attempts

```bash
kubectl apply -f - <<EOF
apiVersion: monitors.demo.io/v1beta1
kind: Ping
metadata:
  name: github-test
spec:
  hostname: "github.com"
  attempts: 5
EOF
```

#### Test Multiple Websites

```bash
kubectl apply -f - <<EOF
apiVersion: monitors.demo.io/v1beta1
kind: Ping
metadata:
  name: microsoft-test
spec:
  hostname: "microsoft.com"
  attempts: 3
EOF
```

### Troubleshooting

#### Docker Issues

- Ensure Docker Desktop is running.
- Verify WSL2 integration is enabled.
- To test Docker functionality within WSL2:
  ```bash
  docker ps
  ```

#### Minikube Issues

- Check the status of Minikube:
  ```bash
  minikube status
  ```
- If not running, start Minikube:
  ```bash
  minikube start --driver=docker
  ```

#### Common Errors

- If encountering Docker permission errors:
  ```bash
  sudo usermod -aG docker $USER
  # Log out and back into WSL
  ```

- If CRDs are not found, ensure they are installed:
  ```bash
  make install
  ```

### Project Structure

```
ping-operator/
├── api/v1beta1/              # API definitions
│   ├── ping_types.go         # Ping resource definition
├── controllers/              # Controller logic
│   ├── ping_controller.go    # Main controller code
├── config/                   # Kubernetes manifests
└── Makefile                  # Build and deployment commands
```

### How It Works

1. The operator monitors for newly created Ping resources.
2. Upon detecting a new Ping resource, it creates a corresponding Kubernetes job.
3. The job runs a pod executing the ping command with the specified parameters.
4. To view the results, check the logs of the created pod.