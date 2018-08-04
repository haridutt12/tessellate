package ansible

import (
	"fmt"
	"log"
		"io/ioutil"
	"encoding/json"
	"bytes"
)

type StateReader interface {
	read() State
}

type FileBased struct {
	inputSource string
}

func (fb *FileBased) read() State {

	var fObj State
	bytes, e := ioutil.ReadFile(fb.inputSource)
	if e != nil {
		log.Println(e)
		return fObj
	}

	if err := json.Unmarshal(bytes, &fObj); err != nil {
		fmt.Println(err)
		log.Printf("invald json file")
		return fObj
	}

	return fObj
}

func genFile(groups map[string][]Instance) {
	var buffer bytes.Buffer
	for groupName, _ := range groups {
		genHeader(&buffer, groupName)
	}
	fmt.Println(buffer.String())
}

func genHeader(buffer *bytes.Buffer, content string) {
	buffer.WriteString("[")
	buffer.WriteString(content)
	buffer.WriteString("]\n")
}