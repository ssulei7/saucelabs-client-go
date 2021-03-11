[![Go Build Validation](https://github.com/ssulei7/saucelabs-client-go/actions/workflows/go.yml/badge.svg)](https://github.com/ssulei7/saucelabs-client-go/actions/workflows/go.yml)

# Saucelabs Go REST Client

## Purpose

A general client that I am currently working on to pull job information within my CI/CD execution

## APIs Covered

### Builds

* Get all builds /rest/v1/user/builds

### Jobs

* Get all jobs /rest/v1/user/jobs

### Example

```go

func main() {
    //provide a sauce key and username in which to retrieve information from
    c := sauce.NewClient(os.Getenv("SAUCE_KEY", "username")
    
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
