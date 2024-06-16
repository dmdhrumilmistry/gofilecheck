# GoFileCheck

Microservice written in GoLang for checking file types using magic bytes

## Installation

### Using GoBinaries

* Install using gobinaries

```bash
curl -sf https://gobinaries.com/dmdhrumilmistry/gofilecheck | sh
```

* Start API

```bash
gofilecheck
```

### Using Golang and Github

* Get and build project

```bash
go install github.com/dmdhrumilmistry/gofilecheck/cmd/gofilecheck@latest
```

* Start API

```bash
gofilecheck
```

### Using Docker

* Download docker file and start containers

```bash
docker compose up -d
```


## Usage

* Send file to API

```bash
curl --location 'http://localhost:3000/api/check' \
--form 'file=@"/path/to/png/file.png"'
```

Response
```json
{
    "extension": ".png",
    "filename": "gofilecheck.png",
    "size": 315659,
    "type": "image/png"
}
```
