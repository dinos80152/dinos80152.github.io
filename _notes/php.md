# PHP

## Linux

### Update PHP version

#### Problem

如果在 OS 上已經有 PHP，直接安裝新版 PHP 會出問題，必須使用 yum replace 去做 PHP 的版本更新

#### Enviroment

作業系統：CentOS 7

#### Solution

1.新增下載位置至 rpm 清單
```bash
rpm -Uvh https://mirror.webtatic.com/yum/el7/epel-release.rpm
rpm -Uvh https://mirror.webtatic.com/yum/el7/webtatic-release.rpm
```

2.安裝 yum replace plugin
```bash
yum install yum-plugin-replace
```

3.安裝 php 5.6
```bash
yum replace –enablerepo=webtatic-testing php-common –replace-with=php56w-common
```

4.查看現在 php 版本
```bash
php -v
```

5.重新 reload php-fpm
```bash
systemctl reload php-fpm
```

## Reference
* https://webtatic.com/packages/php56/