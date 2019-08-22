package http

import (
	"fmt"

	"github.com/protonhq/proton/infra/config"
	"github.com/protonhq/proton/registry"
)

// Init - Initalize HTTP Server
func Init(conf *config.Configuration, ctn *registry.Container) {
	r := NewRouter(ctn)

	portStr := fmt.Sprintf(":%d", conf.Server.Port)

	r.Run(portStr)
}
