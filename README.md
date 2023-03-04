



## MVC架构

- 控制层(Controller)：处理请求和响应。通常控制层是**应用程序的入口点**。在Web应用程序中，它们会响应来自客户端的HTTP请求，并将它们**传递到适当的服务层**，以便进行处理。控制器负责根据请求的类型（GET，POST等）**确定要执行的操作**，并返回响应。
- 服务层(Service)：包含应用程序的**主要业务逻辑**。服务层负责执行控制器所请求的操作，并将其**委托给适当的数据访问层**，以便从持久化存储中检索和保存数据。
- 数据访问层(DAO)：**与数据存储层交互**，包括数据库和其他数据存储介质。数据访问层封装了与数据存储相关的所有细节，并为服务层提供简单而统一的接口，以便查询和更新数据。



### 目录下的相关`.go`文件

- mian.go：调用Gin框架，启动服务

- router.go：将处理程序绑定到路由器路径上，客户端访问路由器路径，绑定后，相当于调用处理程序



### config文件夹

- config.go

### controller文件夹



### service文件夹



### dao文件夹



