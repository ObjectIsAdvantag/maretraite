# Endpoint

http://maretraite.com/api/ 

# Bilan simple

``` http
GET /bilan-simple?naissance="24/12/1971"&media=json

{ date_de_naissance : "24/12/1971"
  taux_plein:171,
  depart_au_plus_tot : {
    age:62
    trimestres_min: 155
  }
 }

GET /faq
[
      { question:"qu'est-ce que le taux plein"
        
      }
]
```
 
``` http
GET /bilan-simple?naissance="24/12/1971"&media=txt

output du shell interactif
``` 

``` http
GET /bilan-simple?naissance="24/12/1971"&media=html

TODO
```
