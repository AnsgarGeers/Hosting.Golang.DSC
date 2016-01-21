package dsc

import(
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "reflect"
)

// File
type File struct {
    // Content specifies the byte content of the file to enforce. Content is only
    // enforced if the Ensure field is set to 'present'
    Content []byte

    // Contentfunc specifies a user defined function that returns byte content
    // to enforce. Contentfunc is only enforced if the Ensure field is set to 'present'.
    // If Contentfunc is supplied the Content field is ignored. Contentfunc is passed
    // the File Resource instance as a function argument.
    Contentfunc func(Resource) ([]byte, error)

    // Ensure specifies the enforcement action of the File resource. Valid options are
    // 'present', 'absent'. This Field is required
    Ensure string

    // Path specifies the filesystem path of the file to enforce. This field is
    // required
    Path string
}

// Apply
func (this *File) Apply() (bool, error) {
    debug(this, 1, "Applying Resource")

    // Run ValidateFields
    fieldsbool, fieldserr := this.ValidateFields()
    if !fieldsbool {
        return fieldsbool, fieldserr
    }

    // Run check, if check returns true we do not need to make any changes
    check, _ := this.Check()
    if check {
        return true, nil
    }

    if this.Ensure == "present" {
        parentpath := filepath.Dir(this.Path)

        if !pathExists(parentpath) {
            os.MkdirAll(parentpath, 0644)
            debug(this, 2, "Created Parent Path "+parentpath)
        }

        if this.Contentfunc != nil {
            funcdata, err := this.Contentfunc(this)
            if err != nil {
                return false, err
            }
            err = ioutil.WriteFile(this.Path, funcdata, 0644)
            if err != nil {
                return false, err
            }
            return true, nil
        } else {
            err := ioutil.WriteFile(this.Path, this.Content, 0644)
            if err != nil {
                return false, err
            }
            return true, nil
        }
    } else if this.Ensure == "absent" {
        err := os.Remove(this.Path)
        if err != nil {
            return false, err
        } else {
            return true, nil
        }
    }
    return false, nil
}

// Check
func (this *File) Check() (bool, error) {
    if this.Ensure == "present" {
        if !pathExists(this.Path) {
            return false, errors.New("The path does not exist or is invalid")
        }

        // retrieve md5 of this.path
        filemd5, err := computeFileMd5(this.Path)
        if err != nil {
            debug(this, 1, err.Error())
            return false, err
        }

        // begin content check
        if this.Contentfunc != nil {
            funcdata, err := this.Contentfunc(this)
            if err != nil {
                return false, err
            }
            if reflect.DeepEqual(filemd5, computeByteMd5(funcdata)) {
                return true, nil
            } else {
                return false, errors.New("Path Content MD5 does not match Content Function MD5")
            }
        } else {
            if reflect.DeepEqual(filemd5, computeByteMd5(this.Content)) {
                return true, nil
            } else {
                return false, errors.New("Path Content MD5 does not match Content MD5")
            }
        }
    } else if this.Ensure == "absent" {
        pathexists := pathExists(this.Path)
        if !pathexists {
            return true, nil
        } else if pathexists {
            return false, errors.New("Resource expected to be absent but exists")
        }
    }

    return false, errors.New("Unknown failure condition.")
}

// Name
func (this *File) Name() string {
    return fmt.Sprintf("File %v", this.Path)
}

// ValidateFields
func (this *File) ValidateFields() (bool, error) {
    // make sure our ensure field is proper
    if this.Ensure != "present" && this.Ensure != "absent" {
        return false, errors.New("Invalid value to 'ensure' resource field. Valid options are 'present' and 'absent'")
    }

    // make sure our path field is not nil
    if &this.Path == nil || this.Path == ""{
        return false, errors.New("Invalid value to 'path' resource field. Path must be supplied")
    }

    return true, nil
}
