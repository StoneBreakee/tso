# tso
Timestamp Oracle
作用是？Oracle存储时间序列数据库？时间序列是否需要归档？rrd？为用户提供时间序列数据的存储和查询？

读取源代码的步骤
1.查看提交历史，了解项目变更
2.回退到最初版本，查看最原始的功能
3.走读过程，写下中文注释
4.走读过程，记录思路与流程

流程
项目工程
tso
├─bench      工作台？
├─client     客户端？
├─proto      协议定义或是解析？
├─tso-server server主程序 
│  └─server
└─util       工具类？

走读走读代码从tso-server/main.go开始
第一个问题： 什么是go flag？
   flag包负责命令行的解析 类似于linux命令的参数解析
   如 ls dir，则flag包负责将命令行中的dir解析出来传递给ls

go语言执行顺序：import -> const -> var -> init() -> main()