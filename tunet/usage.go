package tunet

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func Usage(userName string) error {
	log.Println("*Get Usage")
	userSplit,i,err := CheckOnline(userName)
	if err != nil{
		return err
	}
	if i == 0{
		return fmt.Errorf("没有登录外网")
	}else if i == 1{
		transformBtoG := 1000.0 * 1000.0 * 1000.0
		loginSeconds := timeChange(userSplit[1])
		nowSeconds := timeChange(userSplit[2])
		download := haveUsed(userSplit[3]) / transformBtoG
		upload := haveUsed(userSplit[4]) / transformBtoG
		usage := haveUsed(userSplit[6]) / transformBtoG
		fmt.Println("用户名：", userSplit[0])
		fmt.Println("本次登录前已使用: ", usage, "G")
		fmt.Println("本次登录前还剩余: ", 25.0-usage, "G")
		fmt.Println("本次使用: ", download, "G")
		fmt.Println("本次上传: ", upload, "G")
		fmt.Println("登录时间：", time.Unix(loginSeconds, 0))
		fmt.Println("现在时间：", time.Unix(nowSeconds, 0))
		fmt.Println("本机ip: ", userSplit[8])
		fmt.Println("共使用: ", usage+download, "G")
		fmt.Println("共剩余: ", 25.0-usage-download, "G")
	}else{
		return fmt.Errorf("网络故障")
	}
	return nil
}

func haveUsed(used string) float64 {
	floatUsed, err := strconv.ParseFloat(used, 32)
	if err != nil {
		log.Println("获取流量额失败")
		return 0
	}
	return floatUsed
}

func timeChange(secTime string) int64 {
	thisTime, err := strconv.ParseInt(secTime, 10, 64)
	if err != nil {
		fmt.Println("时间有误")
		return 0
	}
	return thisTime
}
