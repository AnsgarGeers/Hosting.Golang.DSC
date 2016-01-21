package main

import(
    "log"
    "../../Hosting.Golang.DSC/v1"
)

func main() {
    log.Println("Hello, World")

    dsc.LOG_LEVEL = 1

    /*folder := &dsc.Folder{
        Ensure: "absent",
        Path: "c:/tmp/",
        Recurse: true,
    }
    log.Println(folder.Apply())*/

    file := &dsc.File{
        Ensure: "present",
        Path: "c:/tmp/testfile.txt",
        Content: []byte("Some important text"),
    }
    log.Println(file.Apply())

    exec := &dsc.Exec{
        Command: "systeminfo",
    }
    exec.Apply()
}
