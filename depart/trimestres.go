// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Référentiel des trimestres à cotiser en fonction des années de naissance
package depart

import (
	"errors"
	"time"
)

// Le référentiel des trimestres à cotiser selon l'année de naissace, tel qu'en vigueur en Janvier 2016
// voir http://www.la-retraite-en-clair.fr/cid3190611/a-quel-age-peut-partir-la-retraite.html
var refTrimestres ReferentielTrimestres

// Le type TrimestresParPeriode représente le nombre de trimestres à cotiser sur une période d'années de naissance
// Les bornes sont ANNEE_MIN et ANNEE_MAX
type TrimestresParPeriode struct {
	AnneeDebut int // date de naissance inférieure de la période (incluse)
	AnneeFin   int // date de naissance supérieure de la période (excluse)
	Trimestres int // nombre de trimestres à cotiser sur la période
}

var ErrAppelFonctionIncorrect = errors.New("paramètres d'invocation incorrects")

// Le type ReferentielTrimestres regroupe les âges de départ en retraite.
// Ce type permet d'abriter de futures données.
// Par défaut, le référentiel de 2010 est chargé.
type ReferentielTrimestres struct {
	nom        string                 // nom commun du référentiel
	depuis     time.Time              // date depuis laquelle le référentiel est en vigueur
	url        string                 // URL vers le référentiel ministériel
	trimestres []TrimestresParPeriode // stocke le nombre de trimstres à cotiser par périodes de naissance
}

func init() {
	// Chargement du référentiel par défaut
	refTrimestres.nom = "Réforme de 2010"
	refTrimestres.depuis, _ = time.Parse(JJMMAAADateFormat, "01/01/2010") // TODO : A vérifier
	refTrimestres.url = "http://www.la-retraite-en-clair.fr/cid3190611/a-quel-age-peut-partir-la-retraite.html"
	refTrimestres.trimestres = make([]TrimestresParPeriode, 13, 13)

	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{ANNEE_MIN, 1948, 160})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1949, 1950, 161})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1950, 1951, 162})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1951, 1952, 163})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1952, 1953, 164})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1953, 1955, 165})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1955, 1957, 166})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1958, 1961, 167})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1961, 1964, 168})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1964, 1967, 169})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1967, 1970, 170})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1970, 1973, 171})
	refTrimestres.trimestres = append(refTrimestres.trimestres, TrimestresParPeriode{1973, ANNEE_MAX, 172})

}

// Cette fonction retourne le nombre de trimstres à cotiser pour la date de naissance spécifiée
// Le référentiel de trimestres par défaut est interrogé.
func RechercherTrimestre(date time.Time) (int, error) {
	if date.IsZero() {
		return 0, ErrAppelFonctionIncorrect
	}

	for _, p := range refTrimestres.trimestres {
		if isInPeriode(date.Year(), p.AnneeDebut, p.AnneeFin) {
			return p.Trimestres, nil
		}
	}

	return 0, ErrPeriodeNonTrouvee
}

func isInPeriode(annee int, min int, max int) bool {
	return annee >= min && annee <= max
}
