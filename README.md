# nomnom

## Motivation

Custom types and constants are often used in Go for defining enums:

```golang
type ProgrammingLanguage string

const (
	ProgrammingLanguageGo        ProgrammingLanguage = "Go"
	ProgrammingLanguageJava      ProgrammingLanguage = "Java"
	ProgrammingLanguageBrainfuck ProgrammingLanguage = "Brianfuck"
)
```

This allows type safety when using those values:

```golang
type Project struct {
  ImplementedLanguage ProgrammingLanguage
}

var newProject = Project{
  ImplementedLanguage: "Japanese", // will not compile
}
```

However, it's possible to get around these restrictions:

```golang
var newProject = Project{
  ImplementedLanguage: ProgrammingLanguage("Singlish"), // compiles but not a proper enum value
}
```

While it's possible to write functions to check against possible values, it can be cumbersome when
there are many possible enum values:

```golang
func IsProgrammingLanguage(in string) bool {
	switch in {
	case "Go":
		return true
	case "Java":
		return true
	case "Brainfuck":
		return true
	}

	return false
}
```

Furthermore, we'd have to create new checker functions for each type.

nomnom strives to make this usage pattern more maintainable by generating helper functions from
existing enum definitions. It also generates unit tests for the generated functions. See
[fixtures](/gen/fixtures) for examples of generated helper functions and their corresponding tests.
