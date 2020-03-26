package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/roggen-yang/console/peoples"
	"github.com/roggen-yang/console/room"

	"github.com/roggen-yang/console/light"
)

var pathString string
var inDuration time.Duration
var outDuration time.Duration

func main() {
	// 解析命令行参数
	flag.StringVar(&pathString, "path", "", "Set file path")
	flag.DurationVar(&inDuration, "in", 1*time.Second, "Set in room light on delay time")
	flag.DurationVar(&outDuration, "out", 1*time.Second, "Set out room light off delay time")
	flag.Parse()

	// 打开文件
	fi, err := os.Open(pathString)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	// 创建房间
	peo := peoples.NewPeoples()
	li := light.NewLight(inDuration, outDuration)
	room := room.NewRoom(peo, peo, li, li)

	// 创建reader buffer
	br := bufio.NewReader(fi)
	for {
		// 逐行读取
		s, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		// 类型转换string to int
		n, err := strconv.Atoi(string(s))
		if err != nil {
			fmt.Println(err)
			break
		}
		// 根据n的大小判断人的行为（进或者出）
		if n > 0 {
			err = room.ComeIn(n)
			if err != nil {
				fmt.Println(err)
				break
			}
		} else if n < 0 {
			err = room.GetOut(n)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		//fmt.Printf("Input: %v, counter: %v, status: %v \n", n, room.GetCounter(), room.GetLightStatus())
		time.Sleep(1 * time.Second)
		fmt.Println(room.GetLightStatus())
	}
}
