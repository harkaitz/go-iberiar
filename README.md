# IBERIAR.EU

Iberian script transcription program and website. You can
visit it on https://iberiar.eu

## Development

- Before editting the text, install the iberian font
  from [here](./cmd/iberiar/assets/iberian.ttf).
- You need to setup a Go programming environment.

## Testing

Simply run this commands:

    > make
    > ./bin/iberiar
    > firefox http://127.0.0.1:8084

## Todo

- Expose all the sillabaries to the website.
- Add more SVG templates.
- Complete the mapping tables.
- Add references, tables etc to the website.
- A lot.

## Iberian fonts

Copy `html/assets/iberian.ttf` to `/usr/share/fonts/TTF/iberian.ttf`
and execute `fc-cache -fv` to read the source files properly.

You can run the installation scripts:

    sudo ./pkg/i-font-iberian install
    sudo ./pkg/i-font-uniedit-iberian install

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
