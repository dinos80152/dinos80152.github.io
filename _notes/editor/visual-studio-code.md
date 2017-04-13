# Visual Studio Code

## Extension

* Go
* Vim
* Git History (git log)

## Preference

```json
// 將您的設定放入此檔案中以覆寫預設值
{
    "window.zoomLevel": 1,
    "vim.useCtrlKeys": false,
    "terminal.integrated.shell.windows":     "C:\\Program Files\\Git\\bin\\bash.exe"
}
```

## KeyBinding

```json
// Place your key bindings in this file to overwrite the defaults
[
    { "key": "shift+space",            "command": "editor.action.triggerSuggest",
                                     "when": "editorHasCompletionItemProvider && editorTextFocus && !editorReadonly" }
]
```

## Reference

* [Go programming in VS Code](https://code.visualstudio.com/docs/languages/go)