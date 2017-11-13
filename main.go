package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"regexp"
	"time"
	"math/rand"
	
	//"time"

)

type Response struct {
	pattern         *regexp.Regexp
	possibleAnswers []string
}

func substituteWords(answer string, reflectionMap map[string]string) string {
	allWords := strings.Split(answer, " ") // get slices of the words {"words", "in", "sentence"}
	for index, word := range allWords {
		if val, ok := reflectionMap[word]; ok {
			allWords[index] = val // substitite the value
		}
	}
	return strings.Join(allWords, " ") // join back into string "words in sentence"
}

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}


func Eliza(w http.ResponseWriter, r *http.Request){
	t := time.Now()
	

	rand.Seed(time.Now().UnixNano())
	userGuess := r.URL.Query().Get("value")
	
	re := regexp.MustCompile(`I [love|am|Like](.*)`)

	
	answers := []string{`Why do %s?`, `How long have %s?`, `How long have you %s?`,`Why do %s?`}
	
	
	likeResponse := Response{re, answers} // Responses could be read from files.

	// we want to swap certain words so Eliza is actually talking about "you" and not just mimicing what was said, "I",
	// a lot more useful swaps can be found here https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

	reflectionMap := map[string]string{
		"I": "you",
		"was": "were",
		"are": "am",
		"you": "I",
		"like":"liked",
		"love":"love",
		"time": t.Format("3:04PM"),
		// and plenty more
	}

	if likeResponse.pattern.MatchString(userGuess) {
		match := re.FindStringSubmatch(userGuess)             // the pattern matched, so let's look for all the text after I like. ie. everything in (.*) - (.*) means "anything", so this pattern will match any text the user entered.
		topic := match[0]                                     // 0 is the full string, 1 is first match.
		topic = substituteWords(topic, reflectionMap)         // instead of saying "Why do you like you" it maps to "Why do you like me" because of the reflection map.
		index := rand.Intn(len(likeResponse.possibleAnswers)) // index of random possible answer
		chosenResponse := likeResponse.possibleAnswers[index] // pick the answer
		finalResponse := fmt.Sprintf(chosenResponse, topic)   // sub the "topic" into the response - since the response has a %s for this purpose
		fmt.Fprintf(w,finalResponse)                            // we actually want to display this in the html page or template or something.

		// The output will be "How long have you liked me?" or "Why do you like me?"
	} else {
		fmt.Fprintf(w,"I don't know what you said.") // there was no regex match so just say some default answer.
	}
	
	/*
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
		response = "Maaa G, Stress Less I love it hombre"
	}
	if(userGuess=="lets play"){
		
		response = ""
		fmt.Fprintf(w,"Hahah see what i Did: %s",gameGuess)
	
	}
	
	fmt.Fprintf(w,"%s",response)
	*/
	
}

func main() {
	
	//Handles static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/Eliza", Eliza)
    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}