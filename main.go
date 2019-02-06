package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/stevenimle-honeywell/docker-machine-driver-azure2/pkg/drivers/azure"
)

func main() {
	plugin.RegisterDriver(azure.NewDriver("", ""))
}
