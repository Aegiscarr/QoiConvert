package main

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"log"
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
	imgIn       image.Image
	imgOut      *os.File
)

func main() {
	assetFolder, err := os.ReadDir("./input")
	if err != nil {
		fmt.Println("An error occurred reading AssetFolder: ", err)
		return
	}

	for i := range assetFolder {
		itemName := assetFolder[i].Name()
		fmt.Println(fmt.Sprintf("converting '/input/%v'", itemName))
		if strings.Contains(itemName, ".qoi") {

			openedFile, err = os.Open(fmt.Sprintf("./input/%v", itemName))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred opening %v: %v", itemName, err))
				return
			}

			imgIn, err = qoi.Decode(openedFile)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred decoding %v: %v", itemName, err))
				return
			}

			imgOut, err = os.Create(fmt.Sprintf("./output/%v", strings.ReplaceAll(itemName, ".png", ".qoi")))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred creating new file: %v, err", imgOut.Name()))
				return
			}

			err = png.Encode(imgOut, imgIn)
			if err != nil {
				fmt.Println(fmt.Sprintf("An error occurred writing %v, %v", imgOut.Name(), err))
				return
			}
			fmt.Println(fmt.Sprintf("successfully converted './input/%v' to '%v'", itemName, imgOut.Name()))
		} else {
			fmt.Println(fmt.Sprintf("File %v is not a known format file", itemName))
		}
	}

}
