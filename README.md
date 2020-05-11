# Sfcw: the SymFony Cache Watcher

Sfcw is a small Go program that watches your Symfony files and refreshes your cache
when needed, so you don't have to wait when refreshing your browser.

Its goal is to improve the [Developer Experience](https://symfony.com/blog/making-the-symfony-experience-exceptional) with Symfony (DX).   

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/logos/sfcw_400w.png" alt="The Sfcw mascot" align="right" />

## How does it work? 🤔

The program "watches" your files (.env, YAML, twig) and as soon as it detects a
modification, it will call the Symfony `cache:warmup` command to refresh the cache.
It's important to understand that the program will not create nor delete files on
your filesystem by itself.

## Installation 🛠️

You can build the program yourself (this means that you must have a working
Go development environment) or you can [download an executable](https://github.com/strangebuzz/sfcw#downloading-the-executable-).
The program was developed with **go1.14.2**. 

### Compiling the program ⚙️

```
$ git clone git@github.com:strangebuzz/sfcw.git
$ cd sfcw
$ make build 
```

It will build the `sfcw` or `sfcw.exe` executable depending on your platform.

### Downloading the executable 🔽

Here are the executables for the main operating systems:

Operating System | Platform | file       | SHA checksum 
---------------- | -------- | ---------- | ------------
darwin (macOS)   | amd64    | [sfcw](https://sfcw.dev/downloads/darwin/amd64/sfcw) (3.2mo)        | 6670592d4e6a74ba692bdfd912107cba3a6b9bc3ab1f1139778a16c2b730f2cf
linux            | amd64    | [sfcw](https://sfcw.dev/downloads/linux/amd64/sfcw) (3.2mo)         | 68ee5fbe26835b60e027066602fda079d6d997899dd58d6ffccc80b191a2fb1d
windows          | amd64    | [sfcw.exe](https://sfcw.dev/downloads/windows/amd64/sfcw.exe) (3.3mo) | 59420d2ba7c1df9e6afa6746e1bdc3d197792e2263df3cb857cd65d3e6980011

When downloaded, you can check that the executable isn't compromised by comparing
the SHA checksum you get by running the following command and the value displayed
in the previous table.

```
$ shasum -a 256 ./sfcw 
6670592d4e6a74ba692bdfd912107cba3a6b9bc3ab1f1139778a16c2b730f2cf  ./sfcw
```

If you need another executable type, create an issue and point out the operating
system/target platform you want, you will find the possible values in [this article](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures).

For convenience, add `sfcw` in your path so you can access it from everywhere.

💡 The executable is "quite" big (several mo) because it includes the [Go run-time](https://stackoverflow.com/q/28576173/633864)
and hasn't external dependencies.

## Running the watcher ⚡

Now that you have built or downloaded the program, let's try it. If you run it without
arguments, it will display the help message. If you are at the root of your Symfony
application, you can start to watch your project files with the following command:

```
$ sfcw .
```

You should have the following output:

```
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
it's everything you need.

When a change will be detected in your `services.yaml` file, for example, you will
get the following feedback:

```
⬇ Update detected at 17:09:03 > refreshing cache...
  ✅  Done! in 2.43 second(s).
```

Now refresh your page. It should be "fast" as the cache is already refreshed: 🐰

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/doc/img/fast-cache.png" alt="Cache already loaded" align="center" />

Instead of having a "slow" page: 🐌

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/doc/img/slow-cache2.png" alt="Cache refreshed by the browser call" align="center" />

💡 You can also pass a relative path or an absolute path for the first argument:

```
$ sfcw ../strangebuzz.com
$ sfcw /Users/coil/Sites/strangebuzz.com 
```

I run it in the PHPStorm included terminal:

<img src="https://raw.githubusercontent.com/strangebuzz/sfcw/master/doc/img/sfcw-phpstorm-terminal.png" alt="Using sfcw inside a PHPStorn terminal" align="center" />

/‼️\ Be careful that when closing the PHPStorm window, the `sfcw` process won't be 
automatically be killed. /‼️\

## Stopping the watcher ⛔

You can either hit *CTRL+C* or run the kill command with the PID the program has
displayed in the welcome message:

```
$ sudo kill -9 28157
```

## Configuration 🎛️

As we saw previously, if your project has the [Flex directory structure](https://symfony.com/doc/current/setup/flex.html),
the default settings should be OK. The default values will always be set for the
last minor Symfony version, currently 5.1:

Key                 | Default value | Description
------------------- | --------------| -------------------------------------------
console_path        | bin/console   | Relative path to the Symfony console
env                 | dev           | This is the APP_ENV parameter of the Symfony application
debug               | true          | This is the APP_DEBUG parameter of the Symfony application
config_dir          | config        | Relative directory where are stored the configuration files of the Symfony application
translations_dir    | translations  | Relative directory where are stored the translations files of the Symfony application
templates_dir       | templates     | Relative directory where are stored the templates files of the Symfony application
templates_extension | twig          | Default extension for templates files
yaml_extension      | yaml          | Default extension for YAML files, we consider it must be consistent within an application
sleep_time          | 30            | Sleep time between two filesystem checks in milliseconds

If you are not using Flex, you can put a `.sfcw.yaml` file at the root of your project.
Here is the configuration I use for one of my "old" Symfony 4.4 projects:

```
config_dir:       app/config
translations_dir: src/AppBundle/Resources/translations
templates_dir:    src/AppBundle/Resources/views
yaml_extension:   yml
sleep_time:       30
```

💡 The sleep time is the parameter in milliseconds between two filesystem checks.
The lower it is, the faster the cache will be refreshed, but the higher the CPU
will be used. I found 30ms to be a good value for my MacMini 2018 (i7/3,2GHz/16go),
but you probably want to find the best value for your system (with top or htop).

## Todo 📋

- [ ] [Apply the Symfony style for the console output](https://github.com/strangebuzz/sfcw/issues/1) 
- [ ] [Add an option to display the watched file](https://github.com/strangebuzz/sfcw/issues/2)
- [ ] [Add CI with Github actions](https://github.com/strangebuzz/sfcw/issues/3)
- [ ] [Allow having an additional whitelist of custom files to watch](https://github.com/strangebuzz/sfcw/issues/4)
- [ ] Feel free to [create an issue](https://github.com/strangebuzz/sfcw/issues/new) 🙂.

## Notes 📔

I won't do a live update like the Symfony binary. Please watch the repository to
be notified of new releases.

## Contributing 🤝

You are welcome. But don't forget that I want to keep this program very light
with a unique feature. Even it's very young; it's almost "feature complete".

## Fun fact 😄

When I developed `sfcw`, I played a lot with configuration files. One time, I modified
my `.env` file. And it turns out that when I refreshed the browser, the page was
"fast", like 50ms. I repeated the process several times, the same result! 🤔 It took
me some time to realize that an `sfcw` process was still running in the background! 
That's why I couldn't see the "slow" timer. That was it; I had my proof; it works!
™ 😊

## License ™️

This software is published under the [MIT License](LICENSE).
