# 用户相关接口

## 1. 用户注册

### 1.1 接口描述

用户注册

### 1.2 请求方法

POST /v1/register

### 1.3 输入参数

**Body 参数**

| 参数名称         | 必选 | 类型     | 描述    |
|--------------|----|--------|-------|
| stuNum       | 是  | String | 学号    |
| password     | 是  | String | 密码    |
| college      | 是  | String | 学院    |
| major        | 是  | String | 专业    |
| grade        | 是  | String | 年级    |
| name         | 是  | String | 姓名    |
| gender       | 是  | String | 性别    |
| province     | 是  | String | 家乡省份  |
| stuCardPhoto | 是  | String | 学生证照片 |
| phone        | 是  | String | 手机号   |
### 1.4 输出参数

| 参数名称    | 类型     | 描述    |
|---------|--------|-------|
| code    | String | 业务状态码 |
| message | String | 返回消息  |


### 1.5 请求示例

**输入示例**

```bash
curl -X POST \
  -F "stuNum=20999004" \
  -F "password=123Abc!" \
  -F "college=软件学院" \
  -F "major=计算机科学与技术" \
  -F "grade=大四" \
  -F "name=淮屿" \
  -F "gender=男" \
  -F "province=辽宁" \
  -F "stuCardPhoto=@./cardPhoto.jpg" \
  -F "phone=13701010101" \
  http://localhost:8888/v1/user/register
```
**输出示例**

```json
{
  "code": 100001,
  "message": "OK"
}
```