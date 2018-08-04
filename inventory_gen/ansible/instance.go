package ansible

import "strings"

/*
Identify one instance which has keyPair
name of the instance based on the key,
index if there are more than 1 instance of similar type
public and private Ip allocated
*/
type Instance struct {
	key 	   KeyPair
	name       string
	index      int
	privateIp  string
	publicIp   string
}

func (in *Instance) setPrivateIp(privateIp string) {
	in.privateIp = privateIp
}

func (ins *Instance) setPublicIp(publicIp string) {
	ins.publicIp = publicIp
}

func (ins *Instance) setName(name string) {
	ins.name = name
}

func (ins *Instance) setIndex(index int) {
	ins.index = index
}

/*
This function actually looks for the attached private key for a given instance
*/
func (ins *Instance) lookupKey(provider Provider, instanceResource Resource, module Module) {

	for _, publicKeyResourceName := range instanceResource.DependsOn {
		if strings.Contains(publicKeyResourceName, provider.publicKeyIdentifier) {


			keyPair := module.Resources[publicKeyResourceName]
			for _, privateKeyResourceName := range keyPair.DependsOn {
				if strings.Contains(privateKeyResourceName, provider.privateKeyIdentifier) {

					key := KeyPair{}

					key.privateKey = module.Resources[privateKeyResourceName].Primary.Attributes.getPrivateKey()
					key.name = instanceResource.getNameFromKey(publicKeyResourceName)
					ins.key = key
				}
			}
		}
	}
}
