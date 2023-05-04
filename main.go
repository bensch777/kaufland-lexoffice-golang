package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
    "strconv"
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
            fmt.Println("Kaufland: "+record[2]+" "+priceConverter(record[6], record[8]))
            writer.Write([]string{""+record[3]+" "+record[4]+"","Kaufland", ""+record[0]+"", "Bestellnummer: "+record[2]+"", priceConverter(record[6], record[8])})
         } else if strings.Contains(record[5], "Auszahlung") || strings.Contains(record[5], "Payout"){
            fmt.Println("Auszahlung: "+record[0]+ " "+record[13]+" €")
            writer.Write([]string{"Kaufland", "Konto", ""+record[0]+"", "Auszahlung auf Bankkonto", record[13]})
        } 
    }
    writer.Flush()

    fmt.Println("Ausgabe-CSV-Datei erfolgreich erstellt.")
}


func priceConverter (price1Str string, price2Str string) string {
    // Ersetze das Komma durch einen Punkt und konvertiere in eine Dezimalzahl
    price1Str = strings.Replace(price1Str, ",", ".", 1)
    price1, err := strconv.ParseFloat(price1Str, 64)
    if err != nil {
        fmt.Println("Fehler beim Konvertieren von Preis 1:", err)
    }

    price2Str = strings.Replace(price2Str, ",", ".", 1)
    price2, err := strconv.ParseFloat(price2Str, 64)
    if err != nil {
        fmt.Println("Fehler beim Konvertieren von Preis 2:", err)
    }

    // Addiere die Preise
    totalPrice := price1 + price2

    // Konvertiere in das gewünschte Format
    totalPriceStr := strings.Replace(fmt.Sprintf("%.2f", totalPrice), ".", ",", 1)

    return totalPriceStr
}