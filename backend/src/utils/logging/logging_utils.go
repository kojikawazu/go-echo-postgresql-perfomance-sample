package logging_utils

import (
	"log"
	"runtime"
	"time"
)

// getFunctionName 呼び出し元の関数名を取得
func getFunctionName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	return runtime.FuncForPC(pc).Name()
}

// LogStart 処理の開始ログを記録
func LogStart() time.Time {
	functionName := getFunctionName(2)
	log.Printf("[START] %s", functionName)
	return time.Now()
}

// LogEnd 処理の終了ログと実行時間を記録
func LogEnd(start time.Time) {
	functionName := getFunctionName(2)
	duration := time.Since(start)
	log.Printf("[END] %s took %v", functionName, duration)
}

// LogError エラー時の終了ログを記録
func LogError(start time.Time, err error) {
	functionName := getFunctionName(2)
	duration := time.Since(start)
	log.Printf("[ERROR] %s took %v - Error: %v", functionName, duration, err)
}

// LogInfo 通常の情報ログを記録
func LogInfo(message string, details ...interface{}) {
	functionName := getFunctionName(2)
	log.Printf("[INFO] %s: %s %+v", functionName, message, details)
}
