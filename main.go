//Note: Use capital letter to export whet using <`json:"attribute"`>
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type user struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Descrip string `json:"descrip"`
}

//array of users
type users struct {
	UsersArray []user `json:"users"`
}

func main() {
	var file, err = os.Open("./users.json")
	defer file.Close()

	if err != nil {
		log.Fatal("File not found")
	}
	readJSON(file)
}

func readJSON(file *os.File) {
	//var of array
	var usersVar users
	var b, errC = ioutil.ReadAll(file)
	var opc int
	var ban bool
	if errC != nil {
		log.Fatal(errC)
	}

	//pass values ​​from b to usersVar
	json.Unmarshal(b, &usersVar)

	for ban == false {
		fmt.Println("1. View list")
		fmt.Println("2. Search user")
		fmt.Println("3.Exit")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			for i := 0; i < len(usersVar.UsersArray); i++ {
				viewUser(i, usersVar)
			}
		case 2:
			var id int
			fmt.Println("Enter ID")
			fmt.Scan(&id)
			var userSearch = searchUser(id, usersVar)
			if userSearch < 0 {
				fmt.Println("User not found")
			} else {
				viewUser(userSearch, usersVar)
			}
		case 3:
			fmt.Println("E X I T . . .")
			ban = true
		default:
			fmt.Println("Invalid option")
		}
	}
}

func searchUser(id int, userVar users) int {
	for i, v := range userVar.UsersArray {
		if id == v.ID {
			return i
		}
	}
	return -1
}

func viewUser(i int, usersVar users) {
	fmt.Println("User id: ", usersVar.UsersArray[i].ID)
	fmt.Println("User name: ", usersVar.UsersArray[i].User)
	fmt.Println("User pass: ", usersVar.UsersArray[i].Pass)
	fmt.Println("User description: ", usersVar.UsersArray[i].Descrip)
}
