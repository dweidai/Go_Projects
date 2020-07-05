package main

import (
	docconv "code.sajari.com/docconv"
	"fmt"
	gopdf "github.com/jung-kurt/gofpdf"
	unicommon "github.com/unidoc/unipdf/common"
	creator "github.com/unidoc/unipdf/creator"
	pdf "github.com/unidoc/unipdf/model"
	"io/ioutil"
	"os"
	"strings"
)

func usage(){

	fmt.Println("go run merge.go output.pdf input1.pdf input2.pdf ...")
	fmt.Println()
	fmt.Println("go run merge.go -- executive command")
	fmt.Println("output.pdf      -- path and name of the output pdf file")
	fmt.Println("input1.pdf      -- path and name of the first pdf file to be merged")
	fmt.Println("input2.pdf	     -- path and name of the second pdf file to be merged")
	fmt.Println("...		     -- as many input pdf file as possible")
	fmt.Println("If the input file is not pdf, as long as it is DOC, DOCX, PNG, JPEG, GIF, JPG, TXT, HTML, merge.go will convert them to pdf for you")
	os.Exit(1)
}
func main () {
	if os.Args[1] == "-h" || os.Args[1] == "-help"{
		usage()
	} else if len(os.Args) < 4{
		fmt.Println("Incorrect usage")
		fmt.Println()
		usage()
	}
	
	output := ""
	inputFiles := []string{}
	deleteFiles := []string{} //need to be done
	for i := 0; i < len(os.Args); i++ {
		if i == 1{
			output = os.Args[i]
		} else if i > 1 {
			inputFiles = append(inputFiles, os.Args[i])
		}
	}

	// then need to check if those files really exist
	for i := 0; i < len(inputFiles); i++ {
		split := strings.Split(inputFiles[i], ".")
		extension := split[len(split)-1]
		if _, err := os.Stat(inputFiles[i]); os.IsNotExist(err) {
 			fmt.Println(inputFiles[i] + "\tDoes not Exist")
 			os.Exit(1)
		} else if extension != "pdf" {
			fmt.Println("Converting non pdf file to pdf file")
			if extension == "txt" {
				inputFiles[i] = txtToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "docx" {
				inputFiles[i] = docxToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "doc" {
				inputFiles[i] = docToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "jpeg" {
				inputFiles[i] = imageToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "png" {
				inputFiles[i] = imageToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "jpg" {
				inputFiles[i] = imageToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "gif" {
				inputFiles[i] = imageToPDF(inputFiles[i], &deleteFiles)
			} else if extension == "html" {
				inputFiles[i] = htmlToPDF(inputFiles[i], &deleteFiles)
			}
		}
	}

	fmt.Println("Writing to " + output)
	mergePdf(output, inputFiles)
	for file := range deleteFiles{
		err := os.Remove(deleteFiles[file])
		if err !=nil{
			_ = os.Remove(output)
			fmt.Println("Error with deleting temporary conversion files")
		}
	}
	fmt.Println("Done")
}

func imageToPDF(input string, deleteFiles *[]string) string{
	split := strings.Split(input, ".")
	split[len(split)-1] = "pdf"
	toReturn := strings.Join(split[:],".")
	err := helperImagestopdf(input, toReturn)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	*deleteFiles = append(*deleteFiles, toReturn)
	return toReturn
}

func txtToPDF(input string, deleteFiles *[]string) string{
	// read the text file data
	txtStr, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	// create a text file but with pdf extension
	split := strings.Split(input, ".")
	split[len(split)-1] = "pdf"
	toReturn := strings.Join(split[:],".")
	newpdf := gopdf.New("P", "mm", "A4", "")
	newpdf.AddPage()
	// Font
	newpdf.SetFont("Times", "", 12)
	// Output text in a 6 cm width column
	newpdf.MultiCell(180, 5, string(txtStr), "", "", false)
	newpdf.Ln(-1)
	newpdf.OutputFileAndClose(toReturn)
	*deleteFiles = append(*deleteFiles, toReturn)
	return toReturn
}

func docxToPDF(input string, deleteFiles *[]string) string{
	f, err := os.Open(input)
	if err != nil{
		panic(err)
	}
	resp, _ , err := docconv.ConvertDocx(f)
	if err != nil {
		panic(err)
	}
	split := strings.Split(input, ".")
	split[len(split)-1] = "txt"
	toReturn := strings.Join(split[:],".")
	var file, error = os.Create(toReturn)
	if error != nil {
		panic(error)
	}
	defer file.Close()
	file.WriteString(resp)
	*deleteFiles = append(*deleteFiles, toReturn)
	return txtToPDF(toReturn, deleteFiles)
}

func docToPDF(input string, deleteFiles *[]string) string{
	f, err := os.Open(input)
	if err != nil{
		panic(err)
	}
	resp, _ , err := docconv.ConvertDoc(f)
	if err != nil {
		panic(err)
	}
	split := strings.Split(input, ".")
	split[len(split)-1] = "txt"
	toReturn := strings.Join(split[:],".")
	var file, error = os.Create(toReturn)
	if error != nil {
		panic(error)
	}
	defer file.Close()
	file.WriteString(resp)
	*deleteFiles = append(*deleteFiles, toReturn)
	return txtToPDF(toReturn, deleteFiles)
}

func htmlToPDF(input string, deleteFiles *[]string) string{
	f, err := os.Open(input)
	if err != nil{
		panic(err)
	}
	resp, _ , err := docconv.ConvertHTML(f, true)
	if err != nil {
		panic(err)
	}
	split := strings.Split(input, ".")
	split[len(split)-1] = "txt"
	toReturn := strings.Join(split[:],".")
	var file, error = os.Create(toReturn)
	if error != nil {
		panic(error)
	}
	defer file.Close()
	file.WriteString(resp)
	*deleteFiles = append(*deleteFiles, toReturn)
	return txtToPDF(toReturn, deleteFiles)
}

func mergePdf(output string, inputFiles []string) error{
	pdfWriter := pdf.NewPdfWriter()
	for _, inputPath := range inputFiles {
		f, err := os.Open(inputPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer f.Close()
		pdfReader, err := pdf.NewPdfReader(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		isEncrypted, err := pdfReader.IsEncrypted()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if isEncrypted {
			auth, err := pdfReader.Decrypt([]byte(""))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if !auth {
				error := "Cannot merge encrypted, password protected document"
				fmt.Println(error)
				os.Exit(1)
			}
		}

		numPages, err := pdfReader.GetNumPages()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < numPages; i++ {
			pageNum := i + 1

			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = pdfWriter.AddPage(page)
			if err != nil {
				fmt.Println(err)
					os.Exit(1)
			}
		}
	}

	fWrite, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func helperImagestopdf(inputPath string, outputPath string) error {
	c := creator.New()

	imgPath := inputPath
	unicommon.Log.Debug("Image: %s", imgPath)

	img, err := c.NewImageFromFile(imgPath)
	if err != nil {
		unicommon.Log.Debug("Error loading image: %v", err)
		return err
	}
	img.ScaleToWidth(612.0)

	height := 612.0 * img.Height() / img.Width()
	c.SetPageSize(creator.PageSize{612, height})
	c.NewPage()
	img.SetPos(0, 0)
	_ = c.Draw(img)
	err = c.WriteToFile(outputPath)
	return err
}