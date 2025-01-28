package handler

import (
	"encoding/json"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

func (c *CSMSHandler) OnNotifyChargingLimit(chargingStationID string, request *smartcharging.NotifyChargingLimitRequest) (response *smartcharging.NotifyChargingLimitResponse, err error) {
	out,_ := json.Marshal(request)
	logDefault(chargingStationID, request.GetFeatureName()).Infof("%v",string(out))
	return smartcharging.NewNotifyChargingLimitResponse(), nil
}

func (c *CSMSHandler) OnClearedChargingLimit(chargingStationID string, request *smartcharging.ClearedChargingLimitRequest) (response *smartcharging.ClearedChargingLimitResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Warnf("Unsupported feature")
	return &smartcharging.ClearedChargingLimitResponse{}, nil
}

func (c *CSMSHandler) OnNotifyEVChargingNeeds(chargingStationID string, request *smartcharging.NotifyEVChargingNeedsRequest) (response *smartcharging.NotifyEVChargingNeedsResponse, err error) {
	out,_ := json.Marshal(request)
	logDefault(chargingStationID, request.GetFeatureName()).Infof("%v",string(out))
	return smartcharging.NewNotifyEVChargingNeedsResponse(smartcharging.EVChargingNeedsStatusAccepted), nil
}

func (c *CSMSHandler) OnNotifyEVChargingSchedule(chargingStationID string, request *smartcharging.NotifyEVChargingScheduleRequest) (response *smartcharging.NotifyEVChargingScheduleResponse, err error) {
	out,_ := json.Marshal(request)
	logDefault(chargingStationID, request.GetFeatureName()).Infof("%v",string(out))
	return smartcharging.NewNotifyEVChargingScheduleResponse(types.GenericStatusAccepted), nil
}

func (c *CSMSHandler) OnReportChargingProfiles(chargingStationID string, request *smartcharging.ReportChargingProfilesRequest) (response *smartcharging.ReportChargingProfilesResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Warnf("Unsupported feature")
	return &smartcharging.ReportChargingProfilesResponse{}, nil
}

func  (c *CSMSHandler) OnGetCompositeSchedule(chargingStationID string, request *smartcharging.GetCompositeScheduleRequest) (response *smartcharging.GetCompositeScheduleResponse, err error) {
	out,_ := json.Marshal(request)
	logDefault(chargingStationID, request.GetFeatureName()).Infof("%v",string(out))
	return smartcharging.NewGetCompositeScheduleResponse(smartcharging.GetCompositeScheduleStatusAccepted,1), nil
}
