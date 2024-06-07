package configs

import "os"

var dir, _ = os.Getwd()
var KubeConfig = dir + "./configs/admin.conf"
