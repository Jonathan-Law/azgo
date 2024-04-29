# AZGO - Azure DevOps Helper CLI

Azure DevOps Helper CLI is a command-line tool designed to simplify and automate tasks related to managing Azure resources. It provides functionalities such as logging into Azure, managing Azure Container Registries (ACR), and configuring Kubernetes credentials through Azure Kubernetes Service (AKS).

## Features

- **Azure Login**: Authenticate and log into your Azure account directly from the command line.
- **Subscription Management**: List and select Azure subscriptions to manage resource scopes effectively.
- **ACR Management**: List and log into Azure Container Registries associated with your Azure subscription.
- **Kubernetes Configuration**: Automatically set up `kubectl` configurations for interacting with AKS.

## Prerequisites

- **Azure CLI**: Azure DevOps Helper CLI depends on the Azure CLI being installed and properly configured on your system. [Install Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)
- **Go (at least version 1.22.2)**: This tool is written in Go. Make sure you have Go installed on your system. [Install Go](https://golang.org/dl/)
- **Git**: To clone the repository into your local machine. [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Jonathan-Law/azgo.git
   cd azgo
   ```

2. **Build the CLI Tool**:
   ```bash
   go build -o azgo main.go
   ```

3. **Optional: Add to Path**:
   Add the tool to your system path to run it from anywhere on your system.
   ```bash
   sudo cp azgo /usr/local/bin/
   ```

## Usage

### Login to Azure
Start by logging into Azure. This will also allow you to select a subscription and manage ACR and AKS:
```bash
./azgo login
```

### Update Kubernetes Configurations (Unimplemented)
After logging in and selecting your subscription and AKS cluster, you can update your Kubernetes configuration:
```bash
./azgo updatekubeconfig
```

## Contributing
Contributions to the Azure DevOps Helper CLI are welcome! Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to make contributions.

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Additional Resources
For more information on Azure CLI and managing Azure resources, visit the [official Azure documentation](https://learn.microsoft.com/en-us/azure/).

