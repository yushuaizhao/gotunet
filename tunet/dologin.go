package tunet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func DoLogin(user,passwd string)(error){
	log.Println("*Do login")
	logInfo,err := login(user,passwd)
	i:=0
	for err != nil {
		if logInfo == "IP has been online, please logout."{
			break
		}
		log.Println(err)
		time.Sleep(time.Duration(20) * time.Second)
		logInfo,err = login(user,passwd)
		if i > 9 {
			return fmt.Errorf("多次尝试无果 " + logInfo)
		}
		i++
	}
	log.Println(logInfo)
	return nil
}

func login(user,passwd string) (string,error) {
	h := md5.New()
	io.WriteString(h,passwd)
	MD5 := hex.EncodeToString(h.Sum(nil))
	data := map[string][]string{
		"action":   {"login"},
		"username": {user},
		"password": {"{MD5_HEX}" + MD5},
		"ac_id":    {"1"},
	}
	logInfo, err := http.PostForm("https://auth4.tsinghua.edu.cn/do_login.php", data)
	if err != nil {
		log.Println("can't login")
		return "",err
	}
	loginByte, err := ioutil.ReadAll(logInfo.Body)
	defer logInfo.Body.Close()
	if err != nil {
		log.Println("can't login ")
		return "",err
	}
	loginInfo := string(loginByte)
	if loginInfo != "Login is successful."{
		return loginInfo,fmt.Errorf(loginInfo)
	}
	return loginInfo,nil
}
