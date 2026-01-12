package main

type DesktopEntry struct {
	Name                 string
	Type                 string
	Exec                 string
	GenericName          string
	Comment              string
	Categories           []string
	Keywords             []string
	Path                 string
	Icon                 string
	TerminalEnabled      bool
	HiddenEnabled        bool
	StartupNotifyEnabled bool
	NoDisplayEnabled     bool
}
