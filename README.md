# Simple Todolist using Golang IO and Mysql

## HOW TO USE
1. create 2 database with name :
   * `todolist` for product
   * `todolist_test` for test
2. create table in both databsae :
```mysql
CREATE TABLE todolist (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL ,
    author VARCHAR(255) NOT NULL 
)
   ```
3. create environment before running
    * for test : `env=test`
    * for product : `env=product`
4. before running program, make sure your database connection is right.
   you can check on file `utils/database.go`
5. how to running :
   * To running all test, open terminal and type :
   > go test ./...
   * To running programm, open terminal and type :
   > go run main.go