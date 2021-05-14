package initialize

import (
	"fmt"
	"gitee.com/favefan/doconsole/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func TableDBInit(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Registry{},
		models.Host{},
	)
	if err != nil {
		log.Fatalf("Register table failed: %v", err)
	}
	log.Println("Register table success.")

	db.Exec("INSERT INTO `doconsole`.`doconsole_registry` (`created_at`, `updated_at`, `name`) VALUES (?, ?, 'DockerHub')", time.Now(), time.Now())
	//db.Exec("INSERT INTO `doconsole`.`doconsole_host` (`created_at`, `updated_at`, `name`, `via_socket`) VALUES (?, ?, 'local', '1')", time.Now(), time.Now())
	log.Println("Table data init finished.")
}

func DBSetup() *gorm.DB {
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

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.Name,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DatabaseSetting.TablePrefix,
			SingularTable: true,
		},
		Logger: defaultLogger,
	})
	if err != nil {
		log.Fatalf("Database connection setup err: %v", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	sqlDB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	return db
}
