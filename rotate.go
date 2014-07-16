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

var rotation string

func main() {
	fName, fDest, rotation := extract_args()
	fmt.Printf("rotating %s by %f. destination %s\n", fName, rotation, fDest)

	fImg, _ := os.Open(fName)
	defer fImg.Close()		
	width, height := get_size(fName)
	img, _, _ := image.Decode(fImg)
	m := image.NewRGBA(image.Rect(0, 0, height, width))
	graphics.Rotate(m, img, &graphics.RotateOptions{rotation})

	toimg, _ := os.Create(fDest)
	defer toimg.Close()

	jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})
}

func get_size(fName string) (int, int) {
	fImgSize, _ := os.Open(fName)
	defer fImgSize.Close()
    imgConfig, _, err := image.DecodeConfig(fImgSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fName, err)
	}

	return imgConfig.Width, imgConfig.Height
}

func extract_args() (string, string, float64) {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("3 args required: image source, destination and rotation")
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Println("Too many arguments, required: image source, destination and rotation")
		os.Exit(1)
	}
	fName := args[0]
    fDest := args[1]
    rotation := args[2]
    var rotationAngle float64 = 0
    if rotation == "r" {
    	rotationAngle = 1.57079633
    } else {
    	rotationAngle = 4.71238899
    }
    return fName, fDest, rotationAngle
}
