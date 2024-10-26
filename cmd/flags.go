package cmd

import "flag"

func Load() string {
	port := flag.String("port", ":80", "server port")

	flag.Parse()

	return *port
}
