package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var savePOST = func(ctx *gin.Context) {
	formname := ctx.Param("name")
	fullpath, _ := os.Executable()
	dir := filepath.Dir(fullpath)
	filename := filepath.Join(dir, formname+".log")
	err := ctx.Request.ParseForm()
	if err != nil {
		fmt.Print("parse error")
		log.Fatal(err)
	}
	form := ctx.Request.PostForm
	//b, err := json.MarshalIndent(form, "\n", " ")
	b, err := json.Marshal(form)

	if err != nil {
		fmt.Print("marshal error")
		log.Fatal(err)
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("file open error")
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(f, string(b))
	if err != nil {
		fmt.Print("write file error")
		log.Fatal(err)
	}
	ctx.String(200, fmt.Sprintf("保存しました.閉じてOKです。\n%s\n", string(b)))
	//ctx.String(200, "OK")

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	fullpath, _ := os.Executable()
	dir := filepath.Dir(fullpath)
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(304, "/list")
	})
	router.StaticFile("/list", filepath.Join(dir, "index.html"))
	router.Static("/forms", filepath.Join(dir, "forms"))
	router.POST("/save/:name", savePOST)

	fmt.Println("下のURLをブラウザで開いてください")
	fmt.Println("http://localhost:8080/list/")
	router.Run(":8080")
}
