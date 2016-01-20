// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Regroupe les utilitaires concernant la décôte
// voir retraite_personnelle_pourcentage_minoration_applique_taux_plein_bar
package montant

import (
	"time"

	"github.com/ObjectIsAdvantag/retraite/depart"
	"fmt"
)

var ErrTrimestresManquantsHorsLimites = fmt.Errorf("Les trimestres manquants doivent se situer entre 0 et 20")

type TauxDecote struct {
	Nom    string    // nom commun du tableau de taux
	Depuis time.Time // date de naissance à partir de laquelle ces taux s'appliquent
	Jusque time.Time // date de naissance jusque laquelle ces taux s'appliquent
	Taux   [21]float32 // pourcentages de décotes en fonction des trimestres manquants
}

var defaultRef TauxDecote // référentiel post 1952

func init() {
	// Initialisation du référentiel de decôtes par défaut
	defaultRef.Nom = "né après 1952"
	defaultRef.Depuis = depart.StringToTime("01/01/1953")
	defaultRef.Jusque = time.Time{}

	// Extrait de http://www.legislation.cnav.fr/Pages/bareme.aspx?Nom=retraite_personnelle_pourcentage_minoration_applique_taux_plein_bar#toc0
	defaultRef.Taux[0] = 50
	defaultRef.Taux[1] = 49,3125
	defaultRef.Taux[2] = 48,750
	defaultRef.Taux[3] = 48,125
	defaultRef.Taux[4] = 47,500
	defaultRef.Taux[5] = 46,875
	defaultRef.Taux[6] = 46,250
	defaultRef.Taux[7] = 45,625
	defaultRef.Taux[8] = 45,000
	defaultRef.Taux[9] = 44,375
	defaultRef.Taux[10] = 43,750
	defaultRef.Taux[11] = 43,125
	defaultRef.Taux[12] = 42,500
	defaultRef.Taux[13] = 41,875
	defaultRef.Taux[14] = 41,250
	defaultRef.Taux[15] = 40,625
	defaultRef.Taux[16] = 40,000
	defaultRef.Taux[17] = 39,375
	defaultRef.Taux[18] = 38,750
	defaultRef.Taux[19] = 38,125
	defaultRef.Taux[20] = 37,500
}

// Calcule la décôte selon le nombre de trimestres manquants et en fonction de l'année de naissance
// La date de naissance est au format JJ/MM/AAAA, elle permet de sélectionner le bon coefficient de décôte.
// Si la date de naissance n'est pas précisée, le référentiel après 1952 s'applique

func DécotePourTrimestresManquants(trimestres int, dateNaissance string) (float32, error) {
	// TODO Prendre en compte l'année de naissance
	ref := defaultRef

	if trimestres < 0 || trimestres > 20 {
		return 0, ErrTrimestresManquantsHorsLimites
	}

	return ref[trimestres], nil
}