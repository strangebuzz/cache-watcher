————————————————————————————————————————————————————————————————————————————————
                    sfcw : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————


@Bugs
————————————————————————————————————————————————————————————————————————————————


@Todo
————————————————————————————————————————————————————————————————————————————————

* Transfert ConfigDirectory in the config structure
* Watch translations
* Create a parameter for the nesting level to apply
* Make a loop for config glob patterns as they are the same

-- V0.2

* Allow to use custom parameters based on the .sfcw.yml file
* Add some tests
* Transform command arguments into the config


@Nice to have
————————————————————————————————————————————————————————————————————————————————

* Create an option to display the watched file (+ bonus: display as a tree)
* Allow to have an additional whitelist of custom files to watch
* Handle the case where a Symfony command returns an error
* Generate a pid file to be able to manually stop the process if something goes wrong.
* Prevent the user from running sfcw for the same project more than one-time


@Refactoring
————————————————————————————————————————————————————————————————————————————————


@Ideas/Decision to take
————————————————————————————————————————————————————————————————————————————————

* Watch Twig files?
* Handling of XML files? ini files?
* Handling of the old app/console?


@Naming
————————————————————————————————————————————————————————————————————————————————

* sfcw
* Not "cc" because it's the C,C++ compiler.


@Resources
————————————————————————————————————————————————————————————————————————————————

* https://golang.org/pkg/fmt/
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

* Compute the time it takes to get the files list
* Renamed cc to sfcw
* Added handling of absolute paths
* Compute the time that takes to refresh the cache
* Show the help screen when no argument is provided
* Added a sleep time parameter
* Create a config structure
* Run the Symfony commands on a given environment
* Run the Symfony commands with the debug mode or not
* Add the .env files to the watched files
* Gather all the yaml files in config (on three levels)
* Build a simple map with the path as the key and the timestamp as the value
* Get the update timestamp of a file
* Compute the checksum of a file
* Test if the directory contains bin/console
* Run the Symfony console version command and display the output
* Helper for ansi output
* Tested if directory exists
* Get the full path of the working directory
* Nice ansi output for the welcome message
* Introduced a version constant
