# File-Searcher

A CLI-based tool that imitates basic grep functionality. Simple, powerful, made in Go.  

## Installation

**Pre-requisites**: Ensure you have the go compiler installed on your machine.

1. **Clone the repository**

```shell
git clone https://github.com/Prateik-Lohani-07/file-searcher.git
```

2. **Run the build script**

```shell
chmod +x ./build.sh
./build.sh
```

Now you will have an executable `filesearch.exe` in the folder.

3. **Add the output executable to your system path.**

If on linux, move the executable to `/usr/local/bin`. If on windows, add the file to your PATH variable.

On Windows, you may also move it to C drive in a separate folder with any other CLI-tools you may have are added to PATH already.

```shell
# linux
mv filesearch.exe /usr/local/bin
```

```cmd
mkdir C:\cli-tools
move filesearch.exe C:\cli-tools\
```

```powershell
# powershell
New-Item -ItemType Directory C:\cli-tools -Force
Move-Item filesearch.exe C:\cli-tools\
```

Now you can use the `filesearch` command as shown in the next section.

## Usage

```bash
filesearch --dir "<directory>" --query "<keyword>" [flags]
```

### Flags

- `--dir`: Specifies the directory in which search operation is to be performed.
- `--query`: Specifies the keyword that is to be searched for in the search directory.
- `--n`: Whether to show the line and column numbers where the query occurs in the file.

## Example

```shell
filesearch --dir . --query "restaurant"
```

Output:

```shell
[1,28]:Restaurant and ML.txt:   Taking all the data of the restaurants in my area,
```
