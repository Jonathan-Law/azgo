package azureauth

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ListACRs lists all available Azure Container Registries within the selected subscription.
func ListACRs() ([]string, error) {
	cmd := exec.Command("az", "acr", "list", "--query", "[].{Name:name}", "--output", "tsv")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list ACRs: %v\n%s", err, string(output))
	}

	// Parse output to get ACR names
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no ACRs found in the current subscription")
	}

	return lines, nil
}

// ChooseACR prompts the user to select an ACR from the list.
func ChooseACR(acrs []string) (string, error) {
	fmt.Println("Select an Azure Container Registry to log into:")
	for i, acr := range acrs {
		fmt.Printf("%d: %s\n", i+1, acr)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number: ")
	scanner.Scan()
	choice := scanner.Text()
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(acrs) {
		return "", fmt.Errorf("invalid selection")
	}

	return acrs[index-1], nil
}

// ConnectToACR connects to the selected Azure Container Registry.
func ConnectToACR(acrName string) error {
	fmt.Printf("Connecting to ACR: %s...\n", acrName)
	cmd := exec.Command("az", "acr", "login", "--name", acrName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect to ACR: %v\n%s", err, string(output))
	}

	fmt.Println("Connected to ACR successfully.")
	return nil
}
