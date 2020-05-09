# Sfcw: the SymFony Cache Watcher

It's a small Go program that will watch your Symfony files and refresh your cache
when needed so you don't have to wait when refreshing your browser.

It's goal is to improve the DX > [Developper Experience](https://symfony.com/blog/making-the-symfony-experience-exceptional) with Symfony.   

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/logos/sfcw_400w.png" alt="The Sfcw mascot" align="right" />

## How does it work?

It's important to understand that the program will not create nor delete files
on your filesystem. In fact, it will just run the Symfony command `cache:warmup` 
for you.

## Installation

You can either build the program yourself (this means that you must have a working
Go developement environment or you can download an executable. 

### Compiling the program

```terminal
$ git clone git@github.com:strangebuzz/sfcw.git
$ cd sfcw
$ make build 
```

This will build the `sfcw` or `sfcw.exe` executable depending on your platform.

### Downloading the exexcutable

For now, I have built the following executables for the main operating systems:

Operating System | Platform | file       | SHA256 Checksum 
---------------- | -------- | ---------- | ---------------
darwin (macOS)   | amd64    | [sfcw](https://sfcw.dev/downloads/darwin/amd64/sfcw) (3.2mo)        | 6670592d4e6a74ba692bdfd912107cba3a6b9bc3ab1f1139778a16c2b730f2cf
linux            | amd64    | [sfcw](https://sfcw.dev/downloads/linux/amd64/sfcw) (3.2mo)         | 68ee5fbe26835b60e027066602fda079d6d997899dd58d6ffccc80b191a2fb1d
windows          | amd64    | [sfcw.exe](https://sfcw.dev/downloads/linux/amd64/sfcw.exe) (3.3mo) | 59420d2ba7c1df9e6afa6746e1bdc3d197792e2263df3cb857cd65d3e6980011

When downloaded you can check than the executalbe is not compromised by comparing
the sha1 you get by running the following and comparing the value with the one displayed
in the previous table.

```terminal
$ shasum -a 256 ./sfcw 
```

If you need another one excutable type, create an issue and point out the operating
system/target plaftorm you need, you will find the possible values in [this article](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures).

For conveniency, add `sfcw` in your path so you can access it everywhere.

## Running the watcher

Now you have built or downloaded the program. If you run it without argument, it
will display the help. If you are at the root of your Symfony application you can
start to watch your project files with he following command:

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
 > Symfony console path: bin/console
 > Symfony env: Symfony 5.1.0-BETA1 (env: dev, debug: true)
 > 263 file(s) watched in /Users/coil/Sites/strangebuzz.com in 12 millisecond(s).
 > CTRL+C to stop watching or run kill -9 28157.
```

That's it! If you have a Symfony 4 or 5 project with the Flex directory structure,
it's all that you need.

When a change will be detected in your `services.yaml` file for example, you will
get the following feedback:

```terminal
⬇ Update detected at 17:09:03 > refreshing cache...
  ✅  Done! in 2.43 second(s).
```

Now refresh you page. It should be "fast" as the cache is already refreshed. 🎉

Instead of having a "slow" page:

You can also pass a relative path or an absolute path for the first arugument:

```terminal
$ sfcw ../strangebuzz.com
```

```terminal
$ sfcw /Users/coil/Sites/strangebuzz.com 
```

I personnaly run it in the PHPStorm included terminal:

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/doc/img/sfcw-phpstorm-terminal.png" alt="Using sfcw inside a PHPStorn terminal" align="center" />

When closing the PHPStorm window, the `sfcw` process will automatically be killed.

## Stopping the watcher

You can either hit CTRL+C or run the kill command with the PID the program has displayed
in the welcome message:

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


