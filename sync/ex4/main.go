package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("letft.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return Icon(name)
}