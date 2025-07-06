// Copyright (c) 2025 Grigoriy Efimov
//
// Licensed under the MIT License. See LICENSE file in the project root for details.

package main

import (
	"log"

	ink "github.com/dennwc/inkview"
)

func main() {

	app := &App{}

	if err := ink.Run(app); err != nil {
		log.Fatal(err)
	}

}
