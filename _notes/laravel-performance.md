# Laravel Performance

## Database Read/Write

在 Laravel 裡要做讀寫分離很簡單，首先要調整 config 裡的 db 連線，把 host 改為 read 與 write 分開，範例如下

### Example

app/config/database.php

```php
// ...
'mysql' => [
    'read' => [
        'host' => env('DB_HOST_READ', 'localhost'),
    ],
    'write' => [
        'host' => env('DB_HOST_WRITE', 'localhost'),
    ],
    'driver'    => 'mysql',
    'database'  => env('DB_DATABASE', 'forge'),
    'username'  => env('DB_USERNAME', 'forge'),
    'password'  => env('DB_PASSWORD', ''),
    'charset'   => 'utf8',
    'collation' => 'utf8_unicode_ci',
    'prefix'    => '',
    'strict'    => false,
],
// ...
```

.env

```php
// ...
DB_HOST_READ=192.168.1.1
DB_HOST_WRITE=192.168.1.2
DB_DATABASE=homestead
DB_USERNAME=homestead
DB_PASSWORD=secret
// ...
```

這樣在 select 時會使用 read 的 connection, 而做 delete, update, insert 時則會使用 write 的 connection。

那問題來了，因為 database replicate 其實是需要時間的，那在一些需要即時的資料時〈例如搶票〉，為了確保資料沒有落差，我們就不能使用讀寫分離，此時我們可以使用以下語法讓 select 去使用 write connection。

Eloquent

```php
User::onWriteConnection()->find($id);
```

Query Builder

```php
DB::table('users')->selectFromWriteConnection('*')->where('id', $id)->first();
```

若網站的效能是卡在 database，讀寫分離是個可以嘗試改善效能的方法，但也意味著需要額一台機器做原先 database 的 slave。

## Reference

* [Read / Write Connections@laravel.com](http://laravel.com/docs/5.1/database#read-write-connections)
* [onWriteConnection@Laravel API](http://laravel.com/api/5.1/Illuminate/Database/Eloquent/Model.html#method_onWriteConnection)
* [selectFromWriteConnection@Laravel API](http://laravel.com/api/5.1/Illuminate/Database/Connection.html#method_selectFromWriteConnection)