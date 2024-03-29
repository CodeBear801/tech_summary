
## TCP/IP layers
```C++
/*
    ┌────------────┐┌─┬─┬─-┬─┬─-┬─┬─-┬─┬─-┬─┬─-┐
　　│　　　　　　　　││Ｄ│Ｆ│Ｗ│Ｆ│Ｈ│Ｇ│Ｔ│Ｉ│Ｓ│Ｕ│　│
　　│　　　　　　　　││Ｎ│Ｉ│Ｈ│Ｔ│Ｔ│Ｏ│Ｅ│Ｒ│Ｍ│Ｓ│其│
　　│第四层，应用层　││Ｓ│Ｎ│Ｏ│Ｐ│Ｔ│Ｐ│Ｌ│Ｃ│Ｔ│Ｅ│　│
　　│　　　　　　　　││　│Ｇ│Ｉ│　│Ｐ│Ｈ│Ｎ│　│Ｐ│Ｎ│　│
　　│　　　　　　　　││　│Ｅ│Ｓ│　│　│Ｅ│Ｅ│　│　│Ｅ│它│
　　│　　　　　　　　││　│Ｒ│　│　│　│Ｒ│Ｔ│　│　│Ｔ│　│
　　└───────------─┘└─┴─┴─-┴─┴─-┴─┴─-┴─┴─-┴─┴-─┘
　　┌───────-----─┐┌─────────-------┬──--------─────────┐
　　│第三层，传输层　││　　　ＴＣＰ　　　│　　　　ＵＤＰ　　　　│
　　└───────-----─┘└────────-------─┴──────────--------─┘
　　┌───────-----─┐┌───----──┬───---─┬────────-------──┐
　　│　　　　　　　　││　　　　　│ＩＣＭＰ│　　　　　　　　　　│
　　│第二层，网间层　││　　　　　└──---──┘　　　　　　　　　　│
　　│　　　　　　　　││　　　　　　　ＩＰ　　　　　　　　　　　 │
　　└────────-----┘└────────────────────-------------─-┘
　　┌────────-----┐┌─────────-------┬──────--------─────┐
　　│第一层，网络接口││ＡＲＰ／ＲＡＲＰ　│　　　　其它　　　　　│
　　└────────------┘└─────────------┴─────--------──────┘


*/
```

　TCP/IP协议被组织成四个概念层，其中有三层对应于ISO参考模型中的相应层。ICP/IP协议族并不包含物理层和数据链路层，因此它不能独立完成整个计算机网络系统的功能，必须与许多其他的协议协同工作。 
　　TCP/IP分层模型的四个协议层分别完成以下的功能：  
　　第一层:网络接口层  
　　包括用于协作IP数据在已有网络介质上传输的协议。实际上TCP/IP标准并不定义与ISO数据链路层和物理层相对应的功能。相反，它定义像地址解析协议(Address Resolution Protocol,ARP)这样的协议，提供TCP/IP协议的数据结构和实际物理硬件之间的接口。  
　　第二层:网间层  
　　对应于OSI七层参考模型的网络层。本层包含IP协议、RIP协议(Routing Information Protocol，路由信息协议)，负责数据的包装、寻址和路由。同时还包含网间控制报文协议(Internet Control Message Protocol,ICMP)用来提供网络诊断信息。  
　　第三层:传输层  
　　对应于OSI七层参考模型的传输层，它提供两种端到端的通信服务。其中TCP协议(Transmission Control Protocol)提供可靠的数据流运输服务，UDP协议(Use Datagram Protocol)提供不可靠的用户数据报服务。  
　　第四层:应用层  
　　对应于OSI七层参考模型的应用层和表达层。因特网的应用层协议包括Finger、Whois、FTP(文件传输协议)、Gopher、HTTP(超文本传输协议)、Telent(远程终端协议)、SMTP(简单邮件传送协议)、IRC(因特网中继会话)、NNTP（网络新闻传输协议）等。  



<img src="https://www.guru99.com/images/1/093019_0615_TCPIPModelW3.png" alt="osi vs tcpip" width="600"/>
<br/>
