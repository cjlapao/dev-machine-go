package main

import (
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/version"
)

var logger = log.Get()
var ver = version.Get()

func main() {
	logger.Info("testubg")
	ver.Name = "Development Machine Tools"
	ver.Author = "Carlos Lapao"
	ver.License = "MIT"
	ver.Rev = 1
	ver.PrintAnsiHeader()
	ver.PrintHeader()
}
