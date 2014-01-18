#USC Online Judged

uoj的Web端，采用 beego + MongoDB + MySQL 搭建。

网址： www.cs203.net

## 负载

### 数据库需求

- 单条记录 * 数量 * 时间
- 用户：1KB * 200 * 5 = 1MB
- 题目：3KB * 1000 = 3MB
- 作业：10KB * 40 * 5 = 2MB
- 提交：1KB * 30000 * 5 = 150MB
- 消息：2KB * 60000 * 5 = 600MB

PS：平均每条记录有十个字段

###Author

zhrlnt@gmail.com
