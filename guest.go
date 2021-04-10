package main

import (
	"bytes"
	"github.com/syyongx/php2go"
	"golang.org/x/image/bmp"
	"os"
	"time"
)

const (
	typeText  = "text"
	typeFile  = "file"
)


func readClipboard() string {
	contentType, err := Clipboard().ContentType()
	if err != nil {
		println(err.Error())
		return ""
	}
	if contentType == typeText {
		str, err := Clipboard().Text()
		if err != nil {
			println(err.Error())
			Clipboard().Clear()
			Clipboard().SetText(oldWinContent)
			return ""
		}
		return str
	}
	if contentType == "CF_DIBV5" {
		bmpBytes, err := Clipboard().Bitmap()
		if err != nil {
			return ""
		}
		bmpBytesReader := bytes.NewReader(bmpBytes)
		bmpImage, err := bmp.Decode(bmpBytesReader)
		if err != nil {
			return ""
		}
		bmpfile, _ := os.Create(`./bmp.bmp`)
		bmp.Encode(bmpfile, bmpImage)
		bmpfile.Close()
		return ""
	}
	if contentType == typeFile {
		filenames, err := Clipboard().Files()
		if err != nil {
			return ""
		}
		for _, path := range filenames {
			println(path)
		}
		return ""
	}
	return ""
}


func writeClipboard(d string) {
	if err := Clipboard().SetText(d); err != nil {
		println(err.Error())
		return
	}
}


var oldWinContent = ""
var oldLinuxContent = ""

func main() {
	for {
		winContent := readClipboard()
		println(winContent)
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
					writeClipboard(linuxContent)
					php2go.FilePutContents("linuxclipboard.data", "", 0777)
				}
			}
		}

		time.Sleep(time.Duration(600)*time.Millisecond)
	}
}
