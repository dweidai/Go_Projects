package main

import (
	"fmt"
	pdf "github.com/unidoc/unipdf/model"
	"os"
)

func main () {
	if len(os.Args) < 4{
		fmt.Println("Incorrect usage\n")
		fmt.Println("go run merge.go output.pdf input1.pdf input2.pdf ...\n")
		fmt.Println("go run merge.go -- executive command")
		fmt.Println("output.pdf      -- path and name of the output pdf file")
		fmt.Println("input1.pdf      -- path and name of the first pdf file to be merged")
		fmt.Println("input2.pdf	     -- path and name of the second pdf file to be merged")
		fmt.Println("...		     -- as many input pdf file as possible")
		os.Exit(1)
	}
	
	output := ""
	inputFiles := []string{}
	for i := 0; i < len(os.Args); i++ {
		if i == 1{
			output = os.Args[i]
		} else if (i > 1){
			inputFiles = append(inputFiles, os.Args[i])
		}
	}

	// then need to check if those files really exist
	for i := 0; i < len(inputFiles); i++ {
		if _, err := os.Stat(inputFiles[i]); os.IsNotExist(err) {
 			fmt.Println(inputFiles[i] + "\tDoes not Exist")
 			os.Exit(1)
		}		
		
	}

	fmt.Println("Writing to " + output)
	mergePdf(output, inputFiles)
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
			fmt.Println(auth)
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