# Modelbasierte Software-Entwicklung: P2 Translating “generics”

| Name | Kürzel | Matrikelnummer |
| ---- | ------ | -------------- |
| Marco Janotta | jama1028 | 67886 |
| Leonard Bausenwein | bale1017 | 63554 |

Programmausgabe:
```
=========
EXAMPLE 1
=========

-------------------
Generic translation
-------------------
10
10

----------------
Monomorphization
----------------
10
10


=========
EXAMPLE 2
=========

-------------------
Generic translation
-------------------
Not possible because we don't know the types.
Adding two untyped operators is undefined.

----------------
Monomorphization
----------------
3
3.300000


=========
EXAMPLE 3
=========

-------------------
Generic translation
-------------------
x=1, y=2 -> x=2, y=1

----------------
Monomorphization
----------------
x=1, y=2 -> x=2, y=1
a=true, b=false -> a=false, b=true
```
