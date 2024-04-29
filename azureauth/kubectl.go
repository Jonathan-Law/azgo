package azureauth

import (
	"fmt"
	"os/exec"
)

// UpdateKubeConfig updates the kubectl configuration to use the credentials of the specified AKS cluster.
func UpdateKubeConfig(resourceGroup, clusterName string) error {
	fmt.Printf("Updating kubectl config for AKS cluster: %s in resource group: %s...\n", clusterName, resourceGroup)
	cmd := exec.Command("az", "aks", "get-credentials", "--resource-group", resourceGroup, "--name", clusterName, "--overwrite-existing")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to update kubectl config: %v\n%s", err, string(output))
	}
	fmt.Println("kubectl config updated successfully.")
	return nil
}
