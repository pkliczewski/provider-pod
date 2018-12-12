package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pkliczewski/provider-pod/client"
)

func main() {
	ctx := context.Background()

	c, err := client.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Logout(ctx)

	vms, err := c.GetVMs(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print summary per vm (see also: govc/vm/info.go)
	for _, vm := range vms {
		fmt.Printf("%s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName)
	}
}
