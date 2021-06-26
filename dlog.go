package main

import (
	"log"
	"os"
	"time"
)

func init() {
	getDir, _ := os.Getwd()
	// 判断文件夹
	logStatus, _ := existFiles(getDir + "log/")
	if logStatus == false {
		_, _ = createFiles(getDir+"log/", 0755)
	}
}

func Info(accessLog string) {
	if f, err := os.OpenFile("log/"+getCurrentWjDate()+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		log.Println(err)
	} else {
		_, _ = f.WriteString(accessLog + "\n")
	}
}
func getCurrentWjDate() string {
	return time.Now().Format("20060102")
}

func existFiles(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createFiles(path string, perm int) (bool, error) {
	err := os.MkdirAll(path, os.FileMode(perm))
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
