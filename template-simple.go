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
	Simu  Simulation
}

type Simulation struct {
	Trimestres          int    // .Infos.TrimestresTauxPlein - 20
	DeltaTrimestres		int 	// ex : 20
	DecoteParTrimestre  string // le taux de décote par trimestre
	TrimestresCotisés   int
	TrimestresManquants int
	TauxPension         float32
}

var BilanSimple = `
********************************************************************************
Ce bilan vous est proposé gratuitement, sans publicité ni exploitation des données renseignées.
Il a pour objectif de vous aider à appréhender les enjeux principaux de votre future pension.
Nous vous souhaitons de vivre de beaux moments professionnels jusqu'à votre départ en retraite.

Si vous êtes un conseiller Retraite, n'hésitez pas à soumettre vos idées (améliorations, correctifs...) : https://github.com/ObjectIsAdvantag/retraite/issues/1
Si vous êtes développeur, n'hésitez pas à contribuer au projet https://github.com/ObjectIsAdvantag/retraite

Copyright 2016, Stève Sfartz - @ObjectIsAdvantag - License MIT
********************************************************************************


Vos données personnelles :
- vous êtes né le {{ .User.Naissance }}
- vous ne disposez pas de relevé de carrière.


Votre bilan retraite simplifié suite à la réforme 2010 :

- vous devez cotiser {{ .Infos.TrimestresTauxPlein }} trimestres pour une retraite à taux plein
   - une retraite à taux plein correspondant à une pension de l'ordre de 50% de vos 25 meilleurs années.

- vous partirez partir en retraite au plus tôt à {{ .Infos.AgeDépartMin.AgeEnAnneesMois }}, soit le {{ .Infos.DateDépartMin | timeToString }}
   - vous devrez alors avoir cotisé un minmum de {{ .Simu.Trimestres }} trimestres soient {{ .Simu.DeltaTrimestres }} trimestres manquants par rapport à votre taux plein,
   - sans quoi vous devriez repousser votre demande de départ en retraite jusqu'à atteindre ce minimum de {{ .Simu.Trimestres }} trimestres

   - décote : votre pension est diminuée de {{ .Simu.DecoteParTrimestre }}% par trimestre manquant par rapport aux taux plein
   - ex: vous demandez à partir en retraite après le {{ .Infos.DateDépartMin | timeToString }} et avez cotisé {{ .Simu.TrimestresCotisés }} trimestres,
       - soient {{ .Simu.TrimestresManquants }} trimestres manquants par rapport à un taux plein,
       - votre pension serait alors de l'ordre de {{ .Simu.TauxPension }}% de vos 25 meilleures années

- à partir du {{ .Infos.DateTauxPleinAuto | timeToString }}, vous pourrez automatiquement bénéficier d'une retraite à taux plein,
   - et ce, quelque soit votre nombre de trimestres cotisés,
   - car vous aurez atteint l'âge légal de {{ .Infos.AgeTauxPleinAuto.AgeEnAnneesMois }}

- au delà du {{ .Infos.DateDépartExigible | timeToString }} si vous n'avez toujours pas demandé à partir en retraite,
   - votre employeur serait en droit de contraindre ce départ,
   - et vous auriez alors {{ .Infos.AgeDépartExigible.AgeEnAnneesMois }}

- pour un bilan complet, relancer le calculateur avec les données de votre relevé de carrière
	- des relevés de situation vous sont transmis par courrier à partir de vos 40 ans, et tous les 5 ans,
   	- ces relevés font un point sur vos trimestres cotisés,
    - vous pouvez aussi télécharger votre relevé après avoir créé votre espace depuis https://www.lassuranceretraite.fr/
    - à partir de 55 ans, il est aussi possible de prendre rendez-vous avec votre caisse.

Pour en savoir plus :
- https://www.lassuranceretraite.fr/ par la sécurité sociale - CNAV
- http://www.info-retraite.fr/ par le GIP des organismes de retraite de base et complémentaires
- http://www.marel.fr/ : le simulateur M@rel
`
