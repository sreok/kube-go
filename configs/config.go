package configs

import "os"

var dir, _ = os.Getwd()
var KubeConfig = dir + "/configs/admin.conf"
var DbTable = "kube_go"
var Db = "root:123456@/" + DbTable + "?charset=utf8mb4&parseTime=true"
