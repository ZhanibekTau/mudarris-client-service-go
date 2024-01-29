package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	INFO  = "INFO "
	ERROR = "ERROR"

	reset = "\033[0m"
	red   = "\033[31m"
	green = "\033[32m"
	cyan  = "\033[36m"

	timeFormat = "02.01.06. 15:04:05"
)

type LogMessage struct {
	Text string
	Data interface{}
}

type logMessageWithExtra struct {
	*LogMessage
	Datetime time.Time
}

type Log struct{}

func (l *Log) Info(text string, data interface{}) {

	l.createLogMessage(text, data, INFO)

}

func (l *Log) Error(text string, data interface{}) {

	l.createLogMessage(text, data, ERROR)
}

func (l *Log) createLogMessage(text string, data interface{}, level string) {

	var color, tier string

	switch level {
	case INFO:
		color = green
		tier = INFO
	case ERROR:
		color = red
		tier = ERROR
	}

	message := l.getMessage(text, data)
	messageWithExtra := l.addExtraInfo(message)
	marshalled, err := l.marshal(messageWithExtra)
	if err != nil {
		fmt.Println(red + "error while marshalling log message" + reset)
		return
	}
	if len(marshalled) == 0 {
		fmt.Println(red + "marshalled is empty" + reset)
		return
	}

	separator := " | "
	marshalledMessage := strings.Split(marshalled, separator)
	if len(marshalledMessage) != 2 {
		fmt.Println(red + "wrong marshalledMessage length" + reset)
		return
	}

	fmt.Println(color + "[" + tier + "] " + marshalledMessage[0] + separator + cyan +
		marshalledMessage[1] + reset)
}

func (l *Log) getMessage(text string, data interface{}) *LogMessage {

	return &LogMessage{
		Data: data,
		Text: text,
	}
}

func (l *Log) addExtraInfo(message *LogMessage) *logMessageWithExtra {

	currentTime := time.Now()

	return &logMessageWithExtra{
		LogMessage: message,
		Datetime:   currentTime,
	}

}

func (l *Log) marshal(messageWithExtra *logMessageWithExtra) (string, error) {

	dateTimeString := messageWithExtra.Datetime.Format(timeFormat)
	text := messageWithExtra.Text
	data, err := json.Marshal(messageWithExtra.Data)

	if err != nil {
		return "", err
	}

	dataString := string(data)

	var result string

	if dataString != "null" {
		result = fmt.Sprintf("%v | %v: %v", dateTimeString, text, dataString)
		return result, nil
	}

	result = fmt.Sprintf("%v | %v", dateTimeString, text)

	return result, nil
}
