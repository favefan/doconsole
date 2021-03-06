package models

import (
	"gitee.com/favefan/doconsole/global"
	"log"
)

type Registry struct {
	Model
	Name string `gorm:"unique;not null";form:"Name"`
	URL string `gorm:"default:'#';not null";form:"URL"`
	NeedAuth bool `gorm:"default:false;not null";form:"NeedAuth"`
	Username string `form:"Username"`
	Password string	`form:"Password"`
	Comment string `form:"Comment"`
}

func (registry *Registry) Create() error {
	if err := global.GDB.Create(registry).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*Registry) Get(id int) (interface{}, error) {
	if id == -1 {
		var registries []Registry
		if err := global.GDB.Find(&registries).Error; err != nil {
			log.Println(err)
			return nil, err
		}
		return registries, nil
	}
	var registry Registry
	if err := global.GDB.First(&registry, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return registry, nil
}

func (registry *Registry) Update(id interface{}) error {
	if err := global.GDB.Model(&Registry{}).Where("id = ?", id).Updates(registry).Error;
		err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*Registry) Delete(id interface{}) error {
	if err := global.GDB.Where("id = ?", id).Delete(Registry{}).Error;
		err != nil {
		log.Println(err)
		return err
	}
	return nil
}
