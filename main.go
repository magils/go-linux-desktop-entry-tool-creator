package main

import "fmt"

func CreateDesktopEntryFile(entry *DesktopEntry) {

}

func main() {
	entry := DesktopEntry{}

	fmt.Println("==== Desktop Entry Creator ====")

	entry.Name = PromptRequired("Application Name")
	entry.Exec = PromptRequired("Exec command (Full path or command Name)")
	entry.Type = PromptWithDefault("Application Type (Application/Link/Directory)", "Application")
	entry.Icon = PromptWithDefault("Icon", "''")

	fillOptionals := PromptYesNo("\nAll required fields are set. Do you want to fill optionals fields?")

	if fillOptionals {
		fmt.Println()
		entry.GenericName = PromptOptional("Generic Name")
		entry.Comment = PromptOptional("Comment")
	}

	CreateDesktopEntryFile(&entry)
}
