package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

//Get the request from gin context.
//Convert to byte array.
//Unmarshal the request to a map , so that we can modify only specific keys without losing other keys.
//Convert map again to a byte array.
//Set it back to the request.
//This needs to be done as once we read from body , the request is not available to the next handlers.
//Ignore marshalling/unmarshalling errors for simplicity
//Inject a new key "Time" in the request.

func DecryptRequest(ctx *gin.Context) {
	requestBytes := make([]byte, 100)
	bytesRead,_:=ctx.Request.Body.Read(requestBytes)

	requestKeyValueMap := make(map[string]interface{})
	json.Unmarshal(requestBytes[:bytesRead], &requestKeyValueMap)

	for key, val := range requestKeyValueMap {
		fmt.Println("Decrying key ", key)
		requestKeyValueMap[key] = "decrypted_" + fmt.Sprintf("%s",val) //simulate decryption
	}
	requestKeyValueMap["Time"] = time.Now()


	requestBytes , _ = json.Marshal(requestKeyValueMap)
    ctx.Request.Body = io.NopCloser(bytes.NewReader(requestBytes))

    ctx.Next()
}
