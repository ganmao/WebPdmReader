# WebPdmReader
PowerDesigner's file (PDM) reader tool for web display.

# PowerDesigner WEB方式共享查询工具

## 安装方式
绿色软件，直接拷贝使用

1. 手工创建目录：
/data/idx -- 默认的存放索引文件目录
/data/pdm -- 默认的存放PDM文件目录

2. 将需要分析查询的PDM放入/data/pdm

3. 在PDM管理>索引更新 点击刷新PDM索引按钮，自动刷新
路径均为从conf/app.conf配置文件中设置

4. 支持win，linux，mac多系统
WebPdmReader.exe -- windows
WebPdmReader.linux -- linux
WebPdmReader.mac -- mac

## 使用方式
直接运行：WPdmReader.exe

左边栏搜索框可以直接搜索表名（中英文均可）
表格上的搜索框可以根据内容进行模糊搜索

注意：
    更新索引采用覆盖方式，会全部更新。

使用截图
----
![pdm列表](/img/jt1.PNG)
![表信息列表](/img/jt2.PNG)
![表明细字段列表](/img/jt3.PNG)

## 配置说明
conf/app.conf

appname = WPdmReader    应用名称

httpport = 8080         web端口

logLevel = debug        日志级别

IdxPath = data/idx      索引存放目录

PdmPath = data/pdm      pdm文件存放目录

LogPath = logs          日志目录

## 开发说明
* 开发语言
    GOLANG

* 基础框架
    Beego

* 展示模板
    AdminLTE
    
 ## 联系方式
 zdl0812@163.com

