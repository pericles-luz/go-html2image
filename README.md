# go-html2image
Image generator from HTML source. It works only on linux and macOS. Must be adjusted to work on Windows.

## Installation
```bash
go get github.com/pericles-luz/go-html2image
```

It's necessary to install [wkhtmltoimage](https://wkhtmltopdf.org/downloads.html) on your system.

```bash
        sudo apt-get update
        sudo apt-get -y install xfonts-75dpi xfonts-base
        curl --silent --show-error --location --max-redirs 3 --fail --retry 3 --output wkhtmltopdf-linux-amd64.deb https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.jammy_amd64.deb
        sudo dpkg -i wkhtmltopdf-linux-amd64.deb
        sudo ldconfig
        rm wkhtmltopdf-linux-amd64.deb
```
