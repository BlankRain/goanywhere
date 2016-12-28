/**
生成设备编码二维码
**/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)
/* 生成二维码文件
**/
func qrcodeToFile(eCode, dirName, fileName, info string, fileModel bool) error {
	if fileModel == true {
		f := "./" + dirName + "/" + fileName + ".png"
		fmt.Println(f)
		return qrcode.WriteFile(eCode, qrcode.Medium, 512, f)
	} else {
		image := "/" + dirName + "/" + fileName + ".png"
		fmt.Println(buildHtml(eCode, dirName, image, info))
		return nil
	}

}
func buildHtml(eCode, dirName, image, info string) string {
	f := "<div><p>单位编号:%v 设备编码:%v</p><image src='%v'/> <p>%v</p></div>"
	return fmt.Sprintf(f, dirName, eCode, image, info)
}
func main() {
	fmt.Println(os.Args)

	fmt.Println("Start Building Qrcode!")
	model := false
	if len(os.Args) < 2 {
		fmt.Println("請傳參數 pic or something else")
		return
	}

	if os.Args[1] == "pic" {
		model = true
	}
	fmt.Println(`<html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<style>
img{width:200px;height:200px;}
body{font-size:small;}
div{  
border-style: dashed;   
display: inline-block;
padding: 10px;  
margin: 0px;   
border-color: green;height:300px; 
width:250px;    margin-top: 12px;   
margin-bottom: 2.5px;}
</style></head>`)
	lines("./10003.csv", model)
	fmt.Println("Done!")
	defer fmt.Println("</html>")
}
/**
按行解析
**/
func lines(fileName string, model bool) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		l := scanner.Text()
		a := strings.Split(l, ",")
		index += 1
		qrcodeToFile(a[1], a[0], a[1], strings.Join(a[2:], ""), model)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
