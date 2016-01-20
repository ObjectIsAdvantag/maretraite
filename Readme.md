
# Que trouverez-vous ici ?

Quelques outils qui vous permettront de vous éclairer quant à votre future retraite, 
si vous êtes dans le cas d'un salarié du privé qui a cotisé au régime général.

Les calculs sont développés en language Go ou encore #golang. [Un tutoriel](https://gist.github.com/leg0ffant/3bee4829ce2ad8fd026c#file-golang-fr-md) en français pour découvrir le langage.

Vous pouvez utiliser ce projet :
- [x] à partir du code GO, et en adaptant l'exemple placé dans depart_test.go, avec vos données personnelles,
- [ ] ou à partir de la CLI (non implémenté)
- [ ] ou encore à partir de la WebAPI (non implémenté)


## Consulter mes conditions de départ en retraite

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


## Simuler le montant de sa future cotision retraite
 
Voir https://www.lassuranceretraite.fr/portail-info/home/salaries/age-et-montant-de-ma-retraite/quel-sera-montant-ma-retraite/le-calcul-en-detail.html


# Pourquoi ce projet ?

Nous sommes en Janvier 2016, je viens d'avoir 44 ans, et je ne trouve pas d'outil simple pour connaître ma date estimée de départ en retraite à taux plein, ni pour calculer le montant de ma future retraite selon que je parte à l'âge minimal, ou lorsque j'ai atteint le taux plein, à moins d'ailleurs que je n'atteigne l'âge automatique de départ en retraite avant cette date.

Relevant du régime général, je pensais qu'à partir de mon dernier relevé de carrière AGIRC/ARCO, il devait être possible de faire rapidement une simulation de dates et de montants.
Sur ce dernier point, je ne parle pas d'un montant ferme qui nécessiterait des saisies exhaustives de données, mais d'une rapide estimation sur la base de critères simples (du type je continue à travailler au même rythme, je continue à gagner la même somme d'argent).

Des recherche sur Google m'orientent vers des calculateurs complexes où je dois resaisir mon relevé de carrière, ou vers des pages simples mais sans API, enfin une recherche sur Github rapporte 5 résultats avec le mot clef Retraite (spirituel ou des projets de site).

Bref, c'est le moment d'apporter "my 2 euros..."



# Recherches

## Sites

- l'info retraite : [documentation](http://www.info-retraite.fr/) sur les formalités liées à sa retraite, ainsi que le simulateur [M@rel](http://www.marel.fr/), par le Groupement des régimes de retraite obligatoires de base et complémentaires de Sécurité sociale,
- l'assurance retraite : [Calculer](https://www.lassuranceretraite.fr/portail-info/home/salaries/calculer-mon-age-de-depart.html#) son age de départ à la retraite. Il est possible de créer un compte en renseignant ses données personnelles (numéro de sécurité sociale, adresse) 
- ma retraite en clair : [documentation et calculateur](http://www.la-retraite-en-clair.fr/cid3190637/comment-calculer-pension-retraite.html) 
- l'agirc-arcco : documentation et calculateurs sur le [site](http://www.agirc-arrco.fr/), il est possible de se créer un compte depuis le site de l'assurance retraite puis d'accéder aux [services Agirc-arrco](https://services.agirc-arrco.fr) 


## Repos Github

- [sgmap](https://github.com/sgmap/retraite) : code d'un site présentant des démarches personnalisées de départ à la retraite. Le code a été crée par [@xnpore](https://twitter.com/xnopre), Quelques données peuvent pertinentes à ré-exposer sous forme d'API :
   - caisses : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/11.sql
   - départ légal : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/12.sql


# License

MIT, see license file.

Feel free to use, reuse, extend, and contribute



 

