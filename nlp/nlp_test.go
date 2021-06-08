package nlp

import (
	"encoding/json"
	"fmt"
	"testing"
)


func TestNLP_Check(t *testing.T) {
	ins := NewDefaultService()

	check, _ := ins.Check("\\caption{the caption of single sentence does not have period at the end}")

	rst, _ := json.MarshalIndent(check, "", "	")
	fmt.Println(string(rst))

	check, _ = ins.Check("she love dog")

	rst, _ = json.MarshalIndent(check, "", "	")
	fmt.Println(string(rst))
}
