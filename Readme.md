
# Que trouverez-vous ici ?

Des outils qui vous permettront de vous éclairer quant à votre future retraite, 
si vous êtes dans le cas d'un salarié du privé qui a cotisé au régime général.

Vous pouvez consulter votre bilan retraite :
- [X] en lançant l' exécutable maretraite : [télécharger](https://github.com/ObjectIsAdvantag/maretraite/releases) pour linux, windows,
- [X] via docker : docker pull objectisadvantag/maretraite; docker run -it objectisadvantag/maretraite
- [X] à partir du code source : git clone https://github.com/ObjectIsAdvantag/maretraite; make
- [ ] ou encore à partir de la WebAPI (non implémenté).

## Exemple

``` bash
$ ./maretraite

Votre date de naissance (JJ/MM/YYYY): 24/12/1971

Votre bilan retraite simplifié suite à la réforme 2010 :

- vous devrez cotiser 171 trimestres pour toucher une retraite à taux plein
   - une retraite à taux plein correspond à une pension de l'ordre de 50% de vos 25 meilleurs années.

- vous pourrez partir en retraite au plus tôt à 62 ans, le 24/12/2033
   - si vous avez avez cotisé 151 trimestres au minimum
   - sans quoi vous devriez repousser votre demande de départ en retraite

   - principe de la décote :
      - votre pension est diminuée de 0.625 points par trimestre manquant par rapport au taux plein
      - ex: vous demandez à partir en retraite après le 24/12/2033 et avez cotisé 151 trimestres,
            soient 20 trimestres manquants par rapport au taux plein (171 trimestres),
            votre pension serait alors de l'ordre de 37.5% de vos 25 meilleures années

   - principe de la surcote :
      - votre pension est augmentée de 0.625 points par trimestre supplémentaire cotisé
      - ex: vous demandez à partir en retraite après le 24/12/2033 et avez cotisé 179 trimestres,
            soient 8 trimestres supplémentaires par rapport au taux plein (171 trimestres),
            votre pension serait alors de l'ordre de 55% de vos 25 meilleures années

- à partir du 24/12/2038, vous pourrez automatiquement bénéficier d'une retraite à taux plein,
   - et ce, quelque soit votre nombre de trimestres cotisés,
   - car vous aurez atteint l'âge légal de 67 ans

- au delà du 24/12/2041 si vous n'avez toujours pas demandé à partir en retraite,
   - votre employeur serait en droit de contraindre ce départ,
   - et vous auriez alors 70 ans
```

# Pourquoi ce projet ?

Nous sommes en Janvier 2016, je viens d'avoir 44 ans, et je ne trouve pas d'outil simple pour déterminer ma date estimée de départ en retraite à taux plein, ni pour calculer le montant de ma future pension en fonction de ma date de départ en retraite. 

Relevant du régime général, je pensais qu'à partir de mon dernier relevé de carrière AGIRC/ARCO, il devrait être possible de faire rapidement une simulation de dates et de montants.
Je ne parle pas d'un montant - ferme à date - qui nécessiterait de laborieuses saisies de données, mais d'une rapide estimation sur la base de critères simples, 
du type : - je continue à travailler à plein temps, - je perçois une rémunération identique ...

Des recherche sur Google m'orientent vers des calculateurs complexes où je dois resaisir mon relevé de carrière, ou vers des pages simples mais sans API, enfin une recherche sur Github rapporte 5 résultats avec le mot clef Retraite (spirituel ou des projets de site).

Bref, c'est le moment d'apporter "my 2 euros..." en proposant un outil qui offre un aperçu de sa future retraite en moins de 5 minutes.


# Ressources afférentes au départ et calcul de sa retraite

## Sites

- l'info retraite : [documentation](http://www.info-retraite.fr/) sur les formalités liées à sa retraite, ainsi que le simulateur [M@rel](http://www.marel.fr/), par le Groupement des régimes de retraite obligatoires de base et complémentaires de Sécurité sociale,
- l'assurance retraite : [Calculer](https://www.lassuranceretraite.fr/portail-info/home/salaries/calculer-mon-age-de-depart.html#) son age de départ à la retraite. Il est possible de créer un compte en renseignant ses données personnelles (numéro de sécurité sociale, adresse) 
- ma retraite en clair : [documentation et calculateur](http://www.la-retraite-en-clair.fr/cid3190637/comment-calculer-pension-retraite.html) 
- l'agirc-arcco : documentation et calculateurs sur le [site](http://www.agirc-arrco.fr/), il est possible de se créer un compte depuis le site de l'assurance retraite puis d'accéder aux [services Agirc-arrco](https://services.agirc-arrco.fr) 


## Repos Github

- [sgmap](https://github.com/sgmap/retraite) : code d'un site présentant des démarches personnalisées de départ à la retraite. Le code a été crée par [@xnpore](https://twitter.com/xnopre), Quelques données peuvent pertinentes à ré-exposer sous forme d'API :
   - caisses : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/11.sql
   - départ légal : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/12.sql


# Pour les développeurs 

Pré-requis : disposer d'un environnement Go, 
[un tutoriel](https://gist.github.com/leg0ffant/3bee4829ce2ad8fd026c#file-golang-fr-md) en français pour découvrir le langage.

``` bash
> git clone https://github.com/ObjectIsAdvantag/maretraite
> cd maretraite
> make
> 
```


## Calculer des conditions de départ en retraite

A partir de votre date de naissance, déterminez les dates importantes pour votre retraite :
- l'âge minimal auquel vous pourrez prétendre à une retraite (soit parce que vous aurez suffisamment cotisé, soit via le rachat de trimestres)
- l'âge auquel vous serez assurez de pouvoir prendre votre retraite 
A partir de votre relevé de situation individuelle, déterminez à quel moment vous pourrez partir en retraite à taux plein.

En page 2, recherchez le tableau Retraite de base, ligne "Salarié du régime de sécurité sociale (CNAV) - ANNEE". 
Le nombre de trimestres de votre retraite de base, ainsi que l'année sont précisés ici.

Par exemple : vous disposez de 87 trimestres en 2014 (fin 2014 en fait), reporter 87 et 2014.
``` go
// Evaluer son départ en retraite à taux plein avec les informations extraites du relevé de situation individuelle
calcul := CalculerDépartTauxPleinThéorique("DATE DE NAISSANCE", "NOMBRE DE TRIMESTRES", "ANNEE DU RELEVE")
fmt.Printf("SANS interruption de cotisations, vous pourriez partir avec un taux plein le %s, à l'âge de %s, en ayant cotisé %d trimestres au total, soit un reliquat de %d trimestres", 
      calcul.Date, calcul.Age, calcul.TrimestresCotisés, calcul.TrimestresRestants)
```

## Simuler le montant d'une future pension

Les algos permettent de calculer une décôte pour un nombre de trimestres manquants.

Je m'interroge sur la meilleure façon de présenter une information rapidement.
Aussi, j'attendrais d'avoir des contributions sur ce sujet, notamment de la part de conseillers en retraite.

En attendant, voici un pointeur utile : https://www.lassuranceretraite.fr/portail-info/home/salaries/age-et-montant-de-ma-retraite/quel-sera-montant-ma-retraite/le-calcul-en-detail.html


## Disclaimer
 
Le code source a la particularité d'être "orienté français" avec accentutation (types, fonctions, variables). 

C'est une expérience intéressante et parfois bizarre...


# License

MIT, see license file.

Feel free to use, reuse, extend, and contribute



 

