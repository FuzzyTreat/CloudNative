// 1 . DOCKERFILE
// 2. Buildspec som bygger DOCKERFILE
// 2.5           docker login git.systementor.se yacloud yacloud1
// 3.               PUSH IMAGE -> git.systementor.se

package main

import (
	"fmt"

	"github.com/cheynewallace/tabby"
)

func main() {
	fmt.Println("Hej hej kolla vad coolt")
	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT2")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.AddLine("Stefan Holmberg", "Developer", "Whatever2")
	t.Print()
}
