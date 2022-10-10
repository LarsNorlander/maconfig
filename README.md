# M-Audio Oxygen Pro 49 Preset Editor

A tool born out of frustration, this little tool is being built for the purpose of modifying preset files for the Oxygen
Pro 49.

The tool is implemented as a CLI application built in Golang, that makes use of a file with a list of actions. These
actions are meant to manipulate the preset file in a way that makes it easier to do tasks such as setting the CC for
multiple faders in one go.

## Actions

### `hello-world`

The classic hello world. Provide a name, and you'll find a personalized greeting.

#### Parameters

| Name | Values     | Example      | Description           |
|------|------------|--------------|-----------------------|
| name | any string | `'John Doe'` | the person to address |

### `set-cc`

Sets the CC values for a range of control surfaces. Given a list of values, the values will be mapped to the list of
controls specified. The following constraints apply:

* `len(values) >= len(positions)`

#### Parameters

| Name      | Values                                              | Example    | Description                              |
|-----------|-----------------------------------------------------|------------|------------------------------------------|
| control   | `fader`, `knob`                                     | `'fader'`  | the control surface type                 |
| positions | a comma separated list of a single int or int range | `'0,5-10'` | index of control surfaces                |
| values    | a comma separated list of a single int or int range | `'1,2-23'` | CC values to set across control surfaces |
