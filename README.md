<p align="center"><img src="https://i.imgur.com/a1H1sQa.png"/></p>

<p align="center">
  <h1 align="center">BITCOIN-TELEGRAM-BOT</h1>
  <p align="center">
    ·
    <a href="https://github.com/tomassirio/bitcoinTelegramBot/issues">Report Bug</a>
    ·
    <a href="https://github.com/tomassirio/bitcoinTelegramBot/issues">Request Feature</a>
    ·
  </p>
</p>

<p align="center">
  <a href="https://github.com/tomassirio/bitcoinTelegramBot/graphs/contributors"><img src="	https://img.shields.io/github/issues/tomassirio/BitcoinTelegramBot"></a>
  <a href="https://github.com/tomassirio/BitcoinTelegramBot/blob/master/LICENSE"><img src="https://img.shields.io/github/license/tomassirio/BitcoinTelegramBot"></a>
  <a href="https://github.com/tomassirio/bitcoinTelegramBot/network/members"><img src="https://img.shields.io/github/forks/tomassirio/BitcoinTelegramBot"></a>
  <a href="https://img.shields.io/github/stars/tomassirio/bitcoinTelegramBot"><img src="https://img.shields.io/github/stars/tomassirio/BitcoinTelegramBot"></a>
</p>

BitcoinTelegram is a simple Telegram Bot to consult BTC's price over Telegram. It's written in Golang with Telebot's Framework

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#tada-inspiration">Inspiration</a>
    </li>
    <li>
      <a href="#star-getting-started">Getting Started</a>
      <ul>
        <li><a href="#what-you-will-need">What you will need</a></li>
        <li><a href="#computer-installation">Installation</a></li>
        <li><a href="#white_check_mark-add-btb-to-your-telegram-channel">Add BTB to your Telegram Channel</a></li>
      </ul>
    </li>
    <li><a href="#battery-usage">Usage</a></li>
    <li><a href="#building_construction-contribution-guidelines">Contribution Guidelines</a></li>
    <li><a href="bust_in_silhouette-who-am-i">Who Am I?</a></li>
  </ol>
</details>

## :tada: Inspiration

Over the last couple of months I grew fond of Golang, Telegram and Bitcoin. So I decided to combine those three elements to create something that could get a little bit of each world. Hence, Bitcoin-Telegram-Bot (in Golang) was created

## :star: Getting started

### What you will need:

- You are going to need a computer or server where to host the bot.
- Git
- Golang v1.13
- A device with Telegram

### :computer: Installation

Open a Terminal and copy these commands (Linux & Mac devices):

```bash
cd ~
git clone https://github.com/tomassirio/BitcoinTelegramBot.git
cd ./BitcoinTelegramBot
mv .env.example .env
go get github.com/tomassirio/bitcoinTelegram
go run main.go
```

##### Warning: 
This won't work unless you replace the **REPLACE_WITH_TOKEN** on the .env file with the Token granted by @BotFather

### :white_check_mark: Add BTB to your Telegram Channel

Open [@BotFather](https://telegram.me/botfather) on telegram and create a new bot with it's __/newbot__ command.

Assign it a name. This name won't be the one that is shown on each message, so you can name it whatever you want.

@BotFather will grant you a Token. This token is the one that will replace the **REPLACE_WITH_TOKEN** on the .env.example file on this repository. (Don't forget to rename that file to .env)

![image](https://i.imgur.com/RC8anHA.png)

You can also play a little bit more with @BotFather. For example you can use the __/setcommands__ to define the uses your bot has on the '/' icon:

```
price - Gets BTC actual price
historic - Gets a percentage between Today's and Yesterday's price
summary - Gets both the price and historic values
```

![image](https://i.imgur.com/ACmSAF1.png)

## :battery: Usage

Once the bot is running and added to your Telegram Group, you can use any of the following commands:

```sh
    * /price : Get's bitcoin's Last price
    * /historic : Gets a percentage between Today's and Yesterday's price
    * /summary : Gets both the price and historic values
```

## :building_construction: Contribution Guidelines:

-   **_fork_** and **_clone_** this repository
-   Make a new branch using `git checkout -b change/username`
-   Commit the desired changes to that branch
-   Sign off your commits using `git commit -s -m w/signoff`
-   Push your changes to the branch and open a pull request
 -->

## :bust_in_silhouette: Who Am I?

<img src="https://media.discordapp.net/attachments/763140054825697301/763681938652528690/logo-design-branding-logo-tool-open-electronic-1-5f7ed02bc8247.png?width=468&height=468" width="410" height="410" /></p>

  <a href="mailto:tomassirio@gmail.com?Subject=Tomas%20You%20Are%20Amazing!">
      <img src="https://cdn2.downdetector.com/static/uploads/logo/image21.png" width="100"; height="100"/>
  </a>
  <a href="https://www.linkedin.com/in/tomassirio/">
      <img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimage.flaticon.com%2Ficons%2Fpng%2F512%2F174%2F174857.png&f=1&nofb=1" width="100"; height="100"/>
  </a>
  <a href="https://dev.to/tomassirio">
      <img src="https://avatars3.githubusercontent.com/u/13521919?s=280&v=4" width="100"; height="100"/>
  </a>
  <a href="https://www.buymeacoffee.com/tomassirio1">
      <img src="https://i.pinimg.com/originals/60/fd/e8/60fde811b6be57094e0abc69d9c2622a.jpg" width="100"; height="100"/>
  </a>