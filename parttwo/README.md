# 可扩展的分布式系统

分布式系统的好处在于可扩展性，只需要加入新的节点就可以自由扩展集群的性能

##接口和数据存储分离的架构

接口服务层提供了对外的REST接口，而数据服务层则提供数据的存储功能

接口服务和数据服务之间的接口有两种:

    1 数据服务本身提供REST接口
    2 通过MQ消息队列进行通讯

在我的架构中对RabbitMQ的使用分为两种模式，一种模式是向某个Exchange进行一对多的消息群发，另一种模式则是向某个队列进行一对一的消息单发

每一个数据服务节点都需要向所有的接口服务节点通知自身的存在，为此，我们创建一个名为apiServers的Exchange，每一台数据服务节点都会持续向这个Exchange发送心跳消息

所有的接口服务节点在启动之后都会创建一个消息队列来绑定这个Exchange，任何发往这个Exchange的消息都会被转发给它绑定的所有消息队列，也就是说每一个接口服务节点都会收到任意一台数据节点的心跳消息

另外，接口服务需要在收到对象GET请求时定位该对象被保存在哪个数据服务节点上，所以我们还需要创建一个名为dataServers的Exchange

所有的数据服务节点绑定这个Exchange并接收来自接口服务的定位消息，用户该对象的数据服务节点则使用消息单发通知该接口服务节点

对于数据接口来说 和上一版本完全一致，也就是GET和PUT，在这一版本中除了上述的两种方法外，还需要提供一个用于定位的locate接口，用来帮助我们验证架构

    GET /locate/<object_name>

定位的结果

客户端通过GET方法发起对象定位请求，接口服务节点收到该请求后会向后边数据服务层群发一个定位消息，等待数据节点反馈

如果有数据服务节点发回确认消息，则返回该数据节点的地址，如果超过一定时间没有任何反馈，则返回HTTP错误404

##RabbitMQ消息设计

apiServers和dataServers这两个Exchange需要在RabbitMQ上预先创建。

每个接口服务节点在启动后都会创建自己的消息队列并绑定至apiServers Exchange，消息的正文就是该数据服务节点的HTTP监听地址，接口服务节点收到该消息后就会记录这个地址


