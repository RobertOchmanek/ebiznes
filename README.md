# Backend:

## GCC compiller is required to run backend server. To install, follow this guide:
- https://code.visualstudio.com/docs/cpp/config-mingw

## To run backend server execute the following commands in `backend` directory:

- `go mod init (repository name)`, in this case `go mod init github.com/RobertOchmanek/ebiznes_go`
- `go get github.com/labstack/echo/v4`
- `go get github.com/mattn/go-sqlite3`
- `go get -u github.com/jinzhu/gorm`
- `go build main.go`
- `./main`


# Frontend:

## To run frontend client execute the following commands in `frontend` directory:

- `npm start`