package client

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

type Client struct {
	Client *govmomi.Client
}

func (c *Client) GetVMs(ctx context.Context) ([]mo.VirtualMachine, error) {
	client := c.Client

	// Create view of VirtualMachine objects
	m := view.NewManager(client.Client)

	v, err := m.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, err
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		return nil, err
	}

	return vms, nil
}

func (c *Client) Logout(ctx context.Context) error {
	client := c.Client
	return client.Logout(ctx)
}

func NewClient(ctx context.Context) (*Client, error) {
	host := ""
	username := ""
	password := ""
	insecure := true

	u := &url.URL{
		Scheme: "https",
		User:   url.UserPassword(username, password),
		Host:   host,
		Path:   vim25.Path,
	}

	// Connect and log in to ESX or vCenter
	client, err := govmomi.NewClient(ctx, u, insecure)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Client: client,
	}
	return c, nil
}
