package handler

import (
	"encoding/json"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/smartcharging"
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
	logDefault(chargingStationID, request.GetFeatureName()).Warnf("Unsupported feature")
	return &smartcharging.NotifyEVChargingNeedsResponse{}, nil
}

func (c *CSMSHandler) OnNotifyEVChargingSchedule(chargingStationID string, request *smartcharging.NotifyEVChargingScheduleRequest) (response *smartcharging.NotifyEVChargingScheduleResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Warnf("Unsupported feature")
	return &smartcharging.NotifyEVChargingScheduleResponse{}, nil
}

func (c *CSMSHandler) OnReportChargingProfiles(chargingStationID string, request *smartcharging.ReportChargingProfilesRequest) (response *smartcharging.ReportChargingProfilesResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Warnf("Unsupported feature")
	return &smartcharging.ReportChargingProfilesResponse{}, nil
}
