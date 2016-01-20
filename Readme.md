
# Que trouverez-vous ici ?

Quelques outils qui vous permettront de vous éclairer quant à votre future retraite, 
si vous êtes dans le cas d'un salarié du privé qui a cotisé au régime général.

Les calculs sont développés en language Go ou encore #golang. [Un tutoriel](https://gist.github.com/leg0ffant/3bee4829ce2ad8fd026c#file-golang-fr-md) en français pour découvrir le langage.

Vous pouvez utiliser ces codes soit en partant de l'exemple depart_test.go et en saissant vos données personnelles,
ou bien en soumettant vos données à la WebAPI.


## Consulter mes conditions de départ en retraite

A partir de votre relevé de situation individuelle, déterminez à quel moment vous pourrez partir en retraite à taux plein.

```

```


## Simuler le montant de sa future cotision retraite
 
Voir https://www.lassuranceretraite.fr/portail-info/home/salaries/age-et-montant-de-ma-retraite/quel-sera-montant-ma-retraite/le-calcul-en-detail.html


# Pourquoi ?

Nous sommes en Janvier 2016, je viens d'avoir 44 ans, et je ne trouve pas d'outil simple pour calculer le montant de ma future retraite.
A partir de mon dernier relevé de carrière AGIRC/ARCO, il devrait être possible de faire rapidement une simulation.
Je ne parle pas d'un montant ferme mais d'une rapide estimation sur la base de critères simples (du type je continue à travailler au même rythme).
Des recherche sur Google m'orientent vers des calculateurs complexes où je dois resaisir mon relevé de carrière, ou vers des pages simples mais sans API, enfin une recherche sur Github rapporte 5 résultats avec le mot clef Retraite (spirituel ou des projets de site).
Bref, c'est le moment d'apporter my 2 euros...



# Recherches

## Sites

- l'info retraite : [documentation](http://www.info-retraite.fr/) sur les formalités liées à sa retraite, ainsi que le simulateur [M@rel](http://www.marel.fr/) 
- l'assurance retraite : [Calculer](https://www.lassuranceretraite.fr/portail-info/home/salaries/calculer-mon-age-de-depart.html#) son age de départ à la retraite  
- ma retraite en clair : [documentation et calculateur](http://www.la-retraite-en-clair.fr/cid3190637/comment-calculer-pension-retraite.html) 


## Repos Github

- [sgmap](https://github.com/sgmap/retraite) : code d'un site présentant des démarches personnalisées de départ à la retraite. Le code a été crée par [@xnpore](https://twitter.com/xnopre), Quelques données peuvent pertinentes à ré-exposer sous forme d'API :
   - caisses : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/11.sql
   - départ légal : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/12.sql


# License

MIT, see license file.

Feel free to use, reuse, extend, and contribute



 

