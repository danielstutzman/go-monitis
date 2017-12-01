[![GoDoc](https://godoc.org/github.com/danielstutzman/go-monitis?status.svg)](https://godoc.org/github.com/danielstutzman/go-monitis)

# Go SDK for Monitis API

## Supported
* Authentication:
  * Get User AuthToken [(API docs)](http://www.monitis.com/docs/apiActions.html#getAuthToken)
* Contacts API:
  * Get Recent Alerts [(API docs)](http://www.monitis.com/docs/apiActions.html#getRecentAlerts)

## Installation
```
go get -u github.com/danielstutzman/go-monitis
```

## Command-line interface for testing
```
$GOPATH/bin/monitis-cli -apikey API_KEY_HERE -secretkey SECRET_KEY_HERE
```

## See also
* [Monitis REST API documentation](http://www.monitis.com/docs/api.html)
