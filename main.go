package main

import (
	"fmt"
	"gotunet/tunet"
	"io/ioutil"
	"strings"
	"time"
	"log"
	"flag"
)

func main() {
	cycle := flag.Bool("c",true,"循环检测，保持登陆")
	fname := flag.String("f","./passwd","用户名和密码文件，用逗号分开")
	in := flag.Bool("in",false,"登录账户")
	out := flag.Bool("out",false,"登出账户")
	flag.Parse()
	passbyte, err := ioutil.ReadFile(*fname)
	if err != nil {
		fmt.Println("File is wrong")
		panic(err)
	}
	passwd := strings.Split(string(passbyte), ",")
	if *out {
		logout(passwd[0])
		return
	}
	if *cycle {
		cyclelogin(passwd[0],passwd[1])
	}
	if *in {
		login(passwd[0],passwd[1])
	}
	return
}

func login(user,passwd string) error {
	_,i,_ := tunet.CheckOnline(user)
	if i == 0 {
		log.Println("无法访问校外网络")
		tunet.DoLogout(user)
		time.Sleep(time.Duration(15)*time.Second)
	}else if i == -1{
		return fmt.Errorf("网络故障")
	}else{
		return nil
	}
	return tunet.DoLogin(user,passwd)
}

func logout(user string)(error,error){
	return tunet.DoLogout(user)
}

func cyclelogin(user,passwd string) error {
	err := login(user,passwd)
	for {
		if err != nil{
			time.Sleep(time.Duration(1800)*time.Second)
			err = login(user,passwd)
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Duration(150)*time.Second)
		err = login(user,passwd)
	}
	return nil
}
