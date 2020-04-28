
————————————————————————————————————————————————————————————————————————————————
 cc : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————

@Tmp
————————————————————————————————————————————————————————————————————————————————


@Todo
————————————————————————————————————————————————————————————————————————————————

* Récupérer le timestamp d'un fichier yaml
* Construire un tableau simple avec en clé le chemin et en valeur le timestamp
* Lister l'ensemble des fichiers yaml situés dans le répertoire config
* Retourner une erreur si aucun fichier YAML n'est trouvé
* Grab all YAML files in an array
* Add the .env files to the watched files
* Initialize tests
* Help argument


@Refactoring
————————————————————————————————————————————————————————————————————————————————


@Ideas
————————————————————————————————————————————————————————————————————————————————

* idée d'algo:
  - Contruire un tableau avec en clé le chemin du fichier et en valeur la date de modification
  - Si la date de modification est différente alors on lance le cache warmuo

* Take the current directoty by default?
* Handling of XML files?
* Handling of the old app/console?


@Naming
————————————————————————————————————————————————————————————————————————————————

* Sous domaine cc.strangebuzz.com ou meme strangebuzz.com/cc
* symfony.cc -> déjà pris.
* sfcw


@Resources
————————————————————————————————————————————————————————————————————————————————

* https://github.com/spiral/roadrunner Roadrunner contains a file watcher
* https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
* https://blog.golang.org/go1.13-errors


@Done
————————————————————————————————————————————————————————————————————————————————

* Compute the checksum of a file
* Test if the directory contains bin/console
* Run the Symfony console version command and display the output
* Helper for ansi output
* Tested if directory exists
* Get the full path of the working directory
* Nice ansi output for the welcom message
* Introduced a version constant
