// Package all contains the all dialect.
//
//autogenerated:yes
package all

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialect"
	"github.com/bluenviron/gomavlib/v3/pkg/message"
)

// Dialect contains the dialect definition.
var Dialect = dial

// dial is not exposed directly in order not to display it in godoc.
var dial = &dialect.Dialect{
	Version: 2,
	Messages: []message.Message{
		// minimal
		&MessageHeartbeat{},
		&MessageProtocolVersion{},
		// standard
		// common
		&MessageSysStatus{},
		&MessageSystemTime{},
		&MessagePing{},
		&MessageChangeOperatorControl{},
		&MessageChangeOperatorControlAck{},
		&MessageAuthKey{},
		&MessageLinkNodeStatus{},
		&MessageSetMode{},
		&MessageParamRequestRead{},
		&MessageParamRequestList{},
		&MessageParamValue{},
		&MessageParamSet{},
		&MessageGpsRawInt{},
		&MessageGpsStatus{},
		&MessageScaledImu{},
		&MessageRawImu{},
		&MessageRawPressure{},
		&MessageScaledPressure{},
		&MessageAttitude{},
		&MessageAttitudeQuaternion{},
		&MessageLocalPositionNed{},
		&MessageGlobalPositionInt{},
		&MessageRcChannelsScaled{},
		&MessageRcChannelsRaw{},
		&MessageServoOutputRaw{},
		&MessageMissionRequestPartialList{},
		&MessageMissionWritePartialList{},
		&MessageMissionItem{},
		&MessageMissionRequest{},
		&MessageMissionSetCurrent{},
		&MessageMissionCurrent{},
		&MessageMissionRequestList{},
		&MessageMissionCount{},
		&MessageMissionClearAll{},
		&MessageMissionItemReached{},
		&MessageMissionAck{},
		&MessageSetGpsGlobalOrigin{},
		&MessageGpsGlobalOrigin{},
		&MessageParamMapRc{},
		&MessageMissionRequestInt{},
		&MessageSafetySetAllowedArea{},
		&MessageSafetyAllowedArea{},
		&MessageAttitudeQuaternionCov{},
		&MessageNavControllerOutput{},
		&MessageGlobalPositionIntCov{},
		&MessageLocalPositionNedCov{},
		&MessageRcChannels{},
		&MessageRequestDataStream{},
		&MessageDataStream{},
		&MessageManualControl{},
		&MessageRcChannelsOverride{},
		&MessageMissionItemInt{},
		&MessageVfrHud{},
		&MessageCommandInt{},
		&MessageCommandLong{},
		&MessageCommandAck{},
		&MessageCommandCancel{},
		&MessageManualSetpoint{},
		&MessageSetAttitudeTarget{},
		&MessageAttitudeTarget{},
		&MessageSetPositionTargetLocalNed{},
		&MessagePositionTargetLocalNed{},
		&MessageSetPositionTargetGlobalInt{},
		&MessagePositionTargetGlobalInt{},
		&MessageLocalPositionNedSystemGlobalOffset{},
		&MessageHilState{},
		&MessageHilControls{},
		&MessageHilRcInputsRaw{},
		&MessageHilActuatorControls{},
		&MessageOpticalFlow{},
		&MessageGlobalVisionPositionEstimate{},
		&MessageVisionPositionEstimate{},
		&MessageVisionSpeedEstimate{},
		&MessageViconPositionEstimate{},
		&MessageHighresImu{},
		&MessageOpticalFlowRad{},
		&MessageHilSensor{},
		&MessageSimState{},
		&MessageRadioStatus{},
		&MessageFileTransferProtocol{},
		&MessageTimesync{},
		&MessageCameraTrigger{},
		&MessageHilGps{},
		&MessageHilOpticalFlow{},
		&MessageHilStateQuaternion{},
		&MessageScaledImu2{},
		&MessageLogRequestList{},
		&MessageLogEntry{},
		&MessageLogRequestData{},
		&MessageLogData{},
		&MessageLogErase{},
		&MessageLogRequestEnd{},
		&MessageGpsInjectData{},
		&MessageGps2Raw{},
		&MessagePowerStatus{},
		&MessageSerialControl{},
		&MessageGpsRtk{},
		&MessageGps2Rtk{},
		&MessageScaledImu3{},
		&MessageDataTransmissionHandshake{},
		&MessageEncapsulatedData{},
		&MessageDistanceSensor{},
		&MessageTerrainRequest{},
		&MessageTerrainData{},
		&MessageTerrainCheck{},
		&MessageTerrainReport{},
		&MessageScaledPressure2{},
		&MessageAttPosMocap{},
		&MessageSetActuatorControlTarget{},
		&MessageActuatorControlTarget{},
		&MessageAltitude{},
		&MessageResourceRequest{},
		&MessageScaledPressure3{},
		&MessageFollowTarget{},
		&MessageControlSystemState{},
		&MessageBatteryStatus{},
		&MessageAutopilotVersion{},
		&MessageLandingTarget{},
		&MessageFenceStatus{},
		&MessageMagCalReport{},
		&MessageEfiStatus{},
		&MessageEstimatorStatus{},
		&MessageWindCov{},
		&MessageGpsInput{},
		&MessageGpsRtcmData{},
		&MessageHighLatency{},
		&MessageHighLatency2{},
		&MessageVibration{},
		&MessageHomePosition{},
		&MessageSetHomePosition{},
		&MessageMessageInterval{},
		&MessageExtendedSysState{},
		&MessageAdsbVehicle{},
		&MessageCollision{},
		&MessageV2Extension{},
		&MessageMemoryVect{},
		&MessageDebugVect{},
		&MessageNamedValueFloat{},
		&MessageNamedValueInt{},
		&MessageStatustext{},
		&MessageDebug{},
		&MessageSetupSigning{},
		&MessageButtonChange{},
		&MessagePlayTune{},
		&MessageCameraInformation{},
		&MessageCameraSettings{},
		&MessageStorageInformation{},
		&MessageCameraCaptureStatus{},
		&MessageCameraImageCaptured{},
		&MessageFlightInformation{},
		&MessageMountOrientation{},
		&MessageLoggingData{},
		&MessageLoggingDataAcked{},
		&MessageLoggingAck{},
		&MessageVideoStreamInformation{},
		&MessageVideoStreamStatus{},
		&MessageCameraFovStatus{},
		&MessageCameraTrackingImageStatus{},
		&MessageCameraTrackingGeoStatus{},
		&MessageCameraThermalRange{},
		&MessageGimbalManagerInformation{},
		&MessageGimbalManagerStatus{},
		&MessageGimbalManagerSetAttitude{},
		&MessageGimbalDeviceInformation{},
		&MessageGimbalDeviceSetAttitude{},
		&MessageGimbalDeviceAttitudeStatus{},
		&MessageAutopilotStateForGimbalDevice{},
		&MessageGimbalManagerSetPitchyaw{},
		&MessageGimbalManagerSetManualControl{},
		&MessageEscInfo{},
		&MessageEscStatus{},
		&MessageWifiConfigAp{},
		&MessageAisVessel{},
		&MessageUavcanNodeStatus{},
		&MessageUavcanNodeInfo{},
		&MessageParamExtRequestRead{},
		&MessageParamExtRequestList{},
		&MessageParamExtValue{},
		&MessageParamExtSet{},
		&MessageParamExtAck{},
		&MessageObstacleDistance{},
		&MessageOdometry{},
		&MessageTrajectoryRepresentationWaypoints{},
		&MessageTrajectoryRepresentationBezier{},
		&MessageCellularStatus{},
		&MessageIsbdLinkStatus{},
		&MessageCellularConfig{},
		&MessageRawRpm{},
		&MessageUtmGlobalPosition{},
		&MessageDebugFloatArray{},
		&MessageOrbitExecutionStatus{},
		&MessageSmartBatteryInfo{},
		&MessageBatteryInfo{},
		&MessageGeneratorStatus{},
		&MessageActuatorOutputStatus{},
		&MessageTimeEstimateToTarget{},
		&MessageTunnel{},
		&MessageCanFrame{},
		&MessageOnboardComputerStatus{},
		&MessageComponentInformation{},
		&MessageComponentInformationBasic{},
		&MessageComponentMetadata{},
		&MessagePlayTuneV2{},
		&MessageSupportedTunes{},
		&MessageEvent{},
		&MessageCurrentEventSequence{},
		&MessageRequestEvent{},
		&MessageResponseEventError{},
		&MessageIlluminatorStatus{},
		&MessageCanfdFrame{},
		&MessageCanFilterModify{},
		&MessageWheelDistance{},
		&MessageWinchStatus{},
		&MessageOpenDroneIdBasicId{},
		&MessageOpenDroneIdLocation{},
		&MessageOpenDroneIdAuthentication{},
		&MessageOpenDroneIdSelfId{},
		&MessageOpenDroneIdSystem{},
		&MessageOpenDroneIdOperatorId{},
		&MessageOpenDroneIdMessagePack{},
		&MessageOpenDroneIdArmStatus{},
		&MessageOpenDroneIdSystemUpdate{},
		&MessageHygrometerSensor{},
		// uavionix
		&MessageUavionixAdsbOutCfg{},
		&MessageUavionixAdsbOutDynamic{},
		&MessageUavionixAdsbTransceiverHealthReport{},
		// icarous
		&MessageIcarousHeartbeat{},
		&MessageIcarousKinematicBands{},
		// cubepilot
		&MessageCubepilotRawRc{},
		&MessageHerelinkVideoStreamInformation{},
		&MessageHerelinkTelem{},
		&MessageCubepilotFirmwareUpdateStart{},
		&MessageCubepilotFirmwareUpdateResp{},
		// csairlink
		&MessageAirlinkAuth{},
		&MessageAirlinkAuthResponse{},
		&MessageAirlinkEyeGsHolePushRequest{},
		&MessageAirlinkEyeGsHolePushResponse{},
		&MessageAirlinkEyeHp{},
		&MessageAirlinkEyeTurnInit{},
		// ardupilotmega
		&MessageSensorOffsets{},
		&MessageSetMagOffsets{},
		&MessageMeminfo{},
		&MessageApAdc{},
		&MessageDigicamConfigure{},
		&MessageDigicamControl{},
		&MessageMountConfigure{},
		&MessageMountControl{},
		&MessageMountStatus{},
		&MessageFencePoint{},
		&MessageFenceFetchPoint{},
		&MessageAhrs{},
		&MessageSimstate{},
		&MessageHwstatus{},
		&MessageRadio{},
		&MessageLimitsStatus{},
		&MessageWind{},
		&MessageData16{},
		&MessageData32{},
		&MessageData64{},
		&MessageData96{},
		&MessageRangefinder{},
		&MessageAirspeedAutocal{},
		&MessageRallyPoint{},
		&MessageRallyFetchPoint{},
		&MessageCompassmotStatus{},
		&MessageAhrs2{},
		&MessageCameraStatus{},
		&MessageCameraFeedback{},
		&MessageBattery2{},
		&MessageAhrs3{},
		&MessageAutopilotVersionRequest{},
		&MessageRemoteLogDataBlock{},
		&MessageRemoteLogBlockStatus{},
		&MessageLedControl{},
		&MessageMagCalProgress{},
		&MessageEkfStatusReport{},
		&MessagePidTuning{},
		&MessageDeepstall{},
		&MessageGimbalReport{},
		&MessageGimbalControl{},
		&MessageGimbalTorqueCmdReport{},
		&MessageGoproHeartbeat{},
		&MessageGoproGetRequest{},
		&MessageGoproGetResponse{},
		&MessageGoproSetRequest{},
		&MessageGoproSetResponse{},
		&MessageRpm{},
		&MessageDeviceOpRead{},
		&MessageDeviceOpReadReply{},
		&MessageDeviceOpWrite{},
		&MessageDeviceOpWriteReply{},
		&MessageAdapTuning{},
		&MessageVisionPositionDelta{},
		&MessageAoaSsa{},
		&MessageEscTelemetry_1To_4{},
		&MessageEscTelemetry_5To_8{},
		&MessageEscTelemetry_9To_12{},
		&MessageOsdParamConfig{},
		&MessageOsdParamConfigReply{},
		&MessageOsdParamShowConfig{},
		&MessageOsdParamShowConfigReply{},
		&MessageObstacleDistance_3d{},
		&MessageWaterDepth{},
		&MessageMcuStatus{},
		// asluav
		&MessageCommandIntStamped{},
		&MessageCommandLongStamped{},
		&MessageSensPower{},
		&MessageSensMppt{},
		&MessageAslctrlData{},
		&MessageAslctrlDebug{},
		&MessageAsluavStatus{},
		&MessageEkfExt{},
		&MessageAslObctrl{},
		&MessageSensAtmos{},
		&MessageSensBatmon{},
		&MessageFwSoaringData{},
		&MessageSensorpodStatus{},
		&MessageSensPowerBoard{},
		&MessageGsmLinkStatus{},
		&MessageSatcomLinkStatus{},
		&MessageSensorAirflowAngles{},
		// development
		&MessageParamAckTransaction{},
		&MessageAirspeed{},
		&MessageWifiNetworkInfo{},
		&MessageSetVelocityLimits{},
		&MessageVelocityLimits{},
		&MessageFigureEightExecutionStatus{},
		&MessageBatteryStatusV2{},
		&MessageFuelStatus{},
		&MessageGroupStart{},
		&MessageGroupEnd{},
		&MessageRadioRcChannels{},
		&MessageAvailableModes{},
		&MessageCurrentMode{},
		&MessageAvailableModesMonitor{},
		&MessageGnssIntegrity{},
		&MessageTargetAbsolute{},
		&MessageTargetRelative{},
		&MessageControlStatus{},
		// pythonarraytest
		&MessageArrayTest_0{},
		&MessageArrayTest_1{},
		&MessageArrayTest_3{},
		&MessageArrayTest_4{},
		&MessageArrayTest_5{},
		&MessageArrayTest_6{},
		&MessageArrayTest_7{},
		&MessageArrayTest_8{},
		// test
		&MessageTestTypes{},
		// ualberta
		&MessageNavFilterBias{},
		&MessageRadioCalibration{},
		&MessageUalbertaSysStatus{},
		// storm32
		&MessageStorm32GimbalManagerInformation{},
		&MessageStorm32GimbalManagerStatus{},
		&MessageStorm32GimbalManagerControl{},
		&MessageStorm32GimbalManagerControlPitchyaw{},
		&MessageStorm32GimbalManagerCorrectRoll{},
		&MessageQshotStatus{},
		&MessageFrskyPassthroughArray{},
		&MessageParamValueArray{},
		// avssuas
		&MessageAvssPrsSysStatus{},
		&MessageAvssDronePosition{},
		&MessageAvssDroneImu{},
		&MessageAvssDroneOperationMode{},
		// all
	},
}
