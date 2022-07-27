package app

import starter "github.com/tsxylhs/go-starter/starter"

var Server = starter.NewService("code_server")
var WebApi = starter.NewWeb("code_test")
var MqttClient = starter.NewMqttClient("mqtt_Client")
var TcpClient = starter.NewTcpClient("tcp_Client")
