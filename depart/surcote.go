// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Regroupe les utilitaires concernant la surcote
package depart


// Retourne le taux de la surcote par trimestre supplémentaire
func SurcotePourTrimestreSupplementaire() float32 {
	return 0.625
}

// Retourne le taux de la pension pour un nombre de trimestres cotisés au delà du nombre requis pour un taux plein
func CalculerSurcotePourTrimestresSupplementaires(trimestresApresTauxPlein int) float32 {
	// Taux unique depuis 2008, indépendamment de la date de naissance
	return 50.0 + SurcotePourTrimestreSupplementaire() * float32(trimestresApresTauxPlein)
}


