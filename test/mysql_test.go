package test

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"io"
	"os"
	"sync"
	"testing"
)

var DB *sql.DB

var BatchCount int64

func TestMysql(t *testing.T) {
	// 初始化
	var err error
	if err = initMysql(); err != nil {
		fmt.Println("init failed", err)
		return
	}
	var wg sync.WaitGroup
	BatchCount = 100
	ch := make(chan string)
	quitCh := make(chan struct{})
	go execSql(ch, quitCh, &wg)
	if err = readFile("F:\\GoWorkspace\\zvos-edge-command-control\\download\\9999_1.sql", ch, quitCh, &wg); err != nil {
		fmt.Println("readfile err:", err)
		return
	}
	wg.Wait()
	DB.Close()
}
func initMysql() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s&multiStatements=true", "root", "zhjc2022mysql", "10.71.6.157", 30451, "utf8")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 设置参数
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	// DB.SetConnMaxLifetime(30 * time.Second)

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func readFile(fileName string, ch chan string, quitCh chan struct{}, wg *sync.WaitGroup) error {
	fileHandle, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	var index int64
	wg.Add(1)
	reader := bufio.NewReader(fileHandle)
	for {
		var strTotal string
		for index = 0; index < BatchCount; index++ {
			str, err := reader.ReadString(';')
			if err == io.EOF {
				zap.L().Debug("read end")
				if strTotal != "" {
					ch <- strTotal
				}
				goto Loop
			}
			strTotal += str
		}
		select {
		case <-quitCh:
			goto Loop
		case ch <- strTotal:
		}
	}
Loop:
	close(ch)
	return nil
}

func execSql(ch chan string, quitCh chan struct{}, wg *sync.WaitGroup) {
	var count int64
	beishu := int64(1)
	for {
		if data, ok := <-ch; ok {
			// 执行失败如何处理
			_, err := DB.Exec(data)
			if err != nil {
				fmt.Println("error:", count, data, err)
				break
			}
			count += BatchCount
			if count == beishu*100 {
				fmt.Println("count", count)
				beishu += 1
			}
		} else {
			break
		}
	}
	wg.Done()
	defer func() {
		quitCh <- struct{}{}
	}()
}
