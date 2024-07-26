package main

import (
	"GoEdgeKeyGen/third/encode"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const (
	id = "nameless"
	dayFrom = "1949-10-01"
	dayTo = "2077-06-04"
	requestCode = ""
	company = "GoEdge | 分遗产群出品"
	nodes = 0
	updatedAt = 1700000000
	edition = "ultra"
	email = ""
)
func main() {
	fmt.Println("GoEdge License Keygen By @GoEdge | 分遗产群")
	payload := KeyGen{
		Id: id,
		DayFrom: dayFrom,
		DayTo: dayTo,
		MacAddresses: []string{},
		ReqCode: requestCode,
		Hostname: "*",
		Company: company,
		Nodes: nodes,
		UpdatedAt: updatedAt,
		Components: []string{"*"},
		Edition: edition,
		Email: email,
	}
	key, _ := json.Marshal(payload)
	license := encode.Encode(key)
	licenseStr := base64.StdEncoding.EncodeToString(license)
	fmt.Println("↓已产出注册码↓")
	fmt.Println(licenseStr)
	fmt.Scanln()
}


type KeyGen struct{
	Id string `json:"id"`
	DayFrom string `json:"dayFrom"`
	DayTo string `json:"dayTo"`
	MacAddresses []string `json:"macAddresses"`
	ReqCode string `json:"requestCode"`
	Hostname string `json:"hostname"`
	Company string `json:"company"`
	Nodes int64 `json:"nodes"`
	UpdatedAt int64 `json:"updatedAt"`
	Components []string `json:"components"`
	Edition string `json:"edition"`
	Email string `json:"email"`
}