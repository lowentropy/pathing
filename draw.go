package main

import (
	"code.google.com/p/draw2d/draw2d"
	"image"
	"image/png"
)

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

func main() {
	i := image.NewRGBA(image.Rect(0, 0, 200, 200))
	gc := draw2d.NewGraphicContext(i)
	gc.MoveTo(10.0, 10.0)
	gc.LineTo(100.0, 10.0)
	gc.Stroke()
	saveToPngFile("TestPath.png", i)
}
