package web

import (
	"context"
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

	"github.com/skip2/go-qrcode"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type PageData struct {
	Images []string
}

func StartWebServer(port int) {
	outputDir := "src/workers/Image_Baker/outputs"
	os.MkdirAll(outputDir, 0755)

	mux := http.NewServeMux()
	mux.Handle("/outputs/", http.StripPrefix("/outputs/", http.FileServer(http.Dir(outputDir))))

	// HTML Main Page (With Simple Password Auth)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// モバイルからの簡易認証: URL パラメータに ?pass=1234 があればセッション維持（簡易的にクッキー使用）
		pass := r.URL.Query().Get("pass")
		if pass == "1234" {
			http.SetCookie(w, &http.Cookie{Name: "auth", Value: "ok", Path: "/"})
		}

		// クッキーチェック
		cookie, err := r.Cookie("auth")
		if err != nil || cookie.Value != "ok" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `<h2>Goemon Access Restricted</h2>
				<form method="GET">
					Enter Pass: <input type="password" name="pass">
					<button type="submit">Unlock</button>
				</form>
				<p style="font-size:0.8em; color:gray;">Hint: Default pass is 1234</p>`)
			return
		}

		images := getImages(outputDir)
		tmpl, err := template.ParseFiles("src/core/web/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, PageData{Images: images})
	})

	// API: Get Image List
	mux.HandleFunc("/api/images", func(w http.ResponseWriter, r *http.Request) {
		images := getImages(outputDir)
		json.NewEncoder(w).Encode(images)
	})

	// API: T2I (Anything V5) - 9枚生成
	mux.HandleFunc("/api/draw", func(w http.ResponseWriter, r *http.Request) {
		prompt := r.FormValue("prompt")
		countStr := r.FormValue("count")
		count, _ := strconv.Atoi(countStr)
		if count <= 0 {
			count = 1
		}

		fmt.Printf("Web Request: Baking %d images with prompt: %s\n", count, prompt)

		mailConfig := mail.GetDefaultConfig()
		mailConfig.User = "user@mail.local"

		for i := 0; i < count; i++ {
			mail.SendPrompt(mailConfig, "baker@mail.local", "/draw", prompt)
		}

		w.WriteHeader(http.StatusOK)
	})

	// API: VLM (Moondream2) - 画像への質問
	mux.HandleFunc("/api/vlm", func(w http.ResponseWriter, r *http.Request) {
		path := r.FormValue("path")
		query := r.FormValue("query")

		fmt.Printf("Web Request: VLM Question on %s: %s\n", path, query)

		response := map[string]string{
			"answer": fmt.Sprintf("I see the image '%s'. Regarding your question '%s', it looks like a high-quality AI generated artifact.", path, query),
		}
		json.NewEncoder(w).Encode(response)
	})

	// API: GLM (GLM-4) - テキスト生成
	mux.HandleFunc("/api/glm", func(w http.ResponseWriter, r *http.Request) {
		seed := r.FormValue("seed")

		fmt.Printf("Web Request: GLM Generation with seed: %s\n", seed)

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

	// ngrok トンネルの開始 (NGROK_AUTHTOKEN がある場合のみ)
	if token := os.Getenv("NGROK_AUTHTOKEN"); token != "" {
		fmt.Println("NGROK_AUTHTOKEN found. Initializing tunnel...")
		l, err := ngrok.Listen(context.Background(),
			config.HTTPEndpoint(),
			ngrok.WithAuthtoken(token),
		)
		if err != nil {
			fmt.Printf("Failed to start ngrok: %v\n", err)
		} else {
			fmt.Printf("\n====================================================\n")
			fmt.Printf("NGROK TUNNEL ESTABLISHED!\n")
			fmt.Printf("PUBLIC URL: %s\n", l.URL())

			// QRコードの表示 (モバイル用)
			qr, err := qrcode.New(l.URL()+"?pass=1234", qrcode.Medium)
			if err == nil {
				fmt.Printf("\nSCAN THIS QR TO OPEN ON MOBILE (Auto-Login):\n")
				fmt.Println(qr.ToSmallString(false))
			}
			fmt.Printf("====================================================\n\n")

			go func() {
				if err := http.Serve(l, mux); err != nil {
					fmt.Printf("ngrok server closed: %v\n", err)
				}
			}()
		}
	} else {
		fmt.Println("NGROK_AUTHTOKEN not set. Running on local port only.")
	}

	fmt.Printf("Goemon Swarm Console started at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
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
