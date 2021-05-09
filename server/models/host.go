package models

import "log"

type Host struct {
	Model
	Name string `gorm:"unique;not null";form:"Name"`
	ViaSocket bool `gorm:"default:false;not null";form:"ViaSocket"`
	DockerEngineURL string `gorm:"default:'#'";form:"URL"`
	HostIP string `gorm:"default:'#'";form:"HostIP"`
	TLS bool `gorm:"default:false;not null";form:"TLS"`
}

func (host *Host) Create() error {
	if err := db.Create(host).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*Host) Get(id int) (interface{}, error) {
	if id == -1 {
		var hosts []Host
		if err := db.Find(&hosts).Error; err != nil {
			log.Println(err)
			return nil, err
		}
		return hosts, nil
	}
	var host Host
	if err := db.First(&host, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return host, nil
}

func (host *Host) Update(id interface{}) error {
	if err := db.Model(&Host{}).Where("id = ?", id).Updates(host).Error;
		err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*Host) Delete(id interface{}) error {
	if err := db.Where("id = ?", id).Delete(Host{}).Error;
		err != nil {
		log.Println(err)
		return err
	}
	return nil
}
