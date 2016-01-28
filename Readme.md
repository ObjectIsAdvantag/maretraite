
# Que trouverez-vous ici ?

Des outils qui vous permettront de vous �clairer quant � votre future retraite, 
si vous �tes dans le cas d'un salari� du priv� qui a cotis� au r�gime g�n�ral.

Vous pouvez utiliser ce projet :
- [x] ou � partir de la CLI (en cours) : [t�l�charger](https://github.com/ObjectIsAdvantag/retraite/releases) un ex�cutable pour votre plateforme,
- [X] � partir du code GO, et en adaptant l'exemple plac� dans depart_test.go, avec vos donn�es personnelles,
- [ ] ou encore � partir de la WebAPI (non impl�ment�).


# Pourquoi ce projet ?

Nous sommes en Janvier 2016, je viens d'avoir 44 ans, et je ne trouve pas d'outil simple pour d�terminer ma date estim�e de d�part en retraite � taux plein, ni pour calculer le montant de ma future pension en fonction de ma date de d�part en retraite. 

Relevant du r�gime g�n�ral, je pensais qu'� partir de mon dernier relev� de carri�re AGIRC/ARCO, il devrait �tre possible de faire rapidement une simulation de dates et de montants.
Je ne parle pas d'un montant - ferme � date - qui n�cessiterait de laborieuses saisies de donn�es, mais d'une rapide estimation sur la base de crit�res simples, 
du type : - je continue � travailler � plein temps, - je per�ois une r�mun�ration identique ...

Des recherche sur Google m'orientent vers des calculateurs complexes o� je dois resaisir mon relev� de carri�re, ou vers des pages simples mais sans API, enfin une recherche sur Github rapporte 5 r�sultats avec le mot clef Retraite (spirituel ou des projets de site).

Bref, c'est le moment d'apporter "my 2 euros..." en proposant un outil qui offre un aper�u de sa future retraite en moins de 5 minutes.


# Ressources aff�rentes au domaine de la retraite

## Sites

- l'info retraite : [documentation](http://www.info-retraite.fr/) sur les formalit�s li�es � sa retraite, ainsi que le simulateur [M@rel](http://www.marel.fr/), par le Groupement des r�gimes de retraite obligatoires de base et compl�mentaires de S�curit� sociale,
- l'assurance retraite : [Calculer](https://www.lassuranceretraite.fr/portail-info/home/salaries/calculer-mon-age-de-depart.html#) son age de d�part � la retraite. Il est possible de cr�er un compte en renseignant ses donn�es personnelles (num�ro de s�curit� sociale, adresse) 
- ma retraite en clair : [documentation et calculateur](http://www.la-retraite-en-clair.fr/cid3190637/comment-calculer-pension-retraite.html) 
- l'agirc-arcco : documentation et calculateurs sur le [site](http://www.agirc-arrco.fr/), il est possible de se cr�er un compte depuis le site de l'assurance retraite puis d'acc�der aux [services Agirc-arrco](https://services.agirc-arrco.fr) 


## Repos Github

- [sgmap](https://github.com/sgmap/retraite) : code d'un site pr�sentant des d�marches personnalis�es de d�part � la retraite. Le code a �t� cr�e par [@xnpore](https://twitter.com/xnopre), Quelques donn�es peuvent pertinentes � r�-exposer sous forme d'API :
   - caisses : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/11.sql
   - d�part l�gal : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/12.sql


# Pour les d�veloppeurs 

Pr�-requis : disposer d'un environnement Go, 
[un tutoriel](https://gist.github.com/leg0ffant/3bee4829ce2ad8fd026c#file-golang-fr-md) en fran�ais pour d�couvrir le langage.

``` bash
> git clone https://github.com/ObjectIsAdvantag/retraite
> cd retraite
> make
> 
```


## Consulter mes conditions de d�part en retraite

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

## Simuler le montant de sa future cotision retraite

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



 

