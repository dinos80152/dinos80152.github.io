# CentOS 7

## Firewall

從 CentOS7 開始防火牆使用 firewalld 取代 iptables 了，以下介紹一下 firewall 的指令以及設定路徑

### 指令

* firewall-cmd [--zone][--]

zone |
-----|
drop |
block|
public|
external|
dmz|
work|
home|
internal|
trusted|

### 設定路徑

* /ect/firewalld/
* /usr/lib/firewalld/

### 參考

* [Redhat](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Security_Guide/sec-Using_Firewalls.html#sec-Introduction_to_firewalld)
