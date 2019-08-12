package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/focusteach/gfworks/examples/model"
	"github.com/gin-gonic/gin"
)

// GetTopFaces 获取指定的人脸头像数据.
func GetTopFaces(c *gin.Context) {
	top := c.DefaultQuery("top", "10")
	var faces model.Faces

	ntop, _ := strconv.Atoi(top)

	faces.Top = ntop
	for i := 0; i < ntop; i++ {
		faces.Faces = append(faces.Faces, model.Face{
			FaceID:   strconv.Itoa(i),
			FaceName: "t1",
		})
	}
	c.JSON(0, faces)
}

func Upload(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
}
