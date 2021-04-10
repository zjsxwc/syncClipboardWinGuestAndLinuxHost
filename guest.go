package main

import (
	"bytes"
	"github.com/syyongx/php2go"
	"golang.org/x/image/bmp"
	"image/jpeg"
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
		bmpBytes, err = Clipboard().Bitmap()

		if err != nil {
			bmpBytes = nil
			return ""
		}
		bmpBytesReader := bytes.NewReader(bmpBytes)
		//bmpImage, err := bmp.Decode(bmpBytesReader)
		bmpImage, err := bmp.Decode(bmpBytesReader)
		if err != nil {
			bmpImage = nil
			return ""
		}
		jpgfile, _ := os.Create(`./image.jpg`)
		jpeg.Encode(jpgfile, bmpImage, &jpeg.Options{100})
		jpgfile.Close()
		bmpBytes = nil
		bmpImage = nil
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
var bmpBytes []byte
func main() {
	for {
		winContent := readClipboard()
		if len(winContent) > 0 {
			println(winContent)
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
