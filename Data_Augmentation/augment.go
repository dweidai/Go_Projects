package main


import (
  "fmt"
  "image"
  "os"
  "path/filepath"
)

type Pixel struct {
    R int
    G int
    B int
    A int
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
    return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func getFilesFromDir(dirPath string)([]string){
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
   
    return files[1:]
}

func getImageFromFilePath(filePath string) ([][]Pixel, error) {
    f, err := os.Open(filePath)
	if err != nil {
    	return nil, err
	}
	img, _, err := image.Decode(f)

    if err != nil {
        return nil, err
    }

    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    var pixels [][]Pixel
    for y := 0; y < height; y++ {
        var row []Pixel
        for x := 0; x < width; x++ {
            row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
        }
        pixels = append(pixels, row)
    }

    return pixels, nil
}


func main(){
	dataFiles := getFilesFromDir("/Users/dwei/Desktop/chest_xray/val/NORMAL")
	 for _, file := range dataFiles {
    	fmt.Println(file )
    }
	pixels, err := getImageFromFilePath(dataFiles[0])
	if err != nil{
		print(err)
	} else {
		fmt.Println(pixels)
	}
}