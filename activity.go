package flogo_triggerEventMap

import (
	"bytes"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"net/http"
	"strings"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
	strMapUrl := context.GetInput("mapUrl").(string)
	strValueList := context.GetInput("valueList").(string)
	values := strings.Split(strValueList, ",")
	responses := make([]*http.Response,len(values))

	for index := 0; index < len(values);index++  {
		strRequestPayload := map[string]interface{}{"value":values[index],}
		payloadAsBytes, error := json.Marshal(strRequestPayload)

		if error != nil {
			return true, error
		}

		response, requestError := http.Post(strMapUrl, "application/json", bytes.NewBuffer(payloadAsBytes))

		if requestError != nil {
			return true, requestError
		}

		responses[index] = response
	}

	context.SetOutput("output", responses)
	return true, nil
}
