# CacheWatcher

CacheWatcher est un petit program Go qui surveille vos fichiers Symfony et rafraichit
votre cache si besoin afin que vous n'ayez pas à attendre lors du rafraichissment
du navigateur.

Son but est d'améliorer [l'epérience Développeur](https://symfony.com/blog/making-the-symfony-experience-exceptional) avec Symfony (DX).

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/cw_400w.png" alt="La mascotte Cw" align="right" />

## Comment ça marche ? 🤔

Le programme "surveille" vos fichier (.env, YAML, twig) et dès qu'il détecte un
changement, il appelle la commande Symfony `cache:warmup` pour rafraichir le code.
Il est important de comprendre que le programme ne va pas appeller ni créer des 
fichiers sur votre machine par lui même/

## Installation 🛠️

Vous pouvez compiler le programme manuellent (ça implique que vous ayez un environnement
de travail Go fonctionnel) ou vous pouvez [télécharger un exécutable](#downloading-the-executable-)
Ce programme a été dévelppé avec **go1.14.2**.

### Compilation du programme ⚙️

```
$ git clone git@github.com:strangebuzz/cache-watcher.git
$ cd cache-watcher
$ make build 
```

Cela va construire l'exécutable `cw` or `cw.exe` selon votre système d'exploitation.

### Téléchargement de l'éxécutable 🔽

Voici les exécutables des princpaux systèmes d'exploitation :

Système d'exploitation | Platforme | version | file        | SHA checksum 
---------------------- | --------- | ------- | ----------- | -------------------
darwin (macOS)         | amd64     | 0.4.0   | [cw](https://sfcw.dev/downloads/darwin/amd64/cw) (3.2mo)        | b35078644ac3b3f025276a0c5fcd77b3d2c8fe9cd15d136df969772e6f513973 
linux                  | amd64     | 0.4.0   | [cw](https://sfcw.dev/downloads/linux/amd64/cw) (3.2mo)         | cc5c4b828482db2dd00ae5a566799ff9778de4d48dde520e4cb2e867c7ad4182 
windows                | amd64     | 0.4.0   | [cw.exe](https://sfcw.dev/downloads/windows/amd64/cw.exe) (3.3mo) | d244f9322d2d45b60312fafcb4d2d9499b4632d2a652c38f0d86094af90bfcda 

Une fois téléchargé, vous pouvez vérifier que le fichier n'est pas compromis en
comparant le contrôle d'intégrité SHA en executant la commande suivante et la valeur
affichée avec celle du tableau suivant: 

```
$ shasum -a 256 ./cw 
b35078644ac3b3f025276a0c5fcd77b3d2c8fe9cd15d136df969772e6f513973  ./cw
```

Sur Linux et MacOS, donnez le droit d'exécution au programme :

```
$ chmod +x ./cw
```

Si vous avez besoin d'un autre type d'exécutable, vous pouvez créer une issue
en mentionant le système d'exploitation et platforme dont vous avez besoin.
Vous trouverez les valeur possibles dans [cet article](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures).

Par commodité, ajoutez `cw` à votre `path` pour pour y accéder de n'importe ou.

💡 L'exécutable est "assez" gros (plusieurs mo) car il embarque le [run-time Go](https://stackoverflow.com/q/28576173/633864)
et n'a pas de dépendance externe.

## Lancement ⚡

Maintenant que vous avez compilé ou téléchargé le programme, essayons le. Si vous
le lancez sans arguments, il affichera un message d'aide. Si vous êtes la racine
de votre application Symfony, vous pouvez commencer à surveiller vos fichier avec
la commande suivante : 

```
$ cw .
```

Vous devriez avoir la commande suivante.

```
——————————————————————————————————————————————————————————————————————
  CacheWatcher version v0.4.0 by COil - https://www.strangebuzz.com 🐝
——————————————————————————————————————————————————————————————————————
CacheWatcher watches your files (.env, YAML, Twig) and automatically refreshes your application cache.
——————————————————————————————————————————————————————————————————————
 > Project directory: /Users/coil/Sites/strangebuzz.com
 > Symfony console path: bin/console
 > Symfony env: Symfony 5.1.0-BETA1 (env: dev, debug: true)
 > 263 file(s) watched in /Users/coil/Sites/strangebuzz.com in 12 millisecond(s).
 > CTRL+C to stop watching or run kill -9 28157.
```

Et voilà ! Si vous avez un projet Symfony 4 ou 5 avec la structure de répertoire Flex
c'est tout ce sont vous avez besoin. 

Quand une modification sera détectée sur votre fichier `services.yaml` par exemple,
vous aurez le retour suivante : 

```
⬇ Update detected at 17:09:03 > refreshing cache...
  ✅  Done! in 2.43 second(s).
```

Maintenant rafraichissez votre page. Elle devrait être rapide puisque le cache
est déjà prêt : 🐰

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/fast-cache.png" alt="Cache already loaded" align="center" />

Au lien d'avoir une page "lente" : 🐌

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/slow-cache.png" alt="Cache refreshed by the browser call" align="center" />

💡 Vous pouvez aussi passer un chemin relatif ou absolu comme argument : 

```
$ cw ../strangebuzz.com
$ cw /Users/coil/Sites/strangebuzz.com 
```

Je l'utiliser dans le terminal inclue dans PHPStorm : 

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/cw-phpstorm-terminal.png" alt="Using cw inside a PHPStorn terminal" align="center" />

/‼️\ Attention, si vous fermez la fenêtre PHPStorm, le processus `cw` ne sera pas
automatiquement tué. /‼️\

## Arrêt ⛔

Vous pouvez soit utiliser *CTRL+C* ou tuer le processus manuellement grâce au PID
qui a été affiché dans le message d'accueil :

```
$ sudo kill -9 28157
```

## Configuration 🎛️

Comment nous l'avons py précédemment, si votre projet a [une structure Flex](https://symfony.com/doc/current/setup/flex.html),  
les paramètres par défaut devraient être bons. Ces valeurs par défaut seront toujours
adaptées à la dernière version mineure de Symfony, actuellement 5.1 :

Clé                 | Value par défaut | Description
------------------- | -----------------| -------------------------------------------
console_path        | bin/console      | Chemin relatif vers la console Symfony
env                 | dev              | Correspond au paramètre APP_ENV de l'application Symfony 
debug               | true             | Correspond au paramètre APP_DEBUG de l'application Symfony
config_dir          | config           | Chemin relatif ou sont stockés les fichiers de configuration de l'application Symfony
translations_dir    | translations     | Chemin relatif ou sont stockés les fichiers de traductions de l'application Symfony
templates_dir       | templates        | Chemin relatif ou sont stockés les templates de l'application Symfony
templates_extension | twig             | Extension par défaut des templates
yaml_extension      | yaml             | Extension par défau des fichiers YAML, on considère qu'elle cohérente pour l'ensemble de l'application
sleep_time          | 30               | Pause entre deux analyses du système de fichiers en millisecondes

Si vous n'utiliez pas Flex, vous pouvez mettre un fichier `.cw.yaml` à la racine
de votre projet. Voici la configuration que j'utilise pour l'un de mes "anciens"
project Symfony 4.4 :

```
config_dir:       app/config
translations_dir: src/AppBundle/Resources/translations
templates_dir:    src/AppBundle/Resources/views
yaml_extension:   yml
sleep_time:       30
```

💡 Le temps de pause (sleep_time) est le paramètre en millisecondes entre deux
analyses du système de fichiers. Plus petite est la valeur, plus rapide sera le
rafraichissment du cache, mais plus haute sera l'utilisation du processeur. J'ai
constaté que 30ms était un bon compromise pour mon MacMini 2018 (i7 / 3,2GHz / 16go),
mais vous voudrez surement trouver la valeur la plus adaptée à votre système (avec
top ou hop). 

## À faire 📋

- [ ] [Appliquer le style Symfony pour la sortie console](https://github.com/strangebuzz/cache-watcher/issues/1) 
- [ ] [Ajouter une option pour afficher les fichiers surveillés](https://github.com/strangebuzz/cache-watcher/issues/2)
- [ ] [Ajouter une CI avec les actions Github](https://github.com/strangebuzz/cache-watcher/issues/3)
- [ ] [Permettre d'avoir une liste blanch additionnelle de fichiers à surveiller](https://github.com/strangebuzz/cache-watcher/issues/4)
- [ ] Libre à vous de [créer un ticket](https://github.com/strangebuzz/cache-watcher/issues/new) 🙂.

## Notes 📔

Je ne ferai pas de mise à jour en direct comme le binaire Symfony. Merci de surveiller
le dépôt afin d'être notifié de la publication de nouvelles versions.

## Contribuer  🤝

Vous êtes la bienvenue. Mais n'oubliez pas que je veux garder ce programme aussi
léger que possible avec une fonctionnalité unique. Même si ce projet est très récent
je considère presque que toutes les fonctionnalités ont déjà été implémentées. 

## Truc marrant 😄

Quand je développais `cw`, j'ai beaucoup joué avec les fichier de configuration.
Une fois, j'ai modifié un fichier `.env` et il se trouve que quand j'ai rafraichit
la page elle était rapide, genre 50ms. J'ai répété l'opération plusieurs fois, 
même résultat ! 🤔 Ça m'a pris quelques instants avant de comprendre qu'un processus
tournait toujours en tâche de fond, c'est pourquoi je ne pouvais pas constater
un chargement "lent" de ma page. Et voilà, j'avais ma preuve, ça fonctionne ! ™ 😊
 
## Credits ™

* Symfony ™ est une marque déposée de [Symfony SAS](https://symfony.com/license).
* Logo Golang original "Gopher" par [Renee French](http://reneefrench.blogspot.com).

## License ™

Ce logiciel est publié sans la [license MIT](LICENSE).

## Remerciements 👏

* [Johanatan Scheiber](@jschme) pour ces nombreuses relectures de la documention 
  et des articles du blog.
