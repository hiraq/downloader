# Golang Downloader
Golang cli apps to download files from list of urls

# How To Build

```
go build --race
```

Please make sure to install golang in your os!

# How To Use

Create a file (let say, `images.txt`) and the content like :

```
http://2.bp.blogspot.com/-SxX9MXXc4YU/T6BsKJbk6JI/AAAAAAAAAMc/IiQWbkoQtL0/s1600/goraster.png
https://golang.org/doc/gopher/bumper.png
https://pbs.twimg.com/media/BsTtrWiCcAADk5W.jpg
https://lh6.googleusercontent.com/-WhYQjo7IgYc/TunZIZbncTI/AAAAAAAAAE8/lIMOU4bqlXQ/w800-h800/gopher-wallpaper.png
https://lh4.googleusercontent.com/-KvhHWmaPiAg/T434KqiQSkI/AAAAAAAAGTA/z4xcKQCbHyA/w800-h800/Go_code.png
```

After that, use command lines :

```
./downloader -fl /path/to/your/images.txt -sp /path/to/save/downloaded/files/
```

Or if you not build it yet :

```
go run --race *.go -fl /path/to/your/images.txt -sp /path/to/save/downloaded/files/
```

Please dont forget the last trailing slash in `sp` param.

# Rationale

All registered files should be downloaded concurrently not sequential, this will help us to
get more better performance.

Be carefull with goroutines and channels guys, please make sure you test and build with `--race`
parameter to check for race condition.

# Notes

My motivation to create this simple script just to have fun, so dont be too serious here, if you think you can add
more functionality (maybe like hooks?) just fork it, okay?
