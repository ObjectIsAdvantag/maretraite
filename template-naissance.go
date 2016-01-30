// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Ce template est utilisé pour mettre en forme les informations légales de retraite pour une date de naissance.
// Si vous disposez d'un relevé d'informations, utiliser le template "template-releve.go"  pour un bilan plus complet.
package main

import (
	"github.com/ObjectIsAdvantag/maretraite/depart"
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
Votre bilan retraite simplifié suite à la réforme de 2010 :

- vous devez cotiser {{ .Infos.TrimestresTauxPlein }} trimestres pour toucher une retraite à taux plein
   - une retraite à taux plein correspond à une pension de l'ordre de 50% de vos 25 meilleurs années.

- vous pourrez partir en retraite au plus tôt à {{ .Infos.AgeDépartMin.AgeEnAnneesMois }} (après le {{ .Infos.DateDépartMin | timeToString }})
   - et si vous avez cotisé au moins {{ .SimuDecote.Trimestres }} trimestres
   - sans quoi vous devrez repousser votre demande de départ en retraite

   - principe de la décote :
      - votre pension est diminuée de {{ .SimuDecote.EvolParTrimestre }} points par trimestre manquant par rapport au taux plein
      - ainsi: si vous demandiez à partir en retraite après {{ .Infos.AgeDépartMin.AgeEnAnneesMois }} et aviez cotisé {{ .SimuDecote.Trimestres }} trimestres,
            soient {{ .SimuDecote.DeltaTrimestres }} trimestres manquants par rapport au taux plein ({{ .Infos.TrimestresTauxPlein }} trimestres),
            votre pension serait alors de l'ordre de {{ .SimuDecote.TauxPension }}% de vos 25 meilleures années

   - principe de la surcote :
      - votre pension est augmentée de {{ .SimuSurcote.EvolParTrimestre }} points par trimestre supplémentaire cotisé
      - ainsi: si vous demandiez à partir en retraite après {{ .Infos.AgeDépartMin.AgeEnAnneesMois }} et aviez cotisé {{ .SimuSurcote.Trimestres }} trimestres,
            soient {{ .SimuSurcote.DeltaTrimestres }} trimestres supplémentaires par rapport au taux plein ({{ .Infos.TrimestresTauxPlein }} trimestres),
            votre pension serait alors de l'ordre de {{ .SimuSurcote.TauxPension }}% de vos 25 meilleures années

- à partir de {{ .Infos.AgeTauxPleinAuto.AgeEnAnneesMois }} (après le {{ .Infos.DateTauxPleinAuto | timeToString }}), vous pourrez automatiquement bénéficier d'une retraite à taux plein,
   - et ce, quelque soit votre nombre de trimestres cotisés,
   - car vous aurez atteint l'âge légal pour votre année de naissance.

- au delà de {{ .Infos.AgeDépartExigible.AgeEnAnneesMois }} (après le {{ .Infos.DateDépartExigible | timeToString }}), si vous n'êtes toujours pas en retraite,
   - votre employeur sera en droit de vous obliger à partir en retraite.

Pour en savoir plus :
- https://www.lassuranceretraite.fr/ par la sécurité sociale - CNAV
- http://www.info-retraite.fr/ par le GIP des organismes de retraite de base et complémentaires
- http://www.marel.fr/ : le simulateur M@rel
`
