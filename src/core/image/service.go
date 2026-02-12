package image

import (
	"fmt"
	"goemon/src/core/mail"
)

func InitImageService(modelPath string) {
	fmt.Printf("Initializing Image Service with model: %s\n", modelPath)
}

func StartWorkers(bakerConfig, watcherConfig mail.MailConfig) {
	fmt.Println("Starting Image Workers...")
	// Logic to start baker and watcher workers
}
