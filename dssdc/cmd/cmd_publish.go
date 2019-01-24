package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	publishMsg string
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the project by git/fileupload-api.",
	Long: `Publish the project by git/fileupload-api.
You should set msg to use git.
You should install Git to use the function.`,

	Run: func(cmd *cobra.Command, args []string) {
		if publishMsg != "" {
			iDirs := cfg.Monitors.GetStringSlice("includeDirs")
			for _, v := range iDirs {
				runExec("git", "add", v)
			}
			runExec("git", "commit", "-m", publishMsg)
			runExec("git", "push")
			if !allPublishWay {
				return
			}
		}

		tgName := cfg.ProjectName + ".tar.gz"
		fmt.Println("doing")
		pkg(tgName)
		os.Remove(tgName)
	},
}

func runExec(vv string, v ...string) {
	mustLog(append([]string{vv}, v...))
	cmd := exec.Command(vv, v...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = wd
	cmd.Env = os.Environ()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		mustLog("error=>", err.Error())
		return
	}
	err = cmd.Start()
	if err != nil {
		mustLog(err)
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	err = cmd.Wait()
	if err != nil {
		mustLog("cmd wait err ", err)
		return
	}
	if cmd.Process != nil {
		if err = cmd.Process.Kill(); err != nil {
			mustLog("cmd cannot kill ", err)
		}
	}
	cmd.Start()
}
