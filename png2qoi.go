package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/xfmoulet/qoi"
	"golang.org/x/image/webp"
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
		if strings.Contains(itemName, ".png") {

			openedFile, err = os.Open(fmt.Sprintf("./input/%v", itemName))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred opening %v: %v", itemName, err))
				return
			}

			imgIn, err = png.Decode(openedFile)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred decoding %v: %v", itemName, err))
				return
			}

			imgOut, err = os.Create(fmt.Sprintf("./output/%v", strings.ReplaceAll(itemName, ".png", ".qoi")))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred creating new file: %v, err", imgOut.Name()))
				return
			}

			err = qoi.Encode(imgOut, imgIn)
			if err != nil {
				fmt.Println(fmt.Sprintf("An error occurred writing %v, %v", imgOut.Name(), err))
				return
			}
			fmt.Println(fmt.Sprintf("successfully converted './input/%v' to '%v'", itemName, imgOut.Name()))
		} else if strings.Contains(itemName, ".jpg") {
			openedFile, err = os.Open(fmt.Sprintf("./input/%v", itemName))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred opening %v: %v", itemName, err))
				return
			}

			imgIn, err = jpeg.Decode(openedFile)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred decoding %v: %v", itemName, err))
				return
			}

			imgOut, err = os.Create(fmt.Sprintf("./output/%v", strings.ReplaceAll(itemName, ".jpg", ".qoi")))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred creating new file: %v, err", imgOut.Name()))
				return
			}

			err = qoi.Encode(imgOut, imgIn)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred writing %v, %v", imgOut.Name(), err))
				return
			}
			fmt.Println(fmt.Sprintf("successfully converted './input/%v' to '%v'", itemName, imgOut.Name()))
		} else if strings.Contains(itemName, ".jpeg") {
			openedFile, err = os.Open(fmt.Sprintf("./input/%v", itemName))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred opening %v: %v", itemName, err))
				return
			}

			imgIn, err = jpeg.Decode(openedFile)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred decoding %v: %v", itemName, err))
				return
			}

			imgOut, err = os.Create(fmt.Sprintf("./output/%v", strings.ReplaceAll(itemName, ".jpeg", ".qoi")))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred creating new file: %v, err", imgOut.Name()))
				return
			}

			err = qoi.Encode(imgOut, imgIn)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred writing %v, %v", imgOut.Name(), err))
				return
			}
			fmt.Println(fmt.Sprintf("successfully converted './input/%v' to '%v'", itemName, imgOut.Name()))
		} else if strings.Contains(itemName, ".webp") {
			openedFile, err = os.Open(fmt.Sprintf("./input/%v", itemName))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred opening %v: %v", itemName, err))
				return
			}

			imgIn, err = webp.Decode(openedFile)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred decoding %v: %v", itemName, err))
				return
			}

			imgOut, err = os.Create(fmt.Sprintf("./output/%v", strings.ReplaceAll(itemName, ".webp", ".qoi")))
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred creating new file: %v, err", imgOut.Name()))
				return
			}

			err = qoi.Encode(imgOut, imgIn)
			if err != nil {
				log.Fatalln(fmt.Sprintf("An error occurred writing %v, %v", imgOut.Name(), err))
				return
			}
			fmt.Println(fmt.Sprintf("successfully converted './input/%v' to '%v'", itemName, imgOut.Name()))
		} else {
			fmt.Println(fmt.Sprintf("File %v is not a known format file", itemName))
		}
	}

}
