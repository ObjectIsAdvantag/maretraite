
# Que trouverez-vous ici ?

Des outils qui vous permettront de vous �clairer quant � votre future retraite, 
si vous �tes dans le cas d'un salari� du priv� qui a cotis� au r�gime g�n�ral.

Vous pouvez consulter votre bilan retraite :
- [X] en lan�ant l' ex�cutable maretraite : [t�l�charger](https://github.com/ObjectIsAdvantag/maretraite/releases) pour linux, windows,
- [X] via docker : docker pull objectisadvantag/maretraite; docker run -it objectisadvantag/maretraite
- [X] � partir du code source : git clone https://github.com/ObjectIsAdvantag/maretraite; make
- [ ] ou encore � partir de la WebAPI (non impl�ment�).

## Exemple

``` bash
$ ./maretraite

Votre date de naissance (JJ/MM/YYYY): 24/12/1971

Votre bilan retraite simplifi� suite � la r�forme 2010 :

- vous devrez cotiser 171 trimestres pour toucher une retraite � taux plein
   - une retraite � taux plein correspond � une pension de l'ordre de 50% de vos 25 meilleurs ann�es.

- vous pourrez partir en retraite au plus t�t � 62 ans, le 24/12/2033
   - si vous avez avez cotis� 151 trimestres au minimum
   - sans quoi vous devriez repousser votre demande de d�part en retraite

   - principe de la d�cote :
      - votre pension est diminu�e de 0.625 points par trimestre manquant par rapport au taux plein
      - ex: vous demandez � partir en retraite apr�s le 24/12/2033 et avez cotis� 151 trimestres,
            soient 20 trimestres manquants par rapport au taux plein (171 trimestres),
            votre pension serait alors de l'ordre de 37.5% de vos 25 meilleures ann�es

   - principe de la surcote :
      - votre pension est augment�e de 0.625 points par trimestre suppl�mentaire cotis�
      - ex: vous demandez � partir en retraite apr�s le 24/12/2033 et avez cotis� 179 trimestres,
            soient 8 trimestres suppl�mentaires par rapport au taux plein (171 trimestres),
            votre pension serait alors de l'ordre de 55% de vos 25 meilleures ann�es

- � partir du 24/12/2038, vous pourrez automatiquement b�n�ficier d'une retraite � taux plein,
   - et ce, quelque soit votre nombre de trimestres cotis�s,
   - car vous aurez atteint l'�ge l�gal de 67 ans

- au del� du 24/12/2041 si vous n'avez toujours pas demand� � partir en retraite,
   - votre employeur serait en droit de contraindre ce d�part,
   - et vous auriez alors 70 ans
```

# Pourquoi ce projet ?

Nous sommes en Janvier 2016, je viens d'avoir 44 ans, et je ne trouve pas d'outil simple pour d�terminer ma date estim�e de d�part en retraite � taux plein, ni pour calculer le montant de ma future pension en fonction de ma date de d�part en retraite. 

Relevant du r�gime g�n�ral, je pensais qu'� partir de mon dernier relev� de carri�re AGIRC/ARCO, il devrait �tre possible de faire rapidement une simulation de dates et de montants.
Je ne parle pas d'un montant - ferme � date - qui n�cessiterait de laborieuses saisies de donn�es, mais d'une rapide estimation sur la base de crit�res simples, 
du type : - je continue � travailler � plein temps, - je per�ois une r�mun�ration identique ...

Des recherche sur Google m'orientent vers des calculateurs complexes o� je dois resaisir mon relev� de carri�re, ou vers des pages simples mais sans API, enfin une recherche sur Github rapporte 5 r�sultats avec le mot clef Retraite (spirituel ou des projets de site).

Bref, c'est le moment d'apporter "my 2 euros..." en proposant un outil qui offre un aper�u de sa future retraite en moins de 5 minutes.


# Pour les d�veloppeurs 

Pr�-requis : disposer d'un environnement Go, 
[un tutoriel](https://gist.github.com/leg0ffant/3bee4829ce2ad8fd026c#file-golang-fr-md) en fran�ais pour d�couvrir le langage.

``` bash
> git clone https://github.com/ObjectIsAdvantag/maretraite
> cd maretraite
> make
> 
```


## Calculer des conditions de d�part en retraite

A partir de votre date de naissance, d�terminez les dates importantes pour votre retraite :
- l'�ge minimal auquel vous pourrez pr�tendre � une retraite (soit parce que vous aurez suffisamment cotis�, soit via le rachat de trimestres)
- l'�ge auquel vous serez assurez de pouvoir prendre votre retraite 
A partir de votre relev� de situation individuelle, d�terminez � quel moment vous pourrez partir en retraite � taux plein.

En page 2, recherchez le tableau Retraite de base, ligne "Salari� du r�gime de s�curit� sociale (CNAV) - ANNEE". 
Le nombre de trimestres de votre retraite de base, ainsi que l'ann�e sont pr�cis�s ici.

Par exemple : vous disposez de 87 trimestres en 2014 (fin 2014 en fait), reporter 87 et 2014.
``` go
// Evaluer son d�part en retraite � taux plein avec les informations extraites du relev� de situation individuelle
calcul := CalculerD�partTauxPleinTh�orique("DATE DE NAISSANCE", "NOMBRE DE TRIMESTRES", "ANNEE DU RELEVE")
fmt.Printf("SANS interruption de cotisations, vous pourriez partir avec un taux plein le %s, � l'�ge de %s, en ayant cotis� %d trimestres au total, soit un reliquat de %d trimestres", 
      calcul.Date, calcul.Age, calcul.TrimestresCotis�s, calcul.TrimestresRestants)
```

## Simuler le montant d'une future pension

Les algos permettent de calculer une d�c�te pour un nombre de trimestres manquants.

Je m'interroge sur la meilleure fa�on de pr�senter une information rapidement.
Aussi, j'attendrais d'avoir des contributions sur ce sujet, notamment de la part de conseillers en retraite.

En attendant, voici un pointeur utile : https://www.lassuranceretraite.fr/portail-info/home/salaries/age-et-montant-de-ma-retraite/quel-sera-montant-ma-retraite/le-calcul-en-detail.html


## Disclaimer
 
Le code source a la particularit� d'�tre "orient� fran�ais" avec accentutation (types, fonctions, variables). 

C'est une exp�rience int�ressante et parfois bizarre...



# License

MIT, see license file.

Feel free to use, reuse, extend, and contribute



 

