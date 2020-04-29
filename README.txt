
————————————————————————————————————————————————————————————————————————————————
 cc : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————

@Tmp
————————————————————————————————————————————————————————————————————————————————


@Bugs
————————————————————————————————————————————————————————————————————————————————

* The main directory argument doesn't handle absolute paths
* If we pass --help without other argument, we shoud't consider it as the directory
  argument.


@Todo
————————————————————————————————————————————————————————————————————————————————

* Compute the time that takes the refresh
* Renommer cc en sfcw car cc est le compileur c, c++ ! Donc ça ne va pas passer !! lol
* Make a loop for config glob patterns as they are the same

-- V0.2

* Allow to use custom parameters based on the .cc.yml file
* Add some tests


@Nice to have
————————————————————————————————————————————————————————————————————————————————

* Display the time it took to warmup the cache
* Introduce a verbose mode
* Handle the case where a Symfony command returns an error
* Create an option to display the watched file (+ bonus:: display as a tree)
* Prevent the user from running cc for the same project more than one-time
* Generate a pid file in case to be able to manually stop the process if something
  goes wrong.
* Allow to have an additional whitelist of custom files to watch


@Refactoring
————————————————————————————————————————————————————————————————————————————————

* Make a loop for the three patterns in /config


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

* https://golang.org/doc/effective_go.html
* https://www.golangprograms.com/go-language/struct.html
* https://golangbot.com/structs/
* https://yourbasic.org/golang/for-loop/
* https://golangcode.com/sleeping-with-go/
* https://github.com/spiral/roadrunner Roadrunner contains a file watcher
* https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
* https://blog.golang.org/go1.13-errors


@Done
————————————————————————————————————————————————————————————————————————————————

* When run without argument show the help screen
* Added a sleep time parameter
* Create a config structure
* Run the Symfony commands on a given environment
* Run the Symfony commands with the debug mode or not
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
