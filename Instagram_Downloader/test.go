package main

import(
	"fmt"
	//"os/exec"
	"strings"
	"reflect"
	"net/http"
	//"io/ioutil"
)

func test(){
	//curl := exec.Command("curl", "https://www.instagram.com/cindymello/")
	//output, e := curl.Output()
	//fmt.Println(reflect.TypeOf(output).Kind())
	//fmt.Println(curl)
	//fmt.Println(len(output))
	/*x := 11
	for i := x; i < len(output); i++ {
		fmt.Println(i)
		fmt.Println(output[i])
	}*/
	//fmt.Println(e)
	url := "https://www.instagram.com/cindymello/"
	//url := "https://www.instagram.com/cindymello/"
	fmt.Println(url)
	req, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(*req).Kind())
	content := *req
	fmt.Println(content.Status)
	if content.Status == "200 OK"{
		fmt.Println("yayyyy!")
	}
	want := strings.SplitAfterN(url, "/", 4)
	fmt.Println(want[3])
	//ioutil.WriteFile("./dat1", message, 0644)
	/*url = "https://www.instagram.com/cindymelloxx/"
	fmt.Println(url)
	req, err = http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(*req).Kind())

	fmt.Println(*req)*/
	//ioutil.WriteFile("./dat2", message, 0644)
	/*body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(body)
	*/
}