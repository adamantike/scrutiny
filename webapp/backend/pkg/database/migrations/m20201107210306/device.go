package m20201107210306

import (
	"time"
)

type Device struct {
	//GORM attributes, see: http://gorm.io/docs/conventions.html
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	WWN    string `json:"wwn" gorm:"primary_key"`
	HostId string `json:"host_id"`

	DeviceName     string  `json:"device_name"`
	Manufacturer   string  `json:"manufacturer"`
	ModelName      string  `json:"model_name"`
	InterfaceType  string  `json:"interface_type"`
	InterfaceSpeed string  `json:"interface_speed"`
	SerialNumber   string  `json:"serial_number"`
	Firmware       string  `json:"firmware"`
	RotationSpeed  int     `json:"rotational_speed"`
	Capacity       int64   `json:"capacity"`
	FormFactor     string  `json:"form_factor"`
	SmartSupport   bool    `json:"smart_support"`
	DeviceProtocol string  `json:"device_protocol"` //protocol determines which smart attribute types are available (ATA, NVMe, SCSI)
	DeviceType     string  `json:"device_type"`     //device type is used for querying with -d/t flag, should only be used by collector.
	SmartResults   []Smart `gorm:"foreignkey:DeviceWWN" json:"smart_results"`
}
