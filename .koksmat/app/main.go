package main

import (
	"runtime/debug"
	"strings"

	"github.com/365admin/koksmat-door/magicapp"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: Koksmat Door
description: Entry point to the kitchens
---

# Koksmat Door

Koksmat Door is based on the Caddy web server which protects all exposed services.`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("Koksmat Door", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.Execute(name, "Koksmat Door", "")
}
