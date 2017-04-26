# Windows 共用資料夾(Shared folder)至 VirtualBox Linux - 設定與自動掛載

要在 VirtualBox 裡的 Linux 掛載 Windows 的資料夾，必須藉由 VirtualBox 所提供的 Guest Additions 完成。設定開機自動掛載時須先設定 modules-load vboxsf 才不會導致開機失敗。

## Environment

* 本機 OS: Windows 7
* 虛擬機 OS: CentOS 7

## Step

### 掛載 VboxGuestAdditions

1. 在虛擬機器視窗 > 裝置 > CD/DVD 裝置 > 選擇虛擬 CD/DVD 磁碟檔案
1. 選定 C:\Program Files\Oracle\VirtualBox\VboxGuestAdditions.iso

    ![掛載 VboxGuestAdditions.iso](https://farm9.staticflickr.com/8587/16632772349_ffab77e2ae_b_d.jpg)

### 安裝 Guest Additions

1. 進入 CentOS 7，並切換至 root
1. 掛載 cdrom

    ```bash
    mount /dev/cdrom /media/cdrom/
    ```

1. 選擇 Linux 版本的 Guest Additions 安裝

    ```bash
    /media/cdrom/VBoxLinuxAdditions.run
    ```

### 設定共用資料夾

1. Virtual Box 管理視窗 > 選擇虛擬機器 > 設定值 > 共用資料夾 > 加入新的共用資料夾
1. 選擇資料夾路徑，勾選「自動掛載」、「設為永久」

    ![共用資料夾設定](https://farm9.staticflickr.com/8722/16819021845_71920f1842_b_d.jpg)

1. 進入 CentOS，看是否有讀取到共用資料夾，若成功會看到多一資料夾為 sf_

    ```bash
    ls /media/
    cdrom sf_Project
    ```

1. 查看資料夾內容是否有存取進來

    ```bash
    ls /media/sf_Project
    ```

### 掛載

1. 掛載共用資料夾至指定資料夾底下以供 CentOS 使用，**直接使用所設定的資料夾名稱**

    ```bash
    mount -t vboxsf **project** /mnt/win/project/
    ```

### 開機自動掛載

如果直接編輯 /etc/fstab 會導致開機失敗，必須設定先讀取 vboxsf 模組

1. 新增開機讀取模組設定檔: vboxsf.conf

    ```bash
    touch /etc/modules-load.d/vboxsf.conf
    ```

1. 新增一行 vboxsf 至 vboxsf.conf
1. 新增開機掛載至 /etc/fstab

    ```bash
    vim /etc/fstab
    # ([來源][目的地][檔案類型][系統參數][dump][fsck])
    project /mnt/win/project vboxsf defaults 0 0
    ```

1. 重新開機！ 會發現共用資料夾已自動掛載囉！

## Reference

* [larsar@github](https://gist.github.com/larsar/1687725)
* [Richard Turner@askubuntu](http://askubuntu.com/questions/365346/virtualbox-shared-folder-mount-from-fstab-fails-works-once-bootup-is-complete)
* [鳥哥的 Linux 私房菜](http://linux.vbird.org/linux_basic/0230filesystem.php#fstab)