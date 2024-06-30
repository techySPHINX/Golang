package main

import{
	"fmt","log","net/http"

}

func formHandler(w http.ResponseWriter, r "http.Request"){
	if err := r.parseForm(); err != nil{
		fmt.printf(w, "parseForm() err: %v", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, r "http.Request"){
	if r.URL.Path  != '/hello'{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil){
		log.Fatal(err)
	}
}