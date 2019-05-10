daemon-go

守护进程，整个后端启动的第一个进程，

负责启动其他进程
负责管理所有进程状态, 和其他进程间有心跳检测


有客户端可以连接管理，REST API




---



这里，是用来管理整个服务器的工具


管理进程启动，管理进程调度



用k8s的api进行管理最好了




---


每一个进程是一个actor，都有一个信箱channel

actor是异步模型，需要消除同步操作



---

actor实现同步


每一个actor，是一个进程，每个进程都有其标识符，通过标识符可以生成相应的channel，

消息里面携带发送actor的标识符，接收者可根据标识符去回复指定的进程，指定的进程收到相应消息，触发保存的回调
对于指定进程的回复，一个是需要按照相应顺序返回，则可找到相应的回调（这里可以参考对于redis调用的实现，redis是单线程，所以返回的顺序和请求的顺序是一一对应的）
一个是需要在msg中添加相应的标识，则可找到相应的回调,（这里可以参考http协议的实现，还有rpc的实现）

---






首先，
启动redis docker，并尝试订阅 redis channel，每一个进程订阅一个channel到redis，

订阅channel成功后，


启动一个co进程，启动的时候，传入参数，{redis地址，需要订阅的频道，manager的频道，}，co启动成功后，定时向manager发心跳


启动一个robot进程，传入参数，{redis地址，需要订阅的频道，manager的频道，}，robot启动成功后，定时向manager发心跳

启动一个table进程，传入参数，{redis地址，需要订阅的频道，manager的频道，}，table启动成功后，定时向manager发心跳


给robot发命令，告诉其去连接co，并进行游戏，传入参数，{co的频道，}

co接收到robot登陆命令，查找数据库，验证user，去manager请求一个桌子进程的channel，请求到之后，给table的channel转发robot的信息

manager收到一个桌子请求后，找到存活的可用的桌子进程，并返回channel给co



---


可通过docker API来管理，启动docker