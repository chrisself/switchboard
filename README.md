# Switchboard

Switchboard is an articulation set generator for Logic Pro. Instead of mapping articulations to keyswitches using Logic's Articulation Set Editor, articulation mappings are defined using simple JSON files. Switchboard uses these articulation mappings to generate articulation sets on the fly.

## Installation

Switchboard is preloaded with articulation mappings for several sample libraries including BBC Symphony Orchestra, Spitfire Symphony Orchestra, and others. To install ready-made articulation sets, [download the latest release](https://github.com/chrisself/switchboard/releases/latest) and place the files in _User > Music > Audio Music Apps > Articulation Settings_. Browse the catalog of built-in articulation mappings [here](https://github.com/chrisself/switchboard/tree/main/catalog).

## Usage

Switchboard searches the `catalog` directory for articulation mappings. Each file corresponds with an individual patch provided by a sample library. For example, `Horns a4.json` would correspond with a "Horns a4" patch, and the generated articulation set would be written to `Horns a4.plist`.

A patch contains an array of articulation mapping objects, where `name` is the name of the articulation as seen in the plugin, and `keyswitch` is the configured note. Here's a simple example:

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

> [!IMPORTANT]
> Switchboard defines middle C as C3, meaning the first and lowest possible note is C-2.

To generate articulation sets for patches in the catalog, including your own patches, you must have a version of the Go compiler installed. See the [Download and install](https://go.dev/doc/install) guide for detailed instructions. From there, execute:

```
go run main.go
```

## Conventions

The built-in catalog is currently structured to mirror the source sample libraries as closely as possible. No effort is taken to normalize synonymous articulation names (e.g. "Long" and "Sustain") across sample libraries. Additionally, the catalog does not consolidate patches containing identical articulation mappings into one. For example, while BBC Symphony Orchestra section leader patches contain identical articulation mappings an articulation set is still generated for each one. The motivation is reduce confusion in locating the correct articulation set for a patch.

## Extensibility

Because the format is non-proprietary the articulation mappings can feasibly be used in support of other digital audio workstations, in developing control surfaces, etc.
