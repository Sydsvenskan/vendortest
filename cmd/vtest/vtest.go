package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/Sydsvenskan/vendortest"
)

func main() {
	vendortest.StartWithLogger(logrus.StandardLogger())
}
