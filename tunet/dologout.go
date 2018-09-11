package tunet

import (
	"log"
	"net/http"
	"time"
)

func DoLogout(user string) (error,error) {
	log.Println("Do logout")
	err1 := logout("http://auth4.tsinghua.edu.cn/cgi-bin/srun_portal?action=logout&username="+user+"@tsinghua")
	err2 := logout("http://auth4.tsinghua.edu.cn/cgi-bin/srun_portal?action=logout&username=" + user)
	i := 0
	for err1 != nil || err2 != nil{
		time.Sleep(time.Duration(20)*time.Second)
		err1 = logout("http://auth4.tsinghua.edu.cn/cgi-bin/srun_portal?action=logout&username="+user+"@tsinghua")
		err2 = logout("http://auth4.tsinghua.edu.cn/cgi-bin/srun_portal?action=logout&username=" + user)
		if i > 9{
			return err1,err2
		}
	}
	return nil,nil
}

func logout(url string) error {
	_, err := http.Get(url)
	if err != nil {
		log.Println("can't logout")
		return err
	}
	return nil
}
