package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
    "regexp"
)

func main() {
    payloadsFile := flag.String("payloads", "", "Path to the file containing payloads")
    parametersFile := flag.String("parameters", "", "Path to the file containing parameters")
    flag.Parse()

    if *payloadsFile == "" || *parametersFile == "" {
        fmt.Println("Usage: replfuzz -payloads payloads.txt -parameters parameters.txt")
        os.Exit(1)
    }

    payloads, err := readLines(*payloadsFile)
    if err != nil {
        fmt.Printf("Error reading payloads file: %v\n", err)
        os.Exit(1)
    }
    parameters, err := readLines(*parametersFile)
    if err != nil {
        fmt.Printf("Error reading parameters file: %v\n", err)
        os.Exit(1)
    }

    paramMap := make(map[string]string)
    for _, line := range parameters {
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            paramMap[parts[0]] = parts[1]
        }
    }

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        url := scanner.Text()
        for _, payload := range payloads {
            modifiedURL := url
            for key := range paramMap {
                modifiedURL = replaceParameterValue(modifiedURL, key, payload)
            }
            if containsAnyParam(modifiedURL, paramMap) {
                fmt.Println(modifiedURL)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading from standard input: %v\n", err)
        os.Exit(1)
    }
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return lines, nil
}

func replaceParameterValue(url, param, payload string) string {
    // Regular expression to match the parameter and its value
    re := regexp.MustCompile(fmt.Sprintf(`([?&]%s=)[^&]*`, regexp.QuoteMeta(param)))
    return re.ReplaceAllString(url, fmt.Sprintf("$1%s", payload))
}

func containsAnyParam(url string, params map[string]string) bool {
    for key := range params {
        if strings.Contains(url, key+"=") {
            return true
        }
    }
    return false
}
