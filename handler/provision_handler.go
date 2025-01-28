package handler

import (
	"encoding/json"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/provisioning"
)

func (c *CSMSHandler) OnSetVariables(chargingStationID string,request *provisioning.SetVariablesRequest) (response *provisioning.SetVariablesResponse, err error) {
	out,_ := json.Marshal(request)
	logDefault(chargingStationID, request.GetFeatureName()).Infof("%v",string(out))
	var setVariablesResponse = provisioning.NewSetVariablesResponse([]provisioning.SetVariableResult{})
	setVariablesResponse.SetVariableResult = append(setVariablesResponse.SetVariableResult, provisioning.SetVariableResult{})
	return &provisioning.SetVariablesResponse{}, nil
}