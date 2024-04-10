## Minimalist Web Notepad Golang

Minimalist Web Notepad Golang重置版 

最近我在github发现了一个不错的PHP项目[Minimalist Web Notepad](https://github.com/pereorga/minimalist-web-notepad)

用于临时记录与传输纯文本非常方便，简直是像我这种极简主义的最爱

于是我就用go重置了一下

## 运行教程

下载最新的Release，运行二进制文件即可

## 使用方法

1. 访问网页： [https://note.u-web.pp.ua/（演示站不定期崩坏）](https://note.u-web.pp.ua/)
2. 它会随机分配 10 个字符组成的地址，如 https://note.u-web.pp.ua/1234567890 ，如果想指定地址，只需要访问时手动修改，如 https://note.u-web.pp.ua/114514 。
3. 在上面编辑文本
4. 等待一会（几秒，取决于延迟），服务端就会存储网页内容到名为 114514 的文件里。
5. 关闭网页，如果关闭太快，会来不及保存，丢失编辑。
6. 在其他平台再访问同样的网址，就能剪切内容了 ٩۹(๑•̀ω•́ ๑)۶