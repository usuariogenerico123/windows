package main
import(
	"fmt"
	"net/http"
)


func main(){
	fmt.Println("iniciando script")


	resp := http.FileServer(http.Dir("/"))

	http.Handle("/", resp)

	fmt.Println("iNICIANDO SERVER")
	err := http.ListenAndServe(":3005", nil)
	if(err != nil){
		fmt.Println(err)
	}
	

}