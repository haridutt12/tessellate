package ansible

import (
	"strings"
	"strconv"
	"log"
)

/*
This file actually helps to parse the state file which is in JSON format
*/
const SEPERATOR  = "."

type KeyPair struct {
	privateKey string
	name string
}

/*
Consider version and modules present in the state file
*/
type State struct {
	Version int `json:"version"`
	Modules []Module `json:"modules"`
}

func (st *State) getModuleLen() int {
	return len(st.Modules)
}

func (st *State) getModuleFromIndex(index int) Module {
	if (index > st.getModuleLen() - 1) {
		panic("The passed index to get module is greater than total length")
	}
	return st.Modules[index]
}

/*
Find all the resources for a module
*/
type Module struct {
	Resources map[string]Resource `json:"resources"`
}


/*
Resource which has type, has a depends on array and primary resource
*/
type Resource struct {
	Type string `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary PrimaryResource `json:"primary"`
}

/*
Based on the key name get the
*/
func (re *Resource) getNameFromKey(key string) string {
	return strings.Split(key, SEPERATOR)[1]
}

func (re *Resource) getIndexFromKey(key string) int {
	name := strings.Split(key, SEPERATOR)
	index, e := strconv.Atoi(name[len(name)-1])
	if e != nil {
		log.Printf("Only one instance for resource %s", key)
		return 0
	}
	return index
}

type PrimaryResource struct {
	Id string `json:"id"`
	Attributes Attribute `json:"attributes"`
}

type Attribute struct {
	PrivateIp string `json:"private_ip"`
	PublicIp  string `json:"public_ip"`
	PrivateKeyPem string `json:"private_key_pem"`
}

func (at Attribute) getPrivateIp() string {
	return at.PrivateIp
}

func (at Attribute) getPublicIp() string {
	return at.PublicIp
}

func (at Attribute) getPrivateKey() string {
	return at.PrivateKeyPem
}
