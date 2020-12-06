package devices
//
//import (
//	"encoding/json"
//	"github.com/eclipse/paho.clients.golang"
//	"smart_empire/config"
//)
//
//type VibrationSensorMsg struct {
//	Battery int64 		`json:"battery"`
//	Voltage int64 		`json:"voltage"`
//	Linkquality int64 	`json:"linkquality"`
//	angle int64 	`json:"linkquality"`
//	angle_x int64 	`json:"linkquality"`
//	angle_y int64 	`json:"linkquality"`
//	angle_z int64 	`json:"linkquality"`
//	angle_z int64 	`json:"linkquality"`
//	angle_x_absolute int64 	`json:"linkquality"`
//	angle_y_absolute int64 	`json:"linkquality"`
//	strength int64 	`json:"linkquality"`
//	action string 	`json:"linkquality"`
//}
//
//type VibrationSensorType struct {
//	Name string
//	Topic string
//	MsgChan chan VibrationSensorMsg
//}
//
//var VibrationSensor = VibrationSensorType{
//	Name:    "VibrationSensor",
//	Topic:   config.Cfg.MqttClient.Sensors.Vibration.Topic,
//	MsgChan: make(chan VibrationSensorMsg),
//}
//
//func (vs VibrationSensorType) MqttHandler (msg clients.Message) {
//	var sensorMsg VibrationSensorMsg
//	json.Unmarshal(msg.Payload(), &sensorMsg)
//	ds.MsgChan <- sensorMsg
//}
