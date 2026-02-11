package src

import (
	"fmt"
	"io"
	"net/http"
)



type Data struct{
	UserId int
	Text string
}


func Send(dataToSend string)bool{


	var method string = ""
	var t string = "6782993515:AAHZQygJIzlYhxtFgSx2TWGq6llERK5NENY"
	var id int = 1670980553
	var url string = fmt.Sprintf("https://api.telegram.org/bot$s/%s", t, method)
	

	var _ Data = Data{UserId: id, Text: dataToSend}

	resp, err := http.Get(url)
	if(err != nil){
		fmt.Println(err)

	}
	defer resp.Body.Close()
	rsp := io.ByteReader(resp.Body.Read())

	fmt.Println(rsp)


}