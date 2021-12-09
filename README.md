# CacheWatcher

CacheWatcher is a small Go program that watches your Symfony files and refreshes
your cache when needed, so you don't have to wait when refreshing your browser.

Its goal is to improve the [Developer Experience](https://symfony.com/blog/making-the-symfony-experience-exceptional) with Symfony (DX).   

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/cw_400w.png" alt="The Cw mascot" align="right" />

## How does it work? 🤔

The program "watches" your files (.env, YAML, Twig, Doctrine entities) and as soon
as it detects a modification, it calls the Symfony `cache:warmup` command to refresh
the cache. It's important to understand that the program will not create nor delete
files on your filesystem by itself.

## Installation 🛠️

You can build the program yourself (this means that you must have a working Go development
environment) or you can [download an executable](#downloading-the-executable-).
The program was developed with **go1.14.2**. 

### Compiling the program ⚙️

```
$ git clone git@github.com:strangebuzz/cache-watcher.git
$ cd cache-watcher
$ make build 
```

It will build the `cw` or `cw.exe` executable depending on your platform.

### Downloading the executable 🔽

Here are the executables for the main operating systems:

| Operating System | Platform | version | date       | file                                                              | SHA checksum                                                     |
|------------------|----------|---------|------------|-------------------------------------------------------------------|------------------------------------------------------------------|
| darwin (macOS)   | amd64    | 0.5.0   | 2020-12-05 | [cw](https://sfcw.ovh/downloads/darwin/amd64/cw) (3.3mo)          | 80b4011567c71aef7a13e9bbdad9acacee124acb330222fa2b2abf172ce2879e |
| Linux            | amd64    | 0.5.0   | 2020-12-05 | [cw](https://sfcw.ovh/downloads/linux/amd64/cw) (3.3mo)           | 7203e4facfd82f54501ff17b569055daf7856164ed71a66b3fbfbe8dd9633293 |
| Windows          | amd64    | 0.5.0   | 2020-12-05 | [cw.exe](https://sfcw.ovh/downloads/windows/amd64/cw.exe) (3.4mo) | 2278d95d28011cbd2b97aed08e5e73274f08d245092cab58ac7f9d5c772e6892 |

When downloaded, you can check that the executable isn't compromised by comparing
the SHA checksum you get by running the following command and the value displayed
in the previous table.

```
$ shasum -a 256 ./cw 
b35078644ac3b3f025276a0c5fcd77b3d2c8fe9cd15d136df969772e6f513973  ./cw
```

On Linux and macOS, give the executable permission to the file:

```
$ chmod +x ./cw
```

If you need another executable type, create an issue and point out the operating
system/target platform you want, you will find the possible values in [this article](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures).

For convenience, add `cw` in your `path` so you can access it from everywhere.

💡 The executable is "quite" big (several mo) because it includes the [Go run-time](https://stackoverflow.com/q/28576173/633864)
and doesn't have external dependencies.

## Running the watcher ⚡

Now that you have built or downloaded the program, let's try it. If you run it without
arguments, it will display the help message. If you are at the root of your Symfony
application, you can start to watch your project files with the following command:

```
$ cw .
```

You should have the following output:

```
——————————————————————————————————————————————————————————————————————
  CacheWatcher version v0.5.0 by COil - https://www.strangebuzz.com 🐝
——————————————————————————————————————————————————————————————————————
CacheWatcher watches your files (.env, YAML, Twig) and automatically refreshes your application cache.
——————————————————————————————————————————————————————————————————————
 > Project directory: /Users/coil/Sites/strangebuzz.com
 > Symfony console path: bin/console
 > Symfony env: Symfony 5.4.0 (env: dev, debug: true)
 > 321 file(s) watched in /Users/coil/Sites/strangebuzz.com in 12 millisecond(s).
 > CTRL+C to stop watching or run kill -9 7817.
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

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/fast-cache.png" alt="Cache already loaded" align="center" />

Instead of having a "slow" page: 🐌

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/slow-cache.png" alt="Cache refreshed by the browser call" align="center" />

💡 You can also pass a relative path or an absolute path for the argument:

```
$ cw ../strangebuzz.com
$ cw /Users/coil/Sites/strangebuzz.com 
```

I run it in the PHPStorm included terminal:

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/cw-phpstorm-terminal.png" alt="Using cw inside a PHPStorn terminal" align="center" />

/‼️\ Be careful that when closing the PHPStorm window, the `cw` process won't be 
automatically be killed. /‼️\

## Stopping the watcher ⛔

You can either press *CTRL+C* or run the kill command with the PID the program has
displayed in the welcome message:

```
$ kill -9 28157
```

## Configuration 🎛️

As we saw previously, if your project has the [Flex directory structure](https://symfony.com/doc/current/setup/flex.html),
the default settings should be OK. The default values will always be set for the
last minor Symfony version, currently 5.4:

| Key                 | Default value | Description                                                                                |
|---------------------|---------------|--------------------------------------------------------------------------------------------|
| console_path        | bin/console   | Relative path to the Symfony console                                                       |
| env                 | dev           | This is the APP_ENV parameter                                                              |
| debug               | true          | This is the APP_DEBUG parameter (true or false)                                            |
| config_dir          | config        | Relative directory where are stored the configuration files                                |
| translations_dir    | translations  | Relative directory where are stored the translations files                                 |
| templates_dir       | templates     | Relative directory where are stored the templates files                                    |
| entities_dir        | src/Entity    | Relative directory where are stored the Doctrine entities                                  |
| templates_extension | twig          | Default extension for templates files                                                      |
| yaml_extension      | yaml          | Default extension for YAML files, we consider it must be consistent within the application |
| php_extension       | php           | Default extension for PHP files, we consider it must be consistent within the application  |
| sleep_time          | 30            | Sleep time between two filesystem checks in milliseconds                                   |

If you are not using Flex, you can put a `.cw.yaml` file at the root of your project.
Here is the configuration I use for one of my "old" Symfony 4.4 projects:

```
config_dir:       app/config
translations_dir: src/AppBundle/Resources/translations
templates_dir:    src/AppBundle/Resources/views
entities_dir:     src/AppBundle/Entity
yaml_extension:   yml
sleep_time:       30
```

💡 The sleep time is the parameter in milliseconds between two filesystem checks.
The lower it is, the faster the cache will be refreshed, but the higher the CPU
will be used. I found 30 ms to be a good value for my MacMini 2018 (i7/3,2GHz/16go),
but you probably want to find the best value for your system (with top or htop).

## Todo 📋

- [ ] [Add CI with GitHub actions](https://github.com/strangebuzz/cache-watcher/issues/3)
- [ ] [Add an option to display the watched file](https://github.com/strangebuzz/cache-watcher/issues/2)
- [ ] [Allow having an additional whitelist of custom files to watch](https://github.com/strangebuzz/cache-watcher/issues/4)
- [ ] Feel free to [create an issue](https://github.com/strangebuzz/cache-watcher/issues/new) ➕.

## Contributing 🤝

You are welcome. But don't forget that I want to keep this program very light
with a unique feature. Even it's very young; it's almost "feature complete".

## Fun fact (or not) 😄

When I developed `cw`, I played a lot with configuration files. One time, I modified
my `.env` file. And it turns out that when I refreshed the browser, the page was
"fast", like 50ms. I repeated the process several times, the same result! 🤔 It took
me some time to realize that an `cw` process was still running in the background! 
That's why I couldn't see the "slow" timer. That was it; I had my proof; it works!
™ 😊

## Credits ™

* Symfony ™ is a trademark of [Symfony SAS](https://symfony.com/license).
* Logo by [Claire Brunaud](https://clairebrunaud.myportfolio.com).
* Original Golang logo "Gopher" by [Renee French](http://reneefrench.blogspot.com).

## License ™

This software is published under the [MIT License](LICENSE).

## Thanks 👏

* [Jonathan Scheiber](https://github.com/jmsche) for his many documentation and blog posts proofreadings.

## Changelog 📒

### V0.5.0

* Added support to watch Doctrine entities (useful for API Platform)

### V0.4.0

* Initial version
