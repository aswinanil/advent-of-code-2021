package main

import (
    "fmt"
)

func getGammaAndEpilson(report []string) (string, string) {
    var gamma, epilson string

    for i := 0; i < len(report[0]); i++ {
        oneCount, zeroCount := 0, 0

        for j := 0; j < len(report); j++ {
            if GetCharAt(report[j], i) == "0" {
                zeroCount++
            } else {
                oneCount++
            }
        }

        if zeroCount > oneCount {
            gamma += "0"
            epilson += "1"
        } else if zeroCount < oneCount {
            gamma += "1"
            epilson += "0"
        } else {
            fmt.Println("Error: 0s & 1s occurs same amount of times!")
        }
    }

    return gamma, epilson
}

func main() {
    lines := ParseFile(3, false)

    // Part 1
    gamma, epilson := getGammaAndEpilson(lines)
    fmt.Println(GetIntFromBinary(gamma) * GetIntFromBinary(epilson))
}
