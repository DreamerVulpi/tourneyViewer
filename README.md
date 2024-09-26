# TourneyViewer

[![Go Report Card](https://goreportcard.com/badge/github.com/DreamerVulpi/tourneyViewer)](https://goreportcard.com/report/github.com/DreamerVulpi/tourneyViewer) [![Russia](https://upload.wikimedia.org/wikipedia/en/thumb/f/f3/Flag_of_Russia.svg/28px-Flag_of_Russia.svg.png)](#русский)

## English

<img style="padding: 10px" align="right" alt="TourneyBot logo" src="https://i.imgur.com/ny2A2WJ.png" width="250">

TourneyViewer is a project for tournament organizers on the [startgg](https://www.start.gg/) platform that allows you to run a widget to display the actual tournament grid live (for example, using [OBS](https://obsproject.com/)).

Using the open API [startgg](https://www.start.gg/), the program receives data about the tournament grid of the group stage, processes the data and displays it on the page located on the local host of the program.

If you want to help the project, suggest ideas and developments in your [pull requests](https://github.com/DreamerVulpi/tourneyViewer/pulls).  

<br>
<br>

## To be realized in the future

* the normal view of the page;

## Features

* Single and double elimination tournament formats are supported;
* Update widget every 1 minute;
  
## Getting Started

### Installing

0. You need to get:
   * developer token for startgg; [How get?](https://developer.start.gg/docs/authentication/)
   * link to the phaseGroup(bracket) of your tournament;
    ```For example: https://www.start.gg/tournament/wild-hunters-1/event/main-online-crossplatform-event/brackets/1724283/2562131```
1. Download the finished build and create a ```config``` folder in the directory.
    * Create file ```config.toml```
    * Copy the template and fill the previously created file with it:

    ```toml
    link = "your link"
    token = "your token"
    ```

### Usage

0. Change file ```config/viewer.css``` to custumization bracket;
1. Start the project;
2. Go to the website to check ```localhost:7777```
3. Use the widget anywhere. Enjoy the process!

## Русский

<img style="padding: 10px" align="right" alt="TourneyBot logo" src="https://i.imgur.com/ny2A2WJ.png" width="250">

TourneyViewer - это проект для организаторов турниров на платформе [startgg](https://www.start.gg/), который позволяет запустить виджет для отображения актуальной турнирной сетки в прямом эфире (например, с помощью [OBS](https://obsproject.com/)).

Используя открытый API [startgg](https://www.start.gg/), программа получает данные о турнирной сетке группового этапа, обрабатывает их и отображает на странице, расположенной на локальном хосте программы.

Если вы хотите помочь проекту, предлагайте идеи и разработки в [pull requests](https://github.com/DreamerVulpi/tourneyViewer/pulls).    

<br>

## Будет реализовано в будущем

* нормальный вид страницы;
  
## Особенности

* Поддержка форматов Single and double elimination;
* Обновление виджета раз в 1 минуту;

## Начало работы

### Установка

0. Вам необходимо получить:
   * токен разработчика для startgg; [Как получить?](https://developer.start.gg/docs/authentication/)
   * ссылку на турнирную сетку вашего турнира;
    ``Например: https://www.start.gg/tournament/wild-hunters-1/event/main-online-crossplatform-event/brackets/1724283/2562131``.
1. Скачайте готовую сборку и создайте папку ``config'' в директории.
    * Создайте файл ``config.toml''
    * Скопируйте шаблон и заполните им ранее созданный файл:

    ```toml
    link = "ваша ссылка"
    token = "ваш токен"
    ```

### Использование

0. Измените файл ```config/viewer.css``` для кастомизации ветки;
1. Запустите проект;
2. Перейдите на сайт, чтобы проверить ```localhost:7777```.
3. Используйте виджет в любом месте. Наслаждайтесь процессом!
