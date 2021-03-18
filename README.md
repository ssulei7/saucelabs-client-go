[![Go Build Validation](https://github.com/ssulei7/saucelabs-client-go/actions/workflows/go.yml/badge.svg)](https://github.com/ssulei7/saucelabs-client-go/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/ssulei7/saucelabs-client-go/branch/main/graph/badge.svg?token=SSD5JHOL1P)](https://codecov.io/gh/ssulei7/saucelabs-client-go)

# Saucelabs Go REST Client

## Purpose

A general client to interface with the SauceLabs API with Go. 

## APIs Covered

### Builds

* Get all builds /rest/v1/user/builds

### Jobs

* Get all jobs /rest/v1/user/jobs


### Example

```go

func main() {
    //provide a sauce key and username in which to retrieve information from
    c := sauce.NewClient(os.Getenv("SAUCE_KEY"), os.Getenv("SAUCE_USER"), "base url, leave empty string for default")
    
    //retrieve builds, for example...
    builds, err := c.GetBuilds()
    
    if err == nil {
        //do something with builds slice...
    }
}

```

## TODO

Many items are still pending, including but not limited to...

* Operations on sauce types
* Remaining APIs
