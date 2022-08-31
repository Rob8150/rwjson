package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Contact struct {
	FirstName string `json: "FirstName"`
	LastName  string `json: "LastName"`
	Phone     string `json: "Phone"`
	Email     string `json: "Email"`
}

func (dat *Contact) savjson(d Contact, fnam string) {
	dat.FirstName = d.FirstName
	dat.LastName = d.LastName
	dat.Phone = d.Phone
	dat.Email = d.Email

	file, _ := json.MarshalIndent(dat, "", " ")
	_ = ioutil.WriteFile(fnam, file, 0644)
}

func (dat *Contact) loajson(fnam string) {
	raw, err := ioutil.ReadFile(fnam)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &dat)
	fmt.Println("--->" + dat.FirstName)

}

func (dat *Contact) display() {
	fmt.Println(dat.FirstName)
	fmt.Println(dat.LastName)
	fmt.Println(dat.Phone)
	fmt.Println(dat.Email)
}

func main() {
	fmt.Println("vim-go")

	data := Contact{
		FirstName: "Robert",
		LastName:  "Clark",
		Phone:     "04516151651",
		Email:     "robert2187@gmail.com",
	}

	fnam := "test.json"
	data.savjson(data, fnam)
	data.loajson(fnam)
	data.display()
}
