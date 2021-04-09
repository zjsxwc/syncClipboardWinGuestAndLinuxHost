package main

import (
	"time"
	"github.com/atotto/clipboard"

	"github.com/syyongx/php2go"
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
					php2go.FilePutContents("winclipboard.data", "", 0777)
				}
			}
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
