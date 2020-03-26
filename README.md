# console
控制房间的照明系统

## 类图
![image](https://github.com/roggen-yang/console/doc/class.png)

### 测试场景1：有人进入房间，在延迟时间内离开房间，灯一致不亮
> go run main.go -path conf/test1.txt -in 5s -out 5s

### 测试场景2：有人进入房间，停留时间超过延时时间，灯在延时后点亮
> go run main.go -path conf/test2.txt -in 5s -out 5s

### 测试场景3：有人进入房间，在延时时间内离开，同时又有人进来，灯按第二个进入时的延时时间点亮
> go run main.go -path conf/test3.txt -in 5s -out 5s

### 测试场景4：有人进入房间，延时后灯点亮，然后离开房间，延时关灯期间内有人进来，灯不灭
> go run main.go -path conf/test4.txt -in 5s -out 5s

### 测试场景5：有人进入房间，延时后灯亮，然后离开房间，延时后灯灭
> go run main.go -path conf/test5.txt -in 5s -out 5s