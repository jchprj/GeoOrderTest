package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
)

var logPrefix, logDir, logLevel string
var logStd bool

type logWriter struct {
	file     *os.File
	fileTime time.Time
}

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

//New log file everyday
func (w *logWriter) getFile() {
	now := time.Now().Local()
	if dateEqual(now, w.fileTime) {
		return
	}
	if w.file != nil {
		w.file.Close()
	}
	w.fileTime = now
	fileName := fmt.Sprintf("%s/%s%d_%02d_%02d.log", logDir, logPrefix, now.Year(),
		now.Month(), now.Day())

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	var err error
	if w.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664); err != nil {
		fmt.Println("open log file failed:", err)
	}
}

func (w *logWriter) Write(buf []byte) (int, error) {
	if logStd {
		os.Stdout.Write(buf)
	}
	w.getFile()
	if w.file != nil {
		w.file.Write(buf)
	}
	return len(buf), nil
}

func initLog() {
	writer := new(logWriter)
	if logDir != "" {
		// Umask 是权限的补码,用于设置创建文件和文件夹默认权限
		mask := syscall.Umask(0)
		defer syscall.Umask(mask)

		if err := os.MkdirAll(logDir, 0774); err != nil {
			fmt.Println("os.MkdirAll failed:", logDir, err)
		}
	}
	if lv, err := logrus.ParseLevel(logLevel); err == nil {
		logrus.SetLevel(lv)
	} else {
		fmt.Println("parse level failed:", logLevel, err)
	}
	logrus.SetOutput(writer)

	logrus.WithFields(logrus.Fields{
		"logDir":    logDir,
		"level":     logLevel,
		"logPrefix": logPrefix,
		"logStd":    logStd,
	}).Debug("setup log")
}
