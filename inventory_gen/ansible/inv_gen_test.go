package ansible

import (
	"testing"
	)

func TestRead(t *testing.T) {
	r := &FileBased{
		inputSource: "test_data/small_test_state.tfstate",
	}

	prov := Provider{}
	prov.privateKeyIdentifier = "tls_private_key"
	prov.publicKeyIdentifier = "aws_key_pair"
	prov.machineIdentifier = "aws_instance"

	state := prov.getState(r)

	t.Run("Version should be 3 as in the test state file", func (t *testing.T){
		if state.Version != 3 {
			t.Errorf("Test version is not matching")
		}
	})

	t.Run("Number of instances identified should be 12", func (t *testing.T){
		if len(prov.getInstances(state)) != 12 {
			t.Errorf("Number of instances are not identified")
		}
	})

	t.Run("Number of instances identified for DB type should be 3", func (t *testing.T){
		if len(groupInstances(prov.getInstances(state))["db"]) != 3 {
			t.Errorf("Number of instances are not identified as 3")
		}
	})

	t.Run("Number of instances identified for type ANZ is 2", func (t *testing.T){
		if len(groupInstances(prov.getInstances(state))["anz"]) != 2 {
			t.Errorf("Number of instances are not identified as 2")
		}
	})

	t.Run("Number of instances identified for type monitoring is 2", func (t *testing.T){
		if len(groupInstances(prov.getInstances(state))["anz"]) != 2 {
			t.Errorf("Number of instances are not identified as 2 for monnitoring")
		}
	})

	t.Run("Number of instances identified for type api is 2", func (t *testing.T){
		if len(groupInstances(prov.getInstances(state))["api"]) != 2 {
			t.Errorf("Number of instances are not identified as 2 for monnitoring")
		}
	})

	t.Run("Number of instances identified for type dmz is 2", func (t *testing.T){
		if len(groupInstances(prov.getInstances(state))["dmz"]) != 2 {
			t.Errorf("Number of instances are not identified as 2 for DMZ")
		}
	})

	t.Run("Number of instances identified for type bastion is 1", func (t *testing.T){
		if len(groupInstances(prov.getInstances(state))["bastion"]) != 1 {
			t.Errorf("Number of instances are not identified as 1 for bastion")
		}
	})

	instances := prov.getInstances(state)
	groups := groupInstances(instances)
	t.Run("Test for bastion public/private IPs, private key", func (t *testing.T){

		bastion := groups["bastion"][0]

		if bastion.name != "bastion" {
			t.Errorf("Name is not matching bastion")
		}

		if bastion.index != 0 {
			t.Errorf("Index is not zero for bastion host")
		}

		if bastion.privateIp != "10.54.1.226" {
			t.Errorf("Private Ip is not matching for bastion host")
		}

		if bastion.publicIp != "13.56.234.108" {
			t.Errorf("Public Ip is not matching for bastion host")
		}

		//Check the content of private key, for test it is a clear text
		//But in actual it is a hashed value
		if bastion.key.privateKey != "bastion_private_key" {
			t.Errorf("Private key content is not matching for bastion host")
		}
	})
}