# QuickMock
&emsp;&emsp;通过配置json，快速构建http请求接口，方便前端程序猿快速调用接口。

## 配置文件说明
```json
{
    "port":"8080",
    "routers":[
        {
            "name":"/",
            "method":"get,post",
            "response":"{'code':200,'data':'this url is /'}"
        },
        {
            "name":"/hello",
            "method":"get",
            "status_code":"302",
            "response":"https://github.com/champly/quick-mock"
        }
    ]
}
```
- port : 监听的端口号
- routers : 路由配置
> - name : 路由名称
> - method : 请求方式(不区分大小写),可以多种方式,用分号分割
> - status_code : 返回状态码，如果是302或者301则跳转到response中
> - response : 返回内容(会自动把**单引号**替换成**双引号**)
***
## 需要注意的是
1. 路由如果不存在，则返回状态码404;
2. 配置文件必须和执行文件放在同一个目录下面;
3. 配置文件名字必须是config.json,格式为json格式;
4. bin文件夹里面是编译之后的文件，直接下载就可以运行.
***
&emsp;&emsp;如果有什么问题或者建议的话可以联系我的邮箱:champly1993@gmail.com，如果有什么好的建议或者功能都可以通过邮件告诉我，谢谢你的支持!