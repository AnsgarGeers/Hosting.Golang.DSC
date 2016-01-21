package dsc

import(
    "errors"
    "fmt"
    "os"
)

// Folder
type Folder struct{
    // Ensure specifies the enforcement action of the Folder resource. Valid options are
    // 'present', 'absent'. This Field is required
    Ensure string

    // Path specifies the filesystem path of the folder to enforce. This field is
    // required
    Path string

    // Recurse specifies that during Ensure "absent" we are to remove all children
    // from the provided Path
    Recurse bool
}

// Apply
func (this *Folder) Apply() (bool, error) {
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
        mkdirerr := os.MkdirAll(this.Path, 0777)
        if mkdirerr != nil {
            debug(this, 1, mkdirerr.Error())
            return false, mkdirerr
        }
        return true, nil
    } else if this.Ensure == "absent" {
        var rmerr error
        if this.Recurse {
            rmerr = os.RemoveAll(this.Path)
        } else {
            rmerr = os.Remove(this.Path)
        }
        if rmerr != nil {
            debug(this, 1, rmerr.Error())
            return false, rmerr
        }
        return true, nil
    }

    return false, errors.New("Unknown failure condition.")
}

// Check
func (this *Folder) Check() (bool, error) {
    if this.Ensure == "present" {
        if !pathExists(this.Path) {
            return false, errors.New("The path is expected to exist but does not")
        } else if pathExists(this.Path) {
            return true, nil
        }
    } else if this.Ensure == "absent" {
        if pathExists(this.Path) {
            return false, errors.New("The path is expected not to exist but it is present")
        } else if !pathExists(this.Path) {
            return true, nil
        }
    }

    return false, errors.New("Unknown failure condition.")
}

// Name
func (this *Folder) Name() string {
    return fmt.Sprintf("Folder %v", this.Path)
}

// ValidateFields
func (this *Folder) ValidateFields() (bool, error) {
    // make sure our ensure field is proper
    if this.Ensure != "present" && this.Ensure != "absent" {
        return false, errors.New("Invalid value to 'Ensure' resource field. Valid options are 'present' and 'absent'")
    }

    // make sure our path field is not nil
    if &this.Path == nil || this.Path == "" {
        return false, errors.New("Invalid value to 'Path' resource field. Path must be supplied")
    }

    return true, nil
}
