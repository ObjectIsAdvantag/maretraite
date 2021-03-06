// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Point d'entrée du calculateur de retraite pour des salariés du privé.
// Ce programme interactif se lance dans une console.
// Le programme interroge l'utilisateur et retourne ses données de départ en retraite sous forme textuelle.
package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ObjectIsAdvantag/maretraite/depart"
	log "github.com/Sirupsen/logrus"
)

type InfosUser struct {
	Naissance         string
	AnnéeRelevé        int
	TrimestresCotisés int
}

const version = "v0.1"

func (infos InfosUser) sansRelevé() bool {
	return (infos.TrimestresCotisés == 0) && (infos.AnnéeRelevé == 0)
}

// Point d'entrée du programme
func main() {
	// Interroge l'utilisateur
	log.Debugln("Demande des informations : start")
	userData, err := interrogerUtilisateur()
	if err != nil {
		log.Infof("Données incorrectes, le bilan ne peut être généré")
		fmt.Println("Données incorrectes, le bilan ne peut être généré.\nEssayez à nouveau...")
		os.Exit(-1)
	}
	log.Debugln("Demande des informations : ok! ", userData.Naissance, userData.AnnéeRelevé, userData.TrimestresCotisés)

	// Calcule les données de départ en retraite
	log.Debugln("CalculerDépartLégal : start")
	infosLegales, _ := depart.CalculerInfosLégales(userData.Naissance)
	log.Debugln("CalculerDépartLégal : ok! ", infosLegales)

	// Génère le bilan
	log.Debugln("Génération du bilan : start")
	if userData.sansRelevé() {
		genererBilanSimple(userData, infosLegales)
	} else {
		genererBilanComplet(userData, infosLegales)
	}
	log.Debugln("Génération du bilan : ok!")
}

// Demande à l'utilisateur sa date de naissance à minima, ainsi que qq informations de son relevé de carrière
// Une erreur est retournée si aucune date de naissance n'est proposée
func interrogerUtilisateur() (InfosUser, error) {
	// Banner
	fmt.Printf(`
********************************************************************************
Ce bilan vous est proposé gratuitement, sans publicité ni exploitation commerciale de vos données.
Il a pour objectif de vous aider à appréhender les enjeux principaux de votre future pension.
Nous vous souhaitons de vivre de beaux moments professionnels jusqu'à votre départ en retraite.

Si vous êtes un conseiller Retraite, n'hésitez pas à soumettre vos idées (améliorations, correctifs...) :
https://github.com/ObjectIsAdvantag/retraite/issues/1

Si vous êtes développeur, n'hésitez pas à contribuer au projet :
https://github.com/ObjectIsAdvantag/retraite

Copyright 2016, Stève Sfartz - @ObjectIsAdvantag - License MIT - %s
********************************************************************************

`, version)

	// Date de naissance
	var naissance string
	essai, essaiMax := 1, 3
	for {
		fmt.Printf("Votre date de naissance (JJ/MM/YYYY): ")
		fmt.Scanf("%s\n", &naissance)
		_, err := depart.StringToTime(naissance)
		if err == nil {
			log.Debugf("Date saisie : %v", naissance)
			break
		}
		if essai >= essaiMax {
			return InfosUser{}, fmt.Errorf("Impossible de récupérer la date de naissance")
		}

		essai++
		fmt.Printf("Saisie incorrecte, nouvel essai (%d/%d)\n", essai, essaiMax)
	}

	// TODO Relevé de carrière
	//return InfosUser{Naissance:"24/12/1971", TrimestresCotisés:87, DateReleve:2014}, nil

	return InfosUser{Naissance: naissance, TrimestresCotisés: 0, AnnéeRelevé: 0}, nil
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	// [PENDING]
	log.SetLevel(log.WarnLevel)
}

func genererBilanSimple(userData InfosUser, infos depart.InfosDepartEnRetraite) {
	log.Debugln("genererBilanSimple : start")

	t, _ := template.New("BilanSimple").Funcs(template.FuncMap{
		"timeToString": depart.TimeToString,
	}).Parse(BilanDateDeNaissance)

	trimestresDecote := 16 // au maximum pension.DECOTE_MAX_TRIMESTRES = 20 trimestres
	decote, _ := depart.CalculerDécotePourTrimestresManquants(trimestresDecote, userData.Naissance)
	simuDecote := Simulation {
		// la décote est limitée à 20 trimestres
		Trimestres:infos.TrimestresTauxPlein - trimestresDecote,
		DeltaTrimestres:trimestresDecote,
		EvolParTrimestre:fmt.Sprintf("%.3f", decote.Decote.Diminution),
		TauxPension:decote.TauxPension,
	}

	deltaTrimestres := 8
	simuSurcote := Simulation {
		// la surcote n'est pas limitée
		Trimestres:infos.TrimestresTauxPlein + deltaTrimestres,
		DeltaTrimestres:deltaTrimestres,
		EvolParTrimestre:fmt.Sprintf("%.3f", depart.SurcotePourTrimestreSupplementaire()),
		TauxPension:depart.CalculerSurcotePourTrimestresSupplementaires(deltaTrimestres),
	}

	data := TemplateSimpleData{
		User:userData,
		Infos:infos,
		SimuDecote:simuDecote,
		SimuSurcote:simuSurcote,
	}

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalf("Le bilan n'a pas pu être généré, err: ", err)
	}
	log.Debugln("genererBilanSimple : ok !")
}

func genererBilanComplet(userData InfosUser, infos depart.InfosDepartEnRetraite) {
	log.Debugln("genererBilanComplet : start")
	t, _ := template.New("BilanComplet").Parse(BilanReleveDeCarriere)
	log.Debugln("CalculerDépartTauxPleinThéorique : start")
	departTauxPlein, _ := depart.CalculerDépartTauxPlein(userData.Naissance, userData.TrimestresCotisés, userData.AnnéeRelevé)
	log.Debugln("CalculerDépartTauxPleinThéorique : ok! ", departTauxPlein)
	data, err := preparerDonneesPourTemplate(userData, infos, departTauxPlein)
	if err != nil {
		log.Fatalf("Le bilan n'a pas pu être généré, err: ", err)
	}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalf("Le bilan n'a pas pu être généré, err: ", err)
	}
	log.Debugln("genererBilanComplet : ok !")
}

func preparerDonneesPourTemplate(user InfosUser, dep depart.InfosDepartEnRetraite, tauxPlein depart.CalculDépart) (TemplateReleveData, error) {

	dateRelevé, err := depart.VérifierRelevé(user.AnnéeRelevé)
	if err != nil {
		log.Debugf("Impossible de calculer le nombre de trimestres cotisés, à cause de la date du relevé: %v, err: %v", user.AnnéeRelevé, err)
		return TemplateReleveData{}, fmt.Errorf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
	}

	nbTrimestresRestants, err := depart.NombreDeTrimestresEntre(dateRelevé, dep.DateDépartMin)
	if err != nil {
		log.Debugf("Impossible de calculer le nombre de trimestres cotisés, à cause de la date du relevé: %v, err: %v", user.AnnéeRelevé, err)
		return TemplateReleveData{}, fmt.Errorf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
	}

	departMinTrimestresCotises := nbTrimestresRestants + user.TrimestresCotisés
	departMinTrimestresManquants := dep.TrimestresTauxPlein - departMinTrimestresCotises

	calcul, _ := depart.CalculerDécotePourTrimestresManquants(departMinTrimestresManquants, user.Naissance)
	min := InfosDepartMin{
		InfosDepart: InfosDepart{
			DateDépart:        depart.TimeToString(dep.DateDépartMin),
			AgeDépart:         dep.AgeDépartMin,
			TrimestresCotisés: departMinTrimestresCotises,
		},
		TrimestresManquants: departMinTrimestresManquants,
		Decote:              calcul,
	}

	return TemplateReleveData{
		User:      user,
		Infos:     dep,
		DepartMin: min,
		//TauxPlein:plein,
	}, nil

}
