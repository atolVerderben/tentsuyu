# Tentsuyu

This is an extension on top of [Ebiten](https://hajimehoshi.github.io/ebiten/v2/).
Tentsuyu is a tempura dipping sauce, so I thought the name was appropriate.
This came from another project I was working on and decided to pull ou the portions
that could be reusable for other games using ebiten. This is very much a work in progress and mainly used for my personal projects.

## Features

* Camera
* Input Manager
* HUD
* UI Controller
  * Menus
  * Text
  * Drawn Cursor
  * Very basic "text box"
* Tile Map implementation
  * Reads JSON files from Tiled editor
    * Currently only one format, this needs to be expanded
* Image Manager (track 'textures')
* GameObject interface
  * Basic Object implementation
  * Basic Image Options implementation
* Starter Game Struct
