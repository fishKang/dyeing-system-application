package util

import (
	"encoding/json"
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

func NewLogZap() *zap.Logger {
	//日志文件存放目录
	writeSyncer, _ := os.Create("./info.log")
	//编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//日志等级字母大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//AddCaller()为显示文件名和行号
	log := zap.New(core, zap.AddCaller())
	return log
}

func NewSugarLogZap() *zap.SugaredLogger {
	//日志文件存放目录
	writeSyncer := getLogWriter()
	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoder := getEncoder()
	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//AddCaller()为显示文件名和行号
	log := zap.New(core, zap.AddCaller())
	return log.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		// Filename:   "./test.log",
		Filename:   os.Getenv("LOG_NAME"),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	//编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//日志等级字母大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func PrintInfo(log *zap.Logger, serialNum string, str string) {
	log.Info(serialNum, zap.String("msg", str))
}

// func SugarPrintInfo(sugarLog *zap.SugaredLogger, eduChannel entity.Channel, str string) {
// 	// sugarLog.Info(eduChannel.SerialNum,zap.String().Service,eduChannel.Method, zap.String("msg", str))
// 	sugarLog.Infof("%s %s.%s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, str)
// }

func SugarPrintInfo(sugarLog *zap.SugaredLogger, eduChannel entity.Channel, T any) error {
	jsonToMap, err := json.Marshal(T)
	if err != nil {
		sugarLog.Infof("%s %s.%s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, "参数转换失败")
		return errors.New("参数转换失败")
	}
	sugarLog.Infof("%s %s.%s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, jsonToMap)
	return nil
}
func RequestSugarPrintInfo(sugarLog *zap.SugaredLogger, eduChannel entity.Channel, T any) error {
	jsonToMap, err := json.Marshal(T)
	if err != nil {
		sugarLog.Infof("%s %s.%s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, "参数转换失败")
		return errors.New("参数转换失败")
	}
	sugarLog.Infof("%s %s.%s %s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, eduChannel.Method+"入参：", jsonToMap)
	eduChannel.Request = string(jsonToMap)
	return nil
}
func ResponseSugarPrintInfo(sugarLog *zap.SugaredLogger, eduChannel entity.Channel, T any) error {
	jsonToMap, err := json.Marshal(T)
	if err != nil {
		sugarLog.Infof("%s %s.%s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, "参数转换失败")
		return errors.New("参数转换失败")
	}
	sugarLog.Infof("%s %s.%s %s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, eduChannel.Method+"出参：", jsonToMap)
	eduChannel.Response = string(jsonToMap)
	return nil
}

func SugarPrintSimple(sugarLog *zap.SugaredLogger, str string) {
	// sugarLog.Info(eduChannel.SerialNum,zap.String().Service,eduChannel.Method, zap.String("msg", str))
	sugarLog.Infof("%s", str)
}

func PrintPanic(log *zap.Logger, serialNum string, str string) {
	log.Panic(serialNum, zap.String("msg", str))
}

func SugarPrintPanic(sugarLog *zap.SugaredLogger, eduChannel entity.Channel, str string) {
	sugarLog.Infof("%s %s.%s %s", eduChannel.SerialNum, eduChannel.Service, eduChannel.Method, str)
	sugarLog.Panic(eduChannel.SerialNum + " " + eduChannel.Service + "." + eduChannel.Method + " " + str)
}
