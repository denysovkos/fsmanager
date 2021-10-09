# FSManager - Tree view

Simple util to displays the directory structure of a path or of the disk in a drive graphically. If you don't specify a drive or path, this command displays the tree structure beginning with the current directory of the current drive.

# Hot to build

1. Clone this repo
2. Run ``go
env GOOS=linux GOARCH=arm go build
`` Where GOOS your operation system, and GOARCH is architecture of your processor
3. In the folder you will find executable file, so you can start to use

Tip:
```
On Linux and MacOS systems better build command:
go build -ldflags "-s -w"
```

# How to install

1. Run `go install`
2. Then will be available as regular program

# How to use

Just run `fsmanager <PATH_FROM>`

Path from, mean absolute path to the directory. If not specified, will be used path, where fsmanager was executed
