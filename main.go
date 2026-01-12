package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateField(fieldName string, fieldValue string) string {
	return fmt.Sprintf("%s=%s\n", fieldName, fieldValue)
}

func CreateDesktopEntryFile(filename string, content string) {
	err := os.WriteFile(fmt.Sprintf("%s.desktop", filename), []byte(content), 0644)

	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(1)
	}
}

func GenerateContent(entry *DesktopEntry) string {
	builder := new(strings.Builder)

	builder.WriteString("[Desktop Entry]\n")
	builder.WriteString(CreateField("Name", entry.Name))
	builder.WriteString(CreateField("Type", entry.Type))
	builder.WriteString(CreateField("Exec", entry.Exec))

	if entry.GenericName != "" {
		builder.WriteString(CreateField("GenericName", entry.GenericName))
	}

	if entry.Comment != "" {
		builder.WriteString(CreateField("Comment", entry.Comment))
	}

	if len(entry.Categories) > 0 {
		builder.WriteString(CreateField("Categories", strings.Join(entry.Categories, ";")))
	}

	if len(entry.Keywords) > 0 {
		builder.WriteString(CreateField("Keywords", strings.Join(entry.Keywords, ";")))
	}

	if entry.Path != "" {
		builder.WriteString(CreateField("Path", entry.Path))
	}

	if entry.Icon != "" {
		builder.WriteString(CreateField("Icon", entry.Icon))
	}

	builder.WriteString(CreateField("Terminal", strconv.FormatBool(entry.TerminalEnabled)))
	builder.WriteString(CreateField("Hidden", strconv.FormatBool(entry.HiddenEnabled)))
	builder.WriteString(CreateField("StartupNotify", strconv.FormatBool(entry.StartupNotifyEnabled)))
	builder.WriteString(CreateField("NoDisplay", strconv.FormatBool(entry.NoDisplayEnabled)))

	return builder.String()
}

func main() {
	entry := DesktopEntry{}

	fmt.Println("==== Desktop Entry Creator ====")
	filename := PromptRequired("Desktop File Name")
	entry.Name = PromptRequired("Application Name")
	entry.Exec = PromptRequired("Exec command (Full path or command Name)")
	entry.Type = PromptWithDefault("Application Type (Application/Link/Directory)", "Application")
	entry.Icon = PromptWithDefault("Icon", "''")

	fillOptionals := PromptYesNo("\nAll required fields are set. Do you want to fill optionals fields?")

	if fillOptionals {
		fmt.Println()
		entry.GenericName = PromptOptional("Generic Name")
		entry.Comment = PromptOptional("Comment")
		entry.TerminalEnabled = PromptYesNo("Is a terminal application?")
		entry.HiddenEnabled = PromptYesNo("Should be hidden?")
		entry.NoDisplayEnabled = PromptYesNo("Is this app display enabled?")
		entry.StartupNotifyEnabled = PromptYesNo("Should startup notification be enabled?")
	}

	content := GenerateContent(&entry)
	CreateDesktopEntryFile(filename, content)

	fmt.Println("File generated")
}
