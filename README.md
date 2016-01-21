[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/HOSTINGLabs/Hosting.Golang.DSC/v1)

# Golang Desired State Configuration

**this package is experimental and its api has not yet been solidified. Please feel free to contribute**

Golang DSC is a package that implements a Desired State Configuration (a.k.a Configuration Enforcement) primitives. Golang DSC is inspired from Puppet's Configuration DSL and Microsoft Powershell DSC.

For example, let us say instead of writing a file with the traditional boiler plate, we can more accurately describe our file write in a DSL-like syntax:

```go
package main

import(
    "log"
    "github.com/hostinglabs/Hosting.Golang.DSC/v1"
)

func main(){
    // Define our resource
    fileresource := &dsc.File{
        Path: "/tmp/myfile",
        Content: "Hello, World",
        Ensure: "Present",
    }

    // Apply our resource
    err := dsc.Apply(fileresource);
    if err != nil {
        log.Fatal(err)
    }
}
```

In the above example Golang DSC will take care of all the logic to ensure that the file at _Path_ "/tmp/myfile" will _Ensure_ "present" with _Content_ of "Hello World". You do not need to write the complex logic to handle the process of writing this file. If the file already exists with the proper content no action will be taken by applying the resource.

# Built-In and Custom Resources

Golang DSC provides a number of pre-built useful resources for your consumption:

* dsc.File
* dsc.Folder
* dsc.Exec

Creating your own custom resources is as simple as implementing the dsc.Resource interface:

```go
package main

import(
    "log"
    "github.com/hostinglabs/Hosting.Golang.DSC/v1"
)

// Define your resource type
type MyResource struct {
    // Your resource accepts a string parameter
    MyParam1 string

    // Your resource accepts a bool parameter
    MyParam2 bool
}

// Implement the Apply method of dsc.Resource
func (mr *MyResource) Apply() (bool, error) {
    return true, nil
}

// Implement the Check method of dsc.Resource
func (mr *MyResource) Check() (bool, error) {
    return true, nil
}

// Implement the Name method of dsc.Resource
func (mr *MyResource) Name() string {
    return "MyResource"
}

// Implement the ValidateFields method of dsc.Resource
func (mr *MyResource) ValidateFields() (bool, error) {
    return true, nil
}

// main
func main() {
    // Define your resource
    myresource := &MyResource{
        MyParam1: "Hello",
        MyParam2: true,
    }

    // Apply your resource
    err := dsc.Apply(myresource)
    if err != nil {
        log.Fatal(err)
    }
}
```

# Package Versioning

Golang DSC implements versioning by providing subpackages:

* "github.com/hostinglabs/Hosting.Golang.DSC/v1"
* "github.com/hostinglabs/Hosting.Golang.DSC/v2"  <= future release not yet built
