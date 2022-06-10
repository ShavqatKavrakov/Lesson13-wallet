package wallet

import (
	"fmt"
	"testing"
)

func TestService_FindAccoundById_success(t *testing.T) {
	svc := &Service{}
	_, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	_, err1 := svc.FindAccontById(1)
	if err != nil {
		fmt.Println(err1)
	}
}
func TestService_FindAccoundById_notFound(t *testing.T) {
	svc := &Service{}
	_, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	_, err1 := svc.FindAccontById(2)
	if err1 != nil {
		fmt.Println(err1)
	}
}
