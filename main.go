// no file yet

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/plastikov/cyoa/story"
)
func main(){
	file := flag.String("file", "story.json", "The CYOA file")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil{
		fmt.Printf("%s does not exist or cannot be found", err)
	}
	story, err := story.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n\n", story)
}	