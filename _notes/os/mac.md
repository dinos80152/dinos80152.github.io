# Mac

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

* Adjust key repeat speed
  1. System Preferences
  1. Keyboard
  1. Key Repeat

## Path

1. change dir to paths.d

    ```bash
    /ect/paths.d/
    ```

1. add path file, name "go"

    ```bash
    #go
    /usr/local/go/bin
    ```

1. save file, and restart terminal