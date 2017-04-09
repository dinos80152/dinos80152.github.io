# Queue

```sequence
Consumer->Queue Server: Add Job
Worker->Queue Server: Listen
Worker->Queue Server: Get Job
Worker->Worker: Do Something
```

## Scenario

### Real World

百貨公司地下街點餐，

```sequence
顧客(Client)->櫃臺(Web Application Server): 點餐
櫃臺(Web Application Server)->佇列(Queue Server): 紀錄
廚師(Worker)->佇列(Queue Server):拿單
廚師(Worker)->廚師(Worker):煮菜
廚師(Worker)->顧客(Client):通知
```

直接接觸顧客的櫃臺或服務生就如 Web Server，在廚房專心做菜的廚師就像 Worker，而顧客如何知道何時取餐？靠看板上的號碼或是通知器。

### Computer World

使用者從你的網站上傳圖片後，你會將使用者的圖片縮圖及優化處理後上傳至 File Server，若一般的作法會在網站伺服器做處理，若使用者過多很容易造成伺服器爆炸，連網站都掛點。

使用Job Queue的話，假設你現在有三台機器，一台當 Web Server，兩台當 Worker Server。

```sequence
Client->Web Server: Upload Picture
Web Server->File Server: Upload Picture
Web Server->Queue: Add Picture Process Job
WorkerA->Queue: Listen
WorkerA->Queue: Get Job
WorkerA->File Server: Download Picture
WorkerA->WorkerA: Process Picture
WorkerA->File Server: Upload Thumbnail and Optimized Picture
WorkerB->Queue: Listen
WorkerB->Queue: Get Job
WorkerB->File Server: Download Picture
WorkerB->WorkerB: Process Picture
WorkerB->File Server: Upload Thumbnail and Optimized Picture
```

### Details

### What's Queue

* Data Structure: FIFO (First In First Out) 
* Line Up (排隊)

### Why

* 一次做好一件事

### When

* 事情需要專心處理時，事情需要消耗大量資源時。

### Where
* Beanstalk
* Redis
* Iron.io

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

### Laravel Example

[Laravel 5 Example by Dino Lai @github](https://github.com/dinos80152/laravel5-example/tree/master/app/Jobs)


## Beanstalk
[official web site](http://kr.github.io/beanstalkd/)

### Beanstalk Client Reference

* [PHP Client](https://github.com/pda/pheanstalk)
* [Laravel Documentation](http://laravel.com/docs/5.1/queues)

### Tools
[Beanstalk Console](https://github.com/ptrofimov/beanstalk_console)

### Reference
* [http://www.haodaima.net/art/2428782](http://www.haodaima.net/art/2428782)
* [BeanstalkD: An Introduction to queuing by Alister Bulman](http://alister.github.io/presentations/Beanstalkd/)

## Extended Reading
[RabbitMQ Tutorial (Know How Queue Work)](https://www.rabbitmq.com/getstarted.html)