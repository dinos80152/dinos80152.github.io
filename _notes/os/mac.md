# Mac

## Keyboard Shortcut

* [Spectacle](https://www.spectacleapp.com/): Move and Resize Window

## Key Repeat

Mac OS hold key down is show special char default, to change to repeat input char.

* Enable Key Repeat

    ```bash
    defaults write NSGlobalDomain ApplePressAndHoldEnabled -bool false
    ```

* Disable Key Repeat

    ```bash
    defaults write NSGlobalDomain ApplePressAndHoldEnabled -bool true
    ```

* Restart your computer

* Adjust key repeat speed
  1. System Preferences
  1. Keyboard
  1. Key Repeat
  1. Delay Until Repeat

## Chinese Input

* [McBopomofo](https://mcbopomofo.openvanilla.org/): for Old Style Chinese Input like 「ㄅ半」

## Path

1. change dir to paths.d

    ```bash
    /etc/paths.d/
    ```

1. add path file, name "go"

    ```bash
    #go
    /usr/local/go/bin
    ```

1. save file, and restart terminal

## System Limit

### Kernel

* View

    ```sh
    $sudo sysctl -a
    ```

* Adjust

    ```sh
    $sudo sysctl -w kern.maxfiles=1048600
    $sudo sysctl -w kern.maxfilesperproc=1048576
    ```

### Open Files Limit

* Check ulimit

    ```sh
    ulimit -n
    ```

#### Session

* Adjust ulimit

    ```sh
    ulimit -S -n 1048576
    ```

#### System Wide

* Temporarily

    ```sh
    $sudo sysctl -w fs.file-max=1048576
    ```

* Persist
  * Add `fs.file-max=1048576` in `/etc/sysctl.conf`

#### User Level

* Edit `/etc/security/limits.conf`

    ```bash
    <domain>    <type>  <item>  <value>
    dinolai     hard    nofile  4096
    dinolai     soft    nofile  1024
    ```
