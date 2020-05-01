# FreightDistribution

货运配送h5+app 后端

   面向于发货人和货运公司的中介平台，为发货人和货运公司提供了一个交流交易平台，货运公司可以在该平台发布运输信息（支持配送物品种类、支持配送的地点、价格等）。发货人跟据自己的需求进行筛选，选择合适的货运方案，同时支持和货运公司的在线交流，当双方达成一致时，便可进行交易，约定好发货地点和收货地点等信息。货运公司可以根据自己的运送情况对物流信息进行更新。后台具有用户、订单的管理，客服在线服务功能，后台管理员可以对用户、订单进行管理，如进行封号，下线处理，对违规订单进行下架等，在线处理用户发来的各种问题。
   采用前后端分离架构，前端采用uniapp框架, 该框架是一个基于vue开发的，支持一次开发可以编译到多端。我将该项目的前端页面编译成移动端h5和app，用户可以通过app或浏览器进行使用。 
   后端主要分成和数据库进行交互的model层，业务逻辑层（json的序列化与反序列化），和前端交互的路由层以及中间件，工具库（日志，对象存储等）。主要用到的技术有：gin框架+mysql+gorm，REST风格。使用jwt进行用户身份的认证。前后端使用json进行交互。核心功能为基于websocket协议的一对一即时通信，所有客户端连接服务端，通过sync.map存储连接的地址。聊天时用户发送消息到服务端，服务端通过消息中的目标id找到目标连接，并转发消息，实现即时聊天，由于go语言的特性，能支持很高的并发量。
