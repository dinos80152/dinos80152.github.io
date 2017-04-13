# Nginx, Php-fpm

## nginx

### configuration

### log

* /var/log/nginx/

## php-fpm

* php-fpm.ini

### CGI

Common Gateway Interface
> Common Gateway Interface (CGI) is a standard environment for web servers to interface with executable programs installed on a server that generate web pages dynamically. - wikipedia

Cons
> CGI applications run in separate processes, which are created at the start of each request and torn down at the end.

* limits efficiency and scalability
> At high loads, the operating system process creation and destruction overhead becomes significant.

* limits resource reuse techniques (such as reusing database connections, in-memory caching, etc.).

* Load environment variable every time (php.ini, extensions)

### FastCGI

FastCGI uses persistent processes to handle a series of requests. These processes are owned by the FastCGI server, not the web server.

* Socket

```bash
// nginx example
```

* TCP Connection

```bash
// nginx example
```

#### Pros

* Each individual FastCGI process can handle many requests over its lifetime
* Multiple FastCGI servers can be configured, increasing stability and scalability.
* server and application processes to be restarted independently
* Different types of incoming requests can be distributed to specific FastCGI servers which have been equipped to handle those particular types of requests efficiently.
* php解释程序被载入内存而不是每次需要时从存储器读取
* 服务器不用在每次需要时都载入php解释程序

#### log

/var/log/php-fpm/

#### status

<http://example.com/status>