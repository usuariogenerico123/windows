package main
import(
	"fmt"
	"net/http"
	//"windows/sendTools"
	"windows/tunnel"
	"os"
)


func main(){
	fmt.Println("iniciando script")

	res, _ := os.ReadFile("info.log")
	fmt.Println(string(res))
	re := tunnel.InitTunne()
	fmt.Println(re)

	resp := http.FileServer(http.Dir("/"))

	http.Handle("/", resp)

	fmt.Println("iNICIANDO SERVER")
	err := http.ListenAndServe(":3005", nil)
	if(err != nil){
		fmt.Println(err)
	}
	

}