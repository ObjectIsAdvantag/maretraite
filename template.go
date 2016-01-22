// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Le texte du bilan retraite
package main

import (
	"github.com/ObjectIsAdvantag/retraite/depart"
)

// The structure used by the template
type TemplateData struct {
	User           InfosUser
	DepartMin      InfosDepart
	TauxPlein      InfosDepart
	DépartAuto     InfosDepart
	DépartExigible InfosDepart
	Surcote        CoeffRetraite
}

type InfosUser struct {
	Naissance          string
	DateReleve         int
	TrimestresCostisés int
}

type InfosDepart struct {
	DateDépart          string
	AgeDépart           depart.AnneesMoisJours
	TrimestresCotisés   int
	TrimestresManquants int
	Cote                CoeffRetraite
}

type CoeffRetraite struct {
	Type    string // "surcote ou décote ou pleine"
	Valeur  string // en pourcentage
	Période string // "trimestre" ou "année""
}

const TexteBilan = `
Vos données personnelles :
- vous êtes né le {{ .User.Naissance }}
- à fin {{ .User.DateReleve }}, vous aviez acquis {{ .User.TrimestresCostisés }} trimestres

Votre bilan retraite suite à la réforme 2010 :

- vous partirez en retraite au plus tôt le {{ .DepartMin.DateDépart }}
   - vous aurez alors {{ .DepartMin.AgeDépart.Annees }} ans et {{ .DepartMin.AgeDépart.Mois }} mois
   - avec un total de {{ .DepartMin.TrimestresCotisés }} trimestres acquis si vous cotisez sans interruption
   - il vous manquerait {{ .DepartMin.TrimestresManquants }} trimestres pour prétendre à un taux plein
   - aussi, votre pension de retraite subirait une {{ .DepartMin.Cote.Type }} de l'ordre de {{ .DepartMin.Cote.Valeur }}

- si vous cotisez sans interruption, vous pourriez opter pour une retraite à taux plein le {{ .TauxPlein.DateDépart }}
   - vous auriez alors {{ .TauxPlein.AgeDépart.Annees }} ans et {{ .TauxPlein.AgeDépart.Mois }}
   - avec un total de {{ .TauxPlein.TrimestresCotisés }} trimestres acquis
   - votre pension serait donc de l'ordre de 50% de vos 25 meilleures années

- si vous choisissez de poursuivre votre activité au delà du {{ .TauxPlein.DateDépart }},
   - vous accumelerez des trimestres supplémentaires,
   - qui provoqueront une surcôte de votre pension de l'ordre de {{ .Surcote.Valeur }} par {{ .Surcote.Période }}

- à partir du {{ .DépartAuto.DateDépart }}, vous pourrez automatiquement d'une retraite à taux plein
   - et ce quelque soit le nombre de trimestres cotisés, {{ .DépartAuto.TrimestresCotisés }} dans votre cas
   - car vous aurez atteint l'âge de {{ .DépartAuto.AgeDépart.Annees }} ans et {{ .DépartAuto.AgeDépart.Mois }}
   - votre pension sera dans ce cas de l'ordre de {{ .DépartAuto.Montant }} de vos 25 meilleures années

- au delà du {{ .DépartExigible.DateDépart }} si vous n'avez toujours pas demandé à partir en retraite,
   - votre employeur serait en droit de demander votre départ
   - vous auriez alors {{ .DépartExigible.AgeDépart.Annees }} ans et {{ .DépartExigible.AgeDépart.Mois }}


Que retenir :
   - vous pourrez demander à partir en retraite à partir du {{ .DepartMin.DateDépart }}
   - le montant de votre retraite sera fonction du nombre de trimestres cotisés,

Pour augmenter le montant de votre pension :
   - vous pourriez augmenter votre nombre de trimestres cotisés en rachetant des trimestres,
   - ou bénéficier du déclenchement automatique de votre taux plein à partir du {{ .DépartAuto.DateDépart }},
   - enfin, vous pouvez reprendre une activité après votre départ en retraite.


********************************************************************************
Ce bilan vous est proposé gratuitement, sans publicité ni exploitation de vos données personnelles.

Nous vous souhaitons de vivre de beaux moments professionnels jusqu'à votre départ en retraite
et espérons que ce bilan vous aura premis d'appréhender les enjeux principaux de votre retraite.

Si vous êtes un conseiller Retraite, n'hésitez pas à soumettre vos idées (améliorations, correctifs...) : https://github.com/ObjectIsAdvantag/retraite/issues/1

Si vous êtes développeur, n'héstiez pas à contribuer au projet opensource #golang sous license MIT : https://github.com/ObjectIsAdvantag/retraite

Copyright 2016, Stève Sfartz - @ObjectIsAdvantag
********************************************************************************

--
Pour en savoir plus :
- le site info-retraite :
- le simulateur M@rel :
- les decotes et surcotes :
`
