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

func getFilteredReport(filter string, i int, report []string) []string {
    var filteredReport []string

    for _, line := range report {
        if GetCharAt(line, i) == filter {
            filteredReport = append(filteredReport, line)
        }
    }

    return filteredReport
}

func getFilter(isOxygen bool, zeroCount int, oneCount int) string {
    if isOxygen {
        if zeroCount > oneCount {
            return "0"
        } else {
            return "1"
        }
    } else {
        if oneCount < zeroCount {
            return "1"
        } else {
            return "0"
        }
    }
}

func getOxygenOrCO2Rating(report []string, isOxygen bool) string {
    for i := 0; i < len(report[0]) && len(report) > 1; i++ {
        oneCount, zeroCount := 0, 0

        for j := 0; j < len(report); j++ {
            if GetCharAt(report[j], i) == "0" {
                zeroCount++
            } else {
                oneCount++
            }
        }

        filter := getFilter(isOxygen, zeroCount, oneCount)
        report = getFilteredReport(filter, i, report)
    }

    return report[0]
}

func main() {
    lines := ParseFile(3, false)

    // Part 1
    gamma, epilson := getGammaAndEpilson(lines)
    fmt.Println(GetIntFromBinary(gamma) * GetIntFromBinary(epilson))

    // Part 2
    oxygen := getOxygenOrCO2Rating(lines, true)
    co2 := getOxygenOrCO2Rating(lines, false)
    fmt.Println(GetIntFromBinary(oxygen) * GetIntFromBinary(co2))
}
