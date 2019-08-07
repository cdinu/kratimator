# Kratimator

The goal of this project is to find and safely replace the dash in Romanian words with the non-breaking dash.
It is one of my first GoLang projects and also it was built under pressure, so bear with the code clumsiness.

🇷🇴 Scopul acestui proiect este să înlocuiască, într-un mod cât mai sigur, cratima-minus cu cratima-care-nu-desparte-cuvintele (aka non-breaking hyphen). Motivul principal este acela că în mod implicit HTML desparte linia după minus, iar acest lucru nu este corect în scrierea românească. Vizual diferențele sunt minore (minus: `-`, non-breaking-hyphen: `‑`)

Exemple: "să-l" devine "să-l"

Abordarea este ușor naivă, dar prinde majoritatea cazurilor. Sunt trei categorii de înlocuiri:

- cuvinte de sine stătătoare: _într-o_, _mi-au_, etc.
- terminații: _-mi-l_ (_dă-mi-l_)
- excepții: _a VII-a_, 'mă-sa'

## Installation

If you are lazy, go to the [releases page](https://github.com/cdinu/kratimator/releases) and download the binary.

Otherwise, setup a Go dev environment, clone the source, do a `go build`, `go install` and enjoy.


## Usage

After having the executable in your path, use this utility as a command line to tool.

Test it with `kratimator -h`.

By default it does not overwrite files, but adds a funny extension in the same path. There is an option to overwrite the files, but you have to discover it for yourself under the `-h` flag. With power comes responsibility.

Glob patterns work very well with zsh (maybe with other shells, but I haven't tested yet). Something like: 

```
kratimator ~/some/random/path/**/*.jade
```


## License and Contributing

This software is licensed under [GPLv3](https://www.gnu.org/licenses/gpl-3.0.html) License.

Feel free to submit PRs or Issues.


