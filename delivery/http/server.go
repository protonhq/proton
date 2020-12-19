package http

import (
	"fmt"

	"github.com/protonhq/proton/infra/config"
	"github.com/protonhq/proton/registry"
)

// Init - Initalize HTTP Server
func Init(ctn *registry.Container) {
	r := NewRouter(ctn)
	appConfig := ctn.Resolve("config").(*config.Configuration)
	portStr := fmt.Sprintf(":%d", appConfig.Server.Port)
	r.Run(portStr)
}
