<h1 align="center">Silicon Cloud CLI</h1>

<p align="center">
<p>

The Silicon Cloud CLI is an open source tool, you can get the latest version from [GitHub](https://github.com/siliconflow/siliconcloud-cli).

## Introduction
The Silicon Cloud CLI is a command line tool for managing your files on silicon cloud. It provides an easy way to upload,  and manage your silicon cloud files.

## CLI Releases

All releases please [click here](https://github.com/siliconflow/siliconcloud--cli/releases).

## Installation
SiliconCloud-CLI is available on Linux, macOS and Windows platforms.
Binaries for Linux, Windows and Mac are available as tarballs in the [release page](https://github.com/siliconcloud-/siliconcloud--cli/releases).

- Linux
  ```shell
    VERSION=0.0.1
    tar -xzvf siliconcloud-cli-linux-$VERSION-amd64.tar.gz
    install siliconcloud-cli /usr/local/bin
  ```

* Via a GO install

  ```shell
  # NOTE: The dev version will be in effect!
  go install github.com/siliconflow/siliconcloud-cli@latest
  ```

## Building From Source

SiliconCloud-CLI is currently using GO v1.22.X or above.
In order to build it from source you must:

1. Clone the repo
2. Build and run the executable

     ```shell
     make build && ./execs/siliconcloud
     ```

---

## Quick start

### Login
The Silicon Cloud CLI uses api-keys to authenticate client. To login your machine, run the following CLI:

```bash
siliconcloud login
# or using an environment variable
siliconcloud login -k $API_KEY
```

### Logout
To logout your machine, run the following CLI:

```bash
siliconcloud logout
```

### Upload files
To upload files to the silicon cloud, run the following CLI:

```bash
siliconcloud upload -n mymodel -t checkpoint -p /local/path/directory-or-file
```

You can specify model name, model type and path to upload by using the `-n`, `-t` and `-p` flags respectively.

### View Models
To view all your models in the silicon cloud, run the following CLI:

```bash
siliconcloud model ls -t checkpoint
```

You must specify model type by using the `-t` flag.

### View Model Files
To view all files in a model, run the following CLI:

```bash
siliconcloud model ls-files -n mymodel -t checkpoint
```

You must specify model name and model type by using the `-n` and `-t` flags respectively.

### Remove File
To remove file from a model, run the following CLI:

```bash
siliconcloud model rm-file -n mymodel -t checkpoint -f file_name
```

You must specify model name, model type and file name by using the `-n`, `-t` and `-f` flags respectively.

### Remove Model
To remove model from the silicon cloud, run the following CLI:

```bash
siliconcloud model rm -n mymodel -t checkpoint
```

You must specify model name and model type by using the `-n` and `-t` flags respectively.
