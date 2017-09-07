package model

import (
	"github.com/minicool/mms-device-service/src/db"
	"github.com/golang/glog"
	"time"
)

type DeviceRegedit struct {
	ID int					`gorm:"primary_key;not null;"`
	DeviceName string		`gorm:"type:varchar(50);not null;"`
	DeviceType DeviceType	`gorm:"ForeignKey:DeviceTypeId;"`
	DeviceTypeId int32		`gorm:"type:tinyint;"`
	IpAddr string			`gorm:"type:varchar(24);not null;"`
	IpPort int32			`gorm:"type:int;"`
	MacAddr string			`gorm:"type:varchar(50);not null;unique;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func init(){
	db.RegisterModel(&DeviceRegedit{})
}

func CreateTable_DeviceRegedit(){
	err := db.Conn.HasTable(&DeviceType{})
	if err == false{
		db.Conn.CreateTable(&DeviceType{})
		glog.Info("create table DeviceType")
	}else{
		glog.Info("DeviceType table is exists")
	}
}

func AddDeviceRegedit(deviceName string,deviceTypeId int32,ipAddr string,ipPort int32,macAddr string){
	deviceType,err := FindDeviceType_Name(int(deviceTypeId))
	if err != nil {
		return
	}

	regedit := DeviceRegedit{DeviceType:*deviceType,DeviceName:deviceName,IpAddr:ipAddr,IpPort:ipPort,MacAddr:macAddr}

//	db.Conn.Model(&regedit).Related(&user, "DeviceTypeId")
	if db.Conn.NewRecord(regedit) == false {
		glog.Error("DeviceRegedit primary key is blank")
	}else{
		db.Conn.Create(&regedit)
	}
}