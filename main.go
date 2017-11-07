package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"

)

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}


func Eliza(w http.ResponseWriter, r *http.Request){
	
	userGuess:= r.URL.Query().Get("value")
	check:=false
	
	response := ""
	if (userGuess == "Hello"){
		//fmt.Fprintf(w, "Hello, %s!", userGuess)
		response = "Hello Keith, how are you?"
		
	}else{
		response = "Try saying Hello!"
		
	}
	if(strings.Contains(userGuess,"bad")){
		response = "Be happy"
		
	}else if(strings.Contains(userGuess,"good")){
		response = "Maaa G, Stress Less I love it, Lets play a game hombre , Enter a string Il preform some magic on it"
		check = true
		
	}
	
	fmt.Fprintf(w,"%s",response)
	
	if(check == true){
		 response=Reverse(userGuess)
		 fmt.Fprintf(w,"%s",response)
	}
	
	
	
	
}

func main() {
	
	//Handles static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/Eliza", Eliza)
    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}