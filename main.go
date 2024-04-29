package main

import (
	"azgo/azureauth"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "azgo"}
	rootCmd.AddCommand(loginCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func loginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Log in to Azure and set up resources",
		Long: `This command handles the authentication with Azure, allows the user to select a subscription, 
and connects to the Azure Container Registry (ACR) associated with the chosen subscription.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Azure login process...")
			fmt.Println("Logging into Azure...")
			err := azureauth.Authenticate()
			if err != nil {
				if err.Error() == "No subscription found in the CLI profile" {
					fmt.Println("Please run 'az login' to log in to Azure first.")
					os.Exit(1)
				} else if err.Error() == "Invoking Azure CLI failed with the following error: exec: \"az\": executable file not found in $PATH" {
					fmt.Printf("Login failed: %v\n", err)
					fmt.Println("Please ensure the Azure CLI is installed and configured correctly.")
					fmt.Println("Visit the following link for more information on installing and configuring the Azure CLI: https://learn.microsoft.com/en-us/cli/azure/install-azure-cli")
					os.Exit(1)
				} else {
					fmt.Println("Please check your Azure credentials and try again.")
					os.Exit(1)
				}
				return
			}

			/* If I need to gather a selected subscription -- this is how */
			_, err = azureauth.SelectSubscription()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to select subscription: %v\n", err)
				return
			}

			acrs, err := azureauth.ListACRs()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to list ACRs: %v\n", err)
				return
			}

			acrName, err := azureauth.ChooseACR(acrs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to select ACR: %v\n", err)
				return
			}

			if err := azureauth.ConnectToACR(acrName); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to connect to ACR: %v\n", err)
				return
			}

			resourceGroup, clusterName, err := azureauth.GetClusterConfig()
			if err := azureauth.UpdateKubeConfig(resourceGroup, clusterName); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to update kubectl config: %v\n", err)
				return
			}

			fmt.Println("Setup complete.")
		},
	}
}
