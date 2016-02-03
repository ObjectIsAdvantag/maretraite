// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Regroupe les utilitaires concernant la décôte
// voir retraite_personnelle_pourcentage_minoration_applique_taux_plein_bar
package depart

import (
	"time"
	"fmt"
)

var ErrTrimestresManquantsHorsLimites = fmt.Errorf("Les trimestres manquants doivent se situer entre 0 et 20")

type TauxDecote struct {
	Nom        string      // nom commun du tableau de taux
	Depuis     time.Time   // date de naissance à partir de laquelle ces taux s'appliquent
	Jusque     time.Time   // date de naissance jusque laquelle ces taux s'appliquent
	Taux       [21]float32 // pourcentages de décotes en fonction des trimestres manquants
	Minoration float32     // le coefficient de minoration appliqué au taux plein (50 %) par trimestre manquant
	Diminution float32     // la diminution du taux par trimestre manquant
}

// Le type CalculDecote est utilisé pour retourné le résultat du calcul d'une pension de retraite en fonction d'un nombre de trimestres manquants
type CalculDecote struct {
	Decote        	  *TauxDecote // référence vers le taux de decote utilisé
	TauxPension       float32     // le taux de la pension suite à la décote, par exemple 45% par rapport à un taux plein à 50%
	DecoteTotale      float32     // la decote totale appliquée au vu du nombre de trimestres manquants
	TrimestresRetenus int         // le nombre de trimestres retenus pour le calcul
	Correction        bool        // Indique si un plafond a été atteint
}

// La décote est viagère, c'est-à-dire appliquée jusqu'au décès. Elle ne peut excéder 20 trimestres, contrairement à la surcote qui n'est pas plafonnée.
// Extrait de http://www.la-retraite-en-clair.fr/cid3196014/ce-faut-savoir-sur-decote.html
const DECOTE_MAX_TRIMESTRES = 20

var Default TauxDecote // référentiel post 1952

func init() {
	// Initialisation du référentiel de decôtes par défaut
	Default.Nom = "né après 1952"
	Default.Depuis, _ = StringToTime("01/01/1953")
	Default.Jusque, _ = StringToTime(fmt.Sprintf("01/01/%d", ANNEE_MAX))

	// Extrait de http://www.legislation.cnav.fr/Pages/bareme.aspx?Nom=retraite_personnelle_pourcentage_minoration_applique_taux_plein_bar#toc0
	Default.Minoration = 1.25
	Default.Diminution = 0.625

	Default.Taux[0] = 50
	Default.Taux[1] = 49.375
	Default.Taux[2] = 48.750
	Default.Taux[3] = 48.125
	Default.Taux[4] = 47.500
	Default.Taux[5] = 46.875
	Default.Taux[6] = 46.250
	Default.Taux[7] = 45.625
	Default.Taux[8] = 45.000
	Default.Taux[9] = 44.375
	Default.Taux[10] = 43.750
	Default.Taux[11] = 43.125
	Default.Taux[12] = 42.500
	Default.Taux[13] = 41.875
	Default.Taux[14] = 41.250
	Default.Taux[15] = 40.625
	Default.Taux[16] = 40
	Default.Taux[17] = 39.375
	Default.Taux[18] = 38.750
	Default.Taux[19] = 38.125
	Default.Taux[20] = 37.500
}

// Calcule la décôte selon le nombre de trimestres manquants et en fonction de l'année de naissance
// La date de naissance est au format JJ/MM/AAAA, de celle-ci dépend le coefficient de décôte
// Du nombre de trimestres manquants pour un taux plein dépend la valeur totale de la décote
// Retourne le taux de la pension, ainsi que la décote totale.
// Si une borne était dépassée, le nombre de trimestres retenu, et un bool précise si une correction sur le nombre de trimestres a été opérée.
// Si la date de naissance n'est pas précisée, le référentiel après 1952 s'applique
// Si le nombre de trimestres est inférieure à 1, une valeur de 0 trimestres est retournée
// Si le nombre de trimestres est supérieure à 20, une valeur pour 20 trimestres est retournée
func CalculerDécotePourTrimestresManquants(trimestres int, naissanceJJMMAAAA string) (CalculDecote, error) {
	naissance, err := parseDateNaissance(naissanceJJMMAAAA)
	if err != nil {
		return CalculDecote{}, fmt.Errorf("Impossible de calculer la décôte, la date de naissance %s n'est incorrecte, err: %v", naissanceJJMMAAAA, err)
	}

	return calculerDécotePourTrimestresManquantsInternal(trimestres, naissance), nil
}

func calculerDécotePourTrimestresManquantsInternal(trimestres int, naissance time.Time) (calcul CalculDecote) {


	// TODO Prendre en compte l'année de naissance,
	// A ce stade, on considère que l'année de naissance est toujours post 1952
	calcul.Decote = &Default

	if trimestres < 0 {
		calcul.Correction = true
		calcul.TrimestresRetenus = 0
		calcul.DecoteTotale = 0
		calcul.TauxPension = calcul.Decote.Taux[0]
		return
	}

	// Cas de la borne min
	if trimestres > DECOTE_MAX_TRIMESTRES {
		calcul.Correction = true
		calcul.TrimestresRetenus = DECOTE_MAX_TRIMESTRES
		calcul.DecoteTotale = calcul.Decote.Diminution * DECOTE_MAX_TRIMESTRES
		calcul.TauxPension = calcul.Decote.Taux[DECOTE_MAX_TRIMESTRES]
		return
	}

	calcul.Correction = false
	calcul.TrimestresRetenus = trimestres
	calcul.DecoteTotale = calcul.Decote.Diminution * float32(trimestres)
	calcul.TauxPension = calcul.Decote.Taux[trimestres]
	return
}
