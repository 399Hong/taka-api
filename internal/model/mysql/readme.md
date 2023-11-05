To generate model code for table
1. Create a folder with the name of the table.
2. Create the DDL file ( {tableName}.sql ) for the table and place it in the folder created in step 1.
3. Run **goctl model mysql ddl** with the arguments below
    - --src {folderLocation}/{table}.sql
    - --dir {folderLocation}

goctl doesn't allow foreign key definition in DDL to generate code, use --datasource {connectionString}**

For Example:
```bash
 goctl model mysql ddl --src internal/model/mysql/user/user.sql --dir internal/model/mysql/user    
```
Alternatively, to generate from database
```bash
goctl model mysql datasource -url="root:12345678@tcp(127.0.0.1:3306)/taka" -table="company"  -dir=internal/model/mysql/company
```