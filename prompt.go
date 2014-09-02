package goprompt

import (
    "fmt"
    "os"
    "io"
    "bufio"
)

type Result struct {
    Name string
    Value string
    DefaultValue string
    Children map[string]Result
}

func NewResult(name string) Result {
    result := Result {
        Name: name,
        Children: make(map[string]Result),
    }
    return result
}

type Field struct {
    Name string
    Title string
    Value string
    DefaultValue string
    form *Form
    Instructions string
    prompted bool
}

func (field *Field)IsPending() bool {
    return field.Value == "" && field.DefaultValue == ""
}

func (field *Field) Process() Result {
    result := NewResult(field.Name)
    for {
        if field.DefaultValue != "" {
            // print prompt for optional field, includes default value
            field.form.prompt.printf("%s(%s): ", field.Title, field.DefaultValue)
        } else {
            // ask for required field
            field.form.prompt.printf("%s: ", field.Title)
        }
        if field.form.prompt.scanner.Scan() {
            txt := field.form.prompt.scanner.Text()
            field.prompted = true
            field.Value = txt
            if field.IsPending() {
                // user skipped the field
                // print help and ask for a value again
                field.form.prompt.printfl(field.Instructions)

                // iterate again in this input field loop
                continue
            } else {
                if field.Value == "" {
                    field.Value = field.DefaultValue
                }
                // exit this input field loop
                break
            }
        } else { // CONTROL-C
            os.Exit(2)
        }
    } // field input loop
    result.Value = field.Value
    return result
}

type Form struct {
    Name string
    Title string
    Fields []*Field
    prompt *Prompt
}

func (f *Form) Process() Result {
    result := NewResult(f.Name)
    for i, field := range f.Fields {
        field.form = f
        if field.Name == "" {
            field.Name = fmt.Sprintf("field.%d", i)
        }
        if !field.prompted {
            fieldResult :=field.Process()
            result.Children[fieldResult.Name] = fieldResult
        }
    }
    return result
}

func (f *Form) PrintIntro() {
    f.prompt.printfl(f.Title)
    f.prompt.printfl("")
}

type Prompt struct {
    Output io.Writer
    Input io.Reader
    Forms []*Form
    scanner *bufio.Scanner
}

func (p *Prompt) GetOutput() io.Writer {
    output := p.Output
    if output != nil {
        return output
    }
    return os.Stdout
}

func (p *Prompt) GetInput() io.Reader {
    input := p.Input
    if input != nil {
        return input
    }
    return os.Stdin
}

func (p *Prompt) printf(format string, a ...interface{}) {
    output := p.GetOutput()
    fmt.Fprintf(output, format, a...)
}

func (p *Prompt) printfl(format string, a ...interface{}) {
    p.printf(format + "\n", a...)
}

func (p *Prompt) Process() Result {
    result := NewResult("")
    input := p.GetInput()
    p.scanner = bufio.NewScanner(input)
    for i, form:= range p.Forms {
        form.prompt = p
        if form.Name == "" {
            form.Name = fmt.Sprintf("form.%d", i)
        }
        form.PrintIntro()
        formResult := form.Process()
        result.Children[form.Name] = formResult
    }
    return result
}
