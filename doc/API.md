# API说明

关于uoj接口的调用说明

## 概览

- 为方便描述，Host 地址为 localhost 。建议在客户端设置 APP_HOST 常量。
- 列表页每页显示 50 条记录。默认为前 50 条。除第一页以外，获取其他页面需要传入 page 参数。如 localhost/u?page=2 会显示第 51 ~100 条记录。
- 如果不做特殊说明，所有的发送数据均使用POST方法。
- 如果不做特殊说明，所有的操作行为会返回一个 JSON，boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。

## 详细说明

### 获取用户信息

获取用户列表

    localhost/u

获取某个用户的信息

    localhst/u/:uid 

### 用户行为

- 登陆 OJ

		localhost/user/login

	- 使用 POST 方法
	- 参数列表： 
		- userid 	用户ID
		- password 	密码
	- 返回信息：
		- 信息以 JSON 格式返回。boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。
		- "OK" = true 表示登陆成功，Info 字段以 JSON 对象的格式返回用户信息。
		- "OK" = false表示登陆失败，Info 字段返回错误信息代码。

- 退出登陆

		localhost/user/logout

	- 使用 POST 方法
	- 参数列表：没有参数
	- 返回信息：
		- 信息以 JSON 格式返回。boolean 类型的"OK"字段表示操作是否成功，Info 字段显示操作结果。
		- "OK" = true 表示退出成功，Info 字段返回退出的用户。
		- "OK" = false表示退出失败，Info 字段返回错误信息代码。

- 用户注册

		localhost/user/join

	- 使用 POST 方法
	- 参数列表

		| 参数| 描述 | 备注 |  
		|:---|:------|:-----|
		| userid	|  | 必需。不可更改。|
		| nickname  |  默认为 userid。 ||
		| password  |    | 必需。|
		| password2 |    | 必需。|
		| emailaddr |   | 必需。需要验证邮箱。|
		| studentid |    | 必需。不可更改。|
		| realname  |    | 必需。不可更改。|

