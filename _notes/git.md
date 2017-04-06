# Git

## Commit Message

[A Note About Git Commit Messages@tpope](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)

## Git Mirror

```bash
git clone --mirror gitolite@svn.garena.tw:GarenaTW/Authentication
cd Authentication.git/
git remote set-url --push origin git@gitlab.garena.tw:GarenaTW/Authentication.git
git fetch -p origin
git push --mirror
```

[Duplicating a repository@github.com](https://help.github.com/articles/duplicating-a-repository/)
