# Linux

## Signal

### Ctrl+z

<kbd>Ctrl + z </kbd>是 Unix 裡的一種 Signal，目的是為了暫時執行中的程序，並可對執行程序進行前背景的切換

#### Commands

| 指令     | 說明                                                  |
| -------- | ----------------------------------------------------- |
| bg       | 切換至被背景作業                                      |
| fg       | 切回前景作業 (後進先出，先彈出最後進去背景作業的程序) |
| jobs     | 看目前所有運行於背景的程序                            |
| fg %id   | 指定程序至前景作業                                    |
| kill %id | 刪除指定背景程序                                      |

#### Example

1. 執行 ping

    ```bash
    ping http://www.google.com.tw
    ```

2. 中斷程序

    <kbd>Ctrl + z</kbd>

3. 放於背景作業

    ```bash
    bg
    ```

4. 此時 <kbd>Ctrl + c</kbd> 無法中斷程序，可嘗試 cd 到任何地方

5. 打上 jobs 可見

    ```bash
    [1]+ Stopped ping http://www.google.com.tw
    ```

6. 程序挪至前景作業 or 指定 job id 移至前景 or 直接刪除

    ```bash
    fg or fg %1 or kill %1
    ```

7. 若移至前台，可 <kbd>Ctrl + c</kbd> 中斷程序

#### 補充

也可以直接在要執行的程序後面加上 & 就會在背景執行囉

```bash
ping http://www.google.com.tw &
```

#### Reference

* [Unix signal@Wikipedia](http://en.wikipedia.org/wiki/Unix_signalhttp://en.wikipedia.org/wiki/Unix_signal)
* [What is effect of CTRL + Z on a unix\Linux application@superuser](http://superuser.com/questions/476873/what-is-effect-of-ctrl-z-on-a-unix-linux-application)
* [unix-background-job](http://www.thegeekstuff.com/2010/05/unix-background-job/)
* [第十七章、程序管理與 SELinux 初探@鳥哥](http://linux.vbird.org/linux_basic/0440processcontrol.php)

## Linux Tunnel

### 網路架構

```sequence
local->ssh server(128.18.123.111): connect
ssh server(128.18.123.111)->dist server(10.59.1.23;19336): connect
```

### 使用說明

* -N 連線後不執行指令
* -f 連線後背景執行
* -L 啟用SSH Tunnel

```bash
ssh -N -f -L LOCAL_PORT:DEST_HOST:DEST_HOST_PORT SSH_USER@SSH_SERVER_IP
```

### 範例

```bash
ssh -N -f -L 59196:10.59.1.23:19336 dino@124.18.123.111
```

### 查看連線資訊

```bash
netstat -ntulp
```

### 刪除連線

```bash
kill PID
```

### Reference

* [http://gwokae.mewggle.com/wordpress/index.php/archives/678](http://gwokae.mewggle.com/wordpress/index.php/archives/678)
