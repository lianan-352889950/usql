package metacmd

import (
	"github.com/xo/usql/text"
)

// Metacmd represents a command and associated meta information about it.
type Metacmd uint

// Decode converts a command name (or alias) into a Runner.
func Decode(name string, params []string) (Runner, error) {
	mc, ok := cmdMap[name]
	if !ok || name == "" {
		return nil, text.ErrUnknownCommand
	}

	cmd := cmds[mc]
	if cmd.Min > len(params) {
		return nil, text.ErrMissingRequiredArgument
	}

	return RunnerFunc(func(h Handler) (Result, error) {
		p := &Params{h, name, params, Result{}}
		err := cmd.Process(p)
		return p.Result, err
	}), nil
}

// Command types.
const (
	// None is an empty command.
	None Metacmd = iota

	// Question is question meta command (\?)
	Question

	// Quit is the quit meta command (\?).
	Quit

	// Copyright is the copyright meta command (\copyright).
	Copyright

	// Connect is the connect meta command (\c, \connect).
	Connect

	// Disconnect is the disconnect meta command (\Z).
	Disconnect

	// Password is the change password meta command (\password).
	Password

	// ConnectionInfo is the connection info meta command (\conninfo).
	ConnectionInfo

	// Drivers is the driver info meta command (\drivers).
	Drivers

	// Describe is the describe meta command (\d and variants).
	Describe

	// Print is the print query buffer meta command (\p, \print, \raw).
	Print

	// Reset is the reset query buffer meta command (\r, \reset).
	Reset

	// Transact is the transaction meta command (\begin, \commit, \rollback).
	Transact

	// Prompt is the variable prompt meta command (\prompt).
	Prompt

	// SetVar is the set variable meta command (\set).
	SetVar

	// Unset is the variable unset meta command (\unset).
	Unset

	// SetFormatVar is the set format variable meta commands (\pset, \a, \C, \f, \H, \t, \T, \x).
	SetFormatVar
)
