# macOS

## Keyboard Shortcut

- [Rectangle](https://rectangleapp.com/): Move and Resize Window

### iTerm 2

#### Use ⌥ ← and ⌥ → to jump forwards / backwards words

[Use ⌥ ← and ⌥→ to jump forwards / backwards words in iTerm 2, on OS X@coderwall](https://coderwall.com/p/h6yfda/use-and-to-jump-forwards-backwards-words-in-iterm-2-on-os-x)

1. set left ⌥ key to act as an escape character
   1. iterm2 > preferences > profile > keys
   1. left ⌥ key > Esc+
1. set ⌥← to jump backwards
   1. Keyboard Shortcut: ⌥←
   1. Action: Send Escape Sequence
   1. Esc+: b
1. ⌥ → as well

## Key Repeat

Mac OS hold key down is show special char default, to change to repeat input char.

- Enable Key Repeat

  ```bash
  defaults write NSGlobalDomain ApplePressAndHoldEnabled -bool false
  ```

- Disable Key Repeat

  ```bash
  defaults write NSGlobalDomain ApplePressAndHoldEnabled -bool true
  ```

- Restart your computer

- Adjust key repeat speed
  1. System Preferences
  1. Keyboard
  1. Key Repeat
  1. Delay Until Repeat

## Chinese Input

- [McBopomofo](https://mcbopomofo.openvanilla.org/): for Old Style Chinese Input like 「ㄅ半」

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

- View

  ```sh
  $sudo sysctl -a
  ```

- Adjust

  ```sh
  $sudo sysctl -w kern.maxfiles=1048600
  $sudo sysctl -w kern.maxfilesperproc=1048576
  ```

### Open Files Limit

- Check ulimit

  ```sh
  ulimit -n
  ```

#### Session

- Adjust ulimit

  ```sh
  ulimit -S -n 1048576
  ```

#### System Wide

- Temporarily

  ```sh
  $sudo sysctl -w fs.file-max=1048576
  ```

- Persist
  - Add `fs.file-max=1048576` in `/etc/sysctl.conf`

#### User Level

- Edit `/etc/security/limits.conf`

  ```bash
  <domain>    <type>  <item>  <value>
  dinolai     hard    nofile  4096
  dinolai     soft    nofile  1024
  ```
