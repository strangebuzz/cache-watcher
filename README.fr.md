# CacheWatcher

CacheWatcher est un petit programme Go qui observe vos fichiers Symfony et rafraichit
votre cache si besoin afin que vous n'ayez pas à attendre lors du rafraichissement
du navigateur.

Son but est d'améliorer [l'expérience Développeur](https://symfony.com/blog/making-the-symfony-experience-exceptional) avec Symfony (DX).

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/cw_400w.png" alt="La mascotte Cw" align="right" />

## Comment ça marche ? 🤔

Le programme "observe" vos fichiers (.env, YAML, Twig, entités Doctrine) et dès qu'il
détecte un changement, il appelle la commande Symfony `cache:warmup` pour rafraichir
le cache. Il est important de comprendre que le programme ne va ni appeler ni créer
des fichiers sur votre machine par lui-même.

## Installation 🛠️

Vous pouvez compiler le programme manuellement (ça implique que vous ayez un environnement
de travail Go fonctionnel) ou vous pouvez [télécharger un exécutable](#downloading-the-executable-).
Ce programme a été développé avec **go1.14.2**.

### Compilation du programme ⚙️

```
$ git clone git@github.com:strangebuzz/cache-watcher.git
$ cd cache-watcher
$ make build 
```

Cela va construire l'exécutable `cw` or `cw.exe` selon votre système d'exploitation.

### Téléchargement de l'exécutable 🔽

Voici les exécutables des principaux systèmes d'exploitation :

Système d'exploitation | Plateforme | version | fichier     | Contrôle d'intégrité SHA 
---------------------- | ---------- | ------- | ----------- | -------------------
darwin (macOS)         | amd64      | 0.5.0   | [cw](https://sfcw.dev/downloads/darwin/amd64/cw) (3.3mo)          | 80b4011567c71aef7a13e9bbdad9acacee124acb330222fa2b2abf172ce2879e
Linux                  | amd64      | 0.5.0   | [cw](https://sfcw.dev/downloads/linux/amd64/cw) (3.3mo)           | 7203e4facfd82f54501ff17b569055daf7856164ed71a66b3fbfbe8dd9633293
Windows                | amd64      | 0.5.0   | [cw.exe](https://sfcw.dev/downloads/windows/amd64/cw.exe) (3.4mo) | 2278d95d28011cbd2b97aed08e5e73274f08d245092cab58ac7f9d5c772e6892

Une fois téléchargé, vous pouvez vérifier que le fichier n'est pas compromis. Comparez
d'une part, le code de contrôle d'intégrité SHA en exécutant la commande suivante
et d'autre part, la valeur affichée dans le tableau précédent : 

```
$ shasum -a 256 ./cw 
b35078644ac3b3f025276a0c5fcd77b3d2c8fe9cd15d136df969772e6f513973  ./cw
```

Sur Linux et macOS, donnez le droit d'exécution au programme :

```
$ chmod +x ./cw
```

Si vous avez besoin d'un autre type d'exécutable, vous pouvez créer un ticket en
mentionnant le système d'exploitation et plateforme dont vous avez besoin. Vous
trouverez les valeurs possibles dans [cet article](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures).

Par commodité, ajoutez `cw` à votre `path` pour y accéder de n'importe où.

💡 L'exécutable est "assez" gros (plusieurs mo), car il embarque le [run-time Go](https://stackoverflow.com/q/28576173/633864)
et n'a pas de dépendance externe.

## Lancement ⚡

Maintenant que vous avez compilé ou téléchargé le programme, essayons-le. Si vous
le lancez sans arguments, il affiche un message d'aide. Si vous êtes à la racine
de votre application Symfony, vous pouvez commencer à observer vos fichiers avec
la commande suivante : 

```
$ cw .
```

Vous devriez avoir la sortie suivante :

```
——————————————————————————————————————————————————————————————————————
  CacheWatcher version v0.5.0 by COil - https://www.strangebuzz.com 🐝
——————————————————————————————————————————————————————————————————————
CacheWatcher watches your files (.env, YAML, Twig) and automatically refreshes your application cache.
——————————————————————————————————————————————————————————————————————
 > Project directory: /Users/coil/Sites/strangebuzz.com
 > Symfony console path: bin/console
 > Symfony env: Symfony 5.2.0 (env: dev, debug: true)
 > 321 file(s) watched in /Users/coil/Sites/strangebuzz.com in 12 millisecond(s).
 > CTRL+C to stop watching or run kill -9 7817.
```

Et voilà ! Si vous avez un projet Symfony 4 ou 5 avec la structure de répertoire Flex,
c'est tout ce dont vous avez besoin. 

Quand une modification est détectée sur votre fichier `services.yaml` par exemple,
vous avez le retour suivant : 

```
⬇ Update detected at 17:09:03 > refreshing cache...
  ✅  Done! in 2.43 second(s).
```

Maintenant, rafraichissez votre page. Elle devrait être rapide puisque le cache
est déjà prêt : 🐰

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/fast-cache.png" alt="Cache already loaded" align="center" />

Au lieu d'avoir une page "lente" : 🐌

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/slow-cache.png" alt="Cache refreshed by the browser call" align="center" />

💡 Vous pouvez aussi passer un chemin relatif ou absolu comme argument : 

```
$ cw ../strangebuzz.com
$ cw /Users/coil/Sites/strangebuzz.com 
```

Je l'utilise dans le terminal inclus dans PHPStorm : 

<img src="https://raw.githubusercontent.com/strangebuzz/cache-watcher/master/doc/img/cw-phpstorm-terminal.png" alt="Using cw inside a PHPStorn terminal" align="center" />

/‼️\ Attention, si vous fermez la fenêtre PHPStorm, le processus `cw` ne sera pas
automatiquement arrêté. /‼️\

## Arrêt ⛔

Vous pouvez soit utiliser *CTRL+C* ou arrêter le processus manuellement grâce au PID
qui a été affiché dans le message d'accueil :

```
$ kill -9 28157
```

## Configuration 🎛️

Comme nous l'avons vu précédemment, si votre projet a [une structure Flex](https://symfony.com/doc/current/setup/flex.html), 
les paramètres par défaut devraient être bons. Ces paramètres par défaut seront
toujours adaptés à la dernière version mineure de Symfony, actuellement 5.2 :

Clé                 | Valeur par défaut | Description
------------------- | ------------------| -------------------------------------------
console_path        | bin/console       | Chemin relatif vers la console Symfony
env                 | dev               | Correspond au paramètre APP_ENV de l'application Symfony 
debug               | true              | Correspond au paramètre APP_DEBUG de l'application Symfony (true ou false)
config_dir          | config            | Chemin relatif ou sont stockés les fichiers de configuration de l'application Symfony
translations_dir    | translations      | Chemin relatif ou sont stockés les fichiers de traductions de l'application Symfony
templates_dir       | templates         | Chemin relatif ou sont stockés les templates de l'application Symfony
entities_dir        | src/Entity        | Chemin relatif ou sont stockés les entités Doctrine
templates_extension | twig              | Extension par défaut des templates
yaml_extension      | yaml              | Extension par défaut des fichiers YAML, on considère qu'elle est cohérente pour l'ensemble de l'application
sleep_time          | 30                | Pause entre deux analyses du système de fichiers en millisecondes

Si vous n'utilisez pas Flex, vous pouvez mettre un fichier `.cw.yaml` à la racine
de votre projet. Voici la configuration que j'utilise pour un de mes anciens projets
Symfony 4.4 :

```
config_dir:       app/config
translations_dir: src/AppBundle/Resources/translations
templates_dir:    src/AppBundle/Resources/views
yaml_extension:   yml
sleep_time:       30
```

💡 Le temps de pause (sleep_time) est le paramètre en millisecondes entre deux
analyses du système de fichiers. Plus petite est la valeur, plus rapide sera le
rafraichissement du cache, mais plus haute sera l'utilisation du processeur. J'ai
constaté que 30 ms était un bon compromis pour mon Mac Mini 2018 (i7 / 3,2 GHz / 16 go),
mais vous voudrez surement trouver la valeur la plus adaptée à votre système (avec
top ou htop).

## À faire 📋

- [ ] [Ajouter une CI avec les actions GitHub](https://github.com/strangebuzz/cache-watcher/issues/3)
- [ ] [Ajouter une option pour afficher les fichiers surveillés](https://github.com/strangebuzz/cache-watcher/issues/2)
- [ ] [Permettre d'avoir une liste blanche additionnelle de fichiers à observer](https://github.com/strangebuzz/cache-watcher/issues/4)
- [ ] Libre à vous de [créer un ticket](https://github.com/strangebuzz/cache-watcher/issues/new) ➕.

## Notes 📔

Je ne ferai pas de mise à jour en direct comme le binaire Symfony. Merci de surveiller
le dépôt afin d'être notifié de la publication de nouvelles versions.

## Contribuer  🤝

Vous êtes la bienvenue. Mais n'oubliez pas que je veux garder ce programme aussi
léger que possible avec un seul but en tête. Même si ce projet est très récent,
les fonctionnalités principales sont déjà implémentées. 

## Truc marrant (ou pas) 😄

Quand je développais `cw`, j'ai beaucoup joué avec les fichiers de configuration.
Une fois, j'ai modifié un fichier `.env` et il se trouve que quand j'ai rafraîchi
la page, elle était rapide, genre 50ms. J'ai répété l'opération plusieurs fois, 
même résultat ! 🤔 Ça m'a pris quelques instants avant de comprendre qu'un processus
`cw` tournait toujours en tâche de fond. C'est pourquoi je ne pouvais pas constater
un chargement "lent" de ma page. Et voilà, j'avais ma preuve, ça fonctionne ! ™ 😊
 
## Mérites ™

* Symfony ™ est une marque déposée de [Symfony SAS](https://symfony.com/license).
* Logo par [Claire Brunaud](https://clairebrunaud.myportfolio.com).
* Logo Golang original "Gopher" par [Renee French](http://reneefrench.blogspot.com).

## Licence ™

Ce logiciel est publié sous la [licence MIT](LICENSE).

## Remerciements 👏

* [Jonathan Scheiber](https://github.com/jmsche) pour ces nombreuses relectures
  de la documentation et des articles du blog.

## Changelog 📒

### V0.5.0

* Ajout du support pour surveiller les entités Doctrine (utile pour API Platform)

### V0.4.0

* Version initiale
