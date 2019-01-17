# Make Emoji Colorful on Chrome in Ubuntu Linux

I have try `NotoColorEmoji.tff` but emojis are **not complete**, and [twemoji-color-font](https://github.com/eosrei/twemoji-color-font) is **only support firefox**.

Finally, I found [EmojiOne Fonts](https://github.com/emojione/emojione/tree/master/extras/fonts) could be completed my requirement.

## Installation Steps

1. Download newest `emojione-android.ttf` from [emojione-assets](https://github.com/emojione/emojione-assets/releases)
1. Move `emojione-android.ttf` to `~/.local/share/fonts/`
1. Create `70-emojione-color.conf` in `/etc/fonts/conf.avail/`

    ```xml
    <?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE fontconfig SYSTEM "fonts.dtd">
    <fontconfig>

    <!-- Add emoji generic family -->
    <match target="pattern">
        <test qual="any" name="family"><string>emoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <!-- Aliases for the other emoji fonts -->
    <match target="pattern">
        <test qual="any" name="family"><string>Apple Color Emoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Segoe UI Emoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Segoe UI Symbol</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>EmojiOne</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Emoji One</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Android Emoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>NotoColorEmoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Noto Emoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Twemoji</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>EmojiSymbols</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <match target="pattern">
        <test qual="any" name="family"><string>Symbola</string></test>
        <edit name="family" mode="assign" binding="same"><string>Noto Color Emoji</string></edit>
    </match>

    <!-- Do not allow any app to fallback to Symbola, ever -->
    <selectfont>
        <rejectfont>
        <pattern>
            <patelt name="family">
            <string>Symbola</string>
            </patelt>
        </pattern>
        </rejectfont>
    </selectfont>
    </fontconfig>
    ```

1. Create symbolic link to enable font config

    ```shell
    sudo ln -s /etc/fonts/conf.avail/70-emojione-color.conf /etc/fonts/conf.d/70-emojione-color.conf
    ```

1. Regenerate font cache

    ```shell
    fc-cache -vf
    ```

1. Restart Chrome
1. Try it on [unicode full emoji list](https://unicode.org/emoji/charts/full-emoji-list.html)

## Reference

* [EmojiOne](https://www.emojione.com/)
* [How To Enable Color Emoji on Chrome for Linux](https://www.omgubuntu.co.uk/2016/08/enable-color-emoji-linux-google-chrome-noto)
* [Twitter Color Emoji SVGinOT Font](https://github.com/eosrei/twemoji-color-font)
