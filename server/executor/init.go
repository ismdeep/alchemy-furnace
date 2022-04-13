package executor

import (
	"encoding/json"
	"errors"
	"github.com/ismdeep/rand"
	"time"
)

// EOF 结尾标识
const EOF = "EOF_e1cd5d879b3cec1a52d07ab8b39a61e8"

// TypeStdout 标准输出
const TypeStdout = 0

// TypeStderr 标准错误
const TypeStderr = 1

// ExeLog 执行记录
type ExeLog struct {
	Type     int       `json:"type"`      // 内容类型，0标准输出，1标准错误输出
	OutputAt time.Time `json:"output_at"` // 输出时间
	Line     string    `json:"line"`      // 行内容
}

// Executor 执行器
type Executor struct {
	Logs      []*ExeLog
	MsgQueues map[string]chan *ExeLog
	EOF       bool
}

var exeLogCache map[string]*Executor

func init() {
	exeLogCache = make(map[string]*Executor)
}

// GenerateExecutor 生成执行器
func GenerateExecutor() string {
	exeID := rand.HexStr(32)
	exeLogCache[exeID] = &Executor{
		Logs:      make([]*ExeLog, 0),
		MsgQueues: make(map[string]chan *ExeLog),
		EOF:       false,
	}
	return exeID
}

// GenerateListener 生成监听器
func GenerateListener(exeID string) (chan *ExeLog, string, error) {
	executor, ok := exeLogCache[exeID]
	if !ok {
		return nil, "", errors.New("not found")
	}
	listenerID := rand.HexStr(32)
	c := make(chan *ExeLog, len(executor.Logs)+65535)
	for _, v := range executor.Logs {
		c <- v
	}
	executor.MsgQueues[listenerID] = c
	return c, listenerID, nil
}

// DestroyListener 销毁监听器
func DestroyListener(exeID string, listenerID string) error {
	executor, ok := exeLogCache[exeID]
	if !ok {
		return errors.New("executor not found")
	}

	c, ok := executor.MsgQueues[listenerID]
	if !ok {
		return errors.New("listener not found")
	}

	close(c)
	delete(executor.MsgQueues, listenerID)
	return nil
}

// PushMsg 推送消息
func PushMsg(exeID string, t int, content string) {
	executor, ok := exeLogCache[exeID]
	if !ok {
		return
	}

	if content == EOF {
		executor.EOF = true
		for _, q := range executor.MsgQueues {
			close(q)
		}
		return
	}

	l := &ExeLog{
		Type:     t,
		OutputAt: time.Now(),
		Line:     content,
	}
	executor.Logs = append(executor.Logs, l)
	for _, q := range executor.MsgQueues {
		q <- l
	}
	return
}

// DumpLog 转储日志
func DumpLog(exeID string) (string, error) {
	executor, ok := exeLogCache[exeID]
	if !ok {
		return "", errors.New("executor not found")
	}

	bytes, err := json.Marshal(executor.Logs)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
