#### 说明

virtualbox里windows剪切板到linux宿主机剪切板的同步老是失效，

最开始想通过网络来传数据，后来发现每次虚拟机剪切板失效的时候虚拟机与宿主机（10.0.2.2）之间的网络也是不通的，

或者反过来正是因为虚拟机与宿主机网络不通才导致剪切板的同步失效，

但我发现共享目录文件仍旧是有效的，于是有了这个基于共享目录文件来同步虚拟机windows剪切板与linux宿主机剪切板。

使用前最好把virtualbox的共享剪切板关掉。

#### Linux 宿主机



运行

```
./host.sh
```

linux 下编译 host

```
go get github.com/atotto/clipboard
go get github.com/syyongx/php2go

go build -o host-linux host.go
```


#### Win 虚拟机


双击运行

```
guest.bat
```

win下编译 guest

```
go get gopkg.in/Knetic/govaluate.v3
go get github.com/lxn/walk
go get github.com/lxn/win
go get golang.org/x/sys
go get golang.org/x/image
go get github.com/syyongx/php2go

go build -o guest.exe guest.go winclipboard.go
```