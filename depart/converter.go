// Fonction utilitaire de conversion d'une chaine de caractères en date
package depart


import (
	"errors"
	"time"
	"fmt"
)

const JJMMAAADateFormat = "2/1/2006"

var ErrDateVide = errors.New("la date n'est pas renseignée")
var ErrDateFormatInvalide = errors.New("le format de date n'est pas valide")
var ErrDateLimites = errors.New("la date n'est pas entre le 01/01/1900 et aujourd'hui")


func parseDate(input string) (time.Time, error) {
	if input == "" {
		return time.Time{}, ErrDateVide
	}

	// Parser la date
	res, err := time.Parse(JJMMAAADateFormat, input)
	if err != nil {
		return time.Time{}, ErrDateFormatInvalide
	}

	// Vérifier que la date est bien entre le 1/1/1900 et aujourd'hui
	min, _ := time.Parse(JJMMAAADateFormat, fmt.Sprintf("01/01/%d", ANNEE_MIN))
	if res.Before(min) || res.After(time.Now()) {
		return time.Time{}, ErrDateLimites
	}

	return res, nil
}