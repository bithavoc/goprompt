package goprompt

import (
    "fmt"
    "os"
    "io"
    "bufio"
)

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

type Form struct {
    Title string
    Fields []*Field
    prompt *Prompt
}

func (f *Form) PrintIntro() {
    f.prompt.printfl(f.Title)
    f.prompt.printfl("")
}

type Prompt struct {
    Output io.Writer
    Input io.Reader
    Forms []*Form
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

func (p *Prompt) Process() {
    input := p.GetInput()
    for _, form:= range p.Forms {
        form.prompt = p
        form.PrintIntro()
        for _, field := range form.Fields {
            field.form = form
            if !field.prompted {
                scanner := bufio.NewScanner(input)
                for {
                    if field.DefaultValue != "" {
                        // print prompt for optional field, includes default value
                        p.printf("%s(%s): ", field.Title, field.DefaultValue)
                    } else {
                        // ask for required field
                        p.printf("%s: ", field.Title)
                    }
                    if scanner.Scan() {
                        txt := scanner.Text()
                        field.Value = txt
                        field.prompted = true
                        if field.IsPending() {
                            // user skipped the field
                            // print help and ask for a value again
                            p.printfl(field.Instructions)

                            // iterate again in this input field loop
                            continue
                        } else {
                            // exit this input field loop
                            break
                        }
                    } else { // CONTROL-C
                        os.Exit(2)
                    }
                } // field input loop
            }
        }
    }
}
