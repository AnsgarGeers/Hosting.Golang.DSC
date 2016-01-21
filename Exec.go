package dsc

import(
    "fmt"
    "os/exec"
    "strings"
)

// Exec
type Exec struct {
    // Command
    Command string

    // Args
    Args []string
}

// Apply
func (this *Exec) Apply() (bool, error) {
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

	out, err := exec.Command(this.Command, this.Args...).Output();
    if err != nil {
        debug(this, 1, err.Error())
		return false, err
	}
    debug(this, 2, string(out))

    return true, nil
}

// Check
func (this *Exec) Check() (bool, error) {
    // commands never hold state, so we always return false here
    return false, nil
}

// Name
func (this *Exec) Name() string {
    return fmt.Sprintf("Exec %v [%v]", this.Command, strings.Join(this.Args, ","))
}

// ValidateFields
func (this *Exec) ValidateFields() (bool, error) {
    return true, nil
}
