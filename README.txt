
————————————————————————————————————————————————————————————————————————————————
 cc : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————

@Tmp
————————————————————————————————————————————————————————————————————————————————


@Todo
————————————————————————————————————————————————————————————————————————————————

* Tester si le répertoire semble bien contenir un projet Symfony (bin/console)
* Lister l'ensemble des fichiers yaml situés dans le répertoire config
* Retourner une erreur si aucun fichier YAML n'est trouvé
* Sotocker l'ensemble des fichiers YAML dans un tableau
* Afficher le contenu d'une fichier
* Add the .env files to the watched files
* Rendre le répertoire à analyser configurable (.cc.dist)
* Initialiser des tests


@Refactoring
————————————————————————————————————————————————————————————————————————————————

* getSymfonyProjectDir must return an error


@Ideas
————————————————————————————————————————————————————————————————————————————————


@Naming
————————————————————————————————————————————————————————————————————————————————

* symfony.cc -> déjà pris.
* sfcw


@Resources
————————————————————————————————————————————————————————————————————————————————

* https://blog.golang.org/go1.13-errors


@Done
————————————————————————————————————————————————————————————————————————————————

* Tested if directory exists
* Get the full path of the working directory
* Nice ansi output for the welcom message
* Introduced a version constant
