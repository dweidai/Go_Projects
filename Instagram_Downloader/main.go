package main


import(
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
	"net/http"
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
	fmt.Println("You want to download using instagram url or username (url/user)")
	fmt.Scan(&temp)
	if strings.ToLower(temp) == "url"{
		fmt.Println("Please enter an instagram url")
		fmt.Scan(&url)
	} else if strings.ToLower(temp) == "user"{
		fmt.Println("Please enter the instagram username")
		fmt.Scan(&username)
	}
	fmt.Println("Please enter the output path or \"skip\" then the default is current directory")
	fmt.Scan(&outputPath)
	if outputPath == "skip"{
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
	if url == "" && username == ""{
		fmt.Println("The url or the username for Instagram is required")
		flag.Usage()
		os.Exit(1)
	}
}

func newRequest(igusername string, igurl string) (ig_struct, error){
	if igusername != "" && igurl == ""{
		url = fmt.Sprintf("https://www.instagram.com/%s", igusername)
	} else if igusername == "" && igurl != ""{
		var err error
		temp := "https://www.instagram.com/"
		length := len(temp)
		urlCount := len(igurl)
		if urlCount <= length {
			err = errors.New("URL is not long enough to determine username")
		}
		want := strings.SplitAfterN(url, "/", 4)
		if want[0] == "https:/" && want[1] == "/" && want[2] == "www.instagram.com/"{
			username = strings.Trim(want[3], "/")
		} else {
			err = errors.New("Username not found")
		}
		if err != nil{
			return ig_struct{}, err
		}
	}
	req, err := http.Get(url)
	if err != nil{
		return ig_struct{}, err
	}
	content := *req
	if content.Status != "200 OK"{
		fmt.Println(url)
		fmt.Println(content.Status)
		err = errors.New("Username not found or invalid")
	}
	defer req.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&ig); err != nil {
		return ig_struct{}, fmt.Errorf("JSON: %s\n", err.Error())
	}
	return err
}
func main(){
	init_Info()
	ig, err := newRequest(username, url)
	if err != nil{
		fmt.Printf("Fetch: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(ig)
	
}