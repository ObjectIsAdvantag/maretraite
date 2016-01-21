// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Point d'entrée du calculateur de retraite pour des salariés du privé.
// Ce programme interactif se lance dans une console.
// Le programme interroge l'utilisateur et retourne ses données de départ en retraite sous forme textuelle.
package main

import (
	"text/template"
	"os"

	"github.com/ObjectIsAdvantag/retraite/depart"
	log "github.com/Sirupsen/logrus"
)

type templateData struct {
	infos 	depart.InfosDepartEnRetraite
	tauxPlein	depart.CalculDépart
}

const texteBilan = ``

// Point d'entrée du programme
func main() {
	// Interroge l'utilisateur
	log.Debugln("Demande des informations : start")
	dateDeNaissance, acquis, relevé , _ := interrogerUtilisateur()
	log.Debugln("Demande des informations : ok! ", dateDeNaissance, acquis, relevé)

	// Calcule les données de départ en retraite
	log.Debugln("CalculerDépartLégal : start")
	infosLegales, _ := depart.CalculerDépartLégal(dateDeNaissance)
	log.Debugln("CalculerDépartLégal : ok! ", infosLegales)

	log.Debugln("CalculerDépartTauxPleinThéorique : start")
	departTauxPlein, _ := depart.CalculerDépartTauxPleinThéorique(dateDeNaissance, acquis, relevé)
	log.Debugln("CalculerDépartTauxPleinThéorique : ok! ", departTauxPlein)

	// Génère le bilan
	log.Debugln("Génération du bilan : start")
	t , _ := template.New("bilan").Parse(texteBilan)
	data := templateData{ infosLegales, departTauxPlein}
	t.Execute(os.Stdout, data)
	log.Debugln("Génération du bilan : ok !")

}

// Première implémentation statique
func interrogerUtilisateur() (string, int, int, error) {
	return "24/12/1971", 87, 2014, nil
}

func init() {
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	// [PENDING]
	log.SetLevel(log.DebugLevel)
}