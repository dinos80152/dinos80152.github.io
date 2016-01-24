# Linux Tunnel

## 說明

### 網路架構
local -> ssh server(128.18.123.111) -> dest server(10.59.1.23:19336)

### 使用說明
* -N 連線後不執行指令
* -f 連線後背景執行
* -L 啟用SSH Tunnel

```
ssh -N -f -L LOCAL_PORT:DEST_HOST:DEST_HOST_PORT SSH_USER@SSH_SERVER_IP
```

### 範例

```
ssh -N -f -L 59196:10.59.1.23:19336 dino@124.18.123.111
```

### 查看連線資訊

```
netstat -ntulp
```

### 刪除連線

```
kill PID
```

## Reference
http://gwokae.mewggle.com/wordpress/index.php/archives/678
