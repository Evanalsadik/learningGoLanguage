package main

import "fmt"

type Item struct {
	id    int
	name  string
	price int
}

type ItemNotFoundError struct {
	Id int
}

var items = []Item{
	{1, "jeruk", 13400},
	{2, "anggur", 1400},
	{3, "nanas", 3400},
}

func main() {
	result, err := getItemById(0)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("result:", result)
}

func getItemById(id int) (Item, error) {
	for _, item := range items {
		if item.id == id {
			return item, nil
		}
	}
	return Item{}, &ItemNotFoundError{Id: id}
}

func (i *ItemNotFoundError) Error() string {
	return fmt.Sprintf("Item dengan id %d tidak ditemukan.", i.Id)
}
