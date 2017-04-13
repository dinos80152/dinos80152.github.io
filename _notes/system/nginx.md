# Nginx

## Introduction

NGINX 是一個 HTTP 和反向代理伺服器(Reverse Proxy Server)，與知名的 Apache Server 一樣可作為 HTTP Server，但比 Apache 更為快速，節省更多系統資源。

> Apache is like Microsoft Word, it has a million options but you only need six. Nginx does those six things, and it does five of them 50 times faster than Apache. — Chris Lea on nginx and WordPress

## Install

### Enviroment

OS: CentOS 7

### Step

1. 加入下載位置至 RPM 清單

    ```bash
    rpm -Uvh http://nginx.org/packages/centos/7/noarch/RPMS/nginx-release-centos-7-0.el7.ngx.noarch.rpm
    ```

1. 透過安裝工具 yum 安裝 nginx

    ```bash
    yum install nginx
    ```

1. 加入防火牆允許清單並重啟防火牆

    ```bash
    firewall-cmd –permanent –zone=public –add-service=http
    firewall-cmd –permanent –zone=public –add-service=httpd
    firewall-cmd –permanent –zone=public –add-service=https
    firewall-cmd –reload
    ```

1. 啟動 NGINX服務

    ```bash
    systemctl start nginx.service
    ```

## Path

### Config

```bash
/etc/nginx/
```

### Log

```bash
/var/log/nginx
```

## Daemon

1. 新增以下 Shell Script 至 /ect/init.d/nginx (記得確認 nginx.conf 與 bin 檔的位置)

    ```bash
    #!/bin/sh
    #
    # nginx - this script starts and stops the nginx daemon
    #
    # chkconfig:   - 85 15
    # description:  Nginx is an HTTP(S) server, HTTP(S) reverse \
    #               proxy and IMAP/POP3 proxy server
    # processname: nginx
    # config:      /etc/nginx/nginx.conf
    # config:      /etc/sysconfig/nginx
    # pidfile:     /var/run/nginx.pid

    # Source function library.
    . /etc/rc.d/init.d/functions

    # Source networking configuration.
    . /etc/sysconfig/network

    # Check that networking is up.
    [ "$NETWORKING" = "no" ] && exit 0

    nginx="/usr/sbin/nginx"
    prog=$(basename $nginx)

    NGINX_CONF_FILE="/etc/nginx/nginx.conf"

    [ -f /etc/sysconfig/nginx ] && . /etc/sysconfig/nginx

    lockfile=/var/lock/subsys/nginx

    make_dirs() {
    # make required directories
    user=`$nginx -V 2>&1 | grep "configure arguments:" | sed 's/[^*]*--user=\([^ ]*\).*/\1/g' -`
    if [ -z "`grep $user /etc/passwd`" ]; then
        useradd -M -s /bin/nologin $user
    fi
    options=`$nginx -V 2>&1 | grep 'configure arguments:'`
    for opt in $options; do
        if [ `echo $opt | grep '.*-temp-path'` ]; then
            value=`echo $opt | cut -d "=" -f 2`
            if [ ! -d "$value" ]; then
                # echo "creating" $value
                mkdir -p $value && chown -R $user $value
            fi
        fi
    done
    }

    start() {
        [ -x $nginx ] || exit 5
        [ -f $NGINX_CONF_FILE ] || exit 6
        make_dirs
        echo -n $"Starting $prog: "
        daemon $nginx -c $NGINX_CONF_FILE
        retval=$?
        echo
        [ $retval -eq 0 ] && touch $lockfile
        return $retval
    }

    stop() {
        echo -n $"Stopping $prog: "
        killproc $prog -QUIT
        retval=$?
        echo
        [ $retval -eq 0 ] && rm -f $lockfile
        return $retval
    }

    restart() {
        configtest || return $?
        stop
        sleep 1
        start
    }

    reload() {
        configtest || return $?
        echo -n $"Reloading $prog: "
        killproc $nginx -HUP
        RETVAL=$?
        echo
    }

    force_reload() {
        restart
    }

    configtest() {
    $nginx -t -c $NGINX_CONF_FILE
    }

    rh_status() {
        status $prog
    }

    rh_status_q() {
        rh_status >/dev/null 2>&1
    }

    case "$1" in
        start)
            rh_status_q && exit 0
            $1
            ;;
        stop)
            rh_status_q || exit 0
            $1
            ;;
        restart|configtest)
            $1
            ;;
        reload)
            rh_status_q || exit 7
            $1
            ;;
        force-reload)
            force_reload
            ;;
        status)
            rh_status
            ;;
        condrestart|try-restart)
            rh_status_q || exit 0
                ;;
        *)
            echo $"Usage: $0 {start|stop|status|restart|condrestart|try-restart|reload|force-reload|configtest}"
            exit 2
    esac
    ```

1. 修改權限

    ```bash
    chmod +x /etc/init.d/nginx
    ```

1. 設定開機啟動

    ```bash
    systemctl enable nginx
    ```

1. 查看是否成功

    ```bash
    # systemctl list-unit-files | grep nginx
    nginx.service enabled
    ```

## Reference

* [http://www.wikivs.com/wiki/apache_vs_nginx)](http://www.wikivs.com/wiki/apache_vs_nginx)
* [http://wiki.nginx.org/RedHatNginxInitScript](http://wiki.nginx.org/RedHatNginxInitScript)
* [http://www.rackspace.com/knowledge_center/article/centos-adding-an-nginx-init-script](http://www.rackspace.com/knowledge_center/article/centos-adding-an-nginx-init-script)