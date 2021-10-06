package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SzData struct {
	Msg    string `comment:"msg"`
	Status string `comment:"status"`
	Data   []struct {
		Status    string `comment:"status"`
		DispNum   string `comment:"dispNum"`
		Disp_data []struct {
			StdStg string `comment:"StdStg"`
			StdStl string `comment:"StdStl"`
			Loc    string `comment:"loc"`
			Tid    string `comment:"tid"`
			Name   string `comment:"name"`
			Eid    string `comment:"eid"`
			City   string `comment:"city"`
			Type   string `comment:"type"`
			Exp    string `comment:"exp"`
		} `bson:"disp_data"`
	} `bson:"data"`
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds)

	var name, city string
	flag.StringVar(&name, "n", "", "<username>")
	flag.StringVar(&city, "c", "", "<city>")
	flag.Parse()
	if len(os.Args) < 4 {
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		var rs SzData
		fmt.Println(fmt.Sprintf("%-16s%-9s%-9s%-8s%-7s%-7s", "申请编号", "申请姓名", "中签日期", "中签城市", "中签类型", "中签指数"))
		resData := geturl(name, city)
		_ = json.Unmarshal([]byte(resData), &rs)

		for _, itemx := range rs.Data {
			for _, itemy := range itemx.Disp_data {
				fmt.Println(fmt.Sprintf("%-20s%-10s%-13s%-10s%-11s%-10s", itemy.Tid, itemy.Name, itemy.Eid, itemy.City, itemy.Type, itemy.Exp))
			}
		}
	}
}
func geturl(name, city string) string {
	client := &http.Client{}
	//读取配置文件的url, 并提交请求
	url := "https://sp0.baidu.com/9_Q4sjW91Qh3otqbppnN2DJv/pae/common/api/yaohao?name=" + name + "&city=" + city + "&format=json"

	resquest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(fmt.Sprintf("%s", err))
	}
	//处理返回的结果
	response, _ := client.Do(resquest)
	//关闭流
	defer response.Body.Close()
	//检出结果集
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println(err2)
	}
	return string(body)
}
