package main

import (
	"fmt";
	"log";
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request)  {

	if err:= r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParseForm() err:%v",err)
		return
	}
	fmt.Fprintf(w,"POST form Successful\n")

	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"Name :%v\n",name)
	fmt.Fprintf(w,"Address: %v\n",address)



	
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path!="/hello" {
		http.Error(w,"Not Found",http.StatusNotFound)
		return
		
	}
	if r.Method!="GET"{
		http.Error(w,"Method Not Supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello")



	
}

func main()  {

	fileServer :=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Printf("Server Starting At Port 8080\n")

	if err := http.ListenAndServe(":8080",nil); err!=nil {
		log.Fatal(err)
	}
	
}