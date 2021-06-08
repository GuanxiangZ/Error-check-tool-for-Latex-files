package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/parser"
	"io"
	"net/http"
	"os"
)

func Serve(port uint32, file *os.File) error {
	gin.DefaultWriter = io.MultiWriter(file, os.Stderr)
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/api/upload", uploadHandler)
	return r.Run(fmt.Sprintf(":%d", port))
}


type uploadResponse struct {
	Error   uint8  `json:"err"`
	Msg     string `json:"msg"`
	Report  *parser.Report `json:"report"`
}

func uploadHandler(c *gin.Context) {
	var (
		formfile, _ = c.FormFile("file")
		file        io.Reader
		err         error
		content     parser.Report
		header  map[string][]string
		languageCode string
		filename string
	)

	header = c.Request.Header
	if header["languageCode"] != nil && len(header["languageCode"]) >0 {
		languageCode = header["languageCode"][0]
	} else {
		languageCode = "en-US"
	}

	if header["filename"] != nil && len(header["filename"]) >0 {
		filename = header["filename"][0]
	} else {
		filename = "demo.tex"
	}

	file, err = formfile.Open()
	if err != nil {
		c.JSON(http.StatusAccepted, &uploadResponse{
			1,
			err.Error(),
			nil,
		})
	} else {
		content, err = parser.ProcessTeX(file, filename, languageCode, false)
		if err != nil {
			c.JSON(http.StatusAccepted, &uploadResponse{
				Error:   2,
				Msg:     err.Error(),
				Report: nil,
			})
			return
		}

		c.JSON(http.StatusOK, &uploadResponse{
			0,
			"file upload",
			&content,
		})
	}
}
