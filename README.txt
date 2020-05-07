————————————————————————————————————————————————————————————————————————————————
                    Sfcw : The Symfony Cache Watcher
————————————————————————————————————————————————————————————————————————————————


@Bugs
————————————————————————————————————————————————————————————————————————————————


@Todo
————————————————————————————————————————————————————————————————————————————————

* Transfert the repo in sfcw, make it public
* Change the go import


@Questions/Investigations:
————————————————————————————————————————————————————————————————————————————————

* About the main loop: https://stackoverflow.com/q/61657220/633864


@Distribution
————————————————————————————————————————————————————————————————————————————————

* https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
* https://dave.cheney.net/2012/09/08/an-introduction-to-cross-compilation-with-go


@Nice to have
————————————————————————————————————————————————————————————————————————————————

* Apply the Symfony style for the console output (use colored background)
* Add an option to display the watched file (+ bonus: display as a tree)
* Allow to have an additional whitelist of custom files to watch
* Handle the case where a Symfony command returns an error
  > https://stackoverflow.com/q/39508086/633864
* Create a deamon mode: --deamon
  + Generate a pid file to be able to manually stop the process if something goes wrong
* Prevent the user from running sfcw for the same project more than one-time


@Refactoring/Long term
————————————————————————————————————————————————————————————————————————————————


@V2
————————————————————————————————————————————————————————————————————————————————

* Rewrite with https://github.com/fsnotify/fsnotify?


@Ideas/Decision to take
————————————————————————————————————————————————————————————————————————————————

* Create a parameter for the nesting level to apply? Make it generic ?
* Only display the raw go errors with the verbose mode?
* Handling of XML files? ini files?
* Handling of the old app/console?


@Naming
————————————————————————————————————————————————————————————————————————————————

* sfcw
* Not "cc" because it's the C compiler.


@Resources
————————————————————————————————————————————————————————————————————————————————

* https://www.jetbrains.com/help/webstorm/using-file-watchers.html#
* https://golang.org/pkg/time/#Time.Format > Mon Jan 2 15:04:05 -0700 MST 2006
* https://golang.org/pkg/fmt/
* https://stackoverflow.com/q/26804642/633864
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

* Doc to create linux and win exec in Makefile
* Added Symfony5 and 4 fixtures
* Added some tests
* Added the pid number in the welcome message
* Watch Twig files
* Watch translations files
* Transfered the ConfigDirectory constant in the config structure
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
