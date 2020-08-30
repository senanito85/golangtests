package main
import ( "fmt"; "strings" )

func typerCase(name string) {
	fmt.Println("My name is", name)
}

 func converter( module, author string) (s1, s2 string) {
        module = strings.Title(module)
        author = strings.ToUpper(author)

        return module, author
 }

func main() {

		typerCase("Sanan\n")
		typerCase("Jamie\n")
        
        module := "function basics"
        author := "nigel poulton"

        fmt.Println(converter(module, author))
}

