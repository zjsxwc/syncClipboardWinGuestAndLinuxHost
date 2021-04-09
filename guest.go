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
		winContent, _ := clipboard.ReadAll()
		if len(winContent) > 0 {
			if (winContent != oldWinContent) {
				if (winContent != oldLinuxContent) {
					println("send clipboard data to linux")
					println(winContent)
	
					oldWinContent = winContent
					php2go.FilePutContents("winclipboard.data", winContent, 0777)
				}
			}
		}
	
	
		linuxContent, _ := php2go.FileGetContents("linuxclipboard.data")
		if len(linuxContent) > 0 {
			if linuxContent != oldWinContent {
				if (linuxContent != oldLinuxContent) {
					println("write linux data to clipboard")
					println(linuxContent)
	
					oldLinuxContent = linuxContent
					clipboard.WriteAll(linuxContent)
					php2go.FilePutContents("linuxclipboard.data", "", 0777)
				}
			}
		}

		time.Sleep(time.Duration(600)*time.Millisecond)
	}


}
