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
	User            InfosUser
	Infos			depart.InfosDepartEnRetraite
	MinTrimestres	int								// .Infos.TrimestresTauxPlein - 20
	MinDecote		string							//
}



var BilanSimple = `
********************************************************************************
Ce bilan vous est proposé gratuitement, sans publicité ni exploitation des données renseignées.
Il a pour objectif de vous aider à appréhender les enjeux principaux de votre future retraite.
Nous vous souhaitons de vivre de beaux moments professionnels jusqu'à votre départ en retraite,

Si vous êtes un conseiller Retraite, n'hésitez pas à soumettre vos idées (améliorations, correctifs...) : https://github.com/ObjectIsAdvantag/retraite/issues/1
Si vous êtes développeur, n'hésitez pas à contribuer au projet https://github.com/ObjectIsAdvantag/retraite

Copyright 2016, Stève Sfartz - @ObjectIsAdvantag - License MIT
********************************************************************************


Vos données personnelles :
- vous êtes né le {{ .User.Naissance }}
- vous ne disposez pas de relevé de carrière
   - des relevés de situation vous sont transmis automatiquement tous les 5 ans à partir de vos 40 ans
   - vous pouvez aussi générer votre relevé après avoir créé un compte sur le site XXXXXX
   - à partir de 55 ans, il est aussi possible de prendre rendez-vous lien XXXX


Votre bilan retraite suite à la réforme 2010 :

- vous devez cotiser {{ .Infos.TrimestresTauxPlein }} trimestres pour une retraite à taux plein,
   - votre pension à taux plein correspondant à 50% de la rémunération de vos 25 meilleurs années.

- vous partiriez partir en retraite au plus tôt à {{ .Infos.AgeDépartMin.AgeEnAnneesMois }},
   - si vous avez cotisé un minimum de {{ .MinTrimestres }} trimestres
   - avec une décôte de {{ .MinDecote }} par trimestre manquant par rapport au taux plein

   - ainsi, par exemple pour un départ avec 20 trimestres manquants,
   - vous obtiendriez une pension de l'ordre de XX % de vos 25 meilleures années

- à partir du {{ .Infos.DateTauxPleinAuto | timeToString }}, vous pourrez automatiquement bénéficier d'une retraite à taux plein
   - car vous aurez atteint l'âge légal de {{ .Infos.AgeTauxPleinAuto.AgeEnAnneesMois }}

- au delà du {{ .Infos.DateDépartExigible | timeToString }} si vous n'avez toujours pas demandé à partir en retraite,
   - votre employeur serait en droit de demander votre départ
   - et vous auriez alors {{ .Infos.AgeDépartExigible.AgeEnAnneesMois }}


Pour en savoir plus :
- le site info-retraite :
- le simulateur M@rel :
- les decotes et surcotes :
`
