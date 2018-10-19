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

2. 增加/修改ec2 tag
```
URL: /api/v1/aws/ec2
Method: POST
Query: resources<string>   //多个资源时使用","区分, 注意不能有其他空白字符!
Body: map[string]string     //当Key不存在时创建, 存在时覆盖
```
[举例]

3. 删除ec2 tag
URL: /api/v1/aws/ec2
Method: DELETE
Query: resources<string> //多个资源时使用","区分, 注意不能有其他空白字符!
Body: map[string]string  //只有Key,Value都匹配才执行删除操作,而且不删除的时候返回也是正常的！

4. 获取rds信息
```
URL: /api/v1/aws/rds
Method: GET
Query: -
Body: -

5. 获取autoscaling信息
```
URL: /api/v1/aws/autoscaling
Method: GET
Query: -
Body: -

