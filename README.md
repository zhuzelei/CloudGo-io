## 概述
-----
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

## 任务要求
-----
编程 web 应用程序 cloudgo-io。 请在项目 README.MD 给出完成任务的证据！</br>
**基本要求**</br>
1.支持静态文件服务</br>
2.支持简单 js 访问</br>
3.提交表单，并输出一个表格</br>
4.对 /unknown 给出开发中的提示，返回码 5xx</br>

## 实验过程
-----
1.首先克隆所需要的相应的库

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181116181147264.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl8zNjMyODM4MA==,size_16,color_FFFFFF,t_70)

库的地址分别是
```
github.com/codegangsta/negroni
github.com/gorilla/mux   
github.com/unrolled/render
github.com/spf13/pflag
```
克隆指令即`git clone https://github.com/.../...`

2.熟悉基本流程
详细的介绍和入门操作见[潘老师的这篇博客](https://blog.csdn.net/pmlpml/article/details/78539261)，在这里我就不做赘述了


## 实验结果
-----
静态文件访问
</br>
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181116181810806.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl8zNjMyODM4MA==,size_16,color_FFFFFF,t_70)

简单js访问
</br>
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181116181842342.png)


提交表单输出表格
</br>
![在这里插入图片描述](https://img-blog.csdnimg.cn/201811161819184.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl8zNjMyODM4MA==,size_16,color_FFFFFF,t_70)
</br>
![在这里插入图片描述](https://img-blog.csdnimg.cn/20181116181934233.png)

unknown返回5xx代码
</br>
![在这里插入图片描述](https://img-blog.csdnimg.cn/2018111618200570.png)


代码相关具体内容可以查看[我的github](https://github.com/zhuzelei/CloudGo-io)
