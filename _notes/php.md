# PHP

## Linux

### Update PHP version

#### Problem

如果在 OS 上已經有 PHP，直接安裝新版 PHP 會出問題，必須使用 yum replace 去做 PHP 的版本更新

#### Enviroment

作業系統：CentOS 7

#### Solution

1. 新增下載位置至 rpm 清單

    ```bash
    rpm -Uvh https://mirror.webtatic.com/yum/el7/epel-release.rpm
    rpm -Uvh https://mirror.webtatic.com/yum/el7/webtatic-release.rpm
    ```

1. 安裝 yum replace plugin

    ```bash
    yum install yum-plugin-replace
    ```

1. 安裝 php 5.6

    ```bash
    yum replace –enablerepo=webtatic-testing php-common –replace-with=php56w-common
    ```

1. 查看現在 php 版本

    ```bash
    php -v
    ```

1. 重新 reload php-fpm

    ```bash
    systemctl reload php-fpm
    ```

### Reference

* <https://webtatic.com/packages/php56/>

## Programming

### Annotation

最近在碰 Java Spring 和 Hibernate，發現有個 @ 怎麼沒看過，原來叫做 Annotation，查了一下發現 PHP 有些 Framework 也開始使用這東西了，不過是寫在 DocBlockr 裡阿，所以 PHP 的各位們似乎不太喜歡這東西。

#### Reference

* [Annotations in PHP: They Exist](http://www.slideshare.net/rdohms/annotations-in-php-they-exist)
* [PHP Annotations Are a Horrible Idea](http://theunraveler.com/blog/2012/php-annotations-are-a-horrible-idea/)
* [PHP: Annotations are an Abomination](https://r.je/php-annotations-are-an-abomination.html)

### PHPUnit

## Tools

### phpcs

#### Command

* Favor Setting

    ```bash
    phpcs --standard=PSR1,PSR2 -n --ignore=Generic.Files.LineEndings,PSR2.Files.EndFileNewline
    ```

* List Standard Sniffs

    ```bash
    phpcs --standard=PSR1,PSR2 -e
    ```
