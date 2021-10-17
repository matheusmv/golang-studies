# Variable & Type Declarations

 - 4 types of declaration:

     - variable
     - type
     - constant
     - function

## Variable

 - Default values:

     - Integer: 0
     - Float: 0
     - Complex Number: 0 Real and 0 Imaginary Part
     - Byte: 0
     - Rune: 0
     - String: ""
     - Bool: false
     - Array: Every array value to its default
     - Struct: Every field to its default
     - Map: nil
     - Channel: nil
     - Interface: nil
     - Slices: nil
     - Pointer: nil
     - Function: nil

 - Types of declarations:

        var name int = expression

        var name *int = &expression

        var var1, var2 = expression1, expression2

        var1, var2 := expression1, expression2

        name := expression

        name := &expression

 - Custom types:

        type fahrenheit int
	
        type celsius int
