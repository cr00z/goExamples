Before docker run

```
rm -rf ~/Library/Containers/com.docker.docker
mkdir -p ~/goinfre/Docker/Data
ln -s ~/goinfre/Docker ~/Library/Containers/com.docker.docker
...
brew install mysql-client
```

Run mysql

```
$ docker run --name=mysql -d -p 3306:3306 -v ~/goSite/mysqlSite/mysql_data:/var/lib/mysql mysql/mysql-server
$ docker logs mysql
[Entrypoint] GENERATED ROOT PASSWORD: 385U/:2Ng@&CM8K=c9*O.vgeGIL.ww88
$ docker exec -it mysql mysql -uroot -p

```

Add db & table

```
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '123';
mysql> CREATE DATABASE mydb;
mysql> USE mydb;
mysql> CREATE TABLE users
(
    id int auto_increment primary key,
    name varchar(32) null,
    surname varchar(32) null
) CHARACTER SET utf8 COLLATE utf8_general_ci;
INSERT INTO `mydb`.users (id, name, surname) VALUES (1, 'John', 'Doe');
INSERT INTO `mydb`.users (id, name, surname) VALUES (2, 'Jane', 'Doe');
INSERT INTO `mydb`.users (id, name, surname) VALUES (3, 'Джек', 'Доусон');
INSERT INTO `mydb`.users (id, name, surname) VALUES (4, 'Лизель', 'Мемингер');
INSERT INTO `mydb`.users (id, name, surname) VALUES (5, 'Джейн', 'Эйр');
INSERT INTO `mydb`.users (id, name, surname) VALUES (6, 'Мартин', 'Иден');
INSERT INTO `mydb`.users (id, name, surname) VALUES (7, 'Джон', 'Голт');
INSERT INTO `mydb`.users (id, name, surname) VALUES (8, 'Сэмвелл', 'Тарли');
INSERT INTO `mydb`.users (id, name, surname) VALUES (9, 'Гермиона', 'Грейнджер');
mysql> CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
mysql> GRANT ALL PRIVILEGES ON mydb.* TO 'user'@'localhost';
mysql> RENAME USER 'user'@'localhost' TO 'user'@'172.17.0.1';
```