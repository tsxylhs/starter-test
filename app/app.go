package app

import starter "github.com/tsxylhs/go-starter/starter"

var Server = starter.NewService("code_server", true, false)
var WebApi = starter.NewWeb("code_test")
var MqttClient = starter.NewMqttClient("mqtt")
var TcpClient = starter.NewTcpClient("tcp")
