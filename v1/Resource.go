package dsc

// Resource
type Resource interface {
    // Apply processes the resource structure and attempts to enforce the resource
    // state. Should an error occur during enforcement an error interface is returned
    // describing the problem.
    Apply() (bool, error)

    // Check processes the resource structure and checks if the resource is in compliance
    // with the resource definition. Check will return a bool indicating if the resource
    // is in compliance with its structure, and potentially an error describing any failures
    // during the check process
    Check() (bool, error)

    // Name provides a unique name for the specified resource for use in tracking
    // or logging
    Name() string

    // ValidateFields processes the resource structure and ensures that supplied values
    // are valid and proper.
    ValidateFields() (bool, error)
}
