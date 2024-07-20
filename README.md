# subtotal
It is a practical and simple tool for enumerating subdomains with VirusTotal's large database via a free API.
It is a passive recognition, a single HTTP GET request is sent to VT.

## Installation
```
go install github.com/TaveiraTavel/subtotal@latest
```

## Configuration
Sets `VT_API_KEY` environment variable to use.
https://www.virustotal.com/gui/my-apikey


## Usage
```
echo <domain> | subtotal
```
