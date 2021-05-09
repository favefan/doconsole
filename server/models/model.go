package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/favefan/doconsole/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db *gorm.DB
)

// Model is the default model strcut.
type Model struct {
	Id        uint           `gorm:"primaryKey";form:"Id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Setup() {
	var err error

	// 默认logger实现
	defaultLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix,
			SingularTable: true,
		},
		Logger: defaultLogger,
	})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Auto Migrate
	db.AutoMigrate(&Registry{})
	db.AutoMigrate(&Host{})
	db.Exec("INSERT INTO `doconsole`.`doconsole_registry` (`created_at`, `updated_at`, `name`) VALUES (?, ?, 'DockerHub')", time.Now(), time.Now())
	db.Exec("INSERT INTO `doconsole`.`doconsole_host` (`created_at`, `updated_at`, `name`, `via_socket`) VALUES (?, ?, 'local', '1')", time.Now(), time.Now())

}

// CloseDB will not close DB.
func CloseDB() {
	//defer db.Close()
}
