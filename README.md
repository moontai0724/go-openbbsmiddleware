# go-openbbsmiddleware
go implementation of [openbbs-middleware](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96).

這裡是使用 golang 來達成 [openbbs-middleware](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96).

## Getting Started

You can start with the [swagger api](https://api.devptt.site:5000)
and try the api.

You can copy the curl command from the link if you encounter
CORS issue.

You can `telnet api.devptt.site 8888` and check how the api affects the existing pttbbs (in api.devptt.site)

您可以到 [swagger api](https://api.devptt.site:5000/)
並且試著使用 api.

如果您在 swagger 網頁裡遇到 CORS 的問題. 你可以在網頁裡 copy
curl 指令測試.

您可以使用 `telnet api.devptt.site 8888` 確認 api 如何影響既有的 pttbbs (在 api.devptt.site 裡)

## Docker-compose

You can do the following to start with docker-compose:

* copy `docker_compose.env.template` to `docker_compose.env` and modify the settings.
* `./scripts/docker_initbbs.sh [BBSHOME] pttofficialapps/go-pttbbs:latest`
* `docker-compose --env-file docker_compose.env -f docker-compose.yaml up -d`
* register at `http://localhost:3457/account/register`
* login at `http://localhost:3457/account/login`
* `telnet localhost 8888` and use the account that you registered.

您可以使用以下方式來使用 docker-compose:

* 將 `./docker_compose.env.template` copy 到 `./docker_compose.env` 並且更改 BBSHOME 到您所希望的位置.
* `./scripts/docker_initbbs.sh [BBSHOME] pttofficialapps/go-pttbbs:latest`
* `docker-compose --env-file docker_compose.env -f docker-compose.yaml up -d`
* 在 `http://localhost:3457/account/register` 做 register
* 在 `http://localhost:3457/account/login` 做 login
* `telnet localhost 8888` 並且使用您剛剛登錄的帳號使用.
* 第一次使用時. 須先將 SYSOP 和 pttguest 建立起來.

## Discussing / Reviewing / Questioning the code.

Besides creating issues, you can do the following
to discuss / review / question the code:

* `git clone` the repo
* create a review-[topic] branch
* commenting at the specific code-region.
* pull-request
* start discussion.
* close the pr with comments with only the link of the pr in the code-base.

除了開 issues 以外, 您還可以做以下的事情來對於 code 做討論 / review / 提出問題.

* `git clone` 這個 repo.
* 開一個 review-[topic] 的 branch.
* 對於想要討論的部分在 code 裡寫 comments.
* pull-request
* 對於 PR 進行討論.
* 當 PR 關掉時, comments 會留下關於這個 pr 討論的 link.

## Develop

You can start developing by `git clone` this repository.

您可以使用 `git clone` 來一起開發.

## Unit-Test

You can do unit-test with:

你可以做以下的事情來進行 unit-test:

* `./scripts/test.sh`

You can check coverage with:

您可以做以下的事情來進行 coverage-check:

* `./scripts/coverage.sh`

## Swagger

You can run swagger with:

您可以做以下的事情將 swagger 跑起來:

* `sudo apt install python3 python3-setuptools` - We use python to process some transform, and use `python3-setuptools` for package management.
* `sudo python3 ./apidoc/setup.py install` - Install python packages that we used here.
* `./scripts/swagger.sh [host]`
* go to `http://localhost:5000`

## Schema definition

* `https://github.com/Ptt-official-app/go-openbbsmiddleware/tree/main/schema`
