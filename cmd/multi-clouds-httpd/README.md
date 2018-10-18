# clouds-proxy-httpd

## 接口文档


| id | comment | URL | method | query | body | return |
| ------ | ------ | ------ |
| ss |

0. 返回结果
 ```
 Type: json
 example:
 {
    "code": 0, <int>
    "msg": "", <string>
    "err": "",  <string>
    "data": interface{}, <...>任意类型，99%的情况为列表.
}
```

注意: 必须首先校验code和err, 当code=0，并且err=""或"<nil>"时，data才是正常返回的结果。

**其他情况，data可能会有数据，但是不保证正确以及完整！！！**


1. 获取ec2
```
URL: /api/v1/aws/ec2
Method: GET
Query: -
Body: -
```

2. 修改ec2 tag
```
URL: /api/v1/aws/ec2
Method: POST
Query: resources<string>, 多个资源时使用","区分, 注意不能有其他空白字符!
Body: map[string]string
```
[举例]
