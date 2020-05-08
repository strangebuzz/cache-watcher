# Sfcw

Sfcw stands for SymFony Cache Watcher.

It's a small Go program that will watch your Symfony files and refresh your cache
when needed so you don't have to wait when hitting refreshing the page in the browser.

It's goal is to improve your the DX (Developper Experience) with Symfony.   

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/logos/sfcw.png" alt="The Sfcw mascot" align="right" />

## Installation

You can either build the program youself (this means that you must have a working
Go working developement environment or you can download the executables. 

### Compiling the program

```terminal
$ git clone git@github.com:strangebuzz/sfcw.git
$ cd sfcw
$ make build 
```

This will build the `sfcw` or `sfcw.exe` depenfing on your platform.

### Download

For now, I have built the following executables:

For conveniency, add sfcx in your path so you hace access to it everywhere.

If you need another one, create an issue and point out the operating system/target
plaftorm you need, you will find the possible values in [this article](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures).

## Running the watcher

Now you have built or downloaded the program. If you run it without argument, it
will display the help. If you are at the root of you Symfony application you can
run with the following command:

```terminal
$ sfcw .
```

You should have an output like this:

```terminal
——————————————————————————————————————————————————————————————————————
  Sfcw version v0.3.0 by COil - https://www.strangebuzz.com 🐝
——————————————————————————————————————————————————————————————————————
Sfcw watches your files (.env, YAML, Twig) and automatically refreshes your application cache.
——————————————————————————————————————————————————————————————————————
 > Project directory: /Users/coil/Sites/strangebuzz.com
Custom config file not found.
 > Symfony console path: bin/console
 > Symfony env: Symfony 5.1.0-BETA1 (env: dev, debug: true)
 > 263 file(s) watched in /Users/coil/Sites/strangebuzz.com in 12 millisecond(s).
 > CTRL+C to stop watching or run kill -9 28157.
```

That's it! If you have a Symfony4 or 5 project with Flex it's all that you need.

When a file change will be detected in your services.yaml files for example, you
will have a feedback by the program:

```terminal
⬇ Update detected at 17:09:03 > refreshing cache...
  ✅  Done! in 2.43 second(s).
```

You can also pass a relative path or an absolute path for the first arugument:

```terminal
$ sfcw ../strangebuzz.com
```

```terminal
$ sfcw /Users/coil/Sites/strangebuzz.com 
```

## Stopping the watcher

You can either hit CTRL+C or run the kill command with the PID the program has
displayed in the welcome message:

```terminal
$ sudo kill -9 28157
```

## Configuration

## Todo

- [ ] Apply the Symfony style for the console output
- [ ] Add CI with Github actions
- [ ] Add an option to display the watched file
- [ ] Allow to have an additional whitelist of custom files to watch
- [ ] ...feel free to [create an issue](https://github.com/strangebuzz/sfcw/issues/new) 🙂.

## License

This software is published under the [MIT License](LICENSE.md)


