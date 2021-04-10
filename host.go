package main

import (
	"github.com/atotto/clipboard"
	"time"

	"github.com/syyongx/php2go"
	"os/exec"
)

func main() {

	oldWinContent := ""
	oldLinuxContent := ""


	for {

		winContent, _ := php2go.FileGetContents("winclipboard.data")
		if len(winContent) > 0 {
			if winContent != oldWinContent {
				if winContent != oldLinuxContent {
					oldWinContent = winContent
					clipboard.WriteAll(winContent)
					time.Sleep(time.Duration(20)*time.Millisecond)
					php2go.FilePutContents("winclipboard.data", "", 0777)
				}
			}
		}

		if php2go.FileExists("image.jpg") {
			//xclip -selection clipboard -t image/jpeg -i image.jpg
			//xclip -selection clipboard -t image/jpeg -o|tee ggg.jpg >/dev/null
			time.Sleep(time.Duration(200)*time.Millisecond)

			cmdArgs := []string{"xclip", "-selection", "clipboard", "-t", "image/jpeg", "-i", "image.jpg"}
			cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
			in, _ := cmd.StdinPipe()
			cmd.Start()
			in.Close()
			cmd.Wait()



			cmdArgs = []string{"rm", "image.jpg"}
			cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
			in, _ = cmd.StdinPipe()
			cmd.Start()
			in.Close()
			cmd.Wait()

			println("jpg")
		}

		linuxContent, _ := clipboard.ReadAll()
		if len(linuxContent) > 0 {
			if linuxContent != oldLinuxContent {
				if linuxContent != oldWinContent {
					oldLinuxContent = linuxContent
					php2go.FilePutContents("linuxclipboard.data", linuxContent, 0777)
				}
			}
		}

		time.Sleep(time.Duration(600)*time.Millisecond)

	}

}
