# Switchboard

Switchboard is an articulation set generator for Logic Pro. Instead of mapping articulations to keyswitches using Logic's Articulation Set Editor, articulation mappings are defined using simple JSON files. Switchboard uses these articulation mappings to generate articulation sets on the fly.

## Installation

Switchboard is preloaded with articulation mappings for several popular sample libraries including BBC Symphony Orchestra and Spitfire Symphony Orchestra. To install, download the latest release and place the extracted articulation sets in _User > Music > Audio Music Apps > Articulation Settings_.

## Usage

Switchboard searches the `library` directory for articulation mappings. Each file corresponds with an individual patch provided by a sample library. For example, `Horns a4.json` would correspond with a "Horns a4" patch, and the generated articulation set would be written to `Horns a4.plist`.

An articulation mapping is an array of objects, where `name` is the name of the articulation as seen in the plugin, and `keyswitch` is the configured note. Here's a simple example:

```json
[
  {
    "name": "Legato",
    "keyswitch": {
      "note": "C",
      "octave": -2
    }
  }
]
```

Because the format is non-proprietary the articulation mappings can feasibly be used in support of other digital audio workstations, in developing control surfaces, etc.

> [!IMPORTANT]
> Switchboard defines middle C as C3, meaning the first and lowest possible note is C-2.

Finally, to generate articulation sets for all patches in the library, execute:

```
go run main.go
```
