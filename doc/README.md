# API说明

关于uoj接口的调用说明

## 概览

- 为方便描述，Host 地址为 localhost 。建议在客户端设置 APP_HOST 常量。
- 列表项每页显示 50 条记录。默认为前 50 条。除第一页以外，获取其他页面需要传入 page 参数。如 localhost/u?page=2 会显示第 51 ~100 条记录。
- 如果不做特殊说明，所有的发送数据均使用POST方法。
- 如果不做特殊说明，所有的操作行为会返回一个 JSON，boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。
- 除规定的访问方式以外的其他非法访问行为按攻击服务器处理。

## 详细说明

### 基本格式

localhost/u/:uid

- localhost 为主机名，建议使用 APP_HOST 常量
- 类 :uid 以冒号开始的字段表示参数，使用时请替换为实际数据。如 localhost/u/zhr 表示获取zhr的信息。
 

### 获取数据

获取用户信息

```
localhost/u 		// 获取用户列表

localhost/u/:uid 	// 获取某个用户的信息。
````

获取题目信息

```
localhost/p 			// 获取题目列表

localhost/p/:pid		// 获取某个题目的信息

localhost/p_tag/:tag	// 获取特定标签的题目
```

获取动态信息

```
localhost/m 			// 获取所有人的动态

localhost/m_user/:uid	// 查看某人动态。

localhost/m_tag/:tag	// 获取特定标签的动态

localhost/m_pid/:pid	// 与某道题相关的讨论
```

>已登陆用户查看自己的动态时，服务器会返回所有消息，查看他人动态会返回公开消息。无需客户端验证。

获取作业信息

```
localhost/t/ 		// 获取作业列表

localhost/t/:tid/	// 获取指定作业的信息

localhost/t/record	// 获取指定作业的提交记录
```

### 发送数据

用户行为

- 登陆 OJ

		localhost/u_login

	- 使用 POST 方法
	- 参数列表： 
		- userid 	用户ID
		- password 	密码
	- 返回信息：
		- 信息以 JSON 格式返回。boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。
		- "OK" = true 表示登陆成功，Info 字段以 JSON 对象的格式返回用户信息。
		- "OK" = false表示登陆失败，Info 字段返回错误信息代码。

- 退出登陆

		localhost/u_logout

	- 使用 POST 方法
	- 参数列表：没有参数
	- 返回信息：
		- 信息以 JSON 格式返回。boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。
		- "OK" = true 表示退出成功，Info 字段返回退出的用户。
		- "OK" = false表示退出失败，Info 字段返回错误信息代码。

- 用户注册

		localhost/u_join

	- 使用 POST 方法
	- 参数列表

		| 参数| 描述 | 备注 |  
		|:-----:|:------|:-----|
		| userid	|  | 必需。不可更改。|
		| nickname  |  默认为 userid。 ||
		| password  |    | 必需。|
		| password2 |    | 必需。|
		| emailaddr |   | 必需。需要验证邮箱。|
		| studentid |    | 必需。不可更改。|
		| realname  |    | 必需。不可更改。|

	- 返回信息
		- 信息以 JSON 格式返回。boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。
		- "OK" = true 表示注册成功。Info 字段返回用户信息。
		- "OK" = false表示注册失败，Info 字段返回错误信息代码：
			- 201：信息不完整
			- 202：信息不匹配：信息格式有误或两次密码不一致
			- 210：用户已存在

题目

- 提交题目

		localhost/p_submit

	- 使用 POST 方法
	- 参数列表

		| 参数| 描述 | 备注 |  
		|:-----:|:------|:-----|
		| userid	|  | 必需。不可更改。|
		| nickname  |  默认为 userid。 ||
		| password  |    | 必需。|
		| password2 |    | 必需。|
		| emailaddr |   | 必需。需要验证邮箱。|
		| studentid |    | 必需。不可更改。|
		| realname  |    | 必需。不可更改。|

	- 返回信息

- 上传题目

		localhost/p_add

	- 使用 POST 方法
	- 参数列表

		| 参数| 描述 | 备注 |  
		|:-----:|:------|:-----|
		| title	|  | 必需。不可更改。|
		| content  |  默认为 userid。 ||
		| tag  |    | 必需。|
		| password2 |    | 必需。|
		| emailaddr |   | 必需。需要验证邮箱。|
		| studentid |    | 必需。不可更改。|
		| realname  |    | 必需。不可更改。|

	- 返回信息

- 修改题目

动态

- 发出一条动态

作业

- 提交作业题目
- 添加作业
- 修改作业
