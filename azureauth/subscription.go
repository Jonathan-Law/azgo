package azureauth

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
)

func SelectSubscription() (string, error) {
	subClient := subscriptions.NewClient()
	subClient.Authorizer = authorizer

	ctx := context.Background()
	result, err := subClient.List(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to list subscriptions: %v", err)
	}

	fmt.Println("Select a subscription to use:")
	var subs []subscriptions.Subscription
	for result.NotDone() {
		subs = append(subs, result.Values()...)
		result.NextWithContext(ctx)
	}

	for i, sub := range subs {
		fmt.Printf("%d: %s (%s)\n", i+1, *sub.DisplayName, *sub.SubscriptionID)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number: ")
	scanner.Scan()
	choice := scanner.Text()
	index, _ := strconv.Atoi(choice)
	selectedSub := subs[index-1]

	cmd := exec.Command("az", "account", "set", "--subscription", *selectedSub.SubscriptionID)
	if _, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to set subscription: %v", err)
	}

	return *selectedSub.SubscriptionID, nil
}
