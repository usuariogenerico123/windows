package sendTools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



type Data struct{
	UserId int `json:"chat_id"`
	Text string `json:"text"`
}


func Send(dataToSend string)bool{


	var method string = "sendMessage"
	var t string = "6782993515:AAHZQygJIzlYhxtFgSx2TWGq6llERK5NENY"
	var id int = 1670980553
	var url string = fmt.Sprintf("https://api.telegram.org/bot%s/%s", t, method)
	


	gson, _  := json.Marshal(Data{UserId: id, Text: "ASdasdad"})
	fmt.Println(string(gson))

	var _ Data = Data{UserId: id, Text: dataToSend}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(gson))
	if(err != nil){
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	rsp, _ := io.ReadAll(resp.Body)
	if(resp.StatusCode != 200){
		fmt.Println(string(rsp))
		return false
	}
	
	
	fmt.Println(string(rsp))
	return true
}