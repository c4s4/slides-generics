# Les Génériques en Go

Michel Casabianca

casa@sweetohm.net

<https://sweetohm.net/slides/slides-generics/>
---
## Avant Go 1.18

Il a toujours été possible d'écrire du code générique en Go à l'aide du type **interface{}**. Par exemple, on peut écrire une fonction qui affiche *n* fois une valeur quelconque avec :

```go
<? INCLUDE src/interface.go ?>
```
[Sur le Playground](https://go.dev/play/p/6BYtkoXfzNV)

Cet exemple est particulièrement simple car la fonction `fmt.Println()` accepte tout type. Avant Go *1.18*, sa signature était : `func Println(a ...interface{}) (n int, err error)`.

---
D'autre part, on peut définir le type d'un argument avec une interface spécifique. Par exemple:

```go
<? INCLUDE src/interface2.go ?>
```
[Sur le Playground](https://go.dev/play/p/vVi5CyxWD4Q)

Le type `error` est une interface qui définit la méthode `Error() string`. On peut donc envoyer n'importe quoi à la fonction `PrintError()` pourvu que ça implémente une méthode `Error()`.

---
## Le début des ennuis

Supposons que nous voulions écrire une fonction qui renvoie le maximum de deux valeurs. Nous pouvons l'écrire, pour les entiers, comme suit :

```go
<? INCLUDE src/maxint.go ?>
```
[Sur le Playground](https://go.dev/play/p/MeVme43ZZon)

Si nous voulons généraliser cette fonction à d'autres types, les interfaces ne nous sont pas d'un grand secours car aucune fonction ne définit les opérateurs de comparaison. Nous devrons donc **réécrire cette fonction pour tous les types** ! Il serait possible d'accepter en entrée le type `interface{}` mais nous devons alors faire des **assertions sur les types** et cela ne simplifierait pas les chose.

---
## Les Generics à la rescousse

Go *1.18* implémente les *Generics*. On peut maintenant ajouter des paramètres de type (*type parameters* en anglais) à la signature d'une fonction. Pour pouvoir rendre notre fonction `Max()` générique, nous pourrions écrire :

```go
<? INCLUDE src/maxgenerics.go ?>
```
[Sur le Playground](https://go.dev/play/p/JS5qi9JpgCm)

Ainsi, avec le paramètre de type `[N int | float64]`, nous indiquons que les paramètres de la fonction sont du type `int` ou `float64`. À noter que l'on ne peut mélanger les types, donc l'appel `Max(1, 2.0)` provoque une erreur de compilation.

---
## Le retour des interfaces

Il est aussi possible à partir de Go *1.18* de définir des interfaces comme une liste de types. Nous pourrions réécrire l'exemple précédent de la manière suivante :

```go
<? INCLUDE src/maxgenerics2.go ?>
```
[Sur le Playground](https://go.dev/play/p/pq2kbXNwZpV)

---
## Les alias de types

Si nous définissons un alias pour un type, nous pouvons l'englober dans une liste avec le caractère `~`, comme suit :

```go
<? INCLUDE src/maxgenerics3.go ?>
```
[Sur le Playground](https://go.dev/play/p/k8cOiNqF61X)

Ainsi par exemple, `~int` englobe le type `int` mais aussi tous ses alias, dont `Truc`.

---
## Contraintes

Il peut être laborieux de définir ainsi ses propres interfaces avec des listes de types. Le package *golang.org/x/exp/constraints* en propose un certain nombre :

- **Complex** : *~complex64 | ~complex128*
- **Float** : *~float32 | ~float64*
- **Integer** : *Signed | Unsigned*
- **Ordered** : *Integer | Float | ~string*
- **Signed** : *~int | ~int8 | ~int16 | ~int32 | ~int64*
- **Unsigned** : *~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr*

```go
<? INCLUDE src/maxgenerics4.go ?>
```
[Sur le Playground](https://go.dev/play/p/6-B_UhnjlkF)

---
## Instantiation

Il est possible de passer le type d'argument lors de l'appel d'une fonction générique. On pourra par exemple faire l'appel :

```go
m := Max[int](1, 2)
```

L'expression `Max[int]` est une *instantiation* de la fonction générique `Max`. Elle fixe les types des paramètres. On peut par exemple écrire :

```go
MaxFloat := Max[float64]
m := MaxFloat(1.0, 2.0)
```

La fonction `MaxFloat` est maintenant une fonction non générique qui n'accepte que des paramètres de type `float`.

---
## Types avec paramètres de type

Supposons que nous voulions faire la somme des valeurs des éléments d'une liste. Nous pourrions écrire, avec la liste chaînée standard du Go :

```go
<? INCLUDE src/list.gob ?>
```
[Sur le Playground](https://go.dev/play/p/CLY7xgTdKWs)

Cela ne compile pas car on ne peut faire une somme avec le type `interface{}` qui est celui de la valeur des éléments d'une liste : `src/list.go:12:3: invalid operation: sum += e.Value (mismatched types int and any)`. On remarquera au passage que `any` est le nouveau nom pour `interface{}`.

---
Utiliser le type `interface{}` ou `any` est ennuyeux car nous devons caster les valeurs pour pouvoir les utiliser. Il y a bien sûr une solution à base de Generics. Voici une implémentation minimaliste de liste avec des génériques :

```go
type Element[T any] struct {
	Next  *Element[T]
	Value T
}

type List[T any] struct {
	Front *Element[T]
	Last  *Element[T]
}

func (l *List[T]) PushBack(value T) {
	node := &Element[T]{
		Next:  nil,
		Value: value,
	}
	if l.Front == nil {
		l.Front = node
		l.Last = node
	} else {
		l.Last.Next = node
		l.Last = node
	}
}
```

Dans ce code nous avons ajouté des paramètres de type aux définitions types, comme dans `Element[T any]`. Cette notation indique que nous définissons un type `Element` qui contient le type `T` qui peut être quelconque.

---
Nous pouvons maintenant utiliser ces valeurs sans avoir à les caster :

```go
func main() {
	list := &List[int]{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	sum := 0
	for n := list.Front; n != nil; n = n.Next {
		sum += n.Value
	}
	println(sum)
}
```
[Sur le Playground](https://go.dev/play/p/-5lFyeE55hO)

Il est important de noter que nous avons fixé le type de la liste lors de l'instanciation :

```go
list := &List[int]{}
```

Nous avons ainsi indiqué que notre liste contient des `int` et nous pouvons alors les manipuler comme tels.

---
## Inférence de type

Nous avons vu que nous pouvons fixer le type des paramètres lors d'un appel à une fonction générique avec :

```go
m := Max[int](1, 2)
```

Dans ce cas, le compilateur sait le type des paramètres parce qu'on lui indique. Mais lorsque nous écrivons :

```go
m := Max(1, 2)
```

Le compilateur **infère le type des paramètres** de la fonction générique de celui des arguments lors de l'appel. Ce type d'inférence est appelé en anglais *function argument type inference*. Cependant, il est parfois impossible d'inférer le type de la valeur de retour, comme pour la fonction :

```go
func NewT[T any]() *T {
    ...
}
```

Il faudra alors aider le compilateur en procédant à l'*instanciation* de la fonction avant l'appel :

```go
t := NewT[int]()
```

---
## Quand utiliser les Generics ?

La première recommandation est de ne **jamais définir des contraintes avant d'écrire le code**. Cela peut sembler séduisant d'anticiper et de commencer par définir des contraintes, mais c'est inutile.

![](img/boussole.png)

Le cas d'usage des génériques est de factoriser du **code identique dupliqué avec plusieurs types**. C'est une alternative préférables à l'utilisation du type `interface{}` pour des questions de performance, d'occupation mémoire et de simplicité du code. Le cas d'usage typique est celui des structures de données.

---
## Conclusion

Les génériques sont la grande nouveauté du Go *1.18* qui est la release qui a amené le plus de changements depuis que le Go est Open Source. Cependant, cette fonctionnalité n'a pas été encore assez utilisée en production par un grand nombre d'utilisateurs et doit donc être **utilisée avec précaution**, et bien sûr largement couverte de tests.

![Generics Gopher](img/gopher.png)
