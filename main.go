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
        if (strings.Contains(record[1], "Verkaufserlös") || strings.Contains(record[1], "Release") ) && !strings.Contains(record[1], "Storno")  {
            fmt.Println("Kaufland: "+record[3]+" "+record[8]+" "+record[10])
            writer.Write([]string{""+record[38]+" "+record[39]+"","Warenheim", ""+record[0]+"", "Bestellnummer: "+record[3]+"", record[8]})
            //writer.Write([]string{"Kaufland", ""+record[0]+"", "KL-PROVISION-"+record[3]+"", "-"+record[12]})
         }  else if strings.Contains(record[1], "Auszahlung") || strings.Contains(record[1], "Payout"){
            fmt.Println("Auszahlung: "+record[3]+ " "+record[4]+" €")
            //writer.Write([]string{"Kaufland", ""+record[0]+"", "KL-AUSZAHLUNG", record[4]})
        }  else if strings.Contains(record[1], "Zusatzleistungen") {
            fmt.Println("Zusatzleistungen: "+record[3]+ " "+record[4]+" €")
            //writer.Write([]string{"Kaufland", ""+record[0]+"", "KL-SONSTIGES", record[4]})
        }  else if strings.Contains(record[1], "Storno") {
            fmt.Println("Storno: "+record[3])
            //writer.Write([]string{"Kaufland", ""+record[0]+"", "KL-STORNO-"+record[3]+"", record[12]})
        }  else if strings.Contains(record[1], "Lastschrift") {
            fmt.Println("Lastschrift : "+record[3]+ " "+record[4]+" €")
            //writer.Write([]string{"Kaufland", ""+record[0]+"", "KL-LASTSCHRIFT", record[4]})
        }
 
    }
    writer.Flush()

    fmt.Println("Ausgabe-CSV-Datei erfolgreich erstellt.")
}