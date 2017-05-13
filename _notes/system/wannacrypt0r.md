# WannaCrypt0r 2.0

## Translation

### Malware Analysis

1. 最一開始的這隻 mssecsvc.exe 會丟出 tasksche.exe 並執行。

1. 送出 HTTP 請求給特定網域名，確認是否傳播。

1. mssecsvc2.0 服務被創建，這個服務會再次執行 mssecsvc.exe
    1. 這次執行會透過 TCP PORT 445 去嘗試連結子網路(subnet)所有的IP
    1. 連結成功便會開始傳送資料。（這邊可能是感染其他被連結的電腦）

1. tasksche.exe 開始找所有儲存裝置，包含網路資料夾、USB、隨身硬碟

1. 找硬碟裡副檔名符合以下列表的檔案，並以 2048-bit RSA 加密
    * 常見文件檔 (.ppt, .doc, .docx, .xlsx, .sxi)
    * 罕見文件檔 (.sxw, .odt, .hwp)
    * 壓縮檔、影音檔 (.zip, .rar, .tar, .bz2, .mp4, .mkv)
    * 電子郵件相關 (.eml, .msg, .ost, .pst, .edb)
    * 資料庫相關 (.sql, .accdb, .mdb, .dbf, .odb, .myd)
    * 程式碼相關 (.php, .java, .cpp, .pas, .asm)
    * 加解密鑰匙與認證 (.key, .pfx, .pem, .p12, .csr, .gpg, .aes)
    * 設計、圖片、照片 (.vsd, .odg, .raw, .nef, .svg, .psd)
    * 虛擬機器相關 (.vmx, .vmdk, .vdi)

1. 新增一個資料夾"Tor"，裡面有 tor.exe 、 9 個 dll 檔、taskdl.exe、taskse.exe

1. 各程式說明
    * taskse.exe  開啟 @wanadecryptor@.exe 跳出勒索訊息給你看
    * taskdl.exe 刪除暫存檔
    * tasksche.exe 尋找符合格式的檔案並加密

1. 當你想付款的時候就會啟動 Tor.exe (Tor本身無害，他只是被用來創造全匿名的連線，跟他最有關係的是暗網)

1. 用以下三個 windows 指令刪除你的 shadow copy (windows備份和系統還原)
    ![wannacrypt-shadow-copy-delete](https://3.bp.blogspot.com/-Kl4zaGW7jDQ/WRYwl8gD4RI/AAAAAAAAA_s/tZGlUi9jmSMsaX6khcqVi2T1HkcJb-8RwCLcB/s1600/image1.png)

1. 病毒會用以下兩個 windows 指令去用你的隱藏檔和變更檔案存取權限

    ```bash
    # file control
    attrib +h [[Drive:][Path] FileName] [/s[/d]]
    # access control
    icacls . /grant Everyone:F /T /C /Q
    ```

### Kill Switch

1. Cisco Umbrella 的研究人員發現有個奇怪的網址突然大量的被要求解析
    ![wannacrypt-dns-query](https://1.bp.blogspot.com/-jMtZ8ol4fu8/WRYwB-uRldI/AAAAAAAAA_Y/tAFnRICndIUSGHgmv7ffFgl8qoMOFcFOACLcB/s1600/image5.png)

1. 另外也有人分析病毒的行為後，得知病毒有個子程序會送一個HTTP請求給一個網址，若請求失敗，便會傳播，若請求成功，這個子程序就會結束。
    ![wannacrypt-execution](https://3.bp.blogspot.com/-EUgk1JpJjVU/WRYwHbKRO4I/AAAAAAAAA_c/t24Ea80MJOsv5giibrH42V4FjxoAQePywCLcB/s1600/image10.png)

1. 似乎以上兩個線索不謀而合，然後就有勇者花了10.69鎂去註冊了這個網域，讓病毒送請求成功。請求成功，子程序結束，便不會傳播。再次觀察病毒行為，成功的停止了病毒的擴散。
    ![wannacrypt-execution](https://pbs.twimg.com/media/C_pJ_3EXYAA0LAu.jpg)

## Reference

* [Player 3 Has Entered the Game: Say Hello to 'WannaCry'@talosintelligence](http://blog.talosintelligence.com/2017/05/wannacry.html#h.t19xasesgr4p)
* [WannaCry ransomware used in widespread attacks all over the world@Kaspersky Lab](https://securelist.com/blog/incidents/78351/wannacry-ransomware-used-in-widespread-attacks-all-over-the-world/)
* [Cyberattacks in 12 Nations Said to Use Leaked N.S.A. Hacking Tool@Hacker News](https://news.ycombinator.com/item?id=14326439)
* [MalwareTech@Twitter](https://twitter.com/MalwareTechBlog)
* [Infection Live Map](https://intel.malwaretech.com/WannaCrypt.html)
* [Botnet Trackers](https://intel.malwaretech.com/botnet/wcrypt)