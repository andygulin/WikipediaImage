### WikipediaImage 爬取【维基百科】每日图片

#### 爬取地址

- https://zh.wikipedia.org/wiki/Wikipedia:每日图片/2020年1月
- https://zh.wikipedia.org/wiki/Wikipedia:每日图片/2020年2月
- https://zh.wikipedia.org/wiki/Wikipedia:每日图片/2020年3月
- ......

#### 使用方法

```shell
go build

./WikipediaImage 2023 1
./WikipediaImage 2023 2
......
```

图片保存地址：当前目录下store_image文件夹

#### Git Large File
- https://git-lfs.github.com
```shell
git lfs track "*.jpg"
```