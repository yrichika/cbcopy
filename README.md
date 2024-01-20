# cbcopy

日本語の説明は下にあります。

## Description

The command copies the content entered with `|` to the clipboard. It functions similarly to Mac's `pbcopy` for WSL. Please refrain from installing it on Linux systems that are not WSL, as it won't work on standard Linux.

```sh
echo "some text to the clipboard" | cbcopy
# this works too.
cbcopy "some text to the clipboard"
```

## 説明

`|`で入力した内容をクリップボードにコピーするコマンドです。
WSL用にMacの`pbcopy`と同様の動きをするコマンドです。
WSLでないLinuxでは動かないため、WSLでないLinuxにはインストールしないでください。


```sh
echo "some text to the clipboard" | cbcopy
# 引数の値もクリップボードにコピーできます。
cbcopy "some text to the clipboard"
```

