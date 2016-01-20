package mydraw

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"os"
)

func MakeGif(name string, frames int, delay int, framer func(int) (img draw.Image)) {
	//
	log.Println("GIF CREATION START")
	frameList := make([]*image.Paletted, frames)
	delays := make([]int, frames)
	for i := 0; i < frames; i++ {
		frameList[i] = Convert2Paletted(framer(i))
		delays[i] = delay
	}
	log.Println("GIF CREATION COMPLETE")
	//
	log.Println("CREATING FILE")
	file, err := os.Create(name + ".gif")
	if err != nil {
		log.Fatal("FILE CREATE ERR:", err)
	}
	defer file.Close()
	log.Println("GIF ENCODE START")
	if err = gif.EncodeAll(file, &gif.GIF{Image: frameList, Delay: delays, LoopCount: 0}); err != nil {
		log.Println("ENCODE ERR:", err)
	}
	log.Println("GIF ENCODE COMPLETE")
}

func Convert2Paletted(img draw.Image) *image.Paletted {
	bounds := img.Bounds()
	res := image.NewPaletted(bounds, palette.Plan9)
	draw.Draw(res, bounds, img, bounds.Min, draw.Src)
	return res
}
