package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"see"},
	Short:   "See others project.",
	Long:    `See others project.`,

	Run: func(cmd *cobra.Command, args []string) {
		getServer()
	},
}

func getServer() {
	port := os.Getenv("DSSDC_PORT")
	if port == "" {
		port = ":" + strconv.Itoa(cfg.Port)
	}
	// TODO: auto port inc
	log.Println(port)

	// TODO: get

	// TODO: 优化（考虑第三方库）
	// open the web browser
	run, ok := openCommands[runtime.GOOS]
	if !ok {
		mustLog("WARNING", "don't know how to open things on %s platform", runtime.GOOS)
	}
	cmd := exec.Command(run, "http://localhost"+port)
	go cmd.Start()

	fmt.Printf("Listen %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
