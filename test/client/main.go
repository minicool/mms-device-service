package main

import (
	"github.com/golang/glog"
	"os"
	"flag"
	"github.com/minicool/mms-device-service/src/conf"
	"strconv"
	"google.golang.org/grpc"
	pb "github.com/minicool/mms-device-service/proto"
	"golang.org/x/net/context"
)

func main() {
	//Init the command-line flags.
	flag.Parse()
	// user temp logs
	defer glog.Infof("Temp folder for log files: %s\n", os.TempDir())
	// Flushes all pending log I/O.
	defer glog.Flush()

	//load config
	config,err := conf.LoadConfig()
	if err != nil {
		return
	}
	address_port := strconv.Itoa(config.Ser_conf.Port)
	address := config.Ser_conf.Host + ":" + address_port
	glog.Infof("address %s %d",config.Ser_conf.Host,config.Ser_conf.Port)

	//grpc
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		glog.Errorln("Faild to Dial grpc",err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewRFIDDeviceServerClient(conn)

	// 调用方法
	req := &pb.DeviceRegeditRequest{"测试二",pb.DEVICETYPE_CARD,
		"192.168.1.101",33303,"12"}
	res, err := c.DeviceRegedit(context.Background(), req)
	if err != nil {
		glog.Errorln(err)
	}
	glog.Info(res)

}

