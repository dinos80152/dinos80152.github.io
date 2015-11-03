## create repository in gitlab

![gitlab](http://dinos80152.github.io/_data/gitlab.png)

## Set push url

cd /xxx.garena.tw/.git

git remote set-url --push origin git@gitlab.garena.tw:GarenaTW/xxx.garena.tw.git

git push --mirror


## change remote url

本機、測試機、正式機已有 git

```
git remote set-url origin git@gitlab.garena.tw:GarenaTW/xxx.garena.tw.git
```

## add remote origin

測試機無 git

```
git init
```

```
git remote add origin git@gitlab.garena.tw:GarenaTW/xxx.garena.tw.git
```