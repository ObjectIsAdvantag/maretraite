// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Point d'entrée du calculateur de retraite pour des salariés du privé.
// Ce programme interactif se lance dans une console.
// Le programme interroge l'utilisateur et retourne ses données de départ en retraite sous forme textuelle.
package main

import (
	"os"
	"text/template"

	"github.com/ObjectIsAdvantag/retraite/depart"
	log "github.com/Sirupsen/logrus"
	"github.com/ObjectIsAdvantag/retraite/montant"
	"fmt"
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
	data := preparerDonneesPourTemplate(userData, infosLegales, departTauxPlein)
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Warnln("Le bilan n'a pas pu être généré, err: ", err)
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

func preparerDonneesPourTemplate(user InfosUser, dep depart.InfosDepartEnRetraite, tauxPlein depart.CalculDépart) TemplateData {
	decote := CoeffRetraite {
		Type: "décote",
		Période: "trimestre",
		Valeur: fmt.Sprintf("%f%%", montant.Default.Diminution),
	}
	min := InfosDepart{
		DateDépart: depart.TimeToString(dep.DateDépartMin),
		AgeDépart: dep.AgeDépartMin,
		TrimestresCotisés: 99999,
		TrimestresManquants: 99999,
		Cote: decote,
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
	}

}
