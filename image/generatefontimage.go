package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	dpi      = flag.Float64("dpi", 150, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./image/serendity.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 36, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

var text = []string{
	"First step is to acknowledge and call out the bias",
}

func main() {
	flag.Parse()

	var textToDraw = "Later on, do they also have the product/business sense to build it out fully and scale it"
	var textToDrawSize = len(textToDraw)
	var textFontSize = 36.0

	fmt.Println("String Length ", textToDrawSize)

	//Read the font data
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, 1080, 1080))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	//Draw the guidelines
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	if (textToDrawSize > 0) && (textToDrawSize <= 25) {
		c.SetFontSize(textFontSize)
	} else if (textToDrawSize > 25) && (textToDrawSize <= 100) {
		textFontSize = 24
		c.SetFontSize(textFontSize)
	}

	// Draw the text
	pt := freetype.Pt(50, 500+int(c.PointToFixed(textFontSize)>>6))
	_, err = c.DrawString(textToDraw, pt)
	if err != nil {
		log.Println(err)
		return
	}
	//pt.Y += c.PointToFixed(textFontSize * *spacing)
	/*
		for _, s := range text {
			fmt.Println("TEXT LENGTH ", len(s))
			_, err = c.DrawString(s, pt)
			if err != nil {
				log.Println(err)
				return
			}
			pt.Y += c.PointToFixed(*size * *spacing)
		}
	*/

	//Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
}
