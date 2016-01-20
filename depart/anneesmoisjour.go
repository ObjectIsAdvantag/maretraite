package depart


import (
	"time"
	"fmt"
)

// Le type AnneesMois représente une date sous la forme d'un nombre d'années, de mois et de jour
// Exemple : 62 ans et 7 mois
type AnneesMoisJours struct {
	Annees      int		`json:"annees,required"`		// nombre d'années
	Mois 		int		`json:"mois,omitempty"`			// nombre de mois
	Jours 		int 	`json:"jours,omitempty"`		// nombre de jours
}

func (amj AnneesMoisJours) EnAnnees() float32 {
	return float32(amj.Annees) + float32(amj.Mois)/12
}

func GetTime(annees int, mois int, jour int) (time.Time, error) {
	if annees < 0 || annees > 2100 || mois < 1 || mois > 12 || jour < 1 || jour > 32 {
		return time.Time{}, ErrDateFormatInvalide
	}

	return time.Date(annees, time.Month(mois), jour, 0, 0, 0, 0, time.UTC), nil
}




func CalculerAge(depuis time.Time, jusque time.Time) (AnneesMoisJours, error) {

	if jusque.Before(depuis) {
		return AnneesMoisJours{}, fmt.Errorf("la date de fin:%s est avant la date de début: %s", jusque, depuis)
	}

	//

	// Formule de calcul de l'age en année / mois
	// soient AAAA2/MM2/DD2 - AAAA1/MM1/DD1
	//
	// 1. Calculer si un ajustement est nécessaire en fonction de l'écart de jour
	// AAAA3/MM3/DD1 = AAAA2/MM2/DD2 - 1 mois
	// ECART_EN_JOUR = AAAA3/MM2/DD2 - AAAA3/MM3/DD1
	// IF ECART_EN_JOUR > 15 THEN ADJUST_MONTH = -1 ELSE ADJUST_MONTH = 0
	//
	// 2. Calculer l'écart en années / mois
	// IF MM2 < MM1 THEN RETURN AAAA2-AAAA1-1/12+MM2-MM1 + ADJUST_MONTH
	// IF MM2 > MM1 THEN RETURN AAAA2-AAAA1/MM2-MM1 + ADJUST_MONTH
	// IF MM2 == MM1 THEN IF DD2 > DD1 THEN RETURN AAAA2-AAAA1/MM2-MM1 ELSE RETURN AAAA2-AAAA1-1/12+MM2-MM1


	if jusque.Month() < depuis.Month(){
		return AnneesMoisJours{
			Annees: jusque.Year()-depuis.Year()-1,
			Mois: 12+int(jusque.Month())-int(depuis.Month()),
			}, nil
	}

	return AnneesMoisJours{
		Annees: jusque.Year()-depuis.Year(),
		Mois: int(jusque.Month())-int(depuis.Month()),
		}, nil
}