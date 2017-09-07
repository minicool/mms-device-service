package model

import (
	"github.com/minicool/mms-device-service/src/db"
	"github.com/golang/glog"
	"time"
)

type DeviceType struct {
	ID int					`gorm:"primary_key;not null;"`
	TypeName string			`gorm:"type:varchar(50);not null;unique;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}


//func (u DeviceType) TableName() string {
//	if u.TableName() == "admin" {
//		return "admin_users"
//	} else {
//		return "users"
//	}
//}

func init(){
	db.RegisterModel(&DeviceType{})
	//db.Conn.Model(&DeviceRegedit{}).Related(&DeviceType{}, "DeviceRegeditId")
}

func InitDb_DeviceType(){
	db.Conn.Model(&DeviceRegedit{}).Related(&DeviceType{})
	CreateTable_DeviceType()
	AddDeviceType("打印机")
	AddDeviceType("发卡器")
	AddDeviceType("报警器")
}

func CreateTable_DeviceType(){
	err := db.Conn.HasTable(&DeviceRegedit{})
	if err == false{
		db.Conn.CreateTable(&DeviceRegedit{})
		glog.Info("create table DeviceRegedit")
	}else{
		glog.Info("DeviceRegedit table is exists")
	}
}

func AddDeviceType(typeName string){
	deviceType := DeviceType{TypeName:typeName}
	if db.Conn.NewRecord(deviceType) == false {
		glog.Error("deviceType primary key is blank")
	}else{
		db.Conn.Create(&deviceType)
	}
}

func FindDeviceType_Name(id int) (*DeviceType,error){
	deviceType := &DeviceType{ID: id}
	res := db.Conn.Debug().First(&deviceType, "id=?", id)
	err := res.Error
	if err != nil {
		glog.Errorf("Error finding user with id %s: %v", id, res.Error.Error())
	}
	return deviceType, err
}
