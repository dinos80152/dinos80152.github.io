# MySQL

## Storage Engine

> Storage engines are MySQL components that handle the SQL operations for different table types--MySQL 5.6 Reference Manual

## Comparison

InnoDB vs. MyISAM

| MySQL                 | MyISAM                         | InnoDB                               |
| --------------------- | ------------------------------ | ------------------------------------ |
| Full Text Search      | Yes                            | > 5.6.4                              |
| Design                | DBMS                           | RDBMS                                |
| Referential integrity |                                | Yes                                  |
| Transactions          |                                | ACID Transaction                     |
| Foreign Key           |                                | Yes                                  |
| Locking               | Table Level                    | Row Level                            |
| select                | better                         |                                      |
| insert, update        |                                | better                               |
| Storage               | No ordering in storage of data | Row data stored in pages in PK order |

## Reference

* [MyISAM versus InnoDB@stackoverflow](http://stackoverflow.com/questions/20148/myisam-versus-innodb)
* [Rackspace](http://www.rackspace.com/knowledge_center/article/mysql-engines-myisam-vs-innodb)
* [What's the difference between MyISAM and InnoDB?@stackoverflow](http://stackoverflow.com/questions/12614541/whats-the-difference-between-myisam-and-innodb)
* [MySQL 5.6 Reference Manual](http://dev.mysql.com/doc/refman/5.6/en/storage-engines.html)