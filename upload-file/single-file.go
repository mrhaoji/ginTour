package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		log.Println(file.Filename)

		if err != nil {
			c.String(http.StatusBadRequest, "a Bad request")
			return
		}

		filename := file.Filename
		fmt.Println(filename)

		if err := c.SaveUploadedFile(file, "./test/" + filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.String(http.StatusCreated, "upload successful")
	})

	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			filename := file.Filename
			c.SaveUploadedFile(file, "./test/" + filename)
		}

		c.String(http.StatusCreated, "%d files upload successful", len(files))
	})

	r.Run(":8080")
}

//curl -X POST http://localhost:8080/upload \
//-F "file=@/Users/haoji/Downloads/lh.jpg" \
//-H "Content-Type: multipart/form-data"

//curl -X POST http://localhost:8080/uploads \
//-F "upload[]=@/Users/haoji/Downloads/youxi.py" \
//-F "upload[]=@/Users/haoji/Downloads/lehakepu.jpg" \
//-H "Content-Type: multipart/form-data"