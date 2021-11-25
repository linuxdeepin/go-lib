Go 图形处理库, 主要围绕 GdkPixbuf 库进行增强开发. 支持对图片进行
剪切, 翻转, 缩放, 混合, 模糊处理等操作.

API 的命名风格与 github.com/linuxdeepin/go-lib/graphic 库相同, 以 Blur 操作为例:
- **Blur** 对 C.GdkPixbuf 对象进行模糊处理

- **BlurImage** 对目标文件进行模糊处理

- **BlurImageCache** 对目标文件进行模糊处理, 同时将处理后的文件放到缓
  存目录, 下次对同一文件进行相同操作时可以大大提高速度
