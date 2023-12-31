package logic

import (
	"f1/internal/data"
	"f1/internal/entities"
	"fmt"
	"time"
)

func calculateCombination(name string, components []entities.Config) []entities.Config {

	resultCombi := []entities.Config{}

	for i := 0; i < len(components)-1; i++ {
		if components[i].IsActive == false {
			continue
		}
		componentI := components[i]
		for j := i + 1; j < len(components); j++ {
			if components[i].IsActive == false {
				continue
			}
			componentJ := components[j]
			combi := entities.Config{
				CombiCode:  []string{name + "-" + fmt.Sprint(i+1), name + "-" + fmt.Sprint(j+1)},
				Power:      componentI.Power + componentJ.Power,
				Aero:       componentI.Aero + componentJ.Aero,
				LigtWeight: componentI.LigtWeight + componentJ.LigtWeight,
				Grip:       componentI.Grip + componentJ.Grip,
				Cost:       componentI.Cost + componentJ.Cost,
				IsActive:   true,
			}
			resultCombi = append(resultCombi, combi)
		}
	}

	return resultCombi
}

func removeLesserCombi(combinations []entities.Config) []entities.Config {

	goodCombi := []entities.Config{}

	avgConfig := entities.Config{
		Power:      0,
		Aero:       0,
		LigtWeight: 0,
		Grip:       0,
		Cost:       0,
		IsActive:   false,
	}
	for i := 0; i < len(combinations); i++ {
		avgConfig.Cost += combinations[i].Cost
		avgConfig.Power += combinations[i].Power
		avgConfig.Aero += combinations[i].Aero
		avgConfig.LigtWeight += combinations[i].LigtWeight
		avgConfig.Grip += combinations[i].Grip
	}

	avgPointPerCost := (avgConfig.Power + avgConfig.Aero + avgConfig.LigtWeight + avgConfig.Grip) / float64(avgConfig.Cost)

	for i := 0; i < len(combinations); i++ {
		combiPointperCost := (combinations[i].Power + combinations[i].Aero + combinations[i].LigtWeight + combinations[i].Grip) / float64(combinations[i].Cost)
		if combiPointperCost >= (avgPointPerCost * data.Thershold) {
			goodCombi = append(goodCombi, combinations[i])
		}
	}

	return goodCombi
}

func GetIdealConfig(idelConfig entities.Config) (bestConfig *entities.Config, totalPoints float64, idealValueMatchCount int) {

	tm := time.Now()

	avgConfig := entities.Config{
		Power:      0,
		Aero:       0,
		LigtWeight: 0,
		Grip:       0,
	}

	powerCombination := calculateCombination("Power", data.PowerComponents)
	aeroCombination := calculateCombination("Aero", data.AeroComponents)
	ligtWeightCombination := calculateCombination("LigtWeight", data.LightWeightComponents)
	gripCombination := calculateCombination("Grip", data.GripComponents)
	toolCombination := calculateCombination("Tool", data.Tools)
	totalCombi := len(powerCombination) * len(aeroCombination) * len(ligtWeightCombination) * len(gripCombination) * len(toolCombination) * len(data.TeamPrincipals)

	powerCombination = removeLesserCombi(powerCombination)
	aeroCombination = removeLesserCombi(aeroCombination)
	ligtWeightCombination = removeLesserCombi(ligtWeightCombination)
	gripCombination = removeLesserCombi(gripCombination)
	toolCombination = removeLesserCombi(toolCombination)

	tCombiCount := len(powerCombination) * len(aeroCombination) * len(ligtWeightCombination) * len(gripCombination) * len(toolCombination) * len(data.TeamPrincipals)
	// tCombiCount := totalCombi
	// fmt.Println(totalCombi)
	// fmt.Println(tCombiCount)
	fmt.Printf("Reduced to %.2f %% combinations.\n", float64(tCombiCount)/float64(totalCombi)*100)

	count := 0

	idealValueMatchCount = 0
	idealValueMissDiff := 999999.99

	for _, pCombi := range powerCombination {
		for _, aCombi := range aeroCombination {
			for _, lCombi := range ligtWeightCombination {
				for _, gCombi := range gripCombination {
					for _, tCombi := range toolCombination {
						for tCount, teamPrincipal := range data.TeamPrincipals {

							idealValueMCount := 0
							idealValueMDiff := 0.0

							combiCode := pCombi.CombiCode
							combiCode = append(combiCode, aCombi.CombiCode...)
							combiCode = append(combiCode, lCombi.CombiCode...)
							combiCode = append(combiCode, gCombi.CombiCode...)
							combiCode = append(combiCode, tCombi.CombiCode...)

							teamPrincipal.CombiCode = []string{"TeamPrincipal-" + fmt.Sprint(tCount+1)}
							combiCode = append(combiCode, teamPrincipal.CombiCode...)
							count++
							currentCombination := entities.Config{
								CombiCode:  combiCode,
								Cost:       pCombi.Cost + aCombi.Cost + lCombi.Cost + gCombi.Cost + tCombi.Cost + teamPrincipal.Cost,
								Power:      (108.0 + pCombi.Power + aCombi.Power + lCombi.Power + gCombi.Power + tCombi.Power) * teamPrincipal.Power,
								Aero:       (162.0 + pCombi.Aero + aCombi.Aero + lCombi.Aero + gCombi.Aero + tCombi.Aero) * teamPrincipal.Aero,
								LigtWeight: (81.0 + pCombi.LigtWeight + aCombi.LigtWeight + lCombi.LigtWeight + gCombi.LigtWeight + tCombi.LigtWeight) * teamPrincipal.LigtWeight,
								Grip:       (216.0 + pCombi.Grip + aCombi.Grip + lCombi.Grip + gCombi.Grip + tCombi.Grip) * teamPrincipal.Grip,
							}

							score := currentCombination.Power + currentCombination.Aero + currentCombination.LigtWeight + currentCombination.Grip

							if currentCombination.Cost > idelConfig.Cost {
								continue
							}

							// Avg
							avgConfig.Power += currentCombination.Power
							avgConfig.Aero += currentCombination.Aero
							avgConfig.LigtWeight += currentCombination.LigtWeight
							avgConfig.Grip += currentCombination.Grip

							if currentCombination.Power >= idelConfig.Power {
								idealValueMCount++
							} else {
								idealValueMDiff += idelConfig.Power - currentCombination.Power
							}

							if currentCombination.Aero >= idelConfig.Aero {
								idealValueMCount++
							} else {
								idealValueMDiff += idelConfig.Aero - currentCombination.Aero
							}

							if currentCombination.LigtWeight >= idelConfig.LigtWeight {
								idealValueMCount++
							} else {
								idealValueMDiff += idelConfig.LigtWeight - currentCombination.LigtWeight
							}

							if currentCombination.Grip >= idelConfig.Grip {
								idealValueMCount++
							} else {
								idealValueMDiff += idelConfig.Grip - currentCombination.Grip
							}

							if idealValueMCount > idealValueMatchCount {
								idealValueMatchCount = idealValueMCount
								totalPoints = score
								bestConfig = &currentCombination
								idealValueMissDiff = idealValueMDiff
							} else if idealValueMCount == idealValueMatchCount && (idealValueMDiff < idealValueMissDiff || (idealValueMDiff == idealValueMissDiff && score > totalPoints)) {
								totalPoints = score
								bestConfig = &currentCombination
								idealValueMissDiff = idealValueMDiff
							}

							if count%5000000 == 0 {
								fmt.Printf("Count: %d of %d, Elapsed Time: %.2f seconds, Remaining Time: %.2f seconds\n", count/5000000, tCombiCount/5000000, time.Since(tm).Seconds(), (time.Since(tm) / time.Duration(count) * time.Duration(tCombiCount-count)).Seconds())
							}
						}
					}
				}
			}
		}
	}

	avgConfig.Power /= float64(count)
	avgConfig.Aero /= float64(count)
	avgConfig.LigtWeight /= float64(count)
	avgConfig.Grip /= float64(count)

	fmt.Println("Average Combination: Power: ", avgConfig.Power, " Aero: ", avgConfig.Aero, " LigtWeight: ", avgConfig.LigtWeight, " Grip: ", avgConfig.Grip)

	return

}
