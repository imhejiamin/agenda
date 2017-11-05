# agenda

## agenda 概述
命令行实用程序并不是都象 cat、more、grep 是简单命令。go 项目管理程序，类似 java 项目管理 maven、Nodejs 项目管理程序 npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能变化、添加新命令等。因此，符合 OCP 原则 的设计是至关重要的编程需求。

## 任务目标
熟悉 go 命令行工具管理项目。

合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda。

使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码。

项目部署在 Github 上，合适多人协作，特别是代码归并。

支持日志（原则上不使用debug调试程序）。

## agenda需求描述
agenda help ：列出命令说明

agenda register -uUserName --password pass -email=a@xxx.com ：注册用户

agenda help register ：列出 register 命令的描述

agenda cm ... : 创建一个会议

原则上一个命令对应一个业务功能

使用 json 存储 User 和 Meeting 实体

当前用户信息存储在 curUser.txt 中

cmd ：存放命令实现代码

entity ：存放 User 和 Meeting 对象读写与处理逻辑

其他目录 ： 自由添加

使用 log 包记录命令执行情况


## agenda命令

**you can use "agenda [command] help" to get the usage of each parameter of the command.**
  ###   register        
  Register user.  
  
  ###   login          
  Login.
  
   ###  logout          
  Logout.
  ###  addParticipator 
  Add your own meetings' participators.

  ###   clear           
  Clear all meetings you attended or created.

   ### createMeetings  
  Create meetings.
  
  ###   deleteAMeeting  
  Cancel your own meeting by specifying title name.
  
  ###   deleteUser      
  Delete your account.
  
 
  
  ###   listMeetings    
  List all of your own meetings during a time interval.
  
  ###   listUser        
  list all users.
  
 ###   help            
  Help about any command.
  
  ###   quit            
  Quit meetings.
  

  
  ###   rmParticipator 
  Remove your own meetings' participators.

## 部分功能测试
![1](https://github.com/imhejiamin/agenda/blob/master/images/help.png)

![12](https://github.com/imhejiamin/agenda/blob/master/images/login.png)

![13](https://github.com/imhejiamin/agenda/blob/master/images/createmeeting.png)

![14](https://github.com/imhejiamin/agenda/blob/master/images/delete.png)

![5](https://github.com/imhejiamin/agenda/blob/master/images/list.png)

![16](https://github.com/imhejiamin/agenda/blob/master/images/quit.png)

![167](https://github.com/imhejiamin/agenda/blob/master/images/deleteuser.png)


















