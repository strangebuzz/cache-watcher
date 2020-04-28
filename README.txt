
————————————————————————————————————————————————————————————————————————————————
 cc : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————

@Tmp
————————————————————————————————————————————————————————————————————————————————


@Todo
————————————————————————————————————————————————————————————————————————————————

* Test if the directory contains bin/console
* Run the Symfony console and display the output
* Display the Symfony version
* Lister l'ensemble des fichiers yaml situés dans le répertoire config
* Retourner une erreur si aucun fichier YAML n'est trouvé
* Grab all YAML files in an array
* Compute the checksum of a file
* Add the .env files to the watched files
* Initialize tests


@Refactoring
————————————————————————————————————————————————————————————————————————————————



@Ideas
————————————————————————————————————————————————————————————————————————————————

* Take the current directoty by default?
* Handling of XML files?
* Handling of the old app/console?


@Naming
————————————————————————————————————————————————————————————————————————————————

* symfony.cc -> déjà pris.
* sfcw


@Resources
————————————————————————————————————————————————————————————————————————————————

* https://blog.golang.org/go1.13-errors


@Done
————————————————————————————————————————————————————————————————————————————————

* Helper for ansi output
* Tested if directory exists
* Get the full path of the working directory
* Nice ansi output for the welcom message
* Introduced a version constant
