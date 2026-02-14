

package main
import(
	"fmt"
	//"net/http"
	//"windows/sendTools"
	"windows/tunnel"
	
)


func main(){
	fmt.Println("ASJDLAJSLDJALKSDJLAKJSDLKASJDLKJALSDAKDLA")

	
	re := tunnel.InitTunne()
	fmt.Println(re)

	// resp := http.FileServer(http.Dir("/"))

	// http.Handle("/", resp)

	// fmt.Println("iNICIANDO SERVER")
	// err := http.ListenAndServe(":3005", nil)
	// if(err != nil){
	// 	fmt.Println(err)
	// }
	

}