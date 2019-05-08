// Package 只做了一个简单的灰阶图，没有去上色
package main

import (
	"fmt"
	"image"
	"os"

	"image/color"
	"image/jpeg"
	_ "image/jpeg"
)

func AddPhtoFrame() {
	target := "res/ret.jpg"
	//TODO:use the res/gophergala.jpg to generate a image and write to res/m.jpg which is similar to like the logo in the README.md
	file, _ := os.Open("res/gophergala.jpg")
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	// 转化
	rgba := hdImage(img)
	f, _ := os.Create(target)
	defer f.Close()
	// 编码到文件
	jpeg.Encode(f, rgba, nil)
	fmt.Println("W:", width, "H:", height)
}

// 图片灰化处理
func hdImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			_, g, _, a := colorRgb.RGBA()
			g_uint8 := uint8(g >> 8)
			a_uint8 := uint8(a >> 8)
			if i == 0 && j == 0 {
				fmt.Println("g/a", g, a)
				fmt.Println("g/a uint8", g_uint8, a_uint8)
			}
			newRgba.SetRGBA(i, j, color.RGBA{g_uint8, g_uint8, g_uint8, a_uint8})
		}
	}
	return newRgba
}

func main() {
	AddPhtoFrame()
	println(`This final exercise,let's add a photo frame for gala logo!
You should use image package to generate a new iamge from the gala log(which is stored in res/gophergala.jpg,and makes it like the res/m.jpg.
Now edit main.go to complete 'AddPhtoFrame' function,this task has no test,enjoy your trip!`)
}
