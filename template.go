// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Le texte du bilan retraite
package main


const TexteBilan = `
Vos données personnelles :
- vous êtes né le {{ .User.Naissance }}
- à fin {{ .User.DateReleve }}, vous aviez acquis {{ .User.TrimestresAcquis }} trimestres

Votre bilan retraite suite à la réforme 2010 :

- vous partiriez en retraite au plus tôt le {{ TimeToString (.Infos.DateDépartMin) }}

********************************************************************************
Ce bilan vous est proposé gratuitement, sans publicité ni exploitation de vos données personnelles.
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