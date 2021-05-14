package models

import (
	"gitee.com/favefan/doconsole/global"
	"log"
)

type Host struct {
	Model
	Name string `gorm:"unique;not null";form:"Name"`
	ViaSocket bool `gorm:"default:false;not null";form:"ViaSocket"`
	DockerEngineURL string `gorm:"default:'#'";form:"URL"`
	HostIP string `gorm:"default:'#'";form:"HostIP"`
	TLS bool `gorm:"default:false;not null";form:"TLS"`
}

func (host *Host) Create() error {
	if err := global.GDB.Create(host).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*Host) Get(id int) (interface{}, error) {
	if id == -1 {
		var hosts []Host
		if err := global.GDB.Find(&hosts).Error; err != nil {
			log.Println(err)
			return nil, err
		}
		return hosts, nil
	}
	var host Host
	if err := global.GDB.First(&host, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return host, nil
}

func (host *Host) Update(id interface{}) error {
	if err := global.GDB.Model(&Host{}).Where("id = ?", id).Updates(host).Error;
		err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*Host) Delete(id interface{}) error {
	if err := global.GDB.Where("id = ?", id).Delete(Host{}).Error;
		err != nil {
		log.Println(err)
		return err
	}
	return nil
}
