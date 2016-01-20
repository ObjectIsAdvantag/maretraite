package depart

import (
	"testing"
)


func Test1(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1971, 12, 24})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 20})


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
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1971, 12, 24})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 20})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 0, 27}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}


func Test3(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1971, 12, 24})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 25})

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
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1971, 12, 24})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 31})

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
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1971, 12, 24})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2015, 12, 31})

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
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1971, 12, 24})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 1})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 44, 0, 8}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test7(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 7, 4})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42, 0, 0}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test8(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 8, 4})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,1, 0}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test9(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 7, 20})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,0, 16}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}


func Test10(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 8, 20})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,1, 16}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test10b(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 12, 20})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,5, 16}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test10c(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 12, 31})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,5, 27}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test10d(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2017, 1, 1})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,5, 28}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test10e(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2017, 1, 2})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,5, 29}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test10f(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2017, 1, 3})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42,5, 30}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test10g(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2017, 1, 4})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 42, 6, 0}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test11(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 4})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 41, 6, 0}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}


func Test12(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 5})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 41, 6, 1}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test13(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 1, 3})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 41, 5, 30}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test14(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 2, 4})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 41, 7, 0}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test15(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 2, 3})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 41, 6, 30}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}

func Test16(t *testing.T) {
	depuis, _ := AnneesMoisJourToTime(AnneesMoisJours{1974, 7, 4})
	jusque, _ := AnneesMoisJourToTime(AnneesMoisJours{2016, 2, 2})

	age, err := CalculerAge(depuis, jusque)
	if err != nil {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, err:%s:", depuis, jusque, err)
	}

	expected := AnneesMoisJours{ 41, 6, 29}
	if age != expected {
		t.Errorf("calculer age error, depuis: %s, jusque:%s, attendu: %s, obtenu: %s", depuis, jusque, expected, age)
	}
}
