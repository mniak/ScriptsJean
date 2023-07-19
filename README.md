# ScriptsJean

## Script `createdirs`
Creates a structure of directories like the example below. The current date in the format `YY MM DD` will be used.

The directories will be created in the same directory as the binary/executable is placed.

The current version creates directories for all the days of the current month.

```
23 07 09 Frei. Fotos/
├─ 23 07 09 Pingos/
├─ 23 07 09 Flash Navega/
└─ 23 07 09 Ecobie/
   ├─ Ecobie 1/
   ├─ Ecobie 2/
   ├─ Ecobie 3/
   └─ Ecobie 4/
``` 

## Script `wadown`
Downloads images received from whatsapp into `images/phone_number`.

The program will keep running until you close it. If it closes without intervention, it means an unexpected error occurred.

Instead of the phone number, an alias can be used. You can specify aliases creating a file `wadown.config` and filling the `known_phones` as [this](cmd/wadown/wadown.config):

```
known_phones:
  "4915263011023": "Afthas Ardendeun"
  "523120799119": "Guadalupe Ramirez"
```

The phone numbers must have country and area codes. A simple way to know the correct numbers is to check the directories created inside `images/` for not-yet-known phones.

This is an [YAML file](https://en.wikipedia.org/wiki/YAML).

A database file `wadown.db` will be created. If you desire to reconfigure the application (e.g. change the account that is connected), you can simply delete it.