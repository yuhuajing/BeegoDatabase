# BeegoDatabase
```javascript
docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret mysql:latest
```

清除Navicat的数据表数据（删除数据后自增条数的id也会在原基础上增加）

语句会清除所有的表数据

```javascript
TRUNCATE TABLE table_name;
```

gorm的官方档案：

```golang
http://v1.gorm.io/docs/query.html
```
