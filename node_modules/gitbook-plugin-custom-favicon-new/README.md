# gitbook-custom-favicon-new

将自己的图标添加到gitbook主题。
favicon.ico和apple-touch-icon-precomposed-152.png

插件删除位于`"_book/gitbook/images/favicon.ico"`的gitbook favicon，用自己的favicon代替。apple-touch-icon-precomposed-152.png同理
这属于暴力方法，如有更好的方法欢迎指正

## Install via gitbook

### In book.json

* 添加 `custom-favicon-new` 到 `plugins` 数组中
* 在`pluginsConfig`中配置`favicon`的路径

#### book.json
```json
{
	"plugins" : ["custom-favicon-new"],
	"pluginsConfig" : {
		"favicon": "path/favicon.ico",
		"appleTouchIconPrecomposed152": "path/logo.png"
	}
}
```

### using gitbook-cli

```bash
gitbook install
```



