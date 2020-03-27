# Latest Git Commit Id Generator 
generate a go file contains a git commit id.

It is good to generate a go file in a go http server project, that the http server can publish a url to show it's version.

## Install
Using go get
```go
go get github.com/pharosnet/gitcgen
```

Using download bin
* [Darwin](https://github.com/pharosnet/gitcgen/releases/download/v1.0.0/gitcgen_darwin_x64_v1.0.0.zip)
* [Linux](https://github.com/pharosnet/gitcgen/releases/download/v1.0.0/gitcgen_linux_x64_v1.0.0.zip)
* [Windows](https://github.com/pharosnet/gitcgen/releases/download/v1.0.0/gitcgen_win_x64_v1.0.0.zip)


## Usage

### Commands

| Name | short Name | Desc                                             |
| ---- | ---------- | ------------------------------------------------ |
| show | s          | show latest git commit id                        |
| gen  | g          | generate a go file contains latest git commit id |

### show command args:

| Name      | Short Name | Value  | Desc                    |
| --------- | ---------- | ------ | ----------------------- |
| work_tree | w          | string | the path of git project |

gen command args:

| Name      | Short Name | Value  | Desc                                                       |
| --------- | ---------- | ------ | ---------------------------------------------------------- |
| work_tree | w          | string | the path of git project                                    |
| output    | o          | string | the file path of generated go file, such as "./foo/bar.go" |
| short     | s          | bool   | use short git commit id. default is true.                  |



### Examples:

Show latest git commit id

```bash
gitcgen show -w="{TARGET_GIT_PROJECT_LOCAL_PATH}"
```

Generate a go file
```bash
gitcgen gen -w="{TARGET_GIT_PROJECT_LOCAL_PATH}" -o="./foo/bar.go"
```

Using in go generate

```go
//go:generate gitcgen gen -w="./" -o="./versions/git_latest_commit_id.go"
func main() {
    // ...
}
```

