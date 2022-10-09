# Data Structure

## Faders

### Properties

| Name      | Value    |
|-----------|----------|
| Start     | `0x016f` |
| Item Size | 100 B    |
| Amount    | 36       |

### Fields

| Field | Offset | Values |
|-------|--------|--------|
| CHAN  | `0x00` | N/A    |
| CC    | `0x19` | 0-127  |
| MIN   | `0x32` | 0-127  |
| MAX   | `0x4b` | 0-127  |

## Dials

### Properties

| Name      | Value    |
|-----------|----------|
| Start     | `0x731b` |
| Item Size | 100 B    |
| Amount    | 32       |

### Fields

| Field | Offset | Values |
|-------|--------|--------|
| CHAN  | `0x00` | N/A    |
| CC    | `0x19` | 0-127  |
| MIN   | `0x32` | 0-127  |
| MAX   | `0x4b` | 0-127  |

## Buttons

### Properties

| Name      | Value    |
|-----------|----------|
| Start     | `0x0f7f` |
| Item Size | 750 B    |
| Amount    | 32       |

### Fields

These are the universal fields you'll find for all buttons:

| Field | Offset | Values                      |
|-------|--------|-----------------------------|
| CHAN  | `0x00` | N/A                         |
| MODE  | `0x19` | [MODE Values](#mode-values) |

If **MODE** is set to **CC**, you'll get this second set of fields:

| Field   | Offset | Values |
|---------|--------|--------|
| CC      | `0x32` | 0-127  |
| PRESS   | `0x4b` | 0-127  |
| RELEASE | `0x64` | 0-127  |

If **MODE** is set to **CC Cycle**, you'll get this second set of fields:

| Field | Offset | Values                      |
|-------|--------|-----------------------------|
| TYPE  | `0x96` | [TYPE Values](#type-values) |
| CC    | `0x7d` | 0-127                       |

You'll also get this third set of fields for **CC Cycle**:

| Field   | Offset | Values |
|---------|--------|--------|
| VALUE 1 | `0xaf` | 0-127  |
| VALUE 2 | `0xc8` | 0-127  |
| VALUE 3 | `0xe1` | 0-127  |
| VALUE 4 | `0xfa` | 0-127  |

### Values

#### MODE Values

| Byte   | Value    |
|--------|----------|
| `0x00` | CC       |
| `0x01` | CC Cycle |

#### TYPE Values

| Byte   | Value    |
|--------|----------|
| `0x01` | 2 Values |
| `0x02` | 3 Values |
| `0x03` | 4 Values |

## Pads

### Properties

| Name      | Value    |
|-----------|----------|
| Start     | `0x7f9b` |
| Item Size | 950 B    |
| Amount    | 64       |

### Fields

| Field   | Offset | Values                        |
|---------|--------|-------------------------------|
| CHAN    | `0x00` | N/A                           |
| MODE    | `0x4b` | N/A                           |
| COLOR 1 | `0x19` | [COLOR Values](#color-values) |
| COLOR 2 | `0x32` | [COLOR Values](#color-values) |

> CHAN and MODE have not been documented at the current time

### Values

#### COLOR Values

| Byte   | Value      |
|--------|------------|
| `0x00` | Off        |
| `0x01` | Chartreuse |
| `0x02` | Green      |
| `0x03` | Aquamarine |
| `0x04` | Cyan       |
| `0x05` | Azure      |
| `0x06` | Blue       |
| `0x07` | Violet     |
| `0x08` | Magenta    |
| `0x09` | Rose       |
| `0x0a` | Red        |
| `0x0b` | Orange     |
| `0x0c` | Yellow     |
| `0x0d` | White      |

