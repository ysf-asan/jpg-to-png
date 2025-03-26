package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	input_file := "your_file_location.jpg"
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Dosya açılamadı: ", err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Görsel decode edilemedi: ")
		return
	}
	output_file := strings.TrimSuffix(input_file, filepath.Ext(input_file)) + ".png"
	outfile, err := os.Create(output_file)
	if err != nil {
		fmt.Println("Dosya oluşturulamadı: ", err)
		return
	}
	defer outfile.Close()

	err = png.Encode(outfile, img)
	if err != nil {
		fmt.Println("PNG kaydedilemedi", err)
		return
	}
	fmt.Println("Dönüştürme Tamamlandı", output_file)

}
