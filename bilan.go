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
)

type UserData struct {
	Naissance        string
	DateReleve       int
	TrimestresAcquis int
}

type templateData struct {
	User      UserData
	Infos     depart.InfosDepartEnRetraite
	TauxPlein depart.CalculDépart
}

// Point d'entrée du programme
func main() {
	// Interroge l'utilisateur
	log.Debugln("Demande des informations : start")
	userData, _ := interrogerUtilisateur()
	log.Debugln("Demande des informations : ok! ", userData.Naissance, userData.DateReleve, userData.TrimestresAcquis)

	// Calcule les données de départ en retraite
	log.Debugln("CalculerDépartLégal : start")
	infosLegales, _ := depart.CalculerDépartLégal(userData.Naissance)
	log.Debugln("CalculerDépartLégal : ok! ", infosLegales)

	log.Debugln("CalculerDépartTauxPleinThéorique : start")
	departTauxPlein, _ := depart.CalculerDépartTauxPleinThéorique(userData.Naissance, userData.TrimestresAcquis, userData.DateReleve)
	log.Debugln("CalculerDépartTauxPleinThéorique : ok! ", departTauxPlein)

	// Génère le bilan
	log.Debugln("Génération du bilan : start")
	//[PENDING] use template during dev, integrate into binary for production
	t, _ := template.New("bilan").Parse(TexteBilan)
	//	templateFile := "bilan.template"
	//	t, err := template.New(templateFile).ParseFiles(templateFile)
	//	if err != nil {
	//		log.Warnln("Impossible de lire le template: ", templateFile, ", err: ", err)
	//		os.Exit(1)
	//	}
	data := templateData{userData, infosLegales, departTauxPlein}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Warnln("Le bilan n'a pas pu être généré, err: ", err)
	}
	log.Debugln("Génération du bilan : ok !")

}

// Première implémentation statique
func interrogerUtilisateur() (UserData, error) {
	return UserData{Naissance:"24/12/1971", TrimestresAcquis:87, DateReleve:2014}, nil
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	// [PENDING]
	log.SetLevel(log.DebugLevel)
}
