package programs

import (
	"image"
	"image/draw"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/exp/shiny/screen"
)

type GalleryProgram struct {
	buffer screen.Buffer
	images []image.Image
}

func NewGalleryProgram(imageDirs ...string) *GalleryProgram {
	images := make([]image.Image, 0)

	for _, dirPath := range imageDirs {
		filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Printf("failed to access path %q: %v", path, err)
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				log.Printf("failed to open image: %v", err)
				return err
			}
			defer file.Close()

			if d.IsDir() {
				return nil
			}

			img, _, err := image.Decode(file)
			if err != nil {
				log.Printf("failed to load image: %v", err)
				return err
			}

			images = append(images, img)
			return nil
		})
	}

	return &GalleryProgram{
		images: images,
	}
}

func (program *GalleryProgram) Init(buffer screen.Buffer) {
	program.buffer = buffer
	draw.Draw(program.buffer.RGBA(), program.buffer.Bounds(), program.images[0], image.Point{0, 0}, draw.Src)
}

func (program *GalleryProgram) Draw() screen.Buffer {
	return program.buffer
}
