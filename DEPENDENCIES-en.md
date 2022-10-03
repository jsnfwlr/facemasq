# Contributing to the project by adding or improving documentation and i18n support
## Documentation
The documentation is currently limited to the development and installation aspects of the project, with no information about the different features. It is also all in English and the only way to access the documentation is via the git repository, but I hope to add an official documentation system with i18n support before v0.2.0 is published.

If you would like to help with expanding the documentation or translating the existing documentation, you can do so using your web-browser:

Just sign up for a github account if you don't already have one, fork the primary repository, use the github user interface to modify the markdown documents, and open a pull-request to get your changes merged in to the primary repository. If you're adding a translation for exisiting documentation, please do so by copying the file you are translating and changing the [ISO 639-1 Language Code](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) in the filename to match the language you are contributing, and then translating the contents of the copied file.

> **Example**: To translate this DEPENDECIES documentation from English to Spanish, copy `DEPENDENCIES-en.md` to `DEPENDENCIES-es.md`

## Internationalisation (i18n)
The faceMasq interface and error/debug/process messages are all currently in English. At present there are no i18n packages in either the Go API or the VueJS web interface, but I hope to add i18n support to both before v0.2.0 is published.

Once i18n is possible, I will gladly accept contributions adding translations.

If you would like to help with expanding the translations of the interface, you can do so using your web-browser:

Just sign up for a github account if you don't already have one, fork the primary repository, use the github user interface to modify the appropriate i18n files, and open a pull-request to get your changes merged in to the primary repository.

# Contributing to the project by adding or improving features and bug fixes
* ### Git
  Git is required to clone/fork the primary code repository, push your changes to your fork, and open pull-requests.

* ### C/C++ tools and libraries
  `libpcap-dev` is used by the core of faceMasq to perform the network scans.
  
  `gcc`, `make`, `flex`, and `bison` are required to build `libpcap` too.

  *Debian/Ubuntu/Mint*
  ```sh
  sudo apt install build-essential bison flex
  ```

  *Red Hat/CentOS/Fedora* **Note:** These commands need [improving](#documentation)
  ```sh
  sudo yum groupinstall "Development Tools" 
  sudo yum install bison flex libpcap-dev 
  ```

  *Arch/Manjaro* **Note:** These commands need [improving](#documentation)
  ```sh
  pacman -S base-devel bison flex libpcap-dev 
  ```

  *Mac OS* **Note:** These commands need [improving](#documentation)
  ```sh
  brew install make gcc libpcap-dev bison flex
  ```
* ### NodeJS and PNPM
  NodeJS and PNPM are used to build the web interface for faceMasq.
  ```sh
  ```
* ### Docker
  Docker is used to run the database containers needed for testing - `mysql:8-debian` and `postgres:14-alpine`.

# Building your own containers

## Make
Make is used to run the makefile commands.

```sh
sudo apt install make # debian/ubuntu/mint
sudo yum install make # redhat/centos/fedora
pacman -S base-devel  # arch/manjaro
brew install make     # mac os
```

## SemVer Bash 
SemVer Bash is used in Makefile for versioning the project.

```sh
wget https://raw.githubusercontent.com/unforswearing/bash-semver/main/semver.bash
chmod +x semver.bash
sudo mv semver.bash /usr/bin/semver
```


## Multi-arch Docker Buildx Config (Optional)
The Buildx Config is used to build a single container image that supports all of the different CPU architectures faceMasq supports (currently `amd64` and  `arm64`/`aarch64`, although I would like to add `armv6` and `armv7` in the future). It is not a requirement if you just want to build or test for the one CPU architecture.
