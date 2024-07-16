# subtotal
It is a practical and simple tool for enumerating subdomains with VirusTotal's large database via a free API.
It is a passive recognition, a single HTTP GET request is sent to VT.

## API-KEY
https://www.virustotal.com/gui/my-apikey

## Installation
```
go install github.com/TaveiraTavel/subtotal@HEAD
```

## Usage
```
echo example.org | subtotal <api-key>
```
