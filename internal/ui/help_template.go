package ui

var UsageTemplate = `
{{styleTitle " 💊 VITO CLI "}}
{{descStyle "\nUltravitamined CLI tooling for devs. \n"}}
{{styleTitle " USAGE "}}
  {{styleCommand "vito [pill]"}}{{descStyle " Executes a specific PILL (command)."}}
  {{styleCommand "vito [flag]"}}{{descStyle " Executes VITO with a flag."}}

{{styleTitle " PILLS "}}{{range .Commands}}{{if (not .IsAvailableCommand)}} {{continue}} {{end}}{{if (and (not .Hidden) (ne .Name "help"))}}
  {{rpad (styleCommand .Name) 15}} {{descStyle .Short}}{{end}}{{end}}

{{styleTitle " FLAGS "}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

{{descStyle "Use \"vito [command] --help\" for more info about a specific pill."}}
`
