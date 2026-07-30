# 🔐 Pkeys

A Go CLI tool for AES-256 GCM encryption — encrypt and decrypt text, files, or your clipboard, right from the terminal.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue)
![Platforms](https://img.shields.io/badge/Platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

## Table of Contents
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [File Loading](#file-loading)
- [License](#license)

## 📦 Installation

Prebuilt releases are available for Mac, Linux and Windows.

To build yourself:

Clone the repo
```bash
git clone https://github.com/JbrownWFU/pkeys
```

and build
```bash
cd pkeys
go build -o pkeys
```

## ⚙️ Configuration

Pkeys relies on a `PKEYS_KEY` environment variable to store and retrieve your key. Instructions for setting your key can be found in the [generate](#generate) command.

## 🚀 Usage

Pkeys supports the encryption and decryption of text, files, and clipboard contents.

Most commands support redirecting their output to files with the `-o/--output` flag. Files are given default safe names unless a name is specified.

> If both `-o` and `-c` are passed to `encrypt`/`decrypt`, `-o` takes precedence and the clipboard is left untouched.

### Generate
Generate a new unique key
```bash
pkeys generate
```

Use the `-s` flag to generate a key with instructions to set based on your platform.
```bash
pkeys generate -s
```

Redirect to file
```bash
pkeys generate -o key.txt
```

### Encrypt
Encrypt passed text content
```bash
pkeys encrypt "my secret"
```

Use the `-f` flag with a filepath to instead encrypt a file. See [File Loading](#file-loading) for more information.
```bash
pkeys encrypt -f ./docs/mySecret.md
```

Pass the `-c` flag instead to encrypt the contents of your clipboard and overwrite the existing value.
```bash
pkeys encrypt -c
```

Redirect to file
```bash
pkeys encrypt "my secret" -o secret.enc
pkeys encrypt -f mySecret.md -o mySecret.enc
```

### Decrypt
Decrypt passed ciphertext content
```bash
pkeys decrypt <ciphertext>
```

Use the `-f` flag with a filepath to instead decrypt a file. See [File Loading](#file-loading) below.
```bash
pkeys decrypt -f ./docs/mySecret.enc
```

Pass the `-c` flag instead to decrypt the contents of your clipboard and overwrite the existing value.
```bash
pkeys decrypt -c
```

Redirect to file
```bash
pkeys decrypt <ciphertext> -o secret.md
pkeys decrypt -f mySecret.enc -o mySecret.md
```

## 📁 File Loading

File contents are entirely loaded into memory during encryption/decryption.

## 📄 License

This project is licensed under the MIT License.
