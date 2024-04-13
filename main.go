package main

import (
	"flag"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("index.html")

	var random_int int

	flag.IntVar(&random_int, "l", 10, "random int long")

	// 当访问根目录时，生成一个随机字符串，并重定向到"/random"路径
	r.GET("/", func(c *gin.Context) {
		randomString := randomString(random_int)
		c.Redirect(http.StatusFound, "/"+randomString)
	})

	r.GET("/:path", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		path := c.Param("path")
		filePath := "./_tmp_/" + path

		// 检查文件是否存在，如果不存在则创建
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// 创建目录
			if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			// 创建文件
			if _, err := os.Create(filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		}

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		body := string(fileContent)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": path,
			"body":  body,
		})
	})

	r.POST("/:path", func(c *gin.Context) {
		// 读取POST请求的内容
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
			return
		}

		// 获取文件路径
		path := c.Param("path")

		// 定义文件路径
		filePath := "./_tmp_/" + path

		// 创建文件夹，如果不存在的话
		dir := "./_tmp_/"
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		}

		// 写入文件
		err = ioutil.WriteFile(filePath, body, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to file"})
			return
		}

		// 返回成功的响应
		c.JSON(http.StatusOK, gin.H{"status": "Success"})

	})

	var port string

	flag.StringVar(&port, "p", ":80", "port to listen on")

	flag.Parse()

	r.Run(port)
}

// randomString 生成指定长度的随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
