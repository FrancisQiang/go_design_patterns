package decorator

import (
	"fmt"
)

type Logger interface {
	LogErr(content string)
	LogInfo(content string)
	LogWarn(content string)
}

type AdvanceLogger interface {
	Logger
	LogErrWithColor(content string)
	LogInfoWithColor(content string)
	LogWarnWithColor(content string)
}

type MyLogger struct {

}

func (m MyLogger) LogErr(content string) {
	fmt.Println("err:" + content)
}

func (m MyLogger) LogInfo(content string) {
	fmt.Println("info:" + content)
}

func (m MyLogger) LogWarn(content string) {
	fmt.Println("warn:" + content)
}

type MyAdvanceLogger struct {
	logger Logger
}

func (m MyAdvanceLogger) LogErr(content string) {
	m.logger.LogErr(content)
}

func (m MyAdvanceLogger) LogInfo(content string) {
	m.logger.LogInfo(content)
}

func (m MyAdvanceLogger) LogWarn(content string) {
	m.logger.LogWarn(content)
}

func (m MyAdvanceLogger) LogErrWithColor(content string) {
	m.logger.LogErr("red:" + content)
}

func (m MyAdvanceLogger) LogInfoWithColor(content string) {
	m.logger.LogInfo("blue:" + content)
}

func (m MyAdvanceLogger) LogWarnWithColor(content string) {
	m.logger.LogWarn("yellow:" + content)
}


