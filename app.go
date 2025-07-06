// Copyright (c) 2025 Grigoriy Efimov
//
// Licensed under the MIT License. See LICENSE file in the project root for details.

package main

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	ink "github.com/dennwc/inkview"
)

const defaultFontSize = 25

type App struct {
	font  *ink.Font
	fontH int
	fontW int
}

func (a *App) Init() error {
	ink.ClearScreen()
	ink.DrawTopPanel()

	if a.fontH == 0 {
		a.fontH = defaultFontSize
	}

	a.font = ink.OpenFont(ink.DefaultFontMono, a.fontH, true)
	a.font.SetActive(color.RGBA{0, 0, 0, 255})
	a.fontW = ink.CharWidth('a') // Work only for monospace font

	a.Draw()
	ink.Repaint()

	return nil
}

func (a *App) Close() error {
	return nil
}

func (a *App) Draw() {
	ink.ClearScreen()
	ink.DrawTopPanel()
	a.font.SetActive(color.Black)

	screenSize := ink.ScreenSize()

	text := ""

	text += fmt.Sprintln("Device:")
	text += fmt.Sprintln(ink.DeviceKey())
	text += fmt.Sprintln(ink.DeviceModel())
	text += fmt.Sprintln(ink.SoftwareVersion())
	text += fmt.Sprintln(ink.HwAddress())
	text += fmt.Sprintln(ink.GetCurrentLang())
	text += fmt.Sprintln(ink.HardwareType())
	text += fmt.Sprintln(ink.SerialNumber())
	text += fmt.Sprintln("Screen:")
	text += fmt.Sprintf("%dx%d\n", screenSize.X, screenSize.Y)
	text += fmt.Sprintln("Software:")
	text += fmt.Sprintln(ink.SoftwareVersion())
	text += fmt.Sprintln("Sensors:")
	text += fmt.Sprintln(ink.Temperature())

	startPoint := image.Point{screenSize.X / 10, screenSize.Y / 10}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		ink.DrawString(startPoint, line)
		startPoint.Y += a.fontH
	}

	ink.PartialUpdate(image.Rectangle{image.Point{0, 0}, screenSize})

}

func (a *App) Key(e ink.KeyEvent) bool {
	return true
}

func (a *App) Pointer(e ink.PointerEvent) bool {
	return true
}

func (a *App) Touch(e ink.TouchEvent) bool {
	return true
}

func (a *App) Orientation(o ink.Orientation) bool {
	return true
}

func requestNetworkConnection() {
	ink.QueryNetwork()
	err := ink.ConnectDefault()
	if err != nil {
		ink.Exit()
	}
}
