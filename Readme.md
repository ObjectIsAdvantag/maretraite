
# Que trouverez-vous ici ?

Quelques outils qui vous permettront de vous �clairer quant � votre future retraite, 
si vous �tes dans le cas d'un salari� du priv� qui a cotis� au r�gime g�n�ral.

Les calculs sont d�velopp�s en language Go ou encore #golang. [Un tutoriel](https://gist.github.com/leg0ffant/3bee4829ce2ad8fd026c#file-golang-fr-md) en fran�ais pour d�couvrir le langage.

Vous pouvez utiliser ces codes soit en partant de l'exemple depart_test.go et en saissant vos donn�es personnelles,
ou bien en soumettant vos donn�es � la WebAPI.


## Consulter mes conditions de d�part en retraite

A partir de votre relev� de situation individuelle, d�terminez � quel moment vous pourrez partir en retraite � taux plein.

```

```


## Simuler le montant de sa future cotision retraite
 
Voir https://www.lassuranceretraite.fr/portail-info/home/salaries/age-et-montant-de-ma-retraite/quel-sera-montant-ma-retraite/le-calcul-en-detail.html


# Pourquoi ?

Nous sommes en Janvier 2016, je viens d'avoir 44 ans, et je ne trouve pas d'outil simple pour calculer le montant de ma future retraite.
A partir de mon dernier relev� de carri�re AGIRC/ARCO, il devrait �tre possible de faire rapidement une simulation.
Je ne parle pas d'un montant ferme mais d'une rapide estimation sur la base de crit�res simples (du type je continue � travailler au m�me rythme).
Des recherche sur Google m'orientent vers des calculateurs complexes o� je dois resaisir mon relev� de carri�re, ou vers des pages simples mais sans API, enfin une recherche sur Github rapporte 5 r�sultats avec le mot clef Retraite (spirituel ou des projets de site).
Bref, c'est le moment d'apporter my 2 euros...



# Recherches

## Sites

- l'info retraite : [documentation](http://www.info-retraite.fr/) sur les formalit�s li�es � sa retraite, ainsi que le simulateur [M@rel](http://www.marel.fr/) 
- l'assurance retraite : [Calculer](https://www.lassuranceretraite.fr/portail-info/home/salaries/calculer-mon-age-de-depart.html#) son age de d�part � la retraite  
- ma retraite en clair : [documentation et calculateur](http://www.la-retraite-en-clair.fr/cid3190637/comment-calculer-pension-retraite.html) 


## Repos Github

- [sgmap](https://github.com/sgmap/retraite) : code d'un site pr�sentant des d�marches personnalis�es de d�part � la retraite. Le code a �t� cr�e par [@xnpore](https://twitter.com/xnopre), Quelques donn�es peuvent pertinentes � r�-exposer sous forme d'API :
   - caisses : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/11.sql
   - d�part l�gal : https://github.com/sgmap/retraite/blob/master/server/db/evolutions/12.sql


# License

MIT, see license file.

Feel free to use, reuse, extend, and contribute



 

