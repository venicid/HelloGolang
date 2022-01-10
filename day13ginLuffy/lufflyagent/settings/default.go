package settings

import (
	"fmt"
	"os"
)

func DefaultHandle(){
	fmt.Printf("未知命令 : %v\n", os.Args[1])
	fmt.Printf("使用说明 : %s [start|stop]\n", os.Args[0]) // return the program name back to %s
	os.Exit(1)
}