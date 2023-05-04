package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
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
    writer.Write([]string{"Auftraggeber", "Empfänger", "Buchungsdatum", "Verwendungszweck", "Betrag"})
    for _, record := range records {
        if (strings.Contains(record[5], "Verkaufserlös") || strings.Contains(record[5], "Release") ) && !strings.Contains(record[5], "Storno")  {
            fmt.Println("Kaufland: "+record[2]+" "+record[6])
            writer.Write([]string{""+record[3]+" "+record[4]+"","Kaufland", ""+record[0]+"", "Bestellnummer: "+record[2]+"", record[6]})
            
         } 
 
    }
    writer.Flush()

    fmt.Println("Ausgabe-CSV-Datei erfolgreich erstellt.")
}