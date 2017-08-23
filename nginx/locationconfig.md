测试location内部的配置：
proxy_pass只能设置一条
proxy_set_header 同一个变量可以设置多次。nginx会把配置的所有请求头都转发给后端服务器。
后端服务器会取哪个值呢？
例如：
proxy_set_header X-Forwarded-For "1.1.1.1";
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
同时设置这两个头，客户端获取到的X-Forwarded-For的值是什么呢？
发现获取到的值为一个数组，按先后顺序记录了X-Forwarded-For的值。

惊喜不惊喜，意外不意外。


ps：使用beego框架搭建的web服务如何获取请求的请求头信息？
直接在controller中就可以获取到相应信息
header:=c.Ctx.Request.Header
返回的header是一个字典类型，可以直接取某个key所对应的值。
