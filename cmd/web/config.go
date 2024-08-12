package main

import "flag"

type config struct {
	addr      string
	staticDir string
}

func newConfig() config {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	return cfg
}
