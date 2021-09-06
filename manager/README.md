 # MOCK
 是一个支持logicrec项目mock数据的一个模块

 # 支持的功能列表
 1. 用于基于map实现的分布式server同步RPC调用, mock服务
    1. 使用一个并发map做为一个服务注册中心，
    2. key为服务唯一标记，value是一个 server 对象指针列表
    3. 调用时，直接通过key获取server对象直接调用