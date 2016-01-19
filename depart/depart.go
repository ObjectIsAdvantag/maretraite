// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Age légal de départ à la retraite (as of January 2016)
//
// Ce service vous permet de connaitre votre âge légal de départ à la retraite et vous précise :
// - votre âge légal de départ à la retraite
// - le point de départ possible de votre retraite
// - la durée de cotisation pour obtenir une retraite à taux plein
// - l'âge et la date auxquels vous obtiendrez automatiquement une retraite à taux plein
package depart

import (
	"time"
)

// La structure DepartLegal regroupe les données légales de départ à la retraite.
// Cette structure est calculée à partir d'une date de naissance, voir la fonction CalculerDepartLegal
type InfosDepartLegal struct {
	TrimestresRequis	int					`json:"nombreDeTrimestresRequis"`	// nombre de trimestres afin de disposer de sa retraite à taux plein
	AgeDepartMin   		AgeEnAnneesMois		`json:"ageDepartMinimum"`// âge à partir duquel il est possible de percevoir une retraite, mais elle ne sera pas à taux plein si le nombre de trimestres cotisés n'est pas celui requis
	//DateDepartMin		string				`json:"dateDepartMinimum"`// date à partir de laquelle il est possible de percevoir une retraite
	AgeTauxPleinAuto	AgeEnAnneesMois		`json:"ageTauxPleinAutomatique"`// âge à partir duquel il est possible de percevoir une retraite à taux plein même si le nombre de trimestres cotisés n'est pas suffisant
	//DateDepartAuto		string				`json:"dateDepartAutomatique"`// date à partir de laquelle il est possible de percevoir une retraite à taux plein même si le nombre de trimestres cotisés n'est pas suffisant
}

type InfosDepart struct {
	Age 				AgeEnAnneesMois    	// âge au moment du départ
	Date 				time.Time		   	// date en JJ/MM/AAAA
	Trimestres		 	int					// nombre de trimestres total  entre la date de début d'activité et le moment du départ
	TrimestreTravailles	int					// nombre de trimestres travaillés entre la date de début d'activité et le moment du départ
}

const ANNEE_MIN = 1900
const ANNEE_MAX = 2100

// Calcul les informations de départ légale à la retraite, à partir d'une date de naissance au format JJ/MM/AAAA
// Il est possible de spécifier un nombre de trimestres non cotisés, qui viendront décaler d'autant les dates et âges
func CalculerDepartLegal(input string, nonCotises int) (*InfosDepartLegal, error) {
	dateNaissance, err := parseDate(input)
	if err != nil {
		return nil, err
	}

	var res InfosDepartLegal
	res.TrimestresRequis, err = RechercherTrimestre(dateNaissance)
	if err != nil {
		return nil, err
	}

	al, err := RechercherAgeLegal(dateNaissance)
	if err != nil {
		return nil, err
	}

	res.AgeDepartMin = al.AgeDepartMin

	res.AgeTauxPleinAuto = al.AgeTauxPleinAuto

	return &res, nil
}

