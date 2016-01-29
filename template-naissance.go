// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Ce template est utilisé pour mettre en forme les informations légales de retraite pour une date de naissance.
// Si vous disposez d'un relevé d'informations, utiliser le template "template-releve.go"  pour un bilan plus complet.
package main

import (
	"github.com/ObjectIsAdvantag/retraite/depart"
)

// The structure used by the template
type TemplateSimpleData struct {
	User  InfosUser
	Infos depart.InfosDepartEnRetraite
	SimuDecote  Simulation
	SimuSurcote Simulation
}

type Simulation struct {
	Trimestres          int    // .Infos.TrimestresTauxPlein - 20
	DeltaTrimestres		int 	// ex : 20
	EvolParTrimestre  string // le taux de décote par trimestre
	TauxPension         float32
}

var BilanDateDeNaissance = `
Votre bilan retraite simplifié suite à la réforme 2010 :

- vous devrez cotiser {{ .Infos.TrimestresTauxPlein }} trimestres pour toucher une retraite à taux plein
   - une retraite à taux plein correspond à une pension de l'ordre de 50% de vos 25 meilleurs années.

- vous partirez partir en retraite au plus tôt à {{ .Infos.AgeDépartMin.AgeEnAnneesMois }}, soit le {{ .Infos.DateDépartMin | timeToString }}
   - date à laquelle vous devrez avoir cotisé un minmum de {{ .SimuDecote.Trimestres }} trimestres
   - sans quoi vous devriez repousser votre demande de départ en retraite

   - principe de la décote :
      - votre pension est diminuée de {{ .SimuDecote.EvolParTrimestre }} points par trimestre manquant par rapport aux taux plein
      - ex: vous demandez à partir en retraite après le {{ .Infos.DateDépartMin | timeToString }} et avez cotisé {{ .SimuDecote.Trimestres }} trimestres,
            soient {{ .SimuDecote.DeltaTrimestres }} trimestres manquants par rapport au taux plein,
            votre pension serait alors de l'ordre de {{ .SimuDecote.TauxPension }}% de vos 25 meilleures années

   - principe de la surcote :
      - votre pension est augmentée de {{ .SimuSurcote.EvolParTrimestre }} points par trimestre supplémentaire cotisé
      - ex: vous demandez à partir en retraite après le {{ .Infos.DateDépartMin | timeToString }} et avez cotisé {{ .SimuSurcote.Trimestres }} trimestres,
            soient {{ .SimuSurcote.DeltaTrimestres }} trimestres supplémentaires par rapport au taux plein,
            votre pension serait alors de l'ordre de {{ .SimuSurcote.TauxPension }}% de vos 25 meilleures années

- à partir du {{ .Infos.DateTauxPleinAuto | timeToString }}, vous pourrez automatiquement bénéficier d'une retraite à taux plein,
   - et ce, quelque soit votre nombre de trimestres cotisés,
   - car vous aurez atteint l'âge légal de {{ .Infos.AgeTauxPleinAuto.AgeEnAnneesMois }}

- au delà du {{ .Infos.DateDépartExigible | timeToString }} si vous n'avez toujours pas demandé à partir en retraite,
   - votre employeur serait en droit de contraindre ce départ,
   - et vous auriez alors {{ .Infos.AgeDépartExigible.AgeEnAnneesMois }}

Pour en savoir plus :
- https://www.lassuranceretraite.fr/ par la sécurité sociale - CNAV
- http://www.info-retraite.fr/ par le GIP des organismes de retraite de base et complémentaires
- http://www.marel.fr/ : le simulateur M@rel
`
