


[![Build Status](https://travis-ci.org/mchirico/datagrab.svg?branch=master)](https://travis-ci.org/mchirico/datagrab)
[![codecov](https://codecov.io/gh/mchirico/datagrab/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/datagrab)

[![Build Status](https://mchirico.visualstudio.com/datagrab/_apis/build/status/mchirico.datagrab?branchName=master)](https://mchirico.visualstudio.com/datagrab/_build/latest?definitionId=9&branchName=master)


# datagrab



### Checklist:

1. dockerPassword
2. [CodeCov Token](https://codecov.io/gh/mchirico)
3. No Caps in project
4. MONGO_CONNECTION_STRING
5. MONGO_DATABASE 
6. Make Azure Boards Public
7. More Cobra commands? (cobra add say)



## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


# datagrab
