# message queue

#### 简介

使用双向链表实现类似redis的list功能，编程语言用golang实现。

#### 调用示例

* 获取实例

```golang
instance := queue.GetDBInstance("db1")
```

* 向左侧追加

```golang
instance.LPush(0)
instance.LPush("a")
instance.LPush(map[string]int{"age":1})
```

* 向右侧追加

```golang
instance.RPush(0)
instance.RPush("a")
instance.RPush(map[string]int{"age":1})
```

* 取左侧第一个节点值

```golang
val, err := instance.LPop()
```

* 取右侧第一个节点值

```golang
val, err := instance.RPop()
```

* 打印当前列表值(从左向右)

```golang
instance.DisplayQueue()
```

* 获取数据库列表

```golang
 dbSlice := instance.GetDBList()
```

* 删除当前数据库

```golang
 err := instance.FlushDB()
```

#### benchmark

运行`go run main.go`即可获得benchmark结果

* 环境：

```
OS
-------------------------------------------------------------------------------------
             ............                keepchen@keepchen-PC 
         .';;;;;.       .,;,.            -------------------------- 
      .,;;;;;;;.       ';;;;;;;.         OS: Deepin 15.11 x86_64 
    .;::::::::'     .,::;;,''''',.       Model: ESC300 G4 
   ,'.::::::::    .;;'.          ';      Kernel: 4.15.0-30deepin-generic 
  ;'  'cccccc,   ,' :: '..        .:     Uptime: 29 minutes 
 ,,    :ccccc.  ;: .c, '' :.       ,;    Packages: 1831 
.l.     cllll' ., .lc  :; .l'       l.   Shell: zsh 5.3.1 
.c       :lllc  ;cl:  .l' .ll.      :'   Resolution: 1920x1080, 1600x900 
.l        'looc. .   ,o:  'oo'      c,   DE: Deepin 
.o.         .:ool::coc'  .ooo'      o.   WM: KWin 
 ::            .....   .;dddo      ;c    Theme: Deepin [GTK2/3] 
  l:...            .';lddddo.     ,o     Icons: Deepin [GTK2/3] 
   lxxxxxdoolllodxxxxxxxxxc      :l      Terminal: deepin-terminal 
    ,dxxxxxxxxxxxxxxxxxxl.     'o,       CPU: Intel i3-7100 (4) @ 3.9GHz 
      ,dkkkkkkkkkkkkko;.    .;o;         GPU: NVIDIA GeForce GT 710B 
        .;okkkkkdl;.    .,cl:.           Memory: 2772MB / 15990MB 
            .,:cccccccc:,.                
                                         ████████████████████████

Golang
-------------------------------------------------------------------------------------
go version go1.15 linux/amd64

```

* 测试结果：

| 操作/数量 | 10000次    | 100000次    | 1000000次    |
| --------- | ---------- | ----------- | ------------ |
| LPush     | 0.827006ms | 12.559893ms | 140.601657ms |
| RPush     | 1.080579ms | 13.262032ms | 130.045112ms |
| LPop      | 0.675541ms | 7.199598ms  | 57.969342ms  |
| RPop      | 0.723422ms | 9.694183ms  | 71.108501ms  |

