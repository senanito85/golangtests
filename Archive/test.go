
    a := 10
    Switch a  {
    case 10:
    fmt.Println("10")
    case 9:
    fmt.Println("9")
    case 1:
    fmt.Println("1")
    case default:
    fmt.Println("none of these")

}
    a := 10
    Switch a == 10 {
    case true:
    fmt.Println("the value is negative")
    case false:
    fmt.Println("the value is single digit")
    }


    package main

    import (
        "fmt"
        "os"
    )

    type point struct {
        x, y int
    }

    func main() {

        p := point{1, 2}
        fmt.Printf("%v\n", p)

        fmt.Printf("%+v\n", p)

        fmt.Printf("%#v\n", p)

        fmt.Printf("%T\n", p)

        fmt.Printf("%t\n", true)

        fmt.Printf("%d\n", 123)

        fmt.Printf("%b\n", 14)

        fmt.Printf("%c\n", 33)

        fmt.Printf("%x\n", 456)

        fmt.Printf("%f\n", 78.9)

        fmt.Printf("%e\n", 123400000.0)
        fmt.Printf("%E\n", 123400000.0)

        fmt.Printf("%s\n", "\"string\"")

        fmt.Printf("%q\n", "\"string\"")

        fmt.Printf("%x\n", "hex this")

        fmt.Printf("%p\n", &p)

        fmt.Printf("|%6d|%6d|\n", 12, 345)

        fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

        fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

        fmt.Printf("|%6s|%6s|\n", "foo", "b")

        fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

        s := fmt.Sprintf("a %s", "string")
        fmt.Println(s)

        fmt.Fprintf(os.Stderr, "an %s\n", "error")
    }
