package logger

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var nacosLogger *log.Logger

func init() {
	nacosLogger = log.New(os.Stderr, "nacos-", log.LstdFlags)
}

func MkdirIfNecessary(createDir string) error {
	var path string
	var err error
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}

	s := strings.Split(createDir, path)
	startIndex := 0
	dir := ""
	if s[0] == "" {
		startIndex = 1
	} else {
		dir, _ = os.Getwd() //当前的目录
	}
	for i := startIndex; i < len(s); i++ {
		d := dir + path + strings.Join(s[startIndex:i+1], path)
		if _, e := os.Stat(d); os.IsNotExist(e) {
			err = os.Mkdir(d, os.ModePerm) //在当前目录下生成md目录
			if err != nil {
				break
			}
		}
	}
	return err
}

func InitLog(logDir string) error {
	err := MkdirIfNecessary(logDir)
	if err != nil {
		return err
	}
	logDir = logDir + string(os.PathSeparator)
	rl, err := rotatelogs.New(filepath.Join(logDir, "nacos-sdk.log-%Y%m%d%H%M"), rotatelogs.WithRotationTime(time.Hour), rotatelogs.WithMaxAge(48*time.Hour), rotatelogs.WithLinkName(filepath.Join(logDir, "nacos-sdk.log")))
	if err != nil {
		return err
	}
	nacosLogger.SetOutput(rl)
	nacosLogger.SetFlags(log.LstdFlags)
	return nil
}

func Printf(format string, v ...interface{}) {
	nacosLogger.Printf(format, v)
}

func Println(v ...interface{}) {
	nacosLogger.Println(v)
}

func Fatalf(format string, v ...interface{}) {
	nacosLogger.Fatalf(format, v)
}

func Print(v ...interface{}) {
	nacosLogger.Print(v)
}

func Panicf(format string, v ...interface{}) {
	nacosLogger.Panicf(format, v)
}
