package tunnel

import (
	//"bytes"
	//"container/list"
	"fmt"
	"io"
	
	"os/exec"
	
)


func InitTunne()bool{
	//var url string = "ssh -p 443 -R0:127.0.0.1:3333 free.pinggy.io > info.txt"
	//var out bytes.Buffer
	//var port string = ":3333"

	//e := strings.Split(url, " ")
	

	fmt.Println("Asdd")
	cmd := exec.Command("ssh", "-p", "443", "-R0:127.0.0.1:3333", "free.pinggy.io")
	//cmd.Stdout = &out
	stout, err := cmd.StdoutPipe()
	if(err != nil){fmt.Println(err)}


	//----
	//stdin, err := cmd.StdinPipe()
	//if(err != nil){fmt.Println("a")}
	//defer stdin.Close()
	//io.WriteString(stdin, "yes")
	//io.WriteString(stdin, "yes")

	cmd.Start()

	siu, _ := io.ReadAll(stout)
	fmt.Println(string(siu))
	cmd.Wait()


	
	

	return true
}

