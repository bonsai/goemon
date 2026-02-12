package image

import "fmt"

type ImageMetadata struct {
	ID     string
	Path   string
	Prompt string
}

func SaveMetadata(meta ImageMetadata) error {
	fmt.Printf("Saving metadata for image: %s\n", meta.ID)
	return nil
}
