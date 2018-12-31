# Bluetooth Not Working on Ubuntu

## Problem

Turn Bluetooth on by Indicator Applet, but there are no adapters showing on Bluetooth Adapters, and no device showing on Bluetooth Manager.

## Problem Identification

* Check bluetooth process, *OK*.

    ```bash
    $ systemctl status bluetooth
    ● bluetooth.service - Bluetooth service
    Loaded: loaded (/lib/systemd/system/bluetooth.service; enabled; vendor preset: enabled)
    Active: active (running) since Sun 2018-12-30 01:33:45 CST; 1 day 17h ago
        Docs: man:bluetoothd(8)
    Main PID: 16797 (bluetoothd)
    Status: "Running"
        Tasks: 1 (limit: 4915)
    CGroup: /system.slice/bluetooth.service
            └─16797 /usr/lib/bluetooth/bluetoothd
    ```

* Check bluetooth devices by `hciconfig`, **no device found**.

    ```bash
    $ hciconfig -a
    hci0:	Type: Primary  Bus: USB
        BD Address: 00:00:00:00:00:00  ACL MTU: 0:0  SCO MTU: 0:0
        DOWN
        RX bytes:0 acl:0 sco:0 events:0 errors:0
        TX bytes:0 acl:0 sco:0 commands:0 errors:0
        Features: 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
        Packet type: DM1 DH1 HV1
        Link policy:
        Link mode: SLAVE ACCEPT
    ```

* Check bluetooth control interface, **no default controller available**.

    ```bash
    $ bluetoothctl
    Agent registered
    [bluetooth]# scan on
    No default controller available
    ```

## Solution

* Hot remove and reload the kernel module

    ```console
    $ sudo rmmod btusb
    $ sudo modprobe btusb
    ```

* Check bluetooth devices by `hciconfig` again.

    ```bash
    $ hciconfig -a
    hci0:	Type: Primary  Bus: USB
        BD Address: 58:00:E3:C5:C1:E8  ACL MTU: 1024:8  SCO MTU: 50:8
        UP RUNNING
        RX bytes:759 acl:0 sco:0 events:58 errors:0
        TX bytes:4853 acl:0 sco:0 commands:58 errors:0
        Features: 0xff 0xfe 0x8f 0xfe 0xd8 0x3f 0x5b 0x87
        Packet type: DM1 DM3 DM5 DH1 DH3 DH5 HV1 HV2 HV3
        Link policy: RSWITCH HOLD SNIFF
        Link mode: SLAVE ACCEPT
        Name: 'ideapad-510S'
        Class: 0x1c010c
        Service Classes: Rendering, Capturing, Object Transfer
        Device Class: Computer, Laptop
        HCI Version: 4.1 (0x7)  Revision: 0x0
        LMP Version: 4.1 (0x7)  Subversion: 0x25a
        Manufacturer: Qualcomm (29)
    ```

* Use `bluetoothctl` to manipulate bluetooth.

    ```bash
    $ bluetoothctl
    [NEW] Controller 58:00:E3:C5:C1:E8 ideapad-510S [default]
    [NEW] Device 00:11:67:55:66:AA Soundbody
    Agent registered
    [bluetooth]# list
    Controller 58:00:E3:C5:C1:E8 ideapad-510S [default]
    [bluetooth]# devices
    Device 00:11:67:55:66:AA Soundbody
    ```

## Reference

* [Bluetooth Issue: "No default controller available"@archlinux.org](https://bbs.archlinux.org/viewtopic.php?pid=1731068#p1731068)