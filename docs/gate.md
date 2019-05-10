1 管理用户连接session
2 实现response，request





只有gateway进程服务，需要监听各个不同协议server，
其他服务，只需要连接到消息队列，收发消息
连接到数据库，读写数据





接收用户连接，协议，tcp，ws，http


用户每次连接，用户需要携带token，用于验证用户状态

服务器从数据库验证token，验证成功连接成功，绑定user，否则断开连接


有状态的user conn开始通信



