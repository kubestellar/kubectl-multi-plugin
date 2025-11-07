# kubectl-multi

A comprehensive kubectl plugin for multi-cluster operations with KubeStellar. This plugin extends kubectl to work seamlessly across all KubeStellar managed clusters, providing unified views and operations while filtering out workflow staging clusters (WDS).

## Overview

kubectl-multi is a kubectl plugin written in Go that automatically discovers KubeStellar managed clusters and executes kubectl commands across all of them simultaneously. It provides a unified tabular output with cluster context information, making it easy to monitor and manage resources across multiple clusters.

### Key Features

- **Multi-cluster resource viewing**: Get resources from all managed clusters with unified output
- **Cluster context identification**: Each resource shows which cluster it belongs to
- **All kubectl commands**: Supports all major kubectl commands across clusters
- **KubeStellar integration**: Automatically discovers managed clusters via KubeStellar APIs
- **WDS filtering**: Automatically excludes Workload Description Space clusters
- **Familiar syntax**: Uses the same command structure as kubectl


## how to install 

### Downloading step for Linux
```bash
# Step 1: Download kubectl-multi binary for Linux
TAG="v0.0.3"

# Fix: Use ${TAG#v} to remove just 'v', not 'v_'
curl -L -o "kubectl-multi_${TAG#v}_linux_amd64.tar.gz" \
  "https://github.com/kubestellar/kubectl-plugin/releases/download/${TAG}/kubectl-multi_${TAG#v}_linux_amd64.tar.gz"

# Step 2: Extract and install
tar -xzf "kubectl-multi_${TAG#v}_linux_amd64.tar.gz"
sudo mv kubectl-multi /usr/local/bin/kubectl-multi



#to test
kubectl-multi -v

```

### Downloading step for Windows

#### Using PowerShell (Recommended)
```powershell
# Step 1: Set version and download kubectl-multi binary for Windows
$TAG = "v0.0.3"
$VERSION = $TAG.TrimStart('v')

# Download the binary
$url = "https://github.com/kubestellar/kubectl-plugin/releases/download/$TAG/kubectl-multi_${VERSION}_windows_amd64.tar.gz"
Invoke-WebRequest -Uri $url -OutFile "kubectl-multi_${VERSION}_windows_amd64.tar.gz"

# Step 2: Extract the archive
tar -xzf "kubectl-multi_${VERSION}_windows_amd64.tar.gz"

# Step 3: Move to a directory in your PATH (e.g., C:\Program Files\kubectl-multi)
# Create directory if it doesn't exist
New-Item -ItemType Directory -Force -Path "C:\Program Files\kubectl-multi"
Move-Item -Force kubectl-multi.exe "C:\Program Files\kubectl-multi\kubectl-multi.exe"

# Step 4: Add to PATH if not already there
$path = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($path -notlike "*C:\Program Files\kubectl-multi*") {
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$path;C:\Program Files\kubectl-multi",
        "Machine"
    )
    Write-Host "Added kubectl-multi to PATH. Please restart your terminal."
}

# Test installation (restart terminal first)
kubectl-multi -v
```

#### Manual Installation (Alternative)
1. Download the Windows binary from the [releases page](https://github.com/kubestellar/kubectl-plugin/releases/latest)
   - Look for `kubectl-multi_*_windows_amd64.tar.gz`
2. Extract the archive using 7-Zip, WinRAR, or Windows built-in tar
3. Move `kubectl-multi.exe` to a directory in your PATH:
   - Example: `C:\Program Files\kubectl-multi\`
   - Or add it to an existing PATH directory like `C:\Windows\System32\`
4. Add the directory to your PATH environment variable:
   - Search for "Environment Variables" in Windows Settings
   - Edit the "Path" variable under System variables
   - Add the directory containing `kubectl-multi.exe`
5. Restart your terminal and test: `kubectl-multi -v`

### Downloading by brew
```bash 
# This looks for: github.com/kubestellar/homebrew-kubectl-multi 
brew tap kubestellar/kubectl-multi

# This also looks for: github.com/kubestellar/homebrew-kubectl-multi 
brew install kubestellar/kubectl-multi/kubectl-multi

kubectl-multi -v
```


## Quick Start for developer

```bash
# Install the plugin
make install

# To build the binary
make build

# Get nodes from all managed clusters
kubectl multi get nodes

# Get pods from all clusters in all namespaces
kubectl multi get pods -A
```


## Documentation

- **[Installation Guide](docs/installation_guide.md)** - How to install and set up kubectl-multi
- **[Usage Guide](docs/usage_guide.md)** - Detailed usage examples and commands
- **[Architecture Guide](docs/architecture_guide.md)** - Technical architecture and how it works
- **[Development Guide](docs/development_guide.md)** - Contributing and development workflow
- **[API Reference](docs/api_reference.md)** - Code organization and technical implementation

## Tech Stack

- **Go 1.21+**: Primary language for the plugin
- **Cobra**: CLI framework for command structure and parsing
- **Kubernetes client-go**: Official Kubernetes Go client library
- **KubeStellar APIs**: For managed cluster discovery

## Example Output

```
CONTEXT  CLUSTER       NAME                    STATUS  ROLES          AGE    VERSION
its1     cluster1      cluster1-control-plane  Ready   control-plane  6d23h  v1.33.1
its1     cluster2      cluster2-control-plane  Ready   control-plane  6d23h  v1.33.1
its1     its1-cluster  kubeflex-control-plane  Ready   <none>         6d23h  v1.27.2+k3s1
```

## Related Projects

- [KubeStellar](https://github.com/kubestellar/kubestellar) - Multi-cluster configuration management
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) - Kubernetes command-line tool

## Support

For issues and questions:
- File an issue in this repository  
- Check the KubeStellar documentation
- Join the KubeStellar community discussions

## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.