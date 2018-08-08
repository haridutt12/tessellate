package ansible

/*
This structure takes the cloud provider name, resource identifiers for a given cloud
For example :- AWS identifies a VM as aws_instance, similarly DigitalOcean identifies as droplet
Publickey in AWS is identified as aws_key_pair
#TODO based on cloud provider passed as command line, get the Provider
*/
type Provider struct {
	cloudProvider string
	machineIdentifier string
	publicKeyIdentifier string
	privateKeyIdentifier string
}

/*
Get the state from a StateReader which could be a file or tessellate client querying tessellate server
*/
func (pro *Provider) getState(reader StateReader) State {
	return reader.read()
}

/*
Get the VMs with public/private IP, private key in the state file
*/
func (pro *Provider) getResources(state State) map[string]AvailableResources {

	module := state.getModuleFromIndex(0)

	groups := map[string]AvailableResources{}

	instances := []Instance{}

	for resourceIdentifier, oneResource := range module.Resources {
		switch oneResource.Type {
		case pro.machineIdentifier:
			oneInstance := Instance{}

			oneInstance.lookupKey(*pro, oneResource, module)
			oneInstance.setIndex(oneResource.getIndexFromKey(resourceIdentifier))
			oneInstance.setPrivateIp(oneResource.Primary.Attributes.getPrivateIp())
			oneInstance.setPublicIp(oneResource.Primary.Attributes.getPublicIp())
			oneInstance.setName(oneResource.getNameFromKey(resourceIdentifier))

			instances = append(instances, oneInstance)
			addToMap(groups, oneInstance)

		}

		//groups[resourceIdentifier] = AvailableResources{Instances: instances}
	}

	return groups
}

///*
//Group the similar kind of instances for ansible inventory
//*/
func addToMap(groups map[string]AvailableResources, instance Instance) {

		instances := groups[instance.name].Instances
	if _, ok := groups[instance.name]; ok {
		instances = append(instances, instance)
	} else {
		instances = []Instance{instance}
	}

	groups[instance.name] = AvailableResources{instances}

}
