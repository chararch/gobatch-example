# GoBatch 快速入门

本项目演示了如何使用GoBatch进行简单的批处理任务。它从CSV文件读取用户数据，处理后将结果输出到MySQL数据库。

## 其他语言版本

- [English Version](README.md)

## 前提条件

- 系统上已安装Go
- 已安装MySQL数据库

## 设置

1. **创建数据库**：在您的MySQL服务器中设置一个新数据库，例如`gobatch`。
   - **注意**：创建数据库后，请修改`main.go`中的数据库连接详细信息以匹配您的MySQL配置。更新用户名、密码、主机和数据库名称（如果不同于`gobatch`）。

2. **初始化GoBatch运行时表**：使用[schema_mysql.sql](https://github.com/chararch/gobatch/blob/master/sql/schema_mysql.sql)中的SQL脚本创建GoBatch所需的运行时依赖表。

3. **创建用户表**：使用提供的`users.sql`在您的MySQL数据库中创建必要的`users`表。

4. **CSV数据**：确保`users.csv`位于项目目录中。

## 构建和运行

使用提供的`run.sh`脚本构建并运行项目。

```bash
./run.sh
```