package main


import(
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)
	
var (
	url                  string
	username   	         string
	outputPath           string
	downloadCount        int
	takeImages     	     bool
	takeVideos        	 bool
)

func init_Info(){
	flag.StringVar(&url, "url", "", "Ex: https://www.instagram.com/arsenal/")
	flag.StringVar(&username, "username", "", "Ex: cindymello")
	flag.StringVar(&outputPath, "o", "./", "Ex: ~/Downloads or ../pictures")
	flag.IntVar(&downloadCount, "count", 0, "Ex: Download the first n media or the deault is all")
	flag.BoolVar(&takeImages, "images", false, "Grab images, please input y/n yes/no")
	flag.BoolVar(&takeVideos, "videos", false, "Grab videos, please input y/n yes/no")
	flag.Parse()
	var temp string
	fmt.Println("Please enter an instagram url")
	fmt.Scan(&url)
	fmt.Println("Please enter the instagram username")
	fmt.Scan(&username)
	fmt.Println("Please enter the output path or empty then the default is current directory")
	fmt.Scan(&outputPath)
	if outputPath == ""{
		outputPath = "./"
	}
	fmt.Println("Please enter the number of contents your want to download (enter \"all\" to download all)")
	fmt.Scan(&temp)
	if strings.ToLower(temp) == "all"{
		downloadCount = -1
	} else {
		i, err := strconv.Atoi(temp)
		if err != nil{
			fmt.Println(err)
			flag.Usage()
			os.Exit(1)
		}
		downloadCount = i
	}
	fmt.Println("Please enter whether your want to download images (y/n)")
	fmt.Scan(&temp)
	temp = strings.ToLower(temp)
	if temp == "y" || temp == "yes"{
		takeImages = true
	} else if temp == "n" || temp == "no"{
		takeImages = false
	}else{
		takeImages = true
	}
	fmt.Println("Please enter whether your want to download images (y/n)")
	fmt.Scan(&temp)
	temp = strings.ToLower(temp)
	if temp == "y" || temp == "yes"{
		takeVideos = true
	} else if temp == "n" || temp == "no"{
		takeVideos = false
	} else{
		takeVideos = true
	}
	if url == "" || username == ""{
		fmt.Println("The url and the username for Instagram is required")
		flag.Usage()
		os.Exit(1)
	}
}

func main(){
	init_Info()
	fmt.Println(url)
	fmt.Println(username)
	fmt.Println(outputPath)
	fmt.Println(downloadCount)
	fmt.Println(takeImages)
	fmt.Println(takeVideos)
}