package main

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"strings"

	"github.com/xfmoulet/qoi"
)

var (
	assetFolder []fs.DirEntry
	err         error
	i           int
	itemName    string
	openedFile  *os.File
	imgPng      image.Image
	imgQoi      *os.File
)

func main() {
	assetFolder, err := os.ReadDir("./assets")
	if err != nil {
		fmt.Println("An error occurred reading AssetFolder: ", err)
		return
	}

	for i := range assetFolder {
		itemName := assetFolder[i].Name()
		fmt.Println(fmt.Sprintf("converting 'input/%v'", itemName))
		if strings.Contains(itemName, ".png") {

			openedFile, err = os.Open(fmt.Sprintf("input/%v", itemName))
			if err != nil {
				fmt.Printf("An error occurred opening %v: %v", itemName, err)
				return
			}
			imgPng, err = png.Decode(openedFile)
			if err != nil {
				fmt.Printf("An error occurred decoding %v: %v", itemName, err)
				return
			}
			imgQoi, err = os.Create(fmt.Sprintf("output/%v", strings.ReplaceAll(itemName, ".png", ".qoi")))
			if err != nil {
				fmt.Printf("An error occurred creating new file: %v, err", imgQoi.Name())
				return
			}
			err = qoi.Encode(imgQoi, imgPng)
			if err != nil {
				fmt.Printf("An error occurred writing %v, %v", imgQoi.Name(), err)
				return
			}
		} else {
			fmt.Println(fmt.Sprintf("File %v is not a png format file", itemName))
		}
	}

}
