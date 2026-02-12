package main

import (
	"bufio"
	"fmt"
	"goemon/src/cmd/cli"
	"goemon/src/core/agent"
	"goemon/src/core/db"
	"goemon/src/core/image"
	"goemon/src/core/mail"
	"goemon/src/core/web"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Initialize Global DB (Image)
	if err := db.InitGlobalDB(); err != nil {
		fmt.Printf("Warning: Could not initialize Image DB: %v\n", err)
	}
	defer func() {
		if db.GlobalDB != nil {
			db.GlobalDB.Close()
		}
	}()

	// Initialize Global DB (Mail)
	if err := db.InitGlobalMailDB(); err != nil {
		fmt.Printf("Warning: Could not initialize Mail DB: %v\n", err)
	}
	defer func() {
		if db.GlobalMailDB != nil {
			db.GlobalMailDB.Close()
		}
	}()

	// Initialize Global DB (Vector)
	if err := db.InitGlobalVectorDB(); err != nil {
		fmt.Printf("Warning: Could not initialize Vector DB: %v\n", err)
	}
	defer func() {
		if db.GlobalVectorDB != nil {
			db.GlobalVectorDB.Close()
		}
	}()

	// Initialize Image Service
	image.InitImageService("C:/models/vlm/moondream2")

	// Start Image Workers (in background)
	bakerConfig := mail.GetDefaultConfig()
	bakerConfig.User = agent.GetAgentEmail("baker")

	watcherConfig := mail.GetDefaultConfig()
	watcherConfig.User = agent.GetAgentEmail("watcher")

	image.StartWorkers(bakerConfig, watcherConfig)

	// Start Web Console (in background)
	go web.StartWebServer(8080)

	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()
		fmt.Print("Select option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		// --- System & Infrastructure ---
		case "1":
			cli.RunDashboard()
		case "2":
			cli.ShowDataMetrics()
		case "3":
			fmt.Println("Restarting system...")
			runCommand("docker-compose", "down", "--remove-orphans")
			runCommand("docker-compose", "build")
			runCommand("docker-compose", "up", "-d")
			fmt.Println("Waiting for Mail Server to initialize (15s)...")
			time.Sleep(15 * time.Second)
			fmt.Println("Setting up default email account...")
			runCommand("docker", "exec", "mailserver", "setup", "email", "add", "user@mail.local", "password")
			fmt.Println("System Rebuild Complete.")

		// --- Mail Operations ---
		case "4":
			fmt.Println("Receiving mail...")
			cli.RunMailReceiverUI()
		case "5":
			cli.RunMailSenderUI()
		case "6":
			runCommand("docker", "logs", "-f", "mailserver", "--tail", "50")

		// --- AI Workers & Agents ---
		case "7":
			cli.CheckWorkerStatus()
		case "8":
			listModels()
		case "9":
			runCommand("docker", "logs", "-f", "sd-worker", "--tail", "50")
		case "10":
			fmt.Println("Starting Foreman Agent (Background logic)...")
			go agent.StartForeman()
			fmt.Println("Foreman is running.")

		// --- Development & Others ---
		case "11":
			cli.CreateADR()
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("\033[91mInvalid option\033[0m")
		}

		fmt.Println("\nPress Enter to continue...")
		reader.ReadString('\n')
	}
}

func showMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("=== Goemon Core Command Center ===")

	fmt.Println("\n[ System & Infrastructure ]")
	fmt.Println(" 1. System Dashboard (Live)")
	fmt.Println(" 2. Data Metrics (Taijukei)")
	fmt.Println(" 3. Restart System (Compose Up)")

	fmt.Println("\n[ Mail Operations ]")
	fmt.Println(" 4. Receive Mail (Download Images)")
	fmt.Println(" 5. Send Prompt (Email)")
	fmt.Println(" 6. Mail Server Logs (Follow)")

	fmt.Println("\n[ AI Workers & Agents ]")
	fmt.Println(" 7. SD Worker Status (Detailed)")
	fmt.Println(" 8. List Models")
	fmt.Println(" 9. SD Worker Logs (Follow)")
	fmt.Println("10. Start Foreman Agent")

	fmt.Println("\n[ Development & Others ]")
	fmt.Println("11. Create New ADR")
	fmt.Println(" 0. Exit")
	fmt.Println("=================================")
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func listModels() {
	modelDir := filepath.Join("src", "workers", "Image_Baker", "models")
	fmt.Printf("Checking models in: %s\n", modelDir)

	err := filepath.Walk(modelDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Printf("- %-30s (%d MB)\n", info.Name(), info.Size()/(1024*1024))
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error listing models: %v\n", err)
	}
}
