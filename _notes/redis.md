# Redis

## Enviroment

* OS: Ubuntu 14.0.4

## Install

### Download and Install

1. 下載並編譯

    ```bash
    $ wget http://download.redis.io/redis-stable.tar.gz
    $ tar xvzf redis-stable.tar.gz
    # 習慣將第三方軟體放置 /opt
    $ sudo mv redis-stable /opt/
    $ cd /opt/redis-stable
    $ make
    ```

    執行以上後會出現 /opt/redis-stable/src 資料夾

1. 測試安裝環境

    ```bash
    make test
    ```

1. 自動佈署 /usr/local/bin/redis-server, /usr/local/bin/redis-cli

    ```bash
    make install
    ```

1. 測試 server

    ```bash
    sudo redis-server
    ```

1. 測試 client

    ```bash
    $ redis-cli ping
    pong
    ```

1. 都有動那就安裝成功了！

### Configure on Linux

1. 建立 redis dump 資料時，所儲存的資料夾

    ```bash
    sudo mkdir /var/redis/6379
    ```

1. 建立 redis 設定檔

    ```bash
    mkdir /etc/redis
    sudo cp redis.conf /etc/redis/6379.conf
    ```

1. 設定 conf

    ```bash
    > $ sudo vim /etc/redis/6379.conf
    #/etc/redis/6379.conf
    daemonize yes
    pidfile /var/run/redis_6379.pid
    loglevel notice
    logfile "/var/log/redis_6379.log"
    dir /var/redis/6379
    ```

1. 建立 daemon 檔

    ```bash
    sudo cp utils/redis_init_script /etc/init.d/redis_6379
    # 可以 vim 確認設置
    ```

1. 設定 redis 開機啟動

    ```bash
    sudo update-rc.d redis_6379 defaults
    ```

1. 啟動 redis!

    ```bash
    /etc/init.d/redis_6379 start
    ```

1. 恭喜你成功安裝 Redis!

### Test

1. 塞筆資料進去

    ```bash
    $ redis-cli
    redis> set key value
    OK
    ```

1. 取資料出來

    ```bash
    redis> get key
    "value"
    ```

### Reference

* [Redis Quick Start@redis](http://redis.io/topics/quickstart)