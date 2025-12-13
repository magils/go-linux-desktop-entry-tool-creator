package main

type DesktopEntry struct {
	Name                 string
	GenericName          string
	Comment              string
	Exec                 string
	Type                 string
	Categories           []string
	Keywords             []string
	Path                 string
	Icon                 string
	TerminalEnabled      bool
	HiddenEnabled        bool
	StartupNotifyEnabled bool
	NoDisplayEnabled     bool
}
