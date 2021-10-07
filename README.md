# vim-swp-exp

#### 介绍

vim-swp-exp 是一款swp文件泄漏利用工具。

#### 原理

当用vim编辑器打开一个文件时候，vim会自动生成一个.swp文件做备份。如果vim程序正常退出，vim会删除该swp文件。如果vim程序非正常退出，vim自然没有能力删除掉该文件。下次打开vim时，vim会询问用户是否通过swp文件恢复未保存的文本。  
在CTF比赛和真实的渗透测试中，运维人员可能会用vim在服务器上编辑站点的.php文件。只要用vim打开文件，就会有.swp文件生成。本程序每0.5s请求一次目标的url的swp文件，一旦检测到存在swp文件，就将该文件下载下来。使用者可以 vim 一个同名文件，然后按R键让vim根据swp文件恢复文件源码。从而将黑盒测试转变为白盒测试。  
swp文件下载完成后，程序会继续监控，一旦线上的swp文件发生变化，程序也会更新本地的swp文件。

#### 截图

![avatar](https://img-blog.csdnimg.cn/20211007203021543.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)
批量监控:  
![avatar](https://img-blog.csdnimg.cn/20211007202922339.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)

#### 特性

1. -u 监控单个url
2. -f 批量监控url  

#### 安装方法

```
go build
./vim-swp-exp
```