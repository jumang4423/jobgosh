<!-- DO NOT REMOVE - contributor_list:data:start:["jumang4423", "OdriMezanny", "ranon-rat"]:end -->

![README LOGO](_img/bk.png)
# jobgosh | job management tool made with golang for shell

[![Contributor List](https://github.com/jumang4423/jobgosh/actions/workflows/contributor_list.yml/badge.svg)](https://github.com/jumang4423/jobgosh/actions/workflows/contributor_list.yml)

a multi-platform work time management CLI tool to track and improve your day to day workflow

<!-- prettier-ignore-start -->
<!-- DO NOT REMOVE - contributor_list:start -->
## 👥 Contributors


- **[@jumang4423](https://github.com/jumang4423)**

- **[@OdriMezanny](https://github.com/OdriMezanny)**

- **[@ranon-rat](https://github.com/ranon-rat)**

<!-- DO NOT REMOVE - contributor_list:end -->
<!-- prettier-ignore-end -->


## dependancies for build

    - go 
        macOS, brew install go
        debian, sudo apt install go -y
        arch, sudo pacman -S go


## installation

run this command below:

```bash
chmod +x scripts/install.sh
./scripts/install.sh
```
default, this script add path **both bash/zsh & fish**

## usage

- ```jobgosh``` 

    displays help for this tool

- ```jobgosh -times all``` 

    see how long you spent your time for each group

- ```jobgosh -from [YYYY/MM/DD] -to [YYYY/MM/DD]``` 

    more specific option of -t all

    you can choose the duration of time

- ```jobgosh -work [up | down]```

    the current directory will be assigned as work space

    ```up``` to start work

    ```down``` to finish work

    listen, you need to ```down``` unless you want waste your whole work time


## doesn't work? 
try importing them into each shell settings:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.jobgosh" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/jobgosh.fish```

```bash
set PATH ~/.jobgosh : "$PATH" 
```

## development with a docker

### 1. turn true the DOCKER_DEVELOPMENT variable
- in the main.go at var section, there is a variable called ```DOCKER_DEVELOPMENT```
- when its true, u can develop with docker

### 2. launch the docker

```bash
# build go image
sudo docker-compose build
# launch image on background process
sudo docker-compose up -d
# exec go to command using docker envinroment
sudo docker-compose exec jobgosh go run *.go
```


## ERROR MEMOS

- #01
    -t !== month, year or all
- #02
    -w !== up or down
- #03
    log directory error
- #04
    group json not found