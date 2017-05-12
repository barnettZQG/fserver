package main

import (
	"log"

	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/valyala/fasthttp"
)

var root string
var port int

func init() {
	fs := pflag.CommandLine
	fs.StringVar(&root, "root", "/var/www", "root path")
	fs.IntVar(&port, "port", 8080, "fs server listen port")
}
func main() {
	pflag.Parse()
	fs := &fasthttp.FS{
		// Path to directory to serve.
		Root: root,

		// Generate index pages if client requests directory contents.
		GenerateIndexPages: false,

		// Enable transparent compression to save network traffic.
		Compress: true,
	}

	// Create request handler for serving static files.
	h := fs.NewRequestHandler()

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	logrus.Info("f-server listen " + addr)
	// Start the server.
	if err := fasthttp.ListenAndServe(addr, h); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
