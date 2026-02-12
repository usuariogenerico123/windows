package tunnel

import (
	"bytes"
	//"container/list"
	"fmt"
	//"io"
	
	"os/exec"
	"os"
)


func InitTunne()bool{
	//var url string = "ssh -p 443 -R0:127.0.0.1:3333 free.pinggy.io > info.txt"
	var out bytes.Buffer
	//var port string = ":3333"

	//e := strings.Split(url, " ")
	

	fmt.Println("Asdd")
	cmd := exec.Command("ssh", "-p", "443", "-R0:127.0.0.1:3333", "free.pinggy.io")
	cmd.Stdout = &out
	//stout, err := cmd.StdoutPipe()
	//if(err != nil){fmt.Println(err)}



	cmd.Run()


	//fmt.Println(string(out.String()))
	


	os.WriteFile("siu.log", []byte(string(out.String())), 0666)
	
	

	return true
}

