// Copyright (c) 2025 Grigoriy Efimov
//
// Licensed under the MIT License. See LICENSE file in the project root for details.

package main

import (
	"image"
	"image/color"

	ink "github.com/CatInBeard/inkview"
)

const defaultFontSize = 14

type App struct {
	font  *ink.Font
	fontH int
	fontW int
}

func (a *App) Init() error {
	ink.ClearScreen()
	ink.DrawTopPanel()

	a.font = ink.OpenFont(ink.DefaultFontMono, defaultFontSize, true)
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

	text := GetCurrentTranslation("example_text")

	ink.DrawString(image.Point{screenSize.X / 5, screenSize.Y / 2}, text)

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
