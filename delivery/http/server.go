package http

import "fmt"

// Init - Initalize HTTP Server
func Init(port int) {
	r := NewRouter()

	portStr := fmt.Sprintf(":%d", port)

	r.Run(portStr)
}
