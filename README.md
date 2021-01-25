<p align="center"><img src="https://i.imgur.com/a1H1sQa.png"/></p>

<h1 align="center">BITCOIN-TELEGRAM-BOT</h1>

<p align="center">
  <a href="https://github.com/tomassirio/bitcoinTelegramBot/graphs/contributors"><img src="	https://img.shields.io/github/issues/tomassirio/BitcoinTelegramBot"></a>
  <a href="https://github.com/tomassirio/bitcoinTelegramBot/issues"><img src="https://img.shields.io/github/issues/tomassirio/bitcoinTelegramBot"></a>
  <a href="https://github.com/tomassirio/bitcoinTelegramBot/network/members"><img src="https://img.shields.io/github/forks/tomassirio/bitcoinTelegram"></a>
  <a href="https://img.shields.io/github/stars/tomassirio/bitcoinTelegramBot"><img src="https://img.shields.io/github/stars/tomassirio/bitcoinTelegbitcoinTelegramBotram"></a>
</p>

BitcoinTelegram is a simple Telegram Bot to consult BTC's price over Telegram. It's written in Golang with Telebot's Framework

<!-- ![image](https://miro.medium.com/max/8512/0*1YAdWi5ruRiSQDas) -->

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

You can also play a little bit more with @BotFather. For example you can use the __/setcommands__ to define the uses your bot has on the '/' icon:

```
price - Gets BTC actual price
```

## :battery: Usage

Once the bot is running and added to your Telegram Group, you can use any of the following commands:

```sh
    * /price : Get's bitcoin's Last price
    # * $multi-add <element> <element> - adds mutiple elements \in a list
    # * $help: shows you a message with the available commands
    # * $list: lists all the components \in the channel\'s list
    # * $log: a log of the versions
    # * $poll <active_time_in_minutes>: creates a poll on 5 random items of the list. If attribute is not supplied the poll has no limitation of time.
    # * $random: gives you a random component from the list
    # * $remove <component>: removes the desired component
    # * $multi-remove <element> <element> - removes multiple elements \in a list
    # * $remind <time_in_minutes> <component>: adds a component to the list and reminds you of it in n minutes
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
  <a href="https://discord.io/siriobots">
      <img src="https://www.net-aware.org.uk/siteassets/images-and-icons/application-icons/app-icons-discord.png?w=585&scale=down" width="100"; height="100"/>
  </a>
  <a href="https://www.buymeacoffee.com/tomassirio1">
      <img src="https://i.pinimg.com/originals/60/fd/e8/60fde811b6be57094e0abc69d9c2622a.jpg" width="100"; height="100"/>
  </a>