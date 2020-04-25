package main

import(
	"fmt"
	//"os/exec"
	"reflect"
	"net/http"
	//"io/ioutil"
)

func main(){
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
	fmt.Println(url)
	req, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(*req).Kind())
	content := *req
	fmt.Println(content.Status)
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