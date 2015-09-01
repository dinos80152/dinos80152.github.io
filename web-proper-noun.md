# 網站開發專有名詞 For Product Manager


## 前端

### HTML (Hyper Text Markup Language)

```
<html>
  <body>
    <b>Example</b>
  </body>
</html>
```

#### Meta Tag


### CSS (Cascading Style Sheets)
  * 樣式表
  * 位置，顯示，大小，顏色...
  * (動畫)

### Javascript
  * DOM (Document Object Model) 元素控制
  * 動畫
  * 註冊事件
  * 邏輯
  * 傳送資料

http://www.w3.org/TR/DOM-Level-2-Core/introduction.html）

### RWD (Responsive Web Design)
  * 一份 HTML, CSS, Javascript 打遍所有裝置
  * 透過 CSS Media Query 決定在哪個解析度時使用哪一些 CSS
  * ex: http://headshot.garena.tw

```
@media screen and (max-width: 300px) {
    body {
        background-color: lightblue;
    }
}
```

### AWD (Adaptive Web Design)
  * 透過 Client 或 Server 偵測使用者瀏覽器屬性（哪種瀏覽器，裝備，螢幕大小），決定要回傳哪種 HTML, CSS, Javascript 給使用者。
  * 會做手機版，桌機版不同的版面。
  * 更彈性
  * ex: http://thunderstrike.garena.tw

http://blog.catchpoint.com/2014/07/16/adaptive-vs-responsive-web-design-quantifying-difference/

### CDN (Content Delivery Network)
* 多個節點會互相同步，要東西時可以去跟最近的拿，快又有效率。



## Browser


### Browser Cache
  * Ctrl + F5

### IE
  * 模擬 IE 各版本
  * 請調成 IE 11

### Chrome
  * AdBlock
  * ad_ 都會被擋
  * pop_ 也會被擋

### RWD
  * F12



## 後端
* PHP
  * 伺服器端程式語言
  * 跟資料庫做溝通
  * 對資料做處理
  * 產生回傳頁面

* Cache (快取)
  * 減少 Application, Database Server Loading
  * 將第一次計算完的資料暫時存在記憶體或是檔案，這段期間都讀這個檔案所存的結果，不去做程式計算和資料庫存取。
  * 大約，不即時（每X分鐘更新一次）

* Queue
  * 排隊
  * 非同步（不能即時回傳結果）

* Database (資料庫)
  * 存資料的地方
  * 撈資料
  * xxxQL
  * xxxDB
  * 記結果

* Log (記錄檔)
  * 記過程



## 伺服器

### DNS (Domain Name System)
* 找東西

### Load Balance (負載平衡)
* 使用演算法判斷此請求要交給哪台伺服器做處理
* ex: 分配者，主管決定事情要給誰做

https://en.wikipedia.org/wiki/Load_balancing_(computing)

### 常見 Http Error Code
* 400 Bad Request
  * 參數錯（網址錯）
  * ex: xxx.tw?id=abc

* 403 Forbidden
  * Web Server (Nginx) IP 沒開

* 404 Not Found
  * Application Server 沒同步
  * 網址錯

* 500 Internal Server Error
  * 程式寫錯

* 502 Bad Gateway
  * Application Server 爆了
  * Database Server 爆了

* 503 Service Unavailable
  * Load Balance Server 爆了

http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html



## Example


### Code

* index.html

```html
<html>
	<head>
		<title>抽獎</title>
		<link rel="stylesheet" type="text/css" href="index.css">
	</head>
	<body>
		<table>
			<thead>
	     		<tr>
	     			<th>姓名</th>
	     			<th>電話</th>
	     			<th>出生年月日</th>
	     		</tr>
	     	</thead>
	     	<tbody>
	     		<tr>
	     			<td><?php echo $user->name; ?>></td>
	     			<td><?php echo $user->phone; ?></td>
	     			<td><?php echo $user->birthdate; ?></td>
	     		</tr>
     		</tbody>
     	</table>
     	<script src="index.js"></script>
  	</body>
</html>
```

* index.css

```css
table tbody tr th {
	color: red;
}

table tbody tr td {
	color: green;
}
```

* index.js

```javascript
$(function() {
	$('td').click(function() {
		alert('恭喜得獎');
	});
});
```

* index.php

```php
<?php
    $employee_id = rand(1, 320);
    $user = findOnGoogleSpreadsheetById($employee_id);
?>
```

### Rule
* 1 Browser, 1 LB, 2 Ap server, 1 DB, 1 Queue, 1 Cache. Total 7.