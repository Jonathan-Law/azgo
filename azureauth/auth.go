package azureauth

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var authorizer autorest.Authorizer

func Authenticate() error {
	var err error
	authorizer, err = auth.NewAuthorizerFromCLI()
	if err != nil {
		// If the SDK authorizer fails, try invoking `az login`
		fmt.Println("You are not currently authenticated. Attempting to invoke 'az login'...")
		cmd := exec.Command("az", "login")
		output, err := cmd.CombinedOutput()
		if err != nil {
			// If `az login` fails, return the error along with its output
			return fmt.Errorf("failed to invoke 'az login': %s\n%s", err, string(output))
		}

		// Check the output for a successful login message or error handling
		if strings.Contains(string(output), "You have logged in.") {
			fmt.Println("Successfully logged in using Azure CLI.")
			authorizer, err = auth.NewAuthorizerFromCLI()
			if err != nil {
				return err
			}
			return nil
		} else {
			fmt.Println("Successfully logged in using Azure CLI.")
			authorizer, err = auth.NewAuthorizerFromCLI()
			if err != nil {
				return err
			}
			return nil
		}
	}
	// Use authorizer with your Azure SDK clients
	return nil
}
