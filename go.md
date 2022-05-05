# GoRoutine

## Goroutine 简介

#### go 的协程，类似于多线程的操作

* golang 对协程的处理
* 协程 (co-routine) --> goroutine
* 特点
    * 灵活调度，可以经常切换
    * 内存几 kb，可以大量创建

#### GMP

* G --> goroutine 指的是每个一 go 程序
* P --> processor 一个处理器，用来处理 goroutine 协程，可以通过 GOMAXPROCS 来进行设置个数
* M --> thread 系统线程，如果想要执行的话，需要获取 p，然后由 p 来执行 g

```text
    --->  G, G, G, G ---> 全局队列，即正在等待执行的 goroutine 程序

	[G,G,G]	   [G,G,G]     p 的本地队列，成为 LocalP，等待运行的队列
	   |	      |
       P		  P	···    Processor，每个 P 都对应一个 LocalP 队列
       |		  |
	  M0		 M1        内核线程
	--------操作系统调度器----
		cpu	  cpu 	cpu	    硬件 cpu 核心
```

## 调度器的设计策略

#### 复用线程

* work stealing 机制（偷取）
  * 当某个 M1 在执行某个 G 的时候，另外一个 M2 没有 G 可执行，但 M1 的 LocalP 中有很多个 G，那么 M2 会从 M1 的 LocalP 队列中获取一个 G 进行执行，被称为偷取
* hand off 机制，（握手）
  * 当 M1 在执行 G1 时被阻塞了，那么操作系统会唤醒一个空闲线程 M3，执行 M1 的 P 以及 LocalP，即 M3 全面接管了 M1 之前的任务，同时，被阻塞的 G1，会直接绑定在 M1 上等待执行

#### 利用并行

* GOMAXPROCS 限定 P 的个数 = CPU 核数 / 2
* 可以空闲 CPU 交给其他线程来进行使用

#### 抢占

* 在 co-routine 中，如果某个协程 c1 被 CPU 绑定，那么 协程 c2 只能等待 c1 释放 CPU，才能进行执行
* 而在 goroutine，如果某个 G1 被 CPU 绑定，那么 G2 最多只等待 10ms, 如果 10ms 后 G1 还不释放 CPU，那么就进行抢占

#### 全局 G 队列

* 基于 work stealing 机制进行改进
* 即，在  work stealing 机制 中 M2 会从 M1 的 LocalP 中进行偷取，而此种调度情况下，M2 会从 全局 P 队列中进行偷取

## Channel

#### 无缓冲的 channel

* 第一步，两个 goroutine 都到达观点，但都没有进行发送或接受
* 第二步，发送方 goroutine 发送数据，此时，发送方会被锁住，直到接收方 goroutine 读取数据
* 第三步，接收方 goroutine 接收数据，接收方也被锁住，进行数据交换
* 第四步，数据发送完成，接收完成，两个 goroutine 继续执行

#### 有缓冲的 channel

* 第一步，发送方 goroutine 将数据放入管道中，继续执行，也可以继续放数据进入管道，相当于一个缓冲区，直到缓冲区放满，则阻塞发送
* 第二步，接收方 goroutine 从管道中获取数据，两边互不干扰，空管道阻塞读
* 第三步，两个 goroutine 可以无需等待，进行自己的操作

#### 关闭 channel

* channel 不像一个文件一样需要经常关闭，只有确定不会发送任何数据了，才会尝试关闭 channel
* 关闭 channel 后，无法继续向 channel 中写入数据，会引发 panic 错误，导致接收即可返回 0
* 关闭 channel 后可以继续从 channel 中读取数据
* 对于 nil channel，无论收发都会被阻塞
  * nil channel 指的是未初始化的 channel `var c chan int`，即未进行 make 初始化

# Go Module

#### 什么是 Go Module

* Go Module 是 Go 语言的依赖解决方案，发布于 Go1.11，成长与 1.12，丰富于 1.13，在 1.14 正式生产使用
* Go Module 目前集成于 Go 的工具链中，只要安装 Go，自然而然也可以使用 Go Module，而 Go Module 的出现也解决了在 Go 1.11 之前的几个问题
  * Go 长久以来的依赖管理问题
  * 淘汰现有的 ”GOPATH“ 的使用模式
  * 统一社区的其他依赖工具

#### go mod 命令

| 命令            | 作用                           |
| --------------- | ------------------------------ |
| go mod init     | 生成 go.mod 文件               |
| go mod download | 下载 go.mod 文件中所指明的依赖 |
| go mod tidy     | 整理现有的依赖                 |
| go mod graph    | 查看现有的依赖结构             |
| go mod edit     | 编辑 go.mod 文件               |
| go mod vendor   | 导出项目所有依赖的 vendor 目录 |
| go mod verify   | 校验一个模块是否被篡改过       |
| go mod why      | 查看为什么需要依赖某模块       |

#### go mod 环境变量

* `go env 查看环境变量`

  ```shell
  $ go env
  GO111MODULE="auto"
  GOPROXY="https://proxy.golang.org,direct"
  GONOPROXY=""
  GOSUMDB="sum.golang.org"
  GONOSUMDB=""
  GOPRIVATE=""
  ···
  ```

* GO111MODULE
  * go 语言提供了 GO111MODULE 来作为 Go Module 的开关，允许设置以下参数
    * auto 只要包含了  go.mod 就默认开启
    * on 启用 Go Module
    * off 禁用 Go Module
  * 通过 `$ go env -v GO111MODULE=on` 来进行设置
* GOPROXY
  * 设置 Go 的拉取镜像代理，去镜像站拉取模块
  * 主要有以下几个
    * 阿里云 `https://mirrors.aliyun.com/goproxy/`
    * 七牛云 `https://goproxy.cn.direct`
  * `$ go env -w GOPROXY=https://goproxy.cn,direct` 此处 direct 指的是如果设置的源中没有的话，会去默认源中拉取

* GOSUMDB
  * 用于在模块拉去时，保证拉去版本数据未经过篡改，如果发现不一致或有篡改，那么会终止拉取动作
  * 默认值是 sum.golang.org
  * `$ go env -w GOSUMDB=off` 设置关闭校验，但不推荐
* GONOPROXY, GONOSUMDB, GOPRIVATE
  * 私有仓库所设置的源都不会经过 proxy 去拉取且不会校验，而是需要自己配置，可以用通配符
  * `$ go env -w GOPRIVATE="git.example.com,github.com/RexJoush/zinx"`
  * `$ go env -w GOPRIVATE="*.example.com`

