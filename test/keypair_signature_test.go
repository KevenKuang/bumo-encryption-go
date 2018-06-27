// keypair_signature_test
package keypair_signature

import (
	"testing"

	"github.com/bumoproject/bumo-encryption-go/src/keypair"
	"github.com/bumoproject/bumo-encryption-go/src/signature"
)

func Test_Create(t *testing.T) {
	public, private, address, err := keypair.Create()
	if err != nil {
		t.Error(err)
	}
	t.Log("Test_Create succeed", public, private, address)
}

func Test_CheckPrivateKey(t *testing.T) {
	private := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
	if keypair.CheckPrivateKey(private) {
		t.Log("Test_CheckPrivateKey succeed")
	}
	t.Error("Test_CheckPrivateKey failured")
}

func Test_CheckPublicKe(t *testing.T) {
	public := "b00180c2007082d1e2519a0f2d08fd65ba607fe3b8be646192a2f18a5fa0bee8f7a810d011ed"

	if keypair.CheckPublicKey(public) {
		t.Log("Test_CheckPublicKey succeed")
	}
	t.Error("Test_CheckPublicKey failured")
}
func Test_CheckAddress(t *testing.T) {

	address := "buQs9npaCq9mNFZG18qu88ZcmXYqd6bqpTU3"
	if keypair.CheckAddress(address) {
		t.Log("Test_CheckAddress succeed")
	}
	t.Error("Test_CheckAddress failured")

}

func Test_GetEncPublicKey(t *testing.T) {

	private := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
	pubFromPriv, err := keypair.GetEncPublicKey(private)
	if err != nil {
		t.Error(err)
	}
	t.Log("Test_GetEncPublicKey succeed", pubFromPriv)
}
func Test_GetEncAddress(t *testing.T) {

	public := "b00180c2007082d1e2519a0f2d08fd65ba607fe3b8be646192a2f18a5fa0bee8f7a810d011ed"
	addrFromPub, err := keypair.GetEncAddress(public)
	if err != nil {
		t.Error(err)
	}
	t.Log("Test_GetEncPublicKey succeed", addrFromPub)
}

func Test_Sign(t *testing.T) {

	private := "privbvYfqQyG3kZyHE4RX4TYVa32htw8xG4WdpCTrymPUJQ923XkKVbM"
	var message []byte
	sign, err := signature.Sign(private, message)
	if err != nil {
		t.Error(err)
	}
	t.Log("Test_GetEncPublicKey succeed", sign)
}
func Test_Verify(t *testing.T) {

	public := "b00180c2007082d1e2519a0f2d08fd65ba607fe3b8be646192a2f18a5fa0bee8f7a810d011ed"
	sign := "c2e3c76a8d296eaed0517a27e29b8f678a67d8b7fd4116683140496ebdfbacbc0cc1b7d68d80ff107d2e4028fddeecde7aad309f6e332f89c4e7d97117fcb907"
	var message []byte
	if signature.Verify(public, message, sign) {
		t.Log("Test_Verify succeed")
	}
	t.Error("Test_Verify failured")

}
