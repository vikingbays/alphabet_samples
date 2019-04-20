// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package main

import (
	"alphabet/cmd/utils"
	"path/filepath"
	"strings"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir("."))
	lidx := strings.LastIndex(dir, "src")

	projectName := dir[:(lidx - 1)]

	utils.DoCmd("-start", projectName)
}
