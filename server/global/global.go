package global

import (
	"github.com/docker/docker/client"
	"gorm.io/gorm"
)

var (
	GDB     *gorm.DB
	GClient	*client.Client
	// GVA_REDIS  *redis.Client
)
