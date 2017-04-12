# Git

## Commit Message

[A Note About Git Commit Messages@tpope](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)

## Git Mirror

最近想把公司的舊專案慢慢從 Gitolite 換到 Gitlab 上。

研究了一下 git clone mirror 選項，它可以把整個 repo 的資訊都 clone 下來，包括 branches, tags。

如果你有連進去 git server 看過遠端 repo 怎麼存的話，它就是把那整個資料夾給載下來。

我們再用 git push \--mirror 把整包都丟到新的 git server 上面的 repo 裡。

### Example

1. 整包 clone 下來

    ```bash
    git clone --mirror gitolite@git.lab317.org:dinos80152/Authentication
    ```

1. 到 gitlab 建一個空的 repo
1. 進到專案資料夾，設定新的遠端 git repo 位置

    ```bash
    cd Authentication.git/
    git remote set-url --push origin git@gitlab.lab317.org:dinos80152/Authentication.git
    ```

1. local 更新 remote branch，因為是新的 repo，就是將本地的 origin/xxx 都刪囉。-p == \--prune

    ```bash
    git fetch -p origin
    ```

1. 最後將整包 push 上去

    ```bash
    git push --mirror
    ```

1. 可以在 Gitlab 上看到 repo 裡有以前在 Gitolite 的所有紀錄與東西囉。

## Reference

* [Duplicating a repository@github.com](https://help.github.com/articles/duplicating-a-repository/)