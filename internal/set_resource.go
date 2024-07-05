package internal
import(
	"fmt"
	"os"
)

func SetLogFile() (*os.File, error) {
    logFile, err := os.OpenFile("database.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Printf("无法打开日志文件: %v\n", err)
        return nil, err
    }
    return logFile, nil
}