# xform
--
    import "github.com/MJKWoolnough/gopherjs/xform"

Package xform provides some shortcut funcs for various form related activites

## Usage

#### func  InputButton

```go
func InputButton(id, name string) *dom.HTMLInputElement
```
InputButton creates a button input

#### func  InputCheckbox

```go
func InputCheckbox(id string, value bool) *dom.HTMLInputElement
```
InputCheckbox creates a checkbox input

#### func  InputColor

```go
func InputColor(id string) *dom.HTMLInputElement
```
InputColor creates a colour based input box, the workings of which are
implementation specific

#### func  InputDate

```go
func InputDate(id string) *dom.HTMLInputElement
```
InputDate create a date based input, the workings of which are implementation
specific

#### func  InputDateTime

```go
func InputDateTime(id string) *dom.HTMLInputElement
```
InputDateTime create a datetime based input, the workings of which are
implementation specific

#### func  InputDateTimeLocal

```go
func InputDateTimeLocal(id string) *dom.HTMLInputElement
```
InputDateTimeLocal create a local datetime based input, the workings of which
are implementation specific

#### func  InputEmail

```go
func InputEmail(id, value string) *dom.HTMLInputElement
```
InputEmail is a text box that validates as an email address

#### func  InputMonth

```go
func InputMonth(id string) *dom.HTMLInputElement
```
InputMonth creates a month based input box

#### func  InputNumber

```go
func InputNumber(id string, min, max, value float64) *dom.HTMLInputElement
```
InputNumber creates a text input that only allows number to be entered

#### func  InputPassword

```go
func InputPassword(id, value string) *dom.HTMLInputElement
```
InputPassword creates a password input

#### func  InputRadio

```go
func InputRadio(id, name string, value bool) *dom.HTMLInputElement
```
InputRadio create a radio button input

#### func  InputRange

```go
func InputRange(id string, min, max, step, value float64) *dom.HTMLInputElement
```
InputRange creates a sliding rule with which a number in the given range can be
selected

#### func  InputSizeable

```go
func InputSizeable(id, value string) *dom.HTMLSpanElement
```
InputSizeable returns a content-editable span that is style to look a text input
box

#### func  InputSubmit

```go
func InputSubmit(name string) *dom.HTMLInputElement
```
InputSubmit creates a submit input

#### func  InputText

```go
func InputText(id, value string) *dom.HTMLInputElement
```
InputText creates a text input box

#### func  InputTime

```go
func InputTime(id string) *dom.HTMLInputElement
```
InputTime creates a time based input box

#### func  InputURL

```go
func InputURL(id, value string) *dom.HTMLInputElement
```
InputURL is a text box that validates as a URL

#### func  InputUpload

```go
func InputUpload(id string) *dom.HTMLInputElement
```
InputUpload creates an upload input field

#### func  InputWeek

```go
func InputWeek(id string) *dom.HTMLInputElement
```
InputWeek creates a week based input box

#### func  Label

```go
func Label(label, forID string) *dom.HTMLLabelElement
```
Label create a form label

#### func  SelectBox

```go
func SelectBox(id string, values ...Option) *dom.HTMLSelectElement
```
SelectBox provides a select input, filled with the given options

#### func  TextArea

```go
func TextArea(id string, value string) *dom.HTMLTextAreaElement
```
TextArea provides a textarea input

#### type Option

```go
type Option struct {
	Label, Value string
	Selected     bool
}
```

Option is a structure to define a 'select's option.

#### type SizeableList

```go
type SizeableList struct {
	*dom.HTMLDivElement
}
```

SizeableList is a collection of InputSizable elements

#### func  InputSizeableList

```go
func InputSizeableList(values ...string) *SizeableList
```
InputSizeableList creates a list of InputSizeable elements, wrapped in a div

#### func (*SizeableList) Values

```go
func (s *SizeableList) Values() []string
```
Values returns the values of the enclose InputSizeable's
