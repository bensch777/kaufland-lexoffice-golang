package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // Öffnen Sie die Quell-CSV-Datei
    inputFile, err := os.Open("input.csv")
    if err != nil {
        fmt.Println("Fehler beim Öffnen der Quell-CSV-Datei:", err)
        return
    }
    defer inputFile.Close()

    // Öffnen Sie die Ausgabe-CSV-Datei
    outputFile, err := os.Create("output.csv")
    if err != nil {
        fmt.Println("Fehler beim Erstellen der Ausgabe-CSV-Datei:", err)
        return
    }
    defer outputFile.Close()

    // Lesen Sie die Quell-CSV-Datei
    reader := csv.NewReader(inputFile)
    reader.Comma = ';'
    reader.LazyQuotes = true
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Fehler beim Lesen der Quell-CSV-Datei:", err)
        return
    }

    // Schreiben Sie die Ausgabe-CSV-Datei
    writer := csv.NewWriter(outputFile)
    for _, record := range records {
        fmt.Println(record[4])
        if record[2] == "Freigabe" {
            // Teilen Sie den Wert von sum_price_gross und fee_net durch 2 und schreiben Sie sie in separate Zeilen

            writer.Write(record)
        } else {
            writer.Write(record)
        }
    }
    writer.Flush()

    fmt.Println("Ausgabe-CSV-Datei erfolgreich erstellt.")
}