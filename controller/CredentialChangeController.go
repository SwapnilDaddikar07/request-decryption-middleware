package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	request2 "swapnildaddikar07/request-decryption-middleware/request"
	"time"
)

type CredentialChangeController struct {
}

func NewCredentialChangeController() CredentialChangeController {
	return CredentialChangeController{}
}

func (c CredentialChangeController) CredentialChange(ctx *gin.Context) {
	fmt.Println("Change credential request received. Binding request body for further processing.")
	request := request2.CredentialChangeRequest{}
	bindError := ctx.ShouldBindJSON(&request)
	if bindError != nil {
		fmt.Println("Cannot find request. Returning bad request.")
		ctx.AbortWithStatus(400)
		return
	}

	copyInto:=make([]string,1)
	copy(copyInto,[]string{"!23"})

	fmt.Println("Decrypted credentials are ",request)
	//Process the credential change request. And return success if no errors.

	//Simulating credential change processing.
	time.Sleep(time.Second)

	//Success response
	ctx.Status(http.StatusOK)
}
