package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

/*   功能介绍
- 可扩展的Hook机制：允许使用者通过hook的方式将日志分发到任意地方,如本地文件系统、标准输出、logstash、elasticsearch或者mq等
- logrus 内置了两种日志格式,JSONFormatter和TextFormatter还可以自己动手实现接口Formatter,来定义自己的日志格式
- Field机制：logrus 鼓励通过Field机制进行精细化的、结构化的日志记录,而不是通过冗长的消息来记录日志

logrus.WithFields(logrus.Fields{
    "name": "dj",
    "age": 18,
  }).Info("info msg")

logrus.New() 可以自定义logger,创建任意个logger

*/

func init() {
	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	// 输出文件名和方法信息
	logrus.SetReportCaller(true)
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 第三方日志格式 nested-logrus-formatter

	// 设置日志输出位置
	stdout := os.Stdout
	file, _ := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE, 0755)
	logrus.SetOutput(io.MultiWriter(stdout, file))
}

func Info(args ...interface{}) {
	logrus.Infoln(args)
}
