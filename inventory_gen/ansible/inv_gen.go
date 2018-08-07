package ansible

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"path"
	"os"
)

type StateReader interface {
	read() State
}

type FileBased struct {
	inputSource string
}

func (fb *FileBased) read() State {

	var fObj State
	b, e := ioutil.ReadFile(fb.inputSource)
	if e != nil {
		log.Println(e)
		return fObj
	}

	if err := json.Unmarshal(b, &fObj); err != nil {
		log.Println("invald json file")
		return fObj
	}

	return fObj
}

func genFile(groups map[string][]Instance) {
	var buffer bytes.Buffer
	for groupName, instances := range groups {
		genHeader(&buffer, groupName)
		for _, instance := range instances {
			writeSSHKeyToFile(instance.key)
		}
	}
	fmt.Println(buffer.String())
}

func genHeader(buffer *bytes.Buffer, content string) {
	buffer.WriteString("[")
	buffer.WriteString(content)
	buffer.WriteString("]\n")
}

func writeSSHKeyToFile(key KeyPair) {
	parentDir := "/tmp"
	filePath := path.Join(parentDir, fmt.Sprintf("%s.%s", key.name, "pem"))

	if _, err := os.Stat(filePath);!os.IsNotExist(err) {
		log.Println("Key already exists on path")
		return
	}

	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer  f.Close()

	f.WriteString(key.privateKey)

}