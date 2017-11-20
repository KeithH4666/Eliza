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
	
	re := regexp.MustCompile(`(?i)[I|i] [love|like](.*)`)
	if matched, _:= regexp.MatchString(`(?i).*\bHow|Why\b.*`, userGuess);matched{
		answers1 := []string{`I like to keep myself and my answers to myself`,`Why would you wonder such a thing?`,`I'm a chat bot not wikipedia`}
		randindex1 := rand.Intn(len(answers1))
		fmt.Fprintf(w,answers1[randindex1])
		return
	}
	
	if matched, _:= regexp.MatchString(`(?i).*am.*`, userGuess);matched{
		answers3 := []string{`I'm not sure what that feels like, as I am a bot.`,`I have always wondered what human emotions are like.`,`I am unfamiliar with that feeling`}
		randindex3 := rand.Intn(len(answers3))
		fmt.Fprintf(w,answers3[randindex3])
		return
	}
	
	if matched, _:= regexp.MatchString(`(?i).*Hello|Hey|Hi`, userGuess);matched{
		answers4 := []string{`Greetings user, nice to chat!`,`Hello human! My name is Eliza!`,`Hello , my names Eliza! Lets chat!`}
		randindex4 := rand.Intn(len(answers4))
		fmt.Fprintf(w,answers4[randindex4])
		return
	}
	
	if matched, _:= regexp.MatchString(`(?i).*\bfather\b.*`, userGuess);matched{
		fmt.Fprintf(w,"I am your father")
		return
	}
	
	if matched, _:= regexp.MatchString(`(?i).*hate|dont'|dislike.*`, userGuess);matched{
		answers5 := []string{`Don't be so negative!`,`Always look on the bright side of life!`,`I'v learned as a bot to appreciate everything.`}
		randindex5 := rand.Intn(len(answers5))
		fmt.Fprintf(w,answers5[randindex5])
		return
	}

	
	answers := []string{`Why do %s?`, `How long have %s?`,`Why %s?`}
	
	
	likeResponse := Response{re, answers} // Responses could be read from files.

	// we want to swap certain words so Eliza is actually talking about "you" and not just mimicing what was said, "I",
	// a lot more useful swaps can be found here https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

	reflectionMap := map[string]string{
		"I": "you",
		"i": "you",
		"was": "were",
		"are": "am",
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
		responses :=[]string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
		}
		randindex := rand.Intn(len(responses))
		fmt.Fprintf(w,responses[randindex]) // there was no regex match so just say some default answer.
	}
	
}

func main() {
	
	//Seeds for true random
	rand.Seed(time.Now().UnixNano())
	
	//Handles static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/Eliza", Eliza)
	
    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}