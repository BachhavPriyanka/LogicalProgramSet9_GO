// structure to M
package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Friend struct {
	Name                       string
	MobNumber, AlternateMobNum int
	Address                    string
	City                       string
}

func main() {
	Details := []Friend{
		{"Rizwaan", 987878765, 999983663, "Indira", "Hyderbaad"},
		{"Diksha", 9887388743, 999943562, "Society", "Mumbai"},
		{"Mac", 78664799, 34768889, "laxmiRoad", "Pune"},
		{"Ken", 987656649, 87636372, "College-Road", "Nagpur"},
		{"Jenny", 999556789, 9998382722, "Airforce", "Pune"},
		{"Diksha", 964336777, 34768889, "camp", "Mumbai"},
		{"Tom", 789089995, 938829378, "fish-Maeket", "Kokan"},
		{"henrry", 978978900, 34768889, "camp2", "lonawalaa"},
		{"Albert", 898080786, 34768889, "Badi-Chopar", "Jaipur"},
		{"Albert", 898080786, 34768889, "mansorovar", "Jaipur"},
	}

	file, err := os.Create("frienDetails.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := csv.NewWriter(file)

	var data [][]string
	for _, record := range Details {
		row := []string{record.Name, strconv.Itoa(record.MobNumber), strconv.Itoa(record.AlternateMobNum), record.Address, record.City}
		data = append(data, row)
	}
	write.WriteAll(data)

}
