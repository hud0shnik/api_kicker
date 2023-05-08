# 👁️ Api_kicker (не актуально) ⚙️

[![License - BSD 3-Clause](https://img.shields.io/static/v1?label=License&message=BSD+3-Clause&color=%239a68af&style=for-the-badge)](/LICENSE)

<h3 align="left">🛠 Стек технологий:</h3>

<!-- Telegram -->
<a href="https://telegram.org/" target="_blank">
<img src="https://img.icons8.com/color/48/000000/telegram-app--v3.png" alt="telegram" width="40" height="40"/></a>
<!-- Golang -->
<a href="https://golang.org" target="_blank"> 
<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go lang" width="40" height="40"/></a>
<!-- Visual Studio Code -->
<a href="https://code.visualstudio.com/" target="_blank">
<img src="https://img.icons8.com/fluent/48/000000/visual-studio-code-2019.png" alt="vs code" width="40" height="40"/></a>
<!-- Heroku -->
<a href="https://www.heroku.com/" target="_blank"><img src="https://img.icons8.com/color/48/000000/heroku.png" alt="heroku" width="40" height="40"/></a>

<h3 align="left">📄 О самом проекте:</h3>
Веб приложение, которое пинает мои апихи каждый час (проверяя работоспособность). В случае нахождения ошибки, уведомляет меня через телеграм бота (метод POST к апи телеграма). Также само работает как апи для получения статуса других апих.

<b>Потеряло свою актуальность в связи с переходом на vercel в ноябре 2022, т.е больше не работает.</b>

<br/><br/>
Реквест:
<br/><br/>

``` Elixir
GET https://api-kicker.herokuapp.com/status
```


Респонс:

``` JSON
{
  "github_api": "Ok",
  "osu_api": "Ok"
}
```
