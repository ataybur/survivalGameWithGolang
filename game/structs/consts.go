// consts.go
package structs

const (
	REGEX_1  = "There is a ([a-zA-Z ]+) at position ([0-9]+)"
	REGEX_2  = "([a-zA-Z ]+) attack is ([0-9]+)"
	REGEX_3  = "Resources are ([0-9]+) meters away"
	REGEX_4  = "([a-zA-Z ]+) has ([0-9]+) hp"
	REGEX_5  = "([a-zA-Z ]+) is Enemy"
	END_LINE = "\n"
	CONST_1  = "Hero started journey with %d HP!" + END_LINE
	CONST_2  = "Hero defeated %s with %d HP remaining" + END_LINE
	CONST_3  = "Survived" + END_LINE
	CONST_4  = "%s defeated Hero with %d HP remaining" + END_LINE
	CONST_5  = "Hero is Dead!! Last seen at position %d!!" + END_LINE
)

var REGEX_ARR = [5]string{REGEX_1, REGEX_2, REGEX_3, REGEX_4, REGEX_5}

var func_map = map[string]reg_func_interface{
	REGEX_1: reg1,
	REGEX_2: reg2,
	REGEX_3: reg3,
	REGEX_4: reg4,
	REGEX_5: reg5,
}

type reg_func_interface func([]string, string, *Context)
