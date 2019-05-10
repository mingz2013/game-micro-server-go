robot进程，整个项目的客户端，从connector进程接入，可用于压力测试



server是入口

robot，负责引入各个game项目里的robot模块，分发调用，管理robot goroutine



robot, 只是一个综合，用其他的模块的客户端，组成整个robot的流程