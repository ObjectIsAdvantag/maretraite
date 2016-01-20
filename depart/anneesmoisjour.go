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

func AnneesMoisJourToTime(amj AnneesMoisJours) (time.Time, error) {
	if amj.Annees < 0 || amj.Annees > 2100 {
		return time.Time{}, ErrDateFormatInvalide
	}
	if amj.Mois < 1 || amj.Mois > 12 {
		return time.Time{}, ErrDateFormatInvalide
	}
	if amj.Jours < 1 || amj.Jours > 32 {
		return time.Time{}, ErrDateFormatInvalide
	}

	return time.Date(amj.Annees, time.Month(amj.Mois), amj.Jours, 0, 0, 0, 0, time.UTC), nil
}

func TimeToAnneesMoisJour(t time.Time) (AnneesMoisJours, error) {
	if t.IsZero() {
		return AnneesMoisJours{}, ErrDateFormatInvalide
	}

	return AnneesMoisJours{t.Year(), int(t.Month()), t.Day()}, nil
}


func CalculerAge(depuis time.Time, jusque time.Time) (AnneesMoisJours, error) {

	if jusque.Before(depuis) {
		return AnneesMoisJours{}, fmt.Errorf("la date de fin:%s se situe avant la date de début: %s", jusque, depuis)
	}

	// Formule de calcul de l'age en année / mois
	// soient AAAA2/MM2/DD2 - AAAA1/MM1/DD1
	//
	// 1. Se rendre à l'année cible et voir si on a dépassé
	// - si ce n'est pas le cas, ok
	// - si c'est le cas, revenir 1 an en arrière,
	// - mémoriser l'année cible, le bond réalisé en année, et le fait qu'on a dû ou non s'arrêter un an avant
	//
	//
	// 2. Se rendre sur le mois cible candidat et voir si on a dépassé
	// - si ce n'est pas le cas
	// - si c'est le cas, tenir qu'on doit opérer un changement de mois
	// - mémoriser le mois cible et calculer le bond réalisé en mois
	//
	// 3. Se rendre sur le mois cible, et calculer la différence en seconde entre les 2 dates
	// - convertir cette différence en jours

	// 1.
	tempDateCible, _ := AnneesMoisJourToTime( AnneesMoisJours{
		Annees:jusque.Year(),
		Mois:int(depuis.Month()),
		Jours:depuis.Day(),
	})

	anneeCible := jusque.Year()
	changementAnnee := 0
	if tempDateCible.After(jusque) {
		anneeCible--
		changementAnnee = 1
	}

	// 2.
	moisCible := int(jusque.Month())
	changementMois2 := 0
	changementAnnee2 := 0
	if depuis.Day() > jusque.Day() {
		moisCible--
		changementMois2 = 1
		if moisCible == 0 {
			moisCible = 12
			changementAnnee2 = 1
		}
	}
	nbMois := int(jusque.Month()) + 12 * changementAnnee - int(depuis.Month()) - changementMois2

	// 3.
	tempDateCible, _ = AnneesMoisJourToTime( AnneesMoisJours{
		Annees:jusque.Year()-changementAnnee2,
		Mois:moisCible,
		Jours:depuis.Day(),
	})

	deltaJours := jusque.Sub(tempDateCible).Minutes()/60/24

	return AnneesMoisJours{
		Annees: anneeCible - depuis.Year(),
		Mois: nbMois,
		Jours: int(deltaJours),
		}, nil
}