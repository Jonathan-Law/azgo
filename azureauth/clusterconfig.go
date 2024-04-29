package azureauth

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetClusterConfig attempts to automatically fetch the resource group and cluster name from Azure,
// or prompts the user to enter them if unable to automatically determine.
func GetClusterConfig() (resourceGroup, clusterName string, err error) {
	// Attempt to automatically fetch resource group and cluster name
	resourceGroup, clusterName, err = fetchAKSDetails()
	if err == nil {
		return resourceGroup, clusterName, nil
	}

	// If automatic fetch fails, prompt user for details
	fmt.Println("Unable to automatically determine AKS details.")
	return promptForAKSDetails()
}

// fetchAKSDetails tries to find AKS cluster details using Azure CLI
func fetchAKSDetails() (resourceGroup, clusterName string, err error) {
	fmt.Println("Fetching AKS cluster details from Azure...")
	cmd := exec.Command("az", "aks", "list", "--query", "[0].{Name:name,ResourceGroup:resourceGroup}", "--output", "tsv")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch AKS details: %v", err)
	}

	details := strings.Split(strings.TrimSpace(string(output)), "\t")
	if len(details) < 2 {
		return "", "", fmt.Errorf("no AKS clusters found or output format unexpected")
	}

	return details[1], details[0], nil
}

// promptForAKSDetails requests the user to manually input the resource group and cluster name
func promptForAKSDetails() (resourceGroup, clusterName string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the Resource Group Name: ")
	scanner.Scan()
	resourceGroup = scanner.Text()

	fmt.Print("Enter the AKS Cluster Name: ")
	scanner.Scan()
	clusterName = scanner.Text()

	if resourceGroup == "" || clusterName == "" {
		return "", "", fmt.Errorf("resource group and cluster name must not be empty")
	}

	return resourceGroup, clusterName, nil
}
