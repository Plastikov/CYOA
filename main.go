// no file yet

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"github.com/plastikov/cyoa/story"
)
func main(){
	file := flag.String("file", "gopher.json", "The CYOA file")
	flag.Parse()
	fmt.Printf("Opening the CYOA file %s", *file)

	f, err := os.Open(*file)
	if err != nil{
		fmt.Printf("%s does not exist or cannot be found", err)
	}
	d := json.NewDecoder(f)
	var story story.Chapter
	if err := d.Decode(&story); err != nil {
		fmt.Printf("unable to read %s", err)
	}
	fmt.Printf("%v+\n", story)
}
