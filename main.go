package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// model
type TODO struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	Title     string     `json:"title" gorm:"not null;size:32;default:''"`
	Status    bool       `json:"status" gorm:"default:false;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:datetime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:datetime"`
}

func initMysql() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			tmp := time.Now().Local().Format("2006-01-02 15:04:05")
			now, _ := time.ParseInLocation("2006-01-02 15:04:05", tmp, time.Local)
			return now
		},
	})
	return err
}

func main() {
	// 创建数据库连接
	err := initMysql()
	if err != nil {
		panic(err.Error())
	}
	// 创建table
	err = DB.Debug().AutoMigrate(&TODO{})
	if err != nil {
		panic(err.Error())
	}

	g := gin.Default()
	g.Static("/static", "static")
	g.LoadHTMLGlob("templates/*")

	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := g.Group("/v1")
	{
		// 添加代办事项
		v1Group.POST("/todo", func(c *gin.Context) {
			todo := new(TODO)
			if err := c.BindJSON(&todo); err == nil {
				DB.Debug().Create(&todo)
				c.JSON(http.StatusOK, todo)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			}
		})

		// 获取全部代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []TODO
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		// 更新
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			todo := new(TODO)
			id := c.Param("id")

			if err := DB.Where("id = ?", id).Find(todo).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			err := c.BindJSON(todo)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}

			err = DB.Model(todo).Update("status", todo.Status).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todo)
		})

		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			todo := new(TODO)
			id := c.Param("id")

			if err = DB.Where("id=?", id).Delete(todo).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"id": id})
			}
		})
	}

	g.Run(":8100")
}
