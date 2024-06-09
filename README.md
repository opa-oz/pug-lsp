# pug-lsp: Pug Language Protocol Server

An implementation of the Language Protocol Server for [Pug.js](https://pugjs.org)

## Features

`pug-lsp` aims to provide suggestions for you to edit `.pug` in your editor.

### Note

Project is under heavy development. Current functionality can be change in stable release.

### Tags suggestions

Auto suggest list of HTML5 tags.

![tags-suggestions](docs/tags-suggestions.png)

### Attributes suggestions

#### Auto suggest common attributes (such as `style`, `class`, `title`) for tags

![common-attributes](docs/common-attributes.png)

#### Auto suggest events (such as `onclick`, `onenter`) for tags

![events-attributes](docs/events-attributes.png)

#### Auto suggest tag-specific attributes (such as `href` for `a`)

![special-attributes](docs/special-attributes.png)

#### `&attributes` snippet

_Yes, it's [a real feature](https://pugjs.org/language/attributes.html#attributes) of Pug_

![attributes-shortcut](docs/attributes-shortcut.png)


## Thanks

 - [zealot128/tree-sitter-pug](https://github.com/zealot128/tree-sitter-pug)


----

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S1UZ9P7)

