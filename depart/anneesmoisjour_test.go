package depart

import (
	"testing"
)

func Test1(t *testing.T) {
	depuis, err := GetTime(1971, 12, 24)
	if err != nil {
		t.Errorf("calculer age error, erreur conversion err: %s", err)
	}
	jusque, err := GetTime(2016, 1, 20)
	if err != nil {
		t.Errorf("calculer age error, erreur conversion err: %s", err)
	}

	age, err := CalculerAge(depuis, jusque)

	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 0, 27}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test2(t *testing.T) {
	depuis, _ := GetTime(1971, 12, 24)
	jusque, _ := GetTime(2016, 1, 5)

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 0, 12}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test3(t *testing.T) {
	depuis, _ := GetTime(1971, 12, 24)
	jusque, _ := GetTime(2016, 1, 25)

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 1, 1}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test4(t *testing.T) {
	depuis, _ := GetTime(1971, 12, 24)
	jusque, _ := GetTime(2016, 1, 31)

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 1, 7}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test5(t *testing.T) {
	depuis, _ := GetTime(1971, 12, 24)
	jusque, _ := GetTime(2016, 1, 1)

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 0, 7}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test6(t *testing.T) {
	depuis, _ := GetTime(1971, 12, 24)
	jusque, _ := GetTime(2016, 2, 1)

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 1, 8}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

