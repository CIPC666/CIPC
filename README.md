[![pipeline status](https://api.travis-ci.org/33cn/plugin.svg?branch=master)](https://travis-ci.org/33cn/plugin/)
[![Go Report Card](https://goreportcard.com/badge/github.com/33cn/plugin?branch=master)](https://goreportcard.com/report/github.com/33cn/plugin)


# 基于 chain33 区块链开发 框架 开发的 CIPC公有链系统


### 编译

```
git clone https://github.com/CIPC666/CIPC $GOPATH/src/github.com/CIPC666/CIPC
cd $GOPATH/src/github.com/CIPC666/CIPC
go build -i -o cipc
go build -i -o cipc-cli github.com/CIPC666/CIPC/cli
```

### 运行
拷贝编译好的cipc, cipc-cli, cipc.toml这三个文件置于同一个文件夹下，执行：
```
./cipc -f cipc.toml
```
