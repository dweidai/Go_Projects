package main

import(
	"fmt"
	//"os/exec"
	//"reflect"
	//"net/http"
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
	//req, _ := http.NewRequest("GET", url, nil)
	/*res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))*/

}