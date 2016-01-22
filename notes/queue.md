# Queue

## What's Queue

* FIFO (First In First Out) Data Structure
* Line Up (排隊)

## Why

* 一次做好一件事

## When

* 事情需要專心處理時，事情需要消耗大量資源時。

## Where
* Beanstalkd
* Redis
* Iron.io

## How

* 顧客點餐(Client)→服務生紀錄(Web Application Server)→佇列(Queue Server)→廚師煮菜(Job Server)→通知顧客

### Simple Queue Example In PHP

```php
class Queue
{
    $queue = [];

    public function enqueue()
    {
        return $this->queue
    }

    public function dequeue()
    {

    }
}
```

### Laravel