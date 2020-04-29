
————————————————————————————————————————————————————————————————————————————————
 cc : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————

@Tmp
————————————————————————————————————————————————————————————————————————————————


@Todo
————————————————————————————————————————————————————————————————————————————————

* Create a config structure
* Run the Symfony commands on a given environment
* Create a system > terminal package
* Create the help screen
* When run without argument show the help screen


@Nice to have
————————————————————————————————————————————————————————————————————————————————

* Handle the case where a Symfony command returns an error
* Create an option to display the watched file (+ bonus:: display as a tree)
* Add tests


@Refactoring
————————————————————————————————————————————————————————————————————————————————


@Ideas
————————————————————————————————————————————————————————————————————————————————

* Take the current directoty by default?
* Handling of XML files?
* Handling of the old app/console?


@Naming
————————————————————————————————————————————————————————————————————————————————

* Sous domaine cc.strangebuzz.com ou meme strangebuzz.com/en/cc
* symfony.cc -> déjà pris.
* sfcw


@Resources
————————————————————————————————————————————————————————————————————————————————

* https://www.golangprograms.com/go-language/struct.html
* https://yourbasic.org/golang/for-loop/
* https://golangcode.com/sleeping-with-go/
* https://github.com/spiral/roadrunner Roadrunner contains a file watcher
* https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
* https://blog.golang.org/go1.13-errors


@Done
————————————————————————————————————————————————————————————————————————————————

* Add the .env files to the watched files
* Gather all the yaml files in config (on three levels)
* Build a simple map with the c en clé le chemin et en valeur le timestamp
* Get the update timestamp of a file
* Compute the checksum of a file
* Test if the directory contains bin/console
* Run the Symfony console version command and display the output
* Helper for ansi output
* Tested if directory exists
* Get the full path of the working directory
* Nice ansi output for the welcom message
* Introduced a version constant
