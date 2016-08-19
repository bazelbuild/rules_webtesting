// Package errors provides an Error interface that includes the component name with the underlying error.
package errors

import (
	"errors"
	"fmt"
)

// DefaultComp is the default component name used when no other component specified.
const DefaultComp = "web test launcher"

type wtlError struct {
	error
	component string
	permanent bool
}

func (we *wtlError) Component() string {
	return we.component
}

func (we *wtlError) Permanent() bool {
	return we.permanent
}

// Component returns err.Component() if err implements it, otherwise it returns DefaultComp.
// Component is intended for grouping error messages by source.
func Component(err error) string {
	type componentError interface {
		Component() string
	}
	if ce, ok := err.(componentError); ok {
		return ce.Component()
	}
	return DefaultComp
}

// IsPermanent returns err.Permanent() if err implements it, otherwise it returns false.
// IsPermanent is intended for preventing retries when errors occur that are unlikely to go away.
func IsPermanent(err error) bool {
	type permanentError interface {
		Permanent() bool
	}

	pe, ok := err.(permanentError)
	return ok && pe.Permanent()
}

func (we *wtlError) Error() string {
	p := ""
	if we.permanent {
		p = " (permanent)"
	}
	return fmt.Sprintf("[%s%s]: %v", we.component, p, we.error)
}

// New returns an error, e, such that:
//   Permanent(e) is false
//   If Component(err) is not DefaultComp, Component(e) is Component(err)
//   Else If component is "", Component(e) is DefaultComp
//   Else Component(e) is component.
func New(component string, err interface{}) error {
	return createErr(component, err, false)
}

// NewPermanent returns an error, e, such that:
//   Permanent(e) is true
//   If Component(err) is not DefaultComp, Component(e) is Component(err)
//   Else If component is "", Component(e) is DefaultComp
//   Else Component(e) is component.
func NewPermanent(component string, err interface{}) error {
	return createErr(component, err, true)
}

func createErr(component string, err interface{}, permanent bool) error {
	e := func() error {
		switch t := err.(type) {
		case error:
			return t
		case string:
			return errors.New(t)
		default:
			return fmt.Errorf("%v", err)
		}
	}()

	ec := Component(e)
	ep := IsPermanent(e)

	if ec != DefaultComp {
		component = ec
	}

	if component == "" {
		component = DefaultComp
	}

	if ep == permanent && ec == component {
		return e
	}

	if we, ok := e.(*wtlError); ok {
		return createErr(component, we.error, permanent)
	}

	return &wtlError{
		error:     e,
		component: component,
		permanent: permanent,
	}
}
