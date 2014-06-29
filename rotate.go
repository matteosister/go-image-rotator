package main

import (
    "fmt"
    "os"
    //"image/draw"
    "image"
    "image/jpeg"
  	_ "image/png"
    "code.google.com/p/graphics-go/graphics"
)

func get_size(fName string) (int, int) {
	fImgSize, _ := os.Open(fName)
	defer fImgSize.Close()
    imgConfig, _, err := image.DecodeConfig(fImgSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fName, err)
	}

	return imgConfig.Width, imgConfig.Height
}

func main() {
    args := os.Args[1:]
	for _, fName := range args {
		fImg, _ := os.Open(fName)
		defer fImg.Close()		
		width, height := get_size(fName)
		img, _, _ := image.Decode(fImg)
		m := image.NewRGBA(image.Rect(0, 0, height, width))
		graphics.Rotate(m, img, &graphics.RotateOptions{1.57079633})

		toimg, _ := os.Create("new.jpg")
		defer toimg.Close()

		jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})
	}
}
