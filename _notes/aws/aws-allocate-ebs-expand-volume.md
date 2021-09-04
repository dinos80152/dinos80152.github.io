# AWS Allocate EBS Expand Volume

1. Expand EBS Volume
    ![alt tag](https://github.com/dinos80152/dinos80152.github.io/blob/master/_data/aws-ebs-modify-volume.png?raw=true)
1. Connect to instance
1. Find disk space and mount point

    ```sh
    [root@ip-10-103-224-94 /]# df -h
    Filesystem      Size  Used Avail Use% Mounted on
    devtmpfs        187G  104K  187G   1% /dev
    tmpfs           187G  4.0K  187G   1% /dev/shm
    /dev/nvme0n1p1  7.9G  2.0G  5.8G  26% /
    /dev/dm-3       9.8G   93M  9.2G   1% /var/lib/docker/devicemapper/mnt/d3a092243df961a990444ce90710b55cac3418147156d0ac35a8125ef8bf594e
    shm              64M     0   64M   0% /var/lib/docker/containers/7ca9f9cfead972d931d69ab11fb84c3ee5c5af0a96e8c544b39c8581db30ae14/mounts/shm
    /dev/dm-4       9.8G   37M  9.3G   1% /var/lib/docker/devicemapper/mnt/140aa6872d9638ccbd69825a6bbbf68c4f98376f5a71c9006832d82a98b211b7
    shm              64M     0   64M   0% /var/lib/docker/containers/a3615fa2900c14dff8f5cfc75364c452fb15fc9af4b614f291d60892d7e3cbc1/mounts/shm
    /dev/dm-5       9.8G  165M  9.2G   2% /var/lib/docker/devicemapper/mnt/680b7e0fad7001f62626680d0138587c6206aa904718867a2a955f37dd024466
    shm              64M     0   64M   0% /var/lib/docker/containers/f9075f1f8dd3c2046c251f1232dac06d5b3587e2d26f0a0bb3358da1fb39bb39/mounts/shm
    /dev/xvdo       4.9T  1.7T  3.0T  37% /var/lib/docker/plugins/b88a69330fd8d1f46cbc9148f6ca3a7a8a3aa06cc736d253d5bd56a73fe6605a/propagated-mount/volumes/prometheus-vol
    ```

1. Resize mount point by resize2fs

    ```sh
    [root@ip-10-103-224-94 /]# sudo resize2fs -f /dev/xvdo
    resize2fs 1.43.5 (04-Aug-2017)
    Filesystem at /dev/xvdo is mounted on /var/lib/docker/plugins/b88a69330fd8d1f46cbc9148f6ca3a7a8a3aa06cc736d253d5bd56a73fe6605a/propagated-mount/volumes/prometheus-vol; on-line resizing required
    old_desc_blocks = 625, new_desc_blocks = 1250
    The filesystem on /dev/xvdo is now 2621440000 (4k) blocks long.
    ```

1. Check size

    ```sh
    [root@ip-10-103-224-94 /]# df -h
    Filesystem      Size  Used Avail Use% Mounted on
    devtmpfs        187G  104K  187G   1% /dev
    tmpfs           187G  4.0K  187G   1% /dev/shm
    /dev/nvme0n1p1  7.9G  2.0G  5.8G  26% /
    /dev/dm-3       9.8G   93M  9.2G   1% /var/lib/docker/devicemapper/mnt/d3a092243df961a990444ce90710b55cac3418147156d0ac35a8125ef8bf594e
    shm              64M     0   64M   0% /var/lib/docker/containers/7ca9f9cfead972d931d69ab11fb84c3ee5c5af0a96e8c544b39c8581db30ae14/mounts/shm
    /dev/dm-4       9.8G   37M  9.3G   1% /var/lib/docker/devicemapper/mnt/140aa6872d9638ccbd69825a6bbbf68c4f98376f5a71c9006832d82a98b211b7
    shm              64M     0   64M   0% /var/lib/docker/containers/a3615fa2900c14dff8f5cfc75364c452fb15fc9af4b614f291d60892d7e3cbc1/mounts/shm
    /dev/dm-5       9.8G  165M  9.2G   2% /var/lib/docker/devicemapper/mnt/680b7e0fad7001f62626680d0138587c6206aa904718867a2a955f37dd024466
    shm              64M     0   64M   0% /var/lib/docker/containers/f9075f1f8dd3c2046c251f1232dac06d5b3587e2d26f0a0bb3358da1fb39bb39/mounts/shm
    /dev/xvdo       9.7T  1.7T  7.6T  19% /var/lib/docker/plugins/b88a69330fd8d1f46cbc9148f6ca3a7a8a3aa06cc736d253d5bd56a73fe6605a/propagated-mount/volumes/prometheus-vol
    ```

## Reference

- [How to Resize a Volume in AWS@Winston Kotzan's Technology Blog](https://winstonkotzan.com/blog/2017/03/03/how-to-resize-a-volume-in-aws.html)