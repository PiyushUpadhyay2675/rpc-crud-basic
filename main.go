package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body string
}

type API int

var database []Item

func (a *API) GetByName(Title string, reply *Item) error {
	var getItem Item

	for _,val := range database {
		if val.Title == Title {
			getItem = val
			break
		}
	}
	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = edit
			changed = edit
			break
		}
	}

	*reply = changed

	return nil
} 


func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[(idx+1):]...)
			del = val
			break
		}
	}
	*reply = del
	return nil
}

func (a *API) GetDB(Title string, reply *[]Item) error {
	*reply = database
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Error Registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listening Error:", err)
	}

	log.Printf("Serving RPC on port %d", 4040)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error Serving:", err)
	}

	// fmt.Println("Initial Database:", database)
	// a := Item{"first", "a first item"}
	// b:= Item{"second", "a second item"}
	// c := Item{"third", "a third item"}
	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("Second Database:", database)

	// DeleteItem(b)
	// fmt.Println("Third Database:", database)

	// EditItem("third", Item{"fourth", "a fourth item which replaced third"})
	// fmt.Println("Fourth Databse:", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)
}