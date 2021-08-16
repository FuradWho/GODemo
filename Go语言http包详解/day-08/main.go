package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)



func parsingXML()  {

	type server struct {
		XMLName    xml.Name `xml:"server"`
		ServerName string   `xml:"serverName"`
		ServerIP   string   `xml:"serverIP"`
	}

	type Recurlyservers struct {
		XMLName     xml.Name `xml:"servers"`
		Version     string   `xml:"version,attr"`
		Svs         []server `xml:"server"`
		Description string   `xml:",innerxml"`
	}

	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func parsingJSON()  {
	type Server struct {
		ServerName string
		ServerIP   string
	}

	type Serverslice struct {
		Servers []Server
	}
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	
}

func structToJson()  {

	type Order struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Quantity int `json:"quantity"`
		TotalPrice float64 `json:"total_price"`
	}

	o := Order{
		"123",
		"go",
		3,
		30,
	}

	b,err := json.Marshal(o)
	if err != nil{
		panic(err)
	}

	fmt.Printf("%+v\n",o)

	fmt.Printf("%s\n",b)
}


func main() {
	structToJson()
}
