# Selenium with PHPUnit

## Requirements

* Java JDK
* PHP
* PHPUnit

## WebDriver

1. Install

  Download And Install Selenium Standalone Server

  <http://docs.seleniumhq.org/download/>

1. Start

```bash
java -jar selenium-server-standalone-2.52.0.jar
```

## Facebook PHP Web Driver

```bash
php composer.phar require facebook/webdriver
```

## Integrate with PHPUnit

```php
<?php

use Facebook\WebDriver\Remote\DesiredCapabilities;
use Facebook\WebDriver\Remote\RemoteWebDriver;
use Facebook\WebDriver\WebDriverBy;

class RegitsterTest extends PHPUnit_Framework_TestCase
{

    protected $webDriver;

    public function setUp()
    {
        $this->webDriver = RemoteWebDriver::create('http://localhost:4444/wd/hub', DesiredCapabilities::firefox());
    }

    public function tearDown()
    {
        $this->webDriver->quit();
    }
}
```

## Example

<https://github.com/dinos80152/php-tutorial/selenium>

## Tutorial

* [Facebook PHP WebDriver Example](https://github.com/facebook/php-webdriver/blob/community/example.php)(Only WebDriver)
* [Using the Selenium Web Driver API with PHPUnit@sitepoint](http://www.sitepoint.com/using-the-selenium-web-driver-api-with-phpunit/)
* [Working with PHPUnit and Selenium Webdriver@codeception](http://codeception.com/11-12-2013/working-with-phpunit-and-selenium-webdriver.html#.VsIAAFV95QI)

## Documentation

<http://facebook.github.io/php-webdriver/>

## Reference

* [PHP WebDriver](https://github.com/facebook/php-webdriver)
* [Using the Selenium Web Driver API with PHPUnit@sitepoint](http://www.sitepoint.com/using-the-selenium-web-driver-api-with-phpunit/)