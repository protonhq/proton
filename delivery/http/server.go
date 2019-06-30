package http

// Init - Initalize HTTP Server
func Init() {
	r := NewRouter()
	r.Run(":8080")
}