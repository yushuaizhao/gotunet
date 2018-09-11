package tunet

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetUsage()([]string,error){
	userInfo, err := http.Get("http://auth4.tsinghua.edu.cn/rad_user_info.php")
	if err != nil {
		log.Println("can't get the succeed url")
		return nil,err
	}
	userByte, err := ioutil.ReadAll(userInfo.Body)
	defer userInfo.Body.Close()
	if err != nil {
		log.Println("无法获取用户信息")
		return nil,err
	}
	userString := string(userByte)
	userSplit := strings.Split(userString, ",")
	return userSplit,nil
}

func CheckOnline(userName string) ([]string,int,error) {
	log.Println("*Check Online")
	userSplit,err := GetUsage()
	if err != nil{
		return nil,-1,err
	}
	if userSplit[0] == "" {
		return nil,0,nil
	}else if userSplit[0] == userName{
		return userSplit,1 , nil
	}else{
		return nil,-1,nil
	}
	return userSplit,1,nil
}
