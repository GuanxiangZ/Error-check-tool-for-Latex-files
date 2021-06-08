package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	const str = "\\entry{Here is the entry}"
	start := strings.Index(str, "Here is the entry")
	end := strings.LastIndex(str, "Here is the entry")
	//index := latexRe.FindStringIndex(str)
	fmt.Println(str[start: end])
}

func TestSection(t *testing.T) {
	str := ".*\\\\section\\s*\\{([\\S\\s]+)*\\}"
	re := regexp.MustCompile(str)

	teststr := "\\section{asd asdasd}asdasd"
	rst := re.FindAllStringSubmatchIndex(teststr, -1)
	fmt.Println(rst)
	fmt.Println(re.FindStringSubmatch(teststr)[1])
}

func TestFile(t *testing.T) {
	pattern := ".*\\\\begin{document}[\\s\\S]*\\\\end{document}"
	re := regexp.MustCompile(pattern)
	wd, _ := os.Getwd()


	content, _ := ioutil.ReadFile(wd + "/static/demo.tex")
	indexes := re.FindAllSubmatchIndex(content, -1)
	fmt.Println(indexes)
	fmt.Println(string(content[3159: 24235]))
}
