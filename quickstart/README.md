# GoBatch Quickstart

This project demonstrates a simple batch processing task using GoBatch. It reads user data from a CSV file, processes it, and outputs the results to a MySQL database.

## Other Language Versions

- [中文版](README_zh.md)

## Prerequisites

- Go installed on your system
- MySQL database setup

## Setup

1. **Create Database**: Set up a new database in your MySQL server, for example, `gobatch`.
   - **Note**: After creating the database, modify the database connection details in `main.go` to match your MySQL configuration. Update the username, password, host, and database name (if different from `gobatch`).

2. **Initialize GoBatch Runtime Tables**: Use the SQL script from [schema_mysql.sql](https://github.com/chararch/gobatch/blob/master/sql/schema_mysql.sql) to create the runtime dependency tables required by GoBatch.

3. **Create Users Table**: Use the provided `users.sql` to create the necessary `users` table in your MySQL database.

4. **CSV Data**: Ensure `users.csv` is in the project directory.

## Build and Run

Use the provided `run.sh` script to build and run the project.

```bash
./run.sh
```
