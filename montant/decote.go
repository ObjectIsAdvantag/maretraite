// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Regroupe les utilitaires concernant la décôte
// voir retraite_personnelle_pourcentage_minoration_applique_taux_plein_bar
package montant

import (
	"time"

	"fmt"
	"github.com/ObjectIsAdvantag/retraite/depart"
)

var ErrTrimestresManquantsHorsLimites = fmt.Errorf("Les trimestres manquants doivent se situer entre 0 et 20")

type TauxDecote struct {
	Nom    string      // nom commun du tableau de taux
	Depuis time.Time   // date de naissance à partir de laquelle ces taux s'appliquent
	Jusque time.Time   // date de naissance jusque laquelle ces taux s'appliquent
	Taux   [21]float32 // pourcentages de décotes en fonction des trimestres manquants
	Minoration   float32 // le coefficient de minoration appliqué au taux plein (50 %) par trimestre manquant
	Diminution   float32 // la diminution du taux par trimestre manquant

}

var Default TauxDecote // référentiel post 1952

func init() {
	// Initialisation du référentiel de decôtes par défaut
	Default.Nom = "né après 1952"
	Default.Depuis, _ = depart.StringToTime("01/01/1953")
	Default.Jusque, _ = depart.StringToTime(fmt.Sprintf("01/01/%d", depart.ANNEE_MAX))

	// Extrait de http://www.legislation.cnav.fr/Pages/bareme.aspx?Nom=retraite_personnelle_pourcentage_minoration_applique_taux_plein_bar#toc0
	Default.Minoration = 1.25
	Default.Diminution = 0.625

	Default.Taux[0] = 50
	Default.Taux[1] = 49.3125
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
// La date de naissance est au format JJ/MM/AAAA, elle permet de sélectionner le bon coefficient de décôte.
// Si la date de naissance n'est pas précisée, le référentiel après 1952 s'applique
func DécotePourTrimestresManquants(trimestres int, dateNaissance string) (float32, error) {
	// TODO Prendre en compte l'année de naissance
	ref := Default

	if trimestres < 0 || trimestres > 20 {
		return 0, ErrTrimestresManquantsHorsLimites
	}

	return ref.Taux[trimestres], nil
}


