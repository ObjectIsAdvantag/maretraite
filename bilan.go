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

	"github.com/ObjectIsAdvantag/retraite/depart"
	"github.com/ObjectIsAdvantag/retraite/pension"
	log "github.com/Sirupsen/logrus"
)

type InfosUser struct {
	Naissance         string
	DateRelevé        int
	TrimestresCotisés int
}

func (infos InfosUser) sansRelevé() bool {
	return (infos.TrimestresCotisés == 0) && (infos.DateRelevé == 0)
}

// Point d'entrée du programme
func main() {
	// Interroge l'utilisateur
	log.Debugln("Demande des informations : start")
	userData, err := interrogerUtilisateur()
	if err != nil {
		log.Errorf("Données incorrectes, le bilan ne peut être généré, essayez à nouveau...")
		os.Exit(-1)
	}
	log.Debugln("Demande des informations : ok! ", userData.Naissance, userData.DateRelevé, userData.TrimestresCotisés)

	// Calcule les données de départ en retraite
	log.Debugln("CalculerDépartLégal : start")
	infosLegales, _ := depart.CalculerDépartLégal(userData.Naissance)
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
	fmt.Println(`
********************************************************************************
Ce bilan vous est proposé gratuitement, sans publicité ni exploitation des données renseignées.
Il a pour objectif de vous aider à appréhender les enjeux principaux de votre future pension.
Nous vous souhaitons de vivre de beaux moments professionnels jusqu'à votre départ en retraite.

Si vous êtes un conseiller Retraite, n'hésitez pas à soumettre vos idées (améliorations, correctifs...) : https://github.com/ObjectIsAdvantag/retraite/issues/1
Si vous êtes développeur, n'hésitez pas à contribuer au projet https://github.com/ObjectIsAdvantag/retraite

Copyright 2016, Stève Sfartz - @ObjectIsAdvantag - License MIT
********************************************************************************
	`)

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
		fmt.Printf("Saisie incorrecte, nouvel essai %d/%d", essai, essaiMax)
	}

	// TODO Relevé de carrière
	//return InfosUser{Naissance:"24/12/1971", TrimestresCotisés:87, DateReleve:2014}, nil

	return InfosUser{Naissance: naissance, TrimestresCotisés: 0, DateRelevé: 0}, nil
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

	decote := pension.DécotePourTrimestresManquants(infos.TrimestresTauxPlein - pension.DECOTE_MAX_TRIMESTRES, userData.Naissance)
	simuDec := Simulation {
		// la décôte est limitée à 20 trimestres
		Trimestres:infos.TrimestresTauxPlein - pension.DECOTE_MAX_TRIMESTRES,
		DeltaTrimestres:pension.DECOTE_MAX_TRIMESTRES,
		EvolParTrimestre:fmt.Sprintf("%.3f", decote.Decote.Diminution),
		TauxPension:decote.TauxPension,
	}

	deltaTrimestres := 8
	simuSur := Simulation {
		// la surcote n'est pas limitée
		Trimestres:infos.TrimestresTauxPlein + deltaTrimestres,
		DeltaTrimestres:deltaTrimestres,
		EvolParTrimestre:fmt.Sprintf("%.3f", pension.SurcotePourTrimestreSupplementaire()),
		TauxPension:50.0 + float32(deltaTrimestres) * pension.SurcotePourTrimestreSupplementaire(),
	}

	data := TemplateSimpleData{
		User:userData,
		Infos:infos,
		SimuDecote:simuDec,
		SimuSurcote:simuSur,
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
	departTauxPlein, _ := depart.CalculerDépartTauxPleinThéorique(userData.Naissance, userData.TrimestresCotisés, userData.DateRelevé)
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

func preparerDonneesPourTemplate(user InfosUser, dep depart.InfosDepartEnRetraite, tauxPlein depart.CalculDépart) (TemplateData, error) {

	dateRelevé, err := depart.CalculerDateReleve(user.DateRelevé)
	if err != nil {
		log.Debugf("Impossible de calculer le nombre de trimestres cotisés, à cause de la date du relevé: %v, err: %v", user.DateRelevé, err)
		return TemplateData{}, fmt.Errorf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
	}

	delta, err := depart.CalculerDurée(dateRelevé, dep.DateDépartMin)
	if err != nil {
		log.Debugf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
		return TemplateData{}, fmt.Errorf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
	}
	departMinTrimestresCotises := int(delta.AgeEnMoisFloat()/4) + user.TrimestresCotisés
	departMinTrimestresManquants := dep.TrimestresTauxPlein - departMinTrimestresCotises

	calcul := pension.DécotePourTrimestresManquants(departMinTrimestresManquants, user.Naissance)

	//MontantCote: fmt.Sprintf("%.3f%%", minMontantDecote),

	//Valeur: fmt.Sprintf("%.3f%%", pension.Default.Diminution),

	min := InfosDepartMin{
		InfosDepart: InfosDepart{
			DateDépart:        depart.TimeToString(dep.DateDépartMin),
			AgeDépart:         dep.AgeDépartMin,
			TrimestresCotisés: departMinTrimestresCotises,
		},
		TrimestresManquants: departMinTrimestresManquants,
		Decote:              calcul,
	}

	/*
		std := CoeffRetraite {
			Type: "taux plein",
			Période: "aucune",
			Valeur: "0%",
		}

		plein := InfosDepart{
			DateDépart: depart.TimeToString(tauxPlein.Date),
			AgeDépart: tauxPlein.Age,
			TrimestresCotisés: tauxPlein.TrimestresCotisés,
			TrimestresManquants: tauxPlein.TrimestresRestants,
			Cote: std,
		}
	*/

	return TemplateData{
		User:      user,
		Infos:     dep,
		DepartMin: min,
		//TauxPlein:plein,
	}, nil

}
