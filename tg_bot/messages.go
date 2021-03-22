package tg_bot

import (
	"fmt"
	"html"
	"smart_empire/gateway"
	"smart_empire/gateway/devices"
	"smart_empire/gateway/rules"
	"strconv"
)

var temperatureIndicatorsMapping = map[string]string{
	"EXCELLENT": html.UnescapeString("&#" + strconv.Itoa(128994) + ";"),
	"NORMAL":    html.UnescapeString("&#" + strconv.Itoa(128993) + ";"),
	"BAD":       html.UnescapeString("&#" + strconv.Itoa(128308) + ";"),
}

var LightIndicatorsMapping = map[string]string{
	"ON":  html.UnescapeString("&#" + strconv.Itoa(9898) + ";"),
	"OFF": html.UnescapeString("&#" + strconv.Itoa(9899) + ";"),
}

var SirenIndicatorsMapping = map[bool]string{
	true:  html.UnescapeString("&#" + strconv.Itoa(128265) + ";"),
	false: html.UnescapeString("&#" + strconv.Itoa(128263) + ";"),
}

var DefModeIndicatorsMapping = map[string]string{
	"ON":  html.UnescapeString("&#" + strconv.Itoa(128737) + ";"),
	"OFF": html.UnescapeString("&#" + strconv.Itoa(127968) + ";"),
	"HOME": fmt.Sprintf(
		"%s%s",
		html.UnescapeString("&#" + strconv.Itoa(127968) + ";"),
		html.UnescapeString("&#" + strconv.Itoa(128737) + ";"),
	),
}

func getDoorEventMessage(msg devices.DoorSensorMsg) string {
	event := "closed"
	if !msg.Contact {
		event = "opened"
	}
	return fmt.Sprintf("*DoorSensor:* \n%s", event)
}

func getTemperatureEventMessage(msg devices.TemperatureSensorMsg) string {
	return fmt.Sprintf(
		"*TemperatureSensor:* \n"+
			"%.1f ℃ %s / %.1f %% %s",
		msg.Temperature,
		temperatureIndicatorsMapping[msg.TemperatureIndicator],
		msg.Humidity,
		temperatureIndicatorsMapping[msg.HumidityIndicator],
	)
}

func getLightEventMessage(msg devices.LightMsg) string {
	return fmt.Sprintf("*Light:* \n%s", LightIndicatorsMapping[msg.State])
}

func getSirenEventMessage(msg devices.SirenMsg) string {
	return fmt.Sprintf("*Siren:* \n%s", SirenIndicatorsMapping[msg.Value])
}

func getDefModeMessage(option string) string {
	return fmt.Sprintf("*DefMode:* \n%s", DefModeIndicatorsMapping[option])
}

func getAlarmMessage() string {
	return "ALARM! ALARM! ALARM!"
}

func getStatusMsg() string {
	temperatureMsg := fmt.Sprintf(
		"%.1f ℃ %s / %.1f %% %s",
		gateway.TemperatureSensor.GetLatestMsg().Temperature,
		temperatureIndicatorsMapping[gateway.TemperatureSensor.GetLatestMsg().TemperatureIndicator],
		gateway.TemperatureSensor.GetLatestMsg().Humidity,
		temperatureIndicatorsMapping[gateway.TemperatureSensor.GetLatestMsg().HumidityIndicator],
	)
	doorValue := gateway.DoorSensor.GetLatestMsg().Contact
	doorEvent := "closed"
	if !doorValue {
		doorEvent = "opened"
	}
	doorMsg := fmt.Sprintf(
		"%s%s",
		html.UnescapeString("&#"+strconv.Itoa(128682)+";"),
		doorEvent,
	)
	return fmt.Sprintf(
		"%s / %s / %s / %s / %s",
		temperatureMsg,
		DefModeIndicatorsMapping[rules.GetDefMode()],
		SirenIndicatorsMapping[gateway.Siren.GetLatestMsg().Value],
		LightIndicatorsMapping[gateway.Light.GetLatestMsg().State],
		doorMsg,
	)
}
