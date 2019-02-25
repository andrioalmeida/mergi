package main

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/io"
	"image"
	"log"
	"fmt"
)

// Mergi supports url to image, Lets try to merge this placeholders
func main() {
	// Opps Sorry ! let's ignore error for now
	img1, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/500x500"))

	img2, _ := mergi.Import(io.NewBase64Importer("data:image/jpg;base64,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"))
	img3, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/250x250"))

	img4, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/100x100"))
	img5, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/100x100"))
	img6, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/100x100"))
	img7, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/100x100"))
	img8, _ := mergi.Import(io.NewURLImporter("https://via.placeholder.com/100x100"))

	imgs := []image.Image{img1, img2, img3, img4, img5, img6, img7, img8}

	res, err := mergi.Merge("TTBTBBBB", imgs)
	if err != nil {
		log.Fatal(err)
	}
	anim, _ := mergi.Animate([]image.Image{img1, img2}, 20)

	// export image file png/jpg
	mergi.Export(io.NewFileExporter(res, "examples/import_export/res/out.png"))
	mergi.Export(io.NewBase64Exporter("png", res, func(data string) {
		fmt.Println(data)
	}))

	// Export animation gif
	mergi.Export(io.NewAnimationExporter(anim, "examples/import_export/res/out.gif"))
}
