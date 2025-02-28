package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Product{},
			&model.Category{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2025-01-01 10:00:00','2025-01-01 10:00:00','Pen','Pen'),(2,'2025-02-01 15:30:00','2025-02-01 15:30:00','Workbook','Workbook')")

			DB.Exec("INSERT INTO `product`.`product` VALUES (1, '2025-01-01 10:05:00', '2025-01-01 10:10:00', 'Pen', 'A high-quality writing instrument designed for smooth and precise writing.', '/static/image/pen-1.png', 2.50),(2, '2025-02-01 15:35:00', '2025-02-01 15:40:00', 'Workbook', 'Perfect for students and professionals alike, with high-quality pages for notes and sketches.', '/static/image/workbook-1.png', 5.00),(3, '2025-01-01 10:05:00', '2025-01-01 10:10:00', 'Pen', 'A high-quality writing instrument designed for smooth and precise writing.', '/static/image/pen-2.png', 3.00),(4, '2025-02-01 15:35:00', '2025-02-01 15:40:00', 'Workbook', 'Perfect for students and professionals alike, with high-quality pages for notes and sketches.', '/static/image/workbook-2.png', 4.00),(5, '2025-01-01 10:05:00', '2025-01-01 10:10:00', 'Pen', 'A high-quality writing instrument designed for smooth and precise writing.', '/static/image/pen-3.png', 1.80),(6, '2025-02-01 15:35:00', '2025-02-01 15:40:00', 'Workbook', 'Perfect for students and professionals alike, with high-quality pages for notes and sketches.', '/static/image/workbook-3.png', 4.50),(7, '2025-02-01 15:35:00', '2025-02-01 15:40:00', 'Workbook', 'Perfect for students and professionals alike, with high-quality pages for notes and sketches.', '/static/image/workbook-4.png', 4.80)")

			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (1, 1), (2, 2), (3, 1), (4, 2), (5, 1), (6, 2), (7, 2)")

		}
	}
}