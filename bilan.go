// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Point d'entrée du calculateur de retraite pour des salariés du privé.
// Ce programme interactif se lance dans une console.
// Le programme interroge l'utilisateur et retourne ses données de départ en retraite sous forme textuelle.
package main

import (
	"os"
	"fmt"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/ObjectIsAdvantag/retraite/depart"
	"github.com/ObjectIsAdvantag/retraite/montant"
)



// Point d'entrée du programme
func main() {
	// Interroge l'utilisateur
	log.Debugln("Demande des informations : start")
	userData, _ := interrogerUtilisateur()
	log.Debugln("Demande des informations : ok! ", userData.Naissance, userData.DateReleve, userData.TrimestresCostisés)

	// Calcule les données de départ en retraite
	log.Debugln("CalculerDépartLégal : start")
	infosLegales, _ := depart.CalculerDépartLégal(userData.Naissance)
	log.Debugln("CalculerDépartLégal : ok! ", infosLegales)

	log.Debugln("CalculerDépartTauxPleinThéorique : start")
	departTauxPlein, _ := depart.CalculerDépartTauxPleinThéorique(userData.Naissance, userData.TrimestresCostisés, userData.DateReleve)
	log.Debugln("CalculerDépartTauxPleinThéorique : ok! ", departTauxPlein)

	// Génère le bilan
	log.Debugln("Génération du bilan : start")
	t, _ := template.New("bilan").Parse(TexteBilan)
	data, err := preparerDonneesPourTemplate(userData, infosLegales, departTauxPlein)
	if err != nil {
		log.Fatalf("Le bilan n'a pas pu être généré, err: ", err)
	}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalf("Le bilan n'a pas pu être généré, err: ", err)
	}
	log.Debugln("Génération du bilan : ok !")

}

// Première implémentation statique
func interrogerUtilisateur() (InfosUser, error) {
	return InfosUser{Naissance:"24/12/1971", TrimestresCostisés:87, DateReleve:2014}, nil
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	// [PENDING]
	log.SetLevel(log.DebugLevel)
}

func preparerDonneesPourTemplate(user InfosUser, dep depart.InfosDepartEnRetraite, tauxPlein depart.CalculDépart) (TemplateData, error) {
	minDecote := CoeffRetraite {
		Type: "décote",
		Période: "trimestre",
		Valeur: fmt.Sprintf("%.3f%%", montant.Default.Diminution),
	}
	dateRelevé, err := depart.CalculerDateReleve(user.DateReleve)
	if err != nil {
		log.Debugf("Impossible de calculer le nombre de trimestres cotisés, à cause du relevé, err: %v", err)
		return TemplateData{}, fmt.Errorf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
	}
	delta, err := depart.CalculerDurée(dateRelevé, dep.DateDépartMin)
	if err != nil {
		log.Debugf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
		return TemplateData{}, fmt.Errorf("Impossible de calculer le nombre de trimestres cotisés, err: %v", err)
	}
	minCotises := int(delta.EnMois() / 4) + user.TrimestresCostisés
	minManquants := dep.TrimestresTauxPlein - minCotises
	//var dateMinManquants time.Time
	var minMontantDecote float32
	if minManquants <= 20 {
		minMontantDecote, _ = montant.DécotePourTrimestresManquants(minManquants, user.Naissance)
		//dateMinManquants
	} else {
		minMontantDecote, _ =  montant.DécotePourTrimestresManquants(20, user.Naissance)
		//dateMinManquants
	}

	min := InfosDepart{
		DateDépart: depart.TimeToString(dep.DateDépartMin),
		AgeDépart: dep.AgeDépartMin,
		TrimestresCotisés: minCotises,
		TrimestresManquants: minManquants,
		Cote: minDecote,
		MontantCote: fmt.Sprintf("%.3f%%", minMontantDecote),
	}

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

	return TemplateData{
		User:user,
		DepartMin:min,
		TauxPlein:plein,
	}, nil

}
