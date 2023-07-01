## Požadavky
- [Go](https://go.dev) (min. 1.17)
- [Docker](https://docker.com) (na odeslání)

## Instalace
```sh
git clone https://github.com/Tour-de-App/fiber-boilerplate tda-fiber
cd tda-fiber
go get .
```

## Spuštění
### Linux a macOS
```sh
PORT=":3000" go run main.go
```

### Windows
```bat
set PORT=":3000" & go run main.go
```

Aplikace poběží na adrese [`localhost:3000`](http://localhost:3000)

### Docker spuštění
**Windows uživatelé spouštějte ve WSL terminálu**
```
docker build . -t tda-fiber
docker run -p 8080:80 tda-fiber
```

## Struktura
- Složka `views` obsahuje dynamické šablony stránek napsané v jazyce [Pug](https://pugjs.org)
- Soubor `main.go` obsahuje logiku (cesty atd.)
- Složka `public` obsahuje statické soubory (styly apod.)

Pro více informací si pročtěte [oficiální Fiber tutoriál](https://docs.gofiber.io/) (anglicky)
