package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Yusuke Mabuchi",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Meiko Mabuchi",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Ryosuke Mabuchi",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	err = ioutil.WriteFile("./src/chapter3/parse_json/post2.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
	//jsonFile, err := os.Open("./src/chapter3/parse_json/post.json")
	//if err != nil {
	//	fmt.Println("Error opening JSON file:", err)
	//	return
	//}
	//defer jsonFile.Close()
	//
	//decoder := json.NewDecoder(jsonFile)
	//for {
	//	var post Post
	//	err := decoder.Decode(&post)
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Error decoding JSON:", err)
	//		return
	//	}
	//	fmt.Println(post)
	//}

	//jsonData, err := ioutil.ReadAll(jsonFile)
	//if err != nil {
	//	fmt.Println("Error reading JSON data:", err)
	//	return
	//}
	//var post Post
	//json.Unmarshal(jsonData, &post)
	//fmt.Println(post)
}
