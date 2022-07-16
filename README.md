[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/summary/new_code?id=RobertOchmanek_ebiznes)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=RobertOchmanek_ebiznes&metric=bugs)](https://sonarcloud.io/summary/new_code?id=RobertOchmanek_ebiznes)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=RobertOchmanek_ebiznes&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=RobertOchmanek_ebiznes)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=RobertOchmanek_ebiznes&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=RobertOchmanek_ebiznes)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=RobertOchmanek_ebiznes&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=RobertOchmanek_ebiznes)

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