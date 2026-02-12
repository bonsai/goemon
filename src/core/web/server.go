package web

import (
	"encoding/json"
	"fmt"
	"goemon/src/core/agent"
	"goemon/src/core/db"
	"goemon/src/core/mail"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type PageData struct {
	Images []string
}

func StartWebServer(port int) {
	outputDir := "src/workers/Image_Baker/outputs"
	os.MkdirAll(outputDir, 0755)

	http.Handle("/outputs/", http.StripPrefix("/outputs/", http.FileServer(http.Dir(outputDir))))

	// HTML Main Page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		images := getImages(outputDir)
		tmpl, err := template.ParseFiles("src/core/web/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, PageData{Images: images})
	})

	// API: Get Image List
	http.HandleFunc("/api/images", func(w http.ResponseWriter, r *http.Request) {
		images := getImages(outputDir)
		json.NewEncoder(w).Encode(images)
	})

	// API: T2I (Anything V5) - 9枚生成
	http.HandleFunc("/api/draw", func(w http.ResponseWriter, r *http.Request) {
		prompt := r.FormValue("prompt")
		countStr := r.FormValue("count")
		count, _ := strconv.Atoi(countStr)
		if count <= 0 {
			count = 1
		}

		fmt.Printf("Web Request: Baking %d images with prompt: %s\n", count, prompt)

		// 既存のメール送信ロジックを回数分叩く
		config := mail.GetDefaultConfig()
		config.User = "user@mail.local"

		for i := 0; i < count; i++ {
			// 実際には mail.SendPrompt(config, "baker@mail.local", "/draw", prompt)
		}

		w.WriteHeader(http.StatusOK)
	})

	// API: VLM (Moondream2) - 画像への質問
	http.HandleFunc("/api/vlm", func(w http.ResponseWriter, r *http.Request) {
		path := r.FormValue("path")
		query := r.FormValue("query")

		fmt.Printf("Web Request: VLM Question on %s: %s\n", path, query)

		// モック回答（実際には Python Bridge 経由で Moondream2 を叩く）
		response := map[string]string{
			"answer": fmt.Sprintf("I see the image '%s'. Regarding your question '%s', it looks like a high-quality AI generated artifact.", path, query),
		}
		json.NewEncoder(w).Encode(response)
	})

	// API: GLM (GLM-4) - テキスト生成
	http.HandleFunc("/api/glm", func(w http.ResponseWriter, r *http.Request) {
		seed := r.FormValue("seed")

		fmt.Printf("Web Request: GLM Generation with seed: %s\n", seed)

		// SpellWriter を使用して GLM 呼び出し
		sw := agent.NewSpellWriter(db.GlobalVectorDB)
		text, err := sw.CastSpell(seed)
		if err != nil {
			text = "Error generating text: " + err.Error()
		}

		response := map[string]string{
			"text": text,
		}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Printf("Goemon Swarm Console started at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getImages(dir string) []string {
	files, _ := os.ReadDir(dir)
	var images []string
	for _, f := range files {
		if !f.IsDir() && isImage(f.Name()) {
			images = append(images, f.Name())
		}
	}
	sort.Slice(images, func(i, j int) bool {
		return images[i] > images[j]
	})
	if len(images) > 50 {
		images = images[:50]
	}
	return images
}

func isImage(name string) bool {
	ext := filepath.Ext(name)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	}
	return false
}
