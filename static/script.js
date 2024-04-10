// 获取元素引用
var textarea = document.getElementById('content');
var printable = document.getElementById('printable');

// 初始化内容
var content = textarea.value;

// 创建一个函数来定期检查和上传内容
function checkAndUploadContent() {
    var currentContent = textarea.value;
    if (currentContent !== content) {
        // 内容已更改，发送POST请求
        var request = new XMLHttpRequest();
        request.open('POST', window.location.href, true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(currentContent);

        // 更新printable的内容
        while (printable.firstChild) {
            printable.removeChild(printable.firstChild);
        }
        printable.appendChild(document.createTextNode(currentContent));

        // 更新存储的内容
        content = currentContent;
    }

    // 1秒后再次检查
    setTimeout(checkAndUploadContent, 1000);
}

// 开始定期检查和上传内容
checkAndUploadContent();